// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "../src/stBTC.sol";

contract STBTCTest is Test {
    stBTC public token;

    ValidatorRegistry public vr;

    address public admin;

    address public alice;
    address public bob;

    uint256 public constant BTC = 1e8; // sat

    event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares);

    event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id);

    function getAdminAddress(address proxy) internal view returns (address) {
        bytes32 adminSlot = vm.load(proxy, ERC1967Utils.ADMIN_SLOT);
        return address(uint160(uint256(adminSlot)));
    }

    function setUp() public {
        console.log("setUp");

        vr = new ValidatorRegistry();
        vr.setJointPublicKey(hex"4d03b19bf5cafc2c77fcd66f56c55946d7fcbc0855342a6bdf8e37b0a9986e57");

        admin = msg.sender;

        // Deploy stBTC implementation
        stBTC stBtcImplementation = new stBTC();
        bytes memory stBtcData =
            abi.encodeWithSelector(stBTC.initialize.selector, BitcoinNetworkEncoder.Network.Mainnet, vr);
        TransparentUpgradeableProxy stBtcProxy =
            new TransparentUpgradeableProxy(address(stBtcImplementation), admin, stBtcData);
        token = stBTC(address(stBtcProxy));

        alice = makeAddr("alice");
        bob = makeAddr("bob");
    }

    function testTokenInfo() public view {
        console.log("testTokenInfo");

        assertEq(token.name(), "Stroom Bitcoin");
        assertEq(token.symbol(), "stBTC");
        assertEq(token.decimals(), 8);
    }

    function testMintByOwner() public {
        uint256 mintAmount = 10 * BTC;
        bytes32 btcDepositId = keccak256(abi.encodePacked("txHash", uint256(1)));

        assertEq(token.balanceOf(alice), 0, "Alice should have no tokens initially");
        assertEq(token.totalSupply(), 0, "Total supply should be 0 initially");
        assertEq(token.totalShares(), 0, "Total shares should be 0 initially");

        token.mint(mintAmount, alice, btcDepositId);

        uint256 expectedTotalPooledBTC = mintAmount;
        uint256 expectedShares = mintAmount;
        uint256 expectedAliceBalance = mintAmount;

        assertEq(token.balanceOf(alice), expectedAliceBalance, "Alice's balance mismatch after mint");
        assertEq(token.totalSupply(), expectedTotalPooledBTC, "Total supply mismatch after mint");
        assertEq(token.totalShares(), expectedShares, "Total shares mismatch after mint");
    }

    function testMintUpdatesTotalPooledBTCAndShares() public {
        uint256 mintAmount1 = 10 * BTC;
        bytes32 btcDepositId1 = keccak256(abi.encodePacked("txHash1", uint256(1)));

        token.mint(mintAmount1, alice, btcDepositId1);

        uint256 expectedTotalPooledBTC1 = mintAmount1;
        uint256 expectedTotalShares1 = mintAmount1;

        assertEq(token.totalSupply(), expectedTotalPooledBTC1, "Total pooled BTC mismatch after first mint");
        assertEq(token.totalShares(), expectedTotalShares1, "Total shares mismatch after first mint");

        uint256 mintAmount2 = 5 * BTC;
        bytes32 btcDepositId2 = keccak256(abi.encodePacked("txHash2", uint256(2)));

        token.mint(mintAmount2, bob, btcDepositId2);

        uint256 expectedTotalPooledBTC2 = expectedTotalPooledBTC1 + mintAmount2;
        uint256 expectedTotalShares2 =
            expectedTotalShares1 + (mintAmount2 * expectedTotalShares1) / expectedTotalPooledBTC1;

        assertEq(token.totalSupply(), expectedTotalPooledBTC2, "Total pooled BTC mismatch after second mint");
        assertEq(token.totalShares(), expectedTotalShares2, "Total shares mismatch after second mint");

        uint256 expectedSharesBob = (mintAmount2 * token.totalShares()) / token.totalSupply();
        assertEq(token.getShares(bob), expectedSharesBob, "Bob's shares mismatch after second mint");

        assertEq(token.getShares(alice), expectedTotalShares1, "Alice's shares should remain unchanged");
    }

    function testMintWithDuplicateBtcDepositIdFails() public {
        uint256 mintAmount = 10 * BTC;
        bytes32 btcDepositId = keccak256(abi.encodePacked("txHash", uint256(1)));

        token.mint(mintAmount, alice, btcDepositId);

        vm.expectRevert("MINT_ALREADY_PROCESSED");
        token.mint(mintAmount, bob, btcDepositId);

        assertEq(token.getShares(alice), mintAmount, "Alice's shares should remain unchanged");
        assertEq(token.getShares(bob), 0, "Bob should have no shares");
        assertEq(token.totalSupply(), mintAmount, "Total supply should remain unchanged");
    }

    function testMintWithValidSignature() public {
        uint256 mintAmountAlice = 10 * BTC;
        bytes32 btcDepositIdAlice = keccak256(abi.encodePacked("txHashAlice", uint256(1)));
        stBTC.MintInvoice memory invoiceAlice =
            stBTC.MintInvoice({btcDepositId: btcDepositIdAlice, recipient: alice, amount: mintAmountAlice});
        bytes32 invoiceAliceHash = token.getMintInvoiceHash(invoiceAlice);
        console.log("invoice Alice hash");
        emit log_bytes32(invoiceAliceHash);

        uint256 mintAmountBob = 15 * BTC;
        bytes32 btcDepositIdBob = keccak256(abi.encodePacked("txHashBob", uint256(2)));
        stBTC.MintInvoice memory invoiceBob =
            stBTC.MintInvoice({btcDepositId: btcDepositIdBob, recipient: bob, amount: mintAmountBob});
        bytes32 invoiceBobHash = token.getMintInvoiceHash(invoiceBob);
        console.log("invoice Bob hash");
        emit log_bytes32(invoiceBobHash);

        bytes memory signatureAlice =
            hex"564e41f1f0c2cf5026f6e7428cdfd7f735445955baa81597db44a3c8eb4919746a56a396664a08da75c1ad1329c00de8fa1f5311d6ba3609289725dfe1df52a5";
        bytes memory signatureBob =
            hex"7aab8176fdf8216857125755d9f3c346dc3b29b0e6d2b4e49ae9e2661d3f83857213802f1659c7f5af61704311fff30102725caa028628b1606218ebff284c20";

        token.mint(invoiceAlice, signatureAlice);

        assertEq(token.balanceOf(alice), mintAmountAlice, "Alice's balance mismatch after mint");
        assertEq(token.totalSupply(), mintAmountAlice, "Total supply mismatch after mint for Alice");

        token.mint(invoiceBob, signatureBob);

        uint256 expectedTotalSupply = mintAmountAlice + mintAmountBob;
        assertEq(token.balanceOf(bob), mintAmountBob, "Bob's balance mismatch after mint");
        assertEq(token.totalSupply(), expectedTotalSupply, "Total supply mismatch after mint for Bob");
    }

    function testMintWithInvalidSignatureFails() public {
        uint256 mintAmountAlice = 10 * BTC;
        bytes32 btcDepositIdAlice = keccak256(abi.encodePacked("txHashAlice", uint256(1)));
        stBTC.MintInvoice memory invoiceAlice =
            stBTC.MintInvoice({btcDepositId: btcDepositIdAlice, recipient: alice, amount: mintAmountAlice});

        bytes memory invalidSignature =
            hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef123456789abcdef1";

        vm.expectRevert("ValidatorMessageReceiver: INVALID_SIGNATURE");
        token.mint(invoiceAlice, invalidSignature);

        assertEq(token.balanceOf(alice), 0, "Alice's balance should remain 0");
        assertEq(token.totalSupply(), 0, "Total supply should remain 0");
    }

    function testMintWithExistingBtcDepositIdFails() public {
        uint256 mintAmountAlice = 10 * BTC;
        bytes32 btcDepositIdAlice = keccak256(abi.encodePacked("txHashAlice", uint256(1)));
        stBTC.MintInvoice memory invoiceAlice =
            stBTC.MintInvoice({btcDepositId: btcDepositIdAlice, recipient: alice, amount: mintAmountAlice});

        bytes memory validSignatureAlice =
            hex"564e41f1f0c2cf5026f6e7428cdfd7f735445955baa81597db44a3c8eb4919746a56a396664a08da75c1ad1329c00de8fa1f5311d6ba3609289725dfe1df52a5";

        token.mint(invoiceAlice, validSignatureAlice);

        assertEq(token.balanceOf(alice), mintAmountAlice, "Alice's balance mismatch after first mint");

        vm.expectRevert("MINT_ALREADY_PROCESSED");
        token.mint(invoiceAlice, validSignatureAlice);

        assertEq(token.balanceOf(alice), mintAmountAlice, "Alice's balance should remain unchanged");
        assertEq(token.totalSupply(), mintAmountAlice, "Total supply should remain unchanged");
    }

    function testMintRewardsUpdatesTotalPooledBTCAndEmitsEvent() public {
        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 rewardAmount = 5 * BTC;
        uint256 initialNonce = token.totalSupplyUpdateNonce();

        vm.expectEmit(true, true, true, true);
        emit TotalSupplyUpdatedEvent(initialNonce, initialTotalPooledBTC + rewardAmount, token.totalShares());
        token.mintRewards(initialNonce, rewardAmount);

        uint256 updatedTotalPooledBTC = token.totalSupply();
        assertEq(
            updatedTotalPooledBTC, initialTotalPooledBTC + rewardAmount, "Total pooled BTC mismatch after mintRewards"
        );

        assertEq(token.totalSupplyUpdateNonce(), initialNonce + 1, "Nonce mismatch after mintRewards");
    }

    function testMintRewardsWithValidSignature() public {
        vr.setJointPublicKey(hex"10b4a24b61083d0ff552a342d844b3203120e349884d8fa9651cae849d02d920");
        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 rewardAmount = 10 * BTC;
        uint256 initialNonce = token.totalSupplyUpdateNonce();

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(initialNonce, rewardAmount);
        console.log("Total Supply Update Hash");
        emit log_bytes32(totalSupplyUpdateHash);

        bytes memory validSignature =
            hex"4b513184e07f573d12980d4f3925179dac3671bae1af9df3e3272fbf02d6fc6bb1e3f1776997032754d4972c1efb603fc24e478cb09b33d0c68d54c7978b31c8";

        vm.expectEmit(true, true, true, true);
        emit TotalSupplyUpdatedEvent(initialNonce, initialTotalPooledBTC + rewardAmount, token.totalShares());
        token.mintRewards(initialNonce, rewardAmount, validSignature);

        assertEq(
            token.totalSupply(),
            initialTotalPooledBTC + rewardAmount,
            "Total pooled BTC mismatch after mintRewards with valid signature"
        );

        assertEq(
            token.totalSupplyUpdateNonce(), initialNonce + 1, "Nonce mismatch after mintRewards with valid signature"
        );
    }

    function testMintRewardsWithInvalidSignatureFails() public {
        uint256 rewardAmount = 10 * BTC;
        uint256 invalidNonce = token.totalSupplyUpdateNonce();

        bytes memory invalidSignature =
            hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef123456789abcdef1";

        vm.expectRevert("ValidatorMessageReceiver: INVALID_SIGNATURE");
        token.mintRewards(invalidNonce, rewardAmount, invalidSignature);

        assertEq(
            token.totalSupply(), 0, "Total pooled BTC should remain unchanged after mintRewards with invalid signature"
        );

        assertEq(
            token.totalSupplyUpdateNonce(),
            invalidNonce,
            "Nonce should remain unchanged after mintRewards with invalid signature"
        );
    }

    function testMintRewardsWithDuplicateRewardIdFails() public {
        vr.setJointPublicKey(hex"10b4a24b61083d0ff552a342d844b3203120e349884d8fa9651cae849d02d920");

        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 rewardAmount = 10 * BTC;
        uint256 initialNonce = token.totalSupplyUpdateNonce();

        bytes memory validSignature =
            hex"4b513184e07f573d12980d4f3925179dac3671bae1af9df3e3272fbf02d6fc6bb1e3f1776997032754d4972c1efb603fc24e478cb09b33d0c68d54c7978b31c8";

        token.mintRewards(initialNonce, rewardAmount, validSignature);

        assertEq(
            token.totalSupply(),
            initialTotalPooledBTC + rewardAmount,
            "Total pooled BTC mismatch after first mintRewards"
        );

        vm.expectRevert("Invalid update total supply nonce");
        token.mintRewards(initialNonce, rewardAmount, validSignature);

        assertEq(
            token.totalSupply(),
            initialTotalPooledBTC + rewardAmount,
            "Total pooled BTC should remain unchanged after duplicate rewardId"
        );

        assertEq(
            token.totalSupplyUpdateNonce(), initialNonce + 1, "Nonce should remain unchanged after duplicate rewardId"
        );
    }

    function testRedeemValidAmount() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 5 * BTC;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");
        assertEq(token.totalSupply(), mintAmount, "Initial total supply mismatch");

        vm.startPrank(alice);
        vm.expectEmit(true, true, true, true);
        emit RedeemBtcEvent(alice, validBTCAddress, redeemAmount, token.redeemCounter() + 1);
        token.redeem(redeemAmount, validBTCAddress);
        vm.stopPrank();

        assertEq(token.balanceOf(alice), mintAmount - redeemAmount, "Balance mismatch after redeem");

        assertEq(token.totalSupply(), mintAmount - redeemAmount, "Total supply mismatch after redeem");

        assertEq(token.redeemCounter(), 1, "Redeem counter mismatch after redeem");
    }

    function testRedeemBelowMinWithdrawAmountFails() public {
        uint256 mintAmount = 10 * BTC;
        uint256 invalidRedeemAmount = token.minWithdrawAmount() - 1;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert("The sent value must be greater or equal to min withdraw amount");
        token.redeem(invalidRedeemAmount, validBTCAddress);

        assertEq(token.balanceOf(alice), mintAmount, "Balance should remain unchanged after failed redeem");

        assertEq(token.totalSupply(), mintAmount, "Total supply should remain unchanged after failed redeem");
    }

    function testRedeemWithInvalidBTCAddressFails() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 1 * BTC;
        string memory invalidBTCAddress = "InvalidBTCAddress123";

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert("The sent BTC address is not valid");
        token.redeem(redeemAmount, invalidBTCAddress);

        assertEq(token.balanceOf(alice), mintAmount, "Balance should remain unchanged after failed redeem");

        assertEq(token.totalSupply(), mintAmount, "Total supply should remain unchanged after failed redeem");
    }

    function testRedeemReducesTotalPooledBTCAndShares() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 5 * BTC;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 initialShares = token.getShares(alice);
        assertEq(initialTotalPooledBTC, mintAmount, "Initial totalPooledBTC mismatch");
        assertGt(initialShares, 0, "Initial shares should be greater than 0");

        vm.prank(alice);
        token.redeem(redeemAmount, validBTCAddress);

        assertEq(token.totalSupply(), initialTotalPooledBTC - redeemAmount, "Total pooled BTC mismatch after redeem");

        uint256 expectedSharesReduction = (redeemAmount * initialShares) / initialTotalPooledBTC;
        assertEq(token.getShares(alice), initialShares - expectedSharesReduction, "Shares mismatch after redeem");
    }

    function testTransferTokensSuccessfully() public {
        uint256 mintAmount = 10 * BTC;
        uint256 transferAmount = 5 * BTC;

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch for Alice");
        assertEq(token.balanceOf(bob), 0, "Initial balance mismatch for Bob");
        uint256 aliceSharesBefore = token.getShares(alice);
        uint256 bobSharesBefore = token.getShares(bob);

        vm.prank(alice);
        token.transfer(bob, transferAmount);

        assertEq(token.balanceOf(alice), mintAmount - transferAmount, "Alice's balance mismatch after transfer");

        assertEq(token.balanceOf(bob), transferAmount, "Bob's balance mismatch after transfer");

        uint256 expectedAliceShares = aliceSharesBefore - (transferAmount * aliceSharesBefore) / mintAmount;
        assertEq(token.getShares(alice), expectedAliceShares, "Alice's shares mismatch after transfer");

        uint256 expectedBobShares = bobSharesBefore + (transferAmount * aliceSharesBefore) / mintAmount;
        assertEq(token.getShares(bob), expectedBobShares, "Bob's shares mismatch after transfer");
    }

    function testTransferMoreThanBalanceShouldRevert() public {
        uint256 mintAmount = 10 * BTC;
        uint256 transferAmount = 15 * BTC;

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        assertEq(token.balanceOf(alice), mintAmount, "Alice's initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert("INSUFFICIENT_BALANCE");
        token.transfer(bob, transferAmount);

        assertEq(token.balanceOf(alice), mintAmount, "Alice's balance should remain unchanged");
        assertEq(token.balanceOf(bob), 0, "Bob's balance should remain 0");
    }

    function testTransferFrom() public {
        uint256 mintAmount = 10 * BTC;
        uint256 transferAmount = 5 * BTC;

        token.mint(mintAmount, alice, keccak256(abi.encodePacked("deposit1")));

        uint256 initialBalanceAlice = token.balanceOf(alice);
        assertEq(initialBalanceAlice, mintAmount, "Initial balance of Alice incorrect");

        vm.prank(alice);
        token.approve(bob, transferAmount);

        uint256 allowance = token.allowance(alice, bob);
        assertEq(allowance, transferAmount, "Allowance for Bob incorrect");

        vm.prank(bob);
        token.transferFrom(alice, bob, transferAmount);

        uint256 finalBalanceAlice = token.balanceOf(alice);
        assertEq(finalBalanceAlice, mintAmount - transferAmount, "Final balance of Alice incorrect");

        uint256 finalBalanceBob = token.balanceOf(bob);
        assertEq(finalBalanceBob, transferAmount, "Final balance of Bob incorrect");

        uint256 finalAllowance = token.allowance(alice, bob);
        assertEq(finalAllowance, 0, "Final allowance for Bob incorrect");
    }

    function testBalanceOfandTotalSupply() public {
        uint256 mintAmountAlice = 10 * BTC;
        uint256 mintAmountBob = 5 * BTC;

        token.mint(mintAmountAlice, alice, keccak256(abi.encodePacked("deposit1")));
        token.mint(mintAmountBob, bob, keccak256(abi.encodePacked("deposit2")));

        uint256 totalPooledBTC = token.totalSupply();
        uint256 totalShares = token.totalShares();

        uint256 expectedTotalPooledBTC = mintAmountAlice + mintAmountBob;
        uint256 expectedShares = (mintAmountAlice + mintAmountBob);

        assertEq(totalPooledBTC, expectedTotalPooledBTC, "Total pooled BTC mismatch");
        assertGt(totalShares, 0, "Total shares should be greater than 0");

        uint256 expectedBalanceAlice = (token.getShares(alice) * totalPooledBTC) / totalShares;
        uint256 expectedBalanceBob = (token.getShares(bob) * totalPooledBTC) / totalShares;

        assertEq(token.balanceOf(alice), expectedBalanceAlice, "Alice's balance mismatch");
        assertEq(token.balanceOf(bob), expectedBalanceBob, "Bob's balance mismatch");
        assertEq(token.totalSupply(), expectedTotalPooledBTC, "totalSupply does not match _totalPooledBTC");
        assertEq(token.totalShares(), expectedShares, "totalShares does not match expected value");
    }

    function testRedeemAfterMintReward() public {
        uint256 mintAmountAlice = 10 * BTC;
        uint256 mintRewardAmount = 5 * BTC;
        uint256 redeemAmount = 5 * BTC;

        token.mint(mintAmountAlice, alice, keccak256(abi.encodePacked("deposit1")));

        uint256 initialBalanceAlice = token.balanceOf(alice);
        assertEq(initialBalanceAlice, mintAmountAlice, "Alice's initial balance mismatch");

        token.mintRewards(0, mintRewardAmount);

        uint256 updatedBalanceAlice = token.balanceOf(alice);
        uint256 expectedUpdatedBalanceAlice = (mintAmountAlice * token.totalSupply()) / token.totalShares();
        assertEq(updatedBalanceAlice, expectedUpdatedBalanceAlice, "Alice's balance mismatch after mintReward");

        uint256 expectedTotalPooledBTC = token.totalSupply() - redeemAmount;

        vm.prank(alice);
        token.redeem(redeemAmount, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        uint256 remainingBalanceAlice = token.balanceOf(alice);
        uint256 expectedRemainingBalanceAlice = updatedBalanceAlice - redeemAmount;
        assertEq(
            remainingBalanceAlice, expectedRemainingBalanceAlice, "Alice's remaining balance mismatch after redeem"
        );

        assertEq(token.totalSupply(), expectedTotalPooledBTC, "Total pooled BTC mismatch after redeem");
    }

    function testPauseUnpause() public {
        vm.prank(alice);
        vm.expectRevert();
        token.pause();

        token.pause();

        vm.expectRevert();
        token.mint(10 * BTC, alice, keccak256(abi.encodePacked("deposit1")));

        vm.prank(alice);
        vm.expectRevert();
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        vm.expectRevert();
        token.mintRewards(1, 5 * BTC);

        token.unpause();

        token.mint(10 * BTC, alice, keccak256(abi.encodePacked("deposit1")));
        uint256 balanceAfterMint = token.balanceOf(alice);
        assertEq(balanceAfterMint, 10 * BTC, "Mint should work after unpause");

        vm.prank(alice);
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");
        uint256 balanceAfterRedeem = token.balanceOf(alice);
        assertEq(balanceAfterRedeem, 5 * BTC, "Redeem should work after unpause");
    }

    function testEdgeCases() public {
        vm.expectRevert("MINT_AMOUNT_ZERO");
        token.mint(0, alice, keccak256(abi.encodePacked("deposit1")));

        vm.prank(alice);
        vm.expectRevert("The sent value must be greater or equal to min withdraw amount");
        token.redeem(0, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        token.mint(10 * BTC, alice, keccak256(abi.encodePacked("deposit1")));
        assertEq(token.totalShares(), 10 * BTC, "Shares should equal minted amount in the initial state");
        assertEq(token.totalSupply(), 10 * BTC, "Total pooled BTC should equal minted amount");

        vm.prank(alice);
        assertEq(token.balanceOf(alice), 10 * BTC, "Balance should match initial mint");

        vm.expectRevert();
        token.mint(10 * BTC, address(0), keccak256(abi.encodePacked("deposit2")));

        vm.prank(alice);
        vm.expectRevert("The sent BTC address is not valid");
        token.redeem(5 * BTC, "");

        vm.prank(alice);
        vm.expectRevert();
        token.transfer(address(0), 5 * BTC);

        vm.prank(alice);
        vm.expectRevert();
        token.transferFrom(address(0), bob, 5 * BTC);

        vm.prank(alice);
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        assertEq(token.totalShares(), 5 * BTC, "Shares should decrease after redeem");
        assertEq(token.totalSupply(), 5 * BTC, "Total pooled BTC should decrease after redeem");
    }

    function testFuzzMint(uint256 _amount, bytes32 _btcDepositId) public {
        vm.assume(_amount > 0 && _amount < 21_000_000 * BTC);

        address recipient = makeAddr("recipient");

        token.mint(_amount, recipient, _btcDepositId);

        assertEq(token.balanceOf(recipient), _amount, "Recipient balance should match minted amount");
        assertEq(token.totalSupply(), _amount, "Total supply should match minted amount");
    }

    function testFuzzRedeem(uint256 _amount) public {
        vm.assume(_amount > token.minWithdrawAmount() && _amount < 21_000_000 * BTC);

        address recipient = makeAddr("recipient");
        bytes32 depositId = keccak256(abi.encodePacked("deposit"));
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        token.mint(_amount, recipient, depositId);

        vm.prank(recipient);
        token.redeem(_amount, validBTCAddress);

        assertEq(token.balanceOf(recipient), 0, "Recipient balance should be 0 after redeem");
        assertEq(token.totalSupply(), 0, "Total supply should be 0 after redeem");
    }

    function testFuzzTransfer(address _to, uint256 _value) public {
        vm.assume(_value > 0 && _value < 21_000_000 * BTC);

        address sender = makeAddr("sender");
        bytes32 depositId = keccak256(abi.encodePacked("deposit"));

        vm.assume(_to != address(0));

        token.mint(_value, sender, depositId);

        vm.prank(sender);
        token.transfer(_to, _value);

        assertEq(token.balanceOf(_to), _value, "Recipient balance should match transferred amount");
        assertEq(token.balanceOf(sender), 0, "Sender balance should be 0 after transfer");
    }

    function testFuzzMintTransferRedeem(uint256 mintAmount) public {
        vm.assume(mintAmount > token.minWithdrawAmount() * 2 && mintAmount < 21_000_000 * BTC);

        bytes32 depositId = keccak256(abi.encodePacked("aliceDeposit"));
        token.mint(mintAmount, alice, depositId);

        uint256 aliceBalance = token.balanceOf(alice);
        assertEq(aliceBalance, mintAmount, "Alice should receive the minted amount");

        uint256 transferAmount = mintAmount / 2;
        vm.prank(alice);
        token.transfer(bob, transferAmount);

        assertEq(token.balanceOf(alice), mintAmount - transferAmount, "Alice should retain half the tokens");
        assertEq(token.balanceOf(bob), transferAmount, "Bob should receive half the tokens");

        uint256 aliceRemainingBalance = token.balanceOf(alice);
        vm.prank(alice);
        token.redeem(aliceRemainingBalance, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        assertEq(token.balanceOf(alice), 0, "Alice should have zero balance after redeem");

        uint256 bobBalance = token.balanceOf(bob);
        uint256 excessRedeemAmount = bobBalance + 1;
        vm.expectRevert("INSUFFICIENT_BALANCE");
        vm.prank(bob);
        token.redeem(excessRedeemAmount, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        vm.prank(bob);
        token.redeem(bobBalance, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        assertEq(token.balanceOf(bob), 0, "Bob should have zero balance after redeem");

        assertEq(token.totalSupply(), 0, "Total supply should be zero after all redemptions");
    }
}

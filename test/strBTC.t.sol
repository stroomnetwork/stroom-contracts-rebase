// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "../src/strBTC.sol";

contract STRBTCTest is Test {
    strBTC public token;

    ValidatorRegistry public vr;

    address public admin;
    address public pauser;

    address public alice;
    address public bob;

    uint256 public constant BTC = 1e8; // sat

    uint256 public constant INITIAL_SUPPLY = 1000 * BTC;

    event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares);

    event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id);

    function setUp() public {
        console.log("setUp");

        vm.warp(1_729_690_309);

        vr = new ValidatorRegistry();
        vr.setJointPublicKey(hex"9627e95c7c43a6550b0bcc005bbd85de78a1e17285c9acae2349292e78b21c0f");

        admin = msg.sender;
        pauser = makeAddr("Pauser");

        // Deploy strBTC implementation
        strBTC strBtcImplementation = new strBTC();
        bytes memory strBtcData =
            abi.encodeWithSelector(strBTC.initialize.selector, BitcoinNetworkEncoder.Network.Mainnet, vr, admin, pauser);
        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), admin, strBtcData);
        token = strBTC(address(strBtcProxy));

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("initial supply"), recipient: admin, amount: INITIAL_SUPPLY});

        bytes memory signature =
            hex"3c3a9e37d7a8bc3d72be70758e070b31afb248161578ddbc3267c1f7e156ef7beaa20477f82454b5cf23da2d791497c5101cb18c5e48cf1c57054958bdb7662f";

        token.mint(invoice, signature);

        alice = makeAddr("alice");
        bob = makeAddr("bob");
    }

    function testTokenInfo() public view {
        console.log("testTokenInfo");

        assertEq(token.name(), "Stroom Bitcoin");
        assertEq(token.symbol(), "strBTC");
        assertEq(token.decimals(), 8);
    }

    function testMint() public {
        uint256 mintAmount = 10 * BTC;
        bytes32 btcDepositId = keccak256(abi.encodePacked("txHash", uint256(1)));

        assertEq(token.balanceOf(alice), 0, "Alice should have no tokens initially");
        assertEq(token.totalSupply(), INITIAL_SUPPLY, "Total supply should be INITIAL_SUPPLY");
        assertEq(token.totalShares(), INITIAL_SUPPLY, "Total shares should be INITIAL_SUPPLY");

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: btcDepositId, recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"589e1fbad605039213fb205db0e92bef7bfdd27a460c8323735b2e5729cd40e92c5be606fa719d210dabc5c5e2b49788ccb944822cfda8ed003b5b42d4d56450";

        token.mint(invoice, signature);

        uint256 expectedTotalPooledBTC = mintAmount + INITIAL_SUPPLY;
        uint256 expectedShares = mintAmount + INITIAL_SUPPLY;
        uint256 expectedAliceBalance = mintAmount;

        assertEq(token.balanceOf(alice), expectedAliceBalance, "Alice's balance mismatch after mint");
        assertEq(token.totalSupply(), expectedTotalPooledBTC, "Total supply mismatch after mint");
        assertEq(token.totalShares(), expectedShares, "Total shares mismatch after mint");
    }

    function testMintUpdatesTotalPooledBTCAndShares() public {
        uint256 mintAmount1 = 10 * BTC;
        bytes32 btcDepositId1 = keccak256(abi.encodePacked("txHash1", uint256(1)));

        strBTC.MintInvoice memory invoice1 =
            strBTC.MintInvoice({btcDepositId: btcDepositId1, recipient: alice, amount: mintAmount1});
        bytes32 invoiceHash1 = token.getMintInvoiceHash(invoice1);
        console.logBytes32(invoiceHash1);
        bytes memory signature1 =
            hex"47be24aa0a8a343d4b3c90b838a45516611c46cafed4cffe409781bd6d56135faed8e6779db6e881cd31cafd8a07efccd47c9a19572716e895fa69c0ef61c6ab";

        token.mint(invoice1, signature1);

        uint256 expectedTotalPooledBTC1 = mintAmount1 + INITIAL_SUPPLY;
        uint256 expectedTotalShares1 = mintAmount1 + INITIAL_SUPPLY;

        assertEq(token.totalSupply(), expectedTotalPooledBTC1, "Total pooled BTC mismatch after first mint");
        assertEq(token.totalShares(), expectedTotalShares1, "Total shares mismatch after first mint");

        uint256 mintAmount2 = 5 * BTC;
        bytes32 btcDepositId2 = keccak256(abi.encodePacked("txHash2", uint256(2)));

        strBTC.MintInvoice memory invoice2 =
            strBTC.MintInvoice({btcDepositId: btcDepositId2, recipient: bob, amount: mintAmount2});
        bytes32 invoiceHash2 = token.getMintInvoiceHash(invoice2);
        console.logBytes32(invoiceHash2);
        bytes memory signature2 =
            hex"6147dd3c046e5a88b60b272124b871c052527bd0a9c261fc09d776e223c148d94350c0c3884cfdc07e5e1a0dd770c6314e98963900966cb9398104db7d6f9a14";

        token.mint(invoice2, signature2);

        uint256 expectedTotalPooledBTC2 = expectedTotalPooledBTC1 + mintAmount2;
        uint256 expectedTotalShares2 =
            expectedTotalShares1 + (mintAmount2 * expectedTotalShares1) / expectedTotalPooledBTC1;

        assertEq(token.totalSupply(), expectedTotalPooledBTC2, "Total pooled BTC mismatch after second mint");
        assertEq(token.totalShares(), expectedTotalShares2, "Total shares mismatch after second mint");

        uint256 expectedSharesBob = (mintAmount2 * token.totalShares()) / token.totalSupply();
        assertEq(token.getShares(bob), expectedSharesBob, "Bob's shares mismatch after second mint");

        // coment because of stack too deep
        // assertEq(token.getShares(alice), mintAmount1, "Alice's shares should remain unchanged");
    }

    function testMintWithDuplicateBtcDepositIdFails() public {
        uint256 mintAmount = 10 * BTC;
        bytes32 btcDepositId = keccak256(abi.encodePacked("txHash", uint256(1)));

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: btcDepositId, recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"589e1fbad605039213fb205db0e92bef7bfdd27a460c8323735b2e5729cd40e92c5be606fa719d210dabc5c5e2b49788ccb944822cfda8ed003b5b42d4d56450";

        token.mint(invoice, signature);

        vm.expectRevert(strBTC.MintAlreadyProcessed.selector);
        token.mint(invoice, signature);

        assertEq(token.getShares(alice), mintAmount, "Alice's shares should remain unchanged");
        assertEq(token.totalSupply(), mintAmount + INITIAL_SUPPLY, "Total supply should remain unchanged");
    }

    function testMintWithValidSignature() public {
        uint256 mintAmountAlice = 10 * BTC;
        bytes32 btcDepositIdAlice = keccak256(abi.encodePacked("txHashAlice", uint256(1)));
        strBTC.MintInvoice memory invoiceAlice =
            strBTC.MintInvoice({btcDepositId: btcDepositIdAlice, recipient: alice, amount: mintAmountAlice});
        bytes32 invoiceAliceHash = token.getMintInvoiceHash(invoiceAlice);
        console.logBytes32(invoiceAliceHash);

        uint256 mintAmountBob = 15 * BTC;
        bytes32 btcDepositIdBob = keccak256(abi.encodePacked("txHashBob", uint256(2)));
        strBTC.MintInvoice memory invoiceBob =
            strBTC.MintInvoice({btcDepositId: btcDepositIdBob, recipient: bob, amount: mintAmountBob});
        bytes32 invoiceBobHash = token.getMintInvoiceHash(invoiceBob);
        console.logBytes32(invoiceBobHash);

        bytes memory signatureAlice =
            hex"4a40ed483914553a4e84987169cfc27052e2f2b48fdce21b449030b456a2ee3e03408575bc6f2a3252fb6c341a2c890d5712e91357f47c0688a4f0e7e0044303";
        bytes memory signatureBob =
            hex"f29ec11f0f0d8b95d0044f5d778f22fde9d32e1daf12a8ddfb8e0fae35e6e17391d7f5665c9275bff3b31468bf5f14efcd8d0f64443ce2a3f0778b5340b3ed03";

        token.mint(invoiceAlice, signatureAlice);

        assertEq(token.balanceOf(alice), mintAmountAlice, "Alice's balance mismatch after mint");
        assertEq(token.totalSupply(), mintAmountAlice + INITIAL_SUPPLY, "Total supply mismatch after mint for Alice");

        token.mint(invoiceBob, signatureBob);

        uint256 expectedTotalSupply = mintAmountAlice + mintAmountBob + INITIAL_SUPPLY;
        assertEq(token.balanceOf(bob), mintAmountBob, "Bob's balance mismatch after mint");
        assertEq(token.totalSupply(), expectedTotalSupply, "Total supply mismatch after mint for Bob");
    }

    function testMintWithInvalidSignatureFails() public {
        uint256 mintAmountAlice = 10 * BTC;
        bytes32 btcDepositIdAlice = keccak256(abi.encodePacked("txHashAlice", uint256(1)));
        strBTC.MintInvoice memory invoiceAlice =
            strBTC.MintInvoice({btcDepositId: btcDepositIdAlice, recipient: alice, amount: mintAmountAlice});

        bytes memory invalidSignature =
            hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef123456789abcdef1";

        vm.expectRevert(ValidatorMessageReceiver.InvalidValidatorSignature.selector);
        token.mint(invoiceAlice, invalidSignature);

        assertEq(token.balanceOf(alice), 0, "Alice's balance should remain 0");
        assertEq(token.totalSupply(), INITIAL_SUPPLY, "Total supply should remain INITIAL_SUPPLY");
    }

    function testMintRewardsUpdatesTotalPooledBTC() public {
        uint256 initialTotalPooledBTC = token.totalSupply();
        console.log("initialTotalPooledBTC", initialTotalPooledBTC);
        uint256 rewardAmount = 5 * BTC;
        uint256 initialNonce = token.totalSupplyUpdateNonce();

        vm.expectEmit(true, true, true, true);
        emit TotalSupplyUpdatedEvent(initialNonce, initialTotalPooledBTC + rewardAmount, token.totalShares());

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(initialNonce, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory signature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        token.mintRewards(initialNonce, rewardAmount, signature);

        uint256 updatedTotalPooledBTC = token.totalSupply();
        assertEq(
            updatedTotalPooledBTC, initialTotalPooledBTC + rewardAmount, "Total pooled BTC mismatch after mintRewards"
        );

        assertEq(token.totalSupplyUpdateNonce(), initialNonce + 1, "Nonce mismatch after mintRewards");
    }

    function testMintRewardsWithInvalidSignatureFails() public {
        uint256 rewardAmount = 10 * BTC;
        uint256 validNonce = token.totalSupplyUpdateNonce();

        bytes memory invalidSignature =
            hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef123456789abcdef1";

        vm.expectRevert(ValidatorMessageReceiver.InvalidValidatorSignature.selector);
        token.mintRewards(validNonce, rewardAmount, invalidSignature);

        assertEq(
            token.totalSupply(),
            INITIAL_SUPPLY,
            "Total pooled BTC should remain unchanged after mintRewards with invalid signature"
        );

        assertEq(
            token.totalSupplyUpdateNonce(),
            validNonce,
            "Nonce should remain unchanged after mintRewards with invalid signature"
        );
    }

    function testMintRewardsWithDuplicateRewardIdFails() public {
        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 rewardAmount = 10 * BTC;
        uint256 initialNonce = token.totalSupplyUpdateNonce();

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(initialNonce, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);

        bytes memory validSignature =
            hex"53cbf297ec5d591b9f224deb01b0decc0148688cd0826ef3393f7937c94bb66219433284f9968c0b5814627d9713e40b012d04ea26ff78f776752d4e4648bbac";

        token.mintRewards(initialNonce, rewardAmount, validSignature);

        assertEq(
            token.totalSupply(),
            initialTotalPooledBTC + rewardAmount,
            "Total pooled BTC mismatch after first mintRewards"
        );

        vm.expectRevert(strBTC.InvalidTotalSupplyNonce.selector);
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

    function testMintRewardsWithInvalidNonce() public {
        uint256 rewardAmount = 5 * BTC;
        uint256 invalidNonce = 5;

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(invalidNonce, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory invalidSignature =
            hex"7cda4fa045056a574e22f07ef655e74a1640d2afbfb535fb5baf0acd8585085a77d63a352768de3cef51da248944af1a7f809f693b4838a21071ecb2bdc30abf";
        vm.expectRevert(strBTC.InvalidTotalSupplyNonce.selector);
        token.mintRewards(invalidNonce, rewardAmount, invalidSignature);
    }

    function testRedeemValidAmount() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 5 * BTC;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");
        assertEq(token.totalSupply(), mintAmount + INITIAL_SUPPLY, "Initial total supply mismatch");

        vm.startPrank(alice);
        vm.expectEmit(true, true, true, true);
        emit RedeemBtcEvent(alice, validBTCAddress, redeemAmount, token.redeemCounter() + 1);
        token.redeem(redeemAmount, validBTCAddress);
        vm.stopPrank();

        assertEq(token.balanceOf(alice), mintAmount - redeemAmount, "Balance mismatch after redeem");

        assertEq(token.totalSupply(), mintAmount - redeemAmount + INITIAL_SUPPLY, "Total supply mismatch after redeem");

        assertEq(token.redeemCounter(), 1, "Redeem counter mismatch after redeem");
    }

    function testRedeemBelowMinWithdrawAmountFails() public {
        uint256 mintAmount = 10 * BTC;
        uint256 invalidRedeemAmount = token.minWithdrawAmount() - 1;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        token.mint(invoice, signature);

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert(strBTC.AmountBelowMinWithdraw.selector);
        token.redeem(invalidRedeemAmount, validBTCAddress);

        assertEq(token.balanceOf(alice), mintAmount, "Balance should remain unchanged after failed redeem");

        assertEq(
            token.totalSupply(), mintAmount + INITIAL_SUPPLY, "Total supply should remain unchanged after failed redeem"
        );
    }

    function testRedeemWithInvalidBTCAddressFails() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 1 * BTC;
        string memory invalidBTCAddress = "InvalidBTCAddress123";

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        token.mint(invoice, signature);

        assertEq(token.balanceOf(alice), mintAmount, "Initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert(strBTC.InvalidBTCAddress.selector);
        token.redeem(redeemAmount, invalidBTCAddress);

        assertEq(token.balanceOf(alice), mintAmount, "Balance should remain unchanged after failed redeem");

        assertEq(
            token.totalSupply(), mintAmount + INITIAL_SUPPLY, "Total supply should remain unchanged after failed redeem"
        );
    }

    function testRedeemReducesTotalPooledBTCAndShares() public {
        uint256 mintAmount = 10 * BTC;
        uint256 redeemAmount = 5 * BTC;
        string memory validBTCAddress = "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9";

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

        uint256 initialTotalPooledBTC = token.totalSupply();
        uint256 initialShares = token.getShares(alice);
        uint256 aliceBalance = token.balanceOf(alice);

        assertEq(initialTotalPooledBTC, mintAmount + INITIAL_SUPPLY, "Initial totalPooledBTC mismatch");
        assertGt(initialShares, 0, "Initial shares should be greater than 0");

        vm.prank(alice);
        token.redeem(redeemAmount, validBTCAddress);

        assertEq(token.totalSupply(), initialTotalPooledBTC - redeemAmount, "Total pooled BTC mismatch after redeem");

        uint256 expectedRemainingShares = initialShares * (aliceBalance - redeemAmount) / aliceBalance;
        assertEq(token.getShares(alice), expectedRemainingShares, "Shares mismatch after redeem");
    }

    function testTransferTokensSuccessfully() public {
        uint256 mintAmount = 10 * BTC;
        uint256 transferAmount = 5 * BTC;

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

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

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        token.mint(invoice, signature);

        assertEq(token.balanceOf(alice), mintAmount, "Alice's initial balance mismatch");

        vm.prank(alice);
        vm.expectRevert(strBTC.InsufficientBalance.selector);
        token.transfer(bob, transferAmount);

        assertEq(token.balanceOf(alice), mintAmount, "Alice's balance should remain unchanged");
        assertEq(token.balanceOf(bob), 0, "Bob's balance should remain 0");
    }

    function testTransferFrom() public {
        uint256 mintAmount = 10 * BTC;
        uint256 transferAmount = 5 * BTC;

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

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

        strBTC.MintInvoice memory invoiceAlice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmountAlice});
        bytes32 invoiceHashAlice = token.getMintInvoiceHash(invoiceAlice);
        console.logBytes32(invoiceHashAlice);
        bytes memory signatureAlice =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoiceAlice, signatureAlice);

        strBTC.MintInvoice memory invoiceBob =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit2"), recipient: bob, amount: mintAmountBob});
        bytes32 invoiceHashBob = token.getMintInvoiceHash(invoiceBob);
        console.logBytes32(invoiceHashBob);
        bytes memory signatureBob =
            hex"8436ac1084cdca7807c27570d5545b4dafb041481142712af5c2fc29c70ced1e977fabf45ce32aa6043036679e698b1494a7f4eafdab86c75fd02e440d6d138d";
        token.mint(invoiceBob, signatureBob);

        uint256 totalPooledBTC = token.totalSupply();
        uint256 totalShares = token.totalShares();

        uint256 expectedTotalPooledBTC = mintAmountAlice + mintAmountBob + INITIAL_SUPPLY;
        uint256 expectedShares = (mintAmountAlice + mintAmountBob + INITIAL_SUPPLY);

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

        strBTC.MintInvoice memory invoiceAlice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmountAlice});
        bytes32 invoiceHashAlice = token.getMintInvoiceHash(invoiceAlice);
        console.logBytes32(invoiceHashAlice);
        bytes memory signatureAlice =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        token.mint(invoiceAlice, signatureAlice);

        uint256 initialBalanceAlice = token.balanceOf(alice);
        assertEq(initialBalanceAlice, mintAmountAlice, "Alice's initial balance mismatch");

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(0, mintRewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        token.mintRewards(0, mintRewardAmount, validSignature);

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

        vm.prank(pauser);
        token.pause();

        strBTC.MintInvoice memory invoiceAlice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: 10 * BTC});
        bytes32 invoiceHashAlice = token.getMintInvoiceHash(invoiceAlice);
        console.logBytes32(invoiceHashAlice);
        bytes memory signatureAlice =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        token.mint(invoiceAlice, signatureAlice);

        vm.prank(alice);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        bytes32 totalSupplyUpdateHash = token.getTotalSupplyUpdateHash(1, 5 * BTC);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        token.mintRewards(0, 5 * BTC, validSignature);

        vm.prank(pauser);
        token.unpause();

        token.mint(invoiceAlice, signatureAlice);

        uint256 balanceAfterMint = token.balanceOf(alice);
        assertEq(balanceAfterMint, 10 * BTC, "Mint should work after unpause");

        token.mintRewards(0, 5 * BTC, validSignature);

        uint256 balanceAfterMintRewards = token.balanceOf(alice);
        uint256 expectedBalance = (token.getShares(alice) * token.totalSupply()) / token.totalShares();
        assertEq(balanceAfterMintRewards, expectedBalance, "Mint rewards should work after unpause");

        vm.prank(alice);
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");
        uint256 balanceAfterRedeem = token.balanceOf(alice);
        assertEq(balanceAfterRedeem, balanceAfterMintRewards - 5 * BTC, "Redeem should work after unpause");
    }

    function testEdgeCases() public {
        strBTC.MintInvoice memory invoiceZeroAmount =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: 0});
        bytes32 invoiceHashZeroAmount = token.getMintInvoiceHash(invoiceZeroAmount);
        console.logBytes32(invoiceHashZeroAmount);
        bytes memory signatureZeroAmount =
            hex"0a78867e5dced40470397b72963069ecfb66e9a030125974ecfacde8ce4e9ee93a1a01775aaae4b33e0a41e2379cbb46310ebf7fec9d8cf48650753d814b7871";

        vm.expectRevert(strBTC.MintAmountZero.selector);
        token.mint(invoiceZeroAmount, signatureZeroAmount);

        vm.prank(alice);
        vm.expectRevert(strBTC.AmountBelowMinWithdraw.selector);
        token.redeem(0, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: 10 * BTC});
        bytes32 invoiceHash = token.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        token.mint(invoice, signature);

        assertEq(token.totalShares(), 1010 * BTC, "Shares should equal minted amount in the initial state");
        assertEq(token.totalSupply(), 1010 * BTC, "Total pooled BTC should equal minted amount");

        vm.prank(alice);
        assertEq(token.balanceOf(alice), 10 * BTC, "Balance should match initial mint");

        strBTC.MintInvoice memory invoiceZeroAddress =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit2"), recipient: address(0), amount: 10 * BTC});
        bytes32 invoiceHashZeroAddress = token.getMintInvoiceHash(invoiceZeroAddress);
        console.logBytes32(invoiceHashZeroAddress);
        bytes memory signatureZeroAddress =
            hex"522ac790ad6b4a9daec6fc66ad61c24529e668b9de1f0402fb234c0850b0e984b1ece8c6fe4d6b22743686be5a76aa9e41e90b4a8baf3e1e28a9c48b5074ae52";

        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InvalidReceiver.selector, address(0)));
        token.mint(invoiceZeroAddress, signatureZeroAddress);

        vm.prank(alice);
        vm.expectRevert(strBTC.InvalidBTCAddress.selector);
        token.redeem(5 * BTC, "");

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InvalidReceiver.selector, address(0)));
        token.transfer(address(0), 5 * BTC);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientAllowance.selector, alice, 0, 5 * BTC));
        token.transferFrom(address(0), bob, 5 * BTC);

        vm.prank(alice);
        token.redeem(5 * BTC, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");

        assertEq(token.totalShares(), 1005 * BTC, "Shares should decrease after redeem");
        assertEq(token.totalSupply(), 1005 * BTC, "Total pooled BTC should decrease after redeem");
    }

    function testSetMinWithdrawAmount() public {
        uint256 newMinAmount = 2 * BTC;

        vm.prank(alice);
        vm.expectRevert();
        token.setMinWithdrawAmount(newMinAmount);

        vm.prank(admin);
        token.setMinWithdrawAmount(newMinAmount);
        assertEq(token.minWithdrawAmount(), newMinAmount);
    }

    function testRedeemWithInvalidAmount() public {
        uint256 mintAmount = 10 * BTC;

        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

        vm.prank(alice);
        vm.expectRevert(strBTC.InsufficientBalance.selector);
        token.redeem(mintAmount + 1, "1JTFoeWo4xPrQuVidgmzqoVRXLmd8pjtU9");
    }

    function testTotalSharesAndGetShares() public {
        assertEq(token.totalShares(), INITIAL_SUPPLY);
        assertEq(token.getShares(alice), 0);

        uint256 mintAmount = 10 * BTC;
        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount});
        bytes memory signature =
            hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
        token.mint(invoice, signature);

        assertEq(token.totalShares(), mintAmount + INITIAL_SUPPLY);
        assertEq(token.getShares(alice), mintAmount);
    }

    function testGetPooledBTCByShares() public {
        uint256 mintAmount1 = 10 * BTC;
        {
            strBTC.MintInvoice memory invoice1 =
                strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: mintAmount1});
            bytes memory signature1 =
                hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";
            token.mint(invoice1, signature1);
        }

        uint256 mintAmount2 = 15 * BTC;
        {
            strBTC.MintInvoice memory invoice2 =
                strBTC.MintInvoice({btcDepositId: keccak256("deposit2"), recipient: alice, amount: mintAmount2});
            bytes memory signature2 =
                hex"6e554675cfda2312235a42ee7e2fd27f649d2fb65c6e43019bc0adf84f4e6a15052627b0536139d29b00cde592a780805af57e90a96971ac26aea1ada07fa35f";
            token.mint(invoice2, signature2);
        }

        uint256 rewardAmount = 5 * BTC;
        {
            bytes memory validSignature =
                hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";
            token.mintRewards(0, rewardAmount, validSignature);
        }

        uint256 totalSupply = token.totalSupply(); // 30 BTC (10 + 15 + 5)
        uint256 totalShares = token.totalShares(); // 25 BTC (10 + 15)

        uint256 expectedBTC1 = (mintAmount1 * totalSupply) / totalShares;
        assertEq(token.getPooledBTCByShares(mintAmount1), expectedBTC1);

        uint256 expectedBTC2 = (mintAmount2 * totalSupply) / totalShares;
        assertEq(token.getPooledBTCByShares(mintAmount2), expectedBTC2);

        uint256 expectedTotalBTC = ((mintAmount1 + mintAmount2) * totalSupply) / totalShares;
        assertEq(token.getPooledBTCByShares(mintAmount1 + mintAmount2), expectedTotalBTC);

        uint256 expectedShares1 = (expectedBTC1 * totalShares) / totalSupply;
        assertEq(token.getSharesByPooledBTC(expectedBTC1), expectedShares1);

        uint256 expectedShares2 = (expectedBTC2 * totalShares) / totalSupply;
        assertEq(token.getSharesByPooledBTC(expectedBTC2), expectedShares2);

        uint256 expectedTotalShares = ((expectedBTC1 + expectedBTC2) * totalShares) / totalSupply;
        assertEq(token.getSharesByPooledBTC(expectedBTC1 + expectedBTC2), expectedTotalShares);

        assertApproxEqAbs(token.getPooledBTCByShares(token.getSharesByPooledBTC(1 * BTC)), 1 * BTC, 1);
        assertApproxEqAbs(token.getSharesByPooledBTC(token.getPooledBTCByShares(1 * BTC)), 1 * BTC, 1);
    }
}

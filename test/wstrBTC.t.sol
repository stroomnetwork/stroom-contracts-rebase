// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "../src/strBTC.sol";
import "../src/wstrBTC.sol";

contract WstrBTCTest is Test {
    strBTC public strBTCContract;
    wstrBTC public wstrBTCContract;

    ValidatorRegistry public vr;

    address public admin;
    address public pauser;

    address public alice;
    address public bob;

    uint256 public constant BTC = 1e8; // sat
    uint256 public constant INITIAL_SUPPLY = 1000 * BTC;

    function setUp() public {
        console.log("setUp");
        vm.warp(1_729_690_309);

        vr = new ValidatorRegistry();
        vr.setJointPublicKey(hex"9627e95c7c43a6550b0bcc005bbd85de78a1e17285c9acae2349292e78b21c0f");

        admin = msg.sender;
        pauser = makeAddr("Pauser");

        // Deploy strBTC implementation
        strBTC strBTCImplementation = new strBTC();
        bytes memory strBTCData =
            abi.encodeWithSelector(strBTC.initialize.selector, BitcoinNetworkEncoder.Network.Mainnet, vr, admin, pauser);
        TransparentUpgradeableProxy strBTCProxy =
            new TransparentUpgradeableProxy(address(strBTCImplementation), admin, strBTCData);
        strBTCContract = strBTC(address(strBTCProxy));

        // Deploy wstrBTC and link it to strBTC
        wstrBTCContract = new wstrBTC(address(strBTCContract));

        alice = makeAddr("alice");
        bob = makeAddr("bob");

        // Mint initial strBTC for Alice
        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit1"), recipient: alice, amount: INITIAL_SUPPLY});
        bytes memory signature =
            hex"af8d5abb288a9f53b9d847c0bed748dbaadb73dcdb0dd53aeeaaec5620b597e9b1a285de677e5a8a2d267b2f43bc5452eca3a1da96126c1bd7bc1961aaebbc3d";

        strBTCContract.mint(invoice, signature);
    }

    function testFuzzWrapCorrectAmountMinted(uint256 strBTCAmount) public {
        uint256 initialStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 initialWstrBTCBalance = wstrBTCContract.balanceOf(alice);

        vm.assume(strBTCAmount > 0 && strBTCAmount <= initialStrBTCBalance);

        uint256 strBTCBalanceBefore = strBTCContract.balanceOf(address(wstrBTCContract));
        uint256 wstrBTCSupplyBefore = wstrBTCContract.totalSupply();

        uint256 expectedWstrBTC;
        if (wstrBTCSupplyBefore == 0 || strBTCBalanceBefore == 0) {
            expectedWstrBTC = strBTCAmount;
        } else {
            expectedWstrBTC = (strBTCAmount * wstrBTCSupplyBefore) / strBTCBalanceBefore;
        }

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        uint256 mintedWstrBTC = wstrBTCContract.wrap(strBTCAmount);

        assertEq(mintedWstrBTC, expectedWstrBTC, "Incorrect amount of wstrBTC minted");

        uint256 finalStrBTCBalance = strBTCContract.balanceOf(alice);
        assertEq(finalStrBTCBalance, initialStrBTCBalance - strBTCAmount, "strBTC balance did not decrease correctly");

        uint256 finalWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(
            finalWstrBTCBalance, initialWstrBTCBalance + expectedWstrBTC, "wstrBTC balance did not increase correctly"
        );

        uint256 finalContractStrBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));
        assertEq(
            finalContractStrBTCBalance,
            strBTCBalanceBefore + strBTCAmount,
            "strBTC balance of wstrBTC contract incorrect"
        );
    }

    function testWrapZeroAmountReverts() public {
        uint256 strBTCAmount = 0;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        vm.expectRevert(wstrBTC.CannotWrapZero.selector);
        wstrBTCContract.wrap(strBTCAmount);
    }

    function testWrapInsufficientApproval() public {
        uint256 approveAmount = 5 * BTC;
        uint256 wrapAmount = 7 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), approveAmount);

        uint256 allowance = strBTCContract.allowance(alice, address(wstrBTCContract));
        assertEq(allowance, approveAmount, "Allowance incorrect");

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, address(wstrBTCContract), allowance, wrapAmount
            )
        );
        wstrBTCContract.wrap(wrapAmount);
    }

    function testWrapInsufficientBalance() public {
        uint256 wrapAmount = INITIAL_SUPPLY + 1 * BTC;

        uint256 aliceBalance = strBTCContract.balanceOf(alice);
        assertEq(aliceBalance, INITIAL_SUPPLY, "Alice initial balance incorrect");

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), wrapAmount);

        vm.prank(alice);
        vm.expectRevert(strBTC.InsufficientBalance.selector);
        wstrBTCContract.wrap(wrapAmount);
    }

    function testUnwrapSuccess() public {
        uint256 wrapAmount = 5 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), wrapAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.wrap(wrapAmount);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after wrap");

        uint256 strBTCBalanceOfWrapper = strBTCContract.balanceOf(address(wstrBTCContract));
        assertEq(strBTCBalanceOfWrapper, wrapAmount, "wstrBTC contract strBTC balance incorrect");

        vm.prank(alice);
        uint256 strBTCUnwrapped = wstrBTCContract.unwrap(wstrBTCMinted);

        assertEq(strBTCUnwrapped, wrapAmount, "Unwrapped strBTC amount incorrect");

        aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after unwrap");

        uint256 aliceStrBTCBalance = strBTCContract.balanceOf(alice);
        assertEq(aliceStrBTCBalance, INITIAL_SUPPLY, "Alice's strBTC balance incorrect after unwrap");
    }

    function testUnwrapAfterMintRewards() public {
        uint256 initialWrapAmount = 5 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialWrapAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.wrap(initialWrapAmount);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after wrap");

        uint256 rewardAmount = 5 * BTC;

        bytes32 totalSupplyUpdateHash = strBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalSupplyAfterRewards = strBTCContract.totalSupply();
        uint256 totalSharesAfterRewards = strBTCContract.totalShares();

        assertEq(totalSupplyAfterRewards, INITIAL_SUPPLY + rewardAmount, "Total supply incorrect after rewards");

        uint256 expectedStrBTCFromUnwrap = (wstrBTCMinted * totalSupplyAfterRewards) / totalSharesAfterRewards;

        vm.prank(alice);
        uint256 strBTCUnwrapped = wstrBTCContract.unwrap(wstrBTCMinted);

        assertEq(strBTCUnwrapped, expectedStrBTCFromUnwrap, "Unwrapped strBTC amount incorrect after rewards");

        uint256 aliceFinalStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 expectedFinalBalance = INITIAL_SUPPLY + rewardAmount;

        assertEq(aliceFinalStrBTCBalance, expectedFinalBalance, "Alice's strBTC balance incorrect after unwrap");

        uint256 aliceFinalWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceFinalWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after unwrap");
    }

    function testUnwrapZeroWstrBTC() public {
        vm.startPrank(alice);

        uint256 zeroAmount = 0;

        vm.expectRevert(wstrBTC.CannotUnwrapZero.selector);
        wstrBTCContract.unwrap(zeroAmount);

        vm.stopPrank();
    }

    function testUnwrapInsufficientBalance() public {
        vm.startPrank(bob);

        uint256 insufficientAmount = 1 * 10 ** 18;

        vm.expectRevert(wstrBTC.NoWstrBTCSupply.selector);
        wstrBTCContract.unwrap(insufficientAmount);

        vm.stopPrank();
    }

    function testUnwrapCallsTransfer() public {
        uint256 initialStrBTCAmount = 5 * BTC;
        vm.startPrank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialStrBTCAmount);
        uint256 wstrBTCMinted = wstrBTCContract.wrap(initialStrBTCAmount);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after wrap");

        vm.expectCall(
            address(strBTCContract),
            abi.encodeWithSelector(strBTCContract.transfer.selector, alice, initialStrBTCAmount)
        );
        wstrBTCContract.unwrap(wstrBTCMinted);
        vm.stopPrank();
    }

    function testFuzzWrap(uint256 strBTCAmount) public {
        uint256 initialAliceStrBTCAmount = strBTCContract.balanceOf(alice);
        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialContractStrBTCAmount = strBTCContract.balanceOf(address(wstrBTCContract));

        vm.assume(strBTCAmount > 0 && strBTCAmount <= initialAliceStrBTCAmount);

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.wrap(strBTCAmount);

        assertEq(
            strBTCContract.balanceOf(alice), initialAliceStrBTCAmount - strBTCAmount, "Alice's strBTC balance incorrect"
        );
        assertEq(
            wstrBTCContract.balanceOf(alice),
            initialAliceWstrBTCBalance + wstrBTCMinted,
            "Alice's wstrBTC balance incorrect"
        );
        assertEq(
            strBTCContract.balanceOf(address(wstrBTCContract)),
            initialContractStrBTCAmount + strBTCAmount,
            "wstrBTC contract's strBTC balance incorrect"
        );
    }

    function testFuzzUnwrap(uint256 wstrBTCAmount) public {
        uint256 initialAliceStrBTCAmount = strBTCContract.balanceOf(alice);

        vm.startPrank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialAliceStrBTCAmount);

        uint256 wstrBTCMinted = wstrBTCContract.wrap(initialAliceStrBTCAmount);

        vm.assume(wstrBTCAmount > 0 && wstrBTCAmount <= wstrBTCMinted);

        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialContractStrBTCAmount = strBTCContract.balanceOf(address(wstrBTCContract));

        uint256 strBTCUnwrapped = wstrBTCContract.unwrap(wstrBTCAmount);
        vm.stopPrank();

        assertEq(strBTCContract.balanceOf(alice), strBTCUnwrapped, "Alice's strBTC balance incorrect");
        assertEq(
            wstrBTCContract.balanceOf(alice),
            initialAliceWstrBTCBalance - wstrBTCAmount,
            "Alice's wstrBTC balance incorrect"
        );
        assertEq(
            strBTCContract.balanceOf(address(wstrBTCContract)),
            initialContractStrBTCAmount - strBTCUnwrapped,
            "wstrBTC contract's strBTC balance incorrect"
        );
    }

    function testIntegrationWrapUnwrapWithRebase(uint256 wrapAmount) public {
        uint256 initialStrBTC = strBTCContract.balanceOf(alice);
        vm.assume(wrapAmount > 0 && wrapAmount < initialStrBTC);

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), wrapAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.wrap(wrapAmount);

        uint256 initialAliceStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialTotalPooledStrBTC = strBTCContract.totalSupply();

        assertEq(initialAliceStrBTCBalance, initialStrBTC - wrapAmount, "Alice's strBTC balance after wrap incorrect");
        assertEq(initialAliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance after wrap incorrect");

        uint256 rewardAmount = 5 * BTC;
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalPooledStrBTCAfterRebase = strBTCContract.totalSupply();
        uint256 aliceStrBTCBalanceAfterRebase = strBTCContract.balanceOf(alice);

        assertGe(
            totalPooledStrBTCAfterRebase, initialTotalPooledStrBTC, "Total pooled strBTC after rebase should increase"
        );
        assertGe(
            aliceStrBTCBalanceAfterRebase,
            initialAliceStrBTCBalance,
            "Alice's strBTC balance after rebase should increase"
        );

        vm.prank(alice);
        uint256 strBTCUnwrapped = wstrBTCContract.unwrap(wstrBTCMinted);

        uint256 aliceStrBTCBalanceAfterUnwrap = strBTCContract.balanceOf(alice);
        uint256 finalAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);

        assertEq(finalAliceWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after unwrap");

        uint256 expectedBalance = initialStrBTC + rewardAmount;

        assertApproxEqAbs(
            aliceStrBTCBalanceAfterUnwrap,
            expectedBalance,
            10,
            "Alice's strBTC balance after unwrap should include rewards"
        );

        assertEq(
            strBTCUnwrapped,
            (wstrBTCMinted * totalPooledStrBTCAfterRebase) / strBTCContract.totalShares(),
            "Unwrapped strBTC amount is incorrect"
        );
    }

    function testFuzzIntegrationWrapUnwrapWithRebaseForTwoUsers(uint256 aliceWrapAmount, uint256 bobWrapAmount)
        public
    {
        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit2"), recipient: bob, amount: 15 * BTC});
        bytes memory signature =
            hex"afbfdb08ce06fac74578357564a061daf21b5e9e8829b052173c6de7da2db52dbe5cb8ced766d4e3d8c7a0cca7e944f29de838d36673f2b3c5ae3793c8f7d866";

        strBTCContract.mint(invoice, signature);

        vm.assume(aliceWrapAmount > 0 && aliceWrapAmount <= strBTCContract.balanceOf(alice));
        vm.assume(bobWrapAmount > 0 && bobWrapAmount <= strBTCContract.balanceOf(bob));

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), aliceWrapAmount);

        vm.prank(bob);
        strBTCContract.approve(address(wstrBTCContract), bobWrapAmount);

        vm.prank(alice);
        uint256 aliceWstrBTCMinted = wstrBTCContract.wrap(aliceWrapAmount);

        vm.prank(bob);
        uint256 bobWstrBTCMinted = wstrBTCContract.wrap(bobWrapAmount);

        uint256 initialTotalShares = strBTCContract.totalShares();

        uint256 rewardAmount = 5 * BTC;
        bytes32 totalSupplyUpdateHash = strBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalPooledStrBTCAfterRebase = strBTCContract.totalSupply();

        uint256 expectedAliceUnwrapped = (aliceWstrBTCMinted * totalPooledStrBTCAfterRebase) / initialTotalShares;
        uint256 expectedBobUnwrapped = (bobWstrBTCMinted * totalPooledStrBTCAfterRebase) / initialTotalShares;

        vm.prank(alice);
        uint256 aliceStrBTCUnwrapped = wstrBTCContract.unwrap(aliceWstrBTCMinted);

        vm.prank(bob);
        uint256 bobStrBTCUnwrapped = wstrBTCContract.unwrap(bobWstrBTCMinted);

        assertApproxEqAbs(
            aliceStrBTCUnwrapped, expectedAliceUnwrapped, 10, "Alice's unwrapped strBTC amount incorrect after rebase"
        );

        assertApproxEqAbs(
            bobStrBTCUnwrapped, expectedBobUnwrapped, 10, "Bob's unwrapped strBTC amount incorrect after rebase"
        );

        assertEq(wstrBTCContract.balanceOf(alice), 0, "Alice's wstrBTC balance should be zero after unwrap");
        assertEq(wstrBTCContract.balanceOf(bob), 0, "Bob's wstrBTC balance should be zero after unwrap");
    }

    function testWrapWithExcessBalance() public {
        uint256 strBTCAmount = 1 * BTC;

        deal(address(strBTCContract), address(wstrBTCContract), 2000 * BTC);

        uint256 initialStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 initialWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 contractStrBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));

        uint256 totalPooledStrBTC = strBTCContract.totalSupply();
        assertGt(contractStrBTCBalance, totalPooledStrBTC, "Contract balance should exceed total pooled strBTC");

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        wstrBTCContract.wrap(strBTCAmount);

        uint256 finalStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 finalWstrBTCBalance = wstrBTCContract.balanceOf(alice);

        assertEq(finalStrBTCBalance, initialStrBTCBalance - strBTCAmount, "Alice's strBTC balance incorrect after wrap");
        assertGt(finalWstrBTCBalance, initialWstrBTCBalance, "Alice's wstrBTC balance incorrect after wrap");

        uint256 finalContractStrBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));
        assertEq(
            finalContractStrBTCBalance,
            contractStrBTCBalance + strBTCAmount,
            "wstrBTC contract's strBTC balance incorrect after wrap"
        );
    }

    function testZeroSupplyEdgeCase() public {
        uint256 initialWstrBTCTotalSupply = wstrBTCContract.totalSupply();
        uint256 initialWstrBTCBalance = wstrBTCContract.balanceOf(address(this));

        assertEq(initialWstrBTCTotalSupply, 0, "Initial wstrBTC total supply is not zero");
        assertEq(initialWstrBTCBalance, 0, "Initial wstrBTC balance is not zero");

        vm.expectRevert(wstrBTC.NoWstrBTCSupply.selector);
        wstrBTCContract.strBTCPerToken(BTC);

        vm.expectRevert(wstrBTC.NoStrBTCBalance.selector);
        wstrBTCContract.tokensPerStrBTC(BTC);

        uint256 strBTCAmount = 1 * BTC;
        strBTC.MintInvoice memory invoice = strBTC.MintInvoice({
            btcDepositId: keccak256("initial_deposit"),
            recipient: address(this),
            amount: strBTCAmount
        });

        bytes32 invoiceHash = strBTCContract.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature =
            hex"4bfdff7e1b987ba478fdb196010e6ed8be7006e85d9bcec6b9b15887a7b9d9b4e9fd3389a838897952068261c1f49df76ddda11ad0465989049e07cbe73df200";

        strBTCContract.mint(invoice, signature);

        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        uint256 wstrBTCMinted = wstrBTCContract.wrap(strBTCAmount);
        assertEq(wstrBTCMinted, strBTCAmount, "Minted wstrBTC should equal strBTC amount");

        uint256 strBTCUnwrapped = wstrBTCContract.unwrap(wstrBTCMinted);
        assertEq(strBTCUnwrapped, strBTCAmount, "Unwrapped strBTC should equal initial amount");

        uint256 finalWstrBTCTotalSupply = wstrBTCContract.totalSupply();
        assertEq(finalWstrBTCTotalSupply, 0, "wstrBTC total supply after unwrap should be zero");
    }

    function testStrBTCPerTokenAndTokensPerStrBTC() public {
        vm.expectRevert(wstrBTC.NoWstrBTCSupply.selector);
        wstrBTCContract.strBTCPerToken(BTC);

        vm.expectRevert(wstrBTC.NoStrBTCBalance.selector);
        wstrBTCContract.tokensPerStrBTC(BTC);

        uint256 strBTCAmount = 10 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        wstrBTCContract.wrap(strBTCAmount);

        uint256 strBTCPerToken = wstrBTCContract.strBTCPerToken(BTC);
        uint256 tokensPerStrBTC = wstrBTCContract.tokensPerStrBTC(BTC);

        uint256 strBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));
        uint256 wstrBTCSupply = wstrBTCContract.totalSupply();

        uint256 expectedStrBTCPerToken = (strBTCBalance * BTC) / wstrBTCSupply;
        uint256 expectedTokensPerStrBTC = (wstrBTCSupply * BTC) / strBTCBalance;

        assertEq(strBTCPerToken, expectedStrBTCPerToken, "strBTCPerToken calculation is incorrect");
        assertEq(tokensPerStrBTC, expectedTokensPerStrBTC, "tokensPerStrBTC calculation is incorrect");

        uint256 rewardAmount = 5 * BTC;
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";
        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        strBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));
        wstrBTCSupply = wstrBTCContract.totalSupply();

        uint256 strBTCPerTokenAfterRebase = wstrBTCContract.strBTCPerToken(BTC);
        uint256 tokensPerStrBTCAfterRebase = wstrBTCContract.tokensPerStrBTC(BTC);

        expectedStrBTCPerToken = (strBTCBalance * BTC) / wstrBTCSupply;
        expectedTokensPerStrBTC = (wstrBTCSupply * BTC) / strBTCBalance;

        assertEq(
            strBTCPerTokenAfterRebase, expectedStrBTCPerToken, "strBTCPerToken calculation after rebase is incorrect"
        );
        assertEq(
            tokensPerStrBTCAfterRebase, expectedTokensPerStrBTC, "tokensPerStrBTC calculation after rebase is incorrect"
        );
    }
}

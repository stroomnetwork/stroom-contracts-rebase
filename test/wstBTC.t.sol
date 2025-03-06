// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "../src/stBTC.sol";
import "../src/wstBTC.sol";

contract WstBTCTest is Test {
    stBTC public stBTCContract;
    wstBTC public wstBTCContract;

    ValidatorRegistry public vr;

    address public admin;

    address public alice;
    address public bob;

    uint256 public constant BTC = 1e8; // sat

    function setUp() public {
        console.log("setUp");

        vr = new ValidatorRegistry();
        vr.setJointPublicKey(hex"9627e95c7c43a6550b0bcc005bbd85de78a1e17285c9acae2349292e78b21c0f");

        admin = msg.sender;

        // Deploy stBTC implementation
        stBTC stBTCImplementation = new stBTC();
        bytes memory stBTCData =
            abi.encodeWithSelector(stBTC.initialize.selector, BitcoinNetworkEncoder.Network.Mainnet, vr);
        TransparentUpgradeableProxy stBTCProxy =
            new TransparentUpgradeableProxy(address(stBTCImplementation), admin, stBTCData);
        stBTCContract = stBTC(address(stBTCProxy));

        // Deploy wstBTC and link it to stBTC
        wstBTCContract = new wstBTC(address(stBTCContract));

        alice = makeAddr("alice");
        bob = makeAddr("bob");

        // Mint initial stBTC for Alice
        stBTC.MintInvoice memory invoice = stBTC.MintInvoice({
            btcDepositId: keccak256("deposit1"),
            recipient: alice,
            amount: 10 * BTC
        });
        bytes memory signature = hex"41a3536b1cdcaed9205fd3cc79c405c6ae6be89e4acfb5f7298d2f6a17c710bef11896d61e83d936d7da8335f5d3908d8f2cf2e42e07ada17223739419c7001c";

        stBTCContract.mint(invoice, signature);
    }

    function testFuzzWrapCorrectAmountMinted(uint256 stBTCAmount) public {
        uint256 initialStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 initialWstBTCBalance = wstBTCContract.balanceOf(alice);

        vm.assume(stBTCAmount > 0 && stBTCAmount <= initialStBTCBalance);

        uint256 stBTCBalanceBefore = stBTCContract.balanceOf(address(wstBTCContract));
        uint256 wstBTCSupplyBefore = wstBTCContract.totalSupply();

        uint256 expectedWstBTC;
        if (wstBTCSupplyBefore == 0 || stBTCBalanceBefore == 0) {
            expectedWstBTC = stBTCAmount;
        } else {
            expectedWstBTC = (stBTCAmount * wstBTCSupplyBefore) / stBTCBalanceBefore;
        }

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        uint256 mintedWstBTC = wstBTCContract.wrap(stBTCAmount);

        assertEq(mintedWstBTC, expectedWstBTC, "Incorrect amount of wstBTC minted");

        uint256 finalStBTCBalance = stBTCContract.balanceOf(alice);
        assertEq(finalStBTCBalance, initialStBTCBalance - stBTCAmount, "stBTC balance did not decrease correctly");

        uint256 finalWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(finalWstBTCBalance, initialWstBTCBalance + expectedWstBTC, "wstBTC balance did not increase correctly");

        uint256 finalContractStBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        assertEq(
            finalContractStBTCBalance, stBTCBalanceBefore + stBTCAmount, "stBTC balance of wstBTC contract incorrect"
        );
    }

    function testWrapZeroAmountReverts() public {
        uint256 stBTCAmount = 0;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        vm.expectRevert(wstBTC.CannotWrapZero.selector);
        wstBTCContract.wrap(stBTCAmount);
    }

    function testWrapInsufficientApproval() public {
        uint256 approveAmount = 5 * BTC;
        uint256 wrapAmount = 7 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), approveAmount);

        uint256 allowance = stBTCContract.allowance(alice, address(wstBTCContract));
        assertEq(allowance, approveAmount, "Allowance incorrect");

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientAllowance.selector, address(wstBTCContract), allowance, wrapAmount));
        wstBTCContract.wrap(wrapAmount);
    }

    function testWrapInsufficientBalance() public {
        uint256 wrapAmount = 20 * BTC;

        uint256 aliceBalance = stBTCContract.balanceOf(alice);
        assertEq(aliceBalance, 10 * BTC, "Alice initial balance incorrect");

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), wrapAmount);

        vm.prank(alice);
        vm.expectRevert(stBTC.InsufficientBalance.selector);
        wstBTCContract.wrap(wrapAmount);
    }

    function testUnwrapSuccess() public {
        uint256 wrapAmount = 5 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), wrapAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(wrapAmount);

        uint256 aliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(aliceWstBTCBalance, wstBTCMinted, "Alice's wstBTC balance incorrect after wrap");

        uint256 stBTCBalanceOfWrapper = stBTCContract.balanceOf(address(wstBTCContract));
        assertEq(stBTCBalanceOfWrapper, wrapAmount, "wstBTC contract stBTC balance incorrect");

        vm.prank(alice);
        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);

        assertEq(stBTCUnwrapped, wrapAmount, "Unwrapped stBTC amount incorrect");

        aliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(aliceWstBTCBalance, 0, "Alice's wstBTC balance should be zero after unwrap");

        uint256 aliceStBTCBalance = stBTCContract.balanceOf(alice);
        assertEq(aliceStBTCBalance, 10 * BTC, "Alice's stBTC balance incorrect after unwrap");
    }

    function testUnwrapAfterMintRewards() public {
        uint256 initialAliceBalance = 10 * BTC;

        uint256 initialWrapAmount = 5 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), initialWrapAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(initialWrapAmount);

        uint256 aliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(aliceWstBTCBalance, wstBTCMinted, "Alice's wstBTC balance incorrect after wrap");

        uint256 rewardAmount = 5 * BTC;

        bytes32 totalSupplyUpdateHash = stBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature = hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        stBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalSupplyAfterRewards = stBTCContract.totalSupply();
        uint256 totalSharesAfterRewards = stBTCContract.totalShares();

        assertEq(totalSupplyAfterRewards, 15 * BTC, "Total supply incorrect after rewards");

        uint256 expectedStBTCFromUnwrap = (wstBTCMinted * totalSupplyAfterRewards) / totalSharesAfterRewards;
        uint256 expectedRemainingBalanceAfterRebase =
            ((initialAliceBalance - initialWrapAmount) * totalSupplyAfterRewards) / totalSharesAfterRewards;
        uint256 expectedFinalBalance = expectedRemainingBalanceAfterRebase + expectedStBTCFromUnwrap;

        vm.prank(alice);
        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);

        assertEq(stBTCUnwrapped, expectedStBTCFromUnwrap, "Unwrapped stBTC amount incorrect after rewards");

        uint256 aliceFinalStBTCBalance = stBTCContract.balanceOf(alice);
        assertEq(aliceFinalStBTCBalance, expectedFinalBalance, "Alice's stBTC balance incorrect after unwrap");

        aliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(aliceWstBTCBalance, 0, "Alice's wstBTC balance should be zero after unwrap");
    }

    function testUnwrapZeroWstBTC() public {
        vm.startPrank(alice);

        uint256 zeroAmount = 0;

        vm.expectRevert(wstBTC.CannotUnwrapZero.selector);
        wstBTCContract.unwrap(zeroAmount);

        vm.stopPrank();
    }

    function testUnwrapInsufficientBalance() public {
        vm.startPrank(bob);

        uint256 insufficientAmount = 1 * 10 ** 18;

        vm.expectRevert(wstBTC.NoWstBTCSupply.selector);
        wstBTCContract.unwrap(insufficientAmount);

        vm.stopPrank();
    }

    function testUnwrapCallsTransfer() public {
        uint256 initialStBTCAmount = 5 * BTC;
        vm.startPrank(alice);
        stBTCContract.approve(address(wstBTCContract), initialStBTCAmount);
        uint256 wstBTCMinted = wstBTCContract.wrap(initialStBTCAmount);

        uint256 aliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        assertEq(aliceWstBTCBalance, wstBTCMinted, "Alice's wstBTC balance incorrect after wrap");

        vm.expectCall(
            address(stBTCContract), abi.encodeWithSelector(stBTCContract.transfer.selector, alice, initialStBTCAmount)
        );
        wstBTCContract.unwrap(wstBTCMinted);
        vm.stopPrank();
    }

    function testFuzzWrap(uint256 stBTCAmount) public {
        uint256 initialAliceStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 initialAliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        uint256 initialContractStBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));

        vm.assume(stBTCAmount > 0 && stBTCAmount <= initialAliceStBTCBalance);

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(stBTCAmount);

        assertEq(
            stBTCContract.balanceOf(alice), initialAliceStBTCBalance - stBTCAmount, "Alice's stBTC balance incorrect"
        );
        assertEq(
            wstBTCContract.balanceOf(alice),
            initialAliceWstBTCBalance + wstBTCMinted,
            "Alice's wstBTC balance incorrect"
        );
        assertEq(
            stBTCContract.balanceOf(address(wstBTCContract)),
            initialContractStBTCBalance + stBTCAmount,
            "wstBTC contract's stBTC balance incorrect"
        );
    }

    function testFuzzUnwrap(uint256 wstBTCAmount) public {
        uint256 initialAliceStBTCBalance = stBTCContract.balanceOf(alice);

        vm.startPrank(alice);
        stBTCContract.approve(address(wstBTCContract), initialAliceStBTCBalance);

        uint256 wstBTCMinted = wstBTCContract.wrap(initialAliceStBTCBalance);

        vm.assume(wstBTCAmount > 0 && wstBTCAmount <= wstBTCMinted);

        uint256 initialAliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        uint256 initialContractStBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));

        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCAmount);
        vm.stopPrank();

        assertEq(stBTCContract.balanceOf(alice), stBTCUnwrapped, "Alice's stBTC balance incorrect");
        assertEq(
            wstBTCContract.balanceOf(alice),
            initialAliceWstBTCBalance - wstBTCAmount,
            "Alice's wstBTC balance incorrect"
        );
        assertEq(
            stBTCContract.balanceOf(address(wstBTCContract)),
            initialContractStBTCBalance - stBTCUnwrapped,
            "wstBTC contract's stBTC balance incorrect"
        );
    }

    function testIntegrationWrapUnwrapWithRebase(uint256 wrapAmount) public {
        uint256 initialStBTC = stBTCContract.balanceOf(alice);
        vm.assume(wrapAmount > 0 && wrapAmount < initialStBTC);

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), wrapAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(wrapAmount);

        uint256 initialAliceStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 initialAliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        uint256 initialTotalPooledBTC = stBTCContract.totalSupply();

        assertEq(initialAliceStBTCBalance, initialStBTC - wrapAmount, "Alice's stBTC balance after wrap incorrect");
        assertEq(initialAliceWstBTCBalance, wstBTCMinted, "Alice's wstBTC balance after wrap incorrect");

        uint256 rewardAmount = 5 * BTC;
        bytes memory validSignature = hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        stBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalPooledBTCAfterRebase = stBTCContract.totalSupply();
        uint256 aliceStBTCBalanceAfterRebase = stBTCContract.balanceOf(alice);

        assertGe(totalPooledBTCAfterRebase, initialTotalPooledBTC, "Total pooled BTC after rebase should increase");
        assertGe(
            aliceStBTCBalanceAfterRebase, initialAliceStBTCBalance, "Alice's stBTC balance after rebase should increase"
        );

        vm.prank(alice);
        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);

        uint256 aliceStBTCBalanceAfterUnwrap = stBTCContract.balanceOf(alice);
        uint256 finalAliceWstBTCBalance = wstBTCContract.balanceOf(alice);

        assertEq(finalAliceWstBTCBalance, 0, "Alice's wstBTC balance should be zero after unwrap");

        uint256 expectedBalance = initialStBTC + rewardAmount;

        assertApproxEqAbs(
            aliceStBTCBalanceAfterUnwrap,
            expectedBalance,
            10,
            "Alice's stBTC balance after unwrap should include rewards"
        );

        assertEq(
            stBTCUnwrapped,
            (wstBTCMinted * totalPooledBTCAfterRebase) / stBTCContract.totalShares(),
            "Unwrapped stBTC amount is incorrect"
        );
    }

    function testFuzzIntegrationWrapUnwrapWithRebaseForTwoUsers(
        uint256 aliceWrapAmount,
        uint256 bobWrapAmount
    ) public {
        stBTC.MintInvoice memory invoice = stBTC.MintInvoice({
            btcDepositId: keccak256("deposit2"),
            recipient: bob,
            amount: 15 * BTC
        });
        bytes memory signature = hex"afbfdb08ce06fac74578357564a061daf21b5e9e8829b052173c6de7da2db52dbe5cb8ced766d4e3d8c7a0cca7e944f29de838d36673f2b3c5ae3793c8f7d866";
        
        stBTCContract.mint(invoice, signature);

        vm.assume(aliceWrapAmount > 0 && aliceWrapAmount <= stBTCContract.balanceOf(alice));
        vm.assume(bobWrapAmount > 0 && bobWrapAmount <= stBTCContract.balanceOf(bob));

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), aliceWrapAmount);

        vm.prank(bob);
        stBTCContract.approve(address(wstBTCContract), bobWrapAmount);

        vm.prank(alice);
        uint256 aliceWstBTCMinted = wstBTCContract.wrap(aliceWrapAmount);

        vm.prank(bob);
        uint256 bobWstBTCMinted = wstBTCContract.wrap(bobWrapAmount);

        uint256 initialTotalShares = stBTCContract.totalShares();

        uint256 rewardAmount = 5 * BTC;
        bytes32 totalSupplyUpdateHash = stBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature = hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        stBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalPooledBTCAfterRebase = stBTCContract.totalSupply();

        uint256 expectedAliceUnwrapped = (aliceWstBTCMinted * totalPooledBTCAfterRebase) / initialTotalShares;
        uint256 expectedBobUnwrapped = (bobWstBTCMinted * totalPooledBTCAfterRebase) / initialTotalShares;

        vm.prank(alice);
        uint256 aliceStBTCUnwrapped = wstBTCContract.unwrap(aliceWstBTCMinted);

        vm.prank(bob);
        uint256 bobStBTCUnwrapped = wstBTCContract.unwrap(bobWstBTCMinted);

        assertApproxEqAbs(
            aliceStBTCUnwrapped, expectedAliceUnwrapped, 10, "Alice's unwrapped stBTC amount incorrect after rebase"
        );

        assertApproxEqAbs(
            bobStBTCUnwrapped, expectedBobUnwrapped, 10, "Bob's unwrapped stBTC amount incorrect after rebase"
        );

        assertEq(wstBTCContract.balanceOf(alice), 0, "Alice's wstBTC balance should be zero after unwrap");
        assertEq(wstBTCContract.balanceOf(bob), 0, "Bob's wstBTC balance should be zero after unwrap");
    }

    function testWrapWithExcessBalance() public {
        uint256 stBTCAmount = 1 * BTC;

        deal(address(stBTCContract), address(wstBTCContract), 20 * BTC);

        uint256 initialStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 initialWstBTCBalance = wstBTCContract.balanceOf(alice);
        uint256 contractStBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));

        uint256 totalPooledBTC = stBTCContract.totalSupply();
        assertGt(contractStBTCBalance, totalPooledBTC, "Contract balance should exceed total pooled BTC");

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        wstBTCContract.wrap(stBTCAmount);

        uint256 finalStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 finalWstBTCBalance = wstBTCContract.balanceOf(alice);

        assertEq(finalStBTCBalance, initialStBTCBalance - stBTCAmount, "Alice's stBTC balance incorrect after wrap");
        assertGt(finalWstBTCBalance, initialWstBTCBalance, "Alice's wstBTC balance incorrect after wrap");

        uint256 finalContractStBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        assertEq(
            finalContractStBTCBalance,
            contractStBTCBalance + stBTCAmount,
            "wstBTC contract's stBTC balance incorrect after wrap"
        );
    }

    function testZeroSupplyEdgeCase() public {
        uint256 initialWstBTCTotalSupply = wstBTCContract.totalSupply();
        uint256 initialWstBTCBalance = wstBTCContract.balanceOf(address(this));

        assertEq(initialWstBTCTotalSupply, 0, "Initial wstBTC total supply is not zero");
        assertEq(initialWstBTCBalance, 0, "Initial wstBTC balance is not zero");

        vm.expectRevert(wstBTC.NoWstBTCSupply.selector);
        wstBTCContract.stBTCPerToken(BTC);

        vm.expectRevert(wstBTC.NoStBTCBalance.selector);
        wstBTCContract.tokensPerStBTC(BTC);

        uint256 stBTCAmount = 1 * BTC;
        stBTC.MintInvoice memory invoice = stBTC.MintInvoice({
            btcDepositId: keccak256("initial_deposit"),
            recipient: address(this),
            amount: stBTCAmount
        });

        bytes32 invoiceHash = stBTCContract.getMintInvoiceHash(invoice);
        console.logBytes32(invoiceHash);
        bytes memory signature = hex"4bfdff7e1b987ba478fdb196010e6ed8be7006e85d9bcec6b9b15887a7b9d9b4e9fd3389a838897952068261c1f49df76ddda11ad0465989049e07cbe73df200";
       
        stBTCContract.mint(invoice, signature);

        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        uint256 wstBTCMinted = wstBTCContract.wrap(stBTCAmount);
        assertEq(wstBTCMinted, stBTCAmount, "Minted wstBTC should equal stBTC amount");

        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);
        assertEq(stBTCUnwrapped, stBTCAmount, "Unwrapped stBTC should equal initial amount");

        uint256 finalWstBTCTotalSupply = wstBTCContract.totalSupply();
        assertEq(finalWstBTCTotalSupply, 0, "wstBTC total supply after unwrap should be zero");
    }

    function testStBTCPerTokenAndTokensPerStBTC() public {
        vm.expectRevert(wstBTC.NoWstBTCSupply.selector);
        wstBTCContract.stBTCPerToken(BTC);

        vm.expectRevert(wstBTC.NoStBTCBalance.selector);
        wstBTCContract.tokensPerStBTC(BTC);

        uint256 stBTCAmount = 10 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        wstBTCContract.wrap(stBTCAmount);

        uint256 stBTCPerToken = wstBTCContract.stBTCPerToken(BTC);
        uint256 tokensPerStBTC = wstBTCContract.tokensPerStBTC(BTC);

        uint256 stBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        uint256 wstBTCSupply = wstBTCContract.totalSupply();

        uint256 expectedStBTCPerToken = (stBTCBalance * BTC) / wstBTCSupply;
        uint256 expectedTokensPerStBTC = (wstBTCSupply * BTC) / stBTCBalance;

        assertEq(stBTCPerToken, expectedStBTCPerToken, "stBTCPerToken calculation is incorrect");
        assertEq(tokensPerStBTC, expectedTokensPerStBTC, "tokensPerStBTC calculation is incorrect");

        uint256 rewardAmount = 5 * BTC;
        bytes memory validSignature = hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";
        stBTCContract.mintRewards(0, rewardAmount, validSignature);

        stBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        wstBTCSupply = wstBTCContract.totalSupply();

        uint256 stBTCPerTokenAfterRebase = wstBTCContract.stBTCPerToken(BTC);
        uint256 tokensPerStBTCAfterRebase = wstBTCContract.tokensPerStBTC(BTC);

        expectedStBTCPerToken = (stBTCBalance * BTC) / wstBTCSupply;
        expectedTokensPerStBTC = (wstBTCSupply * BTC) / stBTCBalance;

        assertEq(stBTCPerTokenAfterRebase, expectedStBTCPerToken, "stBTCPerToken calculation after rebase is incorrect");
        assertEq(
            tokensPerStBTCAfterRebase, expectedTokensPerStBTC, "tokensPerStBTC calculation after rebase is incorrect"
        );
    }
}

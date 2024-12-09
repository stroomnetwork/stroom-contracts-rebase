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
        stBTCContract.mint(10 * BTC, alice, keccak256("alice_initial_deposit"));
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
        vm.expectRevert("wstBTC: Cannot wrap zero stBTC");
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
        vm.expectRevert();
        wstBTCContract.wrap(wrapAmount);
    }

    function testWrapInsufficientBalance() public {
        uint256 wrapAmount = 20 * BTC;

        uint256 aliceBalance = stBTCContract.balanceOf(alice);
        assertEq(aliceBalance, 10 * BTC, "Alice initial balance incorrect");

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), wrapAmount);

        vm.prank(alice);
        vm.expectRevert();
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
        stBTCContract.mintRewards(0, rewardAmount);

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

        vm.expectRevert("wstBTC: Cannot unwrap zero wstBTC");
        wstBTCContract.unwrap(zeroAmount);

        vm.stopPrank();
    }

    function testUnwrapInsufficientBalance() public {
        vm.startPrank(bob);

        uint256 insufficientAmount = 1 * 10 ** 18;

        vm.expectRevert();
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

    function testIntegrationWrapUnwrapWithRebase() public {
        uint256 initialStBTC = stBTCContract.balanceOf(alice);
        uint256 wrapAmount = 5 * BTC;
        uint256 rewardAmount = 2 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), wrapAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(wrapAmount);

        uint256 initialAliceStBTCBalance = stBTCContract.balanceOf(alice);
        uint256 initialAliceWstBTCBalance = wstBTCContract.balanceOf(alice);
        uint256 initialTotalPooledBTC = stBTCContract.totalSupply();

        assertEq(initialAliceStBTCBalance, initialStBTC - wrapAmount, "Alice's stBTC balance after wrap incorrect");
        assertEq(initialAliceWstBTCBalance, wstBTCMinted, "Alice's wstBTC balance after wrap incorrect");

        stBTCContract.mintRewards(0, rewardAmount);

        uint256 totalPooledBTCAfterRebase = stBTCContract.totalSupply();
        uint256 aliceStBTCBalanceAfterRebase = stBTCContract.balanceOf(alice);

        assertGt(totalPooledBTCAfterRebase, initialTotalPooledBTC, "Total pooled BTC after rebase should increase");
        assertGt(
            aliceStBTCBalanceAfterRebase, initialAliceStBTCBalance, "Alice's stBTC balance after rebase should increase"
        );

        vm.prank(alice);
        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);

        uint256 aliceStBTCBalanceAfterUnwrap = stBTCContract.balanceOf(alice);
        uint256 finalAliceWstBTCBalance = wstBTCContract.balanceOf(alice);

        assertEq(finalAliceWstBTCBalance, 0, "Alice's wstBTC balance should be zero after unwrap");
        assertEq(
            aliceStBTCBalanceAfterUnwrap,
            initialStBTC + rewardAmount,
            "Alice's stBTC balance after unwrap should include rewards"
        );

        assertEq(
            stBTCUnwrapped,
            (wstBTCMinted * totalPooledBTCAfterRebase) / stBTCContract.totalShares(),
            "Unwrapped stBTC amount is incorrect"
        );
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
        uint256 wstBTCMinted = wstBTCContract.wrap(stBTCAmount);

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

        vm.expectRevert("wstBTC: No wstBTC supply");
        wstBTCContract.stBTCPerToken();

        vm.expectRevert("wstBTC: No stBTC balance");
        wstBTCContract.tokensPerStBTC();

        uint256 stBTCAmount = 1 * BTC;
        stBTCContract.mint(stBTCAmount, address(this), keccak256("initial_deposit"));
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        uint256 wstBTCMinted = wstBTCContract.wrap(stBTCAmount);
        assertEq(wstBTCMinted, stBTCAmount, "Minted wstBTC should equal stBTC amount");

        uint256 stBTCUnwrapped = wstBTCContract.unwrap(wstBTCMinted);
        assertEq(stBTCUnwrapped, stBTCAmount, "Unwrapped stBTC should equal initial amount");

        uint256 finalWstBTCTotalSupply = wstBTCContract.totalSupply();
        uint256 finalStBTCTotalSupply = stBTCContract.totalSupply();
        assertEq(finalWstBTCTotalSupply, 0, "wstBTC total supply after unwrap should be zero");
    }

    function testStBTCPerTokenAndTokensPerStBTC() public {
        vm.expectRevert("wstBTC: No wstBTC supply");
        wstBTCContract.stBTCPerToken();

        vm.expectRevert("wstBTC: No stBTC balance");
        wstBTCContract.tokensPerStBTC();

        uint256 stBTCAmount = 10 * BTC;

        vm.prank(alice);
        stBTCContract.approve(address(wstBTCContract), stBTCAmount);

        vm.prank(alice);
        uint256 wstBTCMinted = wstBTCContract.wrap(stBTCAmount);

        uint256 stBTCPerToken = wstBTCContract.stBTCPerToken();
        uint256 tokensPerStBTC = wstBTCContract.tokensPerStBTC();

        uint256 stBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        uint256 wstBTCSupply = wstBTCContract.totalSupply();

        uint256 expectedStBTCPerToken = (stBTCBalance * BTC) / wstBTCSupply;
        uint256 expectedTokensPerStBTC = (wstBTCSupply * BTC) / stBTCBalance;

        assertEq(stBTCPerToken, expectedStBTCPerToken, "stBTCPerToken calculation is incorrect");
        assertEq(tokensPerStBTC, expectedTokensPerStBTC, "tokensPerStBTC calculation is incorrect");

        uint256 rewardAmount = 5 * BTC;
        stBTCContract.mintRewards(0, rewardAmount);

        stBTCBalance = stBTCContract.balanceOf(address(wstBTCContract));
        wstBTCSupply = wstBTCContract.totalSupply();

        uint256 stBTCPerTokenAfterRebase = wstBTCContract.stBTCPerToken();
        uint256 tokensPerStBTCAfterRebase = wstBTCContract.tokensPerStBTC();

        expectedStBTCPerToken = (stBTCBalance * BTC) / wstBTCSupply;
        expectedTokensPerStBTC = (wstBTCSupply * BTC) / stBTCBalance;

        assertEq(stBTCPerTokenAfterRebase, expectedStBTCPerToken, "stBTCPerToken calculation after rebase is incorrect");
        assertEq(
            tokensPerStBTCAfterRebase, expectedTokensPerStBTC, "tokensPerStBTC calculation after rebase is incorrect"
        );
    }
}

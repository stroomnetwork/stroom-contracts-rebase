// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";
import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";
import {IERC4626} from "@openzeppelin/contracts/interfaces/IERC4626.sol";

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

    function testFuzzDepositCorrectAmountMinted(uint256 strBTCAmount) public {
        uint256 initialStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 initialWstrBTCBalance = wstrBTCContract.balanceOf(alice);

        vm.assume(strBTCAmount > 0 && strBTCAmount <= initialStrBTCBalance);

        uint256 expectedWstrBTC = wstrBTCContract.previewDeposit(strBTCAmount);

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        uint256 mintedWstrBTC = wstrBTCContract.deposit(strBTCAmount, alice);

        assertEq(mintedWstrBTC, expectedWstrBTC, "Incorrect amount of wstrBTC minted");

        uint256 finalStrBTCBalance = strBTCContract.balanceOf(alice);
        assertEq(finalStrBTCBalance, initialStrBTCBalance - strBTCAmount, "strBTC balance did not decrease correctly");

        uint256 finalWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(
            finalWstrBTCBalance, initialWstrBTCBalance + expectedWstrBTC, "wstrBTC balance did not increase correctly"
        );

        uint256 finalContractStrBTCBalance = strBTCContract.balanceOf(address(wstrBTCContract));
        assertEq(finalContractStrBTCBalance, strBTCAmount, "strBTC balance of wstrBTC contract incorrect");
    }

    function testDepositZeroAmountSucceeds() public {
        uint256 strBTCAmount = 0;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        uint256 shares = wstrBTCContract.deposit(strBTCAmount, alice);

        assertEq(shares, 0, "Deposit of 0 assets should return 0 shares");
        assertEq(wstrBTCContract.balanceOf(alice), 0, "Alice should have 0 wstrBTC");
    }

    function testDepositInsufficientApproval() public {
        uint256 approveAmount = 5 * BTC;
        uint256 depositAmount = 7 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), approveAmount);

        uint256 allowance = strBTCContract.allowance(alice, address(wstrBTCContract));
        assertEq(allowance, approveAmount, "Allowance incorrect");

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, address(wstrBTCContract), allowance, depositAmount
            )
        );
        wstrBTCContract.deposit(depositAmount, alice);
    }

    function testDepositInsufficientBalance() public {
        uint256 depositAmount = INITIAL_SUPPLY + 1 * BTC;

        uint256 aliceBalance = strBTCContract.balanceOf(alice);
        assertEq(aliceBalance, INITIAL_SUPPLY, "Alice initial balance incorrect");

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), depositAmount);

        vm.prank(alice);
        vm.expectRevert(strBTC.InsufficientBalance.selector);
        wstrBTCContract.deposit(depositAmount, alice);
    }

    function testRedeemSuccess() public {
        uint256 depositAmount = 5 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), depositAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.deposit(depositAmount, alice);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after deposit");

        uint256 strBTCBalanceOfWrapper = strBTCContract.balanceOf(address(wstrBTCContract));
        assertEq(strBTCBalanceOfWrapper, depositAmount, "wstrBTC contract strBTC balance incorrect");

        vm.prank(alice);
        uint256 strBTCRedeemed = wstrBTCContract.redeem(wstrBTCMinted, alice, alice);

        assertEq(strBTCRedeemed, depositAmount, "Redeemed strBTC amount incorrect");

        aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after redeem");

        uint256 aliceStrBTCBalance = strBTCContract.balanceOf(alice);
        assertEq(aliceStrBTCBalance, INITIAL_SUPPLY, "Alice's strBTC balance incorrect after redeem");
    }

    function testRedeemAfterMintRewards() public {
        uint256 initialDepositAmount = 7 * BTC + 101;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialDepositAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.deposit(initialDepositAmount, alice);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after deposit");

        uint256 rewardAmount = 5 * BTC;

        bytes32 totalSupplyUpdateHash = strBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalSupplyAfterRewards = strBTCContract.totalSupply();

        assertEq(totalSupplyAfterRewards, INITIAL_SUPPLY + rewardAmount, "Total supply incorrect after rewards");

        uint256 expectedStrBTCFromRedeem = wstrBTCContract.previewRedeem(wstrBTCMinted);

        vm.prank(alice);
        uint256 strBTCRedeemed = wstrBTCContract.redeem(wstrBTCMinted, alice, alice);

        assertEq(strBTCRedeemed, expectedStrBTCFromRedeem, "Redeemed strBTC amount incorrect after rewards");

        uint256 strBTCBalanceOfWrapper = strBTCContract.balanceOf(address(wstrBTCContract));
        assertLe(strBTCBalanceOfWrapper, 2, "strBTC balance of contract wstrBTC should be at most 2 sat");

        uint256 aliceFinalStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 expectedFinalBalance = INITIAL_SUPPLY + rewardAmount - strBTCBalanceOfWrapper;

        assertLe(expectedFinalBalance - aliceFinalStrBTCBalance, 2, "Alice's strBTC balance incorrect after redeem");

        uint256 aliceFinalWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceFinalWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after redeem");
    }

    function testRedeemZeroWstrBTC() public {
        vm.startPrank(alice);

        uint256 zeroAmount = 0;

        uint256 assets = wstrBTCContract.redeem(zeroAmount, alice, alice);

        assertEq(assets, 0, "Redeem of 0 shares should return 0 assets");

        vm.stopPrank();
    }

    function testRedeemInsufficientBalance() public {
        vm.startPrank(bob);

        uint256 insufficientAmount = 1 * 10 ** 18;

        vm.expectRevert(); // ERC20 insufficient balance error
        wstrBTCContract.redeem(insufficientAmount, bob, bob);

        vm.stopPrank();
    }

    function testRedeemCallsTransfer() public {
        uint256 initialStrBTCAmount = 5 * BTC;
        vm.startPrank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialStrBTCAmount);
        uint256 wstrBTCMinted = wstrBTCContract.deposit(initialStrBTCAmount, alice);

        uint256 aliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        assertEq(aliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance incorrect after deposit");

        vm.expectCall(
            address(strBTCContract),
            abi.encodeWithSelector(strBTCContract.transfer.selector, alice, initialStrBTCAmount)
        );
        wstrBTCContract.redeem(wstrBTCMinted, alice, alice);
        vm.stopPrank();
    }

    function testFuzzDeposit(uint256 strBTCAmount) public {
        uint256 initialAliceStrBTCAmount = strBTCContract.balanceOf(alice);
        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialContractStrBTCAmount = strBTCContract.balanceOf(address(wstrBTCContract));

        vm.assume(strBTCAmount > 0 && strBTCAmount <= initialAliceStrBTCAmount);

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.deposit(strBTCAmount, alice);

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

    function testFuzzRedeem(uint256 wstrBTCAmount) public {
        uint256 initialAliceStrBTCAmount = strBTCContract.balanceOf(alice);

        vm.startPrank(alice);
        strBTCContract.approve(address(wstrBTCContract), initialAliceStrBTCAmount);

        uint256 wstrBTCMinted = wstrBTCContract.deposit(initialAliceStrBTCAmount, alice);

        vm.assume(wstrBTCAmount > 0 && wstrBTCAmount <= wstrBTCMinted);

        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialContractStrBTCAmount = strBTCContract.balanceOf(address(wstrBTCContract));

        uint256 strBTCRedeemed = wstrBTCContract.redeem(wstrBTCAmount, alice, alice);
        vm.stopPrank();

        assertEq(strBTCContract.balanceOf(alice), strBTCRedeemed, "Alice's strBTC balance incorrect");
        assertEq(
            wstrBTCContract.balanceOf(alice),
            initialAliceWstrBTCBalance - wstrBTCAmount,
            "Alice's wstrBTC balance incorrect"
        );
        assertEq(
            strBTCContract.balanceOf(address(wstrBTCContract)),
            initialContractStrBTCAmount - strBTCRedeemed,
            "wstrBTC contract's strBTC balance incorrect"
        );
    }

    function testIntegrationDepositRedeemWithRebase(uint256 depositAmount) public {
        uint256 initialStrBTC = strBTCContract.balanceOf(alice);
        vm.assume(depositAmount > 0 && depositAmount < initialStrBTC);

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), depositAmount);

        vm.prank(alice);
        uint256 wstrBTCMinted = wstrBTCContract.deposit(depositAmount, alice);

        uint256 initialAliceStrBTCBalance = strBTCContract.balanceOf(alice);
        uint256 initialAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);
        uint256 initialTotalPooledStrBTC = strBTCContract.totalSupply();

        assertEq(
            initialAliceStrBTCBalance, initialStrBTC - depositAmount, "Alice's strBTC balance after deposit incorrect"
        );
        assertEq(initialAliceWstrBTCBalance, wstrBTCMinted, "Alice's wstrBTC balance after deposit incorrect");

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
        wstrBTCContract.redeem(wstrBTCMinted, alice, alice);

        uint256 aliceStrBTCBalanceAfterRedeem = strBTCContract.balanceOf(alice);
        uint256 finalAliceWstrBTCBalance = wstrBTCContract.balanceOf(alice);

        assertEq(finalAliceWstrBTCBalance, 0, "Alice's wstrBTC balance should be zero after redeem");

        uint256 expectedBalance = initialStrBTC + rewardAmount;

        assertApproxEqAbs(
            aliceStrBTCBalanceAfterRedeem,
            expectedBalance,
            10,
            "Alice's strBTC balance after redeem should include rewards"
        );
    }

    function testFuzzIntegrationDepositRedeemWithRebaseForTwoUsers(uint256 aliceDepositAmount, uint256 bobDepositAmount)
        public
    {
        strBTC.MintInvoice memory invoice =
            strBTC.MintInvoice({btcDepositId: keccak256("deposit2"), recipient: bob, amount: 15 * BTC});
        bytes memory signature =
            hex"afbfdb08ce06fac74578357564a061daf21b5e9e8829b052173c6de7da2db52dbe5cb8ced766d4e3d8c7a0cca7e944f29de838d36673f2b3c5ae3793c8f7d866";

        strBTCContract.mint(invoice, signature);

        uint256 totalPooledStrBTCBeforeRebase = strBTCContract.totalSupply();

        vm.assume(aliceDepositAmount > 0 && aliceDepositAmount <= strBTCContract.balanceOf(alice));
        vm.assume(bobDepositAmount > 0 && bobDepositAmount <= strBTCContract.balanceOf(bob));

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), aliceDepositAmount);

        vm.prank(bob);
        strBTCContract.approve(address(wstrBTCContract), bobDepositAmount);

        vm.prank(alice);
        uint256 aliceWstrBTCMinted = wstrBTCContract.deposit(aliceDepositAmount, alice);

        vm.prank(bob);
        uint256 bobWstrBTCMinted = wstrBTCContract.deposit(bobDepositAmount, bob);

        uint256 rewardAmount = 5 * BTC;
        bytes32 totalSupplyUpdateHash = strBTCContract.getTotalSupplyUpdateHash(0, rewardAmount);
        console.logBytes32(totalSupplyUpdateHash);
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";

        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalPooledStrBTCAfterRebase = strBTCContract.totalSupply();

        uint256 expectedAliceRedeemed =
            (aliceDepositAmount * totalPooledStrBTCAfterRebase) / totalPooledStrBTCBeforeRebase;
        uint256 expectedBobRedeemed = (bobDepositAmount * totalPooledStrBTCAfterRebase) / totalPooledStrBTCBeforeRebase;

        vm.prank(alice);
        uint256 aliceStrBTCRedeemed = wstrBTCContract.redeem(aliceWstrBTCMinted, alice, alice);

        vm.prank(bob);
        uint256 bobStrBTCRedeemed = wstrBTCContract.redeem(bobWstrBTCMinted, bob, bob);

        assertApproxEqAbs(
            aliceStrBTCRedeemed, expectedAliceRedeemed, 3, "Alice's redeemed strBTC amount incorrect after rebase"
        );

        assertApproxEqAbs(
            bobStrBTCRedeemed, expectedBobRedeemed, 3, "Bob's redeemed strBTC amount incorrect after rebase"
        );

        assertEq(wstrBTCContract.balanceOf(alice), 0, "Alice's wstrBTC balance should be zero after redeem");
        assertEq(wstrBTCContract.balanceOf(bob), 0, "Bob's wstrBTC balance should be zero after redeem");
    }

    function testZeroSupplyEdgeCase() public {
        uint256 initialWstrBTCTotalSupply = wstrBTCContract.totalSupply();
        uint256 initialWstrBTCBalance = wstrBTCContract.balanceOf(address(this));

        assertEq(initialWstrBTCTotalSupply, 0, "Initial wstrBTC total supply is not zero");
        assertEq(initialWstrBTCBalance, 0, "Initial wstrBTC balance is not zero");

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

        uint256 wstrBTCMinted = wstrBTCContract.deposit(strBTCAmount, address(this));
        assertEq(wstrBTCMinted, strBTCAmount, "Minted wstrBTC should equal strBTC amount");

        uint256 strBTCRedeemed = wstrBTCContract.redeem(wstrBTCMinted, address(this), address(this));
        assertEq(strBTCRedeemed, strBTCAmount, "Redeemed strBTC should equal initial amount");

        uint256 finalWstrBTCTotalSupply = wstrBTCContract.totalSupply();
        assertEq(finalWstrBTCTotalSupply, 0, "wstrBTC total supply after redeem should be zero");
    }

    function testConvertToAssetsAndConvertToShares() public {
        uint256 strBTCAmount = 10 * BTC;

        vm.prank(alice);
        strBTCContract.approve(address(wstrBTCContract), strBTCAmount);

        vm.prank(alice);
        wstrBTCContract.deposit(strBTCAmount, alice);

        uint256 assetsFor1Share = wstrBTCContract.convertToAssets(BTC);
        uint256 sharesFor1Asset = wstrBTCContract.convertToShares(BTC);

        assertEq(assetsFor1Share, 1e8, "convertToAssets should return 1e8 before rebase");
        assertEq(sharesFor1Asset, 1e8, "convertToShares should return 1e8 before rebase");

        uint256 rewardAmount = 5 * BTC;
        bytes memory validSignature =
            hex"061a75b07f6a29cea39b16a9d708f2b513efeaef7279f2b2105faec69e06943c5e8a8ec4c134e4298736378eed6f7f0bbd8d706f23fa0f0d1446580313a5d89b";
        strBTCContract.mintRewards(0, rewardAmount, validSignature);

        uint256 totalAssets = wstrBTCContract.totalAssets();
        uint256 totalSupply = wstrBTCContract.totalSupply();

        uint256 assetsFor1ShareAfterRebase = wstrBTCContract.convertToAssets(BTC);
        uint256 sharesFor1AssetAfterRebase = wstrBTCContract.convertToShares(BTC);

        uint256 expectedAssetsFor1Share = (totalAssets * BTC) / totalSupply;
        uint256 expectedSharesFor1Asset = (totalSupply * BTC) / totalAssets;

        assertApproxEqAbs(
            assetsFor1ShareAfterRebase, expectedAssetsFor1Share, 1, "convertToAssets incorrect after rebase"
        );

        assertApproxEqAbs(
            sharesFor1AssetAfterRebase, expectedSharesFor1Asset, 1, "convertToShares incorrect after rebase"
        );
    }

    function testAssetAndTotalAssets() public view {
        assertEq(wstrBTCContract.asset(), address(strBTCContract), "Asset should be strBTC contract");

        uint256 totalAssets = wstrBTCContract.totalAssets();
        assertEq(
            totalAssets, strBTCContract.balanceOf(address(wstrBTCContract)), "totalAssets should match vault balance"
        );
    }
}

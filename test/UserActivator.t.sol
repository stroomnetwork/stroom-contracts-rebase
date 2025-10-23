// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import {Test, console} from "forge-std/Test.sol";
import {UserActivator} from "../src/lib/UserActivator.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";
import {BTCDepositAddressDeriver} from "../lib/blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract UserActivatorTest is Test {
    UserActivator public deriver;
    address public owner;

    function setUp() public {
        owner = address(this);

        // Deploy UserActivator implementation
        UserActivator activatorImpl = new UserActivator();

        // Deploy UserActivator proxy
        bytes memory activatorData = abi.encodeWithSelector(UserActivator.initialize.selector, owner);
        TransparentUpgradeableProxy activatorProxy =
            new TransparentUpgradeableProxy(address(activatorImpl), owner, activatorData);

        deriver = UserActivator(address(activatorProxy));
    }

    function testUserActivation() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));
        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        assertEq(deriver.balanceOf(user), 1, "User should have 1 NFT");
        assertEq(deriver.getUserTokenId(user), 1, "User should have token ID 1");

        string memory btcAddr = deriver.getUserBTCAddress(user);
        assertTrue(bytes(btcAddr).length > 0, "BTC address should be generated");
    }

    function testCannotGetBTCDepositAddressIfNotActivated() public {
        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectRevert(UserActivator.UserNotActivated.selector);
        deriver.getUserBTCAddress(user);
    }

    function testUserActivationSimnet() public {
        deriver.setSeed(
            "sb1p5z8wl5tu7m0d79vzqqsl9gu0x4fkjug857fusx4fl4kfgwh5j25sxv5dv3", BitcoinNetworkEncoder.Network.Simnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));
        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        assertEq(deriver.balanceOf(user), 1, "User should have 1 NFT");
        assertEq(deriver.getUserTokenId(user), 1, "User should have token ID 1");

        string memory btcAddr = deriver.getUserBTCAddress(user);
        assertTrue(bytes(btcAddr).length > 0, "BTC address should be generated");
    }

    function testCannotSetSeedIfNotOwner() public {
        address notOwner = address(0x123);

        vm.startPrank(notOwner);

        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        vm.stopPrank();
    }

    function testCannotBypassOwnerCheckWithBaseContract() public {
        address notOwner = address(0x123);

        vm.startPrank(notOwner);

        // An attempt to bypass the check by casting to a base type
        BTCDepositAddressDeriver baseContract = BTCDepositAddressDeriver(address(deriver));

        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        baseContract.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        vm.stopPrank();
    }

    function testOwnerCanSetSeed() public {
        // By default, in tests, msg.sender is the owner
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        // Check that the seed is set
        assertTrue(deriver.wasSeedSet());
        assertEq(deriver.btcAddr(), "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3");
    }

    function testUserActivationFailsIfSeedNotSet() public {
        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectRevert(UserActivator.SeedNotSet.selector);
        deriver.activateUser(user);
    }

    function testSetDailyActivationLimit() public {
        uint256 newLimit = 200;

        address notOwner = address(0x123);
        vm.startPrank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        deriver.setDailyActivationLimit(newLimit);
        vm.stopPrank();

        vm.expectEmit(true, true, true, true);
        emit UserActivator.DailyActivationLimitUpdated(newLimit);
        deriver.setDailyActivationLimit(newLimit);

        assertEq(deriver.dailyActivationLimit(), newLimit, "Limit should be updated");
    }

    function testSetActivationLimitsEnabled() public {
        address notOwner = address(0x123);
        vm.startPrank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        deriver.setActivationLimitsEnabled(true);
        vm.stopPrank();

        vm.expectEmit(true, true, true, true);
        emit UserActivator.ActivationLimitsEnabled(true);
        deriver.setActivationLimitsEnabled(true);

        assertEq(deriver.activationLimitsEnabled(), true, "Limits should be enabled");

        vm.expectEmit(true, true, true, true);
        emit UserActivator.ActivationLimitsEnabled(false);
        deriver.setActivationLimitsEnabled(false);

        assertEq(deriver.activationLimitsEnabled(), false, "Limits should be disabled");
    }

    function testActivateUserWithLimitsDisabled() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setActivationLimitsEnabled(false);

        for (uint160 i = 1; i <= 150; i++) {
            deriver.activateUser(address(i));
        }

        for (uint160 i = 1; i <= 150; i++) {
            assertEq(deriver.balanceOf(address(i)), 1, "User should have NFT");
        }
    }

    function testActivateUserWithinLimit() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        assertEq(deriver.getRemainingActivationLimit(), 100, "Should have full limit");

        for (uint160 i = 1; i <= 50; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getRemainingActivationLimit(), 50, "Should have half limit left");
        assertEq(deriver.getUsedActivationCount(), 50, "Should track used count");
    }

    function testActivateUserExceedsLimit() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        for (uint160 i = 1; i <= 100; i++) {
            deriver.activateUser(address(i));
        }

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(101)));
    }

    function testActivationLimitResetsNextDay() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        for (uint160 i = 1; i <= 100; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getRemainingActivationLimit(), 0, "Limit should be exhausted");
        assertEq(deriver.getUsedActivationCount(), 100);

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(101)));

        vm.warp(block.timestamp + 1 days + 1);

        assertEq(deriver.getRemainingActivationLimit(), 100, "Limit should be reset");
        assertEq(deriver.getUsedActivationCount(), 0, "Used count should be reset");

        for (uint160 i = 101; i <= 150; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getRemainingActivationLimit(), 50);
    }

    function testLimitsEnabledByDefault() public view {
        assertEq(deriver.activationLimitsEnabled(), true, "Limits should be enabled by default");
        assertEq(deriver.dailyActivationLimit(), 100, "Default limit should be 100");
    }

    function testGetRemainingLimitWhenDisabled() public {
        deriver.setActivationLimitsEnabled(false);

        assertEq(deriver.getRemainingActivationLimit(), type(uint256).max);
    }

    function testCustomActivationLimit() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setDailyActivationLimit(10);

        for (uint160 i = 1; i <= 10; i++) {
            deriver.activateUser(address(i));
        }

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(11)));
    }

    function testIsActivated() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        address user = address(0x1);

        assertFalse(deriver.isActivated(user), "User should not be activated");

        deriver.activateUser(user);

        assertTrue(deriver.isActivated(user), "User should be activated");
    }

    function testNFTIsNonTransferable() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        address alice = address(0x1);
        address bob = address(0x2);

        deriver.activateUser(alice);
        uint256 tokenId = deriver.getUserTokenId(alice);

        assertEq(tokenId, 1, "Alice should have token ID 1");

        vm.prank(alice);
        vm.expectRevert(UserActivator.TokenIsNonTransferable.selector);
        deriver.transferFrom(alice, bob, tokenId);
    }

    function testNFTIsNonBurnable() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        address alice = address(0x1);

        deriver.activateUser(alice);
        uint256 tokenId = deriver.getUserTokenId(alice);
        string memory btcAddress = deriver.getUserBTCAddress(alice);

        assertEq(tokenId, 1, "Alice should have token ID 1");

        vm.prank(alice);
        vm.expectRevert();
        deriver.transferFrom(alice, address(0), tokenId);

        assertTrue(deriver.isActivated(alice), "Alice should still be activated");
        assertEq(deriver.getUserBTCAddress(alice), btcAddress, "BTC address should remain the same");
    }

    function testUserCanOnlyHaveOneNFT() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        address user = address(0x1);

        deriver.activateUser(user);
        assertEq(deriver.balanceOf(user), 1, "User should have 1 NFT");

        vm.expectRevert(UserActivator.UserAlreadyActivated.selector);
        deriver.activateUser(user);
    }

    function testMultipleUsersGetDifferentTokenIds() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setActivationLimitsEnabled(false);

        address alice = address(0x1);
        address bob = address(0x2);
        address charlie = address(0x3);

        deriver.activateUser(alice);
        deriver.activateUser(bob);
        deriver.activateUser(charlie);

        assertEq(deriver.getUserTokenId(alice), 1, "Alice should have token ID 1");
        assertEq(deriver.getUserTokenId(bob), 2, "Bob should have token ID 2");
        assertEq(deriver.getUserTokenId(charlie), 3, "Charlie should have token ID 3");

        string memory btcAlice = deriver.getUserBTCAddress(alice);
        string memory btcBob = deriver.getUserBTCAddress(bob);
        string memory btcCharlie = deriver.getUserBTCAddress(charlie);

        assertFalse(
            keccak256(bytes(btcAlice)) == keccak256(bytes(btcBob)), "Alice and Bob should have different BTC addresses"
        );
        assertFalse(
            keccak256(bytes(btcBob)) == keccak256(bytes(btcCharlie)),
            "Bob and Charlie should have different BTC addresses"
        );
    }

    function testUpgradePreservesState() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setDailyActivationLimit(200);
        deriver.setActivationLimitsEnabled(false);

        address alice = address(0x1);
        address bob = address(0x2);

        deriver.activateUser(alice);
        deriver.activateUser(bob);

        uint256 aliceTokenId = deriver.getUserTokenId(alice);
        uint256 bobTokenId = deriver.getUserTokenId(bob);
        string memory aliceBtc = deriver.getUserBTCAddress(alice);
        string memory bobBtc = deriver.getUserBTCAddress(bob);

        UserActivator newImplementation = new UserActivator();

        address proxyAddress = address(deriver);
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address proxyAdmin = address(uint160(uint256(vm.load(proxyAddress, adminSlot))));

        vm.prank(proxyAdmin);
        (bool success,) = proxyAddress.call(
            abi.encodeWithSignature("upgradeToAndCall(address,bytes)", address(newImplementation), "")
        );
        assertTrue(success, "Upgrade should succeed");

        assertEq(deriver.wasSeedSet(), true, "Seed should still be set");
        assertEq(
            deriver.btcAddr(),
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "BTC addr should be preserved"
        );
        assertEq(deriver.dailyActivationLimit(), 200, "Limit should be preserved");
        assertEq(deriver.activationLimitsEnabled(), false, "Enabled flag should be preserved");

        assertTrue(deriver.isActivated(alice), "Alice should still be activated");
        assertTrue(deriver.isActivated(bob), "Bob should still be activated");
        assertEq(deriver.getUserTokenId(alice), aliceTokenId, "Alice token ID should be preserved");
        assertEq(deriver.getUserTokenId(bob), bobTokenId, "Bob token ID should be preserved");
        assertEq(deriver.getUserBTCAddress(alice), aliceBtc, "Alice BTC address should be preserved");
        assertEq(deriver.getUserBTCAddress(bob), bobBtc, "Bob BTC address should be preserved");

        address charlie = address(0x3);
        deriver.activateUser(charlie);
        assertTrue(deriver.isActivated(charlie), "Should be able to activate new users after upgrade");
    }

    function testCannotReinitializeAfterUpgrade() public {
        UserActivator newImplementation = new UserActivator();

        address proxyAddress = address(deriver);
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address proxyAdmin = address(uint160(uint256(vm.load(proxyAddress, adminSlot))));

        vm.prank(proxyAdmin);
        (bool success,) = proxyAddress.call(
            abi.encodeWithSignature("upgradeToAndCall(address,bytes)", address(newImplementation), "")
        );
        assertTrue(success, "Upgrade should succeed");

        vm.expectRevert();
        deriver.initialize(owner);
    }

    function testUpgradeTokenIdCounterContinues() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setActivationLimitsEnabled(false);

        for (uint160 i = 1; i <= 5; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getUserTokenId(address(5)), 5, "5th user should have token ID 5");

        UserActivator newImplementation = new UserActivator();
        address proxyAddress = address(deriver);
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address proxyAdmin = address(uint160(uint256(vm.load(proxyAddress, adminSlot))));

        vm.prank(proxyAdmin);
        (bool success,) = proxyAddress.call(
            abi.encodeWithSignature("upgradeToAndCall(address,bytes)", address(newImplementation), "")
        );
        assertTrue(success);

        address newUser = address(0x6);
        deriver.activateUser(newUser);

        assertEq(deriver.getUserTokenId(newUser), 6, "New user should get token ID 6 (counter continued)");
    }

    function testUpgradeWithActiveLimits() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3", BitcoinNetworkEncoder.Network.Testnet
        );

        for (uint160 i = 1; i <= 30; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getUsedActivationCount(), 30, "Should have 30 activations used");
        assertEq(deriver.getRemainingActivationLimit(), 70, "Should have 70 remaining");

        UserActivator newImplementation = new UserActivator();
        address proxyAddress = address(deriver);
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address proxyAdmin = address(uint160(uint256(vm.load(proxyAddress, adminSlot))));

        vm.prank(proxyAdmin);
        (bool success,) = proxyAddress.call(
            abi.encodeWithSignature("upgradeToAndCall(address,bytes)", address(newImplementation), "")
        );
        assertTrue(success);

        assertEq(deriver.getUsedActivationCount(), 30, "Used count should be preserved after upgrade");
        assertEq(deriver.getRemainingActivationLimit(), 70, "Remaining limit should be preserved");

        for (uint160 i = 31; i <= 50; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getUsedActivationCount(), 50, "Should have 50 total after more activations");
    }
}

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
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));

        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        assertEq(deriver.getBTCDepositAddress(user), "tb1p8pjjwryjq9d7tke50ndcd97kqxkeztk4k85lzg7l2nynektg9zdsq836sr");
    }

    function testCannotGetBTCDepositAddressIfNotActivated() public {
        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectRevert("User is not activated");
        deriver.getBTCDepositAddress(user);
    }

    function testUserActivationSimnet() public {
        deriver.setSeed(
            "sb1p5z8wl5tu7m0d79vzqqsl9gu0x4fkjug857fusx4fl4kfgwh5j25sxv5dv3",
            "sb1pfusykjdt46ktwq03d20uqqf94uh9487344wr3q5v9szzsxnjdfkszvtlt8",
            BitcoinNetworkEncoder.Network.Simnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));

        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        //console.log("user deposit address", deriver.getBTCDepositAddress(user));
        assertEq(deriver.getBTCDepositAddress(user), "sb1pupljglwqunp2q22sahwjvmxwmxxyj3aatnhemxlneae3l03w2k5sf8r34a");
    }

    function testCannotSetSeedIfNotOwner() public {
        address notOwner = address(0x123);

        vm.startPrank(notOwner);

        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
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
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        vm.stopPrank();
    }

    function testOwnerCanSetSeed() public {
        // By default, in tests, msg.sender is the owner
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        // Check that the seed is set
        assertTrue(deriver.wasSeedSet());
        assertEq(deriver.btcAddr1(), "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3");
    }

    function testUserActivationFailsIfSeedNotSet() public {
        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectRevert("Seed must be set before activating users");
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
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setActivationLimitsEnabled(false);

        for (uint160 i = 0; i < 150; i++) {
            deriver.activateUser(address(i));
        }

        for (uint160 i = 0; i < 150; i++) {
            assertTrue(deriver.activatedAddresses(address(i)));
        }
    }

    function testActivateUserWithinLimit() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        assertEq(deriver.getRemainingActivationLimit(), 100, "Should have full limit");

        for (uint160 i = 0; i < 50; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getRemainingActivationLimit(), 50, "Should have half limit left");
        assertEq(deriver.getUsedActivationCount(), 50, "Should track used count");
    }

    function testActivateUserExceedsLimit() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        for (uint160 i = 0; i < 100; i++) {
            deriver.activateUser(address(i));
        }

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(100)));
    }

    function testActivationLimitResetsNextDay() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        for (uint160 i = 0; i < 100; i++) {
            deriver.activateUser(address(i));
        }

        assertEq(deriver.getRemainingActivationLimit(), 0, "Limit should be exhausted");
        assertEq(deriver.getUsedActivationCount(), 100);

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(100)));

        vm.warp(block.timestamp + 1 days + 1);

        assertEq(deriver.getRemainingActivationLimit(), 100, "Limit should be reset");
        assertEq(deriver.getUsedActivationCount(), 0, "Used count should be reset");

        for (uint160 i = 100; i < 150; i++) {
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
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        deriver.setDailyActivationLimit(10);

        for (uint160 i = 0; i < 10; i++) {
            deriver.activateUser(address(i));
        }

        vm.expectRevert(UserActivator.DailyActivationLimitExceeded.selector);
        deriver.activateUser(address(uint160(10)));
    }
}

// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/strBTC.sol";
import "../src/wBTCConverter.sol";
import "../src/lib/UserActivator.sol";
import "../src/lib/ValidatorRegistry.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title SetTimelockAsAdminAndOwner
 * @dev Script for transferring admin roles to timelock controller
 *
 * Environment variables:
 * - TIMELOCK_ADDRESS: address of the deployed timelock controller
 * - STRBTC_PROXY_ADDRESS: address of the strBTC proxy
 * - STRBTC_PROXY_ADMIN_ADDRESS: address of the strBTC ProxyAdmin
 * - USER_ACTIVATOR_ADDRESS: address of the UserActivator contract
 * - VALIDATOR_REGISTRY_ADDRESS: address of the ValidatorRegistry contract
 * - WBTC_CONVERTER_ADDRESS: address of the wBTCConverter proxy
 * - WBTC_CONVERTER_PROXY_ADMIN_ADDRESS: address of the wBTCConverter ProxyAdmin
 */
contract SetTimelockAsAdminAndOwner is Script {
    address timelock;
    address strBtcProxy;
    address strBtcProxyAdmin;
    address userActivator;
    address validatorRegistry;
    address wbtcConverter;
    address wbtcConverterProxyAdmin;

    function setUp() public {
        timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");
        strBtcProxy = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
        strBtcProxyAdmin = vm.envAddress("APP_ETH_STRBTC_PROXY_ADMIN_ADDRESS");
        userActivator = vm.envAddress("APP_ETH_USER_ACTIVATOR_ADDRESS");
        validatorRegistry = vm.envAddress("APP_ETH_VALIDATOR_REGISTRY_ADDRESS");
        wbtcConverter = vm.envAddress("APP_ETH_WBTC_CONVERTER_ADDRESS");
        wbtcConverterProxyAdmin = vm.envAddress("APP_ETH_WBTC_CONVERTER_PROXY_ADMIN_ADDRESS");

        console.log("=== Transfer Admin Roles to Timelock ===");
        console.log("Timelock address:", timelock);
        console.log("strBTC proxy:", strBtcProxy);
        console.log("strBTC ProxyAdmin:", strBtcProxyAdmin);
        console.log("UserActivator:", userActivator);
        console.log("ValidatorRegistry:", validatorRegistry);
        console.log("wBTCConverter:", wbtcConverter);
        console.log("wBTCConverter ProxyAdmin:", wbtcConverterProxyAdmin);
    }

    function run() external {
        vm.startBroadcast();

        // 1. Transfer strBTC admin role (DEFAULT_ADMIN_ROLE)
        _transferStrBTCAdmin(timelock);

        // 2. Transfer upgrade admin for strBTC proxy
        _transferProxyAdminOwnership(strBtcProxyAdmin, "strBTC", timelock);

        // 3. Transfer wbtcConverter admin role (DEFAULT_ADMIN_ROLE)
        _transferWBTCConverterAdmin(timelock);

        // 4. Transfer upgrade admin for wBTCConverter proxy
        _transferProxyAdminOwnership(wbtcConverterProxyAdmin, "wBTCConverter", timelock);

        // 5. Transfer UserActivator ownership
        _transferUserActivatorOwnership(timelock);

        // 6. Transfer ValidatorRegistry ownership
        _transferValidatorRegistryOwnership(timelock);

        vm.stopBroadcast();

        console.log("=== All admin roles transferred to timelock! ===");
    }

    function _transferStrBTCAdmin(address _timelock) internal {
        strBTC token = strBTC(strBtcProxy);
        bytes32 adminRole = token.DEFAULT_ADMIN_ROLE();

        console.log("Transferring strBTC admin role...");

        // Grant admin role to timelock
        token.grantRole(adminRole, _timelock);
        console.log("    Granted DEFAULT_ADMIN_ROLE to timelock");

        // Renounce admin role from current deployer
        token.renounceRole(adminRole, msg.sender);
        console.log("    Renounced DEFAULT_ADMIN_ROLE from deployer");

        // Verify transfer
        require(token.hasRole(adminRole, _timelock), "Failed to grant admin role to timelock");
        require(!token.hasRole(adminRole, msg.sender), "Failed to renounce admin role from deployer");
        console.log("    strBTC admin role successfully transferred");
    }

    function _transferProxyAdminOwnership(address proxyAdminAddress, string memory contractName, address _timelock)
        internal
    {
        console.log(string.concat("Transferring ", contractName, " ProxyAdmin ownership..."));

        ProxyAdmin proxyAdmin = ProxyAdmin(proxyAdminAddress);

        // Transfer ownership of ProxyAdmin to timelock
        proxyAdmin.transferOwnership(_timelock);
        console.log(string.concat(contractName, " ProxyAdmin ownership transferred to timelock"));

        // Verify transfer
        require(proxyAdmin.owner() == _timelock, "Failed to transfer ProxyAdmin ownership");
        console.log(string.concat(contractName, " upgrade admin successfully transferred"));
    }

    function _transferWBTCConverterAdmin(address _timelock) internal {
        WBTCConverter converter = WBTCConverter(wbtcConverter);
        bytes32 adminRole = converter.DEFAULT_ADMIN_ROLE();

        console.log("Transferring wBTCConverter admin role...");

        // Grant admin role to timelock
        converter.grantRole(adminRole, _timelock);
        console.log("    Granted DEFAULT_ADMIN_ROLE to timelock");

        // Renounce admin role from current deployer
        converter.renounceRole(adminRole, msg.sender);
        console.log("    Renounced DEFAULT_ADMIN_ROLE from deployer");

        // Verify transfer
        require(converter.hasRole(adminRole, _timelock), "Failed to grant admin role to timelock");
        require(!converter.hasRole(adminRole, msg.sender), "Failed to renounce admin role from deployer");
        console.log("    wBTCConverter admin role successfully transferred");
    }

    function _transferUserActivatorOwnership(address _timelock) internal {
        console.log("Transferring UserActivator ownership...");

        UserActivator activator = UserActivator(userActivator);

        // Transfer ownership to timelock
        activator.transferOwnership(_timelock);
        console.log("    UserActivator ownership transferred to timelock");

        // Verify transfer
        require(activator.owner() == _timelock, "Failed to transfer UserActivator ownership");
        console.log("    UserActivator ownership successfully transferred");
    }

    function _transferValidatorRegistryOwnership(address _timelock) internal {
        console.log("Transferring ValidatorRegistry ownership...");

        ValidatorRegistry registry = ValidatorRegistry(validatorRegistry);

        // Transfer ownership to timelock
        registry.transferOwnership(_timelock);
        console.log("    ValidatorRegistry ownership transferred to timelock");

        // Verify transfer
        require(registry.owner() == _timelock, "Failed to transfer ValidatorRegistry ownership");
        console.log("    ValidatorRegistry ownership successfully transferred");
    }
}

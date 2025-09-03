// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";

import "../src/strBTC.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract UpgradeScript is Script {
    function run() external {
        // Get deployed contract addresses from environment
        address strBTCProxy = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
        address proxyAdmin = vm.envAddress("APP_ETH_STRBTC_PROXY_ADMIN_ADDRESS");

        console.log("strBTC Proxy:", strBTCProxy);
        console.log("Proxy Admin:", proxyAdmin);

        vm.startBroadcast();

        // Deploy new implementations
        strBTC newStrBTCImpl = new strBTC();

        ProxyAdmin admin = ProxyAdmin(proxyAdmin);

        // Encode the reinitializeV2() function call
        bytes memory reinitializeData = abi.encodeWithSignature("reinitializeV2()");

        // Upgrade and call reinitializeV2 in one transaction
        admin.upgradeAndCall(ITransparentUpgradeableProxy(strBTCProxy), address(newStrBTCImpl), reinitializeData);

        vm.stopBroadcast();

        console.log("Upgrade completed successfully!");
        console.log("New implementation addresses:");
        console.log("strBTC:", address(newStrBTCImpl));
    }
}

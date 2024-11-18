// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Script.sol";

import "../src/stBTC.sol";
import "../src/wstBTC.sol";
import "../src/lib/UserActivator.sol";
import "../src/lib/ValidatorRegistry.sol";
import "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import { Strings as OpenZeppelinStrings } from "@openzeppelin/contracts/utils/Strings.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployScript is Script {
    BitcoinNetworkEncoder.Network private network;

    function setUp() public {
        network = BitcoinNetworkEncoder.Network(vm.envUint("BITCOIN_NETWORK"));
        console.log("Deploying with Bitcoin network:", uint(network));
    }

    function run() public {
        vm.startBroadcast();

        ValidatorRegistry vr = new ValidatorRegistry();

        // Deploy stBTC implementation
        stBTC stBtcImplementation = new stBTC();
        stBtcImplementation.initialize(network, vr);

        address owner = vm.envAddress("OWNER_ADDRESS");

        // Deploy stBTC proxy
        bytes memory stBtcData = abi.encodeWithSelector(stBTC.initialize.selector, network, vr);
        TransparentUpgradeableProxy stBtcProxy = new TransparentUpgradeableProxy(address(stBtcImplementation), owner, stBtcData);
        stBTC stBtcContract = stBTC(address(stBtcProxy));

        // Deploy wstBTC
        wstBTC wstBtcContract = new wstBTC(address(stBtcContract));

        UserActivator activator = new UserActivator();

        vm.stopBroadcast();

        console.log();
        console.log("Deployed contract addresses as envs:");
        console.log();
        console.logString(string.concat("APP_ETH_STBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(stBtcContract))))));
        console.logString(string.concat("APP_ETH_WSTBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(wstBtcContract))))));
        console.logString(string.concat("APP_ETH_VALIDATOR_REGISTRY_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(vr))))));
        console.logString(string.concat("APP_ETH_USER_ACTIVATOR_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(activator))))));
    }
}
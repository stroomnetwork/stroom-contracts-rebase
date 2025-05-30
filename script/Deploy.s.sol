// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Script.sol";

import "../src/strBTC.sol";
import "../src/wstrBTC.sol";
import "../src/lib/UserActivator.sol";
import "../src/lib/ValidatorRegistry.sol";
import "../src/wBTCConverter.sol";
import "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import {Strings as OpenZeppelinStrings} from "@openzeppelin/contracts/utils/Strings.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployScript is Script {
    BitcoinNetworkEncoder.Network private network;
    address admin = vm.envAddress("ADMIN_ADDRESS"); // timelock address
    address pauser = vm.envAddress("PAUSER_ADDRESS");
    address manager = vm.envAddress("MANAGER_ADDRESS");
    address wbtcAddress = vm.envAddress("WBTC_ADDRESS");

    function setUp() public {
        network = BitcoinNetworkEncoder.Network(vm.envUint("BITCOIN_NETWORK"));
        console.log("Deploying with Bitcoin network:", uint256(network));
    }

    function run() public {
        vm.startBroadcast();

        ValidatorRegistry vr = new ValidatorRegistry();

        // Deploy strBTC implementation
        strBTC strBtcImplementation = new strBTC();

        // Deploy strBTC proxy
        bytes memory strBtcData = abi.encodeWithSelector(strBTC.initialize.selector, network, vr, admin, pauser);
        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), admin, strBtcData);
        strBTC strBtcContract = strBTC(address(strBtcProxy));

        // Deploy wstrBTC
        wstrBTC wstrBtcContract = new wstrBTC(address(strBtcContract));

        UserActivator activator = new UserActivator();

        // Deploy wBTCConverter implementation
        WBTCConverter wBtcConverterImplementation = new WBTCConverter();

        // Deploy wBTCConverter proxy
        bytes memory wBtcConverterData = abi.encodeWithSelector(
            WBTCConverter.initialize.selector, wbtcAddress, address(strBtcContract), admin, manager, pauser
        );

        TransparentUpgradeableProxy wBtcConverterProxy =
            new TransparentUpgradeableProxy(address(wBtcConverterImplementation), admin, wBtcConverterData);
        WBTCConverter wBtcConverterContract = WBTCConverter(address(wBtcConverterProxy));

        // Grant CONVERTER_ROLE to wBTCConverter in strBTC
        strBtcContract.grantRole(strBtcContract.CONVERTER_ROLE(), address(wBtcConverterContract));

        vm.stopBroadcast();

        console.log();
        console.log("Deployed contract addresses as envs:");
        console.log();
        console.logString(
            string.concat(
                "APP_ETH_STRBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(strBtcContract))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_WSTRBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(wstrBtcContract))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_VALIDATOR_REGISTRY_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(vr))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_USER_ACTIVATOR_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(activator))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_WBTC_CONVERTER_ADDRESS=",
                OpenZeppelinStrings.toHexString(uint256(uint160(address(wBtcConverterContract))))
            )
        );
    }
}

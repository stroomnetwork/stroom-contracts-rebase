// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {WBTCConverterImmutable} from "../src/converters/wBTCConverterImmutable.sol";
import {CBBTCConverterImmutable} from "../src/converters/cbBTCConverterImmutable.sol";

contract DeployConverterScript is Script {
    function run() external {
        address wbtc = vm.envAddress("WBTC_ADDRESS");
        address cbBTC = vm.envAddress("CBBTC_ADDRESS");
        address strBTC = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
        address timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");

        vm.startBroadcast();
        WBTCConverterImmutable wbtcConverter = new WBTCConverterImmutable(wbtc, strBTC, timelock);
        CBBTCConverterImmutable cbbtcConverter = new CBBTCConverterImmutable(cbBTC, strBTC, timelock);
        vm.stopBroadcast();

        console.log("WBTCConverter deployed at:", address(wbtcConverter));
        console.log("CBBTCConverter deployed at:", address(cbbtcConverter));
        console.log("WBTC address:", wbtc);
        console.log("CBBTC address:", cbBTC);
        console.log("STRBTC address:", strBTC);
        console.log("Timelock address:", timelock);
    }
}

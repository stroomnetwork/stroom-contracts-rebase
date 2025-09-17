// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {CBBTCConverterImmutable} from "../src/converters/cbBTCConverterImmutable.sol";

contract DeployConverterScript is Script {
    function run() external {
        address cbBTC = vm.envAddress("APP_ETH_CBBTC_ADDRESS");
        address strBTC = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
        address timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");

        vm.startBroadcast();
        CBBTCConverterImmutable converter = new CBBTCConverterImmutable(cbBTC, strBTC, timelock);
        vm.stopBroadcast();

        console.log("Converter deployed at:", address(converter));
        console.log("CBBTC address:", cbBTC);
        console.log("STRBTC address:", strBTC);
        console.log("Timelock address:", timelock);
    }
}
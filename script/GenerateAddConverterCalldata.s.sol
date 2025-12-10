// script/GenerateAddConverterCalldata.s.sol
// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import "../src/lib/TimelockController.sol";
import {strBTC} from "../src/strBTC.sol";
// to run:
// forge script script/GenerateAddConverterCalldata.s.sol

contract GenerateAddConverterCalldataScript is Script {
    address public timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");
    uint256 public delay = vm.envUint("TIMELOCK_DELAY");

    address public strbtc = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
    address public converter = vm.envAddress("CONVERTER_ADDRESS");

    bytes32 public PREDECESSOR = vm.envOr("PREDECESSOR", bytes32(0));
    bytes32 public SALT = vm.envOr("SALT", keccak256("AddConverter2025"));

    function run() public view {
        console.log("=== GENERATION OF CALLDATA FOR ADD CONVERTER OPERATION ===\n");

        bytes memory addConverterCalldata = abi.encodeWithSelector(strBTC.addConverter.selector, converter);

        bytes32 operationId = keccak256(abi.encode(strbtc, 0, addConverterCalldata, PREDECESSOR, SALT));

        console.log("Function: addConverter(address)");
        console.log("Params:");
        console.log("  converter:", converter);
        console.log("");
        console.log("Target:", strbtc);
        console.log("Value: 0");
        console.log("Data:", vm.toString(addConverterCalldata));
        console.log("Predecessor:", vm.toString(PREDECESSOR));
        console.log("Salt:", vm.toString(SALT));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID:", vm.toString(operationId));
    }
}


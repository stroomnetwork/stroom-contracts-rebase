// script/governance/GenerateUniversalCalldata.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "../src/lib/TimelockController.sol";

contract GenerateUniversalCalldataScript is Script {
    address public timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");
    uint256 public delay = vm.envUint("TIMELOCK_DELAY");

    address public targetContract = vm.envAddress("TARGET_CONTRACT");
    string public functionSignature = vm.envString("FUNCTION_SIGNATURE");
    string public functionParams = vm.envString("FUNCTION_PARAMS"); // JSON-like format

    // look at .env_example for example values
    // then run:
    // forge script script/GenerateCalldata.s.sol

    bytes32 public PREDECESSOR = vm.envOr("PREDECESSOR", bytes32(0)); // this is the operation ID of another operation that MUST be performed before
    bytes32 public SALT = vm.envOr("SALT", bytes32(0)); // this is a random value that is used to make the operation ID unique

    function run() public view {
        bytes memory functionCalldata = _generateFunctionCalldata();

        bytes32 operationId = keccak256(abi.encode(targetContract, 0, functionCalldata, PREDECESSOR, SALT));

        _printResults(functionCalldata, operationId);
    }

    function _generateFunctionCalldata() internal view returns (bytes memory) {
        bytes4 selector = bytes4(keccak256(bytes(functionSignature)));

        bytes memory encodedParams = _parseAndEncodeParams();

        return abi.encodePacked(selector, encodedParams);
    }

    function _parseAndEncodeParams() internal view returns (bytes memory) {
        string[] memory params = _splitString(functionParams, ",");

        if (params.length == 0) {
            return "";
        }

        if (params.length == 1) {
            return _encodeParam(params[0]);
        }

        if (params.length == 2) {
            return abi.encode(_parseValue(params[0]), _parseValue(params[1]));
        }

        if (params.length == 3) {
            return abi.encode(_parseValue(params[0]), _parseValue(params[1]), _parseValue(params[2]));
        }

        revert("Unsupported number of parameters");
    }

    function _encodeParam(string memory param) internal pure returns (bytes memory) {
        return abi.encode(_parseValue(param));
    }

    function _parseValue(string memory value) internal pure returns (bytes32) {
        value = _trim(value);

        if (_startsWith(value, "0x")) {
            if (bytes(value).length == 42) {
                address addr = vm.parseAddress(value);
                return bytes32(uint256(uint160(addr)));
            } else {
                return vm.parseBytes32(value);
            }
        }

        return bytes32(vm.parseUint(value));
    }

    function _splitString(string memory str, string memory delimiter) internal pure returns (string[] memory) {
        if (bytes(str).length == 0) {
            string[] memory empty = new string[](0);
            return empty;
        }

        // Simple implementation for comma-separated values
        // This is basic - extend for more complex parsing
        bytes memory strBytes = bytes(str);
        bytes memory delimBytes = bytes(delimiter);

        // Count delimiters
        uint256 count = 1;
        for (uint256 i = 0; i < strBytes.length; i++) {
            if (strBytes[i] == delimBytes[0]) {
                count++;
            }
        }

        string[] memory result = new string[](count);
        uint256 currentIndex = 0;
        uint256 start = 0;

        for (uint256 i = 0; i <= strBytes.length; i++) {
            if (i == strBytes.length || strBytes[i] == delimBytes[0]) {
                bytes memory part = new bytes(i - start);
                for (uint256 j = start; j < i; j++) {
                    part[j - start] = strBytes[j];
                }
                result[currentIndex] = string(part);
                currentIndex++;
                start = i + 1;
            }
        }

        return result;
    }

    function _trim(string memory str) internal pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        uint256 start = 0;
        uint256 end = strBytes.length;

        // Remove leading spaces
        while (start < end && strBytes[start] == 0x20) {
            start++;
        }

        // Remove trailing spaces
        while (end > start && strBytes[end - 1] == 0x20) {
            end--;
        }

        bytes memory result = new bytes(end - start);
        for (uint256 i = start; i < end; i++) {
            result[i - start] = strBytes[i];
        }

        return string(result);
    }

    function _startsWith(string memory str, string memory prefix) internal pure returns (bool) {
        bytes memory strBytes = bytes(str);
        bytes memory prefixBytes = bytes(prefix);

        if (strBytes.length < prefixBytes.length) {
            return false;
        }

        for (uint256 i = 0; i < prefixBytes.length; i++) {
            if (strBytes[i] != prefixBytes[i]) {
                return false;
            }
        }

        return true;
    }

    function _printResults(
        bytes memory functionCalldata,
        bytes32 operationId
    ) internal view {
        console.log("OPERATION SUMMARY:");
        console.log("Timelock Address:", timelock);
        console.log("");
        console.log("Function Signature:", functionSignature);
        console.log("Function Parameters:", functionParams);
        console.log("");
        console.log("Target:", targetContract);
        console.log("Value: 0");
        console.log("Data:", vm.toString(functionCalldata));
        console.log("Predecessor:", vm.toString(PREDECESSOR));
        console.log("Salt:", vm.toString(SALT));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID: ", vm.toString(operationId));
    }
}

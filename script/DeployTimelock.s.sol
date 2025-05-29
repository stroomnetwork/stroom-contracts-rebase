// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/lib/TimelockController.sol";

/**
 * @title DeployTimelockScript
 * @dev Script for deploying StroomTimelockController
 *
 * Environment variables:
 * - ADMIN_ADDRESS: address that will receive DEFAULT_ADMIN_ROLE (optional)
 * - PROPOSER_ADDRESSES: addresses that will receive PROPOSER_ROLE (comma separated)
 * - EXECUTOR_ADDRESSES: addresses that will receive EXECUTOR_ROLE (comma separated)
 * - TIMELOCK_DELAY: delay in seconds (default: 1 day)
 */
contract DeployTimelockScript is Script {
    function run() external {
        // Get addresses from environment
        address admin = vm.envOr("ADMIN_ADDRESS", address(0));
        string memory proposersStr = vm.envString("PROPOSER_ADDRESSES");
        string memory executorsStr = vm.envString("EXECUTOR_ADDRESSES");

        // Get delay (default: 1 day)
        uint256 delay = vm.envOr("TIMELOCK_DELAY", uint256(1 days));

        // Parse proposers addresses
        address[] memory proposers = _parseAddresses(proposersStr);
        console.log("Number of proposers:", proposers.length);
        for (uint256 i = 0; i < proposers.length; i++) {
            console.log("Proposer", i, ":", proposers[i]);
        }

        // Parse executors addresses
        address[] memory executors = _parseAddresses(executorsStr);
        console.log("Number of executors:", executors.length);
        for (uint256 i = 0; i < executors.length; i++) {
            console.log("Executor", i, ":", executors[i]);
        }

        console.log("Timelock delay:", delay, "seconds");
        console.log("Admin (if exists):", admin);

        // Start deployment transaction
        vm.startBroadcast();

        StroomTimelockController timelock = new StroomTimelockController(delay, proposers, executors, admin);

        vm.stopBroadcast();

        console.log("--------------------------------------------------");
        console.log("StroomTimelockController deployed at:", address(timelock));
        console.log("--------------------------------------------------");
    }

    /**
     * @dev Parses a string of addresses separated by commas
     * @param addressesStr String of addresses (0x123,0x456,...)
     * @return Array of addresses
     */
    function _parseAddresses(string memory addressesStr) private pure returns (address[] memory) {
        uint256 count = 1;
        for (uint256 i = 0; i < bytes(addressesStr).length; i++) {
            if (bytes(addressesStr)[i] == ",") {
                count++;
            }
        }

        address[] memory addresses = new address[](count);

        uint256 index = 0;
        uint256 startPos = 0;
        for (uint256 i = 0; i <= bytes(addressesStr).length; i++) {
            if (i == bytes(addressesStr).length || bytes(addressesStr)[i] == ",") {
                string memory addrStr = _substring(addressesStr, startPos, i - startPos);
                addresses[index] = _parseHexAddress(addrStr);
                index++;
                startPos = i + 1;
            }
        }

        return addresses;
    }

    /**
     * @dev Extracts a substring from a string
     * @param str Input string
     * @param startIndex Starting index
     * @param length Length of substring
     */
    function _substring(string memory str, uint256 startIndex, uint256 length) private pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = strBytes[startIndex + i];
        }
        return string(result);
    }

    /**
     * @dev Parses an address in hex format
     * @param hexStr String of address in hex format
     */
    function _parseHexAddress(string memory hexStr) private pure returns (address) {
        bytes memory strBytes = bytes(hexStr);
        uint256 result = 0;
        uint256 start = 0;

        // Skip "0x" if it exists
        if (strBytes.length >= 2 && strBytes[0] == "0" && strBytes[1] == "x") {
            start = 2;
        }

        // Convert each character
        for (uint256 i = start; i < strBytes.length; i++) {
            result *= 16;
            uint8 digit = uint8(strBytes[i]);

            // 0-9
            if (digit >= 48 && digit <= 57) {
                result += (digit - 48);
            }
            // A-F
            else if (digit >= 65 && digit <= 70) {
                result += (digit - 55);
            }
            // a-f
            else if (digit >= 97 && digit <= 102) {
                result += (digit - 87);
            }
        }

        return address(uint160(result));
    }
}

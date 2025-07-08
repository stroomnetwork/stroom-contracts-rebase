// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/lib/TimelockController.sol";
import {Strings as OpenZeppelinStrings} from "@openzeppelin/contracts/utils/Strings.sol";

/**
 * @title DeployTimelockScript
 * @dev Script for deploying StroomTimelockController
 *
 * Environment variables:
 * - ADMIN_ADDRESS: address that will receive DEFAULT_ADMIN_ROLE (optional)
 * - PROPOSER_ADDRESS: address that will receive PROPOSER_ROLE
 * - EXECUTOR_ADDRESS: address that will receive EXECUTOR_ROLE
 * - TIMELOCK_DELAY: delay in seconds (default: 1 day)
 */
contract DeployTimelockScript is Script {
    function run() external {
        // Get addresses from environment
        address admin = vm.envAddress("TIMELOCK_MULTISIG_ADMIN_ADDRESS");
        address[] memory proposer = new address[](1);
        proposer[0] = vm.envAddress("PROPOSER_ADDRESS");
        address[] memory executor = new address[](1);
        executor[0] = admin;

        uint256 delay = vm.envOr("TIMELOCK_DELAY", uint256(1 days));

        console.log("Timelock delay:", delay, "seconds");
        console.log("Proposer:", proposer[0]);
        console.log("Executor:", executor[0]);
        console.log("Admin (if exists):", admin);

        // Start deployment transaction
        vm.startBroadcast();

        StroomTimelockController timelock = new StroomTimelockController(delay, proposer, executor, admin);

        vm.stopBroadcast();

        console.log("--------------------------------------------------");
        console.logString(
            string.concat(
                "APP_ETH_TIMELOCK_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(timelock))))
            )
        );
        console.log("--------------------------------------------------");
    }
}

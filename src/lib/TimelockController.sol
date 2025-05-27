// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title StroomTimelockController
 * @notice This contract inherits the TimelockController from OpenZeppelin
 * and allows you to add custom logic as needed
 */
contract StroomTimelockController is TimelockController {
    /**
     * @dev Constructor for initializing the StroomTimelockController
     * @param minDelay Minimum delay for operations (in seconds)
     * @param proposers Array of addresses that will receive the PROPOSER_ROLE role
     * @param executors Array of addresses that will receive the EXECUTOR_ROLE role
     * @param admin Address that will receive the DEFAULT_ADMIN_ROLE
     */
    constructor(uint256 minDelay, address[] memory proposers, address[] memory executors, address admin)
        TimelockController(minDelay, proposers, executors, admin)
    {}

    /**
     * @dev Version of the contract for identification
     * @return String with the contract version
     */
    function version() external pure returns (string memory) {
        return "1.0.0";
    }
}

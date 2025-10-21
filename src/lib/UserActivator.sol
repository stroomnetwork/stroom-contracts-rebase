// SPDX-License-Identifier: MIT

pragma solidity 0.8.27;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract UserActivator is Initializable, BTCDepositAddressDeriver, OwnableUpgradeable {
    error DailyActivationLimitExceeded();

    event UserAddressActivated(address userETHAddress);
    event DailyActivationLimitUpdated(uint256 newLimit);
    event ActivationLimitsEnabled(bool enabled);

    struct DailyActivationTracker {
        uint256 count;
        uint256 lastResetDay;
    }

    // Mapping of activated addresses.
    mapping(address => bool) public activatedAddresses;

    uint256 public dailyActivationLimit;
    bool public activationLimitsEnabled;
    DailyActivationTracker private dailyActivations;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Initializes the contract
     * @param _owner The owner of the contract
     */
    function initialize(address _owner) public initializer {
        __Ownable_init(_owner);

        // Initialize BTCDepositAddressDeriver state
        wasSeedSet = false;

        // Initialize UserActivator state
        dailyActivationLimit = 100;
        activationLimitsEnabled = false; // disabled by default
    }

    /**
     * @dev Activates a user by setting their address as activated.
     * @param _userAddress The address of the user to be activated.
     * Emits a `UserAddressActivated` event with the user address and their BTC deposit address.
     * Reverts if the user address is already activated.
     */
    function activateUser(address _userAddress) public {
        require(activatedAddresses[_userAddress] == false, "User is already activated");
        require(wasSeedSet, "Seed must be set before activating users");

        if (activationLimitsEnabled) {
            _checkAndUpdateActivationLimit();
        }

        activatedAddresses[_userAddress] = true;

        emit UserAddressActivated(_userAddress);
    }

    /**
     * @dev Sets the seed for the BTC deposit address deriver.
     * @param _btcAddr1 The first BTC taproot address.
     * @param _btcAddr2 The second BTC taproot address.
     * @param _network The network of the BTC addresses.
     */
    function setSeed(string calldata _btcAddr1, string calldata _btcAddr2, BitcoinNetworkEncoder.Network _network)
        public
        override
        onlyOwner
    {
        BTCDepositAddressDeriver.setSeed(_btcAddr1, _btcAddr2, _network);
    }

    function getBTCDepositAddress(address ethAddress) public view override returns (string memory) {
        if (!activatedAddresses[ethAddress]) {
            revert("User is not activated");
        }
        return BTCDepositAddressDeriver.getBTCDepositAddress(ethAddress);
    }

    /**
     * @notice Sets the daily activation limit
     * @param _limit The new daily limit
     * @dev Only the contract owner can call this function
     */
    function setDailyActivationLimit(uint256 _limit) external onlyOwner {
        dailyActivationLimit = _limit;
        emit DailyActivationLimitUpdated(_limit);
    }

    /**
     * @notice Enables or disables activation limits
     * @param _enabled Whether to enable activation limits
     * @dev Only the contract owner can call this function
     */
    function setActivationLimitsEnabled(bool _enabled) external onlyOwner {
        activationLimitsEnabled = _enabled;
        emit ActivationLimitsEnabled(_enabled);
    }

    /**
     * @notice Returns the remaining activation limit for today
     * @return The remaining number of activations allowed today
     */
    function getRemainingActivationLimit() public view returns (uint256) {
        if (!activationLimitsEnabled) return type(uint256).max;

        uint256 currentDay = block.timestamp / 1 days;

        if (dailyActivations.lastResetDay < currentDay) {
            return dailyActivationLimit;
        }

        if (dailyActivations.count >= dailyActivationLimit) {
            return 0;
        }

        return dailyActivationLimit - dailyActivations.count;
    }

    /**
     * @notice Returns the used activation count for today
     * @return The number of activations already used today
     */
    function getUsedActivationCount() public view returns (uint256) {
        uint256 currentDay = block.timestamp / 1 days;

        if (dailyActivations.lastResetDay < currentDay) {
            return 0;
        }

        return dailyActivations.count;
    }

    /**
     * @dev Checks and updates activation limit
     */
    function _checkAndUpdateActivationLimit() internal {
        uint256 currentDay = block.timestamp / 1 days;

        if (dailyActivations.lastResetDay < currentDay) {
            dailyActivations.count = 0;
            dailyActivations.lastResetDay = currentDay;
        }

        if (dailyActivations.count >= dailyActivationLimit) {
            revert DailyActivationLimitExceeded();
        }

        dailyActivations.count += 1;
    }
}

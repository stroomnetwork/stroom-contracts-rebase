// SPDX-License-Identifier: MIT

pragma solidity 0.8.27;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {
    ERC721EnumerableUpgradeable
} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract UserActivator is Initializable, BTCDepositAddressDeriver, ERC721EnumerableUpgradeable, OwnableUpgradeable {
    error DailyActivationLimitExceeded();
    error UserAlreadyActivated();
    error TokenIsNonTransferable();
    error SeedNotSet();
    error UserNotActivated();

    event UserAddressActivated(address userETHAddress);
    event DailyActivationLimitUpdated(uint256 newLimit);
    event ActivationLimitsEnabled(bool enabled);

    struct DailyActivationTracker {
        uint256 count;
        uint256 lastResetDay;
    }

    uint256 public dailyActivationLimit;
    bool public activationLimitsEnabled;
    DailyActivationTracker private dailyActivations;

    uint256 private _nextTokenId;

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
        __ERC721_init("Stroom Activation NFT", "strNFT");
        __ERC721Enumerable_init();
        __Ownable_init(_owner);

        // Initialize BTCDepositAddressDeriver state
        wasSeedSet = false;

        // Initialize UserActivator state
        dailyActivationLimit = 100;
        activationLimitsEnabled = true; // enabled by default
        _nextTokenId = 1; // Start from token ID 1
    }

    /**
     * @dev Activates a user by minting them a non-transferable NFT
     * @param _userAddress The address of the user to be activated
     * Emits a `UserAddressActivated` event with the user address
     * Reverts if the user address is already activated
     */
    function activateUser(address _userAddress) public {
        if (balanceOf(_userAddress) != 0) revert UserAlreadyActivated();
        if (!wasSeedSet) revert SeedNotSet();

        if (activationLimitsEnabled) {
            _checkAndUpdateActivationLimit();
        }

        uint256 tokenId = _nextTokenId;
        _nextTokenId++;

        _safeMint(_userAddress, tokenId);

        emit UserAddressActivated(_userAddress);
    }

    /**
     * @dev Sets the seed for the BTC deposit address deriver
     * @param _btcAddr The BTC taproot address
     * @param _network The network of the BTC address
     */
    function setSeed(string calldata _btcAddr, BitcoinNetworkEncoder.Network _network) public override onlyOwner {
        BTCDepositAddressDeriver.setSeed(_btcAddr, _network);
    }

    /**
     * @notice Check if a user is activated
     * @param userAddress The user's Ethereum address
     * @return True if user is activated, false otherwise
     */
    function isActivated(address userAddress) public view returns (bool) {
        return balanceOf(userAddress) > 0;
    }

    /**
     * @notice Gets the token ID for a user address
     * @param userAddress The user's Ethereum address
     * @return The NFT token ID (0 if user is not activated)
     */
    function getUserTokenId(address userAddress) public view returns (uint256) {
        uint256 balance = balanceOf(userAddress);
        if (balance == 0) {
            return 0;
        }

        // User has exactly 1 NFT, get the token ID
        return tokenOfOwnerByIndex(userAddress, 0);
    }

    /**
     * @notice Gets the BTC deposit address for a user
     * @param userAddress The user's Ethereum address
     * @return The Bitcoin deposit address
     */
    function getUserBTCAddress(address userAddress) public view returns (string memory) {
        uint256 tokenId = getUserTokenId(userAddress);
        if (tokenId == 0) revert UserNotActivated();

        return BTCDepositAddressDeriver.getBTCDepositAddress(tokenId);
    }

    /**
     * @notice Override to make tokens non-transferable and non-burnable (soulbound)
     * @dev Only minting is allowed. Transfers and burns are blocked to preserve BTC address mapping
     */
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721EnumerableUpgradeable)
        returns (address)
    {
        address previousOwner = super._update(to, tokenId, auth);

        // Only allow minting (previousOwner == address(0))
        // Note: to != address(0) is already validated by OpenZeppelin's _mint()
        // Block transfers and burns (previousOwner != address(0))
        if (previousOwner != address(0)) {
            revert TokenIsNonTransferable();
        }

        return previousOwner;
    }

    /**
     * @dev Override required by Solidity for ERC721Enumerable
     */
    function supportsInterface(bytes4 interfaceId) public view override(ERC721EnumerableUpgradeable) returns (bool) {
        return super.supportsInterface(interfaceId);
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

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "./IStrBTC.sol";

/**
 * @title WBTCConverter
 * @notice Contract for conversion between WBTC and strBTC at dynamic rates
 * @dev Allows users to exchange WBTC for strBTC and vice versa with separate rates
 */
contract WBTCConverter is PausableUpgradeable, AccessControlUpgradeable {
    error AmountMustBeGreaterThanZero();
    error ConversionResultedInZeroTokens();
    error InsufficientWBTCBalance();
    error NumeratorMustBeGreaterThanZero();
    error DenominatorMustBeGreaterThanZero();
    error MintingLimitExceeded();

    IStrBTC public strbtc;
    IERC20 public wbtc;

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    // Incoming exchange rate (WBTC -> strBTC)
    uint256 public incomingRateNumerator; // strBTC
    uint256 public incomingRateDenominator; // WBTC

    // Outgoing exchange rate (strBTC -> WBTC)
    uint256 public outgoingRateNumerator; // strBTC
    uint256 public outgoingRateDenominator; // WBTC

    uint256 public totalMinted;
    uint256 public totalBurned;
    uint256 public mintingLimit;

    event WBTCConverted(address indexed user, uint256 wbtcAmount, uint256 strbtcAmount);
    event StrBTCConverted(address indexed user, uint256 strbtcAmount, uint256 wbtcAmount);
    event IncomingRateUpdated(uint256 numerator, uint256 denominator, address updater);
    event OutgoingRateUpdated(uint256 numerator, uint256 denominator, address updater);
    event MintingLimitUpdated(uint256 newLimit, address updater);

    /**
     * @notice Initializes the WBTCConverter contract
     * @param _wbtc wBTC Token Address
     * @param _strbtc strBTC Token Address
     */
    function initialize(address _wbtc, address _strbtc, address admin, address manager, address pauser)
        external
        initializer
    {
        PausableUpgradeable.__Pausable_init();
        AccessControlUpgradeable.__AccessControl_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(MANAGER_ROLE, manager);
        _grantRole(PAUSER_ROLE, pauser);

        wbtc = IERC20(_wbtc);
        strbtc = IStrBTC(_strbtc);

        // Initial rates 1:1
        incomingRateNumerator = 1;
        incomingRateDenominator = 1;
        outgoingRateNumerator = 1;
        outgoingRateDenominator = 1;

        totalMinted = 0;
        totalBurned = 0;
        mintingLimit = type(uint256).max;
    }

    /**
     * @notice Returns the current amount of minted strBTC through this contract
     * @return The net amount currently minted (totalMinted - totalBurned)
     */
    function currentlyMinted() public view returns (uint256) {
        return totalMinted - totalBurned;
    }

    /**
     * @notice Converts wBTC to strBTC at the incoming rate
     * @param wbtcAmount The amount of wBTC to convert
     * @return The amount of strBTC received
     */
    function convertWBTCToStrBTC(uint256 wbtcAmount) external whenNotPaused returns (uint256) {
        if (wbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 strbtcAmount = (wbtcAmount * incomingRateNumerator) / incomingRateDenominator;
        if (strbtcAmount == 0) revert ConversionResultedInZeroTokens();

        if (currentlyMinted() + strbtcAmount > mintingLimit) revert MintingLimitExceeded();

        wbtc.transferFrom(msg.sender, address(this), wbtcAmount);

        strbtc.converterMint(msg.sender, strbtcAmount);

        totalMinted += strbtcAmount;

        emit WBTCConverted(msg.sender, wbtcAmount, strbtcAmount);

        return strbtcAmount;
    }

    /**
     * @notice Converts strBTC back to wBTC at the outgoing rate
     * @param strbtcAmount The amount of strBTC to convert
     * @return The amount of wBTC received
     */
    function convertStrBTCToWBTC(uint256 strbtcAmount) external whenNotPaused returns (uint256) {
        if (strbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 wbtcAmount = (strbtcAmount * outgoingRateDenominator) / outgoingRateNumerator;
        if (wbtcAmount == 0) revert ConversionResultedInZeroTokens();

        if (wbtc.balanceOf(address(this)) < wbtcAmount) revert InsufficientWBTCBalance();

        strbtc.transferFrom(msg.sender, address(this), strbtcAmount);

        strbtc.converterBurn(address(this), strbtcAmount);

        totalBurned += strbtcAmount;

        wbtc.transfer(msg.sender, wbtcAmount);

        emit StrBTCConverted(msg.sender, strbtcAmount, wbtcAmount);

        return wbtcAmount;
    }

    /**
     * @notice Updates the incoming exchange rate (WBTC -> strBTC)
     * @param numerator The numerator of the exchange rate (number of strBTC)
     * @param denominator The denominator of the exchange rate (number of wBTC)
     */
    function setIncomingRate(uint256 numerator, uint256 denominator) external onlyRole(MANAGER_ROLE) {
        if (numerator == 0) revert NumeratorMustBeGreaterThanZero();
        if (denominator == 0) revert DenominatorMustBeGreaterThanZero();

        incomingRateNumerator = numerator;
        incomingRateDenominator = denominator;

        emit IncomingRateUpdated(numerator, denominator, msg.sender);
    }

    /**
     * @notice Updates the outgoing exchange rate (strBTC -> WBTC)
     * @param numerator The numerator of the exchange rate (number of strBTC)
     * @param denominator The denominator of the exchange rate (number of wBTC)
     */
    function setOutgoingRate(uint256 numerator, uint256 denominator) external onlyRole(MANAGER_ROLE) {
        if (numerator == 0) revert NumeratorMustBeGreaterThanZero();
        if (denominator == 0) revert DenominatorMustBeGreaterThanZero();

        outgoingRateNumerator = numerator;
        outgoingRateDenominator = denominator;

        emit OutgoingRateUpdated(numerator, denominator, msg.sender);
    }

    /**
     * @notice Updates both incoming and outgoing rates to the same value
     * @param numerator The numerator of the exchange rate (number of strBTC)
     * @param denominator The denominator of the exchange rate (number of wBTC)
     */
    function setCommonExchangeRate(uint256 numerator, uint256 denominator) external onlyRole(MANAGER_ROLE) {
        if (numerator == 0) revert NumeratorMustBeGreaterThanZero();
        if (denominator == 0) revert DenominatorMustBeGreaterThanZero();

        incomingRateNumerator = numerator;
        incomingRateDenominator = denominator;
        outgoingRateNumerator = numerator;
        outgoingRateDenominator = denominator;

        emit IncomingRateUpdated(numerator, denominator, msg.sender);
        emit OutgoingRateUpdated(numerator, denominator, msg.sender);
    }

    /**
     * @notice Sets the minting limit for strBTC through this converter
     * @param newLimit The new limit for total minted - total burned
     */
    function setMintingLimit(uint256 newLimit) external onlyRole(MANAGER_ROLE) {
        mintingLimit = newLimit;
        emit MintingLimitUpdated(newLimit, msg.sender);
    }

    /**
     * @notice Pauses all conversions
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice Resumes all conversions
     */
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @notice Emergency withdrawal of tokens from the contract
     * @param token The address of the token to withdraw
     * @param amount The amount of tokens
     */
    function emergencyWithdraw(address token, uint256 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        IERC20(token).transfer(msg.sender, amount);
    }
}

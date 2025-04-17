// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "./strBTC.sol";

/**
 * @title WBTCConverter
 * @notice Contract for conversion between WBTC and strBTC at a dynamic rate
 * @dev Allows users to exchange WBTC for strBTC and vice versa
 */
contract WBTCConverter is PausableUpgradeable, AccessControlUpgradeable {
    error AmountMustBeGreaterThanZero();
    error ConversionResultedInZeroTokens();
    error InsufficientWBTCBalance();
    error NumeratorMustBeGreaterThanZero();
    error DenominatorMustBeGreaterThanZero();

    strBTC public strbtc;
    IERC20 public wbtc;

    bytes32 public constant RATE_SETTER_ROLE = keccak256("RATE_SETTER_ROLE");

    // The exchange rate is represented as the ratio X/Y
    uint256 public exchangeRateNumerator; // X (strBTC)
    uint256 public exchangeRateDenominator; // Y (wBTC)

    event WBTCConverted(address indexed user, uint256 wbtcAmount, uint256 strbtcAmount);
    event StrBTCConverted(address indexed user, uint256 strbtcAmount, uint256 wbtcAmount);
    event ExchangeRateUpdated(uint256 numerator, uint256 denominator, address updater);

    /**
     * @notice Initializes the WBTCConverter contract
     * @param _wbtc wBTC Token Addresses
     * @param _strbtc strBTC Token Addresses
     */
    function initialize(address _wbtc, address _strbtc) external initializer {
        PausableUpgradeable.__Pausable_init();
        AccessControlUpgradeable.__AccessControl_init();

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(RATE_SETTER_ROLE, msg.sender);

        wbtc = IERC20(_wbtc);
        strbtc = strBTC(_strbtc);

        // Initial rate 1:1
        exchangeRateNumerator = 1;
        exchangeRateDenominator = 1;
    }

    /**
     * @notice Converts wBTC to strBTC at the current rate
     * @param wbtcAmount The amount of wBTC to convert
     * @return The amount of strBTC received
     */
    function convertWBTCToStrBTC(uint256 wbtcAmount) external whenNotPaused returns (uint256) {
        if (wbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 strbtcAmount = (wbtcAmount * exchangeRateNumerator) / exchangeRateDenominator;
        if (strbtcAmount == 0) revert ConversionResultedInZeroTokens();

        wbtc.transferFrom(msg.sender, address(this), wbtcAmount);

        strbtc.converterMint(msg.sender, strbtcAmount);

        emit WBTCConverted(msg.sender, wbtcAmount, strbtcAmount);

        return strbtcAmount;
    }

    /**
     * @notice Converts strBTC back to wBTC at the current rate
     * @param strbtcAmount The amount of strBTC to convert
     * @return The amount of wBTC received
     */
    function convertStrBTCToWBTC(uint256 strbtcAmount) external whenNotPaused returns (uint256) {
        if (strbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 wbtcAmount = (strbtcAmount * exchangeRateDenominator) / exchangeRateNumerator;
        if (wbtcAmount == 0) revert ConversionResultedInZeroTokens();

        if (wbtc.balanceOf(address(this)) < wbtcAmount) revert InsufficientWBTCBalance();

        strbtc.transferFrom(msg.sender, address(this), strbtcAmount);

        strbtc.converterBurn(address(this), strbtcAmount);

        wbtc.transfer(msg.sender, wbtcAmount);

        emit StrBTCConverted(msg.sender, strbtcAmount, wbtcAmount);

        return wbtcAmount;
    }

    /**
     * @notice Updates the exchange rate
     * @param numerator The numerator of the exchange rate (number of strBTC)
     * @param denominator The denominator of the exchange rate (number of wBTC)
     */
    function setExchangeRate(uint256 numerator, uint256 denominator) external onlyRole(RATE_SETTER_ROLE) {
        if (numerator == 0) revert NumeratorMustBeGreaterThanZero();
        if (denominator == 0) revert DenominatorMustBeGreaterThanZero();

        exchangeRateNumerator = numerator;
        exchangeRateDenominator = denominator;

        emit ExchangeRateUpdated(numerator, denominator, msg.sender);
    }

    /**
     * @notice Pauses all conversions
     */
    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    /**
     * @notice Resumes all conversions
     */
    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
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

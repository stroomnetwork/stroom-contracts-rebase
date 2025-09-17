// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../IStrBTC.sol";

/**
 * @title WBTCConverterImmutable
 * @notice Contract for conversion between WBTC and strBTC at immutable rates
 * @dev Allows users to exchange WBTC for strBTC and vice versa with separate rates
 */
contract WBTCConverterImmutable {
    error AmountMustBeGreaterThanZero();
    error ConversionResultedInZeroTokens();
    error InsufficientWBTCBalance();
    error MintingLimitExceeded();
    error NotWithdrawer();

    IStrBTC public immutable strbtc;
    IERC20 public immutable wbtc;
    address public immutable withdrawer;
    uint256 public immutable mintingLimit;

    uint256 public immutable rateNumerator;
    uint256 public immutable rateDenominator;

    uint256 public totalMinted;
    uint256 public totalBurned;

    event WBTCConverted(address indexed user, uint256 wbtcAmount, uint256 strbtcAmount);
    event StrBTCConverted(address indexed user, uint256 strbtcAmount, uint256 wbtcAmount);

    modifier onlyWithdrawer() {
        if (msg.sender != withdrawer) revert NotWithdrawer();
        _;
    }

    /**
     * @notice constructor of the WBTCConverterImmutable contract
     * @param _wbtc wBTC Token Address
     * @param _strbtc strBTC Token Address
     */
    constructor(address _wbtc, address _strbtc, address _withdrawer) {
        wbtc = IERC20(_wbtc);
        strbtc = IStrBTC(_strbtc);

        withdrawer = _withdrawer;

        rateNumerator = 999;
        rateDenominator = 1000;

        mintingLimit = 500 * 10 ** 8; // 500 wBTC

        totalMinted = 0;
        totalBurned = 0;
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
    function convertWBTCToStrBTC(uint256 wbtcAmount) external returns (uint256) {
        if (wbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 strbtcAmount = (wbtcAmount * rateNumerator) / rateDenominator;
        if (strbtcAmount == 0) revert ConversionResultedInZeroTokens();
        if (currentlyMinted() + strbtcAmount > mintingLimit) revert MintingLimitExceeded();

        totalMinted += strbtcAmount;

        wbtc.transferFrom(msg.sender, address(this), wbtcAmount);
        strbtc.converterMint(msg.sender, strbtcAmount);

        emit WBTCConverted(msg.sender, wbtcAmount, strbtcAmount);

        return strbtcAmount;
    }

    /**
     * @notice Converts strBTC back to wBTC at the outgoing rate
     * @param strbtcAmount The amount of strBTC to convert
     * @return The amount of wBTC received
     */
    function convertStrBTCToWBTC(uint256 strbtcAmount) external returns (uint256) {
        if (strbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 wbtcAmount = (strbtcAmount * rateNumerator) / rateDenominator;
        if (wbtcAmount == 0) revert ConversionResultedInZeroTokens();
        if (wbtc.balanceOf(address(this)) < wbtcAmount) revert InsufficientWBTCBalance();

        totalBurned += strbtcAmount;

        strbtc.transferFrom(msg.sender, address(this), strbtcAmount);
        strbtc.converterBurn(address(this), strbtcAmount);
        wbtc.transfer(msg.sender, wbtcAmount);

        emit StrBTCConverted(msg.sender, strbtcAmount, wbtcAmount);

        return wbtcAmount;
    }

    /**
     * @notice Withdrawal collected fees from the contract
     * @param token The address of the token to withdraw
     * @param recipient The address of the recipient
     * @param amount The amount of tokens
     */
    function withdraw(address token, address recipient, uint256 amount) external onlyWithdrawer {
        IERC20(token).transfer(recipient, amount);
    }
}

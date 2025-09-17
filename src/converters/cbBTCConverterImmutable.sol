// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../IStrBTC.sol";

/**
 * @title CBBTCConverterImmutable
 * @notice Contract for conversion between cbBTC and strBTC at immutable rates
 * @dev Allows users to exchange cbBTC for strBTC and vice versa with separate rates
 */
contract CBBTCConverterImmutable {
    error AmountMustBeGreaterThanZero();
    error ConversionResultedInZeroTokens();
    error InsufficientCBBTCBalance();
    error MintingLimitExceeded();
    error NotWithdrawer();

    IStrBTC public immutable strbtc;
    IERC20 public immutable cbbtc;
    address public immutable withdrawer;
    uint256 public immutable mintingLimit;

    uint256 public immutable rateNumerator;
    uint256 public immutable rateDenominator;

    uint256 public totalMinted;
    uint256 public totalBurned;

    event CBBTCConverted(address indexed user, uint256 cbbtcAmount, uint256 strbtcAmount);
    event StrBTCConverted(address indexed user, uint256 strbtcAmount, uint256 cbbtcAmount);

    modifier onlyWithdrawer() {
        if (msg.sender != withdrawer) revert NotWithdrawer();
        _;
    }

    /**
     * @notice constructor of the CBBTCConverterImmutable contract
     * @param _cbbtc cbBTC Token Address
     * @param _strbtc strBTC Token Address
     * @param _withdrawer Address authorized to withdraw fees
     */
    constructor(address _cbbtc, address _strbtc, address _withdrawer) {
        cbbtc = IERC20(_cbbtc);
        strbtc = IStrBTC(_strbtc);

        withdrawer = _withdrawer;

        rateNumerator = 999;
        rateDenominator = 1000;

        mintingLimit = 500 * 10 ** 8; // 500 cbBTC

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
     * @notice Converts cbBTC to strBTC at the incoming rate
     * @param cbbtcAmount The amount of cbBTC to convert
     * @return The amount of strBTC received
     */
    function convertCBBTCToStrBTC(uint256 cbbtcAmount) external returns (uint256) {
        if (cbbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 strbtcAmount = (cbbtcAmount * rateNumerator) / rateDenominator;
        if (strbtcAmount == 0) revert ConversionResultedInZeroTokens();
        if (currentlyMinted() + strbtcAmount > mintingLimit) revert MintingLimitExceeded();

        totalMinted += strbtcAmount;

        cbbtc.transferFrom(msg.sender, address(this), cbbtcAmount);
        strbtc.converterMint(msg.sender, strbtcAmount);

        emit CBBTCConverted(msg.sender, cbbtcAmount, strbtcAmount);

        return strbtcAmount;
    }

    /**
     * @notice Converts strBTC back to cbBTC at the outgoing rate
     * @param strbtcAmount The amount of strBTC to convert
     * @return The amount of cbBTC received
     */
    function convertStrBTCToCBBTC(uint256 strbtcAmount) external returns (uint256) {
        if (strbtcAmount == 0) revert AmountMustBeGreaterThanZero();

        uint256 cbbtcAmount = (strbtcAmount * rateNumerator) / rateDenominator;
        if (cbbtcAmount == 0) revert ConversionResultedInZeroTokens();
        if (cbbtc.balanceOf(address(this)) < cbbtcAmount) revert InsufficientCBBTCBalance();

        totalBurned += strbtcAmount;

        strbtc.transferFrom(msg.sender, address(this), strbtcAmount);
        strbtc.converterBurn(address(this), strbtcAmount);
        cbbtc.transfer(msg.sender, cbbtcAmount);

        emit StrBTCConverted(msg.sender, strbtcAmount, cbbtcAmount);

        return cbbtcAmount;
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

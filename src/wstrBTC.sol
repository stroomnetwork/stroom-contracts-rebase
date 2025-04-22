// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "../src/IStrBTC.sol";

contract wstrBTC is ERC20Permit {
    IStrBTC public strBTC;

    error CannotWrapZero();
    error CannotUnwrapZero();
    error NoWstrBTCSupply();
    error NoStrBTCBalance();

    /**
     * @param _strBTC address of the strBTC token to wrap
     */
    constructor(address _strBTC) ERC20("Wrapped Stroom Bitcoin", "wstrBTC") ERC20Permit("Wrapped Stroom Bitcoin") {
        strBTC = IStrBTC(_strBTC);
    }

    /**
     * @notice Wraps strBTC into wstrBTC.
     * @param strBTCAmount Amount of strBTC to wrap.
     * @return Amount of wstrBTC minted.
     */
    function wrap(uint256 strBTCAmount) external returns (uint256) {
        if (strBTCAmount == 0) revert CannotWrapZero();

        uint256 strBTCBalance = strBTC.balanceOf(address(this));
        uint256 wstrBTCSupply = totalSupply();

        uint256 wstrBTCAmount;
        if (wstrBTCSupply == 0 || strBTCBalance == 0) {
            // Initial wrap: 1:1 ratio
            wstrBTCAmount = strBTCAmount;
        } else {
            wstrBTCAmount = (strBTCAmount * wstrBTCSupply) / strBTCBalance;
        }

        strBTC.transferFrom(msg.sender, address(this), strBTCAmount);
        _mint(msg.sender, wstrBTCAmount);

        return wstrBTCAmount;
    }

    /**
     * @notice Unwraps wstrBTC back into strBTC.
     * @param wstrBTCAmount Amount of wstrBTC to unwrap.
     * @return Amount of strBTC returned.
     */
    function unwrap(uint256 wstrBTCAmount) external returns (uint256) {
        if (wstrBTCAmount == 0) revert CannotUnwrapZero();

        uint256 strBTCBalance = strBTC.balanceOf(address(this));
        uint256 wstrBTCSupply = totalSupply();

        if (wstrBTCSupply == 0) revert NoWstrBTCSupply();

        uint256 strBTCAmount = (wstrBTCAmount * strBTCBalance) / wstrBTCSupply;

        _burn(msg.sender, wstrBTCAmount);
        strBTC.transfer(msg.sender, strBTCAmount);

        return strBTCAmount;
    }

    /**
     * @notice Get the current amount of strBTC for amount of wstrBTC.
     * @return Amount of strBTC per amount of wstrBTC.
     */
    function strBTCPerToken(uint256 amount) external view returns (uint256) {
        uint256 strBTCBalance = strBTC.balanceOf(address(this));
        uint256 wstrBTCSupply = totalSupply();
        if (wstrBTCSupply == 0) revert NoWstrBTCSupply();

        return (strBTCBalance * amount) / wstrBTCSupply;
    }

    /**
     * @notice Get the current amount of wstrBTC for amount of strBTC.
     * @return Amount of wstrBTC per amount of strBTC.
     */
    function tokensPerStrBTC(uint256 amount) external view returns (uint256) {
        uint256 strBTCBalance = strBTC.balanceOf(address(this));
        uint256 wstrBTCSupply = totalSupply();
        if (strBTCBalance == 0) revert NoStrBTCBalance();

        return (wstrBTCSupply * amount) / strBTCBalance;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "../src/IStBTC.sol";

contract wstBTC is ERC20Permit {
    IStBTC public stBTC;

    error CannotWrapZero();
    error CannotUnwrapZero();
    error NoWstBTCSupply();
    error NoStBTCBalance();

    /**
     * @param _stBTC address of the stBTC token to wrap
     */
    constructor(address _stBTC) ERC20("Wrapped Stroom Bitcoin", "wstBTC") ERC20Permit("Wrapped Stroom Bitcoin") {
        stBTC = IStBTC(_stBTC);
    }

    /**
     * @notice Wraps stBTC into wstBTC.
     * @param stBTCAmount Amount of stBTC to wrap.
     * @return Amount of wstBTC minted.
     */
    function wrap(uint256 stBTCAmount) external returns (uint256) {
        if (stBTCAmount == 0) revert CannotWrapZero();

        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();

        uint256 wstBTCAmount;
        if (wstBTCSupply == 0 || stBTCBalance == 0) {
            // Initial wrap: 1:1 ratio
            wstBTCAmount = stBTCAmount;
        } else {
            wstBTCAmount = (stBTCAmount * wstBTCSupply) / stBTCBalance;
        }

        stBTC.transferFrom(msg.sender, address(this), stBTCAmount);
        _mint(msg.sender, wstBTCAmount);

        return wstBTCAmount;
    }

    /**
     * @notice Unwraps wstBTC back into stBTC.
     * @param wstBTCAmount Amount of wstBTC to unwrap.
     * @return Amount of stBTC returned.
     */
    function unwrap(uint256 wstBTCAmount) external returns (uint256) {
        if (wstBTCAmount == 0) revert CannotUnwrapZero();

        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();

        if (wstBTCSupply == 0) revert NoWstBTCSupply();

        uint256 stBTCAmount = (wstBTCAmount * stBTCBalance) / wstBTCSupply;

        _burn(msg.sender, wstBTCAmount);
        stBTC.transfer(msg.sender, stBTCAmount);

        return stBTCAmount;
    }

    /**
     * @notice Get the current amount of stBTC for amount of wstBTC.
     * @return Amount of stBTC per amount of wstBTC.
     */
    function stBTCPerToken(uint256 amount) external view returns (uint256) {
        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();
        if (wstBTCSupply == 0) revert NoWstBTCSupply();

        return (stBTCBalance * amount) / wstBTCSupply;
    }

    /**
     * @notice Get the current amount of wstBTC for amount of stBTC.
     * @return Amount of wstBTC per amount of stBTC.
     */
    function tokensPerStBTC(uint256 amount) external view returns (uint256) {
        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();
        if (stBTCBalance == 0) revert NoStBTCBalance();

        return (wstBTCSupply * amount) / stBTCBalance;
    }
}

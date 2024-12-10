// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "../src/IStBTC.sol";

contract wstBTC is ERC20Permit {
    IStBTC public stBTC;

    uint256 public constant BTC = 1e8; // sat

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
        require(stBTCAmount > 0, "wstBTC: Cannot wrap zero stBTC");

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
        require(wstBTCAmount > 0, "wstBTC: Cannot unwrap zero wstBTC");

        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();

        require(wstBTCSupply > 0, "wstBTC: No wstBTC supply");

        uint256 stBTCAmount = (wstBTCAmount * stBTCBalance) / wstBTCSupply;

        _burn(msg.sender, wstBTCAmount);
        stBTC.transfer(msg.sender, stBTCAmount);

        return stBTCAmount;
    }

    /**
     * @notice Get the current amount of stBTC for 1 wstBTC.
     * @return Amount of stBTC per 1 wstBTC.
     */
    function stBTCPerToken(uint256 amount) external view returns (uint256) {
        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();
        require(wstBTCSupply > 0, "wstBTC: No wstBTC supply");

        return (stBTCBalance * amount) / wstBTCSupply;
    }

    /**
     * @notice Get the current amount of wstBTC for 1 stBTC.
     * @return Amount of wstBTC per 1 stBTC.
     */
    function tokensPerStBTC(uint256 amount) external view returns (uint256) {
        uint256 stBTCBalance = stBTC.balanceOf(address(this));
        uint256 wstBTCSupply = totalSupply();
        require(stBTCBalance > 0, "wstBTC: No stBTC balance");

        return (wstBTCSupply * amount) / stBTCBalance;
    }
}

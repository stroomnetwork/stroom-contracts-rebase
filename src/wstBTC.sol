// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "../src/IStBTC.sol";

contract wstBTC is ERC20Permit {
    IStBTC public stBTC;

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

        uint256 wstBTCAmount = getSharesByPooledBTC(stBTCAmount);
        _mint(msg.sender, wstBTCAmount);

        stBTC.transferFrom(msg.sender, address(this), stBTCAmount);
        return wstBTCAmount;
    }

    /**
     * @notice Unwraps wstBTC back into stBTC.
     * @param wstBTCAmount Amount of wstBTC to unwrap.
     * @return Amount of stBTC returned.
     */
    function unwrap(uint256 wstBTCAmount) external returns (uint256) {
        require(wstBTCAmount > 0, "wstBTC: Cannot unwrap zero wstBTC");

        uint256 stBTCAmount = getPooledBTCByShares(wstBTCAmount);
        _burn(msg.sender, wstBTCAmount);

        stBTC.transfer(msg.sender, stBTCAmount);
        return stBTCAmount;
    }

    /**
     * @notice Calculates the number of shares corresponding to a given amount of staked BTC (stBTC).
     * @param btcAmount The amount of stBTC to convert to shares.
     * @return The number of shares equivalent to the given stBTC amount.
     */
    function getSharesByPooledBTC(uint256 btcAmount) public view returns (uint256) {
        uint256 totalShares = stBTC.totalShares();
        uint256 totalPooledBTC = stBTC.totalSupply();
        require(totalShares > 0 && totalPooledBTC > 0, "wstBTC: Invalid totalShares or totalPooledBTC");

        return (btcAmount * totalShares) / totalPooledBTC;
    }

    /**
     * @notice Calculates the amount of stBTC corresponding to a given number of shares.
     * @param sharesAmount The number of shares to convert to stBTC.
     * @return The amount of stBTC equivalent to the given shares.
     */
    function getPooledBTCByShares(uint256 sharesAmount) public view returns (uint256) {
        uint256 totalShares = stBTC.totalShares();
        uint256 totalPooledBTC = stBTC.totalSupply();
        require(totalShares > 0 && totalPooledBTC > 0, "wstBTC: Invalid totalShares or totalPooledBTC");

        return (sharesAmount * totalPooledBTC) / totalShares;
    }

    /**
     * @notice Get the current amount of stBTC for 1 wstBTC.
     * @return Amount of stBTC per 1 wstBTC.
     */
    function stBTCPerToken() external view returns (uint256) {
        return getPooledBTCByShares(1 ether);
    }

    /**
     * @notice Get the current amount of wstBTC for 1 stBTC.
     * @return Amount of wstBTC per 1 stBTC.
     */
    function tokensPerStBTC() external view returns (uint256) {
        return getSharesByPooledBTC(1 ether);
    }
}

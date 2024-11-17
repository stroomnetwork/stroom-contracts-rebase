// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "../src/IStBTC.sol";

contract wstBTC is ERC20Permit {
    IStBTC public stBTC;

    /**
     * @param _stBTC address of the stBTC token to wrap
     */
    constructor(IStBTC _stBTC)
        ERC20("Wrapped Stroom Bitcoin", "wstBTC")
        ERC20Permit("Wrapped Stroom Bitcoin")
    {
        stBTC = _stBTC;
    }

    /**
     * @notice Wraps stBTC into wstBTC.
     * @param _stBTCAmount Amount of stBTC to wrap.
     * @return Amount of wstBTC minted.
     */
    function wrap(uint256 _stBTCAmount) external returns (uint256) {
        require(_stBTCAmount > 0, "wstBTC: Cannot wrap zero stBTC");

        uint256 wstBTCAmount = getSharesByPooledBTC(_stBTCAmount);
        _mint(msg.sender, wstBTCAmount);

        stBTC.transferFrom(msg.sender, address(this), _stBTCAmount);
        return wstBTCAmount;
    }

    /**
     * @notice Unwraps wstBTC back into stBTC.
     * @param _wstBTCAmount Amount of wstBTC to unwrap.
     * @return Amount of stBTC returned.
     */
    function unwrap(uint256 _wstBTCAmount) external returns (uint256) {
        require(_wstBTCAmount > 0, "wstBTC: Cannot unwrap zero wstBTC");

        uint256 stBTCAmount = getPooledBTCByShares(_wstBTCAmount);
        _burn(msg.sender, _wstBTCAmount);

        stBTC.transfer(msg.sender, stBTCAmount);
        return stBTCAmount;
    }

    function getSharesByPooledBTC(uint256 _btcAmount) public view returns (uint256) {
        uint256 totalShares = stBTC.totalShares();
        uint256 totalPooledBTC = stBTC.totalSupply(); 
        require(totalShares > 0 && totalPooledBTC > 0, "wstBTC: Invalid totalShares or totalPooledBTC");

        return (_btcAmount * totalShares) / totalPooledBTC;
    }

    function getPooledBTCByShares(uint256 _sharesAmount) public view returns (uint256) {
        uint256 totalShares = stBTC.totalShares();
        uint256 totalPooledBTC = stBTC.totalSupply();
        require(totalShares > 0 && totalPooledBTC > 0, "wstBTC: Invalid totalShares or totalPooledBTC");

        return (_sharesAmount * totalPooledBTC) / totalShares;
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

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC4626.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract wstrBTC is ERC4626, ERC20Permit {
    /**
     * @param _strBTC address of the strBTC token (underlying asset)
     */
    constructor(address _strBTC)
        ERC4626(IERC20(_strBTC))
        ERC20("Wrapped Stroom Bitcoin", "wstrBTC")
        ERC20Permit("Wrapped Stroom Bitcoin")
    {}

    /**
     * @notice Override decimals to match underlying asset (strBTC has 8 decimals)
     */
    function decimals() public view override(ERC4626, ERC20) returns (uint8) {
        return ERC4626.decimals();
    }

    /**
     * @notice Returns the total amount of underlying strBTC held by the vault
     */
    function totalAssets() public view override returns (uint256) {
        return IERC20(asset()).balanceOf(address(this));
    }
}

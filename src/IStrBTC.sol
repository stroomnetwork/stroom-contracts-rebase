// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

interface IStrBTC {
    function totalSupply() external view returns (uint256);
    function totalShares() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function transfer(address recipient, uint256 amount) external returns (bool);
    function getSharesByPooledBTC(uint256 btcAmount) external view returns (uint256);
    function getPooledBTCByShares(uint256 sharesAmount) external view returns (uint256);
    function converterMint(address to, uint256 amount) external;
    function converterBurn(address from, uint256 amount) external;
}

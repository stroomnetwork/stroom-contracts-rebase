// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

interface IStBTC {
    function totalSupply() external view returns (uint256);
    function totalShares() external view returns (uint256);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function transfer(address recipient, uint256 amount) external returns (bool);
}
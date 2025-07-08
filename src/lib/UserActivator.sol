// SPDX-License-Identifier: MIT

pragma solidity 0.8.27;

import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract UserActivator is BTCDepositAddressDeriver, Ownable {
    event UserAddressActivated(address userETHAddress);

    // Mapping of activated addresses.
    mapping(address => bool) public activatedAddresses;

    constructor() BTCDepositAddressDeriver() Ownable(msg.sender) {}

    /**
     * @dev Activates a user by setting their address as activated.
     * @param _userAddress The address of the user to be activated.
     * Emits a `UserAddressActivated` event with the user address and their BTC deposit address.
     * Reverts if the user address is already activated.
     */
    function activateUser(address _userAddress) public {
        require(activatedAddresses[_userAddress] == false, "User is already activated");

        activatedAddresses[_userAddress] = true;

        emit UserAddressActivated(_userAddress);
    }

    /**
     * @dev Sets the seed for the BTC deposit address deriver.
     * @param _btcAddr1 The first BTC taproot address.
     * @param _btcAddr2 The second BTC taproot address.
     * @param _network The network of the BTC addresses.
     */
    function setSeed(string calldata _btcAddr1, string calldata _btcAddr2, BitcoinNetworkEncoder.Network _network)
        public
        override
        onlyOwner
    {
        BTCDepositAddressDeriver.setSeed(_btcAddr1, _btcAddr2, _network);
    }
}

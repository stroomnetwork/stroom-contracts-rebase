// SPDX-License-Identifier: MIT

pragma solidity ^0.8.27;

import {BitcoinUtils} from "blockchain-tools/src/BitcoinUtils.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

/// @title BitcoinUtilsWrapper
/// @notice Wrapper contract for the BitcoinUtils library
/// @dev This contract is required for the correct generation of Go bindings through abigen
contract BitcoinUtilsWrapper {
    using BitcoinUtils for BitcoinNetworkEncoder.Network;

    function validateBitcoinAddress(
        BitcoinNetworkEncoder.Network network,
        string calldata BTCAddress
    ) external view returns (bool) {
        return network.validateBitcoinAddress(BTCAddress);
    }

    function equalBytes(bytes memory one, bytes memory two) external pure returns (bool) {
        return BitcoinUtils.equalBytes(one, two);
    }

    function alphabetCheck(bytes memory BTCAddress) external pure returns (bool) {
        return BitcoinUtils.alphabetCheck(BTCAddress);
    }

    function validateBech32Checksum(string memory btcAddress) external view returns (bool) {
        return BitcoinUtils.validateBech32Checksum(btcAddress);
    }

    function validateBase58Checksum(string calldata btcAddress) external view returns (bool) {
        return BitcoinUtils.validateBase58Checksum(btcAddress);
    }

    function BECH32_ALPHABET_MAP(bytes1 char) external view returns (uint8) {
        return BitcoinUtils.BECH32_ALPHABET_MAP(char);
    }

    function polymodStep(uint256 pre) external pure returns (uint256) {
        return BitcoinUtils.polymodStep(pre);
    }

    function prefixChk(bytes memory prefix) external pure returns (uint256) {
        return BitcoinUtils.prefixChk(prefix);
    }
} 
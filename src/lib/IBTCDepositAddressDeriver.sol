// SPDX-License-Identifier: MIT

pragma solidity 0.8.27;

import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

/**
 * @title IBTCDepositAddressDeriver
 * @notice Interface for BTC deposit address derivation functionality
 */
interface IBTCDepositAddressDeriver {
    /// @notice Emitted when the seed is changed
    event SeedChanged(string btcAddr, string hrp);

    /// @notice Returns whether the seed has been set
    function wasSeedSet() external view returns (bool);

    /// @notice Returns the BTC address
    function btcAddr() external view returns (string memory);

    /// @notice Returns the network HRP (Human Readable Part)
    function networkHrp() external view returns (string memory);

    /// @notice Returns the x coordinate of the public key
    function p1x() external view returns (uint256);

    /// @notice Returns the y coordinate of the public key
    function p1y() external view returns (uint256);

    /**
     * @notice Sets the seed for BTC address derivation
     * @param _btcAddr The BTC taproot address
     * @param _network The network of the BTC address
     */
    function setSeed(string calldata _btcAddr, BitcoinNetworkEncoder.Network _network) external;

    /**
     * @notice Derives a BTC deposit address from an index
     * @param index The index to derive the address from
     * @return The derived BTC deposit address
     */
    function getBTCDepositAddress(uint256 index) external view returns (string memory);

    /**
     * @notice Parses a BTC taproot address to extract public key coordinates
     * @param _hrp The Human Readable Part (network prefix)
     * @param _bitcoinAddress The BTC taproot address
     * @return x The x coordinate of the public key
     * @return y The y coordinate of the public key
     */
    function parseBTCTaprootAddress(string memory _hrp, string calldata _bitcoinAddress)
        external
        pure
        returns (uint256 x, uint256 y);
}


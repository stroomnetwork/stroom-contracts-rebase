// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "@safe-global/safe-contracts/contracts/Safe.sol";
import "@safe-global/safe-contracts/contracts/proxies/SafeProxyFactory.sol";
import "@safe-global/safe-contracts/contracts/handler/CompatibilityFallbackHandler.sol";

/// @title SafeTestUtils
/// @notice Utility functions for testing Safe transactions
contract SafeTestUtils is Test {
    /// @notice Structure for Safe transactions
    struct SafeTx {
        address to;
        uint256 value;
        bytes data;
        Enum.Operation operation;
        uint256 safeTxGas;
        uint256 baseGas;
        uint256 gasPrice;
        address gasToken;
        address refundReceiver;
        uint256 nonce;
    }

    /// @notice Structure for storing signer data for sorting
    struct SignerData {
        address addr;
        uint256 privateKey;
    }

    /// @notice Create a SafeTx structure with basic parameters
    /// @dev Most parameters are set to 0 by default
    function createSafeTx(address to, uint256 value, bytes memory data, Enum.Operation operation, uint256 nonce)
        public
        pure
        returns (SafeTx memory)
    {
        return SafeTx({
            to: to,
            value: value,
            data: data,
            operation: operation,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: address(0),
            nonce: nonce
        });
    }

    /// @notice Calculate the hash of the Safe transaction
    function calculateSafeTxHash(Safe safe, SafeTx memory safeTx) public view returns (bytes32) {
        return safe.getTransactionHash(
            safeTx.to,
            safeTx.value,
            safeTx.data,
            safeTx.operation,
            safeTx.safeTxGas,
            safeTx.baseGas,
            safeTx.gasPrice,
            safeTx.gasToken,
            payable(safeTx.refundReceiver),
            safeTx.nonce
        );
    }

    /// @notice Sort private keys and generate corresponding addresses in ascending order
    /// @param privateKeys Array of unsorted private keys
    /// @return sortedKeys Sorted array of private keys
    /// @return sortedAddresses Corresponding sorted array of addresses
    function sortSigners(uint256[] memory privateKeys)
        public
        pure
        returns (uint256[] memory sortedKeys, address[] memory sortedAddresses)
    {
        if (privateKeys.length == 0) revert("No private keys provided");

        // Create array of signer data
        SignerData[] memory signers = new SignerData[](privateKeys.length);
        for (uint256 i = 0; i < privateKeys.length; i++) {
            signers[i] = SignerData({addr: vm.addr(privateKeys[i]), privateKey: privateKeys[i]});
        }

        // Sort by address (simple bubble sort)
        for (uint256 i = 0; i < signers.length - 1; i++) {
            for (uint256 j = 0; j < signers.length - i - 1; j++) {
                if (signers[j].addr > signers[j + 1].addr) {
                    SignerData memory temp = signers[j];
                    signers[j] = signers[j + 1];
                    signers[j + 1] = temp;
                }
            }
        }

        // Extract sorted keys and addresses
        sortedKeys = new uint256[](privateKeys.length);
        sortedAddresses = new address[](privateKeys.length);

        for (uint256 i = 0; i < signers.length; i++) {
            sortedKeys[i] = signers[i].privateKey;
            sortedAddresses[i] = signers[i].addr;
        }

        return (sortedKeys, sortedAddresses);
    }

    /// @notice Generate signatures for the transaction
    function signTransaction(bytes32 txHash, uint256[] memory privateKeys) public pure returns (bytes memory) {
        if (privateKeys.length == 0) revert("No private keys provided");

        // Sort private keys by address
        (uint256[] memory sortedKeys,) = sortSigners(privateKeys);

        bytes memory signatures;

        for (uint256 i = 0; i < sortedKeys.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(sortedKeys[i], txHash);
            signatures = abi.encodePacked(signatures, r, s, v);
        }

        return signatures;
    }

    /// @notice Execute a Safe transaction
    function execTransaction(Safe safe, SafeTx memory safeTx, bytes memory signatures) public returns (bool) {
        return safe.execTransaction(
            safeTx.to,
            safeTx.value,
            safeTx.data,
            safeTx.operation,
            safeTx.safeTxGas,
            safeTx.baseGas,
            safeTx.gasPrice,
            safeTx.gasToken,
            payable(safeTx.refundReceiver),
            signatures
        );
    }

    /// @notice Execute the full transaction process: create, sign and execute
    function execTransactionWithSigners(
        Safe safe,
        address to,
        uint256 value,
        bytes memory data,
        Enum.Operation operation,
        uint256[] memory privateKeys
    ) public returns (bool) {
        uint256 nonce = safe.nonce();
        SafeTx memory safeTx = createSafeTx(to, value, data, operation, nonce);
        bytes32 txHash = calculateSafeTxHash(safe, safeTx);
        bytes memory signatures = signTransaction(txHash, privateKeys);

        return execTransaction(safe, safeTx, signatures);
    }

    /// @notice Create and sign a contract call
    function createContractCallTx(Safe safe, address contractAddress, bytes memory data)
        public
        view
        returns (SafeTx memory)
    {
        return createSafeTx(contractAddress, 0, data, Enum.Operation.Call, safe.nonce());
    }

    /// @notice Execute a contract call through Safe
    function execContractCallWithSigners(
        Safe safe,
        address contractAddress,
        bytes memory data,
        uint256[] memory privateKeys
    ) public returns (bool) {
        return execTransactionWithSigners(safe, contractAddress, 0, data, Enum.Operation.Call, privateKeys);
    }

    /// @notice Deploy Safe infrastructure and create Safe with specified owners
    /// @param ownerPrivateKeys Array of private keys for Safe owners
    /// @param threshold Required number of signatures
    /// @return safe Deployed Safe contract
    /// @return sortedKeys Sorted private keys corresponding to Safe owners
    /// @return sortedOwners Sorted addresses of Safe owners
    function deploySafe(uint256[] memory ownerPrivateKeys, uint256 threshold)
        public
        returns (Safe safe, uint256[] memory sortedKeys, address[] memory sortedOwners)
    {
        // Sort signers
        (sortedKeys, sortedOwners) = sortSigners(ownerPrivateKeys);

        // Deploy Safe infrastructure
        CompatibilityFallbackHandler handler = new CompatibilityFallbackHandler();
        SafeProxyFactory safeFactory = new SafeProxyFactory();
        Safe safeSingleton = new Safe();

        // Setup Safe initialization data
        bytes memory initializer = abi.encodeWithSelector(
            Safe.setup.selector,
            sortedOwners, // _owners
            threshold, // _threshold
            address(0), // to
            bytes(""), // data
            address(handler), // fallbackHandler
            address(0), // paymentToken
            0, // payment
            payable(address(0)) // paymentReceiver
        );

        // Deploy Safe proxy
        safe = Safe(payable(safeFactory.createProxyWithNonce(address(safeSingleton), initializer, 0)));

        return (safe, sortedKeys, sortedOwners);
    }
}

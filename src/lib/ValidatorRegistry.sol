// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 STROOM <info@stroom.fi>

pragma solidity ^0.8.27;

import "openzeppelin-contracts/contracts/access/Ownable.sol";
import "bip340-solidity/src/Bip340Ecrec.sol";

contract ValidatorRegistry is Ownable, Bip340Ecrec {
    // Combined multisig key of all validators
    bytes32 public jointPublicKey;

    // Constant for the update validator public key message
    bytes public constant MESSAGE_UPDATE_JOINT_PUBLIC_KEY = "STROOM_UPDATE_JOINT_PUBLIC_KEY";

    // Mapping of used nonces
    mapping(bytes32 => bool) public usedNonces;

    error NonceAlreadyUsed(bytes32 nonce);
    error InvalidSignature();
    error NoJointPublicKey();

    // Event to log the update of the validator public key
    event JointPublicKeyUpdated(bytes32 newjointPublicKey);

    constructor() Ownable(msg.sender) {}

    // Function to update the validator public key
    function setJointPublicKey(bytes32 _jointPublicKey) public onlyOwner {
        jointPublicKey = _jointPublicKey;
        emit JointPublicKeyUpdated(_jointPublicKey);
    }

    // Function to update the validator public key with a signature
    function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes32 _nonce, bytes calldata signature) public {
        if (usedNonces[_nonce]) {
            revert NonceAlreadyUsed(_nonce);
        }
        if (!validateMessage(MESSAGE_UPDATE_JOINT_PUBLIC_KEY, abi.encodePacked(_jointPublicKey, _nonce), signature)) {
            revert InvalidSignature();
        }
        jointPublicKey = _jointPublicKey;
        usedNonces[_nonce] = true;
        emit JointPublicKeyUpdated(_jointPublicKey);
    }

    // Function to get the hash of a message based on its type
    function getMessageHash(bytes memory prefix, bytes memory data) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(keccak256(prefix), data));
    }

    // Function to validate a message
    function validateMessageHash(bytes32 hash, bytes calldata signature) public view returns (bool) {
        if (jointPublicKey == 0) {
            revert NoJointPublicKey();
        }

        // slice signature in two bytes32 blocks, rx and s
        uint256 rx = uint256(bytes32(signature[0:32]));
        uint256 s = uint256(bytes32(signature[32:64]));

        return verify(uint256(jointPublicKey), rx, s, hash);
    }

    // Function to validate a message
    function validateMessage(bytes memory prefix, bytes memory data, bytes calldata signature)
        public
        view
        returns (bool)
    {
        bytes32 hash = getMessageHash(prefix, data);
        return validateMessageHash(hash, signature);
    }
}

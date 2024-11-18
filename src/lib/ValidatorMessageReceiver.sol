// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 STROOM <info@stroom.fi>

pragma solidity ^0.8.27;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "bip340-solidity/src/Bip340Ecrec.sol";

import "./ValidatorRegistry.sol";

contract ValidatorMessageReceiver is OwnableUpgradeable, Bip340Ecrec {
    // Add a state variable for the ValidatorRegistry contract
    ValidatorRegistry public validatorRegistry;

    // Update the constructor to accept the ValidatorRegistry contract
    function initialize(ValidatorRegistry _validatorRegistry) public onlyInitializing {
        OwnableUpgradeable.__Ownable_init(msg.sender);
        validatorRegistry = _validatorRegistry;
    }

    modifier onlyValidator(bytes memory prefix, bytes memory data, bytes calldata signature) {
        require(
            validatorRegistry.validateMessage(prefix, data, signature), "ValidatorMessageReceiver: INVALID_SIGNATURE"
        );
        _;
    }
}

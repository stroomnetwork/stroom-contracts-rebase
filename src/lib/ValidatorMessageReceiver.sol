// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 STROOM <info@stroom.fi>

pragma solidity ^0.8.27;

import "./ValidatorRegistry.sol";

contract ValidatorMessageReceiver {
    // Add a state variable for the ValidatorRegistry contract
    ValidatorRegistry public validatorRegistry;

    error InvalidValidatorSignature();

    // Update the constructor to accept the ValidatorRegistry contract
    function initialize(ValidatorRegistry _validatorRegistry) public {
        validatorRegistry = _validatorRegistry;
    }

    modifier onlyValidator(bytes memory prefix, bytes memory data, bytes calldata signature) {
        if (!validatorRegistry.validateMessage(prefix, data, signature)) {
            revert InvalidValidatorSignature();
        }
        _;
    }
}

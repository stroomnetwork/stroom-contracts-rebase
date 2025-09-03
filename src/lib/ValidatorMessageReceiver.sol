// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 STROOM <info@stroom.fi>

pragma solidity 0.8.27;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorRegistry.sol";

contract ValidatorMessageReceiver is Initializable {
    // Add a state variable for the ValidatorRegistry contract
    ValidatorRegistry public validatorRegistry;

    error InvalidValidatorSignature();

    // Internal initialization function - can only be called once during contract initialization
    function __ValidatorMessageReceiver_init(ValidatorRegistry _validatorRegistry) internal onlyInitializing {
        validatorRegistry = _validatorRegistry;
    }

    modifier onlyValidator(bytes memory prefix, bytes memory data, bytes calldata signature) {
        if (!validatorRegistry.validateMessage(prefix, data, signature)) {
            revert InvalidValidatorSignature();
        }
        _;
    }
}

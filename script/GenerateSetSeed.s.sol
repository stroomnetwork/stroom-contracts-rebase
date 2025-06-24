// script/GenerateSetSeedCalldata.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "../src/lib/TimelockController.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract GenerateSetSeedCalldataScript is Script {
    address public timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");
    uint256 public delay = vm.envUint("TIMELOCK_DELAY");

    address public userActivator = vm.envAddress("APP_ETH_USER_ACTIVATOR_ADDRESS");
    address public validatorRegistry = vm.envAddress("APP_ETH_VALIDATOR_REGISTRY_ADDRESS");

    string public taprootAddress1 = vm.envString("TAPROOT_ADDRESS1");
    string public taprootAddress2 = vm.envString("TAPROOT_ADDRESS2");
    bytes32 public jointPublicKey = vm.envBytes32("JOINT_PUBLIC_KEY");
    uint8 public network = uint8(vm.envUint("BITCOIN_NETWORK"));


    function run() public view {
        console.log("=== GENERATION OF CALLDATA FOR SETSEED OPERATIONS ===\n");

        bytes memory setSeedCalldata = abi.encodeWithSelector(
            BTCDepositAddressDeriver.setSeed.selector, 
            taprootAddress1, 
            taprootAddress2, 
            BitcoinNetworkEncoder.Network(network)
        );

        _printCalldataInfo(
            "setSeed(string,string,uint8)",
            userActivator,
            setSeedCalldata,
            keccak256("SetSeed2025"),
            string.concat(
                "\n  taprootAddress1: ", taprootAddress1, "\n",
                "  taprootAddress2: ", taprootAddress2, "\n",
                "  network: ", vm.toString(network), "\n"
            )
        );

        console.log("\n=== GENERATION OF CALLDATA FOR SETJOINTKEY OPERATIONS ===\n");

        bytes memory setJointKeyCalldata = abi.encodeWithSelector(
            ValidatorRegistry.setJointPublicKey.selector, 
            jointPublicKey
        );

        _printCalldataInfo(
            "setJointPublicKey(bytes32)",
            validatorRegistry,
            setJointKeyCalldata,
            keccak256("SetJointKey2025"),
            string.concat("  jointPublicKey: ", vm.toString(jointPublicKey))
        );
    }

    function _printCalldataInfo(
        string memory functionSignature,
        address target,
        bytes memory functionCalldata,
        bytes32 salt,
        string memory params
    ) internal view {
        bytes32 operationId = keccak256(abi.encode(target, 0, functionCalldata, bytes32(0), salt));

        console.log("Function:", functionSignature);
        console.log("Params:");
        console.log(params);
        console.log("");
        console.log("Target:", target);
        console.log("Value: 0");
        console.log("Data:", vm.toString(functionCalldata));
        console.log("Predecessor: 0x0000000000000000000000000000000000000000000000000000000000000000");
        console.log("Salt:", vm.toString(salt));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID:", vm.toString(operationId));
    }
}

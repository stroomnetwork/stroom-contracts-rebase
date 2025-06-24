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
        
        _generateSetSeedCalldata();
        
        console.log("\n=== GENERATION OF CALLDATA FOR SETJOINTKEY OPERATIONS ===\n");
        
        _generateSetJointKeyCalldata();
    }

    function _generateSetSeedCalldata() internal view {
        BitcoinNetworkEncoder.Network _network = BitcoinNetworkEncoder.Network(network);
        
        bytes memory functionCalldata = abi.encodeWithSelector(
            BTCDepositAddressDeriver.setSeed.selector,
            taprootAddress1,
            taprootAddress2,
            _network
        );

        bytes32 salt = keccak256("SetSeed2025");

        bytes32 operationId = keccak256(abi.encode(
            userActivator, 0, functionCalldata, bytes32(0), salt
        ));

        console.log("Function: setSeed(string,string,uint8)");
        console.log("Params:");
        console.log("  taprootAddress1:", taprootAddress1);
        console.log("  taprootAddress2:", taprootAddress2);
        console.log("  network:", network);
        console.log("");
        console.log("Target:", userActivator);
        console.log("Value: 0");
        console.log("Data:", vm.toString(functionCalldata));
        console.log("Predecessor: 0x0000000000000000000000000000000000000000000000000000000000000000");
        console.log("Salt:", vm.toString(salt));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID:", vm.toString(operationId));
    }

    function _generateSetJointKeyCalldata() internal view {
        bytes memory functionCalldata = abi.encodeWithSelector(
            ValidatorRegistry.setJointPublicKey.selector,
            jointPublicKey
        );

        bytes32 salt = keccak256("SetJointKey2025");
        
        bytes32 operationId = keccak256(abi.encode(
            validatorRegistry, 0, functionCalldata, bytes32(0), salt
        ));

        console.log("Function: setJointPublicKey(bytes32)");
        console.log("Params:");
        console.log("  jointPublicKey:", vm.toString(jointPublicKey));
        console.log("");
        console.log("Target:", validatorRegistry);
        console.log("Value: 0");
        console.log("Data:", vm.toString(functionCalldata));
        console.log("Predecessor: 0x0000000000000000000000000000000000000000000000000000000000000000");
        console.log("Salt:", vm.toString(salt));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID:", vm.toString(operationId));
    }
}
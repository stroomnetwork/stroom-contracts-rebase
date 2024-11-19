// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract SetSeedScript is Script {
    uint8 private network;
    address private userActivatorContract;
    address private validatorRegistryContract;
    string private taprootAddress1;
    string private taprootAddress2;
    bytes32 private jointNetworkKey;
    BitcoinNetworkEncoder.Network _network;

    function setUp() public {
        network = uint8(vm.envUint("BITCOIN_NETWORK"));
        _network = BitcoinNetworkEncoder.Network(network);
        userActivatorContract = vm.envAddress("APP_ETH_USER_ACTIVATOR_ADDRESS");
        validatorRegistryContract = vm.envAddress("APP_ETH_VALIDATOR_REGISTRY_ADDRESS");
        taprootAddress1 = vm.envString("TAPROOT_ADDRESS1");
        taprootAddress2 = vm.envString("TAPROOT_ADDRESS2");
        jointNetworkKey = vm.envBytes32("JOINT_PUBLIC_KEY");
    }

    function run() public {
        vm.startBroadcast();

        BTCDepositAddressDeriver activator = BTCDepositAddressDeriver(userActivatorContract);
        ValidatorRegistry vr = ValidatorRegistry(validatorRegistryContract);

        activator.setSeed(taprootAddress1, taprootAddress2, _network);
        vr.setJointPublicKey(jointNetworkKey);

        vm.stopBroadcast();
    }
}

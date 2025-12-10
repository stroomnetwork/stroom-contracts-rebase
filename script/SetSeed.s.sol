// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import {BTCDepositAddressDeriver} from "blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract SetSeedScript is Script {
    uint8 private network;
    address private btcDeriverContract;
    address private validatorRegistryContract;
    string private taprootAddress;
    bytes32 private jointNetworkKey;
    BitcoinNetworkEncoder.Network _network;

    function setUp() public {
        network = uint8(vm.envUint("BITCOIN_NETWORK"));
        _network = BitcoinNetworkEncoder.Network(network);
        btcDeriverContract = vm.envAddress("APP_ETH_BTC_DERIVER_ADDRESS");
        validatorRegistryContract = vm.envAddress("APP_ETH_VALIDATOR_REGISTRY_ADDRESS");
        taprootAddress = vm.envString("TAPROOT_ADDRESS");
        jointNetworkKey = vm.envBytes32("JOINT_PUBLIC_KEY");
    }

    function run() public {
        vm.startBroadcast();

        BTCDepositAddressDeriver activator = BTCDepositAddressDeriver(btcDeriverContract);
        ValidatorRegistry vr = ValidatorRegistry(validatorRegistryContract);

        activator.setSeed(taprootAddress, _network);
        vr.setJointPublicKey(jointNetworkKey);

        vm.stopBroadcast();
    }
}

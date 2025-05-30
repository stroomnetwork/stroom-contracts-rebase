// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test} from "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../src/strBTC.sol";
import "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract ValidatorRegistryTest is Test {
    ValidatorRegistry public vr;
    strBTC public token;

    address public admin;
    address public pauser;
    address public alice;

    function setUp() public {
        vr = new ValidatorRegistry();
        admin = msg.sender;
        pauser = makeAddr("Pauser");
        // Deploy strBTC
        strBTC strBtcImplementation = new strBTC();
        bytes memory strBtcData =
            abi.encodeWithSelector(strBTC.initialize.selector, BitcoinNetworkEncoder.Network.Testnet, vr, admin, pauser);
        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), admin, strBtcData);

        token = strBTC(address(strBtcProxy));

        alice = makeAddr("Alice");
    }

    function testCannotReuseNonceForJointPublicKeyUpdate() public {
        bytes32 initialKey = hex"b0bd465d2251011baa0672dabd8243b7d4761116c590ec006565c8e65d8410ff";
        bytes32 newKey = hex"1c2bf8395c38340251a46d1c59e584b1350397d6404b38b745469631e9ada49a";
        bytes32 nonce = hex"0000000000000000000000000000000000000000000000000000000000000001";

        bytes memory signature =
            hex"c7b7ccf883dd9130308713a6400cc05687b8a074ff5763974b68f20d6e628247d44fcf15328c7d931aa045e9998c04772ff5b6ca544ae781893e68e64b918a8e";

        // Set initial key
        vr.setJointPublicKey(initialKey);

        // First key update - should succeed
        vr.setJointPublicKeySigned(newKey, nonce, signature);
        assertEq(vr.jointPublicKey(), newKey, "Key should be updated");
        assertTrue(vr.usedNonces(nonce), "Nonce should be marked as used");

        // Try to reuse the same nonce - should be rejected
        vm.expectRevert(abi.encodeWithSelector(ValidatorRegistry.NonceAlreadyUsed.selector, nonce));
        vr.setJointPublicKeySigned(initialKey, nonce, signature);
    }

    function testPreventReplayAttack() public {
        bytes32 key1 = hex"b0bd465d2251011baa0672dabd8243b7d4761116c590ec006565c8e65d8410ff";
        bytes32 key2 = hex"1c2bf8395c38340251a46d1c59e584b1350397d6404b38b745469631e9ada49a";

        bytes32 nonce1 = hex"0000000000000000000000000000000000000000000000000000000000000001";
        bytes32 nonce2 = hex"0000000000000000000000000000000000000000000000000000000000000002";

        bytes memory signature1 =
            hex"c7b7ccf883dd9130308713a6400cc05687b8a074ff5763974b68f20d6e628247d44fcf15328c7d931aa045e9998c04772ff5b6ca544ae781893e68e64b918a8e";
        bytes memory signature2 =
            hex"c7bc9d74ada826f2826f2964aadc9eebc5e404ae704042c133d2c14442c2fed854ff0b9e565c53c0543e2d4eed3d4e0d07943bd38aaca360cd543286f6fbb509";

        // Setting K1
        vr.setJointPublicKey(key1);
        assertEq(vr.jointPublicKey(), key1, "Initial key should be set");

        // Try to update to K2 by nonAuthorized address
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, alice));
        vr.setJointPublicKey(key2);

        // Update to K2
        vr.setJointPublicKeySigned(key2, nonce1, signature1);
        assertEq(vr.jointPublicKey(), key2, "Key should be updated to K2");

        // Update back to K1
        vr.setJointPublicKeySigned(key1, nonce2, signature2);
        assertEq(vr.jointPublicKey(), key1, "Key should be updated back to K1");

        // Try to reuse signature1 - should be rejected
        vm.expectRevert(abi.encodeWithSelector(ValidatorRegistry.NonceAlreadyUsed.selector, nonce1));
        vr.setJointPublicKeySigned(key2, nonce1, signature1);
    }

    function testInvalidSignature() public {
        bytes32 key1 = hex"b0bd465d2251011baa0672dabd8243b7d4761116c590ec006565c8e65d8410ff";
        bytes32 key2 = hex"1c2bf8395c38340251a46d1c59e584b1350397d6404b38b745469631e9ada49a";
        bytes32 nonce = hex"0000000000000000000000000000000000000000000000000000000000000001";

        // Invalid signature
        bytes memory invalidSignature =
            hex"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

        vr.setJointPublicKey(key1);

        // Try to update with invalid signature
        vm.expectRevert(ValidatorRegistry.InvalidSignature.selector);
        vr.setJointPublicKeySigned(key2, nonce, invalidSignature);
    }

    function testNoJointPublicKey() public {
        bytes32 key = hex"1c2bf8395c38340251a46d1c59e584b1350397d6404b38b745469631e9ada49a";
        bytes32 nonce = hex"0000000000000000000000000000000000000000000000000000000000000001";
        bytes memory signature =
            hex"c7b7ccf883dd9130308713a6400cc05687b8a074ff5763974b68f20d6e628247d44fcf15328c7d931aa045e9998c04772ff5b6ca544ae781893e68e64b918a8e";

        // Try to validate message without setting jointPublicKey
        vm.expectRevert(ValidatorRegistry.NoJointPublicKey.selector);
        vr.setJointPublicKeySigned(key, nonce, signature);
    }
}

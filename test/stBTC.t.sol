// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "../src/stBTC.sol";
import "../src/wstBTC.sol";
import "../src/IStBTC.sol";

contract STBTCTest is Test {
    stBTC public token;
    wstBTC public wtoken;

    ValidatorRegistry public vr;

    address public admin;

    address public alice;
    address public bob;

    uint256 public constant BTC = 1e8; // sat

    function getAdminAddress(address proxy) internal view returns (address) {
        bytes32 adminSlot = vm.load(proxy, ERC1967Utils.ADMIN_SLOT);
        return address(uint160(uint256(adminSlot)));
    }

    function setUp() public {
        console.log("setUp");

        vr = new ValidatorRegistry();
        vr.setJointPublicKey(hex"7febbae0513352eb5e95991108c70bd5b3a773f4f4d278b7766d537a8b79030c");

        admin = msg.sender;

        // Deploy stBTC implementation
        stBTC stBtcImplementation = new stBTC();
        bytes memory stBtcData = abi.encodeWithSelector(
            stBTC.initialize.selector,
            BitcoinNetworkEncoder.Network.Mainnet,
            vr
        );
        TransparentUpgradeableProxy stBtcProxy = new TransparentUpgradeableProxy(
                address(stBtcImplementation),
                admin,
                stBtcData
            );
        token = stBTC(address(stBtcProxy));

        // Deploy wstBTC
        wtoken = new wstBTC(IStBTC(address(token)));

        alice = makeAddr("alice");
        bob = makeAddr("bob");
    }

    function testTokenInfo() public view {
        console.log("testTokenInfo");

        assertEq(token.name(), "Stroom Bitcoin");
        assertEq(token.symbol(), "stBTC");
        assertEq(token.decimals(), 8);
    }

    function testMint() public {
        console.log("testMint");
        assertEq(token.balanceOf(admin), 0);

        token.mint(1 * BTC, admin, keccak256("deadbeef"));
        assertEq(token.balanceOf(admin), 1 * BTC);
    }

    function testPause() public {
        console.log("testPause");

        token.mint(1 * BTC, admin, keccak256("deadbeef1"));

        token.pause();

        assertTrue(token.paused(), "token should be paused");

        vm.expectRevert(bytes4(keccak256("EnforcedPause()")));
        token.mint(1 * BTC, admin, keccak256("deadbeef2"));

        token.unpause();

        token.mint(1 * BTC, admin, keccak256("deadbeef2"));
    }

}
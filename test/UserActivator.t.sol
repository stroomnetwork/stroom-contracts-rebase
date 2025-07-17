// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import {Test, console} from "forge-std/Test.sol";
import {UserActivator} from "../src/lib/UserActivator.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";
import {BTCDepositAddressDeriver} from "../lib/blockchain-tools/src/BTCDepositAddressDeriver.sol";
import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";

contract UserActivatorTest is Test {
    UserActivator public deriver;

    function setUp() public {
        deriver = new UserActivator();
    }

    function testUserActivation() public {
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));

        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        assertEq(deriver.getBTCDepositAddress(user), "tb1p8pjjwryjq9d7tke50ndcd97kqxkeztk4k85lzg7l2nynektg9zdsq836sr");
    }

    function testCannotGetBTCDepositAddressIfNotActivated() public {
        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectRevert("User is not activated");
        deriver.getBTCDepositAddress(user);
    }

    function testUserActivationSimnet() public {
        deriver.setSeed(
            "sb1p5z8wl5tu7m0d79vzqqsl9gu0x4fkjug857fusx4fl4kfgwh5j25sxv5dv3",
            "sb1pfusykjdt46ktwq03d20uqqf94uh9487344wr3q5v9szzsxnjdfkszvtlt8",
            BitcoinNetworkEncoder.Network.Simnet
        );

        address user = 0x1EaCa1277BcDFa83E60658D8938B3D63cD3E63C1;

        vm.expectEmit(address(deriver));

        emit UserActivator.UserAddressActivated(user);

        deriver.activateUser(user);

        //console.log("user deposit address", deriver.getBTCDepositAddress(user));
        assertEq(deriver.getBTCDepositAddress(user), "sb1pupljglwqunp2q22sahwjvmxwmxxyj3aatnhemxlneae3l03w2k5sf8r34a");
    }

    function testCannotSetSeedIfNotOwner() public {
        address notOwner = address(0x123);

        vm.startPrank(notOwner);

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        vm.stopPrank();
    }

    function testCannotBypassOwnerCheckWithBaseContract() public {
        address notOwner = address(0x123);

        vm.startPrank(notOwner);

        // An attempt to bypass the check by casting to a base type
        BTCDepositAddressDeriver baseContract = BTCDepositAddressDeriver(address(deriver));

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        baseContract.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        vm.stopPrank();
    }

    function testOwnerCanSetSeed() public {
        // By default, in tests, msg.sender is the owner
        deriver.setSeed(
            "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3",
            "tb1psfpmk6v8cvd8kr4rdda0l8gwyn42v5yfjlqkhnureprgs5tuumkqvdkewz",
            BitcoinNetworkEncoder.Network.Testnet
        );

        // Check that the seed is set
        assertTrue(deriver.wasSeedSet());
        assertEq(deriver.btcAddr1(), "tb1p7g532zgvuzv8fz3hs02wvn2almqh8qyvz4xdr564nannkxh28kdq62ewy3");
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test, console2} from "forge-std/Test.sol";
import {StroomTimelockController} from "../src/lib/TimelockController.sol";
import {strBTC} from "../src/strBTC.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title TimelockUpgradeTest
 * @notice Tests to verify the ability of StroomTimelockController to manage contract updates
 */
contract TimelockUpgradeTest is Test {
    StroomTimelockController public timelock;
    strBTC public tokenImplementation;
    strBTC public token;
    ValidatorRegistry public validatorRegistry;
    TransparentUpgradeableProxy public proxy;
    ProxyAdmin public proxyAdmin;

    address public admin;
    address public proposer;
    address public executor;
    address public user;

    uint256 public constant TIMELOCK_DELAY = 2 days;

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;
    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");

    function setUp() public {
        admin = address(0x1);
        proposer = address(0x2);
        executor = address(0x3);
        user = address(0x4);

        address[] memory proposers = new address[](1);
        proposers[0] = proposer;

        address[] memory executors = new address[](1);
        executors[0] = executor;

        timelock = new StroomTimelockController(TIMELOCK_DELAY, proposers, executors, admin);

        validatorRegistry = new ValidatorRegistry();
        tokenImplementation = new strBTC();

        bytes memory initData =
            abi.encodeWithSelector(strBTC.initialize.selector, BitcoinNetworkEncoder.Network.Mainnet, validatorRegistry);
        proxy = new TransparentUpgradeableProxy(address(tokenImplementation), address(timelock), initData); // set timelock as admin

        token = strBTC(address(proxy));

        address proxyAdminAddress = address(uint160(uint256(vm.load(address(proxy), ERC1967Utils.ADMIN_SLOT))));
        proxyAdmin = ProxyAdmin(proxyAdminAddress);
    }

    function testTimelockIsProxyAdmin() public view {
        assertEq(proxyAdmin.owner(), address(timelock), "Timelock should be owner of ProxyAdmin");

        assertTrue(address(timelock) != address(0), "Timelock should exist");
        assertTrue(address(proxy) != address(0), "Proxy should exist");
        assertTrue(address(proxyAdmin) != address(0), "ProxyAdmin should exist");
    }

    function testTimelockCanScheduleUpgrade() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(proposer);

        bytes32 operationId =
            timelock.hashOperation(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));

        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);

        vm.stopPrank();

        assertTrue(timelock.isOperationPending(operationId));
        assertFalse(timelock.isOperationReady(operationId));
    }

    function testTimelockCanExecuteUpgradeAfterDelay() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(proposer);
        bytes32 operationId =
            timelock.hashOperation(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));

        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();

        assertFalse(timelock.isOperationReady(operationId));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        assertTrue(timelock.isOperationReady(operationId));

        vm.prank(executor);
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));

        assertTrue(timelock.isOperationDone(operationId));
    }

    function testDirectUpgradeFailsWithoutTimelock() public {
        strBTC newImplementation = new strBTC();

        vm.startPrank(user);
        vm.expectRevert();
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(proxy)), address(newImplementation), "");
        vm.stopPrank();

        vm.startPrank(admin);
        vm.expectRevert();
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(proxy)), address(newImplementation), "");
        vm.stopPrank();
    }

    function testCancelScheduledUpgrade() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(proposer);
        bytes32 operationId =
            timelock.hashOperation(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));

        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);

        timelock.cancel(operationId);
        vm.stopPrank();

        assertFalse(timelock.isOperationPending(operationId));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
        vm.stopPrank();
    }

    function testOnlyProposerCanScheduleUpgrade() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert();
        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.startPrank(proposer);
        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();
    }

    function testOnlyExecutorCanExecuteUpgrade() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(proposer);
        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(proposer);
        vm.expectRevert();
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert();
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
        vm.stopPrank();

        vm.prank(executor);
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
    }

    function testTimelockDelayProtectsAgainstRushUpgrades() public {
        strBTC newImplementation = new strBTC();

        bytes memory upgradeData = abi.encodeWithSelector(
            ProxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(address(proxy)),
            address(newImplementation),
            ""
        );

        vm.startPrank(proposer);
        timelock.schedule(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
        vm.stopPrank();

        vm.warp(block.timestamp + TIMELOCK_DELAY - 1);

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.execute(address(proxyAdmin), 0, upgradeData, bytes32(0), bytes32(uint256(1)));
        vm.stopPrank();
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test, console2} from "forge-std/Test.sol";
import {StroomTimelockController} from "../src/lib/TimelockController.sol";
import {strBTC} from "../src/strBTC.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";

contract TimeLockControllerTest is Test {
    StroomTimelockController public timelock;
    strBTC public token;
    ValidatorRegistry public validatorRegistry;
    address public admin;
    address public proposer;
    address public executor;
    address public multisig;
    address public converter;
    uint256 public constant TIMELOCK_DELAY = 2 days;

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;
    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");
    bytes32 public constant CANCELLER_ROLE = keccak256("CANCELLER_ROLE");
    bytes32 public constant CONVERTER_ROLE = keccak256("CONVERTER_ROLE");

    function setUp() public {
        admin = address(0x1);
        proposer = address(0x2);
        executor = address(0x3);
        multisig = address(0x4);
        converter = address(0x5);

        address[] memory proposers = new address[](2);
        proposers[0] = proposer;
        proposers[1] = multisig;

        address[] memory executors = new address[](2);
        executors[0] = executor;
        executors[1] = multisig;

        timelock = new StroomTimelockController(TIMELOCK_DELAY, proposers, executors, admin);

        validatorRegistry = new ValidatorRegistry();

        token = new strBTC();
        token.initialize(BitcoinNetworkEncoder.Network.Mainnet, validatorRegistry);

        token.grantRole(DEFAULT_ADMIN_ROLE, admin);

        token.revokeRole(DEFAULT_ADMIN_ROLE, address(this));

        // Transfer admin role to timelock
        vm.startPrank(admin);
        token.grantRole(DEFAULT_ADMIN_ROLE, address(timelock));
        vm.stopPrank();
    }

    function testTimelockControllerInitialization() public view {
        assertTrue(address(timelock) != address(0), "Timelock address should not be zero");
        assertEq(timelock.getMinDelay(), TIMELOCK_DELAY, "Incorrect timelock delay");

        bytes memory testData = abi.encodeWithSignature("test()");
        bytes32 operationId = timelock.hashOperation(address(token), 0, testData, bytes32(0), bytes32(0));

        assertFalse(timelock.isOperation(operationId), "No operations should exist initially");
        assertFalse(timelock.isOperationPending(operationId), "No pending operations should exist initially");
        assertFalse(timelock.isOperationReady(operationId), "No ready operations should exist initially");
        assertFalse(timelock.isOperationDone(operationId), "No done operations should exist initially");
    }

    function testTimelockRolesCorrectlyAssigned() public view {
        assertTrue(timelock.hasRole(PROPOSER_ROLE, proposer));
        assertTrue(timelock.hasRole(PROPOSER_ROLE, multisig));
        assertTrue(timelock.hasRole(EXECUTOR_ROLE, executor));
        assertTrue(timelock.hasRole(EXECUTOR_ROLE, multisig));
        assertTrue(timelock.hasRole(CANCELLER_ROLE, proposer));
        assertTrue(timelock.hasRole(CANCELLER_ROLE, multisig));
        assertTrue(timelock.hasRole(DEFAULT_ADMIN_ROLE, admin));
        assertTrue(timelock.hasRole(DEFAULT_ADMIN_ROLE, address(timelock)));
    }

    function testTimelockDelayCorrectlySet() public view {
        assertEq(timelock.getMinDelay(), TIMELOCK_DELAY);
    }

    function testTimelockIsOnlyDefaultAdmin() public {
        vm.prank(admin);
        token.revokeRole(DEFAULT_ADMIN_ROLE, admin);

        assertTrue(token.hasRole(DEFAULT_ADMIN_ROLE, address(timelock)));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, admin));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, proposer));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, executor));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, address(this)));
    }

    function testCriticalFunctionsRequireTimelock() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testCriticalFunctionsExecutableAfterDelay() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(proposer);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        assertTrue(timelock.isOperationPending(operationId));
        assertFalse(timelock.isOperationReady(operationId));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        assertTrue(timelock.isOperationReady(operationId));

        vm.prank(executor);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
        assertTrue(timelock.isOperationDone(operationId));
    }

    function testScheduledOperationCanBeCancelled() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(proposer);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);

        timelock.cancel(operationId);
        vm.stopPrank();

        assertFalse(timelock.isOperationPending(operationId));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testOnlyProposerCanSchedule() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(executor);
        vm.expectRevert();
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.startPrank(proposer);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        assertTrue(timelock.isOperationPending(operationId));
    }

    function testOnlyExecutorCanExecute() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(proposer);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(proposer);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        vm.prank(executor);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));

        assertTrue(timelock.isOperationDone(operationId));
        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testMultisigAsProposer() public view {
        assertTrue(timelock.hasRole(PROPOSER_ROLE, multisig));
    }

    function testMultisigAsExecutor() public view {
        assertTrue(timelock.hasRole(EXECUTOR_ROLE, multisig));
    }

    function testMultisigCanProposeOperation() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(multisig);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        assertTrue(timelock.isOperationPending(operationId));
    }

    function testMultisigCanExecuteOperation() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(multisig);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.prank(multisig);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
        assertTrue(timelock.isOperationDone(operationId));
    }

    function testUpdatingFeesThroughTimelock() public {
        uint256 newMaxRewardPercent = 50; // 0.5%
        bytes memory data = abi.encodeWithSelector(token.setMaxRewardPercent.selector, newMaxRewardPercent);

        _scheduleAndExecuteOperation(data);

        assertEq(token.maxRewardPercent(), newMaxRewardPercent);
    }

    function testUpdatingMinimumStakeThroughTimelock() public {
        uint256 newMinWithdraw = 10000; // 0.0001 BTC
        bytes memory data = abi.encodeWithSelector(token.setMinWithdrawAmount.selector, newMinWithdraw);

        _scheduleAndExecuteOperation(data);

        assertEq(token.minWithdrawAmount(), newMinWithdraw);
    }

    function testEmergencyPauseThroughTimelock() public {
        bytes memory data = abi.encodeWithSelector(token.pause.selector);

        _scheduleAndExecuteOperation(data);

        assertTrue(token.paused());
    }

    function testDeployNewTimelockController() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;

        address[] memory executors = new address[](1);
        executors[0] = executor;

        StroomTimelockController newTimelock = new StroomTimelockController(TIMELOCK_DELAY, proposers, executors, admin);

        assertTrue(address(newTimelock) != address(0));
        assertTrue(newTimelock.hasRole(PROPOSER_ROLE, proposer));
        assertTrue(newTimelock.hasRole(EXECUTOR_ROLE, executor));
    }

    function testTransferDefaultAdminRoleToNewTimelock() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;

        address[] memory executors = new address[](1);
        executors[0] = executor;

        StroomTimelockController newTimelock = new StroomTimelockController(TIMELOCK_DELAY, proposers, executors, admin);

        bytes memory data = abi.encodeWithSelector(token.grantRole.selector, DEFAULT_ADMIN_ROLE, address(newTimelock));

        _scheduleAndExecuteOperation(data);

        assertTrue(token.hasRole(DEFAULT_ADMIN_ROLE, address(newTimelock)));
    }

    // ======= Helper function =======

    function _scheduleAndExecuteOperation(bytes memory data) internal {
        vm.startPrank(proposer);
        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(executor);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(timelock.isOperationDone(operationId));
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test, console2} from "forge-std/Test.sol";
import {strBTC} from "../src/strBTC.sol";
import {StroomTimelockController} from "../src/lib/TimelockController.sol";
import {ValidatorRegistry} from "../src/lib/ValidatorRegistry.sol";
import {BitcoinNetworkEncoder} from "blockchain-tools/src/BitcoinNetworkEncoder.sol";
import {SafeTestUtils} from "./utils/SafeTestUtils.sol";
import {Safe} from "@safe-global/safe-contracts/contracts/Safe.sol";

contract MultisigTest is Test {
    strBTC public token;
    StroomTimelockController public timelock;
    ValidatorRegistry public validatorRegistry;
    Safe public safe;
    SafeTestUtils public safeUtils;

    address public admin;
    address public converter;
    address public pauser;
    uint256 public constant TIMELOCK_DELAY = 2 days;
    uint256 public constant SAFE_THRESHOLD = 2;

    uint256[] public ownerPrivateKeys;
    address[] public owners;

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;
    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");
    bytes32 public constant CANCELLER_ROLE = keccak256("CANCELLER_ROLE");
    bytes32 public constant CONVERTER_ROLE = keccak256("CONVERTER_ROLE");

    function setUp() public {
        admin = address(0x9999);
        converter = address(0x8888);
        pauser = address(0x7777);

        uint256[] memory keys = new uint256[](3);
        keys[0] = 0x1;
        keys[1] = 0x2;
        keys[2] = 0x3;

        safeUtils = new SafeTestUtils();

        // Deploy Safe with utilities
        (safe, ownerPrivateKeys, owners) = safeUtils.deploySafe(keys, SAFE_THRESHOLD);

        // Deploy timelock
        address[] memory proposers = new address[](1);
        proposers[0] = address(safe);

        address[] memory executors = new address[](1);
        executors[0] = admin;

        timelock = new StroomTimelockController(TIMELOCK_DELAY, proposers, executors, admin);

        vm.prank(admin);
        timelock.grantRole(CANCELLER_ROLE, admin);

        // Deploy strBTC
        validatorRegistry = new ValidatorRegistry();
        token = new strBTC();
        token.initialize(BitcoinNetworkEncoder.Network.Mainnet, validatorRegistry, address(timelock), pauser);
    }

    function testProposers() public view {
        assertTrue(timelock.hasRole(PROPOSER_ROLE, address(safe)));
        assertFalse(timelock.hasRole(PROPOSER_ROLE, admin));
    }

    function testExecutors() public view {
        assertTrue(timelock.hasRole(EXECUTOR_ROLE, admin));
        assertFalse(timelock.hasRole(EXECUTOR_ROLE, address(safe)));
    }

    function testCancellers() public view {
        assertTrue(timelock.hasRole(CANCELLER_ROLE, admin));
        assertTrue(timelock.hasRole(CANCELLER_ROLE, address(safe)));
    }

    function testTimelockControlsStrBTCAdmin() public view {
        assertTrue(token.hasRole(DEFAULT_ADMIN_ROLE, address(timelock)));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, address(safe)));
        assertFalse(token.hasRole(DEFAULT_ADMIN_ROLE, admin));
    }

    function testSafeOwnersConfiguration() public view {
        address[] memory actualOwners = safe.getOwners();
        assertEq(actualOwners.length, owners.length);

        for (uint256 i = 0; i < owners.length; i++) {
            assertTrue(safe.isOwner(owners[i]));
        }
    }

    function testSafeThresholdConfiguration() public view {
        assertEq(safe.getThreshold(), SAFE_THRESHOLD);
    }

    function testSafeCanSetMaxRewardPercent() public {
        uint256 newMaxRewardPercent = 50; // 0.5%
        bytes memory data = abi.encodeWithSelector(token.setMaxRewardPercent.selector, newMaxRewardPercent);

        _scheduleAndExecuteOperationThroughSafe(address(token), data);

        assertEq(token.maxRewardPercent(), newMaxRewardPercent);
    }

    function testSafeCanSetMinWithdrawAmount() public {
        uint256 newMinWithdraw = 10000; // 0.0001 BTC
        bytes memory data = abi.encodeWithSelector(token.setMinWithdrawAmount.selector, newMinWithdraw);

        _scheduleAndExecuteOperationThroughSafe(address(token), data);

        assertEq(token.minWithdrawAmount(), newMinWithdraw);
    }

    function testSafeCanAddConverter() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        _scheduleAndExecuteOperationThroughSafe(address(token), data);

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testSafeCanRemoveConverter() public {
        bytes memory addData = abi.encodeWithSelector(token.addConverter.selector, converter);
        _scheduleAndExecuteOperationThroughSafe(address(token), addData);

        bytes memory removeData = abi.encodeWithSelector(token.removeConverter.selector, converter);
        _scheduleAndExecuteOperationThroughSafe(address(token), removeData);

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testSafeCanSetMinTimeBetweenRewards() public {
        uint256 newMinTime = 24 hours;
        bytes memory data = abi.encodeWithSelector(token.setMinTimeBetweenRewards.selector, newMinTime);

        _scheduleAndExecuteOperationThroughSafe(address(token), data);

        assertEq(token.minTimeBetweenRewards(), newMinTime);
    }

    function testSingleSignatureCannotSchedule() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory timelockData = abi.encodeWithSelector(
            timelock.schedule.selector, address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY
        );

        uint256[] memory singleKey = new uint256[](1);
        singleKey[0] = ownerPrivateKeys[0];

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, singleKey);
    }

    function testThresholdSignaturesCanSchedule() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory timelockData = abi.encodeWithSelector(
            timelock.schedule.selector, address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY
        );

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        bool success = safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, thresholdKeys);
        assertTrue(success);

        bytes32 operationId = timelock.hashOperation(address(token), 0, data, bytes32(0), bytes32(0));
        assertTrue(timelock.isOperationPending(operationId));
    }

    function testInvalidSignatureCannotSchedule() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory timelockData = abi.encodeWithSelector(
            timelock.schedule.selector, address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY
        );

        uint256[] memory invalidKeys = new uint256[](SAFE_THRESHOLD);
        invalidKeys[0] = 0x999; // Invalid key
        invalidKeys[1] = 0x888; // Invalid key

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, invalidKeys);
    }

    function testNonceProtectionAgainstReplay() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory timelockData = abi.encodeWithSelector(
            timelock.schedule.selector, address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY
        );

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        bool success = safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, thresholdKeys);
        assertTrue(success);

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, thresholdKeys);
    }

    function testOnlyAdminCanExecute() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        _scheduleOperationThroughSafe(address(token), data);

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(owners[0]);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        vm.startPrank(admin);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testAdminCannotExecuteBeforeDelay() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        _scheduleOperationThroughSafe(address(token), data);

        vm.startPrank(admin);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testAdminCanCancelScheduledOperation() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes32 operationId = _scheduleOperationThroughSafe(address(token), data);

        vm.startPrank(admin);
        timelock.cancel(operationId);
        vm.stopPrank();

        assertFalse(timelock.isOperationPending(operationId));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(admin);
        vm.expectRevert();
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testNonOwnerCannotCreateSafeTransaction() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory timelockData = abi.encodeWithSelector(
            timelock.schedule.selector, address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY
        );

        uint256[] memory nonOwnerKey = new uint256[](1);
        nonOwnerKey[0] = 0x999; // Not an owner

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, nonOwnerKey);
    }

    function testDirectCallsToStrBTCFail() public {
        vm.startPrank(address(safe));
        vm.expectRevert();
        token.addConverter(converter);
        vm.stopPrank();

        vm.startPrank(admin);
        vm.expectRevert();
        token.addConverter(converter);
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testBypassTimelockAttemptsFail() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(token), data, thresholdKeys);

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testSafeCannotExecuteDirectly() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        _scheduleOperationThroughSafe(address(token), data);

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        bytes memory executeData =
            abi.encodeWithSelector(timelock.execute.selector, address(token), 0, data, bytes32(0), bytes32(0));

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(timelock), executeData, thresholdKeys);

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testSeparationOfConcerns() public view {
        assertTrue(timelock.hasRole(PROPOSER_ROLE, address(safe)));
        assertFalse(timelock.hasRole(EXECUTOR_ROLE, address(safe)));

        assertTrue(timelock.hasRole(EXECUTOR_ROLE, admin));
        assertTrue(timelock.hasRole(CANCELLER_ROLE, admin));
        assertFalse(timelock.hasRole(PROPOSER_ROLE, admin));
    }

    function testAdminCannotBypassMultisig() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(admin);
        vm.expectRevert();
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testSafeCannotBypassTimelock() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        vm.expectRevert();
        safeUtils.execContractCallWithSigners(safe, address(token), data, thresholdKeys);

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testMaliciousAdminProtection() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        vm.startPrank(admin);
        vm.expectRevert();
        timelock.schedule(address(token), 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);
        vm.stopPrank();

        assertFalse(token.hasRole(CONVERTER_ROLE, converter));
    }

    function testCompleteWorkflow() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        // Step 1: Safe schedules operation
        bytes32 operationId = _scheduleOperationThroughSafe(address(token), data);
        assertTrue(timelock.isOperationPending(operationId));

        // Step 2: Wait for delay
        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);
        assertTrue(timelock.isOperationReady(operationId));

        // Step 3: Admin executes
        vm.startPrank(admin);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
        assertTrue(timelock.isOperationDone(operationId));
    }

    function testMultipleProposalsWorkflow() public {
        address converter2 = address(0x7777);

        bytes memory data1 = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory data2 = abi.encodeWithSelector(token.addConverter.selector, converter2);

        bytes32 operationId1 = _scheduleOperationThroughSafe(address(token), data1);
        bytes32 operationId2 = _scheduleOperationThroughSafe(address(token), data2);

        assertTrue(timelock.isOperationPending(operationId1));
        assertTrue(timelock.isOperationPending(operationId2));

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(admin);
        timelock.execute(address(token), 0, data1, bytes32(0), bytes32(0));
        timelock.execute(address(token), 0, data2, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
        assertTrue(token.hasRole(CONVERTER_ROLE, converter2));
    }

    function testAdminSelectiveExecution() public {
        address converter2 = address(0x7777);

        bytes memory data1 = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes memory data2 = abi.encodeWithSelector(token.addConverter.selector, converter2);

        bytes32 operationId1 = _scheduleOperationThroughSafe(address(token), data1);
        bytes32 operationId2 = _scheduleOperationThroughSafe(address(token), data2);

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(admin);
        timelock.execute(address(token), 0, data1, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
        assertFalse(token.hasRole(CONVERTER_ROLE, converter2));
        assertTrue(timelock.isOperationDone(operationId1));
        assertTrue(timelock.isOperationReady(operationId2)); // Still ready for execution
    }

    function testDelayedExecutionWorkflow() public {
        bytes memory data = abi.encodeWithSelector(token.addConverter.selector, converter);

        bytes32 operationId = _scheduleOperationThroughSafe(address(token), data);

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1 days);

        assertTrue(timelock.isOperationReady(operationId));

        vm.startPrank(admin);
        timelock.execute(address(token), 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(token.hasRole(CONVERTER_ROLE, converter));
    }

    // ======= Helper functions =======

    function _scheduleOperationThroughSafe(address target, bytes memory data) internal returns (bytes32) {
        bytes memory timelockData =
            abi.encodeWithSelector(timelock.schedule.selector, target, 0, data, bytes32(0), bytes32(0), TIMELOCK_DELAY);

        uint256[] memory thresholdKeys = new uint256[](SAFE_THRESHOLD);
        for (uint256 i = 0; i < SAFE_THRESHOLD; i++) {
            thresholdKeys[i] = ownerPrivateKeys[i];
        }

        bool success = safeUtils.execContractCallWithSigners(safe, address(timelock), timelockData, thresholdKeys);
        assertTrue(success);

        return timelock.hashOperation(target, 0, data, bytes32(0), bytes32(0));
    }

    function _scheduleAndExecuteOperationThroughSafe(address target, bytes memory data) internal {
        bytes32 operationId = _scheduleOperationThroughSafe(target, data);

        vm.warp(block.timestamp + TIMELOCK_DELAY + 1);

        vm.startPrank(admin);
        timelock.execute(target, 0, data, bytes32(0), bytes32(0));
        vm.stopPrank();

        assertTrue(timelock.isOperationDone(operationId));
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";
import "../src/strBTC.sol";
import "../src/wBTCConverter.sol";
import "../src/lib/ValidatorRegistry.sol";

contract MockWBTC is ERC20 {
    constructor() ERC20("Wrapped BTC", "WBTC") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract WBTCConverterTest is Test {
    WBTCConverter public wbtcConverter;
    strBTC public strbtc;
    IERC20 public wbtc;

    ValidatorRegistry public validatorRegistry;
    address public owner;
    address public user1;
    address public user2;
    address public rateSetter;

    uint256 public constant INITIAL_WBTC_SUPPLY = 1000 * 10 ** 8; // 1000 WBTC

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        rateSetter = makeAddr("rateSetter");

        MockWBTC mockWBTC = new MockWBTC();
        mockWBTC.mint(address(this), INITIAL_WBTC_SUPPLY);
        mockWBTC.mint(user1, INITIAL_WBTC_SUPPLY);
        mockWBTC.mint(user2, INITIAL_WBTC_SUPPLY);
        wbtc = IERC20(address(mockWBTC));

        BitcoinNetworkEncoder.Network network = BitcoinNetworkEncoder.Network.Testnet;

        validatorRegistry = new ValidatorRegistry();

        strBTC strBtcImplementation = new strBTC();

        bytes memory strBtcData = abi.encodeWithSelector(strBTC.initialize.selector, network, validatorRegistry);

        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), owner, strBtcData);
        strbtc = strBTC(address(strBtcProxy));

        WBTCConverter wbtcConverterImplementation = new WBTCConverter();

        bytes memory wbtcConverterData =
            abi.encodeWithSelector(WBTCConverter.initialize.selector, address(wbtc), address(strbtc));

        TransparentUpgradeableProxy wbtcConverterProxy =
            new TransparentUpgradeableProxy(address(wbtcConverterImplementation), owner, wbtcConverterData);
        wbtcConverter = WBTCConverter(address(wbtcConverterProxy));

        strbtc.grantRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));

        wbtcConverter.grantRole(wbtcConverter.RATE_SETTER_ROLE(), rateSetter);

        vm.prank(user1);
        wbtc.approve(address(wbtcConverter), type(uint256).max);
        vm.prank(user2);
        wbtc.approve(address(wbtcConverter), type(uint256).max);
    }

    function testInitialization() public view {
        // Check successful initialization with correct addresses
        assertEq(address(wbtcConverter.wbtc()), address(wbtc), "Incorrect WBTC address");
        assertEq(address(wbtcConverter.strbtc()), address(strbtc), "Incorrect strBTC address");

        // Check initial exchange rate (should be 1:1)
        assertEq(wbtcConverter.exchangeRateNumerator(), 1, "Incorrect exchange rate numerator");
        assertEq(wbtcConverter.exchangeRateDenominator(), 1, "Incorrect exchange rate denominator");

        // Check proper role configuration
        bytes32 adminRole = wbtcConverter.DEFAULT_ADMIN_ROLE();
        bytes32 rateSetterRole = wbtcConverter.RATE_SETTER_ROLE();
        bytes32 converterRole = strbtc.CONVERTER_ROLE();

        // Check that owner has admin role
        assertTrue(wbtcConverter.hasRole(adminRole, owner), "Owner does not have admin role");

        // Check that rateSetter has rate setter role
        assertTrue(wbtcConverter.hasRole(rateSetterRole, rateSetter), "RateSetter does not have correct role");

        // Check CONVERTER_ROLE assignment
        assertTrue(
            strbtc.hasRole(converterRole, address(wbtcConverter)), "Converter does not have CONVERTER_ROLE in strBTC"
        );

        // Check: regular users should not have special roles
        assertFalse(wbtcConverter.hasRole(adminRole, user1), "Regular user should not have admin role");
        assertFalse(wbtcConverter.hasRole(rateSetterRole, user1), "Regular user should not have rate setter role");

        // Check wbtc allowance for conversion
        assertEq(
            wbtc.allowance(user1, address(wbtcConverter)),
            type(uint256).max,
            "User does not have sufficient allowance for WBTC"
        );
    }

    function testFuzzConvertWBTCToStrBTC(uint256 amount) public {
        vm.assume(amount > 0 && amount <= 1000 * 10 ** 8);

        uint256 user1WbtcBalanceBefore = wbtc.balanceOf(user1);
        uint256 user1StrbtcBalanceBefore = strbtc.balanceOf(user1);
        uint256 converterWbtcBalanceBefore = wbtc.balanceOf(address(wbtcConverter));

        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 expectedStrbtc =
            (amount * wbtcConverter.exchangeRateNumerator()) / wbtcConverter.exchangeRateDenominator();

        assertEq(strbtcReceived, expectedStrbtc, "Received strBTC amount incorrect");

        assertEq(wbtc.balanceOf(user1), user1WbtcBalanceBefore - amount, "User WBTC balance not reduced correctly");
        assertEq(
            strbtc.balanceOf(user1),
            user1StrbtcBalanceBefore + expectedStrbtc,
            "User strBTC balance not increased correctly"
        );
        assertEq(
            wbtc.balanceOf(address(wbtcConverter)),
            converterWbtcBalanceBefore + amount,
            "Converter WBTC balance not increased correctly"
        );
    }

    function testFuzzConvertWBTCToStrBTCWithCustomRate(
        uint256 amountInput,
        uint256 numeratorInput,
        uint256 denominatorInput
    ) public {
        uint256 amount = bound(amountInput, 1, 1000 * 10 ** 8);
        uint256 numerator = bound(numeratorInput, 1, 200);
        uint256 denominator = bound(denominatorInput, 1, 200);

        // Ensure conversion doesn't result in zero tokens due to rounding
        vm.assume((amount * numerator) / denominator > 0);

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(numerator, denominator);

        uint256 user1WbtcBalanceBefore = wbtc.balanceOf(user1);
        uint256 user1StrbtcBalanceBefore = strbtc.balanceOf(user1);

        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 expectedStrbtc = (amount * numerator) / denominator;

        assertEq(strbtcReceived, expectedStrbtc, "Received strBTC amount incorrect with custom rate");

        assertEq(wbtc.balanceOf(user1), user1WbtcBalanceBefore - amount, "User WBTC balance not reduced correctly");
        assertEq(
            strbtc.balanceOf(user1),
            user1StrbtcBalanceBefore + expectedStrbtc,
            "User strBTC balance not increased correctly with custom rate"
        );
    }

    function testFuzzConvertStrBTCToWBTC(uint256 amount) public {
        vm.assume(amount > 0 && amount <= 1000 * 10 ** 8);

        // First convert WBTC to strBTC to have some strBTC balance
        uint256 initialConversion = 1000 * 10 ** 8;
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(initialConversion);

        uint256 user1WbtcBalanceBefore = wbtc.balanceOf(user1);
        uint256 user1StrbtcBalanceBefore = strbtc.balanceOf(user1);
        uint256 converterWbtcBalanceBefore = wbtc.balanceOf(address(wbtcConverter));

        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), amount);

        vm.prank(user1);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(amount);

        uint256 expectedWbtc =
            (amount * wbtcConverter.exchangeRateDenominator()) / wbtcConverter.exchangeRateNumerator();

        assertEq(wbtcReceived, expectedWbtc, "Received WBTC amount incorrect");

        assertEq(
            wbtc.balanceOf(user1), user1WbtcBalanceBefore + expectedWbtc, "User WBTC balance not increased correctly"
        );
        assertEq(
            strbtc.balanceOf(user1), user1StrbtcBalanceBefore - amount, "User strBTC balance not reduced correctly"
        );
        assertEq(
            wbtc.balanceOf(address(wbtcConverter)),
            converterWbtcBalanceBefore - expectedWbtc,
            "Converter WBTC balance not reduced correctly"
        );
    }

    function testFuzzConvertStrBTCToWBTCWithCustomRate(
        uint256 amountInput,
        uint256 numeratorInput,
        uint256 denominatorInput
    ) public {
        uint256 amount = bound(amountInput, 1, 100 * 10 ** 8);
        uint256 numerator = bound(numeratorInput, 1, 200);
        uint256 denominator = bound(denominatorInput, 1, 200);

        vm.assume((amount * denominator) / numerator > 0);
        vm.assume((amount * numerator) / denominator > 0);

        uint256 initialConversion = 1000 * 10 ** 8;
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(initialConversion);

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(numerator, denominator);

        uint256 expectedWbtc = (amount * denominator) / numerator;

        vm.assume(wbtc.balanceOf(address(wbtcConverter)) >= expectedWbtc);

        uint256 user1WbtcBalanceBefore = wbtc.balanceOf(user1);
        uint256 user1StrbtcBalanceBefore = strbtc.balanceOf(user1);

        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), amount);

        vm.prank(user1);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(amount);

        assertEq(wbtcReceived, expectedWbtc, "Received WBTC amount incorrect with custom rate");

        assertEq(
            wbtc.balanceOf(user1), user1WbtcBalanceBefore + expectedWbtc, "User WBTC balance not increased correctly"
        );
        assertEq(
            strbtc.balanceOf(user1), user1StrbtcBalanceBefore - amount, "User strBTC balance not reduced correctly"
        );
    }

    function testZeroAmountConversion() public {
        vm.prank(user1);
        vm.expectRevert(WBTCConverter.AmountMustBeGreaterThanZero.selector);
        wbtcConverter.convertWBTCToStrBTC(0);

        vm.prank(user1);
        vm.expectRevert(WBTCConverter.AmountMustBeGreaterThanZero.selector);
        wbtcConverter.convertStrBTCToWBTC(0);
    }

    function testConverterRoleRequired() public {
        uint256 amount = 1 * 10 ** 8; // 1 BTC

        strbtc.revokeRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));

        vm.prank(user1);

        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(amount);

        strbtc.grantRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));
    }

    function testInsufficientWBTCBalance() public {
        uint256 amount = 1 * 10 ** 8; // 1 BTC

        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), wbtc.balanceOf(address(wbtcConverter)));

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertStrBTCToWBTC(amount);
    }

    function testSetExchangeRate() public {
        assertEq(wbtcConverter.exchangeRateNumerator(), 1, "Initial numerator should be 1");
        assertEq(wbtcConverter.exchangeRateDenominator(), 1, "Initial denominator should be 1");

        uint256 newNumerator = 105;
        uint256 newDenominator = 100;

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.exchangeRateNumerator(), newNumerator, "Numerator not updated correctly");
        assertEq(wbtcConverter.exchangeRateDenominator(), newDenominator, "Denominator not updated correctly");

        newNumerator = 95;
        newDenominator = 100;

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.exchangeRateNumerator(), newNumerator, "Numerator not updated correctly");
        assertEq(wbtcConverter.exchangeRateDenominator(), newDenominator, "Denominator not updated correctly");
    }

    function testSetExchangeRateWithZeroNumerator() public {
        // Try to set rate with zero numerator (should revert)
        uint256 newNumerator = 0;
        uint256 newDenominator = 100;

        vm.prank(rateSetter);
        vm.expectRevert();
        wbtcConverter.setExchangeRate(newNumerator, newDenominator);

        // Try to set rate with zero denominator (should revert)
        newNumerator = 100;
        newDenominator = 0;

        vm.prank(rateSetter);
        vm.expectRevert();
        wbtcConverter.setExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.exchangeRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.exchangeRateDenominator(), 1, "Denominator should not change");
    }

    function testSetExchangeRateUnauthorized() public {
        // Try to set rate as unauthorized user (should revert)
        uint256 newNumerator = 105;
        uint256 newDenominator = 100;

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.setExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.exchangeRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.exchangeRateDenominator(), 1, "Denominator should not change");
    }

    function testRoleManagement() public {
        // Define roles
        bytes32 adminRole = wbtcConverter.DEFAULT_ADMIN_ROLE();
        bytes32 rateSetterRole = wbtcConverter.RATE_SETTER_ROLE();

        // Test initial role setup
        assertTrue(wbtcConverter.hasRole(adminRole, owner), "Owner should have admin role");
        assertTrue(wbtcConverter.hasRole(rateSetterRole, rateSetter), "RateSetter should have rate setter role");
        assertFalse(wbtcConverter.hasRole(adminRole, user1), "User1 should not have admin role");
        assertFalse(wbtcConverter.hasRole(rateSetterRole, user1), "User1 should not have rate setter role");

        // Test granting roles
        vm.prank(owner);
        wbtcConverter.grantRole(rateSetterRole, user1);
        assertTrue(wbtcConverter.hasRole(rateSetterRole, user1), "User1 should now have rate setter role");

        // Test revoking roles
        vm.prank(owner);
        wbtcConverter.revokeRole(rateSetterRole, user1);
        assertFalse(wbtcConverter.hasRole(rateSetterRole, user1), "User1 should no longer have rate setter role");

        // Test unauthorized role granting
        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.grantRole(rateSetterRole, user2);
    }

    function testRenounceRole() public {
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.RATE_SETTER_ROLE(), user1);
        assertTrue(
            wbtcConverter.hasRole(wbtcConverter.RATE_SETTER_ROLE(), user1), "TestUser should have rate setter role"
        );

        vm.startPrank(user1);
        wbtcConverter.renounceRole(wbtcConverter.RATE_SETTER_ROLE(), user1);
        vm.stopPrank();

        assertFalse(
            wbtcConverter.hasRole(wbtcConverter.RATE_SETTER_ROLE(), user1),
            "TestUser should no longer have rate setter role after renouncing"
        );
    }

    function testAdminRoleManagement() public {
        // Owner can grant admin role to another account
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2);
        assertTrue(wbtcConverter.hasRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2), "User2 should now have admin role");

        // New admin can grant rate setter role
        vm.prank(user2);
        wbtcConverter.grantRole(wbtcConverter.RATE_SETTER_ROLE(), user1);
        assertTrue(
            wbtcConverter.hasRole(wbtcConverter.RATE_SETTER_ROLE(), user1),
            "User1 should have rate setter role granted by new admin"
        );

        // New admin can revoke rate setter role
        vm.prank(user2);
        wbtcConverter.revokeRole(wbtcConverter.RATE_SETTER_ROLE(), user1);
        assertFalse(
            wbtcConverter.hasRole(wbtcConverter.RATE_SETTER_ROLE(), user1),
            "User1 should no longer have rate setter role after revocation by new admin"
        );
    }

    function testConverterRoleOnStrBTC() public {
        assertTrue(
            strbtc.hasRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter)),
            "Converter should have CONVERTER_ROLE on strBTC"
        );

        // Admin can revoke converter role
        vm.prank(owner);
        strbtc.revokeRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));
        assertFalse(
            strbtc.hasRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter)),
            "Converter should no longer have CONVERTER_ROLE after revocation"
        );

        // Test that conversion fails when role is revoked
        uint256 amount = 1 * 10 ** 8; // 1 BTC

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(amount);
    }

    function testPauseContract() public {
        assertFalse(wbtcConverter.paused(), "Contract should not be paused after initialization");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.pause();

        vm.prank(owner);
        wbtcConverter.pause();

        assertTrue(wbtcConverter.paused(), "Contract should be paused");

        uint256 amount = 1 * 10 ** 8; // 1 BTC

        vm.prank(user1);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        wbtcConverter.convertWBTCToStrBTC(amount);

        vm.prank(user1);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        wbtcConverter.convertStrBTCToWBTC(amount);
    }

    function testUnpauseContract() public {
        vm.prank(owner);
        wbtcConverter.pause();
        assertTrue(wbtcConverter.paused(), "Contract should be paused");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.unpause();

        vm.prank(owner);
        wbtcConverter.unpause();

        assertFalse(wbtcConverter.paused(), "Contract should not be paused after unpause");

        uint256 amount = 1 * 10 ** 8; // 1 BTC

        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        assertEq(strbtcReceived, amount, "Conversion should work after unpausing");
    }

    function testGrantPauserRole() public {
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2);

        assertTrue(
            wbtcConverter.hasRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2), "User should have DEFAULT_ADMIN_ROLE"
        );

        vm.prank(user2);
        wbtcConverter.pause();
        assertTrue(wbtcConverter.paused(), "User with PAUSER_ROLE should be able to pause");
    }

    function testEmergencyWithdrawalOnlyAdmin() public {
        uint256 amount = 1 * 10 ** 8; // 1 BTC
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 converterBalance = wbtc.balanceOf(address(wbtcConverter));
        assertEq(converterBalance, amount, "Converter should have some WBTC");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.emergencyWithdraw(address(wbtc), converterBalance);

        uint256 ownerBalanceBefore = wbtc.balanceOf(owner);

        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), converterBalance);

        assertEq(wbtc.balanceOf(address(wbtcConverter)), 0, "Converter should have zero WBTC balance after withdrawal");
        assertEq(wbtc.balanceOf(owner), ownerBalanceBefore + converterBalance, "Owner should receive withdrawn funds");
    }

    function testEmergencyWithdrawalFullWorkflow() public {
        uint256 amount = 5 * 10 ** 8; // 5 BTC
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        assertEq(strbtc.balanceOf(user1), amount, "User should have received strBTC");
        assertEq(wbtc.balanceOf(address(wbtcConverter)), amount, "Converter should have received WBTC");

        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), amount);

        assertEq(wbtc.balanceOf(address(wbtcConverter)), 0, "Converter should have no WBTC after emergency withdrawal");

        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), amount);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertStrBTCToWBTC(amount);

        vm.prank(owner);
        wbtc.transfer(address(wbtcConverter), amount);

        assertEq(wbtc.balanceOf(address(wbtcConverter)), amount, "Converter should have WBTC after replenishment");

        uint256 userWbtcBalanceBefore = wbtc.balanceOf(user1);

        vm.prank(user1);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(amount);

        assertEq(wbtcReceived, amount, "User should receive the correct amount of WBTC");
        assertEq(wbtc.balanceOf(user1), userWbtcBalanceBefore + amount, "User should have received WBTC");
        assertEq(strbtc.balanceOf(user1), 0, "User should have no strBTC left");
    }

    function testEmergencyWithdrawalPartial() public {
        uint256 amount = 10 * 10 ** 8; // 10 BTC
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 withdrawAmount = 5 * 10 ** 8; // 5 BTC
        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), withdrawAmount);

        assertEq(
            wbtc.balanceOf(address(wbtcConverter)), amount - withdrawAmount, "Converter should have remaining WBTC"
        );

        uint256 smallAmount = 1 * 10 ** 8; // 1 BTC
        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), smallAmount);

        vm.prank(user1);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(smallAmount);

        assertEq(wbtcReceived, smallAmount, "User should receive the correct amount of WBTC");

        uint256 largeAmount = 5 * 10 ** 8; // 5 BTC
        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), largeAmount);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertStrBTCToWBTC(largeAmount);
    }

    function testEmergencyWithdrawalDifferentToken() public {
        MockWBTC newMockWBTC = new MockWBTC();
        uint256 amount = 1000 * 10 ** 18;
        newMockWBTC.mint(address(wbtcConverter), amount);

        assertEq(newMockWBTC.balanceOf(address(wbtcConverter)), amount, "Converter should have mock tokens");

        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(newMockWBTC), amount);

        assertEq(
            newMockWBTC.balanceOf(address(wbtcConverter)), 0, "Converter should have no mock tokens after withdrawal"
        );
        assertEq(newMockWBTC.balanceOf(owner), amount, "Owner should have received mock tokens");
    }

    function testExtremeExchangeRates() public {
        uint256 highNumerator = 1000000; // 1,000,000
        uint256 normalDenominator = 1;

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(highNumerator, normalDenominator);

        uint256 smallWbtc = 1 * 10 ** 6; // 0.01 BTC
        vm.prank(user1);
        uint256 largeStrbtc = wbtcConverter.convertWBTCToStrBTC(smallWbtc);

        assertEq(largeStrbtc, smallWbtc * highNumerator / normalDenominator, "High rate conversion incorrect");

        uint256 lowNumerator = 1;
        uint256 highDenominator = 1000000; // 1,000,000

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(lowNumerator, highDenominator);

        uint256 largeWbtc = 100 * 10 ** 8; // 100 BTC
        vm.prank(user1);
        uint256 tinyStrbtc = wbtcConverter.convertWBTCToStrBTC(largeWbtc);

        assertEq(tinyStrbtc, largeWbtc * lowNumerator / highDenominator, "Low rate conversion incorrect");
    }

    function testVerySmallConversion() public {
        uint256 numerator = 1;
        uint256 denominator = 1000;

        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(numerator, denominator);

        uint256 smallAmount = denominator - 1;

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(smallAmount);

        uint256 barelyValidAmount = denominator;

        vm.prank(user1);
        uint256 received = wbtcConverter.convertWBTCToStrBTC(barelyValidAmount);
        assertEq(received, 1, "Minimum conversion should result in 1 token");
    }

    function testMaximumValuesConversion() public {
        vm.prank(rateSetter);
        wbtcConverter.setExchangeRate(1, 1);

        uint256 largeAmount = type(uint64).max;

        MockWBTC(address(wbtc)).mint(user1, largeAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), largeAmount);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(largeAmount);
        vm.stopPrank();

        assertEq(strbtcReceived, largeAmount, "Large amount conversion should work correctly");
    }

    function testUnauthorizedFunctionCalls() public {
        bytes32 rateSetterRole = wbtcConverter.RATE_SETTER_ROLE();

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.emergencyWithdraw(address(wbtc), 1 * 10 ** 8);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.setExchangeRate(2, 1);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.grantRole(rateSetterRole, user2);
    }

    function testConversionWithoutApproval() public {
        uint256 amount = 1 * 10 ** 8; // 1 BTC

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), 0);
        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(amount);
        vm.stopPrank();

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), amount);
        wbtcConverter.convertWBTCToStrBTC(amount);

        strbtc.approve(address(wbtcConverter), 0);
        vm.expectRevert();
        wbtcConverter.convertStrBTCToWBTC(amount);
        vm.stopPrank();
    }
}

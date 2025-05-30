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
    address public user3;
    address public manager;
    address public pauser;

    uint256 public constant INITIAL_WBTC_SUPPLY = 1000 * 10 ** 8; // 1000 WBTC

    bool public useFork;
    uint256 public forkId;
    string public MAINNET_RPC_URL;

    address public constant WBTC_MAINNET = 0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599;

    uint8 public constant FORK_FUZZING_RUNS = 3;

    function setUp() public {
        useFork = vm.envOr("USE_FORK", false);

        if (useFork) {
            MAINNET_RPC_URL = vm.envOr("MAINNET_RPC_URL", string(""));
            require(bytes(MAINNET_RPC_URL).length > 0, "MAINNET_RPC_URL must be set when USE_FORK=true");

            forkId = vm.createFork(MAINNET_RPC_URL);
            vm.selectFork(forkId);

            wbtc = IERC20(WBTC_MAINNET);
        }

        owner = address(this);
        pauser = makeAddr("Pauser");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        manager = makeAddr("manager");

        if (!useFork) {
            MockWBTC mockWBTC = new MockWBTC();
            mockWBTC.mint(address(this), INITIAL_WBTC_SUPPLY);
            mockWBTC.mint(user1, INITIAL_WBTC_SUPPLY);
            mockWBTC.mint(user2, INITIAL_WBTC_SUPPLY);
            mockWBTC.mint(user3, INITIAL_WBTC_SUPPLY);
            wbtc = IERC20(address(mockWBTC));
        } else {
            uint256 amountPerAccount = 1000 * 10 ** 8;
            deal(address(wbtc), address(this), amountPerAccount);
            deal(address(wbtc), user1, amountPerAccount);
            deal(address(wbtc), user2, amountPerAccount);
            deal(address(wbtc), user3, amountPerAccount);
        }

        BitcoinNetworkEncoder.Network network = BitcoinNetworkEncoder.Network.Testnet;

        validatorRegistry = new ValidatorRegistry();

        strBTC strBtcImplementation = new strBTC();

        bytes memory strBtcData =
            abi.encodeWithSelector(strBTC.initialize.selector, network, validatorRegistry, owner, pauser);

        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), owner, strBtcData);
        strbtc = strBTC(address(strBtcProxy));

        WBTCConverter wbtcConverterImplementation = new WBTCConverter();

        bytes memory wbtcConverterData = abi.encodeWithSelector(
            WBTCConverter.initialize.selector, address(wbtc), address(strbtc), owner, manager, pauser
        );

        TransparentUpgradeableProxy wbtcConverterProxy =
            new TransparentUpgradeableProxy(address(wbtcConverterImplementation), owner, wbtcConverterData);
        wbtcConverter = WBTCConverter(address(wbtcConverterProxy));

        strbtc.grantRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));

        vm.prank(user1);
        wbtc.approve(address(wbtcConverter), type(uint256).max);
        vm.prank(user2);
        wbtc.approve(address(wbtcConverter), type(uint256).max);
        vm.prank(user3);
        wbtc.approve(address(wbtcConverter), type(uint256).max);
    }

    function testInitialization() public view {
        assertEq(address(wbtcConverter.wbtc()), address(wbtc), "Incorrect WBTC address");
        assertEq(address(wbtcConverter.strbtc()), address(strbtc), "Incorrect strBTC address");

        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Incorrect exchange rate numerator");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Incorrect exchange rate denominator");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Incorrect exchange rate numerator");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Incorrect exchange rate denominator");

        bytes32 adminRole = wbtcConverter.DEFAULT_ADMIN_ROLE();
        bytes32 managerRole = wbtcConverter.MANAGER_ROLE();
        bytes32 converterRole = strbtc.CONVERTER_ROLE();

        assertTrue(wbtcConverter.hasRole(adminRole, owner), "Owner does not have admin role");

        assertTrue(wbtcConverter.hasRole(managerRole, manager), "manager does not have correct role");

        assertTrue(
            strbtc.hasRole(converterRole, address(wbtcConverter)), "Converter does not have CONVERTER_ROLE in strBTC"
        );

        assertFalse(wbtcConverter.hasRole(adminRole, user1), "Regular user should not have admin role");
        assertFalse(wbtcConverter.hasRole(managerRole, user1), "Regular user should not have rate setter role");

        assertEq(
            wbtc.allowance(user1, address(wbtcConverter)),
            type(uint256).max,
            "User does not have sufficient allowance for WBTC"
        );
    }

    function testConvertWBTCToStrBTC(uint256 amount) public {
        if (useFork) {
            uint256[] memory amounts = new uint256[](3);
            amounts[0] = 0.1 * 10 ** 8; // 0.1 BTC
            amounts[1] = 1 * 10 ** 8; // 1 BTC
            amounts[2] = 2 * 10 ** 8; // 2 BTC

            for (uint8 i = 0; i < amounts.length; i++) {
                _testConvertWBTCToStrBTCWithAmount(amounts[i]);
            }
        } else {
            vm.assume(amount > 0 && amount <= 1000 * 10 ** 8);
            _testConvertWBTCToStrBTCWithAmount(amount);
        }
    }

    function _testConvertWBTCToStrBTCWithAmount(uint256 amount) internal {
        uint256 user1WbtcBalanceBefore = wbtc.balanceOf(user1);
        uint256 user1StrbtcBalanceBefore = strbtc.balanceOf(user1);
        uint256 converterWbtcBalanceBefore = wbtc.balanceOf(address(wbtcConverter));

        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 expectedStrbtc =
            (amount * wbtcConverter.incomingRateNumerator()) / wbtcConverter.incomingRateDenominator();

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

    function testConvertWBTCToStrBTCWithCustomRate(
        uint256 amountInput,
        uint256 numeratorInput,
        uint256 denominatorInput
    ) public {
        if (useFork) {
            uint256[] memory amounts = new uint256[](3);
            amounts[0] = 0.1 * 10 ** 8;
            amounts[1] = 1 * 10 ** 8;
            amounts[2] = 2 * 10 ** 8;

            uint256[] memory numerators = new uint256[](3);
            uint256[] memory denominators = new uint256[](3);

            numerators[0] = 105;
            denominators[0] = 100;
            numerators[1] = 100;
            denominators[1] = 105;
            numerators[2] = 200;
            denominators[2] = 100;

            for (uint8 i = 0; i < FORK_FUZZING_RUNS; i++) {
                _testConvertWithCustomRate(amounts[i % 3], numerators[i % 3], denominators[i % 3]);
            }
        } else {
            uint256 amount = bound(amountInput, 1, 1000 * 10 ** 8);
            uint256 numerator = bound(numeratorInput, 1, 200);
            uint256 denominator = bound(denominatorInput, 1, 200);

            vm.assume((amount * numerator) / denominator > 0);

            _testConvertWithCustomRate(amount, numerator, denominator);
        }
    }

    function _testConvertWithCustomRate(uint256 amount, uint256 numerator, uint256 denominator) internal {
        vm.prank(manager);
        wbtcConverter.setIncomingRate(numerator, denominator);

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

    function testZeroAmountConversion() public {
        vm.prank(user1);
        vm.expectRevert(WBTCConverter.AmountMustBeGreaterThanZero.selector);
        wbtcConverter.convertWBTCToStrBTC(0);

        vm.prank(user1);
        vm.expectRevert(WBTCConverter.AmountMustBeGreaterThanZero.selector);
        wbtcConverter.convertStrBTCToWBTC(0);
    }

    function testConverterRoleRequired() public {
        uint256 amount = 1 * 10 ** 8;
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        strbtc.revokeRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));

        vm.prank(user1);

        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(amount);

        strbtc.grantRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));

        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);
    }

    function testInsufficientWBTCBalance() public {
        uint256 amount = 1 * 10 ** 8;

        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), wbtc.balanceOf(address(wbtcConverter)));

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertStrBTCToWBTC(amount);
    }

    function testSetExchangeRate() public {
        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Initial incoming numerator should be 1");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Initial incoming denominator should be 1");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Initial outgoing numerator should be 1");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Initial outgoing denominator should be 1");

        uint256 newNumerator = 105;
        uint256 newDenominator = 100;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.incomingRateNumerator(), newNumerator, "Incoming numerator not updated correctly");
        assertEq(wbtcConverter.incomingRateDenominator(), newDenominator, "Incoming denominator not updated correctly");
        assertEq(wbtcConverter.outgoingRateNumerator(), newNumerator, "Outgoing numerator not updated correctly");
        assertEq(wbtcConverter.outgoingRateDenominator(), newDenominator, "Outgoing denominator not updated correctly");

        newNumerator = 95;
        newDenominator = 100;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.incomingRateNumerator(), newNumerator, "Incoming numerator not updated correctly");
        assertEq(wbtcConverter.incomingRateDenominator(), newDenominator, "Incoming denominator not updated correctly");
        assertEq(wbtcConverter.outgoingRateNumerator(), newNumerator, "Outgoing numerator not updated correctly");
        assertEq(wbtcConverter.outgoingRateDenominator(), newDenominator, "Outgoing denominator not updated correctly");
    }

    function testSetIncomingRate() public {
        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Initial incoming numerator should be 1");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Initial incoming denominator should be 1");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Initial outgoing numerator should be 1");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Initial outgoing denominator should be 1");

        uint256 newIncomingNumerator = 110;
        uint256 newIncomingDenominator = 100;

        vm.prank(manager);
        wbtcConverter.setIncomingRate(newIncomingNumerator, newIncomingDenominator);

        assertEq(
            wbtcConverter.incomingRateNumerator(), newIncomingNumerator, "Incoming numerator not updated correctly"
        );
        assertEq(
            wbtcConverter.incomingRateDenominator(),
            newIncomingDenominator,
            "Incoming denominator not updated correctly"
        );
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Outgoing numerator should remain unchanged");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Outgoing denominator should remain unchanged");
    }

    function testSetOutgoingRate() public {
        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Initial incoming numerator should be 1");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Initial incoming denominator should be 1");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Initial outgoing numerator should be 1");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Initial outgoing denominator should be 1");

        uint256 newOutgoingNumerator = 95;
        uint256 newOutgoingDenominator = 100;

        vm.prank(manager);
        wbtcConverter.setOutgoingRate(newOutgoingNumerator, newOutgoingDenominator);

        assertEq(
            wbtcConverter.outgoingRateNumerator(), newOutgoingNumerator, "Outgoing numerator not updated correctly"
        );
        assertEq(
            wbtcConverter.outgoingRateDenominator(),
            newOutgoingDenominator,
            "Outgoing denominator not updated correctly"
        );
        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Incoming numerator should remain unchanged");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Incoming denominator should remain unchanged");
    }

    function testConversionWithDifferentRates() public {
        uint256 wbtcAmount = 1 * 10 ** 8; // 1 WBTC

        vm.prank(manager);
        wbtcConverter.setIncomingRate(105, 100);

        vm.prank(manager);
        wbtcConverter.setOutgoingRate(110, 100);

        vm.startPrank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(wbtcAmount);

        uint256 expectedStrBTC = (wbtcAmount * 105) / 100;
        assertEq(strbtcReceived, expectedStrBTC, "Received strBTC amount incorrect with custom incoming rate");

        uint256 expectedWbtcReturn = (strbtcReceived * 100) / 110;

        strbtc.approve(address(wbtcConverter), strbtcReceived);
        uint256 wbtcReturned = wbtcConverter.convertStrBTCToWBTC(strbtcReceived);
        vm.stopPrank();

        assertEq(wbtcReturned, expectedWbtcReturn, "Returned WBTC amount incorrect with custom outgoing rate");

        assertLt(wbtcReturned, wbtcAmount, "With different rates, returned WBTC should be less than original amount");
    }

    function testSetExchangeRateWithZeroNumerator() public {
        uint256 newNumerator = 0;
        uint256 newDenominator = 100;

        vm.prank(manager);
        vm.expectRevert();
        wbtcConverter.setCommonExchangeRate(newNumerator, newDenominator);

        newNumerator = 100;
        newDenominator = 0;

        vm.prank(manager);
        vm.expectRevert();
        wbtcConverter.setCommonExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Denominator should not change");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Denominator should not change");
    }

    function testSetExchangeRateUnauthorized() public {
        uint256 newNumerator = 105;
        uint256 newDenominator = 100;

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.setCommonExchangeRate(newNumerator, newDenominator);

        assertEq(wbtcConverter.incomingRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.incomingRateDenominator(), 1, "Denominator should not change");
        assertEq(wbtcConverter.outgoingRateNumerator(), 1, "Numerator should not change");
        assertEq(wbtcConverter.outgoingRateDenominator(), 1, "Denominator should not change");
    }

    function testRoleManagement() public {
        bytes32 adminRole = wbtcConverter.DEFAULT_ADMIN_ROLE();
        bytes32 managerRole = wbtcConverter.MANAGER_ROLE();

        assertTrue(wbtcConverter.hasRole(adminRole, owner), "Owner should have admin role");
        assertTrue(wbtcConverter.hasRole(managerRole, manager), "manager should have rate setter role");
        assertFalse(wbtcConverter.hasRole(adminRole, user1), "User1 should not have admin role");
        assertFalse(wbtcConverter.hasRole(managerRole, user1), "User1 should not have rate setter role");

        vm.prank(owner);
        wbtcConverter.grantRole(managerRole, user1);
        assertTrue(wbtcConverter.hasRole(managerRole, user1), "User1 should now have rate setter role");

        vm.prank(owner);
        wbtcConverter.revokeRole(managerRole, user1);
        assertFalse(wbtcConverter.hasRole(managerRole, user1), "User1 should no longer have rate setter role");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.grantRole(managerRole, user2);
    }

    function testRenounceRole() public {
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.MANAGER_ROLE(), user1);
        assertTrue(wbtcConverter.hasRole(wbtcConverter.MANAGER_ROLE(), user1), "TestUser should have rate setter role");

        vm.startPrank(user1);
        wbtcConverter.renounceRole(wbtcConverter.MANAGER_ROLE(), user1);
        vm.stopPrank();

        assertFalse(
            wbtcConverter.hasRole(wbtcConverter.MANAGER_ROLE(), user1),
            "TestUser should no longer have rate setter role after renouncing"
        );
    }

    function testAdminRoleManagement() public {
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2);
        assertTrue(wbtcConverter.hasRole(wbtcConverter.DEFAULT_ADMIN_ROLE(), user2), "User2 should now have admin role");

        vm.prank(user2);
        wbtcConverter.grantRole(wbtcConverter.MANAGER_ROLE(), user1);
        assertTrue(
            wbtcConverter.hasRole(wbtcConverter.MANAGER_ROLE(), user1),
            "User1 should have rate setter role granted by new admin"
        );

        vm.prank(user2);
        wbtcConverter.revokeRole(wbtcConverter.MANAGER_ROLE(), user1);
        assertFalse(
            wbtcConverter.hasRole(wbtcConverter.MANAGER_ROLE(), user1),
            "User1 should no longer have rate setter role after revocation by new admin"
        );
    }

    function testConverterRoleOnStrBTC() public {
        assertTrue(
            strbtc.hasRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter)),
            "Converter should have CONVERTER_ROLE on strBTC"
        );

        vm.prank(owner);
        strbtc.revokeRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter));
        assertFalse(
            strbtc.hasRole(strbtc.CONVERTER_ROLE(), address(wbtcConverter)),
            "Converter should no longer have CONVERTER_ROLE after revocation"
        );

        uint256 amount = 1 * 10 ** 8;

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(amount);
    }

    function testPauseContract() public {
        assertFalse(wbtcConverter.paused(), "Contract should not be paused after initialization");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.pause();

        vm.prank(pauser);
        wbtcConverter.pause();

        assertTrue(wbtcConverter.paused(), "Contract should be paused");

        uint256 amount = 1 * 10 ** 8;

        vm.prank(user1);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        wbtcConverter.convertWBTCToStrBTC(amount);

        vm.prank(user1);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        wbtcConverter.convertStrBTCToWBTC(amount);
    }

    function testUnpauseContract() public {
        vm.prank(pauser);
        wbtcConverter.pause();
        assertTrue(wbtcConverter.paused(), "Contract should be paused");

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.unpause();

        vm.prank(pauser);
        wbtcConverter.unpause();

        assertFalse(wbtcConverter.paused(), "Contract should not be paused after unpause");

        uint256 amount = 1 * 10 ** 8;

        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        assertEq(strbtcReceived, amount, "Conversion should work after unpausing");
    }

    function testGrantPauserRole() public {
        vm.prank(owner);
        wbtcConverter.grantRole(wbtcConverter.PAUSER_ROLE(), user2);

        assertTrue(wbtcConverter.hasRole(wbtcConverter.PAUSER_ROLE(), user2), "User should have PAUSER_ROLE");

        vm.prank(user2);
        wbtcConverter.pause();
        assertTrue(wbtcConverter.paused(), "User with PAUSER_ROLE should be able to pause");
    }

    function testEmergencyWithdrawalOnlyAdmin() public {
        uint256 amount = 1 * 10 ** 8;
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
        uint256 amount = 5 * 10 ** 8;
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
        uint256 amount = 10 * 10 ** 8;
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 withdrawAmount = 5 * 10 ** 8;
        vm.prank(owner);
        wbtcConverter.emergencyWithdraw(address(wbtc), withdrawAmount);

        assertEq(
            wbtc.balanceOf(address(wbtcConverter)), amount - withdrawAmount, "Converter should have remaining WBTC"
        );

        uint256 smallAmount = 1 * 10 ** 8;
        vm.prank(user1);
        strbtc.approve(address(wbtcConverter), smallAmount);

        vm.prank(user1);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(smallAmount);

        assertEq(wbtcReceived, smallAmount, "User should receive the correct amount of WBTC");

        uint256 largeAmount = 5 * 10 ** 8;
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
        uint256 highNumerator = 1000000;
        uint256 normalDenominator = 1;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(highNumerator, normalDenominator);

        uint256 smallWbtc = 1 * 10 ** 6;
        vm.prank(user1);
        uint256 largeStrbtc = wbtcConverter.convertWBTCToStrBTC(smallWbtc);

        assertEq(largeStrbtc, smallWbtc * highNumerator / normalDenominator, "High rate conversion incorrect");

        uint256 lowNumerator = 1;
        uint256 highDenominator = 1000000;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(lowNumerator, highDenominator);

        uint256 largeWbtc = 100 * 10 ** 8;
        vm.prank(user1);
        uint256 tinyStrbtc = wbtcConverter.convertWBTCToStrBTC(largeWbtc);

        assertEq(tinyStrbtc, largeWbtc * lowNumerator / highDenominator, "Low rate conversion incorrect");
    }

    function testVerySmallConversion() public {
        uint256 numerator = 1;
        uint256 denominator = 1000;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(numerator, denominator);

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
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 1);
        uint256 largeAmount = 100000 * 10 ** 8;

        deal(address(wbtc), user1, largeAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), largeAmount);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(largeAmount);
        vm.stopPrank();

        assertEq(strbtcReceived, largeAmount, "Large amount conversion should work correctly");
    }

    function testUnauthorizedFunctionCalls() public {
        bytes32 managerRole = wbtcConverter.MANAGER_ROLE();

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.emergencyWithdraw(address(wbtc), 1 * 10 ** 8);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.setCommonExchangeRate(2, 1);

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.grantRole(managerRole, user2);
    }

    function testConversionWithoutApproval() public {
        uint256 amount = 1 * 10 ** 8;

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

    function testFullConversionCycle() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(98, 100);

        uint256 initialAmount = 10 * 10 ** 8;
        uint256 initialUserWbtcBalance = wbtc.balanceOf(user1);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), initialAmount);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(initialAmount);
        vm.stopPrank();

        uint256 expectedStrBtc = (initialAmount * 98) / 100;
        assertEq(strbtcReceived, expectedStrBtc, "strBTC amount incorrect");
        assertEq(strbtc.balanceOf(user1), expectedStrBtc, "User strBTC balance incorrect");
        assertEq(wbtc.balanceOf(user1), initialUserWbtcBalance - initialAmount, "User WBTC balance incorrect");
        assertEq(wbtc.balanceOf(address(wbtcConverter)), initialAmount, "Converter WBTC balance incorrect");

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(100, 99);

        vm.startPrank(user1);
        strbtc.approve(address(wbtcConverter), strbtcReceived);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(strbtcReceived);
        vm.stopPrank();

        uint256 expectedWbtc = (strbtcReceived * 99) / 100;
        assertEq(wbtcReceived, expectedWbtc, "WBTC amount incorrect");
        assertEq(
            wbtc.balanceOf(user1), initialUserWbtcBalance - initialAmount + expectedWbtc, "Final WBTC balance incorrect"
        );
        assertEq(strbtc.balanceOf(user1), 0, "Final strBTC balance incorrect");
        assertEq(
            wbtc.balanceOf(address(wbtcConverter)),
            initialAmount - expectedWbtc,
            "Final converter WBTC balance incorrect"
        );
    }

    function testMultipleUsersConversion() public {
        uint256 initialAmount1 = 5 * 10 ** 8;
        uint256 initialAmount2 = 10 * 10 ** 8;

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), initialAmount1);
        uint256 strbtcReceived1 = wbtcConverter.convertWBTCToStrBTC(initialAmount1);
        vm.stopPrank();

        vm.startPrank(user2);
        wbtc.approve(address(wbtcConverter), initialAmount2);
        uint256 strbtcReceived2 = wbtcConverter.convertWBTCToStrBTC(initialAmount2);
        vm.stopPrank();

        assertEq(strbtc.balanceOf(user1), strbtcReceived1, "User1 strBTC balance incorrect");
        assertEq(strbtc.balanceOf(user2), strbtcReceived2, "User2 strBTC balance incorrect");
        assertEq(
            wbtc.balanceOf(address(wbtcConverter)), initialAmount1 + initialAmount2, "Converter WBTC balance incorrect"
        );

        vm.startPrank(user1);
        strbtc.approve(address(wbtcConverter), strbtcReceived1);
        uint256 wbtcReceived1 = wbtcConverter.convertStrBTCToWBTC(strbtcReceived1);
        vm.stopPrank();

        assertEq(wbtcReceived1, initialAmount1, "User1 WBTC amount incorrect");
        assertEq(strbtc.balanceOf(user1), 0, "User1 strBTC balance should be zero");

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(2, 1);

        vm.startPrank(user2);
        strbtc.approve(address(wbtcConverter), strbtcReceived2);
        uint256 wbtcReceived2 = wbtcConverter.convertStrBTCToWBTC(strbtcReceived2);
        vm.stopPrank();

        assertEq(wbtcReceived2, initialAmount2 / 2, "User2 WBTC amount incorrect with new rate");
        assertEq(strbtc.balanceOf(user2), 0, "User2 strBTC balance should be zero");

        assertEq(
            wbtc.balanceOf(address(wbtcConverter)),
            initialAmount1 + initialAmount2 - wbtcReceived1 - wbtcReceived2,
            "Final converter WBTC balance incorrect"
        );
    }

    function testLargeScaleConversionWorkflow() public {
        address[] memory users = new address[](5);
        for (uint256 i = 0; i < 5; i++) {
            users[i] = address(uint160(uint256(keccak256(abi.encodePacked("user", i)))));
            deal(address(wbtc), users[i], 100 * 10 ** 8);
        }

        uint256[] memory strbtcBalances = new uint256[](5);

        // All users convert different amounts
        for (uint256 i = 0; i < 5; i++) {
            uint256 amount = (i + 1) * 10 * 10 ** 8;

            vm.startPrank(users[i]);
            wbtc.approve(address(wbtcConverter), amount);
            strbtcBalances[i] = wbtcConverter.convertWBTCToStrBTC(amount);
            vm.stopPrank();

            assertEq(strbtcBalances[i], amount, "User should receive equivalent strBTC");
        }

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(2, 1);

        // Another batch of conversions with new rate
        for (uint256 i = 0; i < 3; i++) {
            uint256 amount = (i + 1) * 5 * 10 ** 8;

            vm.startPrank(users[i]);
            wbtc.approve(address(wbtcConverter), amount);
            uint256 additionalStrBTC = wbtcConverter.convertWBTCToStrBTC(amount);
            strbtcBalances[i] += additionalStrBTC;
            vm.stopPrank();

            assertEq(additionalStrBTC, amount * 2, "User should receive strBTC at new rate");
        }

        // Some users convert back
        for (uint256 i = 0; i < 5; i += 2) {
            vm.startPrank(users[i]);
            strbtc.approve(address(wbtcConverter), strbtcBalances[i]);
            uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(strbtcBalances[i]);
            vm.stopPrank();

            uint256 expectedWBTC = strbtcBalances[i] / 2;
            assertEq(wbtcReceived, expectedWBTC, "User should receive WBTC at current rate");
        }

        assertEq(strbtc.balanceOf(users[1]), strbtcBalances[1], "User 1 should still have strBTC");
        assertEq(strbtc.balanceOf(users[3]), strbtcBalances[3], "User 3 should still have strBTC");
    }

    function testOverflowProtection() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(2, 1);

        uint256 normalAmount = 100;
        deal(address(wbtc), user1, normalAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), normalAmount);
        uint256 received = wbtcConverter.convertWBTCToStrBTC(normalAmount);
        assertEq(received, normalAmount * 2, "Should convert correctly");
        vm.stopPrank();

        uint256 maxValue = type(uint256).max;
        uint256 overflowAmount = maxValue / 2 + 1;

        deal(address(wbtc), user2, overflowAmount);

        vm.startPrank(user2);
        wbtc.approve(address(wbtcConverter), overflowAmount);
        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(overflowAmount);
        vm.stopPrank();
    }

    function testExtremelySmallAmount() public {
        uint256 oneWei = 1;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 2);

        uint256 minRequired = 2;

        vm.startPrank(user1);

        vm.expectRevert(WBTCConverter.ConversionResultedInZeroTokens.selector);
        wbtcConverter.convertWBTCToStrBTC(oneWei);

        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(minRequired);
        assertEq(strbtcReceived, minRequired / 2, "Should receive non-zero amount");

        strbtc.approve(address(wbtcConverter), oneWei);

        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(oneWei);
        assertEq(wbtcReceived, oneWei * 2, "Should receive double the amount with 1:2 rate");

        vm.stopPrank();
    }

    function testExtremelyLargeAmount() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 1);

        uint256 largeAmount = type(uint128).max;

        deal(address(wbtc), user1, largeAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), largeAmount);

        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(largeAmount);
        assertEq(strbtcReceived, largeAmount, "Large amount should convert correctly");

        strbtc.approve(address(wbtcConverter), strbtcReceived);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(strbtcReceived);
        assertEq(wbtcReceived, largeAmount, "Converting back should return the same large amount");

        vm.stopPrank();

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(2, 1);

        uint256 overflowRiskAmount = type(uint256).max / 2 + 1;

        deal(address(wbtc), user2, overflowRiskAmount);

        vm.startPrank(user2);
        wbtc.approve(address(wbtcConverter), overflowRiskAmount);

        vm.expectRevert();
        wbtcConverter.convertWBTCToStrBTC(overflowRiskAmount);

        vm.stopPrank();
    }

    function testBoundaryExchangeRates() public {
        vm.startPrank(manager);

        vm.expectRevert(WBTCConverter.DenominatorMustBeGreaterThanZero.selector);
        wbtcConverter.setCommonExchangeRate(1, 0);

        vm.expectRevert(WBTCConverter.NumeratorMustBeGreaterThanZero.selector);
        wbtcConverter.setCommonExchangeRate(0, 1);

        wbtcConverter.setCommonExchangeRate(1, type(uint256).max);

        wbtcConverter.setCommonExchangeRate(type(uint256).max, 1);

        vm.stopPrank();

        uint256 amount = 1000 * 10 ** 8;
        deal(address(wbtc), user1, amount);

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, type(uint256).max);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), amount);

        vm.expectRevert(WBTCConverter.ConversionResultedInZeroTokens.selector);
        wbtcConverter.convertWBTCToStrBTC(amount);

        uint256 safeAmount = type(uint128).max;
        deal(address(wbtc), user1, safeAmount);
        wbtc.approve(address(wbtcConverter), safeAmount);

        vm.expectRevert(WBTCConverter.ConversionResultedInZeroTokens.selector);
        wbtcConverter.convertWBTCToStrBTC(safeAmount);

        vm.stopPrank();
    }

    function testDivisibilityBoundaries() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(2, 3);

        uint256 evenAmount = 3;

        uint256 oddAmount = 2;

        deal(address(wbtc), user1, evenAmount + oddAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), type(uint256).max);

        uint256 strbtcFromEven = wbtcConverter.convertWBTCToStrBTC(evenAmount);
        assertEq(strbtcFromEven, (evenAmount * 2) / 3, "Even division should work");

        uint256 strbtcFromOdd = wbtcConverter.convertWBTCToStrBTC(oddAmount);
        assertEq(strbtcFromOdd, (oddAmount * 2) / 3, "Should handle division with remainder");

        vm.stopPrank();

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(3, 2);

        vm.startPrank(user1);
        deal(address(wbtc), user1, 5);
        wbtcConverter.convertWBTCToStrBTC(5);

        strbtc.approve(address(wbtcConverter), 5);

        uint256 wbtcFromEven = wbtcConverter.convertStrBTCToWBTC(3);
        assertEq(wbtcFromEven, (3 * 2) / 3, "Even division in reverse should work");

        uint256 wbtcFromOdd = wbtcConverter.convertStrBTCToWBTC(2);
        uint256 two = 2;
        uint256 expectedWbtcFromOdd = (two * two) / 3;
        assertEq(wbtcFromOdd, expectedWbtcFromOdd, "Should handle reverse division with remainder");

        vm.stopPrank();
    }

    function testRateChangeDecimals() public {
        uint256 initialAmount = 10 * 10 ** 8;

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 1);

        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(initialAmount);

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1000000, 999999);

        uint256 preciseAmount = 999999;
        deal(address(wbtc), user2, preciseAmount);

        vm.startPrank(user2);
        MockWBTC(address(wbtc)).approve(address(wbtcConverter), preciseAmount);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(preciseAmount);
        vm.stopPrank();

        assertEq(strbtcReceived, 1000000, "Should receive precise amount based on rate");

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(9999999, 10000000);

        uint256 inputAmount = 9999999;
        deal(address(wbtc), user3, inputAmount);

        vm.startPrank(user3);
        wbtc.approve(address(wbtcConverter), inputAmount);
        uint256 user3StrBtcReceived = wbtcConverter.convertWBTCToStrBTC(inputAmount);

        strbtc.approve(address(wbtcConverter), user3StrBtcReceived);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(user3StrBtcReceived);
        vm.stopPrank();

        assertEq(wbtcReceived, 9999998, "Should receive correct amount based on conversion calculation");
    }

    function testPrecisionLossLimits() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 1000000000);

        uint256 justUnderMin = 999999999;
        deal(address(wbtc), user1, justUnderMin);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), justUnderMin);

        vm.expectRevert(WBTCConverter.ConversionResultedInZeroTokens.selector);
        wbtcConverter.convertWBTCToStrBTC(justUnderMin);
        vm.stopPrank();

        uint256 minRequired = 1000000000;
        deal(address(wbtc), user1, minRequired);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), minRequired);

        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(minRequired);
        assertEq(strbtcReceived, 1, "Should receive 1 token at minimum threshold");

        vm.stopPrank();

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 1);

        uint256 standardAmount = 1000000000;
        deal(address(wbtc), user3, standardAmount);

        vm.startPrank(user3);
        wbtc.approve(address(wbtcConverter), standardAmount);
        uint256 strbtcAmount = wbtcConverter.convertWBTCToStrBTC(standardAmount);
        assertEq(strbtcAmount, standardAmount, "Should receive same amount at 1:1 rate");
        vm.stopPrank();

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1000000000, 1);

        uint256 smallAmount = 1;
        deal(address(wbtc), address(wbtcConverter), 1000000000);

        vm.startPrank(user3);
        strbtc.approve(address(wbtcConverter), strbtcAmount);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(strbtcAmount);
        vm.stopPrank();

        assertEq(wbtcReceived, smallAmount, "Should receive minimal amount at 1B:1 rate");
    }

    function testRedemptionRoundingErrors() public {
        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(3, 10);

        uint256 initialAmount = 10;
        deal(address(wbtc), user1, initialAmount);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), initialAmount);

        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(initialAmount);
        assertEq(strbtcReceived, 3, "Should receive 3 strBTC for 10 WBTC");

        strbtc.approve(address(wbtcConverter), strbtcReceived);
        uint256 wbtcReceived = wbtcConverter.convertStrBTCToWBTC(strbtcReceived);

        assertEq(wbtcReceived, 10, "Converting back should yield original amount");

        vm.stopPrank();

        vm.prank(manager);
        wbtcConverter.setCommonExchangeRate(1, 3);

        deal(address(wbtc), user1, 10);

        vm.startPrank(user1);
        wbtc.approve(address(wbtcConverter), 10);

        uint256 strbtcReceived2 = wbtcConverter.convertWBTCToStrBTC(10);
        assertEq(strbtcReceived2, 3, "Should receive 3 strBTC due to truncation");

        strbtc.approve(address(wbtcConverter), strbtcReceived2);
        uint256 wbtcReceived2 = wbtcConverter.convertStrBTCToWBTC(strbtcReceived2);

        assertEq(wbtcReceived2, 9, "Should receive 9 WBTC due to previous truncation");
        assertLt(wbtcReceived2, 10, "Should have lost some value due to rounding");

        vm.stopPrank();
    }

    function testTokenTracking() public {
        uint256 initialMinted = wbtcConverter.totalMinted();
        uint256 initialBurned = wbtcConverter.totalBurned();
        uint256 initialNetMinted = wbtcConverter.currentlyMinted();

        uint256 amount = 1 * 10 ** 8;
        vm.prank(user1);
        uint256 strbtcReceived = wbtcConverter.convertWBTCToStrBTC(amount);

        assertEq(wbtcConverter.totalMinted(), initialMinted + strbtcReceived, "Total minted not updated correctly");
        assertEq(wbtcConverter.totalBurned(), initialBurned, "Total burned should not change");
        assertEq(
            wbtcConverter.currentlyMinted(), initialNetMinted + strbtcReceived, "Currently minted not updated correctly"
        );

        vm.startPrank(user1);
        strbtc.approve(address(wbtcConverter), strbtcReceived);
        wbtcConverter.convertStrBTCToWBTC(strbtcReceived);
        vm.stopPrank();

        assertEq(
            wbtcConverter.totalMinted(), initialMinted + strbtcReceived, "Total minted should not change after burn"
        );
        assertEq(wbtcConverter.totalBurned(), initialBurned + strbtcReceived, "Total burned not updated correctly");
        assertEq(wbtcConverter.currentlyMinted(), initialNetMinted, "Currently minted should return to initial value");
    }

    function testSetMintingLimit() public {
        uint256 initialLimit = wbtcConverter.mintingLimit();

        assertEq(initialLimit, type(uint256).max, "Initial minting limit should be max uint256");

        uint256 newLimit = 100 * 10 ** 8;

        vm.prank(user1);
        vm.expectRevert();
        wbtcConverter.setMintingLimit(newLimit);

        vm.prank(manager);
        wbtcConverter.setMintingLimit(newLimit);

        assertEq(wbtcConverter.mintingLimit(), newLimit, "Minting limit not updated correctly");
    }

    function testMintingLimitEnforcement() public {
        uint256 mintLimit = 2 * 10 ** 8;
        vm.prank(manager);
        wbtcConverter.setMintingLimit(mintLimit);

        uint256 amount1 = 1 * 10 ** 8;
        vm.prank(user1);
        wbtcConverter.convertWBTCToStrBTC(amount1);

        uint256 amount2 = 1.5 * 10 ** 8;
        vm.prank(user2);
        vm.expectRevert(WBTCConverter.MintingLimitExceeded.selector);
        wbtcConverter.convertWBTCToStrBTC(amount2);

        uint256 remainingLimit = mintLimit - amount1;
        vm.prank(user2);
        wbtcConverter.convertWBTCToStrBTC(remainingLimit);

        assertEq(wbtcConverter.currentlyMinted(), mintLimit, "Current minted amount should equal the limit");

        uint256 minAmount = 1;
        vm.prank(user3);
        vm.expectRevert(WBTCConverter.MintingLimitExceeded.selector);
        wbtcConverter.convertWBTCToStrBTC(minAmount);

        uint256 burnAmount = 0.5 * 10 ** 8;
        vm.startPrank(user2);
        strbtc.approve(address(wbtcConverter), burnAmount);
        wbtcConverter.convertStrBTCToWBTC(burnAmount);
        vm.stopPrank();

        vm.prank(user3);
        wbtcConverter.convertWBTCToStrBTC(burnAmount);
    }

    function testComplexScenarioWithRatesAndLimits() public {
        vm.startPrank(manager);
        wbtcConverter.setIncomingRate(120, 100);
        wbtcConverter.setOutgoingRate(100, 90);
        wbtcConverter.setMintingLimit(10 * 10 ** 8);
        vm.stopPrank();

        uint256 user1Amount = 2 * 10 ** 8;
        uint256 user2Amount = 3 * 10 ** 8;

        vm.prank(user1);
        uint256 user1StrBTC = wbtcConverter.convertWBTCToStrBTC(user1Amount);
        assertEq(user1StrBTC, (user1Amount * 120) / 100, "User1 received incorrect strBTC amount");

        vm.prank(user2);
        uint256 user2StrBTC = wbtcConverter.convertWBTCToStrBTC(user2Amount);
        assertEq(user2StrBTC, (user2Amount * 120) / 100, "User2 received incorrect strBTC amount");

        uint256 expectedMinted = user1StrBTC + user2StrBTC;
        assertEq(wbtcConverter.currentlyMinted(), expectedMinted, "Currently minted tracked incorrectly");

        uint256 user1ConvertBack = user1StrBTC / 2;
        vm.startPrank(user1);
        strbtc.approve(address(wbtcConverter), user1ConvertBack);
        uint256 user1WBTCReceived = wbtcConverter.convertStrBTCToWBTC(user1ConvertBack);
        vm.stopPrank();

        assertEq(
            user1WBTCReceived, (user1ConvertBack * 90) / 100, "User1 received incorrect WBTC amount on conversion back"
        );

        expectedMinted = expectedMinted - user1ConvertBack;
        assertEq(wbtcConverter.currentlyMinted(), expectedMinted, "Currently minted incorrect after burn");

        uint256 remainingStrBTC = wbtcConverter.mintingLimit() - wbtcConverter.currentlyMinted();
        uint256 maxMoreWBTC = (remainingStrBTC * 100) / 120;

        uint256 largeAmount = maxMoreWBTC + 1 * 10 ** 7;

        vm.prank(user3);
        vm.expectRevert(WBTCConverter.MintingLimitExceeded.selector);
        wbtcConverter.convertWBTCToStrBTC(largeAmount);

        vm.prank(user3);
        wbtcConverter.convertWBTCToStrBTC(maxMoreWBTC);

        uint256 finalMinted = wbtcConverter.currentlyMinted();
        assertApproxEqAbs(finalMinted, wbtcConverter.mintingLimit(), 1, "Should reach approximately the minting limit");
    }
}

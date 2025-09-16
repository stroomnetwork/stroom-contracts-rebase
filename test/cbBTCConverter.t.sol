// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {BitcoinNetworkEncoder} from "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";
import "../src/strBTC.sol";
import "../src/converters/cbBTCConverterImmutable.sol";
import "../src/lib/ValidatorRegistry.sol";

contract MockCBBTC is ERC20 {
    constructor() ERC20("Coinbase Wrapped BTC", "cbBTC") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract CBBTCConverterTest is Test {
    CBBTCConverterImmutable public cbbtcConverter;
    strBTC public strbtc;
    IERC20 public cbbtc;

    ValidatorRegistry public validatorRegistry;
    address public owner;
    address public user1;
    address public user2;
    address public user3;
    address public manager;
    address public pauser;
    address public withdrawer;

    uint256 public constant INITIAL_CBBTC_SUPPLY = 1000 * 10 ** 8;

    function setUp() public {
        owner = makeAddr("owner");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        manager = makeAddr("manager");
        pauser = makeAddr("pauser");
        withdrawer = makeAddr("withdrawer");

        vm.startPrank(owner);

        validatorRegistry = new ValidatorRegistry();
        validatorRegistry.setJointPublicKey(
            bytes32(uint256(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef))
        );

        strBTC strBTCImpl = new strBTC();
        bytes memory strBTCData = abi.encodeWithSignature(
            "initialize(uint8,address,address,address)",
            uint8(BitcoinNetworkEncoder.Network.Mainnet),
            address(validatorRegistry),
            owner,
            pauser
        );

        TransparentUpgradeableProxy strBTCProxy =
            new TransparentUpgradeableProxy(address(strBTCImpl), owner, strBTCData);

        strbtc = strBTC(address(strBTCProxy));

        MockCBBTC mockCBBTC = new MockCBBTC();
        cbbtc = IERC20(address(mockCBBTC));

        cbbtcConverter = new CBBTCConverterImmutable(address(cbbtc), address(strbtc), withdrawer);

        strbtc.addConverter(address(cbbtcConverter));

        MockCBBTC(address(cbbtc)).mint(user1, INITIAL_CBBTC_SUPPLY / 3);
        MockCBBTC(address(cbbtc)).mint(user2, INITIAL_CBBTC_SUPPLY / 3);
        MockCBBTC(address(cbbtc)).mint(user3, INITIAL_CBBTC_SUPPLY / 3);

        vm.stopPrank();
    }

    function testInitialState() public view {
        assertEq(address(cbbtcConverter.strbtc()), address(strbtc));
        assertEq(address(cbbtcConverter.cbbtc()), address(cbbtc));
        assertEq(cbbtcConverter.withdrawer(), withdrawer);
        assertEq(cbbtcConverter.mintingLimit(), 500 * 10 ** 8);
        assertEq(cbbtcConverter.totalMinted(), 0);
        assertEq(cbbtcConverter.totalBurned(), 0);
        assertEq(cbbtcConverter.currentlyMinted(), 0);

        assertEq(cbbtcConverter.incomingRateNumerator(), 999);
        assertEq(cbbtcConverter.incomingRateDenominator(), 1000);
        assertEq(cbbtcConverter.outgoingRateNumerator(), 1000);
        assertEq(cbbtcConverter.outgoingRateDenominator(), 999);
    }

    function testConvertCBBTCToStrBTC() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;
        uint256 expectedStrBTC = (cbbtcAmount * 999) / 1000;

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);

        uint256 strbtcReceived = cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);

        assertEq(strbtcReceived, expectedStrBTC);
        assertEq(strbtc.balanceOf(user1), expectedStrBTC);
        assertEq(cbbtc.balanceOf(user1), INITIAL_CBBTC_SUPPLY / 3 - cbbtcAmount);
        assertEq(cbbtc.balanceOf(address(cbbtcConverter)), cbbtcAmount);
        assertEq(cbbtcConverter.totalMinted(), expectedStrBTC);
        assertEq(cbbtcConverter.currentlyMinted(), expectedStrBTC);
        vm.stopPrank();
    }

    function testConvertStrBTCToCBBTC() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;
        uint256 strbtcAmount = (cbbtcAmount * 999) / 1000;

        // First convert cbBTC to strBTC
        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();

        // Then convert strBTC back to cbBTC
        vm.startPrank(user1);
        strbtc.approve(address(cbbtcConverter), strbtcAmount);

        uint256 cbbtcReceived = cbbtcConverter.convertStrBTCToCBBTC(strbtcAmount);

        uint256 expectedCBBTC = (strbtcAmount * 999) / 1000;
        assertEq(cbbtcReceived, expectedCBBTC);
        assertEq(cbbtc.balanceOf(user1), INITIAL_CBBTC_SUPPLY / 3 - cbbtcAmount + expectedCBBTC);
        assertEq(strbtc.balanceOf(user1), 0);
        assertEq(cbbtcConverter.totalBurned(), strbtcAmount);
        assertEq(cbbtcConverter.currentlyMinted(), 0);
        vm.stopPrank();
    }

    function testConvertCBBTCToStrBTCWithAmount(uint256 amount) public {
        vm.assume(amount > 0 && amount <= 100 * 10 ** 8);
        vm.assume(amount >= 1000);
        vm.assume(amount <= cbbtc.balanceOf(user1));

        uint256 expectedStrBTC = (amount * 999) / 1000;
        vm.assume(expectedStrBTC > 0);

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), amount);

        uint256 strbtcReceived = cbbtcConverter.convertCBBTCToStrBTC(amount);

        assertEq(strbtcReceived, expectedStrBTC);
        assertEq(strbtc.balanceOf(user1), expectedStrBTC);
        vm.stopPrank();
    }

    function testConvertZeroAmount() public {
        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), 0);

        vm.expectRevert(CBBTCConverterImmutable.AmountMustBeGreaterThanZero.selector);
        cbbtcConverter.convertCBBTCToStrBTC(0);
        vm.stopPrank();
    }

    function testConvertWithoutApproval() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;

        vm.startPrank(user1);
        vm.expectRevert();
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();
    }

    function testMintingLimit() public {
        uint256 limit = cbbtcConverter.mintingLimit();
        uint256 cbbtcAmount = (limit * 1000) / 999 + 1000;

        MockCBBTC(address(cbbtc)).mint(user1, cbbtcAmount);

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);

        vm.expectRevert(CBBTCConverterImmutable.MintingLimitExceeded.selector);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();
    }

    function testInsufficientCBBTCBalance() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;
        uint256 strbtcAmount = (cbbtcAmount * 999) / 1000;

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();

        vm.startPrank(withdrawer);
        cbbtcConverter.withdraw(address(cbbtc), withdrawer, cbbtcAmount);
        vm.stopPrank();

        vm.startPrank(user1);
        strbtc.approve(address(cbbtcConverter), strbtcAmount);

        vm.expectRevert(CBBTCConverterImmutable.InsufficientCBBTCBalance.selector);
        cbbtcConverter.convertStrBTCToCBBTC(strbtcAmount);
        vm.stopPrank();
    }

    function testWithdraw() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();

        uint256 balanceBefore = cbbtc.balanceOf(withdrawer);

        vm.startPrank(withdrawer);
        cbbtcConverter.withdraw(address(cbbtc), withdrawer, cbbtcAmount);
        vm.stopPrank();

        assertEq(cbbtc.balanceOf(withdrawer), balanceBefore + cbbtcAmount);
        assertEq(cbbtc.balanceOf(address(cbbtcConverter)), 0);
    }

    function testWithdrawNotWithdrawer() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;

        vm.startPrank(user1);
        vm.expectRevert(CBBTCConverterImmutable.NotWithdrawer.selector);
        cbbtcConverter.withdraw(address(cbbtc), user1, cbbtcAmount);
        vm.stopPrank();
    }

    function testMultipleConversions() public {
        uint256 cbbtcAmount1 = 1 * 10 ** 8;
        uint256 cbbtcAmount2 = 2 * 10 ** 8;

        uint256 expectedStrBTC1 = (cbbtcAmount1 * 999) / 1000;
        uint256 expectedStrBTC2 = (cbbtcAmount2 * 999) / 1000;
        uint256 totalExpected = expectedStrBTC1 + expectedStrBTC2;

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount1);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount1);
        vm.stopPrank();

        vm.startPrank(user2);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount2);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount2);
        vm.stopPrank();

        assertEq(strbtc.balanceOf(user1), expectedStrBTC1);
        assertEq(strbtc.balanceOf(user2), expectedStrBTC2);
        assertEq(cbbtcConverter.totalMinted(), totalExpected);
        assertEq(cbbtcConverter.currentlyMinted(), totalExpected);
    }

    function testEvents() public {
        uint256 cbbtcAmount = 1 * 10 ** 8;
        uint256 expectedStrBTC = (cbbtcAmount * 999) / 1000;

        vm.startPrank(user1);
        cbbtc.approve(address(cbbtcConverter), cbbtcAmount);

        vm.expectEmit(true, false, false, true);
        emit CBBTCConverterImmutable.CBBTCConverted(user1, cbbtcAmount, expectedStrBTC);
        cbbtcConverter.convertCBBTCToStrBTC(cbbtcAmount);
        vm.stopPrank();
    }
}

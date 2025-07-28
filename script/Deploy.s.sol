// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.27;

import "forge-std/Script.sol";

import "../src/strBTC.sol";
import "../src/wstrBTC.sol";
import "../src/lib/UserActivator.sol";
import "../src/lib/ValidatorRegistry.sol";
import "../src/wBTCConverterImmutable.sol";
import "../lib/blockchain-tools/src/BitcoinNetworkEncoder.sol";

import {Strings as OpenZeppelinStrings} from "@openzeppelin/contracts/utils/Strings.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployScript is Script {
    BitcoinNetworkEncoder.Network private network;
    address admin = vm.envAddress("ADMIN_ADDRESS"); // msg.sender
    address pauser = vm.envAddress("PAUSER_ADDRESS");
    address manager = vm.envAddress("MANAGER_ADDRESS");
    address wbtcAddress = vm.envAddress("WBTC_ADDRESS");
    address timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");

    function setUp() public {
        network = BitcoinNetworkEncoder.Network(vm.envUint("BITCOIN_NETWORK"));
        console.log("Deploying with Bitcoin network:", uint256(network));
    }

    function run() public {
        vm.startBroadcast();

        ValidatorRegistry vr = new ValidatorRegistry();

        // Deploy strBTC implementation
        strBTC strBtcImplementation = new strBTC();

        // Deploy strBTC proxy
        bytes memory strBtcData = abi.encodeWithSelector(strBTC.initialize.selector, network, vr, admin, pauser);
        TransparentUpgradeableProxy strBtcProxy =
            new TransparentUpgradeableProxy(address(strBtcImplementation), admin, strBtcData);
        strBTC strBtcContract = strBTC(address(strBtcProxy));

        // Get strBTC ProxyAdmin address
        address strBtcProxyAdmin = _getProxyAdmin(address(strBtcProxy));

        // Deploy wstrBTC
        wstrBTC wstrBtcContract = new wstrBTC(address(strBtcContract));

        UserActivator activator = new UserActivator();

        // Deploy wBTCConverter
        WBTCConverterImmutable wBtcConverterContract =
            new WBTCConverterImmutable(wbtcAddress, address(strBtcContract), timelock);

        // Grant CONVERTER_ROLE to wBTCConverter in strBTC
        strBtcContract.grantRole(strBtcContract.CONVERTER_ROLE(), address(wBtcConverterContract));

        vm.stopBroadcast();

        console.log();
        console.log("Deployed contract addresses as envs:");
        console.log();
        console.logString(
            string.concat(
                "APP_ETH_STRBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(strBtcContract))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_WSTRBTC_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(wstrBtcContract))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_VALIDATOR_REGISTRY_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(vr))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_USER_ACTIVATOR_ADDRESS=", OpenZeppelinStrings.toHexString(uint256(uint160(address(activator))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_WBTC_CONVERTER_ADDRESS=",
                OpenZeppelinStrings.toHexString(uint256(uint160(address(wBtcConverterContract))))
            )
        );
        console.logString(
            string.concat(
                "APP_ETH_STRBTC_PROXY_ADMIN_ADDRESS=",
                OpenZeppelinStrings.toHexString(uint256(uint160(strBtcProxyAdmin)))
            )
        );
    }

    /**
     * @dev Get ProxyAdmin address from TransparentUpgradeableProxy
     * @param proxyAddress Address of the proxy contract
     * @return Address of the ProxyAdmin contract
     */
    function _getProxyAdmin(address proxyAddress) private view returns (address) {
        // ERC1967 admin storage slot
        // keccak256("eip1967.proxy.admin") - 1
        bytes32 ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;

        return address(uint160(uint256(vm.load(proxyAddress, ADMIN_SLOT))));
    }
}

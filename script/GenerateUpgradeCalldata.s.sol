// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "forge-std/Script.sol";
import "forge-std/console.sol";

contract GenerateUpgradeCalldata is Script {
    // Timelock context
    address public timelock = vm.envAddress("APP_ETH_TIMELOCK_ADDRESS");
    uint256 public delay = vm.envUint("TIMELOCK_DELAY");

    // ProxyAdmin target and proxy/impl
    address public proxyAdmin = vm.envAddress("APP_ETH_STRBTC_PROXY_ADMIN_ADDRESS");
    address public strBTCProxy = vm.envAddress("APP_ETH_STRBTC_ADDRESS");
    address public newStrBTCImpl = vm.envAddress("NEW_STRBTC_IMPL");

    // Timelock schedule params
    bytes32 public PREDECESSOR = vm.envOr("PREDECESSOR", bytes32(0));
    bytes32 public SALT = vm.envOr("SALT", bytes32(0));

    function run() external view {
        bytes memory callData = abi.encodeWithSignature("reinitializeV2()");

        bytes memory upgradeAndCallData =
            abi.encodeWithSignature("upgradeAndCall(address,address,bytes)", strBTCProxy, newStrBTCImpl, callData);

        bytes32 operationId = keccak256(abi.encode(proxyAdmin, 0, upgradeAndCallData, PREDECESSOR, SALT));

        console.log("Target:", proxyAdmin);
        console.log("Value: 0");
        console.log("Data:", vm.toString(upgradeAndCallData));
        console.log("Predecessor:", vm.toString(PREDECESSOR));
        console.log("Salt:", vm.toString(SALT));
        console.log("Delay:", delay);
        console.log("");
        console.log("Operation ID:", vm.toString(operationId));
    }
}

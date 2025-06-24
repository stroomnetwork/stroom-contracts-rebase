// script/CreateStrBTCWBTCPool.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";

interface IUniswapV3Factory {
    function createPool(address tokenA, address tokenB, uint24 fee) external returns (address pool);
    function getPool(address tokenA, address tokenB, uint24 fee) external view returns (address);
}

interface IUniswapV3Pool {
    function initialize(uint160 sqrtPriceX96) external;
    function token0() external view returns (address);
    function token1() external view returns (address);
}

interface INonfungiblePositionManager {
    struct MintParams {
        address token0;
        address token1;
        uint24 fee;
        int24 tickLower;
        int24 tickUpper;
        uint256 amount0Desired;
        uint256 amount1Desired;
        uint256 amount0Min;
        uint256 amount1Min;
        address recipient;
        uint256 deadline;
    }

    function mint(MintParams calldata params)
        external
        payable
        returns (uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1);
}

interface IERC20 {
    function approve(address spender, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

contract CreateStrBTCWBTCPoolScript is Script {
    address constant FACTORY = 0x0227628f3F023bb0B980b67D528571c95c6DaC1c;
    address constant POSITION_MANAGER = 0x1238536071E1c677A632429e3655c799b22cDA52;

    address public wbtc = vm.envAddress("WBTC_ADDRESS");
    address public strbtc = vm.envAddress("APP_ETH_STRBTC_ADDRESS");

    uint24 constant FEE = 500;

    int24 constant TICK_LOWER = -500; // ≈ 0.95
    int24 constant TICK_UPPER = 500; // ≈ 1.05

    function run() external {
        vm.startBroadcast();

        address pool = _createPool();
        _addLiquidity(pool);

        vm.stopBroadcast();
    }

    function _createPool() internal returns (address pool) {
        IUniswapV3Factory factory = IUniswapV3Factory(FACTORY);

        pool = factory.getPool(wbtc, strbtc, FEE);

        if (pool != address(0)) {
            console.log("Pool already exists:", pool);
            return pool;
        }

        pool = factory.createPool(wbtc, strbtc, FEE);
        console.log("Pool created:", pool);

        // Initialize withrice price
        // sqrtPriceX96 = sqrt(price) * 2^96
        // For price = 1: sqrtPriceX96 = 79228162514264337593543950336
        uint160 sqrtPriceX96 = 79228162514264337593543950336;

        IUniswapV3Pool(pool).initialize(sqrtPriceX96);

        console.log("Pool initialized with 1:1 price");

        return pool;
    }

    function _addLiquidity(address pool) internal {
        IUniswapV3Pool poolContract = IUniswapV3Pool(pool);

        address token0 = poolContract.token0();
        address token1 = poolContract.token1();

        uint256 amount0Desired = 100000000000; // 1000 token (8 decimals)
        uint256 amount1Desired = 100000000000; // 1000 token (8 decimals)

        // Approve tokens
        IERC20(token0).approve(POSITION_MANAGER, amount0Desired);
        IERC20(token1).approve(POSITION_MANAGER, amount1Desired);

        INonfungiblePositionManager.MintParams memory params = INonfungiblePositionManager.MintParams({
            token0: token0,
            token1: token1,
            fee: FEE,
            tickLower: TICK_LOWER,
            tickUpper: TICK_UPPER,
            amount0Desired: amount0Desired,
            amount1Desired: amount1Desired,
            amount0Min: amount0Desired * 95 / 100, // 5% slippage
            amount1Min: amount1Desired * 95 / 100, // 5% slippage
            recipient: msg.sender,
            deadline: block.timestamp + 300
        });

        (uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1) =
            INonfungiblePositionManager(POSITION_MANAGER).mint(params);

        console.log("Liquidity added:");
        console.log("Position ID:", tokenId);
        console.log("Liquidity:", liquidity);
        console.log("Amount0 used:", amount0);
        console.log("Amount1 used:", amount1);
        console.log("Tick range:", TICK_LOWER);
        console.log("to", TICK_UPPER);
        console.log("Price range: ~0.95 to ~1.05");
    }
}

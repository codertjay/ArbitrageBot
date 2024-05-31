// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {DeployArbitrage} from "script/DeployFlashLoanArbitrage.s.sol";
import {FlashLoanArbitrage} from "src/FlashLoanArbitrage.sol";
import {Test, console} from "forge-std/Test.sol";

contract ArbitrageTest is Test {
    //    DeployArbitrage deployer;
    FlashLoanArbitrage public arbitrage;

    address public immutable i_deployed_arbitrage = 0x98aD59d5d63e5951dF8572b68772eA5b401c12ff;

    address public immutable i_startSwapAddress = 0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff;
    address public immutable i_endSwapAddress = 0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506;
    address public constant i_token0 = 0xc2132D05D31c914a87C6611C10748AEb04B58e8F;
    address public constant i_token1 = 0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270;

    uint256 public constant i_arbitrage_amount = 1e18;

    function setUp() external {
        arbitrage = FlashLoanArbitrage(payable(i_deployed_arbitrage));
    }

    function testArbitrage() public {
        console.log("Arbitrage address: %s", address(arbitrage));
        console.log("Arbitrage owner: %s", address(arbitrage.owner()));

        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        arbitrage.makeFlashLoan(i_startSwapAddress, i_endSwapAddress, i_token0, i_token1, i_arbitrage_amount);

        vm.stopBroadcast();

    }
}

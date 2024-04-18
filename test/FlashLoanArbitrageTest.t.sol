// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {DeployArbitrage} from "script/DeployFlashLoanArbitrage.s.sol";
import {FlashLoanArbitrage} from "src/FlashLoanArbitrage.sol";
import {Test, console} from "forge-std/Test.sol";

contract ArbitrageTest is Test {
    //    DeployArbitrage deployer;
    FlashLoanArbitrage public arbitrage;

    address public immutable i_deployed_arbitrage = 0x3dd2C59690B65b6bB13c14C169832A84FE121A10;

    address public immutable i_startSwapAddress = 0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2;
    address public immutable i_endSwapAddress = 0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506;
    address public constant i_token0 = 0x03AA6298F1370642642415EDC0db8b957783e8D6;
    address public constant i_token1 = 0x2dDb89a10Bf2020d8CaE7C5d239b6F38bE9d91D9;

    uint256 public constant i_arbitrage_amount = 1e18;

    function setUp() external {
        arbitrage = new FlashLoanArbitrage();
        //        arbitrage = Arbitrage(payable(i_deployed_arbitrage));
    }

    function testArbitrage() public {
        vm.startPrank(address(arbitrage.owner()));
        arbitrage.makeFlashLoan(i_startSwapAddress, i_endSwapAddress, i_token0, i_token1, i_arbitrage_amount);
        vm.stopPrank();

        console.log("Arbitrage address: %s", address(arbitrage));
        console.log("Arbitrage address: %s", address(arbitrage.owner()));
    }
}

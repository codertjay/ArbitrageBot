// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {DeployArbitrage} from "script/DeployFlashLoanArbitrage.s.sol";
import {FlashLoanArbitrage, IERC20} from "src/FlashLoanArbitrage.sol";
import {Test, console} from "forge-std/Test.sol";


contract ArbitrageTest is Test {
    //    DeployArbitrage deployer;
    FlashLoanArbitrage public arbitrage;

    address public immutable i_deployed_arbitrage = 0xA617BE9Cc7bC4bEc277Fb50e20C9376556a279d4;

    address public immutable i_startSwapAddress = 0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506;
    address public immutable i_endSwapAddress = 0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff;
    address public constant i_token0 = 0xc2132D05D31c914a87C6611C10748AEb04B58e8F;
    address public constant i_token1 = 0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270;

    uint256 public constant i_arbitrage_amount = 1e6;

    function setUp() external {
        // arbitrage = FlashLoanArbitrage(payable(i_deployed_arbitrage));
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        arbitrage = new FlashLoanArbitrage();
        vm.stopBroadcast();
    }

    function testArbitrage() public {
        FlashLoanArbitrage.MakeInput[] memory inputs = new FlashLoanArbitrage.MakeInput[](1);

        inputs[0] = FlashLoanArbitrage.MakeInput(
            keccak256(abi.encodePacked("M1")),
            keccak256(abi.encodePacked("M2")),
            i_token1,
            i_arbitrage_amount,
            5
        );


        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        arbitrage.milking(inputs);
        vm.stopBroadcast();
    }

    function testWithdrawERCR20Token() public {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        IERC20 usdt = IERC20(i_token0);

        // Check initial contract balance
        uint256 initialBalance = usdt.balanceOf(address(arbitrage));
        console.log("Initial contract balance: ", initialBalance);

        // Check initial balance of the recipient
        uint256 initialRecipientBalance = usdt.balanceOf(
            address(arbitrage.owner())
        );
        console.log("Initial recipient balance: ", initialRecipientBalance);

        arbitrage.withdrawERC20(i_token0, initialBalance);

        // Check final contract balance
        uint256 finalBalance = usdt.balanceOf(address(arbitrage));
        console.log("Final contract balance: ", finalBalance);

        // Check final balance of the recipient
        uint256 finalRecipientBalance = usdt.balanceOf(
            address(arbitrage.owner())
        );
        console.log("Final recipient balance: ", finalRecipientBalance);

        vm.stopBroadcast();

        // Assert the transfer was successful
        assertEq(
            finalRecipientBalance,
            initialRecipientBalance + initialBalance
        );
    }

    function testCheckProfitability() public {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        // Encode the string keys "M1" and "M2" to bytes32
        bytes32 key1 = keccak256(abi.encodePacked("M1"));
        bytes32 key2 = keccak256(abi.encodePacked("M2"));
        bytes32 key3 = keccak256(abi.encodePacked("M3"));

        console.logBytes32(key1);
        console.logBytes32(key2);
        console.logBytes32(key3);

        arbitrage.checkProfitability(
            key1,
            key2,
            i_token1,
            i_arbitrage_amount,
            500
        );

        vm.stopBroadcast();
    }



}

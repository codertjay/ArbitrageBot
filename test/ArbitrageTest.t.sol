// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {DeployArbitrage} from "script/DeployArbitrage.s.sol";
import {Arbitrage} from "src/Arbitrage.sol";
import {Test, console} from "forge-std/Test.sol";


contract ArbitrageTest is Test {
//    DeployArbitrage deployer;
    Arbitrage public arbitrage;


    address public immutable i_USDC_contract_address = 0x94a9D9AC8a22534E3FaCa9F4e7F2E2cf85d5E4C8;
    address public immutable i_deployed_arbitrage = 0xe3aFaaECEDda24a870700da1AF2b295eE496378F;
    address public  immutable i_startSwapAddress = "";
    address public  immutable i_endSwapAddress = "";
    address  public constant i_token0 = "";
    address  public constant i_token1 = "";


    uint256 public constant i_arbitrage_amount = 10000000;


    function setUp() external {
//    deployer = new Arbitrage();
        arbitrage = new Arbitrage(i_deployed_arbitrage);
    }

    function testArbitrage() public {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        arbitrage.makeFlashLoan(
            _startSwapAddress,
            _endSwapAddress,
            _token0,
            _token1,
            i_arbitrage_amount
        );
        vm.stopBroadcast();


        console.log("Arbitrage address: %s", address(arbitrage));
    }


}

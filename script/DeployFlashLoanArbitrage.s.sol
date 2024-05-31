// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/Test.sol";
import {FlashLoanArbitrage} from "../src/FlashLoanArbitrage.sol";

contract DeployArbitrage is Script {

    function run() external returns (FlashLoanArbitrage) {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        FlashLoanArbitrage arbitrage = new FlashLoanArbitrage();
        vm.stopBroadcast();
        console.log("Deployed Arbitrage at address: %s", address(arbitrage));
        return arbitrage;
    }
}

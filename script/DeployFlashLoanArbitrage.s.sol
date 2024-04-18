// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/Test.sol";
import {FlashLoanArbitrage} from "../src/FlashLoanArbitrage.sol";

contract DeployArbitrage is Script {
    address public constant I_POOL_ADDRESS_PROVIDER = 0x012bAC54348C0E635dCAc9D5FB99f06F24136C9A;

    function run() external returns (FlashLoanArbitrage) {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        FlashLoanArbitrage arbitrage = new FlashLoanArbitrage();
        vm.stopBroadcast();
        console.log("Deployed Arbitrage at address: %s", address(arbitrage));
        return arbitrage;
    }
}

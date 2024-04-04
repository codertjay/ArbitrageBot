// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {DeployFlashLoan} from "script/DeployFlashLoan.s.sol";
import {FlashLoan} from "src/FlashLoan.sol";
import {Test, console} from "forge-std/Test.sol";


contract FlashLoanTest is Test {
//    DeployFlashLoan deployer;
    FlashLoan public flashLoan;


    address public immutable i_USDC_contract_address = 0x94a9D9AC8a22534E3FaCa9F4e7F2E2cf85d5E4C8;
    address public immutable i_deployed_flashLoan = 0xe3aFaaECEDda24a870700da1AF2b295eE496378F;

    uint256 public constant i_flashLoan_amount = 10000000;

    function setUp() external {
//    deployer = new DeployFlashLoan();
        flashLoan = new FlashLoan(i_deployed_flashLoan);
    }

    function testFlashLoan() public {
        console.log("FlashLoan address: %s", address(flashLoan));
    }


    function testFlashLoanRevertsWhenYouDontPayEnough() public {
        flashLoan.requestFlashLoan(i_USDC_contract_address, i_flashLoan_amount);
    }

    function testGetUSDCBalanceISNotZero() public {
        uint256 flashLoanUsdBalance = flashLoan.getBalance(i_USDC_contract_address);
        console.log("USDC balance: %s", flashLoanUsdBalance);
    }

}

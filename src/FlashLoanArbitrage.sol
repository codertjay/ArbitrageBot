// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@balancer-labs/v2-interfaces/contracts/vault/IVault.sol";
import "@balancer-labs/v2-interfaces/contracts/vault/IFlashLoanRecipient.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

contract FlashLoanArbitrage is IFlashLoanRecipient {
    //////////////////
    // Errors   ///
    //////////////////
    error Arbitrage__OnlyOwner();

    ////////////////////////////////////
    // Modifiers   ///
    ////////////////////////////////////

    modifier onlyOwner() {
        if (msg.sender != owner) {
            revert Arbitrage__OnlyOwner();
        }
        _;
    }

    modifier onlyVault() {
        if (msg.sender != address(vault)) {
            revert Arbitrage__OnlyOwner();
        }
        _;
    }

    IVault private constant vault = IVault(0xBA12222222228d8Ba445958a75a0704d566BF2C8);

    address private constant  main_token = 0xc2132D05D31c914a87C6611C10748AEb04B58e8F;

    ////////////////////////////////////
    // State Declaration ///
    ////////////////////////////////////

    address public owner;


    struct DecentralizedExchange {
        string name;
        address routerV2;
        address factoryAddress;
    }

    mapping(bytes32 => DecentralizedExchange) private dexAddresses;

    ////////////////////////////////////
    // FUNCTIONS   ///
    ////////////////////////////////////

    constructor() {
        owner = msg.sender;

        // Initialize DEX addresses
        dexAddresses[keccak256(abi.encodePacked("M1"))] = DecentralizedExchange({
            name: "QUICKSWAP",
            routerV2: 0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff,
            factoryAddress: 0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32
        });

        dexAddresses[keccak256(abi.encodePacked("M2"))] = DecentralizedExchange({
            name: "SUSHISWAP",
            routerV2: 0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506,
            factoryAddress: 0xc35DADB65012eC5796536bD9864eD8773aBc74C4
        });

        dexAddresses[keccak256(abi.encodePacked("M3"))] = DecentralizedExchange({
            name: "DFYN",
            routerV2: 0xA102072A4C07F06EC3B4900FDC4C7B80b6c57429,
            factoryAddress: 0xE7Fb3e833eFE5F9c441105EB65Ef8b261266423B
        });
    }

    ////////////////////////////////////
    // EXTERNAL FUNCTIONS   ///
    ////////////////////////////////////

    /*
    @param _startOnUniswap : the router to start the arbitrage on
    @param _endSwapAddress : the router to end the arbitrage on
    @param _token0 : the address of the first token
    @param _token1 : the address of the second token
    */
    function makeFlashLoan(
        bytes32 _Key1,
        bytes32 _Key2,
        address _token1FirstPart,
        uint256 _flashAmount,
        uint256 gp
    ) external onlyOwner {
        bytes memory data = abi.encode(_Key1, _Key2, _token1FirstPart);

        // Token to flash loan, by default we are flash loaning 1 token.
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(main_token);

        // Flash loan amount.
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = _flashAmount;

//        vault.flashLoan(this, tokens, amounts, data);
//      Adjust gas price
        uint256 gasPrice = tx.gasprice + gp * 10 ** 9; // Current gas price + 5 Gwei
        bytes memory vault_data = abi.encodeWithSelector(vault.flashLoan.selector, this, tokens, amounts, data);
        (bool success,) = address(vault).call{gas: gasleft()}(vault_data);

        if (!success) {
            revert("Flash loan failed");
        }
    }

    /*
    @param tokens : the tokens to flash loan
    @param amounts : the amounts to flash loan
    @param feeAmounts : the fee amounts to flash loan
    @param userData : the data to pass to the flash loan
    */
    function receiveFlashLoan(
        IERC20[] memory tokens,
        uint256[] memory amounts,
        uint256[] memory feeAmounts,
        bytes memory userData
    ) external override onlyVault {
        uint256 flashAmount = amounts[0];

        (bytes32 startSwapKey, bytes32 endSwapKey,  bytes16 _token1, bytes4 _token2) =
                            abi.decode(userData, (bytes32, bytes32, bytes16, bytes4));

        address token1 = address(uint160(bytes20(abi.encodePacked(_token1, _token2))));


        address startSwapAddress = dexAddresses[startSwapKey].routerV2;
        address endSwapAddress = dexAddresses[endSwapKey].routerV2;

        // Make the Arbitrage Logic
        address[] memory path = new address[](2);

        path[0] = main_token;
        path[1] = token1;
        _swapTokens(path, flashAmount, 0, startSwapAddress);

        path[0] = token1;
        path[1] = main_token;

        uint256 amountOutMin = flashAmount + feeAmounts[0];
        _swapTokens(path, IERC20(token1).balanceOf(address(this)), amountOutMin, endSwapAddress);


        require(IERC20(main_token).balanceOf(address(this)) >= amountOutMin, "Arbitrage failed");


        bool transferSuccess = IERC20(main_token).transfer(address(vault), flashAmount + feeAmounts[0]);
        require(transferSuccess, "Transfer to vault failed");

    }

    /*
    @param path : the path to swap for and against which are two tokens
    @param _amountIn : the amount to swap in
    @param _amountOut : the amount to swap out
    @param _routerAddress : the router address to swap on which can be uniswap or any v2 contract address
    */
    function _swapTokens(address[] memory _path, uint256 _amountIn, uint256 _amountOut, address _routerAddress)
    internal onlyOwner
    {

        bool success = IERC20(_path[0]).approve(address(_routerAddress), _amountIn);


        IUniswapV2Router02(_routerAddress).swapExactTokensForTokens(
            _amountIn,
            _amountOut, // accept any amount of output tokens
            _path,
            address(this),
            (block.timestamp + 1200)
        );

    }


    struct ArbitrageResult {
        bool isProfitable;
        string direction;
        uint256 percentageProfit;
    }

    function getDecimals(address token) internal view returns (uint8) {
        return IERC20(token).decimals();
    }

    function checkProfitability(
        bytes32 _Key1,
        bytes32 _Key2,
        address _token1,
        uint256 _flashAmount,
        uint256 _threshold
    ) public view returns (ArbitrageResult memory) {
        ArbitrageResult memory result;

        /* Making the code obscure*/
        address token1 = _token1;


        address startSwapAddress = dexAddresses[_Key1].routerV2;
        address endSwapAddress = dexAddresses[_Key2].routerV2;

        address[] memory path = new address[](2);
        path[0] = main_token;
        path[1] = token1;

        uint256 _thresholdScaled = _threshold * 100;

        uint8 token0Decimals = getDecimals(main_token);
        uint8 token1Decimals = getDecimals(token1);


        uint256[] memory startSwapAmount = IUniswapV2Router02(startSwapAddress).getAmountsOut(_flashAmount, path);
        uint256[] memory endSwapAmount = IUniswapV2Router02(endSwapAddress).getAmountsOut(_flashAmount, path);


        uint256 startSwapPrice = (startSwapAmount[1] * 10 ** uint256(token0Decimals)) / startSwapAmount[0];
        uint256 endSwapPrice = (endSwapAmount[1] * 10 ** uint256(token0Decimals)) / endSwapAmount[0];


        uint256 TX_FEE = 3; // 0.3% fee, represented as 3 for easier calculations with integers

        if (startSwapPrice > endSwapPrice) {
            uint256 effStartSwapPrice = startSwapPrice * (1000 - TX_FEE) / 1000;
            uint256 effEndSwapPrice = endSwapPrice * (1000 - TX_FEE) / 1000;
            uint256 percentageDifference = ((effStartSwapPrice - effEndSwapPrice) * 10000) / effEndSwapPrice; // keeping two decimals


            if (percentageDifference >= _thresholdScaled) {
                return ArbitrageResult(
                    true,
                    "ATOB",
                    percentageDifference
                );

            }
        } else if (endSwapPrice > startSwapPrice) {
            uint256 effEndSwapPrice = endSwapPrice * (1000 - TX_FEE) / 1000;
            uint256 effStartSwapPrice = startSwapPrice * (1000 - TX_FEE) / 1000;
            uint256 percentageDifference = ((effEndSwapPrice - effStartSwapPrice) * 10000) / effStartSwapPrice;

            if (percentageDifference >= _thresholdScaled) {
                return ArbitrageResult(
                    true,
                    "BTOA",
                    percentageDifference
                );
            }
        }

        result.isProfitable = false;
        result.direction = "";
        result.percentageProfit = 0;
        return result;
    }


    function estimateGasCost() public view returns (uint256) {
        // Gas estimation logic (e.g., specific to the operations performed)
        uint256 gasUsed = 31000 + 100000; // Example estimation
        return gasUsed * tx.gasprice;
    }


    function withdrawEther(uint256 _amount) external onlyOwner {
        require(address(this).balance >= _amount, "Insufficient balance");
        payable(owner).transfer(_amount);
    }

    function withdrawERC20(address _token, uint256 _amount) external onlyOwner {
        uint256 contractBalance = IERC20(_token).balanceOf(address(this));
        require(contractBalance >= _amount, "Insufficient contract balance");
        IERC20(_token).transfer(owner, _amount);
    }

    function withdrawAllERC20(address _token) external onlyOwner {
        IERC20(_token).transfer(owner, IERC20(_token).balanceOf(address(this)));
    }


    receive() external payable {}
}

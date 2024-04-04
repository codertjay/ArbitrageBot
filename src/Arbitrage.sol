// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@balancer-labs/v2-interfaces/contracts/vault/IVault.sol";
import "@balancer-labs/v2-interfaces/contracts/vault/IFlashLoanRecipient.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

contract Arbitrage is IFlashLoanRecipient {
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

    ////////////////////////////////////
    // State Declaration ///
    ////////////////////////////////////

    address public owner;
    mapping(address => IUniswapV2Router02) public routers;
    address[] public routerAddresses;

    ////////////////////////////////////
    // FUNCTIONS   ///
    ////////////////////////////////////

    constructor() {
        owner = msg.sender;
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
        address _startSwapAddress,
        address _endSwapAddress,
        address _token0,
        address _token1,
        uint256 _flashAmount
    ) external onlyOwner {
        bytes memory data = abi.encode(_startSwapAddress, _endSwapAddress, _token0, _token1);

        // Token to flash loan, by default we are flash loaning 1 token.
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(_token0);

        // Flash loan amount.
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = _flashAmount;

        vault.flashLoan(this, tokens, amounts, data);
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

        (address startSwapAddress, address endSwapAddress, address token0, address token1) =
            abi.decode(userData, (address, address, address, address));

        // Make the Arbitrage Logic
        address[] memory path = new address[](2);

        path[0] = token0;
        path[1] = token1;
        _swapTokens(path, flashAmount, 0, startSwapAddress);

        path[0] = token1;
        path[1] = token0;
        _swapTokens(path, IERC20(token1).balanceOf(address(this)), flashAmount, endSwapAddress);

        // Repay the Flash Loan
        IERC20(token0).transfer(address(vault), flashAmount);
        IERC20(token0).transfer(owner, IERC20(token0).balanceOf(address(this)));
    }

    /*
    @param path : the path to swap for and against which are two tokens
    @param _amountIn : the amount to swap in
    @param _amountOut : the amount to swap out
    @param _routerAddress : the router address to swap on which can be uniswap or any v2 contract address
    */
    function _swapTokens(address[] memory _path, uint256 _amountIn, uint256 _amountOut, address _routerAddress)
        internal
    {
        require(IERC20(_path[0]).approve(_routerAddress, _amountIn), "Router approval failed.");

        IUniswapV2Router02(_routerAddress).swapExactTokensForTokens(
            _amountIn,
            _amountOut, // accept any amount of output tokens
            _path,
            address(this),
            (block.timestamp + 1200)
        );
    }

    function withdrawEther(uint256 _amount) external onlyOwner {
        require(address(this).balance >= _amount, "Insufficient balance");
        payable(owner).transfer(_amount);
    }

    receive() external payable {}
}

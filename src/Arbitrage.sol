// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IERC20.sol";


contract Arbitrage {
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
    function printMoney(
        address _startSwapAddress,
        address _endSwapAddress,
        address _token0,
        address _token1,
        uint256 _flashAmount
    ) external onlyOwner {
        // Make the Arbitrage Logic
        address[] memory path = new address[](2);

        path[0] = _token0;
        path[1] = _token1;
        _swapTokens(path, _flashAmount, 0, _startSwapAddress);

        path[0] = _token1;
        path[1] = _token0;
        _swapTokens(path, IERC20(_token1).balanceOf(address(this)), _flashAmount, _endSwapAddress);

        IERC20(_token0).transfer(owner, IERC20(_token0).balanceOf(address(this)));
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

    function withdrawERC20(address _token, uint256 _amount) external onlyOwner {
        IERC20(_token).transfer(owner, _amount);
    }

    receive() external payable {}
}

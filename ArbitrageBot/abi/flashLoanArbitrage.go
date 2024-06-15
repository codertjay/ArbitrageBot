// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrageABI

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FlashLoanArbitrageArbitrageResult is an auto generated low-level Go binding around an user-defined struct.
type FlashLoanArbitrageArbitrageResult struct {
	IsProfitable     bool
	Direction        string
	PercentageProfit *big.Int
}

// FlashLoanArbitrageMakeInput is an auto generated low-level Go binding around an user-defined struct.
type FlashLoanArbitrageMakeInput struct {
	Key1            [32]byte
	Key2            [32]byte
	Token1FirstPart common.Address
	FlashAmount     *big.Int
	Gp              *big.Int
}

// FlashLoanArbitrageMetaData contains all meta data concerning the FlashLoanArbitrage contract.
var FlashLoanArbitrageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"_swapTokens\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_routerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkProfitability\",\"inputs\":[{\"name\":\"_Key1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_Key2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFlashLoanArbitrage.ArbitrageResult\",\"components\":[{\"name\":\"isProfitable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"direction\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"percentageProfit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateGasCost\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"milking\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple[]\",\"internalType\":\"structFlashLoanArbitrage.MakeInput[]\",\"components\":[{\"name\":\"_Key1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_Key2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_token1FirstPart\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveFlashLoan\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"feeAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEther\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"Arbitrage__OnlyOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Arbitrage__OnlyVault\",\"inputs\":[]}]",
}

// FlashLoanArbitrageABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashLoanArbitrageMetaData.ABI instead.
var FlashLoanArbitrageABI = FlashLoanArbitrageMetaData.ABI

// FlashLoanArbitrage is an auto generated Go binding around an Ethereum contract.
type FlashLoanArbitrage struct {
	FlashLoanArbitrageCaller     // Read-only binding to the contract
	FlashLoanArbitrageTransactor // Write-only binding to the contract
	FlashLoanArbitrageFilterer   // Log filterer for contract events
}

// FlashLoanArbitrageCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLoanArbitrageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanArbitrageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLoanArbitrageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanArbitrageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLoanArbitrageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanArbitrageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLoanArbitrageSession struct {
	Contract     *FlashLoanArbitrage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FlashLoanArbitrageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLoanArbitrageCallerSession struct {
	Contract *FlashLoanArbitrageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// FlashLoanArbitrageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLoanArbitrageTransactorSession struct {
	Contract     *FlashLoanArbitrageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FlashLoanArbitrageRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLoanArbitrageRaw struct {
	Contract *FlashLoanArbitrage // Generic contract binding to access the raw methods on
}

// FlashLoanArbitrageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLoanArbitrageCallerRaw struct {
	Contract *FlashLoanArbitrageCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLoanArbitrageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLoanArbitrageTransactorRaw struct {
	Contract *FlashLoanArbitrageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLoanArbitrage creates a new instance of FlashLoanArbitrage, bound to a specific deployed contract.
func NewFlashLoanArbitrage(address common.Address, backend bind.ContractBackend) (*FlashLoanArbitrage, error) {
	contract, err := bindFlashLoanArbitrage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLoanArbitrage{FlashLoanArbitrageCaller: FlashLoanArbitrageCaller{contract: contract}, FlashLoanArbitrageTransactor: FlashLoanArbitrageTransactor{contract: contract}, FlashLoanArbitrageFilterer: FlashLoanArbitrageFilterer{contract: contract}}, nil
}

// NewFlashLoanArbitrageCaller creates a new read-only instance of FlashLoanArbitrage, bound to a specific deployed contract.
func NewFlashLoanArbitrageCaller(address common.Address, caller bind.ContractCaller) (*FlashLoanArbitrageCaller, error) {
	contract, err := bindFlashLoanArbitrage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanArbitrageCaller{contract: contract}, nil
}

// NewFlashLoanArbitrageTransactor creates a new write-only instance of FlashLoanArbitrage, bound to a specific deployed contract.
func NewFlashLoanArbitrageTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLoanArbitrageTransactor, error) {
	contract, err := bindFlashLoanArbitrage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanArbitrageTransactor{contract: contract}, nil
}

// NewFlashLoanArbitrageFilterer creates a new log filterer instance of FlashLoanArbitrage, bound to a specific deployed contract.
func NewFlashLoanArbitrageFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLoanArbitrageFilterer, error) {
	contract, err := bindFlashLoanArbitrage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLoanArbitrageFilterer{contract: contract}, nil
}

// bindFlashLoanArbitrage binds a generic wrapper to an already deployed contract.
func bindFlashLoanArbitrage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlashLoanArbitrageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanArbitrage *FlashLoanArbitrageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanArbitrage.Contract.FlashLoanArbitrageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanArbitrage *FlashLoanArbitrageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.FlashLoanArbitrageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanArbitrage *FlashLoanArbitrageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.FlashLoanArbitrageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanArbitrage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.contract.Transact(opts, method, params...)
}

// CheckProfitability is a free data retrieval call binding the contract method 0x347b29a2.
//
// Solidity: function checkProfitability(bytes32 _Key1, bytes32 _Key2, address _token1, uint256 _flashAmount, uint256 _threshold) view returns((bool,string,uint256))
func (_FlashLoanArbitrage *FlashLoanArbitrageCaller) CheckProfitability(opts *bind.CallOpts, _Key1 [32]byte, _Key2 [32]byte, _token1 common.Address, _flashAmount *big.Int, _threshold *big.Int) (FlashLoanArbitrageArbitrageResult, error) {
	var out []interface{}
	err := _FlashLoanArbitrage.contract.Call(opts, &out, "checkProfitability", _Key1, _Key2, _token1, _flashAmount, _threshold)

	if err != nil {
		return *new(FlashLoanArbitrageArbitrageResult), err
	}

	out0 := *abi.ConvertType(out[0], new(FlashLoanArbitrageArbitrageResult)).(*FlashLoanArbitrageArbitrageResult)

	return out0, err

}

// CheckProfitability is a free data retrieval call binding the contract method 0x347b29a2.
//
// Solidity: function checkProfitability(bytes32 _Key1, bytes32 _Key2, address _token1, uint256 _flashAmount, uint256 _threshold) view returns((bool,string,uint256))
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) CheckProfitability(_Key1 [32]byte, _Key2 [32]byte, _token1 common.Address, _flashAmount *big.Int, _threshold *big.Int) (FlashLoanArbitrageArbitrageResult, error) {
	return _FlashLoanArbitrage.Contract.CheckProfitability(&_FlashLoanArbitrage.CallOpts, _Key1, _Key2, _token1, _flashAmount, _threshold)
}

// CheckProfitability is a free data retrieval call binding the contract method 0x347b29a2.
//
// Solidity: function checkProfitability(bytes32 _Key1, bytes32 _Key2, address _token1, uint256 _flashAmount, uint256 _threshold) view returns((bool,string,uint256))
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerSession) CheckProfitability(_Key1 [32]byte, _Key2 [32]byte, _token1 common.Address, _flashAmount *big.Int, _threshold *big.Int) (FlashLoanArbitrageArbitrageResult, error) {
	return _FlashLoanArbitrage.Contract.CheckProfitability(&_FlashLoanArbitrage.CallOpts, _Key1, _Key2, _token1, _flashAmount, _threshold)
}

// EstimateGasCost is a free data retrieval call binding the contract method 0xfa4cc7c6.
//
// Solidity: function estimateGasCost() view returns(uint256)
func (_FlashLoanArbitrage *FlashLoanArbitrageCaller) EstimateGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlashLoanArbitrage.contract.Call(opts, &out, "estimateGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasCost is a free data retrieval call binding the contract method 0xfa4cc7c6.
//
// Solidity: function estimateGasCost() view returns(uint256)
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) EstimateGasCost() (*big.Int, error) {
	return _FlashLoanArbitrage.Contract.EstimateGasCost(&_FlashLoanArbitrage.CallOpts)
}

// EstimateGasCost is a free data retrieval call binding the contract method 0xfa4cc7c6.
//
// Solidity: function estimateGasCost() view returns(uint256)
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerSession) EstimateGasCost() (*big.Int, error) {
	return _FlashLoanArbitrage.Contract.EstimateGasCost(&_FlashLoanArbitrage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLoanArbitrage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) Owner() (common.Address, error) {
	return _FlashLoanArbitrage.Contract.Owner(&_FlashLoanArbitrage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerSession) Owner() (common.Address, error) {
	return _FlashLoanArbitrage.Contract.Owner(&_FlashLoanArbitrage.CallOpts)
}

// SwapTokens is a paid mutator transaction binding the contract method 0xf3a75160.
//
// Solidity: function _swapTokens(address[] _path, uint256 _amountIn, uint256 _amountOut, address _routerAddress) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) SwapTokens(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _amountOut *big.Int, _routerAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "_swapTokens", _path, _amountIn, _amountOut, _routerAddress)
}

// SwapTokens is a paid mutator transaction binding the contract method 0xf3a75160.
//
// Solidity: function _swapTokens(address[] _path, uint256 _amountIn, uint256 _amountOut, address _routerAddress) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) SwapTokens(_path []common.Address, _amountIn *big.Int, _amountOut *big.Int, _routerAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.SwapTokens(&_FlashLoanArbitrage.TransactOpts, _path, _amountIn, _amountOut, _routerAddress)
}

// SwapTokens is a paid mutator transaction binding the contract method 0xf3a75160.
//
// Solidity: function _swapTokens(address[] _path, uint256 _amountIn, uint256 _amountOut, address _routerAddress) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) SwapTokens(_path []common.Address, _amountIn *big.Int, _amountOut *big.Int, _routerAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.SwapTokens(&_FlashLoanArbitrage.TransactOpts, _path, _amountIn, _amountOut, _routerAddress)
}

// Milking is a paid mutator transaction binding the contract method 0x0f1080cc.
//
// Solidity: function milking((bytes32,bytes32,address,uint256,uint256)[] input) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) Milking(opts *bind.TransactOpts, input []FlashLoanArbitrageMakeInput) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "milking", input)
}

// Milking is a paid mutator transaction binding the contract method 0x0f1080cc.
//
// Solidity: function milking((bytes32,bytes32,address,uint256,uint256)[] input) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) Milking(input []FlashLoanArbitrageMakeInput) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.Milking(&_FlashLoanArbitrage.TransactOpts, input)
}

// Milking is a paid mutator transaction binding the contract method 0x0f1080cc.
//
// Solidity: function milking((bytes32,bytes32,address,uint256,uint256)[] input) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) Milking(input []FlashLoanArbitrageMakeInput) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.Milking(&_FlashLoanArbitrage.TransactOpts, input)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) ReceiveFlashLoan(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "receiveFlashLoan", tokens, amounts, feeAmounts, userData)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) ReceiveFlashLoan(tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.ReceiveFlashLoan(&_FlashLoanArbitrage.TransactOpts, tokens, amounts, feeAmounts, userData)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) ReceiveFlashLoan(tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.ReceiveFlashLoan(&_FlashLoanArbitrage.TransactOpts, tokens, amounts, feeAmounts, userData)
}

// WithdrawAllERC20 is a paid mutator transaction binding the contract method 0x857abbd4.
//
// Solidity: function withdrawAllERC20(address _token) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) WithdrawAllERC20(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "withdrawAllERC20", _token)
}

// WithdrawAllERC20 is a paid mutator transaction binding the contract method 0x857abbd4.
//
// Solidity: function withdrawAllERC20(address _token) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) WithdrawAllERC20(_token common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawAllERC20(&_FlashLoanArbitrage.TransactOpts, _token)
}

// WithdrawAllERC20 is a paid mutator transaction binding the contract method 0x857abbd4.
//
// Solidity: function withdrawAllERC20(address _token) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) WithdrawAllERC20(_token common.Address) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawAllERC20(&_FlashLoanArbitrage.TransactOpts, _token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "withdrawERC20", _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawERC20(&_FlashLoanArbitrage.TransactOpts, _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawERC20(&_FlashLoanArbitrage.TransactOpts, _token, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) WithdrawEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "withdrawEther", _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawEther(&_FlashLoanArbitrage.TransactOpts, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.WithdrawEther(&_FlashLoanArbitrage.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) Receive() (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.Receive(&_FlashLoanArbitrage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) Receive() (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.Receive(&_FlashLoanArbitrage.TransactOpts)
}

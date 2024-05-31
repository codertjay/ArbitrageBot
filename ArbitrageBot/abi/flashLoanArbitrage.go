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

// FlashLoanArbitrageMetaData contains all meta data concerning the FlashLoanArbitrage contract.
var FlashLoanArbitrageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"makeFlashLoan\",\"inputs\":[{\"name\":\"_startSwapAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_endSwapAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_flashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveFlashLoan\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"feeAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routerAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"routers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniswapV2Router02\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEther\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"Arbitrage__OnlyOwner\",\"inputs\":[]}]",
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

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCaller) RouterAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FlashLoanArbitrage.contract.Call(opts, &out, "routerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) RouterAddresses(arg0 *big.Int) (common.Address, error) {
	return _FlashLoanArbitrage.Contract.RouterAddresses(&_FlashLoanArbitrage.CallOpts, arg0)
}

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerSession) RouterAddresses(arg0 *big.Int) (common.Address, error) {
	return _FlashLoanArbitrage.Contract.RouterAddresses(&_FlashLoanArbitrage.CallOpts, arg0)
}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCaller) Routers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FlashLoanArbitrage.contract.Call(opts, &out, "routers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) Routers(arg0 common.Address) (common.Address, error) {
	return _FlashLoanArbitrage.Contract.Routers(&_FlashLoanArbitrage.CallOpts, arg0)
}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_FlashLoanArbitrage *FlashLoanArbitrageCallerSession) Routers(arg0 common.Address) (common.Address, error) {
	return _FlashLoanArbitrage.Contract.Routers(&_FlashLoanArbitrage.CallOpts, arg0)
}

// MakeFlashLoan is a paid mutator transaction binding the contract method 0xaeb03424.
//
// Solidity: function makeFlashLoan(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _flashAmount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactor) MakeFlashLoan(opts *bind.TransactOpts, _startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _flashAmount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.contract.Transact(opts, "makeFlashLoan", _startSwapAddress, _endSwapAddress, _token0, _token1, _flashAmount)
}

// MakeFlashLoan is a paid mutator transaction binding the contract method 0xaeb03424.
//
// Solidity: function makeFlashLoan(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _flashAmount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageSession) MakeFlashLoan(_startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _flashAmount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.MakeFlashLoan(&_FlashLoanArbitrage.TransactOpts, _startSwapAddress, _endSwapAddress, _token0, _token1, _flashAmount)
}

// MakeFlashLoan is a paid mutator transaction binding the contract method 0xaeb03424.
//
// Solidity: function makeFlashLoan(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _flashAmount) returns()
func (_FlashLoanArbitrage *FlashLoanArbitrageTransactorSession) MakeFlashLoan(_startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _flashAmount *big.Int) (*types.Transaction, error) {
	return _FlashLoanArbitrage.Contract.MakeFlashLoan(&_FlashLoanArbitrage.TransactOpts, _startSwapAddress, _endSwapAddress, _token0, _token1, _flashAmount)
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

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

// ArbitrageMetaData contains all meta data concerning the Arbitrage contract.
var ArbitrageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"printMoney\",\"inputs\":[{\"name\":\"_startSwapAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_endSwapAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_printAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"routerAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"routers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniswapV2Router02\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEther\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"Arbitrage__OnlyOwner\",\"inputs\":[]}]",
}

// ArbitrageABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbitrageMetaData.ABI instead.
var ArbitrageABI = ArbitrageMetaData.ABI

// Arbitrage is an auto generated Go binding around an Ethereum contract.
type Arbitrage struct {
	ArbitrageCaller     // Read-only binding to the contract
	ArbitrageTransactor // Write-only binding to the contract
	ArbitrageFilterer   // Log filterer for contract events
}

// ArbitrageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrageSession struct {
	Contract     *Arbitrage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitrageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrageCallerSession struct {
	Contract *ArbitrageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbitrageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrageTransactorSession struct {
	Contract     *ArbitrageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbitrageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrageRaw struct {
	Contract *Arbitrage // Generic contract binding to access the raw methods on
}

// ArbitrageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrageCallerRaw struct {
	Contract *ArbitrageCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrageTransactorRaw struct {
	Contract *ArbitrageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrage creates a new instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrage(address common.Address, backend bind.ContractBackend) (*Arbitrage, error) {
	contract, err := bindArbitrage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbitrage{ArbitrageCaller: ArbitrageCaller{contract: contract}, ArbitrageTransactor: ArbitrageTransactor{contract: contract}, ArbitrageFilterer: ArbitrageFilterer{contract: contract}}, nil
}

// NewArbitrageCaller creates a new read-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageCaller(address common.Address, caller bind.ContractCaller) (*ArbitrageCaller, error) {
	contract, err := bindArbitrage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageCaller{contract: contract}, nil
}

// NewArbitrageTransactor creates a new write-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrageTransactor, error) {
	contract, err := bindArbitrage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageTransactor{contract: contract}, nil
}

// NewArbitrageFilterer creates a new log filterer instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrageFilterer, error) {
	contract, err := bindArbitrage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrageFilterer{contract: contract}, nil
}

// bindArbitrage binds a generic wrapper to an already deployed contract.
func bindArbitrage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.ArbitrageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageSession) Owner() (common.Address, error) {
	return _Arbitrage.Contract.Owner(&_Arbitrage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Owner() (common.Address, error) {
	return _Arbitrage.Contract.Owner(&_Arbitrage.CallOpts)
}

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_Arbitrage *ArbitrageCaller) RouterAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "routerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_Arbitrage *ArbitrageSession) RouterAddresses(arg0 *big.Int) (common.Address, error) {
	return _Arbitrage.Contract.RouterAddresses(&_Arbitrage.CallOpts, arg0)
}

// RouterAddresses is a free data retrieval call binding the contract method 0xd2752680.
//
// Solidity: function routerAddresses(uint256 ) view returns(address)
func (_Arbitrage *ArbitrageCallerSession) RouterAddresses(arg0 *big.Int) (common.Address, error) {
	return _Arbitrage.Contract.RouterAddresses(&_Arbitrage.CallOpts, arg0)
}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_Arbitrage *ArbitrageCaller) Routers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "routers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_Arbitrage *ArbitrageSession) Routers(arg0 common.Address) (common.Address, error) {
	return _Arbitrage.Contract.Routers(&_Arbitrage.CallOpts, arg0)
}

// Routers is a free data retrieval call binding the contract method 0x80dd9a1f.
//
// Solidity: function routers(address ) view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Routers(arg0 common.Address) (common.Address, error) {
	return _Arbitrage.Contract.Routers(&_Arbitrage.CallOpts, arg0)
}

// PrintMoney is a paid mutator transaction binding the contract method 0xd874cc20.
//
// Solidity: function printMoney(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _printAmount) returns()
func (_Arbitrage *ArbitrageTransactor) PrintMoney(opts *bind.TransactOpts, _startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _printAmount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "printMoney", _startSwapAddress, _endSwapAddress, _token0, _token1, _printAmount)
}

// PrintMoney is a paid mutator transaction binding the contract method 0xd874cc20.
//
// Solidity: function printMoney(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _printAmount) returns()
func (_Arbitrage *ArbitrageSession) PrintMoney(_startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _printAmount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.PrintMoney(&_Arbitrage.TransactOpts, _startSwapAddress, _endSwapAddress, _token0, _token1, _printAmount)
}

// PrintMoney is a paid mutator transaction binding the contract method 0xd874cc20.
//
// Solidity: function printMoney(address _startSwapAddress, address _endSwapAddress, address _token0, address _token1, uint256 _printAmount) returns()
func (_Arbitrage *ArbitrageTransactorSession) PrintMoney(_startSwapAddress common.Address, _endSwapAddress common.Address, _token0 common.Address, _token1 common.Address, _printAmount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.PrintMoney(&_Arbitrage.TransactOpts, _startSwapAddress, _endSwapAddress, _token0, _token1, _printAmount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_Arbitrage *ArbitrageTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "withdrawERC20", _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_Arbitrage *ArbitrageSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.WithdrawERC20(&_Arbitrage.TransactOpts, _token, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount) returns()
func (_Arbitrage *ArbitrageTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.WithdrawERC20(&_Arbitrage.TransactOpts, _token, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_Arbitrage *ArbitrageTransactor) WithdrawEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "withdrawEther", _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_Arbitrage *ArbitrageSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.WithdrawEther(&_Arbitrage.TransactOpts, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_Arbitrage *ArbitrageTransactorSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.WithdrawEther(&_Arbitrage.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageSession) Receive() (*types.Transaction, error) {
	return _Arbitrage.Contract.Receive(&_Arbitrage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageTransactorSession) Receive() (*types.Transaction, error) {
	return _Arbitrage.Contract.Receive(&_Arbitrage.TransactOpts)
}

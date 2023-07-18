// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package root

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

// RootERC20PredicateMetaData contains all meta data concerning the RootERC20Predicate contract.
var RootERC20PredicateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAP_TOKEN_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childERC20Predicate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childTokenTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitHelper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStateSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newExitHelper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newChildERC20Predicate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newChildTokenTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nativeTokenRootAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"rootToken\",\"type\":\"address\"}],\"name\":\"mapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onL2StateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rootTokenToChildToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateSender\",\"outputs\":[{\"internalType\":\"contractIStateSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RootERC20PredicateABI is the input ABI used to generate the binding from.
// Deprecated: Use RootERC20PredicateMetaData.ABI instead.
var RootERC20PredicateABI = RootERC20PredicateMetaData.ABI

// RootERC20Predicate is an auto generated Go binding around an Ethereum contract.
type RootERC20Predicate struct {
	RootERC20PredicateCaller     // Read-only binding to the contract
	RootERC20PredicateTransactor // Write-only binding to the contract
	RootERC20PredicateFilterer   // Log filterer for contract events
}

// RootERC20PredicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootERC20PredicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootERC20PredicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootERC20PredicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootERC20PredicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootERC20PredicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootERC20PredicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootERC20PredicateSession struct {
	Contract     *RootERC20Predicate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RootERC20PredicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootERC20PredicateCallerSession struct {
	Contract *RootERC20PredicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RootERC20PredicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootERC20PredicateTransactorSession struct {
	Contract     *RootERC20PredicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RootERC20PredicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootERC20PredicateRaw struct {
	Contract *RootERC20Predicate // Generic contract binding to access the raw methods on
}

// RootERC20PredicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootERC20PredicateCallerRaw struct {
	Contract *RootERC20PredicateCaller // Generic read-only contract binding to access the raw methods on
}

// RootERC20PredicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootERC20PredicateTransactorRaw struct {
	Contract *RootERC20PredicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootERC20Predicate creates a new instance of RootERC20Predicate, bound to a specific deployed contract.
func NewRootERC20Predicate(address common.Address, backend bind.ContractBackend) (*RootERC20Predicate, error) {
	contract, err := bindRootERC20Predicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootERC20Predicate{RootERC20PredicateCaller: RootERC20PredicateCaller{contract: contract}, RootERC20PredicateTransactor: RootERC20PredicateTransactor{contract: contract}, RootERC20PredicateFilterer: RootERC20PredicateFilterer{contract: contract}}, nil
}

// NewRootERC20PredicateCaller creates a new read-only instance of RootERC20Predicate, bound to a specific deployed contract.
func NewRootERC20PredicateCaller(address common.Address, caller bind.ContractCaller) (*RootERC20PredicateCaller, error) {
	contract, err := bindRootERC20Predicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateCaller{contract: contract}, nil
}

// NewRootERC20PredicateTransactor creates a new write-only instance of RootERC20Predicate, bound to a specific deployed contract.
func NewRootERC20PredicateTransactor(address common.Address, transactor bind.ContractTransactor) (*RootERC20PredicateTransactor, error) {
	contract, err := bindRootERC20Predicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateTransactor{contract: contract}, nil
}

// NewRootERC20PredicateFilterer creates a new log filterer instance of RootERC20Predicate, bound to a specific deployed contract.
func NewRootERC20PredicateFilterer(address common.Address, filterer bind.ContractFilterer) (*RootERC20PredicateFilterer, error) {
	contract, err := bindRootERC20Predicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateFilterer{contract: contract}, nil
}

// bindRootERC20Predicate binds a generic wrapper to an already deployed contract.
func bindRootERC20Predicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RootERC20PredicateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootERC20Predicate *RootERC20PredicateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootERC20Predicate.Contract.RootERC20PredicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootERC20Predicate *RootERC20PredicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.RootERC20PredicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootERC20Predicate *RootERC20PredicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.RootERC20PredicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootERC20Predicate *RootERC20PredicateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootERC20Predicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootERC20Predicate *RootERC20PredicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootERC20Predicate *RootERC20PredicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCaller) DEPOSITSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "DEPOSIT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateSession) DEPOSITSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.DEPOSITSIG(&_RootERC20Predicate.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) DEPOSITSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.DEPOSITSIG(&_RootERC20Predicate.CallOpts)
}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCaller) MAPTOKENSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "MAP_TOKEN_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateSession) MAPTOKENSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.MAPTOKENSIG(&_RootERC20Predicate.CallOpts)
}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) MAPTOKENSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.MAPTOKENSIG(&_RootERC20Predicate.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCaller) WITHDRAWSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "WITHDRAW_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateSession) WITHDRAWSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.WITHDRAWSIG(&_RootERC20Predicate.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) WITHDRAWSIG() ([32]byte, error) {
	return _RootERC20Predicate.Contract.WITHDRAWSIG(&_RootERC20Predicate.CallOpts)
}

// ChildERC20Predicate is a free data retrieval call binding the contract method 0xd57184e4.
//
// Solidity: function childERC20Predicate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCaller) ChildERC20Predicate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "childERC20Predicate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildERC20Predicate is a free data retrieval call binding the contract method 0xd57184e4.
//
// Solidity: function childERC20Predicate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) ChildERC20Predicate() (common.Address, error) {
	return _RootERC20Predicate.Contract.ChildERC20Predicate(&_RootERC20Predicate.CallOpts)
}

// ChildERC20Predicate is a free data retrieval call binding the contract method 0xd57184e4.
//
// Solidity: function childERC20Predicate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) ChildERC20Predicate() (common.Address, error) {
	return _RootERC20Predicate.Contract.ChildERC20Predicate(&_RootERC20Predicate.CallOpts)
}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCaller) ChildTokenTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "childTokenTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) ChildTokenTemplate() (common.Address, error) {
	return _RootERC20Predicate.Contract.ChildTokenTemplate(&_RootERC20Predicate.CallOpts)
}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) ChildTokenTemplate() (common.Address, error) {
	return _RootERC20Predicate.Contract.ChildTokenTemplate(&_RootERC20Predicate.CallOpts)
}

// ExitHelper is a free data retrieval call binding the contract method 0x95c7041c.
//
// Solidity: function exitHelper() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCaller) ExitHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "exitHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExitHelper is a free data retrieval call binding the contract method 0x95c7041c.
//
// Solidity: function exitHelper() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) ExitHelper() (common.Address, error) {
	return _RootERC20Predicate.Contract.ExitHelper(&_RootERC20Predicate.CallOpts)
}

// ExitHelper is a free data retrieval call binding the contract method 0x95c7041c.
//
// Solidity: function exitHelper() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) ExitHelper() (common.Address, error) {
	return _RootERC20Predicate.Contract.ExitHelper(&_RootERC20Predicate.CallOpts)
}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCaller) RootTokenToChildToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "rootTokenToChildToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) RootTokenToChildToken(arg0 common.Address) (common.Address, error) {
	return _RootERC20Predicate.Contract.RootTokenToChildToken(&_RootERC20Predicate.CallOpts, arg0)
}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) RootTokenToChildToken(arg0 common.Address) (common.Address, error) {
	return _RootERC20Predicate.Contract.RootTokenToChildToken(&_RootERC20Predicate.CallOpts, arg0)
}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCaller) StateSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootERC20Predicate.contract.Call(opts, &out, "stateSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) StateSender() (common.Address, error) {
	return _RootERC20Predicate.Contract.StateSender(&_RootERC20Predicate.CallOpts)
}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() view returns(address)
func (_RootERC20Predicate *RootERC20PredicateCallerSession) StateSender() (common.Address, error) {
	return _RootERC20Predicate.Contract.StateSender(&_RootERC20Predicate.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootToken, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactor) Deposit(opts *bind.TransactOpts, rootToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.contract.Transact(opts, "deposit", rootToken, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootToken, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateSession) Deposit(rootToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.Deposit(&_RootERC20Predicate.TransactOpts, rootToken, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootToken, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactorSession) Deposit(rootToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.Deposit(&_RootERC20Predicate.TransactOpts, rootToken, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address rootToken, address receiver, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactor) DepositTo(opts *bind.TransactOpts, rootToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.contract.Transact(opts, "depositTo", rootToken, receiver, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address rootToken, address receiver, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateSession) DepositTo(rootToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.DepositTo(&_RootERC20Predicate.TransactOpts, rootToken, receiver, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address rootToken, address receiver, uint256 amount) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactorSession) DepositTo(rootToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.DepositTo(&_RootERC20Predicate.TransactOpts, rootToken, receiver, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newStateSender, address newExitHelper, address newChildERC20Predicate, address newChildTokenTemplate, address nativeTokenRootAddress) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactor) Initialize(opts *bind.TransactOpts, newStateSender common.Address, newExitHelper common.Address, newChildERC20Predicate common.Address, newChildTokenTemplate common.Address, nativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.contract.Transact(opts, "initialize", newStateSender, newExitHelper, newChildERC20Predicate, newChildTokenTemplate, nativeTokenRootAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newStateSender, address newExitHelper, address newChildERC20Predicate, address newChildTokenTemplate, address nativeTokenRootAddress) returns()
func (_RootERC20Predicate *RootERC20PredicateSession) Initialize(newStateSender common.Address, newExitHelper common.Address, newChildERC20Predicate common.Address, newChildTokenTemplate common.Address, nativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.Initialize(&_RootERC20Predicate.TransactOpts, newStateSender, newExitHelper, newChildERC20Predicate, newChildTokenTemplate, nativeTokenRootAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newStateSender, address newExitHelper, address newChildERC20Predicate, address newChildTokenTemplate, address nativeTokenRootAddress) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactorSession) Initialize(newStateSender common.Address, newExitHelper common.Address, newChildERC20Predicate common.Address, newChildTokenTemplate common.Address, nativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.Initialize(&_RootERC20Predicate.TransactOpts, newStateSender, newExitHelper, newChildERC20Predicate, newChildTokenTemplate, nativeTokenRootAddress)
}

// MapToken is a paid mutator transaction binding the contract method 0xf4a120f7.
//
// Solidity: function mapToken(address rootToken) returns(address)
func (_RootERC20Predicate *RootERC20PredicateTransactor) MapToken(opts *bind.TransactOpts, rootToken common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.contract.Transact(opts, "mapToken", rootToken)
}

// MapToken is a paid mutator transaction binding the contract method 0xf4a120f7.
//
// Solidity: function mapToken(address rootToken) returns(address)
func (_RootERC20Predicate *RootERC20PredicateSession) MapToken(rootToken common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.MapToken(&_RootERC20Predicate.TransactOpts, rootToken)
}

// MapToken is a paid mutator transaction binding the contract method 0xf4a120f7.
//
// Solidity: function mapToken(address rootToken) returns(address)
func (_RootERC20Predicate *RootERC20PredicateTransactorSession) MapToken(rootToken common.Address) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.MapToken(&_RootERC20Predicate.TransactOpts, rootToken)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactor) OnL2StateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _RootERC20Predicate.contract.Transact(opts, "onL2StateReceive", arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_RootERC20Predicate *RootERC20PredicateSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.OnL2StateReceive(&_RootERC20Predicate.TransactOpts, arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_RootERC20Predicate *RootERC20PredicateTransactorSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _RootERC20Predicate.Contract.OnL2StateReceive(&_RootERC20Predicate.TransactOpts, arg0, sender, data)
}

// RootERC20PredicateERC20DepositIterator is returned from FilterERC20Deposit and is used to iterate over the raw logs and unpacked data for ERC20Deposit events raised by the RootERC20Predicate contract.
type RootERC20PredicateERC20DepositIterator struct {
	Event *RootERC20PredicateERC20Deposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootERC20PredicateERC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootERC20PredicateERC20Deposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootERC20PredicateERC20Deposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootERC20PredicateERC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootERC20PredicateERC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootERC20PredicateERC20Deposit represents a ERC20Deposit event raised by the RootERC20Predicate contract.
type RootERC20PredicateERC20Deposit struct {
	RootToken  common.Address
	ChildToken common.Address
	Depositor  common.Address
	Receiver   common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposit is a free log retrieval operation binding the contract event 0x8be9001bb612c7123a1861dc0d9d94e683261f6cbbd7c7438b708975bc4908a3.
//
// Solidity: event ERC20Deposit(address indexed rootToken, address indexed childToken, address depositor, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) FilterERC20Deposit(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (*RootERC20PredicateERC20DepositIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.FilterLogs(opts, "ERC20Deposit", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateERC20DepositIterator{contract: _RootERC20Predicate.contract, event: "ERC20Deposit", logs: logs, sub: sub}, nil
}

// WatchERC20Deposit is a free log subscription operation binding the contract event 0x8be9001bb612c7123a1861dc0d9d94e683261f6cbbd7c7438b708975bc4908a3.
//
// Solidity: event ERC20Deposit(address indexed rootToken, address indexed childToken, address depositor, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) WatchERC20Deposit(opts *bind.WatchOpts, sink chan<- *RootERC20PredicateERC20Deposit, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.WatchLogs(opts, "ERC20Deposit", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootERC20PredicateERC20Deposit)
				if err := _RootERC20Predicate.contract.UnpackLog(event, "ERC20Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Deposit is a log parse operation binding the contract event 0x8be9001bb612c7123a1861dc0d9d94e683261f6cbbd7c7438b708975bc4908a3.
//
// Solidity: event ERC20Deposit(address indexed rootToken, address indexed childToken, address depositor, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) ParseERC20Deposit(log types.Log) (*RootERC20PredicateERC20Deposit, error) {
	event := new(RootERC20PredicateERC20Deposit)
	if err := _RootERC20Predicate.contract.UnpackLog(event, "ERC20Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootERC20PredicateERC20WithdrawIterator is returned from FilterERC20Withdraw and is used to iterate over the raw logs and unpacked data for ERC20Withdraw events raised by the RootERC20Predicate contract.
type RootERC20PredicateERC20WithdrawIterator struct {
	Event *RootERC20PredicateERC20Withdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootERC20PredicateERC20WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootERC20PredicateERC20Withdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootERC20PredicateERC20Withdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootERC20PredicateERC20WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootERC20PredicateERC20WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootERC20PredicateERC20Withdraw represents a ERC20Withdraw event raised by the RootERC20Predicate contract.
type RootERC20PredicateERC20Withdraw struct {
	RootToken  common.Address
	ChildToken common.Address
	Withdrawer common.Address
	Receiver   common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterERC20Withdraw is a free log retrieval operation binding the contract event 0x9c4f744b2e971d7058a9d8f43977e0e17bf7d57a48659f0e18541b7ee3d022e4.
//
// Solidity: event ERC20Withdraw(address indexed rootToken, address indexed childToken, address withdrawer, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) FilterERC20Withdraw(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (*RootERC20PredicateERC20WithdrawIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.FilterLogs(opts, "ERC20Withdraw", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateERC20WithdrawIterator{contract: _RootERC20Predicate.contract, event: "ERC20Withdraw", logs: logs, sub: sub}, nil
}

// WatchERC20Withdraw is a free log subscription operation binding the contract event 0x9c4f744b2e971d7058a9d8f43977e0e17bf7d57a48659f0e18541b7ee3d022e4.
//
// Solidity: event ERC20Withdraw(address indexed rootToken, address indexed childToken, address withdrawer, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) WatchERC20Withdraw(opts *bind.WatchOpts, sink chan<- *RootERC20PredicateERC20Withdraw, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.WatchLogs(opts, "ERC20Withdraw", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootERC20PredicateERC20Withdraw)
				if err := _RootERC20Predicate.contract.UnpackLog(event, "ERC20Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Withdraw is a log parse operation binding the contract event 0x9c4f744b2e971d7058a9d8f43977e0e17bf7d57a48659f0e18541b7ee3d022e4.
//
// Solidity: event ERC20Withdraw(address indexed rootToken, address indexed childToken, address withdrawer, address indexed receiver, uint256 amount)
func (_RootERC20Predicate *RootERC20PredicateFilterer) ParseERC20Withdraw(log types.Log) (*RootERC20PredicateERC20Withdraw, error) {
	event := new(RootERC20PredicateERC20Withdraw)
	if err := _RootERC20Predicate.contract.UnpackLog(event, "ERC20Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootERC20PredicateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RootERC20Predicate contract.
type RootERC20PredicateInitializedIterator struct {
	Event *RootERC20PredicateInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootERC20PredicateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootERC20PredicateInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootERC20PredicateInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootERC20PredicateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootERC20PredicateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootERC20PredicateInitialized represents a Initialized event raised by the RootERC20Predicate contract.
type RootERC20PredicateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RootERC20Predicate *RootERC20PredicateFilterer) FilterInitialized(opts *bind.FilterOpts) (*RootERC20PredicateInitializedIterator, error) {

	logs, sub, err := _RootERC20Predicate.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateInitializedIterator{contract: _RootERC20Predicate.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RootERC20Predicate *RootERC20PredicateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RootERC20PredicateInitialized) (event.Subscription, error) {

	logs, sub, err := _RootERC20Predicate.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootERC20PredicateInitialized)
				if err := _RootERC20Predicate.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RootERC20Predicate *RootERC20PredicateFilterer) ParseInitialized(log types.Log) (*RootERC20PredicateInitialized, error) {
	event := new(RootERC20PredicateInitialized)
	if err := _RootERC20Predicate.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootERC20PredicateTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the RootERC20Predicate contract.
type RootERC20PredicateTokenMappedIterator struct {
	Event *RootERC20PredicateTokenMapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootERC20PredicateTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootERC20PredicateTokenMapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootERC20PredicateTokenMapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootERC20PredicateTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootERC20PredicateTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootERC20PredicateTokenMapped represents a TokenMapped event raised by the RootERC20Predicate contract.
type RootERC20PredicateTokenMapped struct {
	RootToken  common.Address
	ChildToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_RootERC20Predicate *RootERC20PredicateFilterer) FilterTokenMapped(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address) (*RootERC20PredicateTokenMappedIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.FilterLogs(opts, "TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return &RootERC20PredicateTokenMappedIterator{contract: _RootERC20Predicate.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_RootERC20Predicate *RootERC20PredicateFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *RootERC20PredicateTokenMapped, rootToken []common.Address, childToken []common.Address) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _RootERC20Predicate.contract.WatchLogs(opts, "TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootERC20PredicateTokenMapped)
				if err := _RootERC20Predicate.contract.UnpackLog(event, "TokenMapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenMapped is a log parse operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_RootERC20Predicate *RootERC20PredicateFilterer) ParseTokenMapped(log types.Log) (*RootERC20PredicateTokenMapped, error) {
	event := new(RootERC20PredicateTokenMapped)
	if err := _RootERC20Predicate.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

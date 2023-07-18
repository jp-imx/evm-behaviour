// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package child

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

// ChildERC20PredicateMetaData contains all meta data concerning the ChildERC20Predicate contract.
var ChildERC20PredicateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"only\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2ERC20Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2ERC20Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"L2TokenMapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALLOWLIST_PRECOMPILE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCKLIST_PRECOMPILE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAP_TOKEN_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TRANSFER_PRECOMPILE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TRANSFER_PRECOMPILE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"READ_ADDRESSLIST_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_PKCHECK_PRECOMPILE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_PKCHECK_PRECOMPILE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childTokenTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newL2StateSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newStateReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRootERC20Predicate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newChildTokenTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newNativeTokenRootAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2StateSender\",\"outputs\":[{\"internalType\":\"contractIStateSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootERC20Predicate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rootTokenToChildToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChildERC20\",\"name\":\"childToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChildERC20\",\"name\":\"childToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChildERC20PredicateABI is the input ABI used to generate the binding from.
// Deprecated: Use ChildERC20PredicateMetaData.ABI instead.
var ChildERC20PredicateABI = ChildERC20PredicateMetaData.ABI

// ChildERC20Predicate is an auto generated Go binding around an Ethereum contract.
type ChildERC20Predicate struct {
	ChildERC20PredicateCaller     // Read-only binding to the contract
	ChildERC20PredicateTransactor // Write-only binding to the contract
	ChildERC20PredicateFilterer   // Log filterer for contract events
}

// ChildERC20PredicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildERC20PredicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20PredicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildERC20PredicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20PredicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildERC20PredicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20PredicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildERC20PredicateSession struct {
	Contract     *ChildERC20Predicate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChildERC20PredicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildERC20PredicateCallerSession struct {
	Contract *ChildERC20PredicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ChildERC20PredicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildERC20PredicateTransactorSession struct {
	Contract     *ChildERC20PredicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ChildERC20PredicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildERC20PredicateRaw struct {
	Contract *ChildERC20Predicate // Generic contract binding to access the raw methods on
}

// ChildERC20PredicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildERC20PredicateCallerRaw struct {
	Contract *ChildERC20PredicateCaller // Generic read-only contract binding to access the raw methods on
}

// ChildERC20PredicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildERC20PredicateTransactorRaw struct {
	Contract *ChildERC20PredicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildERC20Predicate creates a new instance of ChildERC20Predicate, bound to a specific deployed contract.
func NewChildERC20Predicate(address common.Address, backend bind.ContractBackend) (*ChildERC20Predicate, error) {
	contract, err := bindChildERC20Predicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildERC20Predicate{ChildERC20PredicateCaller: ChildERC20PredicateCaller{contract: contract}, ChildERC20PredicateTransactor: ChildERC20PredicateTransactor{contract: contract}, ChildERC20PredicateFilterer: ChildERC20PredicateFilterer{contract: contract}}, nil
}

// NewChildERC20PredicateCaller creates a new read-only instance of ChildERC20Predicate, bound to a specific deployed contract.
func NewChildERC20PredicateCaller(address common.Address, caller bind.ContractCaller) (*ChildERC20PredicateCaller, error) {
	contract, err := bindChildERC20Predicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateCaller{contract: contract}, nil
}

// NewChildERC20PredicateTransactor creates a new write-only instance of ChildERC20Predicate, bound to a specific deployed contract.
func NewChildERC20PredicateTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildERC20PredicateTransactor, error) {
	contract, err := bindChildERC20Predicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateTransactor{contract: contract}, nil
}

// NewChildERC20PredicateFilterer creates a new log filterer instance of ChildERC20Predicate, bound to a specific deployed contract.
func NewChildERC20PredicateFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildERC20PredicateFilterer, error) {
	contract, err := bindChildERC20Predicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateFilterer{contract: contract}, nil
}

// bindChildERC20Predicate binds a generic wrapper to an already deployed contract.
func bindChildERC20Predicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChildERC20PredicateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildERC20Predicate *ChildERC20PredicateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildERC20Predicate.Contract.ChildERC20PredicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildERC20Predicate *ChildERC20PredicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.ChildERC20PredicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildERC20Predicate *ChildERC20PredicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.ChildERC20PredicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildERC20Predicate *ChildERC20PredicateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildERC20Predicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildERC20Predicate *ChildERC20PredicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildERC20Predicate *ChildERC20PredicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.contract.Transact(opts, method, params...)
}

// ALLOWLISTPRECOMPILE is a free data retrieval call binding the contract method 0x55b01e4d.
//
// Solidity: function ALLOWLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) ALLOWLISTPRECOMPILE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "ALLOWLIST_PRECOMPILE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOWLISTPRECOMPILE is a free data retrieval call binding the contract method 0x55b01e4d.
//
// Solidity: function ALLOWLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) ALLOWLISTPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.ALLOWLISTPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// ALLOWLISTPRECOMPILE is a free data retrieval call binding the contract method 0x55b01e4d.
//
// Solidity: function ALLOWLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) ALLOWLISTPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.ALLOWLISTPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// BLOCKLISTPRECOMPILE is a free data retrieval call binding the contract method 0x07b3e252.
//
// Solidity: function BLOCKLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) BLOCKLISTPRECOMPILE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "BLOCKLIST_PRECOMPILE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BLOCKLISTPRECOMPILE is a free data retrieval call binding the contract method 0x07b3e252.
//
// Solidity: function BLOCKLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) BLOCKLISTPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.BLOCKLISTPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// BLOCKLISTPRECOMPILE is a free data retrieval call binding the contract method 0x07b3e252.
//
// Solidity: function BLOCKLIST_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) BLOCKLISTPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.BLOCKLISTPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) DEPOSITSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "DEPOSIT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateSession) DEPOSITSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.DEPOSITSIG(&_ChildERC20Predicate.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) DEPOSITSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.DEPOSITSIG(&_ChildERC20Predicate.CallOpts)
}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) MAPTOKENSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "MAP_TOKEN_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateSession) MAPTOKENSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.MAPTOKENSIG(&_ChildERC20Predicate.CallOpts)
}

// MAPTOKENSIG is a free data retrieval call binding the contract method 0xf6451255.
//
// Solidity: function MAP_TOKEN_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) MAPTOKENSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.MAPTOKENSIG(&_ChildERC20Predicate.CallOpts)
}

// NATIVETOKENCONTRACT is a free data retrieval call binding the contract method 0x3b878c22.
//
// Solidity: function NATIVE_TOKEN_CONTRACT() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) NATIVETOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "NATIVE_TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKENCONTRACT is a free data retrieval call binding the contract method 0x3b878c22.
//
// Solidity: function NATIVE_TOKEN_CONTRACT() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) NATIVETOKENCONTRACT() (common.Address, error) {
	return _ChildERC20Predicate.Contract.NATIVETOKENCONTRACT(&_ChildERC20Predicate.CallOpts)
}

// NATIVETOKENCONTRACT is a free data retrieval call binding the contract method 0x3b878c22.
//
// Solidity: function NATIVE_TOKEN_CONTRACT() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) NATIVETOKENCONTRACT() (common.Address, error) {
	return _ChildERC20Predicate.Contract.NATIVETOKENCONTRACT(&_ChildERC20Predicate.CallOpts)
}

// NATIVETRANSFERPRECOMPILE is a free data retrieval call binding the contract method 0x284017f5.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) NATIVETRANSFERPRECOMPILE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "NATIVE_TRANSFER_PRECOMPILE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETRANSFERPRECOMPILE is a free data retrieval call binding the contract method 0x284017f5.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) NATIVETRANSFERPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.NATIVETRANSFERPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// NATIVETRANSFERPRECOMPILE is a free data retrieval call binding the contract method 0x284017f5.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) NATIVETRANSFERPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.NATIVETRANSFERPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// NATIVETRANSFERPRECOMPILEGAS is a free data retrieval call binding the contract method 0x947287cf.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) NATIVETRANSFERPRECOMPILEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "NATIVE_TRANSFER_PRECOMPILE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NATIVETRANSFERPRECOMPILEGAS is a free data retrieval call binding the contract method 0x947287cf.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateSession) NATIVETRANSFERPRECOMPILEGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.NATIVETRANSFERPRECOMPILEGAS(&_ChildERC20Predicate.CallOpts)
}

// NATIVETRANSFERPRECOMPILEGAS is a free data retrieval call binding the contract method 0x947287cf.
//
// Solidity: function NATIVE_TRANSFER_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) NATIVETRANSFERPRECOMPILEGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.NATIVETRANSFERPRECOMPILEGAS(&_ChildERC20Predicate.CallOpts)
}

// READADDRESSLISTGAS is a free data retrieval call binding the contract method 0x5ea5df79.
//
// Solidity: function READ_ADDRESSLIST_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) READADDRESSLISTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "READ_ADDRESSLIST_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// READADDRESSLISTGAS is a free data retrieval call binding the contract method 0x5ea5df79.
//
// Solidity: function READ_ADDRESSLIST_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateSession) READADDRESSLISTGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.READADDRESSLISTGAS(&_ChildERC20Predicate.CallOpts)
}

// READADDRESSLISTGAS is a free data retrieval call binding the contract method 0x5ea5df79.
//
// Solidity: function READ_ADDRESSLIST_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) READADDRESSLISTGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.READADDRESSLISTGAS(&_ChildERC20Predicate.CallOpts)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) SYSTEM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "SYSTEM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) SYSTEM() (common.Address, error) {
	return _ChildERC20Predicate.Contract.SYSTEM(&_ChildERC20Predicate.CallOpts)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) SYSTEM() (common.Address, error) {
	return _ChildERC20Predicate.Contract.SYSTEM(&_ChildERC20Predicate.CallOpts)
}

// VALIDATORPKCHECKPRECOMPILE is a free data retrieval call binding the contract method 0xe0563ab1.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) VALIDATORPKCHECKPRECOMPILE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "VALIDATOR_PKCHECK_PRECOMPILE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORPKCHECKPRECOMPILE is a free data retrieval call binding the contract method 0xe0563ab1.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) VALIDATORPKCHECKPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.VALIDATORPKCHECKPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// VALIDATORPKCHECKPRECOMPILE is a free data retrieval call binding the contract method 0xe0563ab1.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) VALIDATORPKCHECKPRECOMPILE() (common.Address, error) {
	return _ChildERC20Predicate.Contract.VALIDATORPKCHECKPRECOMPILE(&_ChildERC20Predicate.CallOpts)
}

// VALIDATORPKCHECKPRECOMPILEGAS is a free data retrieval call binding the contract method 0x97e5230d.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) VALIDATORPKCHECKPRECOMPILEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "VALIDATOR_PKCHECK_PRECOMPILE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORPKCHECKPRECOMPILEGAS is a free data retrieval call binding the contract method 0x97e5230d.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateSession) VALIDATORPKCHECKPRECOMPILEGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.VALIDATORPKCHECKPRECOMPILEGAS(&_ChildERC20Predicate.CallOpts)
}

// VALIDATORPKCHECKPRECOMPILEGAS is a free data retrieval call binding the contract method 0x97e5230d.
//
// Solidity: function VALIDATOR_PKCHECK_PRECOMPILE_GAS() view returns(uint256)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) VALIDATORPKCHECKPRECOMPILEGAS() (*big.Int, error) {
	return _ChildERC20Predicate.Contract.VALIDATORPKCHECKPRECOMPILEGAS(&_ChildERC20Predicate.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) WITHDRAWSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "WITHDRAW_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateSession) WITHDRAWSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.WITHDRAWSIG(&_ChildERC20Predicate.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) WITHDRAWSIG() ([32]byte, error) {
	return _ChildERC20Predicate.Contract.WITHDRAWSIG(&_ChildERC20Predicate.CallOpts)
}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) ChildTokenTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "childTokenTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) ChildTokenTemplate() (common.Address, error) {
	return _ChildERC20Predicate.Contract.ChildTokenTemplate(&_ChildERC20Predicate.CallOpts)
}

// ChildTokenTemplate is a free data retrieval call binding the contract method 0xb68ad1e4.
//
// Solidity: function childTokenTemplate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) ChildTokenTemplate() (common.Address, error) {
	return _ChildERC20Predicate.Contract.ChildTokenTemplate(&_ChildERC20Predicate.CallOpts)
}

// L2StateSender is a free data retrieval call binding the contract method 0x1bc114ba.
//
// Solidity: function l2StateSender() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) L2StateSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "l2StateSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2StateSender is a free data retrieval call binding the contract method 0x1bc114ba.
//
// Solidity: function l2StateSender() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) L2StateSender() (common.Address, error) {
	return _ChildERC20Predicate.Contract.L2StateSender(&_ChildERC20Predicate.CallOpts)
}

// L2StateSender is a free data retrieval call binding the contract method 0x1bc114ba.
//
// Solidity: function l2StateSender() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) L2StateSender() (common.Address, error) {
	return _ChildERC20Predicate.Contract.L2StateSender(&_ChildERC20Predicate.CallOpts)
}

// RootERC20Predicate is a free data retrieval call binding the contract method 0x71cf93b7.
//
// Solidity: function rootERC20Predicate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) RootERC20Predicate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "rootERC20Predicate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootERC20Predicate is a free data retrieval call binding the contract method 0x71cf93b7.
//
// Solidity: function rootERC20Predicate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) RootERC20Predicate() (common.Address, error) {
	return _ChildERC20Predicate.Contract.RootERC20Predicate(&_ChildERC20Predicate.CallOpts)
}

// RootERC20Predicate is a free data retrieval call binding the contract method 0x71cf93b7.
//
// Solidity: function rootERC20Predicate() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) RootERC20Predicate() (common.Address, error) {
	return _ChildERC20Predicate.Contract.RootERC20Predicate(&_ChildERC20Predicate.CallOpts)
}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) RootTokenToChildToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "rootTokenToChildToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) RootTokenToChildToken(arg0 common.Address) (common.Address, error) {
	return _ChildERC20Predicate.Contract.RootTokenToChildToken(&_ChildERC20Predicate.CallOpts, arg0)
}

// RootTokenToChildToken is a free data retrieval call binding the contract method 0x7efab4f5.
//
// Solidity: function rootTokenToChildToken(address ) view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) RootTokenToChildToken(arg0 common.Address) (common.Address, error) {
	return _ChildERC20Predicate.Contract.RootTokenToChildToken(&_ChildERC20Predicate.CallOpts, arg0)
}

// StateReceiver is a free data retrieval call binding the contract method 0x05dc2e8f.
//
// Solidity: function stateReceiver() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCaller) StateReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20Predicate.contract.Call(opts, &out, "stateReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateReceiver is a free data retrieval call binding the contract method 0x05dc2e8f.
//
// Solidity: function stateReceiver() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateSession) StateReceiver() (common.Address, error) {
	return _ChildERC20Predicate.Contract.StateReceiver(&_ChildERC20Predicate.CallOpts)
}

// StateReceiver is a free data retrieval call binding the contract method 0x05dc2e8f.
//
// Solidity: function stateReceiver() view returns(address)
func (_ChildERC20Predicate *ChildERC20PredicateCallerSession) StateReceiver() (common.Address, error) {
	return _ChildERC20Predicate.Contract.StateReceiver(&_ChildERC20Predicate.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newL2StateSender, address newStateReceiver, address newRootERC20Predicate, address newChildTokenTemplate, address newNativeTokenRootAddress) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactor) Initialize(opts *bind.TransactOpts, newL2StateSender common.Address, newStateReceiver common.Address, newRootERC20Predicate common.Address, newChildTokenTemplate common.Address, newNativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _ChildERC20Predicate.contract.Transact(opts, "initialize", newL2StateSender, newStateReceiver, newRootERC20Predicate, newChildTokenTemplate, newNativeTokenRootAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newL2StateSender, address newStateReceiver, address newRootERC20Predicate, address newChildTokenTemplate, address newNativeTokenRootAddress) returns()
func (_ChildERC20Predicate *ChildERC20PredicateSession) Initialize(newL2StateSender common.Address, newStateReceiver common.Address, newRootERC20Predicate common.Address, newChildTokenTemplate common.Address, newNativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.Initialize(&_ChildERC20Predicate.TransactOpts, newL2StateSender, newStateReceiver, newRootERC20Predicate, newChildTokenTemplate, newNativeTokenRootAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address newL2StateSender, address newStateReceiver, address newRootERC20Predicate, address newChildTokenTemplate, address newNativeTokenRootAddress) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactorSession) Initialize(newL2StateSender common.Address, newStateReceiver common.Address, newRootERC20Predicate common.Address, newChildTokenTemplate common.Address, newNativeTokenRootAddress common.Address) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.Initialize(&_ChildERC20Predicate.TransactOpts, newL2StateSender, newStateReceiver, newRootERC20Predicate, newChildTokenTemplate, newNativeTokenRootAddress)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactor) OnStateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _ChildERC20Predicate.contract.Transact(opts, "onStateReceive", arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_ChildERC20Predicate *ChildERC20PredicateSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.OnStateReceive(&_ChildERC20Predicate.TransactOpts, arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactorSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.OnStateReceive(&_ChildERC20Predicate.TransactOpts, arg0, sender, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address childToken, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactor) Withdraw(opts *bind.TransactOpts, childToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.contract.Transact(opts, "withdraw", childToken, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address childToken, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateSession) Withdraw(childToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.Withdraw(&_ChildERC20Predicate.TransactOpts, childToken, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address childToken, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactorSession) Withdraw(childToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.Withdraw(&_ChildERC20Predicate.TransactOpts, childToken, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address childToken, address receiver, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactor) WithdrawTo(opts *bind.TransactOpts, childToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.contract.Transact(opts, "withdrawTo", childToken, receiver, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address childToken, address receiver, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateSession) WithdrawTo(childToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.WithdrawTo(&_ChildERC20Predicate.TransactOpts, childToken, receiver, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address childToken, address receiver, uint256 amount) returns()
func (_ChildERC20Predicate *ChildERC20PredicateTransactorSession) WithdrawTo(childToken common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20Predicate.Contract.WithdrawTo(&_ChildERC20Predicate.TransactOpts, childToken, receiver, amount)
}

// ChildERC20PredicateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChildERC20Predicate contract.
type ChildERC20PredicateInitializedIterator struct {
	Event *ChildERC20PredicateInitialized // Event containing the contract specifics and raw log

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
func (it *ChildERC20PredicateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20PredicateInitialized)
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
		it.Event = new(ChildERC20PredicateInitialized)
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
func (it *ChildERC20PredicateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20PredicateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20PredicateInitialized represents a Initialized event raised by the ChildERC20Predicate contract.
type ChildERC20PredicateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChildERC20PredicateInitializedIterator, error) {

	logs, sub, err := _ChildERC20Predicate.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateInitializedIterator{contract: _ChildERC20Predicate.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChildERC20PredicateInitialized) (event.Subscription, error) {

	logs, sub, err := _ChildERC20Predicate.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20PredicateInitialized)
				if err := _ChildERC20Predicate.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) ParseInitialized(log types.Log) (*ChildERC20PredicateInitialized, error) {
	event := new(ChildERC20PredicateInitialized)
	if err := _ChildERC20Predicate.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20PredicateL2ERC20DepositIterator is returned from FilterL2ERC20Deposit and is used to iterate over the raw logs and unpacked data for L2ERC20Deposit events raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2ERC20DepositIterator struct {
	Event *ChildERC20PredicateL2ERC20Deposit // Event containing the contract specifics and raw log

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
func (it *ChildERC20PredicateL2ERC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20PredicateL2ERC20Deposit)
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
		it.Event = new(ChildERC20PredicateL2ERC20Deposit)
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
func (it *ChildERC20PredicateL2ERC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20PredicateL2ERC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20PredicateL2ERC20Deposit represents a L2ERC20Deposit event raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2ERC20Deposit struct {
	RootToken  common.Address
	ChildToken common.Address
	Sender     common.Address
	Receiver   common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2ERC20Deposit is a free log retrieval operation binding the contract event 0xdf34f3a3ed8bedc14a4b284ebaee5374d55b64bac6a84c270dabe8fd6b4cdafd.
//
// Solidity: event L2ERC20Deposit(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) FilterL2ERC20Deposit(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (*ChildERC20PredicateL2ERC20DepositIterator, error) {

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

	logs, sub, err := _ChildERC20Predicate.contract.FilterLogs(opts, "L2ERC20Deposit", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateL2ERC20DepositIterator{contract: _ChildERC20Predicate.contract, event: "L2ERC20Deposit", logs: logs, sub: sub}, nil
}

// WatchL2ERC20Deposit is a free log subscription operation binding the contract event 0xdf34f3a3ed8bedc14a4b284ebaee5374d55b64bac6a84c270dabe8fd6b4cdafd.
//
// Solidity: event L2ERC20Deposit(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) WatchL2ERC20Deposit(opts *bind.WatchOpts, sink chan<- *ChildERC20PredicateL2ERC20Deposit, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChildERC20Predicate.contract.WatchLogs(opts, "L2ERC20Deposit", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20PredicateL2ERC20Deposit)
				if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2ERC20Deposit", log); err != nil {
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

// ParseL2ERC20Deposit is a log parse operation binding the contract event 0xdf34f3a3ed8bedc14a4b284ebaee5374d55b64bac6a84c270dabe8fd6b4cdafd.
//
// Solidity: event L2ERC20Deposit(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) ParseL2ERC20Deposit(log types.Log) (*ChildERC20PredicateL2ERC20Deposit, error) {
	event := new(ChildERC20PredicateL2ERC20Deposit)
	if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2ERC20Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20PredicateL2ERC20WithdrawIterator is returned from FilterL2ERC20Withdraw and is used to iterate over the raw logs and unpacked data for L2ERC20Withdraw events raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2ERC20WithdrawIterator struct {
	Event *ChildERC20PredicateL2ERC20Withdraw // Event containing the contract specifics and raw log

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
func (it *ChildERC20PredicateL2ERC20WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20PredicateL2ERC20Withdraw)
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
		it.Event = new(ChildERC20PredicateL2ERC20Withdraw)
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
func (it *ChildERC20PredicateL2ERC20WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20PredicateL2ERC20WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20PredicateL2ERC20Withdraw represents a L2ERC20Withdraw event raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2ERC20Withdraw struct {
	RootToken  common.Address
	ChildToken common.Address
	Sender     common.Address
	Receiver   common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2ERC20Withdraw is a free log retrieval operation binding the contract event 0xa0923f060a16fc784558d43de424ffde7b01643de5e5d335851b9df94c76bb27.
//
// Solidity: event L2ERC20Withdraw(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) FilterL2ERC20Withdraw(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (*ChildERC20PredicateL2ERC20WithdrawIterator, error) {

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

	logs, sub, err := _ChildERC20Predicate.contract.FilterLogs(opts, "L2ERC20Withdraw", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateL2ERC20WithdrawIterator{contract: _ChildERC20Predicate.contract, event: "L2ERC20Withdraw", logs: logs, sub: sub}, nil
}

// WatchL2ERC20Withdraw is a free log subscription operation binding the contract event 0xa0923f060a16fc784558d43de424ffde7b01643de5e5d335851b9df94c76bb27.
//
// Solidity: event L2ERC20Withdraw(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) WatchL2ERC20Withdraw(opts *bind.WatchOpts, sink chan<- *ChildERC20PredicateL2ERC20Withdraw, rootToken []common.Address, childToken []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChildERC20Predicate.contract.WatchLogs(opts, "L2ERC20Withdraw", rootTokenRule, childTokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20PredicateL2ERC20Withdraw)
				if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2ERC20Withdraw", log); err != nil {
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

// ParseL2ERC20Withdraw is a log parse operation binding the contract event 0xa0923f060a16fc784558d43de424ffde7b01643de5e5d335851b9df94c76bb27.
//
// Solidity: event L2ERC20Withdraw(address indexed rootToken, address indexed childToken, address sender, address indexed receiver, uint256 amount)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) ParseL2ERC20Withdraw(log types.Log) (*ChildERC20PredicateL2ERC20Withdraw, error) {
	event := new(ChildERC20PredicateL2ERC20Withdraw)
	if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2ERC20Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20PredicateL2TokenMappedIterator is returned from FilterL2TokenMapped and is used to iterate over the raw logs and unpacked data for L2TokenMapped events raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2TokenMappedIterator struct {
	Event *ChildERC20PredicateL2TokenMapped // Event containing the contract specifics and raw log

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
func (it *ChildERC20PredicateL2TokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20PredicateL2TokenMapped)
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
		it.Event = new(ChildERC20PredicateL2TokenMapped)
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
func (it *ChildERC20PredicateL2TokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20PredicateL2TokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20PredicateL2TokenMapped represents a L2TokenMapped event raised by the ChildERC20Predicate contract.
type ChildERC20PredicateL2TokenMapped struct {
	RootToken  common.Address
	ChildToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2TokenMapped is a free log retrieval operation binding the contract event 0x46bd56f98e1b14fd35691959270a6e1edf7cb8fcd489e0f9dda89e46c0d1ff0d.
//
// Solidity: event L2TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) FilterL2TokenMapped(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address) (*ChildERC20PredicateL2TokenMappedIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _ChildERC20Predicate.contract.FilterLogs(opts, "L2TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return &ChildERC20PredicateL2TokenMappedIterator{contract: _ChildERC20Predicate.contract, event: "L2TokenMapped", logs: logs, sub: sub}, nil
}

// WatchL2TokenMapped is a free log subscription operation binding the contract event 0x46bd56f98e1b14fd35691959270a6e1edf7cb8fcd489e0f9dda89e46c0d1ff0d.
//
// Solidity: event L2TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) WatchL2TokenMapped(opts *bind.WatchOpts, sink chan<- *ChildERC20PredicateL2TokenMapped, rootToken []common.Address, childToken []common.Address) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _ChildERC20Predicate.contract.WatchLogs(opts, "L2TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20PredicateL2TokenMapped)
				if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2TokenMapped", log); err != nil {
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

// ParseL2TokenMapped is a log parse operation binding the contract event 0x46bd56f98e1b14fd35691959270a6e1edf7cb8fcd489e0f9dda89e46c0d1ff0d.
//
// Solidity: event L2TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildERC20Predicate *ChildERC20PredicateFilterer) ParseL2TokenMapped(log types.Log) (*ChildERC20PredicateL2TokenMapped, error) {
	event := new(ChildERC20PredicateL2TokenMapped)
	if err := _ChildERC20Predicate.contract.UnpackLog(event, "L2TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

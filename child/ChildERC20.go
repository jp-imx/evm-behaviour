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

// ChildERC20MetaData contains all meta data concerning the ChildERC20 contract.
var ChildERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"invalidateNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"predicate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChildERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ChildERC20MetaData.ABI instead.
var ChildERC20ABI = ChildERC20MetaData.ABI

// ChildERC20 is an auto generated Go binding around an Ethereum contract.
type ChildERC20 struct {
	ChildERC20Caller     // Read-only binding to the contract
	ChildERC20Transactor // Write-only binding to the contract
	ChildERC20Filterer   // Log filterer for contract events
}

// ChildERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ChildERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildERC20Session struct {
	Contract     *ChildERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildERC20CallerSession struct {
	Contract *ChildERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ChildERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildERC20TransactorSession struct {
	Contract     *ChildERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ChildERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ChildERC20Raw struct {
	Contract *ChildERC20 // Generic contract binding to access the raw methods on
}

// ChildERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildERC20CallerRaw struct {
	Contract *ChildERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ChildERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildERC20TransactorRaw struct {
	Contract *ChildERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewChildERC20 creates a new instance of ChildERC20, bound to a specific deployed contract.
func NewChildERC20(address common.Address, backend bind.ContractBackend) (*ChildERC20, error) {
	contract, err := bindChildERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildERC20{ChildERC20Caller: ChildERC20Caller{contract: contract}, ChildERC20Transactor: ChildERC20Transactor{contract: contract}, ChildERC20Filterer: ChildERC20Filterer{contract: contract}}, nil
}

// NewChildERC20Caller creates a new read-only instance of ChildERC20, bound to a specific deployed contract.
func NewChildERC20Caller(address common.Address, caller bind.ContractCaller) (*ChildERC20Caller, error) {
	contract, err := bindChildERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildERC20Caller{contract: contract}, nil
}

// NewChildERC20Transactor creates a new write-only instance of ChildERC20, bound to a specific deployed contract.
func NewChildERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ChildERC20Transactor, error) {
	contract, err := bindChildERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildERC20Transactor{contract: contract}, nil
}

// NewChildERC20Filterer creates a new log filterer instance of ChildERC20, bound to a specific deployed contract.
func NewChildERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ChildERC20Filterer, error) {
	contract, err := bindChildERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildERC20Filterer{contract: contract}, nil
}

// bindChildERC20 binds a generic wrapper to an already deployed contract.
func bindChildERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChildERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildERC20 *ChildERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildERC20.Contract.ChildERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildERC20 *ChildERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildERC20.Contract.ChildERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildERC20 *ChildERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildERC20.Contract.ChildERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildERC20 *ChildERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildERC20 *ChildERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildERC20 *ChildERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildERC20 *ChildERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildERC20 *ChildERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.Allowance(&_ChildERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildERC20 *ChildERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.Allowance(&_ChildERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildERC20 *ChildERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildERC20 *ChildERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.BalanceOf(&_ChildERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildERC20 *ChildERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.BalanceOf(&_ChildERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildERC20 *ChildERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildERC20 *ChildERC20Session) Decimals() (uint8, error) {
	return _ChildERC20.Contract.Decimals(&_ChildERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildERC20 *ChildERC20CallerSession) Decimals() (uint8, error) {
	return _ChildERC20.Contract.Decimals(&_ChildERC20.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildERC20 *ChildERC20Caller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildERC20 *ChildERC20Session) GetNonce(user common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.GetNonce(&_ChildERC20.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildERC20 *ChildERC20CallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _ChildERC20.Contract.GetNonce(&_ChildERC20.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildERC20 *ChildERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildERC20 *ChildERC20Session) Name() (string, error) {
	return _ChildERC20.Contract.Name(&_ChildERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildERC20 *ChildERC20CallerSession) Name() (string, error) {
	return _ChildERC20.Contract.Name(&_ChildERC20.CallOpts)
}

// Predicate is a free data retrieval call binding the contract method 0xe6198705.
//
// Solidity: function predicate() view returns(address)
func (_ChildERC20 *ChildERC20Caller) Predicate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "predicate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Predicate is a free data retrieval call binding the contract method 0xe6198705.
//
// Solidity: function predicate() view returns(address)
func (_ChildERC20 *ChildERC20Session) Predicate() (common.Address, error) {
	return _ChildERC20.Contract.Predicate(&_ChildERC20.CallOpts)
}

// Predicate is a free data retrieval call binding the contract method 0xe6198705.
//
// Solidity: function predicate() view returns(address)
func (_ChildERC20 *ChildERC20CallerSession) Predicate() (common.Address, error) {
	return _ChildERC20.Contract.Predicate(&_ChildERC20.CallOpts)
}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() view returns(address)
func (_ChildERC20 *ChildERC20Caller) RootToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "rootToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() view returns(address)
func (_ChildERC20 *ChildERC20Session) RootToken() (common.Address, error) {
	return _ChildERC20.Contract.RootToken(&_ChildERC20.CallOpts)
}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() view returns(address)
func (_ChildERC20 *ChildERC20CallerSession) RootToken() (common.Address, error) {
	return _ChildERC20.Contract.RootToken(&_ChildERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildERC20 *ChildERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildERC20 *ChildERC20Session) Symbol() (string, error) {
	return _ChildERC20.Contract.Symbol(&_ChildERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildERC20 *ChildERC20CallerSession) Symbol() (string, error) {
	return _ChildERC20.Contract.Symbol(&_ChildERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildERC20 *ChildERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildERC20 *ChildERC20Session) TotalSupply() (*big.Int, error) {
	return _ChildERC20.Contract.TotalSupply(&_ChildERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildERC20 *ChildERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ChildERC20.Contract.TotalSupply(&_ChildERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Approve(&_ChildERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Approve(&_ChildERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Burn(&_ChildERC20.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Burn(&_ChildERC20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildERC20 *ChildERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.DecreaseAllowance(&_ChildERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.DecreaseAllowance(&_ChildERC20.TransactOpts, spender, subtractedValue)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_ChildERC20 *ChildERC20Transactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_ChildERC20 *ChildERC20Session) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildERC20.Contract.ExecuteMetaTransaction(&_ChildERC20.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_ChildERC20 *ChildERC20TransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildERC20.Contract.ExecuteMetaTransaction(&_ChildERC20.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildERC20 *ChildERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.IncreaseAllowance(&_ChildERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.IncreaseAllowance(&_ChildERC20.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address rootToken_, string name_, string symbol_, uint8 decimals_) returns()
func (_ChildERC20 *ChildERC20Transactor) Initialize(opts *bind.TransactOpts, rootToken_ common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "initialize", rootToken_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address rootToken_, string name_, string symbol_, uint8 decimals_) returns()
func (_ChildERC20 *ChildERC20Session) Initialize(rootToken_ common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ChildERC20.Contract.Initialize(&_ChildERC20.TransactOpts, rootToken_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address rootToken_, string name_, string symbol_, uint8 decimals_) returns()
func (_ChildERC20 *ChildERC20TransactorSession) Initialize(rootToken_ common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ChildERC20.Contract.Initialize(&_ChildERC20.TransactOpts, rootToken_, name_, symbol_, decimals_)
}

// InvalidateNext is a paid mutator transaction binding the contract method 0x9b77ef11.
//
// Solidity: function invalidateNext(uint256 offset) returns()
func (_ChildERC20 *ChildERC20Transactor) InvalidateNext(opts *bind.TransactOpts, offset *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "invalidateNext", offset)
}

// InvalidateNext is a paid mutator transaction binding the contract method 0x9b77ef11.
//
// Solidity: function invalidateNext(uint256 offset) returns()
func (_ChildERC20 *ChildERC20Session) InvalidateNext(offset *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.InvalidateNext(&_ChildERC20.TransactOpts, offset)
}

// InvalidateNext is a paid mutator transaction binding the contract method 0x9b77ef11.
//
// Solidity: function invalidateNext(uint256 offset) returns()
func (_ChildERC20 *ChildERC20TransactorSession) InvalidateNext(offset *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.InvalidateNext(&_ChildERC20.TransactOpts, offset)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Mint(&_ChildERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Mint(&_ChildERC20.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Transfer(&_ChildERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.Transfer(&_ChildERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.TransferFrom(&_ChildERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ChildERC20 *ChildERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildERC20.Contract.TransferFrom(&_ChildERC20.TransactOpts, from, to, amount)
}

// ChildERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ChildERC20 contract.
type ChildERC20ApprovalIterator struct {
	Event *ChildERC20Approval // Event containing the contract specifics and raw log

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
func (it *ChildERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20Approval)
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
		it.Event = new(ChildERC20Approval)
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
func (it *ChildERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20Approval represents a Approval event raised by the ChildERC20 contract.
type ChildERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChildERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChildERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChildERC20ApprovalIterator{contract: _ChildERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChildERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChildERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20Approval)
				if err := _ChildERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) ParseApproval(log types.Log) (*ChildERC20Approval, error) {
	event := new(ChildERC20Approval)
	if err := _ChildERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChildERC20 contract.
type ChildERC20InitializedIterator struct {
	Event *ChildERC20Initialized // Event containing the contract specifics and raw log

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
func (it *ChildERC20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20Initialized)
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
		it.Event = new(ChildERC20Initialized)
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
func (it *ChildERC20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20Initialized represents a Initialized event raised by the ChildERC20 contract.
type ChildERC20Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChildERC20 *ChildERC20Filterer) FilterInitialized(opts *bind.FilterOpts) (*ChildERC20InitializedIterator, error) {

	logs, sub, err := _ChildERC20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChildERC20InitializedIterator{contract: _ChildERC20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChildERC20 *ChildERC20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChildERC20Initialized) (event.Subscription, error) {

	logs, sub, err := _ChildERC20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20Initialized)
				if err := _ChildERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ChildERC20 *ChildERC20Filterer) ParseInitialized(log types.Log) (*ChildERC20Initialized, error) {
	event := new(ChildERC20Initialized)
	if err := _ChildERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20MetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the ChildERC20 contract.
type ChildERC20MetaTransactionExecutedIterator struct {
	Event *ChildERC20MetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *ChildERC20MetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20MetaTransactionExecuted)
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
		it.Event = new(ChildERC20MetaTransactionExecuted)
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
func (it *ChildERC20MetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20MetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20MetaTransactionExecuted represents a MetaTransactionExecuted event raised by the ChildERC20 contract.
type ChildERC20MetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildERC20 *ChildERC20Filterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*ChildERC20MetaTransactionExecutedIterator, error) {

	logs, sub, err := _ChildERC20.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &ChildERC20MetaTransactionExecutedIterator{contract: _ChildERC20.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildERC20 *ChildERC20Filterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *ChildERC20MetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _ChildERC20.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20MetaTransactionExecuted)
				if err := _ChildERC20.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildERC20 *ChildERC20Filterer) ParseMetaTransactionExecuted(log types.Log) (*ChildERC20MetaTransactionExecuted, error) {
	event := new(ChildERC20MetaTransactionExecuted)
	if err := _ChildERC20.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ChildERC20 contract.
type ChildERC20TransferIterator struct {
	Event *ChildERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ChildERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildERC20Transfer)
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
		it.Event = new(ChildERC20Transfer)
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
func (it *ChildERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildERC20Transfer represents a Transfer event raised by the ChildERC20 contract.
type ChildERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChildERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChildERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChildERC20TransferIterator{contract: _ChildERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChildERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChildERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildERC20Transfer)
				if err := _ChildERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildERC20 *ChildERC20Filterer) ParseTransfer(log types.Log) (*ChildERC20Transfer, error) {
	event := new(ChildERC20Transfer)
	if err := _ChildERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

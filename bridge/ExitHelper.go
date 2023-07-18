// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// IExitHelperBatchExitInput is an auto generated low-level Go binding around an user-defined struct.
type IExitHelperBatchExitInput struct {
	BlockNumber  *big.Int
	LeafIndex    *big.Int
	UnhashedLeaf []byte
	Proof        [][32]byte
}

// ExitHelperMetaData contains all meta data concerning the ExitHelper contract.
var ExitHelperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExitProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"unhashedLeaf\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIExitHelper.BatchExitInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"}],\"name\":\"batchExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointManager\",\"outputs\":[{\"internalType\":\"contractICheckpointManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"unhashedLeaf\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICheckpointManager\",\"name\":\"newCheckpointManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"processedExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExitHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ExitHelperMetaData.ABI instead.
var ExitHelperABI = ExitHelperMetaData.ABI

// ExitHelper is an auto generated Go binding around an Ethereum contract.
type ExitHelper struct {
	ExitHelperCaller     // Read-only binding to the contract
	ExitHelperTransactor // Write-only binding to the contract
	ExitHelperFilterer   // Log filterer for contract events
}

// ExitHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExitHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExitHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExitHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExitHelperSession struct {
	Contract     *ExitHelper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExitHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExitHelperCallerSession struct {
	Contract *ExitHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExitHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExitHelperTransactorSession struct {
	Contract     *ExitHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExitHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExitHelperRaw struct {
	Contract *ExitHelper // Generic contract binding to access the raw methods on
}

// ExitHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExitHelperCallerRaw struct {
	Contract *ExitHelperCaller // Generic read-only contract binding to access the raw methods on
}

// ExitHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExitHelperTransactorRaw struct {
	Contract *ExitHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExitHelper creates a new instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelper(address common.Address, backend bind.ContractBackend) (*ExitHelper, error) {
	contract, err := bindExitHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExitHelper{ExitHelperCaller: ExitHelperCaller{contract: contract}, ExitHelperTransactor: ExitHelperTransactor{contract: contract}, ExitHelperFilterer: ExitHelperFilterer{contract: contract}}, nil
}

// NewExitHelperCaller creates a new read-only instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperCaller(address common.Address, caller bind.ContractCaller) (*ExitHelperCaller, error) {
	contract, err := bindExitHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExitHelperCaller{contract: contract}, nil
}

// NewExitHelperTransactor creates a new write-only instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ExitHelperTransactor, error) {
	contract, err := bindExitHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExitHelperTransactor{contract: contract}, nil
}

// NewExitHelperFilterer creates a new log filterer instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ExitHelperFilterer, error) {
	contract, err := bindExitHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExitHelperFilterer{contract: contract}, nil
}

// bindExitHelper binds a generic wrapper to an already deployed contract.
func bindExitHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExitHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExitHelper *ExitHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExitHelper.Contract.ExitHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExitHelper *ExitHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExitHelper.Contract.ExitHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExitHelper *ExitHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExitHelper.Contract.ExitHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExitHelper *ExitHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExitHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExitHelper *ExitHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExitHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExitHelper *ExitHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExitHelper.Contract.contract.Transact(opts, method, params...)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperCaller) CheckpointManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExitHelper.contract.Call(opts, &out, "checkpointManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperSession) CheckpointManager() (common.Address, error) {
	return _ExitHelper.Contract.CheckpointManager(&_ExitHelper.CallOpts)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperCallerSession) CheckpointManager() (common.Address, error) {
	return _ExitHelper.Contract.CheckpointManager(&_ExitHelper.CallOpts)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperCaller) ProcessedExits(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ExitHelper.contract.Call(opts, &out, "processedExits", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _ExitHelper.Contract.ProcessedExits(&_ExitHelper.CallOpts, arg0)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperCallerSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _ExitHelper.Contract.ProcessedExits(&_ExitHelper.CallOpts, arg0)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperTransactor) BatchExit(opts *bind.TransactOpts, inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "batchExit", inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.Contract.BatchExit(&_ExitHelper.TransactOpts, inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperTransactorSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.Contract.BatchExit(&_ExitHelper.TransactOpts, inputs)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperTransactor) Exit(opts *bind.TransactOpts, blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "exit", blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.Contract.Exit(&_ExitHelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperTransactorSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.Contract.Exit(&_ExitHelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperTransactor) Initialize(opts *bind.TransactOpts, newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "initialize", newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.Contract.Initialize(&_ExitHelper.TransactOpts, newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperTransactorSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.Contract.Initialize(&_ExitHelper.TransactOpts, newCheckpointManager)
}

// ExitHelperExitProcessedIterator is returned from FilterExitProcessed and is used to iterate over the raw logs and unpacked data for ExitProcessed events raised by the ExitHelper contract.
type ExitHelperExitProcessedIterator struct {
	Event *ExitHelperExitProcessed // Event containing the contract specifics and raw log

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
func (it *ExitHelperExitProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExitHelperExitProcessed)
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
		it.Event = new(ExitHelperExitProcessed)
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
func (it *ExitHelperExitProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExitHelperExitProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExitHelperExitProcessed represents a ExitProcessed event raised by the ExitHelper contract.
type ExitHelperExitProcessed struct {
	Id         *big.Int
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExitProcessed is a free log retrieval operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) FilterExitProcessed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*ExitHelperExitProcessedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ExitHelper.contract.FilterLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return &ExitHelperExitProcessedIterator{contract: _ExitHelper.contract, event: "ExitProcessed", logs: logs, sub: sub}, nil
}

// WatchExitProcessed is a free log subscription operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) WatchExitProcessed(opts *bind.WatchOpts, sink chan<- *ExitHelperExitProcessed, id []*big.Int, success []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ExitHelper.contract.WatchLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExitHelperExitProcessed)
				if err := _ExitHelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
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

// ParseExitProcessed is a log parse operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) ParseExitProcessed(log types.Log) (*ExitHelperExitProcessed, error) {
	event := new(ExitHelperExitProcessed)
	if err := _ExitHelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExitHelperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExitHelper contract.
type ExitHelperInitializedIterator struct {
	Event *ExitHelperInitialized // Event containing the contract specifics and raw log

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
func (it *ExitHelperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExitHelperInitialized)
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
		it.Event = new(ExitHelperInitialized)
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
func (it *ExitHelperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExitHelperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExitHelperInitialized represents a Initialized event raised by the ExitHelper contract.
type ExitHelperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExitHelper *ExitHelperFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExitHelperInitializedIterator, error) {

	logs, sub, err := _ExitHelper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExitHelperInitializedIterator{contract: _ExitHelper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExitHelper *ExitHelperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExitHelperInitialized) (event.Subscription, error) {

	logs, sub, err := _ExitHelper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExitHelperInitialized)
				if err := _ExitHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ExitHelper *ExitHelperFilterer) ParseInitialized(log types.Log) (*ExitHelperInitialized, error) {
	event := new(ExitHelperInitialized)
	if err := _ExitHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

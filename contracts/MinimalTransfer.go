// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minimaltransfer

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// SimpleTransferABI is the input ABI used to generate the binding from.
const SimpleTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repeatCount\",\"type\":\"uint256\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleTransferFuncSigs maps the 4-byte function signature to its string representation.
var SimpleTransferFuncSigs = map[string]string{
	"8892170a": "run(address,address,uint256)",
}

// SimpleTransferBin is the compiled bytecode used for deploying new contracts.
var SimpleTransferBin = "0x608060405234801561001057600080fd5b5061016a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638892170a14610030575b600080fd5b61004361003e3660046100d1565b610045565b005b60005b818110156100af57826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161009591815260200190565b60405180910390a3806100a78161010d565b915050610048565b50505050565b80356001600160a01b03811681146100cc57600080fd5b919050565b6000806000606084860312156100e657600080fd5b6100ef846100b5565b92506100fd602085016100b5565b9150604084013590509250925092565b60006001820161012d57634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220a88e043aee0940c10b9519f98e9128ec577dc27fc271809415917fb042d56b0064736f6c634300080e0033"

// DeploySimpleTransfer deploys a new Conflux contract, binding an instance of SimpleTransfer to it.
func DeploySimpleTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.UnsignedTransaction, *types.Hash, *SimpleTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTransferABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleTransferBin), backend)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &SimpleTransfer{SimpleTransferCaller: SimpleTransferCaller{contract: contract}, SimpleTransferTransactor: SimpleTransferTransactor{contract: contract}, SimpleTransferFilterer: SimpleTransferFilterer{contract: contract}}, nil
}

// SimpleTransfer is an auto generated Go binding around an Conflux contract.
type SimpleTransfer struct {
	SimpleTransferCaller     // Read-only binding to the contract
	SimpleTransferTransactor // Write-only binding to the contract
	SimpleTransferFilterer   // Log filterer for contract events
}

// SimpleTransferCaller is an auto generated read-only Go binding around an Conflux contract.
type SimpleTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransferBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type SimpleTransferBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransferTransactor is an auto generated write-only Go binding around an Conflux contract.
type SimpleTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransferBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type SimpleTransferBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransferFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type SimpleTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransferSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type SimpleTransferSession struct {
	Contract     *SimpleTransfer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleTransferCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type SimpleTransferCallerSession struct {
	Contract *SimpleTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SimpleTransferTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type SimpleTransferTransactorSession struct {
	Contract     *SimpleTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SimpleTransferRaw is an auto generated low-level Go binding around an Conflux contract.
type SimpleTransferRaw struct {
	Contract *SimpleTransfer // Generic contract binding to access the raw methods on
}

// SimpleTransferCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type SimpleTransferCallerRaw struct {
	Contract *SimpleTransferCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransferTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type SimpleTransferTransactorRaw struct {
	Contract *SimpleTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleTransfer creates a new instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransfer(address types.Address, backend bind.ContractBackend) (*SimpleTransfer, error) {
	contract, err := bindSimpleTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleTransfer{SimpleTransferCaller: SimpleTransferCaller{contract: contract}, SimpleTransferTransactor: SimpleTransferTransactor{contract: contract}, SimpleTransferFilterer: SimpleTransferFilterer{contract: contract}}, nil
}

// NewSimpleTransferCaller creates a new read-only instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransferCaller(address types.Address, caller bind.ContractCaller) (*SimpleTransferCaller, error) {
	contract, err := bindSimpleTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferCaller{contract: contract}, nil
}

// NewSimpleTransferTransactor creates a new write-only instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransferTransactor(address types.Address, transactor bind.ContractTransactor) (*SimpleTransferTransactor, error) {
	contract, err := bindSimpleTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferTransactor{contract: contract}, nil
}

// NewSimpleTransferFilterer creates a new log filterer instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransferFilterer(address types.Address, filterer bind.ContractFilterer) (*SimpleTransferFilterer, error) {
	contract, err := bindSimpleTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferFilterer{contract: contract}, nil
}

// NewSimpleTransferCaller creates a new read-only instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransferBulkCaller(address types.Address, caller bind.ContractCaller) (*SimpleTransferBulkCaller, error) {
	contract, err := bindSimpleTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferBulkCaller{contract: contract}, nil
}

// NewSimpleTransferBulkTransactor creates a new write-only instance of SimpleTransfer, bound to a specific deployed contract.
func NewSimpleTransferBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*SimpleTransferBulkTransactor, error) {
	contract, err := bindSimpleTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferBulkTransactor{contract: contract}, nil
}

// bindSimpleTransfer binds a generic wrapper to an already deployed contract.
func bindSimpleTransfer(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleTransfer *SimpleTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleTransfer.Contract.SimpleTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleTransfer *SimpleTransferRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.SimpleTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleTransfer *SimpleTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.SimpleTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleTransfer *SimpleTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleTransfer *SimpleTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleTransfer *SimpleTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.contract.Transact(opts, method, params...)
}

// Run is a paid mutator transaction binding the contract method 0x8892170a.
//
// Solidity: function run(address from, address to, uint256 repeatCount) returns()
func (_SimpleTransfer *SimpleTransferTransactor) Run(opts *bind.TransactOpts, from common.Address, to common.Address, repeatCount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.contract.Transact(opts, "run", from, to, repeatCount)
}

// Run is a paid mutator transaction binding the contract method 0x8892170a.
//
// Solidity: function run(address from, address to, uint256 repeatCount) returns()
func (_SimpleTransfer *SimpleTransferBulkTransactor) Run(opts *bind.TransactOpts, from common.Address, to common.Address, repeatCount *big.Int) types.UnsignedTransaction {
	return _SimpleTransfer.contract.GenUnsignedTransaction(opts, "run", from, to, repeatCount)
}

// Run is a paid mutator transaction binding the contract method 0x8892170a.
//
// Solidity: function run(address from, address to, uint256 repeatCount) returns()
func (_SimpleTransfer *SimpleTransferSession) Run(from common.Address, to common.Address, repeatCount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.Run(&_SimpleTransfer.TransactOpts, from, to, repeatCount)
}

// Run is a paid mutator transaction binding the contract method 0x8892170a.
//
// Solidity: function run(address from, address to, uint256 repeatCount) returns()
func (_SimpleTransfer *SimpleTransferTransactorSession) Run(from common.Address, to common.Address, repeatCount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _SimpleTransfer.Contract.Run(&_SimpleTransfer.TransactOpts, from, to, repeatCount)
}

// SimpleTransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleTransfer contract.
type SimpleTransferTransferIterator struct {
	Event *SimpleTransferTransfer // Event containing the contract specifics and raw log

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
func (it *SimpleTransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTransferTransfer)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleTransferTransfer)
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
func (it *SimpleTransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTransferTransfer represents a Transfer event raised by the SimpleTransfer contract.
type SimpleTransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// SimpleTransferTransferOrChainReorg represents a Transfer subscription event raised by the SimpleTransfer contract.
type SimpleTransferTransferOrChainReorg struct {
	Event      *SimpleTransferTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SimpleTransfer *SimpleTransferFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTransferTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, err := _SimpleTransfer.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTransferTransferIterator{contract: _SimpleTransfer.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SimpleTransfer *SimpleTransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTransferTransferOrChainReorg, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleTransfer.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTransferTransferOrChainReorg)
				event.Event = new(SimpleTransferTransfer)

				if log.ChainReorg == nil {
					if err := _SimpleTransfer.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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
func (_SimpleTransfer *SimpleTransferFilterer) ParseTransfer(log types.Log) (*SimpleTransferTransfer, error) {
	event := new(SimpleTransferTransfer)
	if err := _SimpleTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

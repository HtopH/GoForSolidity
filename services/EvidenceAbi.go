// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package services

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
)

// EvidenceMetaData contains all meta data concerning the Evidence contract.
var EvidenceMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"validateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"addr\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"SaveFileEvent\",\"type\":\"event\"}]",
}

// EvidenceABI is the input ABI used to generate the binding from.
// Deprecated: Use EvidenceMetaData.ABI instead.
var EvidenceABI = EvidenceMetaData.ABI

// Evidence is an auto generated Go binding around an Ethereum contract.
type Evidence struct {
	EvidenceCaller     // Read-only binding to the contract
	EvidenceTransactor // Write-only binding to the contract
	EvidenceFilterer   // Log filterer for contract events
}

// EvidenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvidenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvidenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvidenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvidenceSession struct {
	Contract     *Evidence         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvidenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvidenceCallerSession struct {
	Contract *EvidenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EvidenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvidenceTransactorSession struct {
	Contract     *EvidenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EvidenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvidenceRaw struct {
	Contract *Evidence // Generic contract binding to access the raw methods on
}

// EvidenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvidenceCallerRaw struct {
	Contract *EvidenceCaller // Generic read-only contract binding to access the raw methods on
}

// EvidenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvidenceTransactorRaw struct {
	Contract *EvidenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvidence creates a new instance of Evidence, bound to a specific deployed contract.
func NewEvidence(address common.Address, backend bind.ContractBackend) (*Evidence, error) {
	contract, err := bindEvidence(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Evidence{EvidenceCaller: EvidenceCaller{contract: contract}, EvidenceTransactor: EvidenceTransactor{contract: contract}, EvidenceFilterer: EvidenceFilterer{contract: contract}}, nil
}

// NewEvidenceCaller creates a new read-only instance of Evidence, bound to a specific deployed contract.
func NewEvidenceCaller(address common.Address, caller bind.ContractCaller) (*EvidenceCaller, error) {
	contract, err := bindEvidence(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvidenceCaller{contract: contract}, nil
}

// NewEvidenceTransactor creates a new write-only instance of Evidence, bound to a specific deployed contract.
func NewEvidenceTransactor(address common.Address, transactor bind.ContractTransactor) (*EvidenceTransactor, error) {
	contract, err := bindEvidence(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvidenceTransactor{contract: contract}, nil
}

// NewEvidenceFilterer creates a new log filterer instance of Evidence, bound to a specific deployed contract.
func NewEvidenceFilterer(address common.Address, filterer bind.ContractFilterer) (*EvidenceFilterer, error) {
	contract, err := bindEvidence(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvidenceFilterer{contract: contract}, nil
}

// bindEvidence binds a generic wrapper to an already deployed contract.
func bindEvidence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidence *EvidenceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidence.Contract.EvidenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidence *EvidenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidence.Contract.EvidenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidence *EvidenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidence.Contract.EvidenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evidence *EvidenceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evidence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evidence *EvidenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evidence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evidence *EvidenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evidence.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Evidence *EvidenceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evidence.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Evidence *EvidenceSession) Owner() (common.Address, error) {
	return _Evidence.Contract.Owner(&_Evidence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Evidence *EvidenceCallerSession) Owner() (common.Address, error) {
	return _Evidence.Contract.Owner(&_Evidence.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string hash) view returns(string, string, string)
func (_Evidence *EvidenceCaller) GetData(opts *bind.CallOpts, hash string) (string, string, string, error) {
	var out []interface{}
	err := _Evidence.contract.Call(opts, &out, "getData", hash)

	if err != nil {
		return *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string hash) view returns(string, string, string)
func (_Evidence *EvidenceSession) GetData(hash string) (string, string, string, error) {
	return _Evidence.Contract.GetData(&_Evidence.CallOpts, hash)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string hash) view returns(string, string, string)
func (_Evidence *EvidenceCallerSession) GetData(hash string) (string, string, string, error) {
	return _Evidence.Contract.GetData(&_Evidence.CallOpts, hash)
}

// ValidateHash is a free data retrieval call binding the contract method 0xc2c1a1ed.
//
// Solidity: function validateHash(string hash) view returns()
func (_Evidence *EvidenceCaller) ValidateHash(opts *bind.CallOpts, hash string) error {
	var out []interface{}
	err := _Evidence.contract.Call(opts, &out, "validateHash", hash)

	if err != nil {
		return err
	}

	return err

}

// ValidateHash is a free data retrieval call binding the contract method 0xc2c1a1ed.
//
// Solidity: function validateHash(string hash) view returns()
func (_Evidence *EvidenceSession) ValidateHash(hash string) error {
	return _Evidence.Contract.ValidateHash(&_Evidence.CallOpts, hash)
}

// ValidateHash is a free data retrieval call binding the contract method 0xc2c1a1ed.
//
// Solidity: function validateHash(string hash) view returns()
func (_Evidence *EvidenceCallerSession) ValidateHash(hash string) error {
	return _Evidence.Contract.ValidateHash(&_Evidence.CallOpts, hash)
}

// Allow is a paid mutator transaction binding the contract method 0xff9913e8.
//
// Solidity: function allow(address addr) returns()
func (_Evidence *EvidenceTransactor) Allow(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Evidence.contract.Transact(opts, "allow", addr)
}

// Allow is a paid mutator transaction binding the contract method 0xff9913e8.
//
// Solidity: function allow(address addr) returns()
func (_Evidence *EvidenceSession) Allow(addr common.Address) (*types.Transaction, error) {
	return _Evidence.Contract.Allow(&_Evidence.TransactOpts, addr)
}

// Allow is a paid mutator transaction binding the contract method 0xff9913e8.
//
// Solidity: function allow(address addr) returns()
func (_Evidence *EvidenceTransactorSession) Allow(addr common.Address) (*types.Transaction, error) {
	return _Evidence.Contract.Allow(&_Evidence.TransactOpts, addr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address addr) returns()
func (_Evidence *EvidenceTransactor) Deny(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Evidence.contract.Transact(opts, "deny", addr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address addr) returns()
func (_Evidence *EvidenceSession) Deny(addr common.Address) (*types.Transaction, error) {
	return _Evidence.Contract.Deny(&_Evidence.TransactOpts, addr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address addr) returns()
func (_Evidence *EvidenceTransactorSession) Deny(addr common.Address) (*types.Transaction, error) {
	return _Evidence.Contract.Deny(&_Evidence.TransactOpts, addr)
}

// Save is a paid mutator transaction binding the contract method 0x9b6ec14c.
//
// Solidity: function save(string hash, string addr, string timestamp) returns()
func (_Evidence *EvidenceTransactor) Save(opts *bind.TransactOpts, hash string, addr string, timestamp string) (*types.Transaction, error) {
	return _Evidence.contract.Transact(opts, "save", hash, addr, timestamp)
}

// Save is a paid mutator transaction binding the contract method 0x9b6ec14c.
//
// Solidity: function save(string hash, string addr, string timestamp) returns()
func (_Evidence *EvidenceSession) Save(hash string, addr string, timestamp string) (*types.Transaction, error) {
	return _Evidence.Contract.Save(&_Evidence.TransactOpts, hash, addr, timestamp)
}

// Save is a paid mutator transaction binding the contract method 0x9b6ec14c.
//
// Solidity: function save(string hash, string addr, string timestamp) returns()
func (_Evidence *EvidenceTransactorSession) Save(hash string, addr string, timestamp string) (*types.Transaction, error) {
	return _Evidence.Contract.Save(&_Evidence.TransactOpts, hash, addr, timestamp)
}

// EvidenceSaveFileEventIterator is returned from FilterSaveFileEvent and is used to iterate over the raw logs and unpacked data for SaveFileEvent events raised by the Evidence contract.
type EvidenceSaveFileEventIterator struct {
	Event *EvidenceSaveFileEvent // Event containing the contract specifics and raw log

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
func (it *EvidenceSaveFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvidenceSaveFileEvent)
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
		it.Event = new(EvidenceSaveFileEvent)
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
func (it *EvidenceSaveFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvidenceSaveFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvidenceSaveFileEvent represents a SaveFileEvent event raised by the Evidence contract.
type EvidenceSaveFileEvent struct {
	Hash      string
	Addr      string
	Timestamp string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSaveFileEvent is a free log retrieval operation binding the contract event 0xa8f77fe76a8b376eb233491a8640d0d92e846650d433ce784698d344b3389a6b.
//
// Solidity: event SaveFileEvent(string hash, string addr, string timestamp)
func (_Evidence *EvidenceFilterer) FilterSaveFileEvent(opts *bind.FilterOpts) (*EvidenceSaveFileEventIterator, error) {

	logs, sub, err := _Evidence.contract.FilterLogs(opts, "SaveFileEvent")
	if err != nil {
		return nil, err
	}
	return &EvidenceSaveFileEventIterator{contract: _Evidence.contract, event: "SaveFileEvent", logs: logs, sub: sub}, nil
}

// WatchSaveFileEvent is a free log subscription operation binding the contract event 0xa8f77fe76a8b376eb233491a8640d0d92e846650d433ce784698d344b3389a6b.
//
// Solidity: event SaveFileEvent(string hash, string addr, string timestamp)
func (_Evidence *EvidenceFilterer) WatchSaveFileEvent(opts *bind.WatchOpts, sink chan<- *EvidenceSaveFileEvent) (event.Subscription, error) {

	logs, sub, err := _Evidence.contract.WatchLogs(opts, "SaveFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvidenceSaveFileEvent)
				if err := _Evidence.contract.UnpackLog(event, "SaveFileEvent", log); err != nil {
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

// ParseSaveFileEvent is a log parse operation binding the contract event 0xa8f77fe76a8b376eb233491a8640d0d92e846650d433ce784698d344b3389a6b.
//
// Solidity: event SaveFileEvent(string hash, string addr, string timestamp)
func (_Evidence *EvidenceFilterer) ParseSaveFileEvent(log types.Log) (*EvidenceSaveFileEvent, error) {
	event := new(EvidenceSaveFileEvent)
	if err := _Evidence.contract.UnpackLog(event, "SaveFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

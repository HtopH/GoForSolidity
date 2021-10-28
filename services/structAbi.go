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

// StructAbiMetaData contains all meta data concerning the StructAbi contract.
var StructAbiMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"SM3Key\",\"type\":\"bytes32\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\"},{\"name\":\"storeTime\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SM3Key\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SM3Key\",\"type\":\"bytes32\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"evidneceHash\",\"type\":\"bytes32\"},{\"name\":\"storeTime\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"evidences\",\"outputs\":[{\"name\":\"evidenceHash\",\"type\":\"bytes32\"},{\"name\":\"storeTime\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"SM3Key\",\"type\":\"bytes32\"}],\"name\":\"EvidenceStoreLog\",\"type\":\"event\"}]",
}

// StructAbiABI is the input ABI used to generate the binding from.
// Deprecated: Use StructAbiMetaData.ABI instead.
var StructAbiABI = StructAbiMetaData.ABI

// StructAbi is an auto generated Go binding around an Ethereum contract.
type StructAbi struct {
	StructAbiCaller     // Read-only binding to the contract
	StructAbiTransactor // Write-only binding to the contract
	StructAbiFilterer   // Log filterer for contract events
}

// StructAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructAbiSession struct {
	Contract     *StructAbi        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructAbiCallerSession struct {
	Contract *StructAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StructAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructAbiTransactorSession struct {
	Contract     *StructAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StructAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructAbiRaw struct {
	Contract *StructAbi // Generic contract binding to access the raw methods on
}

// StructAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructAbiCallerRaw struct {
	Contract *StructAbiCaller // Generic read-only contract binding to access the raw methods on
}

// StructAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructAbiTransactorRaw struct {
	Contract *StructAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructAbi creates a new instance of StructAbi, bound to a specific deployed contract.
func NewStructAbi(address common.Address, backend bind.ContractBackend) (*StructAbi, error) {
	contract, err := bindStructAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StructAbi{StructAbiCaller: StructAbiCaller{contract: contract}, StructAbiTransactor: StructAbiTransactor{contract: contract}, StructAbiFilterer: StructAbiFilterer{contract: contract}}, nil
}

// NewStructAbiCaller creates a new read-only instance of StructAbi, bound to a specific deployed contract.
func NewStructAbiCaller(address common.Address, caller bind.ContractCaller) (*StructAbiCaller, error) {
	contract, err := bindStructAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructAbiCaller{contract: contract}, nil
}

// NewStructAbiTransactor creates a new write-only instance of StructAbi, bound to a specific deployed contract.
func NewStructAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*StructAbiTransactor, error) {
	contract, err := bindStructAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructAbiTransactor{contract: contract}, nil
}

// NewStructAbiFilterer creates a new log filterer instance of StructAbi, bound to a specific deployed contract.
func NewStructAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*StructAbiFilterer, error) {
	contract, err := bindStructAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructAbiFilterer{contract: contract}, nil
}

// bindStructAbi binds a generic wrapper to an already deployed contract.
func bindStructAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StructAbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructAbi *StructAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructAbi.Contract.StructAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructAbi *StructAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructAbi.Contract.StructAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructAbi *StructAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructAbi.Contract.StructAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructAbi *StructAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructAbi *StructAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructAbi *StructAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructAbi.Contract.contract.Transact(opts, method, params...)
}

// Evidences is a free data retrieval call binding the contract method 0xd84c09de.
//
// Solidity: function evidences(bytes32 ) view returns(bytes32 evidenceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiCaller) Evidences(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EvidenceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	var out []interface{}
	err := _StructAbi.contract.Call(opts, &out, "evidences", arg0)

	outstruct := new(struct {
		EvidenceHash [32]byte
		StoreTime    [32]byte
		Owner        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EvidenceHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.StoreTime = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Evidences is a free data retrieval call binding the contract method 0xd84c09de.
//
// Solidity: function evidences(bytes32 ) view returns(bytes32 evidenceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiSession) Evidences(arg0 [32]byte) (struct {
	EvidenceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	return _StructAbi.Contract.Evidences(&_StructAbi.CallOpts, arg0)
}

// Evidences is a free data retrieval call binding the contract method 0xd84c09de.
//
// Solidity: function evidences(bytes32 ) view returns(bytes32 evidenceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiCallerSession) Evidences(arg0 [32]byte) (struct {
	EvidenceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	return _StructAbi.Contract.Evidences(&_StructAbi.CallOpts, arg0)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 SM3Key) view returns(bytes32 evidneceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiCaller) Query(opts *bind.CallOpts, SM3Key [32]byte) (struct {
	EvidneceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	var out []interface{}
	err := _StructAbi.contract.Call(opts, &out, "query", SM3Key)

	outstruct := new(struct {
		EvidneceHash [32]byte
		StoreTime    [32]byte
		Owner        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EvidneceHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.StoreTime = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 SM3Key) view returns(bytes32 evidneceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiSession) Query(SM3Key [32]byte) (struct {
	EvidneceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	return _StructAbi.Contract.Query(&_StructAbi.CallOpts, SM3Key)
}

// Query is a free data retrieval call binding the contract method 0xa3b61a9c.
//
// Solidity: function query(bytes32 SM3Key) view returns(bytes32 evidneceHash, bytes32 storeTime, address owner)
func (_StructAbi *StructAbiCallerSession) Query(SM3Key [32]byte) (struct {
	EvidneceHash [32]byte
	StoreTime    [32]byte
	Owner        common.Address
}, error) {
	return _StructAbi.Contract.Query(&_StructAbi.CallOpts, SM3Key)
}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(bytes32 SM3Key) view returns(bool)
func (_StructAbi *StructAbiCaller) Verify(opts *bind.CallOpts, SM3Key [32]byte) (bool, error) {
	var out []interface{}
	err := _StructAbi.contract.Call(opts, &out, "verify", SM3Key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(bytes32 SM3Key) view returns(bool)
func (_StructAbi *StructAbiSession) Verify(SM3Key [32]byte) (bool, error) {
	return _StructAbi.Contract.Verify(&_StructAbi.CallOpts, SM3Key)
}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(bytes32 SM3Key) view returns(bool)
func (_StructAbi *StructAbiCallerSession) Verify(SM3Key [32]byte) (bool, error) {
	return _StructAbi.Contract.Verify(&_StructAbi.CallOpts, SM3Key)
}

// Store is a paid mutator transaction binding the contract method 0x5d8f2640.
//
// Solidity: function store(bytes32 SM3Key, bytes32 evidenceHash, bytes32 storeTime) returns()
func (_StructAbi *StructAbiTransactor) Store(opts *bind.TransactOpts, SM3Key [32]byte, evidenceHash [32]byte, storeTime [32]byte) (*types.Transaction, error) {
	return _StructAbi.contract.Transact(opts, "store", SM3Key, evidenceHash, storeTime)
}

// Store is a paid mutator transaction binding the contract method 0x5d8f2640.
//
// Solidity: function store(bytes32 SM3Key, bytes32 evidenceHash, bytes32 storeTime) returns()
func (_StructAbi *StructAbiSession) Store(SM3Key [32]byte, evidenceHash [32]byte, storeTime [32]byte) (*types.Transaction, error) {
	return _StructAbi.Contract.Store(&_StructAbi.TransactOpts, SM3Key, evidenceHash, storeTime)
}

// Store is a paid mutator transaction binding the contract method 0x5d8f2640.
//
// Solidity: function store(bytes32 SM3Key, bytes32 evidenceHash, bytes32 storeTime) returns()
func (_StructAbi *StructAbiTransactorSession) Store(SM3Key [32]byte, evidenceHash [32]byte, storeTime [32]byte) (*types.Transaction, error) {
	return _StructAbi.Contract.Store(&_StructAbi.TransactOpts, SM3Key, evidenceHash, storeTime)
}

// StructAbiEvidenceStoreLogIterator is returned from FilterEvidenceStoreLog and is used to iterate over the raw logs and unpacked data for EvidenceStoreLog events raised by the StructAbi contract.
type StructAbiEvidenceStoreLogIterator struct {
	Event *StructAbiEvidenceStoreLog // Event containing the contract specifics and raw log

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
func (it *StructAbiEvidenceStoreLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StructAbiEvidenceStoreLog)
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
		it.Event = new(StructAbiEvidenceStoreLog)
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
func (it *StructAbiEvidenceStoreLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StructAbiEvidenceStoreLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StructAbiEvidenceStoreLog represents a EvidenceStoreLog event raised by the StructAbi contract.
type StructAbiEvidenceStoreLog struct {
	SM3Key [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEvidenceStoreLog is a free log retrieval operation binding the contract event 0xe522fd7256a54d741ec6a891748c0ce0cdd1170fd8288f5aeaec966fd8489887.
//
// Solidity: event EvidenceStoreLog(bytes32 SM3Key)
func (_StructAbi *StructAbiFilterer) FilterEvidenceStoreLog(opts *bind.FilterOpts) (*StructAbiEvidenceStoreLogIterator, error) {

	logs, sub, err := _StructAbi.contract.FilterLogs(opts, "EvidenceStoreLog")
	if err != nil {
		return nil, err
	}
	return &StructAbiEvidenceStoreLogIterator{contract: _StructAbi.contract, event: "EvidenceStoreLog", logs: logs, sub: sub}, nil
}

// WatchEvidenceStoreLog is a free log subscription operation binding the contract event 0xe522fd7256a54d741ec6a891748c0ce0cdd1170fd8288f5aeaec966fd8489887.
//
// Solidity: event EvidenceStoreLog(bytes32 SM3Key)
func (_StructAbi *StructAbiFilterer) WatchEvidenceStoreLog(opts *bind.WatchOpts, sink chan<- *StructAbiEvidenceStoreLog) (event.Subscription, error) {

	logs, sub, err := _StructAbi.contract.WatchLogs(opts, "EvidenceStoreLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StructAbiEvidenceStoreLog)
				if err := _StructAbi.contract.UnpackLog(event, "EvidenceStoreLog", log); err != nil {
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

// ParseEvidenceStoreLog is a log parse operation binding the contract event 0xe522fd7256a54d741ec6a891748c0ce0cdd1170fd8288f5aeaec966fd8489887.
//
// Solidity: event EvidenceStoreLog(bytes32 SM3Key)
func (_StructAbi *StructAbiFilterer) ParseEvidenceStoreLog(log types.Log) (*StructAbiEvidenceStoreLog, error) {
	event := new(StructAbiEvidenceStoreLog)
	if err := _StructAbi.contract.UnpackLog(event, "EvidenceStoreLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

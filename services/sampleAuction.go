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

// SampleAuctionMetaData contains all meta data concerning the SampleAuction contract.
var SampleAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_biddingTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AutionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bigger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amout\",\"type\":\"uint256\"}],\"name\":\"HighestBidCreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_biddingTime\",\"type\":\"uint256\"}],\"name\":\"reSetEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SampleAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SampleAuctionMetaData.ABI instead.
var SampleAuctionABI = SampleAuctionMetaData.ABI

// SampleAuction is an auto generated Go binding around an Ethereum contract.
type SampleAuction struct {
	SampleAuctionCaller     // Read-only binding to the contract
	SampleAuctionTransactor // Write-only binding to the contract
	SampleAuctionFilterer   // Log filterer for contract events
}

// SampleAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleAuctionSession struct {
	Contract     *SampleAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleAuctionCallerSession struct {
	Contract *SampleAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SampleAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleAuctionTransactorSession struct {
	Contract     *SampleAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SampleAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleAuctionRaw struct {
	Contract *SampleAuction // Generic contract binding to access the raw methods on
}

// SampleAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleAuctionCallerRaw struct {
	Contract *SampleAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SampleAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleAuctionTransactorRaw struct {
	Contract *SampleAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleAuction creates a new instance of SampleAuction, bound to a specific deployed contract.
func NewSampleAuction(address common.Address, backend bind.ContractBackend) (*SampleAuction, error) {
	contract, err := bindSampleAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleAuction{SampleAuctionCaller: SampleAuctionCaller{contract: contract}, SampleAuctionTransactor: SampleAuctionTransactor{contract: contract}, SampleAuctionFilterer: SampleAuctionFilterer{contract: contract}}, nil
}

// NewSampleAuctionCaller creates a new read-only instance of SampleAuction, bound to a specific deployed contract.
func NewSampleAuctionCaller(address common.Address, caller bind.ContractCaller) (*SampleAuctionCaller, error) {
	contract, err := bindSampleAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleAuctionCaller{contract: contract}, nil
}

// NewSampleAuctionTransactor creates a new write-only instance of SampleAuction, bound to a specific deployed contract.
func NewSampleAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleAuctionTransactor, error) {
	contract, err := bindSampleAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleAuctionTransactor{contract: contract}, nil
}

// NewSampleAuctionFilterer creates a new log filterer instance of SampleAuction, bound to a specific deployed contract.
func NewSampleAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleAuctionFilterer, error) {
	contract, err := bindSampleAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleAuctionFilterer{contract: contract}, nil
}

// bindSampleAuction binds a generic wrapper to an already deployed contract.
func bindSampleAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleAuction *SampleAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleAuction.Contract.SampleAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleAuction *SampleAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleAuction.Contract.SampleAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleAuction *SampleAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleAuction.Contract.SampleAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleAuction *SampleAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleAuction *SampleAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleAuction *SampleAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleAuction.Contract.contract.Transact(opts, method, params...)
}

// AuctionEnd is a free data retrieval call binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() view returns(uint256)
func (_SampleAuction *SampleAuctionCaller) AuctionEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SampleAuction.contract.Call(opts, &out, "auctionEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEnd is a free data retrieval call binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() view returns(uint256)
func (_SampleAuction *SampleAuctionSession) AuctionEnd() (*big.Int, error) {
	return _SampleAuction.Contract.AuctionEnd(&_SampleAuction.CallOpts)
}

// AuctionEnd is a free data retrieval call binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() view returns(uint256)
func (_SampleAuction *SampleAuctionCallerSession) AuctionEnd() (*big.Int, error) {
	return _SampleAuction.Contract.AuctionEnd(&_SampleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SampleAuction *SampleAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SampleAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SampleAuction *SampleAuctionSession) Beneficiary() (common.Address, error) {
	return _SampleAuction.Contract.Beneficiary(&_SampleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SampleAuction *SampleAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _SampleAuction.Contract.Beneficiary(&_SampleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SampleAuction *SampleAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SampleAuction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SampleAuction *SampleAuctionSession) HighestBid() (*big.Int, error) {
	return _SampleAuction.Contract.HighestBid(&_SampleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SampleAuction *SampleAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _SampleAuction.Contract.HighestBid(&_SampleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SampleAuction *SampleAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SampleAuction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SampleAuction *SampleAuctionSession) HighestBidder() (common.Address, error) {
	return _SampleAuction.Contract.HighestBidder(&_SampleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SampleAuction *SampleAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _SampleAuction.Contract.HighestBidder(&_SampleAuction.CallOpts)
}

// Auctionend is a paid mutator transaction binding the contract method 0xb9d1fb61.
//
// Solidity: function auctionend() returns()
func (_SampleAuction *SampleAuctionTransactor) Auctionend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleAuction.contract.Transact(opts, "auctionend")
}

// Auctionend is a paid mutator transaction binding the contract method 0xb9d1fb61.
//
// Solidity: function auctionend() returns()
func (_SampleAuction *SampleAuctionSession) Auctionend() (*types.Transaction, error) {
	return _SampleAuction.Contract.Auctionend(&_SampleAuction.TransactOpts)
}

// Auctionend is a paid mutator transaction binding the contract method 0xb9d1fb61.
//
// Solidity: function auctionend() returns()
func (_SampleAuction *SampleAuctionTransactorSession) Auctionend() (*types.Transaction, error) {
	return _SampleAuction.Contract.Auctionend(&_SampleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SampleAuction *SampleAuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleAuction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SampleAuction *SampleAuctionSession) Bid() (*types.Transaction, error) {
	return _SampleAuction.Contract.Bid(&_SampleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SampleAuction *SampleAuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _SampleAuction.Contract.Bid(&_SampleAuction.TransactOpts)
}

// ReSetEnd is a paid mutator transaction binding the contract method 0x4f1a81f2.
//
// Solidity: function reSetEnd(uint256 _biddingTime) returns()
func (_SampleAuction *SampleAuctionTransactor) ReSetEnd(opts *bind.TransactOpts, _biddingTime *big.Int) (*types.Transaction, error) {
	return _SampleAuction.contract.Transact(opts, "reSetEnd", _biddingTime)
}

// ReSetEnd is a paid mutator transaction binding the contract method 0x4f1a81f2.
//
// Solidity: function reSetEnd(uint256 _biddingTime) returns()
func (_SampleAuction *SampleAuctionSession) ReSetEnd(_biddingTime *big.Int) (*types.Transaction, error) {
	return _SampleAuction.Contract.ReSetEnd(&_SampleAuction.TransactOpts, _biddingTime)
}

// ReSetEnd is a paid mutator transaction binding the contract method 0x4f1a81f2.
//
// Solidity: function reSetEnd(uint256 _biddingTime) returns()
func (_SampleAuction *SampleAuctionTransactorSession) ReSetEnd(_biddingTime *big.Int) (*types.Transaction, error) {
	return _SampleAuction.Contract.ReSetEnd(&_SampleAuction.TransactOpts, _biddingTime)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SampleAuction *SampleAuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleAuction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SampleAuction *SampleAuctionSession) Withdraw() (*types.Transaction, error) {
	return _SampleAuction.Contract.Withdraw(&_SampleAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SampleAuction *SampleAuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SampleAuction.Contract.Withdraw(&_SampleAuction.TransactOpts)
}

// SampleAuctionAutionEndedIterator is returned from FilterAutionEnded and is used to iterate over the raw logs and unpacked data for AutionEnded events raised by the SampleAuction contract.
type SampleAuctionAutionEndedIterator struct {
	Event *SampleAuctionAutionEnded // Event containing the contract specifics and raw log

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
func (it *SampleAuctionAutionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleAuctionAutionEnded)
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
		it.Event = new(SampleAuctionAutionEnded)
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
func (it *SampleAuctionAutionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleAuctionAutionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleAuctionAutionEnded represents a AutionEnded event raised by the SampleAuction contract.
type SampleAuctionAutionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAutionEnded is a free log retrieval operation binding the contract event 0xed2d35808168e05b4ad5305e19a084792967e2121f9cd44615d41f982d130792.
//
// Solidity: event AutionEnded(address winner, uint256 amount)
func (_SampleAuction *SampleAuctionFilterer) FilterAutionEnded(opts *bind.FilterOpts) (*SampleAuctionAutionEndedIterator, error) {

	logs, sub, err := _SampleAuction.contract.FilterLogs(opts, "AutionEnded")
	if err != nil {
		return nil, err
	}
	return &SampleAuctionAutionEndedIterator{contract: _SampleAuction.contract, event: "AutionEnded", logs: logs, sub: sub}, nil
}

// WatchAutionEnded is a free log subscription operation binding the contract event 0xed2d35808168e05b4ad5305e19a084792967e2121f9cd44615d41f982d130792.
//
// Solidity: event AutionEnded(address winner, uint256 amount)
func (_SampleAuction *SampleAuctionFilterer) WatchAutionEnded(opts *bind.WatchOpts, sink chan<- *SampleAuctionAutionEnded) (event.Subscription, error) {

	logs, sub, err := _SampleAuction.contract.WatchLogs(opts, "AutionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleAuctionAutionEnded)
				if err := _SampleAuction.contract.UnpackLog(event, "AutionEnded", log); err != nil {
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

// ParseAutionEnded is a log parse operation binding the contract event 0xed2d35808168e05b4ad5305e19a084792967e2121f9cd44615d41f982d130792.
//
// Solidity: event AutionEnded(address winner, uint256 amount)
func (_SampleAuction *SampleAuctionFilterer) ParseAutionEnded(log types.Log) (*SampleAuctionAutionEnded, error) {
	event := new(SampleAuctionAutionEnded)
	if err := _SampleAuction.contract.UnpackLog(event, "AutionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SampleAuctionHighestBidCreasedIterator is returned from FilterHighestBidCreased and is used to iterate over the raw logs and unpacked data for HighestBidCreased events raised by the SampleAuction contract.
type SampleAuctionHighestBidCreasedIterator struct {
	Event *SampleAuctionHighestBidCreased // Event containing the contract specifics and raw log

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
func (it *SampleAuctionHighestBidCreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleAuctionHighestBidCreased)
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
		it.Event = new(SampleAuctionHighestBidCreased)
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
func (it *SampleAuctionHighestBidCreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleAuctionHighestBidCreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleAuctionHighestBidCreased represents a HighestBidCreased event raised by the SampleAuction contract.
type SampleAuctionHighestBidCreased struct {
	Bigger common.Address
	Amout  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidCreased is a free log retrieval operation binding the contract event 0x0b88694b74574a91345396046e7ac24145ee58e01e6c350e8a88e6fce55ad928.
//
// Solidity: event HighestBidCreased(address bigger, uint256 amout)
func (_SampleAuction *SampleAuctionFilterer) FilterHighestBidCreased(opts *bind.FilterOpts) (*SampleAuctionHighestBidCreasedIterator, error) {

	logs, sub, err := _SampleAuction.contract.FilterLogs(opts, "HighestBidCreased")
	if err != nil {
		return nil, err
	}
	return &SampleAuctionHighestBidCreasedIterator{contract: _SampleAuction.contract, event: "HighestBidCreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidCreased is a free log subscription operation binding the contract event 0x0b88694b74574a91345396046e7ac24145ee58e01e6c350e8a88e6fce55ad928.
//
// Solidity: event HighestBidCreased(address bigger, uint256 amout)
func (_SampleAuction *SampleAuctionFilterer) WatchHighestBidCreased(opts *bind.WatchOpts, sink chan<- *SampleAuctionHighestBidCreased) (event.Subscription, error) {

	logs, sub, err := _SampleAuction.contract.WatchLogs(opts, "HighestBidCreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleAuctionHighestBidCreased)
				if err := _SampleAuction.contract.UnpackLog(event, "HighestBidCreased", log); err != nil {
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

// ParseHighestBidCreased is a log parse operation binding the contract event 0x0b88694b74574a91345396046e7ac24145ee58e01e6c350e8a88e6fce55ad928.
//
// Solidity: event HighestBidCreased(address bigger, uint256 amout)
func (_SampleAuction *SampleAuctionFilterer) ParseHighestBidCreased(log types.Log) (*SampleAuctionHighestBidCreased, error) {
	event := new(SampleAuctionHighestBidCreased)
	if err := _SampleAuction.contract.UnpackLog(event, "HighestBidCreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

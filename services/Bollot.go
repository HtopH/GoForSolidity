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

// BollotMetaData contains all meta data concerning the Bollot contract.
var BollotMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"voteCount\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"name\":\"winner\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"weight\",\"type\":\"uint8\"},{\"name\":\"voted\",\"type\":\"bool\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"vote\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposal\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// BollotABI is the input ABI used to generate the binding from.
// Deprecated: Use BollotMetaData.ABI instead.
var BollotABI = BollotMetaData.ABI

// Bollot is an auto generated Go binding around an Ethereum contract.
type Bollot struct {
	BollotCaller     // Read-only binding to the contract
	BollotTransactor // Write-only binding to the contract
	BollotFilterer   // Log filterer for contract events
}

// BollotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BollotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BollotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BollotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BollotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BollotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BollotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BollotSession struct {
	Contract     *Bollot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BollotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BollotCallerSession struct {
	Contract *BollotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BollotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BollotTransactorSession struct {
	Contract     *BollotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BollotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BollotRaw struct {
	Contract *Bollot // Generic contract binding to access the raw methods on
}

// BollotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BollotCallerRaw struct {
	Contract *BollotCaller // Generic read-only contract binding to access the raw methods on
}

// BollotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BollotTransactorRaw struct {
	Contract *BollotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBollot creates a new instance of Bollot, bound to a specific deployed contract.
func NewBollot(address common.Address, backend bind.ContractBackend) (*Bollot, error) {
	contract, err := bindBollot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bollot{BollotCaller: BollotCaller{contract: contract}, BollotTransactor: BollotTransactor{contract: contract}, BollotFilterer: BollotFilterer{contract: contract}}, nil
}

// NewBollotCaller creates a new read-only instance of Bollot, bound to a specific deployed contract.
func NewBollotCaller(address common.Address, caller bind.ContractCaller) (*BollotCaller, error) {
	contract, err := bindBollot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BollotCaller{contract: contract}, nil
}

// NewBollotTransactor creates a new write-only instance of Bollot, bound to a specific deployed contract.
func NewBollotTransactor(address common.Address, transactor bind.ContractTransactor) (*BollotTransactor, error) {
	contract, err := bindBollot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BollotTransactor{contract: contract}, nil
}

// NewBollotFilterer creates a new log filterer instance of Bollot, bound to a specific deployed contract.
func NewBollotFilterer(address common.Address, filterer bind.ContractFilterer) (*BollotFilterer, error) {
	contract, err := bindBollot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BollotFilterer{contract: contract}, nil
}

// bindBollot binds a generic wrapper to an already deployed contract.
func bindBollot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BollotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bollot *BollotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bollot.Contract.BollotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bollot *BollotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bollot.Contract.BollotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bollot *BollotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bollot.Contract.BollotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bollot *BollotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bollot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bollot *BollotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bollot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bollot *BollotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bollot.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Bollot *BollotCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bollot.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Bollot *BollotSession) Chairperson() (common.Address, error) {
	return _Bollot.Contract.Chairperson(&_Bollot.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Bollot *BollotCallerSession) Chairperson() (common.Address, error) {
	return _Bollot.Contract.Chairperson(&_Bollot.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint8 voteCount)
func (_Bollot *BollotCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount uint8
}, error) {
	var out []interface{}
	err := _Bollot.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint8 voteCount)
func (_Bollot *BollotSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount uint8
}, error) {
	return _Bollot.Contract.Proposals(&_Bollot.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint8 voteCount)
func (_Bollot *BollotCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount uint8
}, error) {
	return _Bollot.Contract.Proposals(&_Bollot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint8 weight, bool voted, address delegate, uint8 vote)
func (_Bollot *BollotCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   uint8
	Voted    bool
	Delegate common.Address
	Vote     uint8
}, error) {
	var out []interface{}
	err := _Bollot.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   uint8
		Voted    bool
		Delegate common.Address
		Vote     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint8 weight, bool voted, address delegate, uint8 vote)
func (_Bollot *BollotSession) Voters(arg0 common.Address) (struct {
	Weight   uint8
	Voted    bool
	Delegate common.Address
	Vote     uint8
}, error) {
	return _Bollot.Contract.Voters(&_Bollot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint8 weight, bool voted, address delegate, uint8 vote)
func (_Bollot *BollotCallerSession) Voters(arg0 common.Address) (struct {
	Weight   uint8
	Voted    bool
	Delegate common.Address
	Vote     uint8
}, error) {
	return _Bollot.Contract.Voters(&_Bollot.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 name)
func (_Bollot *BollotCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bollot.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 name)
func (_Bollot *BollotSession) WinnerName() ([32]byte, error) {
	return _Bollot.Contract.WinnerName(&_Bollot.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 name)
func (_Bollot *BollotCallerSession) WinnerName() ([32]byte, error) {
	return _Bollot.Contract.WinnerName(&_Bollot.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint8 winner)
func (_Bollot *BollotCaller) WinningProposal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bollot.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint8 winner)
func (_Bollot *BollotSession) WinningProposal() (uint8, error) {
	return _Bollot.Contract.WinningProposal(&_Bollot.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint8 winner)
func (_Bollot *BollotCallerSession) WinningProposal() (uint8, error) {
	return _Bollot.Contract.WinningProposal(&_Bollot.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Bollot *BollotTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Bollot.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Bollot *BollotSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Bollot.Contract.Delegate(&_Bollot.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Bollot *BollotTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Bollot.Contract.Delegate(&_Bollot.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Bollot *BollotTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Bollot.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Bollot *BollotSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Bollot.Contract.GiveRightToVote(&_Bollot.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Bollot *BollotTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Bollot.Contract.GiveRightToVote(&_Bollot.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 proposal) returns()
func (_Bollot *BollotTransactor) Vote(opts *bind.TransactOpts, proposal uint8) (*types.Transaction, error) {
	return _Bollot.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 proposal) returns()
func (_Bollot *BollotSession) Vote(proposal uint8) (*types.Transaction, error) {
	return _Bollot.Contract.Vote(&_Bollot.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 proposal) returns()
func (_Bollot *BollotTransactorSession) Vote(proposal uint8) (*types.Transaction, error) {
	return _Bollot.Contract.Vote(&_Bollot.TransactOpts, proposal)
}

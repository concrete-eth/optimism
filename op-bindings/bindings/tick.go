// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TickMetaData contains all meta data concerning the Tick contract.
var TickMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b506000608081905260a052600160c0526100293361002e565b61007e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60805160a05160c05161098a6100ad600039600061034e01526000610325015260006102fc015261098a6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b146100d0578063d4b839921461010f578063e591b2821461012f578063f2fde38b1461014a57600080fd5b80633eaf5d9f1461008d57806354fd4d5014610097578063715018a6146100b5578063776d1a01146100bd575b600080fd5b61009561015d565b005b61009f6102f5565b6040516100ac9190610712565b60405180910390f35b610095610398565b6100956100cb366004610763565b6103ac565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ac565b6001546100ea9073ffffffffffffffffffffffffffffffffffffffff1681565b6100ea73deaddeaddeaddeaddeaddeaddeaddeaddead000181565b610095610158366004610763565b6103fb565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610205576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f5469636b3a206f6e6c7920746865206465706f7369746f72206163636f756e7460448201527f2063616e207469636b000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1661022457565b60015460408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3eaf5d9f00000000000000000000000000000000000000000000000000000000179052905160009273ffffffffffffffffffffffffffffffffffffffff16916102a2916107a0565b6000604051808303816000865af19150503d80600081146102df576040519150601f19603f3d011682016040523d82523d6000602084013e6102e4565b606091505b50509050806102f257600080fd5b50565b60606103207f00000000000000000000000000000000000000000000000000000000000000006104af565b6103497f00000000000000000000000000000000000000000000000000000000000000006104af565b6103727f00000000000000000000000000000000000000000000000000000000000000006104af565b604051602001610384939291906107bc565b604051602081830303815290604052905090565b6103a06105ec565b6103aa600061066d565b565b6103b46105ec565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6104036105ec565b73ffffffffffffffffffffffffffffffffffffffff81166104a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101fc565b6102f28161066d565b6060816000036104f257505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561051c578061050681610861565b91506105159050600a836108c8565b91506104f6565b60008167ffffffffffffffff811115610537576105376108dc565b6040519080825280601f01601f191660200182016040528015610561576020820181803683370190505b5090505b84156105e45761057660018361090b565b9150610583600a86610922565b61058e906030610936565b60f81b8183815181106105a3576105a361094e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506105dd600a866108c8565b9450610565565b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101fc565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005b838110156106fd5781810151838201526020016106e5565b8381111561070c576000848401525b50505050565b60208152600082518060208401526107318160408501602087016106e2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561077557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461079957600080fd5b9392505050565b600082516107b28184602087016106e2565b9190910192915050565b600084516107ce8184602089016106e2565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161080a816001850160208a016106e2565b600192019182015283516108258160028401602088016106e2565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361089257610892610832565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826108d7576108d7610899565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008282101561091d5761091d610832565b500390565b60008261093157610931610899565b500690565b6000821982111561094957610949610832565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// TickABI is the input ABI used to generate the binding from.
// Deprecated: Use TickMetaData.ABI instead.
var TickABI = TickMetaData.ABI

// TickBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TickMetaData.Bin instead.
var TickBin = TickMetaData.Bin

// DeployTick deploys a new Ethereum contract, binding an instance of Tick to it.
func DeployTick(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tick, error) {
	parsed, err := TickMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TickBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tick{TickCaller: TickCaller{contract: contract}, TickTransactor: TickTransactor{contract: contract}, TickFilterer: TickFilterer{contract: contract}}, nil
}

// Tick is an auto generated Go binding around an Ethereum contract.
type Tick struct {
	TickCaller     // Read-only binding to the contract
	TickTransactor // Write-only binding to the contract
	TickFilterer   // Log filterer for contract events
}

// TickCaller is an auto generated read-only Go binding around an Ethereum contract.
type TickCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TickTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TickFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TickSession struct {
	Contract     *Tick             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TickCallerSession struct {
	Contract *TickCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TickTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TickTransactorSession struct {
	Contract     *TickTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickRaw is an auto generated low-level Go binding around an Ethereum contract.
type TickRaw struct {
	Contract *Tick // Generic contract binding to access the raw methods on
}

// TickCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TickCallerRaw struct {
	Contract *TickCaller // Generic read-only contract binding to access the raw methods on
}

// TickTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TickTransactorRaw struct {
	Contract *TickTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTick creates a new instance of Tick, bound to a specific deployed contract.
func NewTick(address common.Address, backend bind.ContractBackend) (*Tick, error) {
	contract, err := bindTick(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tick{TickCaller: TickCaller{contract: contract}, TickTransactor: TickTransactor{contract: contract}, TickFilterer: TickFilterer{contract: contract}}, nil
}

// NewTickCaller creates a new read-only instance of Tick, bound to a specific deployed contract.
func NewTickCaller(address common.Address, caller bind.ContractCaller) (*TickCaller, error) {
	contract, err := bindTick(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TickCaller{contract: contract}, nil
}

// NewTickTransactor creates a new write-only instance of Tick, bound to a specific deployed contract.
func NewTickTransactor(address common.Address, transactor bind.ContractTransactor) (*TickTransactor, error) {
	contract, err := bindTick(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TickTransactor{contract: contract}, nil
}

// NewTickFilterer creates a new log filterer instance of Tick, bound to a specific deployed contract.
func NewTickFilterer(address common.Address, filterer bind.ContractFilterer) (*TickFilterer, error) {
	contract, err := bindTick(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TickFilterer{contract: contract}, nil
}

// bindTick binds a generic wrapper to an already deployed contract.
func bindTick(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TickABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tick *TickRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tick.Contract.TickCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tick *TickRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.Contract.TickTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tick *TickRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tick.Contract.TickTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tick *TickCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tick.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tick *TickTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tick *TickTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tick.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_Tick *TickCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tick.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_Tick *TickSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _Tick.Contract.DEPOSITORACCOUNT(&_Tick.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_Tick *TickCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _Tick.Contract.DEPOSITORACCOUNT(&_Tick.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tick *TickCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tick.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tick *TickSession) Owner() (common.Address, error) {
	return _Tick.Contract.Owner(&_Tick.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tick *TickCallerSession) Owner() (common.Address, error) {
	return _Tick.Contract.Owner(&_Tick.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Tick *TickCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tick.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Tick *TickSession) Target() (common.Address, error) {
	return _Tick.Contract.Target(&_Tick.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_Tick *TickCallerSession) Target() (common.Address, error) {
	return _Tick.Contract.Target(&_Tick.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tick *TickCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tick.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tick *TickSession) Version() (string, error) {
	return _Tick.Contract.Version(&_Tick.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tick *TickCallerSession) Version() (string, error) {
	return _Tick.Contract.Version(&_Tick.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tick *TickTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tick *TickSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tick.Contract.RenounceOwnership(&_Tick.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tick *TickTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tick.Contract.RenounceOwnership(&_Tick.TransactOpts)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Tick *TickTransactor) SetTarget(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Tick.contract.Transact(opts, "setTarget", _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Tick *TickSession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Tick.Contract.SetTarget(&_Tick.TransactOpts, _target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address _target) returns()
func (_Tick *TickTransactorSession) SetTarget(_target common.Address) (*types.Transaction, error) {
	return _Tick.Contract.SetTarget(&_Tick.TransactOpts, _target)
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Tick *TickTransactor) Tick(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.contract.Transact(opts, "tick")
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Tick *TickSession) Tick() (*types.Transaction, error) {
	return _Tick.Contract.Tick(&_Tick.TransactOpts)
}

// Tick is a paid mutator transaction binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() returns()
func (_Tick *TickTransactorSession) Tick() (*types.Transaction, error) {
	return _Tick.Contract.Tick(&_Tick.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tick *TickTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tick.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tick *TickSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tick.Contract.TransferOwnership(&_Tick.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tick *TickTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tick.Contract.TransferOwnership(&_Tick.TransactOpts, newOwner)
}

// TickOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tick contract.
type TickOwnershipTransferredIterator struct {
	Event *TickOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TickOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TickOwnershipTransferred)
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
		it.Event = new(TickOwnershipTransferred)
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
func (it *TickOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TickOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TickOwnershipTransferred represents a OwnershipTransferred event raised by the Tick contract.
type TickOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tick *TickFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TickOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tick.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TickOwnershipTransferredIterator{contract: _Tick.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tick *TickFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TickOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tick.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TickOwnershipTransferred)
				if err := _Tick.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tick *TickFilterer) ParseOwnershipTransferred(log types.Log) (*TickOwnershipTransferred, error) {
	event := new(TickOwnershipTransferred)
	if err := _Tick.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

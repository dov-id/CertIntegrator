// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ICertIntegratorData is an auto generated low-level Go binding around an user-defined struct.
type ICertIntegratorData struct {
	BlockNumber *big.Int
	Root        [32]byte
}

// CertIntegratorContractMetaData contains all meta data concerning the CertIntegratorContract contract.
var CertIntegratorContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"courses_\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"states_\",\"type\":\"bytes32[]\"}],\"name\":\"updateCourseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"course_\",\"type\":\"address\"}],\"name\":\"getData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structICertIntegrator.Data[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"course_\",\"type\":\"address\"}],\"name\":\"getDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"course_\",\"type\":\"address\"}],\"name\":\"getLastData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structICertIntegrator.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CertIntegratorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use CertIntegratorContractMetaData.ABI instead.
var CertIntegratorContractABI = CertIntegratorContractMetaData.ABI

// CertIntegratorContract is an auto generated Go binding around an Ethereum contract.
type CertIntegratorContract struct {
	CertIntegratorContractCaller     // Read-only binding to the contract
	CertIntegratorContractTransactor // Write-only binding to the contract
	CertIntegratorContractFilterer   // Log filterer for contract events
}

// CertIntegratorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertIntegratorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertIntegratorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertIntegratorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertIntegratorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertIntegratorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertIntegratorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertIntegratorContractSession struct {
	Contract     *CertIntegratorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CertIntegratorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertIntegratorContractCallerSession struct {
	Contract *CertIntegratorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CertIntegratorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertIntegratorContractTransactorSession struct {
	Contract     *CertIntegratorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CertIntegratorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertIntegratorContractRaw struct {
	Contract *CertIntegratorContract // Generic contract binding to access the raw methods on
}

// CertIntegratorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertIntegratorContractCallerRaw struct {
	Contract *CertIntegratorContractCaller // Generic read-only contract binding to access the raw methods on
}

// CertIntegratorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertIntegratorContractTransactorRaw struct {
	Contract *CertIntegratorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertIntegratorContract creates a new instance of CertIntegratorContract, bound to a specific deployed contract.
func NewCertIntegratorContract(address common.Address, backend bind.ContractBackend) (*CertIntegratorContract, error) {
	contract, err := bindCertIntegratorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertIntegratorContract{CertIntegratorContractCaller: CertIntegratorContractCaller{contract: contract}, CertIntegratorContractTransactor: CertIntegratorContractTransactor{contract: contract}, CertIntegratorContractFilterer: CertIntegratorContractFilterer{contract: contract}}, nil
}

// NewCertIntegratorContractCaller creates a new read-only instance of CertIntegratorContract, bound to a specific deployed contract.
func NewCertIntegratorContractCaller(address common.Address, caller bind.ContractCaller) (*CertIntegratorContractCaller, error) {
	contract, err := bindCertIntegratorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertIntegratorContractCaller{contract: contract}, nil
}

// NewCertIntegratorContractTransactor creates a new write-only instance of CertIntegratorContract, bound to a specific deployed contract.
func NewCertIntegratorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CertIntegratorContractTransactor, error) {
	contract, err := bindCertIntegratorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertIntegratorContractTransactor{contract: contract}, nil
}

// NewCertIntegratorContractFilterer creates a new log filterer instance of CertIntegratorContract, bound to a specific deployed contract.
func NewCertIntegratorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CertIntegratorContractFilterer, error) {
	contract, err := bindCertIntegratorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertIntegratorContractFilterer{contract: contract}, nil
}

// bindCertIntegratorContract binds a generic wrapper to an already deployed contract.
func bindCertIntegratorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertIntegratorContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertIntegratorContract *CertIntegratorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertIntegratorContract.Contract.CertIntegratorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertIntegratorContract *CertIntegratorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.CertIntegratorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertIntegratorContract *CertIntegratorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.CertIntegratorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertIntegratorContract *CertIntegratorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertIntegratorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertIntegratorContract *CertIntegratorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertIntegratorContract *CertIntegratorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.contract.Transact(opts, method, params...)
}

// ContractData is a free data retrieval call binding the contract method 0x0e54d2bd.
//
// Solidity: function contractData(address , uint256 ) view returns(uint256 blockNumber, bytes32 root)
func (_CertIntegratorContract *CertIntegratorContractCaller) ContractData(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Root        [32]byte
}, error) {
	var out []interface{}
	err := _CertIntegratorContract.contract.Call(opts, &out, "contractData", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Root        [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ContractData is a free data retrieval call binding the contract method 0x0e54d2bd.
//
// Solidity: function contractData(address , uint256 ) view returns(uint256 blockNumber, bytes32 root)
func (_CertIntegratorContract *CertIntegratorContractSession) ContractData(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Root        [32]byte
}, error) {
	return _CertIntegratorContract.Contract.ContractData(&_CertIntegratorContract.CallOpts, arg0, arg1)
}

// ContractData is a free data retrieval call binding the contract method 0x0e54d2bd.
//
// Solidity: function contractData(address , uint256 ) view returns(uint256 blockNumber, bytes32 root)
func (_CertIntegratorContract *CertIntegratorContractCallerSession) ContractData(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Root        [32]byte
}, error) {
	return _CertIntegratorContract.Contract.ContractData(&_CertIntegratorContract.CallOpts, arg0, arg1)
}

// GetData is a free data retrieval call binding the contract method 0x38266b22.
//
// Solidity: function getData(address course_) view returns((uint256,bytes32)[])
func (_CertIntegratorContract *CertIntegratorContractCaller) GetData(opts *bind.CallOpts, course_ common.Address) ([]ICertIntegratorData, error) {
	var out []interface{}
	err := _CertIntegratorContract.contract.Call(opts, &out, "getData", course_)

	if err != nil {
		return *new([]ICertIntegratorData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICertIntegratorData)).(*[]ICertIntegratorData)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x38266b22.
//
// Solidity: function getData(address course_) view returns((uint256,bytes32)[])
func (_CertIntegratorContract *CertIntegratorContractSession) GetData(course_ common.Address) ([]ICertIntegratorData, error) {
	return _CertIntegratorContract.Contract.GetData(&_CertIntegratorContract.CallOpts, course_)
}

// GetData is a free data retrieval call binding the contract method 0x38266b22.
//
// Solidity: function getData(address course_) view returns((uint256,bytes32)[])
func (_CertIntegratorContract *CertIntegratorContractCallerSession) GetData(course_ common.Address) ([]ICertIntegratorData, error) {
	return _CertIntegratorContract.Contract.GetData(&_CertIntegratorContract.CallOpts, course_)
}

// GetDataLength is a free data retrieval call binding the contract method 0x77b0bc13.
//
// Solidity: function getDataLength(address course_) view returns(uint256)
func (_CertIntegratorContract *CertIntegratorContractCaller) GetDataLength(opts *bind.CallOpts, course_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CertIntegratorContract.contract.Call(opts, &out, "getDataLength", course_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDataLength is a free data retrieval call binding the contract method 0x77b0bc13.
//
// Solidity: function getDataLength(address course_) view returns(uint256)
func (_CertIntegratorContract *CertIntegratorContractSession) GetDataLength(course_ common.Address) (*big.Int, error) {
	return _CertIntegratorContract.Contract.GetDataLength(&_CertIntegratorContract.CallOpts, course_)
}

// GetDataLength is a free data retrieval call binding the contract method 0x77b0bc13.
//
// Solidity: function getDataLength(address course_) view returns(uint256)
func (_CertIntegratorContract *CertIntegratorContractCallerSession) GetDataLength(course_ common.Address) (*big.Int, error) {
	return _CertIntegratorContract.Contract.GetDataLength(&_CertIntegratorContract.CallOpts, course_)
}

// GetLastData is a free data retrieval call binding the contract method 0xc07eef72.
//
// Solidity: function getLastData(address course_) view returns((uint256,bytes32))
func (_CertIntegratorContract *CertIntegratorContractCaller) GetLastData(opts *bind.CallOpts, course_ common.Address) (ICertIntegratorData, error) {
	var out []interface{}
	err := _CertIntegratorContract.contract.Call(opts, &out, "getLastData", course_)

	if err != nil {
		return *new(ICertIntegratorData), err
	}

	out0 := *abi.ConvertType(out[0], new(ICertIntegratorData)).(*ICertIntegratorData)

	return out0, err

}

// GetLastData is a free data retrieval call binding the contract method 0xc07eef72.
//
// Solidity: function getLastData(address course_) view returns((uint256,bytes32))
func (_CertIntegratorContract *CertIntegratorContractSession) GetLastData(course_ common.Address) (ICertIntegratorData, error) {
	return _CertIntegratorContract.Contract.GetLastData(&_CertIntegratorContract.CallOpts, course_)
}

// GetLastData is a free data retrieval call binding the contract method 0xc07eef72.
//
// Solidity: function getLastData(address course_) view returns((uint256,bytes32))
func (_CertIntegratorContract *CertIntegratorContractCallerSession) GetLastData(course_ common.Address) (ICertIntegratorData, error) {
	return _CertIntegratorContract.Contract.GetLastData(&_CertIntegratorContract.CallOpts, course_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertIntegratorContract *CertIntegratorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CertIntegratorContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertIntegratorContract *CertIntegratorContractSession) Owner() (common.Address, error) {
	return _CertIntegratorContract.Contract.Owner(&_CertIntegratorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CertIntegratorContract *CertIntegratorContractCallerSession) Owner() (common.Address, error) {
	return _CertIntegratorContract.Contract.Owner(&_CertIntegratorContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertIntegratorContract *CertIntegratorContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertIntegratorContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertIntegratorContract *CertIntegratorContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.RenounceOwnership(&_CertIntegratorContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CertIntegratorContract *CertIntegratorContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.RenounceOwnership(&_CertIntegratorContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertIntegratorContract *CertIntegratorContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CertIntegratorContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertIntegratorContract *CertIntegratorContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.TransferOwnership(&_CertIntegratorContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CertIntegratorContract *CertIntegratorContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.TransferOwnership(&_CertIntegratorContract.TransactOpts, newOwner)
}

// UpdateCourseState is a paid mutator transaction binding the contract method 0x2c832fab.
//
// Solidity: function updateCourseState(address[] courses_, bytes32[] states_) returns()
func (_CertIntegratorContract *CertIntegratorContractTransactor) UpdateCourseState(opts *bind.TransactOpts, courses_ []common.Address, states_ [][32]byte) (*types.Transaction, error) {
	return _CertIntegratorContract.contract.Transact(opts, "updateCourseState", courses_, states_)
}

// UpdateCourseState is a paid mutator transaction binding the contract method 0x2c832fab.
//
// Solidity: function updateCourseState(address[] courses_, bytes32[] states_) returns()
func (_CertIntegratorContract *CertIntegratorContractSession) UpdateCourseState(courses_ []common.Address, states_ [][32]byte) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.UpdateCourseState(&_CertIntegratorContract.TransactOpts, courses_, states_)
}

// UpdateCourseState is a paid mutator transaction binding the contract method 0x2c832fab.
//
// Solidity: function updateCourseState(address[] courses_, bytes32[] states_) returns()
func (_CertIntegratorContract *CertIntegratorContractTransactorSession) UpdateCourseState(courses_ []common.Address, states_ [][32]byte) (*types.Transaction, error) {
	return _CertIntegratorContract.Contract.UpdateCourseState(&_CertIntegratorContract.TransactOpts, courses_, states_)
}

// CertIntegratorContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CertIntegratorContract contract.
type CertIntegratorContractOwnershipTransferredIterator struct {
	Event *CertIntegratorContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CertIntegratorContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertIntegratorContractOwnershipTransferred)
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
		it.Event = new(CertIntegratorContractOwnershipTransferred)
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
func (it *CertIntegratorContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertIntegratorContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertIntegratorContractOwnershipTransferred represents a OwnershipTransferred event raised by the CertIntegratorContract contract.
type CertIntegratorContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CertIntegratorContract *CertIntegratorContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CertIntegratorContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertIntegratorContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CertIntegratorContractOwnershipTransferredIterator{contract: _CertIntegratorContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CertIntegratorContract *CertIntegratorContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CertIntegratorContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CertIntegratorContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertIntegratorContractOwnershipTransferred)
				if err := _CertIntegratorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CertIntegratorContract *CertIntegratorContractFilterer) ParseOwnershipTransferred(log types.Log) (*CertIntegratorContractOwnershipTransferred, error) {
	event := new(CertIntegratorContractOwnershipTransferred)
	if err := _CertIntegratorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

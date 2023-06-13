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

// ITokenFactoryDeployTokenContractParams is an auto generated low-level Go binding around an user-defined struct.
type ITokenFactoryDeployTokenContractParams struct {
	TokenContractId *big.Int
	TokenName       string
	TokenSymbol     string
}

// TokenFactoryContractMetaData contains all meta data concerning the TokenFactoryContract contract.
var TokenFactoryContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenContractAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenContractId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITokenFactory.DeployTokenContractParams\",\"name\":\"tokenContractParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"__TokenFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenContractId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structITokenFactory.DeployTokenContractParams\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"deployTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"setNewImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenContractByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContractsBeacon\",\"outputs\":[{\"internalType\":\"contractProxyBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenFactoryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFactoryContractMetaData.ABI instead.
var TokenFactoryContractABI = TokenFactoryContractMetaData.ABI

// TokenFactoryContract is an auto generated Go binding around an Ethereum contract.
type TokenFactoryContract struct {
	TokenFactoryContractCaller     // Read-only binding to the contract
	TokenFactoryContractTransactor // Write-only binding to the contract
	TokenFactoryContractFilterer   // Log filterer for contract events
}

// TokenFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFactoryContractSession struct {
	Contract     *TokenFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFactoryContractCallerSession struct {
	Contract *TokenFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TokenFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFactoryContractTransactorSession struct {
	Contract     *TokenFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TokenFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFactoryContractRaw struct {
	Contract *TokenFactoryContract // Generic contract binding to access the raw methods on
}

// TokenFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFactoryContractCallerRaw struct {
	Contract *TokenFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFactoryContractTransactorRaw struct {
	Contract *TokenFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactoryContract creates a new instance of TokenFactoryContract, bound to a specific deployed contract.
func NewTokenFactoryContract(address common.Address, backend bind.ContractBackend) (*TokenFactoryContract, error) {
	contract, err := bindTokenFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContract{TokenFactoryContractCaller: TokenFactoryContractCaller{contract: contract}, TokenFactoryContractTransactor: TokenFactoryContractTransactor{contract: contract}, TokenFactoryContractFilterer: TokenFactoryContractFilterer{contract: contract}}, nil
}

// NewTokenFactoryContractCaller creates a new read-only instance of TokenFactoryContract, bound to a specific deployed contract.
func NewTokenFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*TokenFactoryContractCaller, error) {
	contract, err := bindTokenFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractCaller{contract: contract}, nil
}

// NewTokenFactoryContractTransactor creates a new write-only instance of TokenFactoryContract, bound to a specific deployed contract.
func NewTokenFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFactoryContractTransactor, error) {
	contract, err := bindTokenFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractTransactor{contract: contract}, nil
}

// NewTokenFactoryContractFilterer creates a new log filterer instance of TokenFactoryContract, bound to a specific deployed contract.
func NewTokenFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFactoryContractFilterer, error) {
	contract, err := bindTokenFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractFilterer{contract: contract}, nil
}

// bindTokenFactoryContract binds a generic wrapper to an already deployed contract.
func bindTokenFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactoryContract *TokenFactoryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactoryContract.Contract.TokenFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactoryContract *TokenFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TokenFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactoryContract *TokenFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TokenFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactoryContract *TokenFactoryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactoryContract *TokenFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactoryContract *TokenFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_TokenFactoryContract *TokenFactoryContractCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_TokenFactoryContract *TokenFactoryContractSession) BaseTokenContractsURI() (string, error) {
	return _TokenFactoryContract.Contract.BaseTokenContractsURI(&_TokenFactoryContract.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) BaseTokenContractsURI() (string, error) {
	return _TokenFactoryContract.Contract.BaseTokenContractsURI(&_TokenFactoryContract.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_TokenFactoryContract *TokenFactoryContractCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_TokenFactoryContract *TokenFactoryContractSession) GetTokenContractsCount() (*big.Int, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsCount(&_TokenFactoryContract.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsCount(&_TokenFactoryContract.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCaller) GetTokenContractsImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "getTokenContractsImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractSession) GetTokenContractsImpl() (common.Address, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsImpl(&_TokenFactoryContract.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) GetTokenContractsImpl() (common.Address, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsImpl(&_TokenFactoryContract.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_TokenFactoryContract *TokenFactoryContractCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_TokenFactoryContract *TokenFactoryContractSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsPart(&_TokenFactoryContract.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_TokenFactoryContract *TokenFactoryContractCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _TokenFactoryContract.Contract.GetTokenContractsPart(&_TokenFactoryContract.CallOpts, offset_, limit_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractSession) Owner() (common.Address, error) {
	return _TokenFactoryContract.Contract.Owner(&_TokenFactoryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) Owner() (common.Address, error) {
	return _TokenFactoryContract.Contract.Owner(&_TokenFactoryContract.CallOpts)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCaller) TokenContractByIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "tokenContractByIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_TokenFactoryContract *TokenFactoryContractSession) TokenContractByIndex(arg0 *big.Int) (common.Address, error) {
	return _TokenFactoryContract.Contract.TokenContractByIndex(&_TokenFactoryContract.CallOpts, arg0)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) TokenContractByIndex(arg0 *big.Int) (common.Address, error) {
	return _TokenFactoryContract.Contract.TokenContractByIndex(&_TokenFactoryContract.CallOpts, arg0)
}

// TokenContractsBeacon is a free data retrieval call binding the contract method 0x1e7292cf.
//
// Solidity: function tokenContractsBeacon() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCaller) TokenContractsBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactoryContract.contract.Call(opts, &out, "tokenContractsBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContractsBeacon is a free data retrieval call binding the contract method 0x1e7292cf.
//
// Solidity: function tokenContractsBeacon() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractSession) TokenContractsBeacon() (common.Address, error) {
	return _TokenFactoryContract.Contract.TokenContractsBeacon(&_TokenFactoryContract.CallOpts)
}

// TokenContractsBeacon is a free data retrieval call binding the contract method 0x1e7292cf.
//
// Solidity: function tokenContractsBeacon() view returns(address)
func (_TokenFactoryContract *TokenFactoryContractCallerSession) TokenContractsBeacon() (common.Address, error) {
	return _TokenFactoryContract.Contract.TokenContractsBeacon(&_TokenFactoryContract.CallOpts)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xcbba3695.
//
// Solidity: function __TokenFactory_init(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) TokenFactoryInit(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "__TokenFactory_init", baseTokenContractsURI_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xcbba3695.
//
// Solidity: function __TokenFactory_init(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractSession) TokenFactoryInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TokenFactoryInit(&_TokenFactoryContract.TransactOpts, baseTokenContractsURI_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xcbba3695.
//
// Solidity: function __TokenFactory_init(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) TokenFactoryInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TokenFactoryInit(&_TokenFactoryContract.TransactOpts, baseTokenContractsURI_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0x0fc35e56.
//
// Solidity: function deployTokenContract((uint256,string,string) params_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) DeployTokenContract(opts *bind.TransactOpts, params_ ITokenFactoryDeployTokenContractParams) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "deployTokenContract", params_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0x0fc35e56.
//
// Solidity: function deployTokenContract((uint256,string,string) params_) returns()
func (_TokenFactoryContract *TokenFactoryContractSession) DeployTokenContract(params_ ITokenFactoryDeployTokenContractParams) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.DeployTokenContract(&_TokenFactoryContract.TransactOpts, params_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0x0fc35e56.
//
// Solidity: function deployTokenContract((uint256,string,string) params_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) DeployTokenContract(params_ ITokenFactoryDeployTokenContractParams) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.DeployTokenContract(&_TokenFactoryContract.TransactOpts, params_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactoryContract *TokenFactoryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.RenounceOwnership(&_TokenFactoryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.RenounceOwnership(&_TokenFactoryContract.TransactOpts)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.SetBaseTokenContractsURI(&_TokenFactoryContract.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.SetBaseTokenContractsURI(&_TokenFactoryContract.TransactOpts, baseTokenContractsURI_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) SetNewImplementation(opts *bind.TransactOpts, newImplementation_ common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "setNewImplementation", newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_TokenFactoryContract *TokenFactoryContractSession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.SetNewImplementation(&_TokenFactoryContract.TransactOpts, newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.SetNewImplementation(&_TokenFactoryContract.TransactOpts, newImplementation_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactoryContract *TokenFactoryContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TransferOwnership(&_TokenFactoryContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactoryContract *TokenFactoryContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactoryContract.Contract.TransferOwnership(&_TokenFactoryContract.TransactOpts, newOwner)
}

// TokenFactoryContractBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the TokenFactoryContract contract.
type TokenFactoryContractBaseTokenContractsURIUpdatedIterator struct {
	Event *TokenFactoryContractBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryContractBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryContractBaseTokenContractsURIUpdated)
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
		it.Event = new(TokenFactoryContractBaseTokenContractsURIUpdated)
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
func (it *TokenFactoryContractBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryContractBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryContractBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the TokenFactoryContract contract.
type TokenFactoryContractBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_TokenFactoryContract *TokenFactoryContractFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*TokenFactoryContractBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _TokenFactoryContract.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractBaseTokenContractsURIUpdatedIterator{contract: _TokenFactoryContract.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_TokenFactoryContract *TokenFactoryContractFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *TokenFactoryContractBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenFactoryContract.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryContractBaseTokenContractsURIUpdated)
				if err := _TokenFactoryContract.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
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

// ParseBaseTokenContractsURIUpdated is a log parse operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_TokenFactoryContract *TokenFactoryContractFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*TokenFactoryContractBaseTokenContractsURIUpdated, error) {
	event := new(TokenFactoryContractBaseTokenContractsURIUpdated)
	if err := _TokenFactoryContract.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenFactoryContract contract.
type TokenFactoryContractInitializedIterator struct {
	Event *TokenFactoryContractInitialized // Event containing the contract specifics and raw log

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
func (it *TokenFactoryContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryContractInitialized)
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
		it.Event = new(TokenFactoryContractInitialized)
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
func (it *TokenFactoryContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryContractInitialized represents a Initialized event raised by the TokenFactoryContract contract.
type TokenFactoryContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenFactoryContract *TokenFactoryContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenFactoryContractInitializedIterator, error) {

	logs, sub, err := _TokenFactoryContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractInitializedIterator{contract: _TokenFactoryContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenFactoryContract *TokenFactoryContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenFactoryContractInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenFactoryContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryContractInitialized)
				if err := _TokenFactoryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenFactoryContract *TokenFactoryContractFilterer) ParseInitialized(log types.Log) (*TokenFactoryContractInitialized, error) {
	event := new(TokenFactoryContractInitialized)
	if err := _TokenFactoryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenFactoryContract contract.
type TokenFactoryContractOwnershipTransferredIterator struct {
	Event *TokenFactoryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenFactoryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryContractOwnershipTransferred)
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
		it.Event = new(TokenFactoryContractOwnershipTransferred)
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
func (it *TokenFactoryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryContractOwnershipTransferred represents a OwnershipTransferred event raised by the TokenFactoryContract contract.
type TokenFactoryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactoryContract *TokenFactoryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenFactoryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFactoryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractOwnershipTransferredIterator{contract: _TokenFactoryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactoryContract *TokenFactoryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenFactoryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFactoryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryContractOwnershipTransferred)
				if err := _TokenFactoryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenFactoryContract *TokenFactoryContractFilterer) ParseOwnershipTransferred(log types.Log) (*TokenFactoryContractOwnershipTransferred, error) {
	event := new(TokenFactoryContractOwnershipTransferred)
	if err := _TokenFactoryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryContractTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the TokenFactoryContract contract.
type TokenFactoryContractTokenContractDeployedIterator struct {
	Event *TokenFactoryContractTokenContractDeployed // Event containing the contract specifics and raw log

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
func (it *TokenFactoryContractTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryContractTokenContractDeployed)
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
		it.Event = new(TokenFactoryContractTokenContractDeployed)
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
func (it *TokenFactoryContractTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryContractTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryContractTokenContractDeployed represents a TokenContractDeployed event raised by the TokenFactoryContract contract.
type TokenFactoryContractTokenContractDeployed struct {
	NewTokenContractAddr common.Address
	TokenContractParams  ITokenFactoryDeployTokenContractParams
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTokenContractDeployed is a free log retrieval operation binding the contract event 0x9b4e4b661d7431d4aa910f9c83e3ee6fc76f67022d19f6ab3900809d2c67c911.
//
// Solidity: event TokenContractDeployed(address newTokenContractAddr, (uint256,string,string) tokenContractParams)
func (_TokenFactoryContract *TokenFactoryContractFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts) (*TokenFactoryContractTokenContractDeployedIterator, error) {

	logs, sub, err := _TokenFactoryContract.contract.FilterLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return &TokenFactoryContractTokenContractDeployedIterator{contract: _TokenFactoryContract.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x9b4e4b661d7431d4aa910f9c83e3ee6fc76f67022d19f6ab3900809d2c67c911.
//
// Solidity: event TokenContractDeployed(address newTokenContractAddr, (uint256,string,string) tokenContractParams)
func (_TokenFactoryContract *TokenFactoryContractFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *TokenFactoryContractTokenContractDeployed) (event.Subscription, error) {

	logs, sub, err := _TokenFactoryContract.contract.WatchLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryContractTokenContractDeployed)
				if err := _TokenFactoryContract.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
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

// ParseTokenContractDeployed is a log parse operation binding the contract event 0x9b4e4b661d7431d4aa910f9c83e3ee6fc76f67022d19f6ab3900809d2c67c911.
//
// Solidity: event TokenContractDeployed(address newTokenContractAddr, (uint256,string,string) tokenContractParams)
func (_TokenFactoryContract *TokenFactoryContractFilterer) ParseTokenContractDeployed(log types.Log) (*TokenFactoryContractTokenContractDeployed, error) {
	event := new(TokenFactoryContractTokenContractDeployed)
	if err := _TokenFactoryContract.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

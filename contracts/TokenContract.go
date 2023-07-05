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

// ITokenContractTokenContractInitParams is an auto generated low-level Go binding around an user-defined struct.
type ITokenContractTokenContractInitParams struct {
	TokenName        string
	TokenSymbol      string
	TokenFactoryAddr common.Address
	Admin            common.Address
}

// TokenContractMetaData contains all meta data concerning the TokenContract contract.
var TokenContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"SuccessfullyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenFactoryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"internalType\":\"structITokenContract.TokenContractInitParams\",\"name\":\"initParams_\",\"type\":\"tuple\"}],\"name\":\"__TokenContract_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"existingTokenURIs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserTokenIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setNewAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenContractMetaData.ABI instead.
var TokenContractABI = TokenContractMetaData.ABI

// TokenContract is an auto generated Go binding around an Ethereum contract.
type TokenContract struct {
	TokenContractCaller     // Read-only binding to the contract
	TokenContractTransactor // Write-only binding to the contract
	TokenContractFilterer   // Log filterer for contract events
}

// TokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenContractSession struct {
	Contract     *TokenContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenContractCallerSession struct {
	Contract *TokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenContractTransactorSession struct {
	Contract     *TokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenContractRaw struct {
	Contract *TokenContract // Generic contract binding to access the raw methods on
}

// TokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenContractCallerRaw struct {
	Contract *TokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// TokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenContractTransactorRaw struct {
	Contract *TokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenContract creates a new instance of TokenContract, bound to a specific deployed contract.
func NewTokenContract(address common.Address, backend bind.ContractBackend) (*TokenContract, error) {
	contract, err := bindTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenContract{TokenContractCaller: TokenContractCaller{contract: contract}, TokenContractTransactor: TokenContractTransactor{contract: contract}, TokenContractFilterer: TokenContractFilterer{contract: contract}}, nil
}

// NewTokenContractCaller creates a new read-only instance of TokenContract, bound to a specific deployed contract.
func NewTokenContractCaller(address common.Address, caller bind.ContractCaller) (*TokenContractCaller, error) {
	contract, err := bindTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenContractCaller{contract: contract}, nil
}

// NewTokenContractTransactor creates a new write-only instance of TokenContract, bound to a specific deployed contract.
func NewTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenContractTransactor, error) {
	contract, err := bindTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenContractTransactor{contract: contract}, nil
}

// NewTokenContractFilterer creates a new log filterer instance of TokenContract, bound to a specific deployed contract.
func NewTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenContractFilterer, error) {
	contract, err := bindTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenContractFilterer{contract: contract}, nil
}

// bindTokenContract binds a generic wrapper to an already deployed contract.
func bindTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenContract *TokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenContract.Contract.TokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenContract *TokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContract.Contract.TokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenContract *TokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenContract.Contract.TokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenContract *TokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenContract *TokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenContract *TokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenContract *TokenContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenContract *TokenContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenContract.Contract.BalanceOf(&_TokenContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenContract *TokenContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenContract.Contract.BalanceOf(&_TokenContract.CallOpts, owner)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_TokenContract *TokenContractCaller) ExistingTokenURIs(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "existingTokenURIs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_TokenContract *TokenContractSession) ExistingTokenURIs(arg0 string) (bool, error) {
	return _TokenContract.Contract.ExistingTokenURIs(&_TokenContract.CallOpts, arg0)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_TokenContract *TokenContractCallerSession) ExistingTokenURIs(arg0 string) (bool, error) {
	return _TokenContract.Contract.ExistingTokenURIs(&_TokenContract.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TokenContract.Contract.GetApproved(&_TokenContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TokenContract.Contract.GetApproved(&_TokenContract.CallOpts, tokenId)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_TokenContract *TokenContractCaller) GetUserTokenIDs(opts *bind.CallOpts, userAddr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "getUserTokenIDs", userAddr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_TokenContract *TokenContractSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _TokenContract.Contract.GetUserTokenIDs(&_TokenContract.CallOpts, userAddr_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_TokenContract *TokenContractCallerSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _TokenContract.Contract.GetUserTokenIDs(&_TokenContract.CallOpts, userAddr_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenContract *TokenContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenContract *TokenContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TokenContract.Contract.IsApprovedForAll(&_TokenContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenContract *TokenContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TokenContract.Contract.IsApprovedForAll(&_TokenContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenContract *TokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenContract *TokenContractSession) Name() (string, error) {
	return _TokenContract.Contract.Name(&_TokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenContract *TokenContractCallerSession) Name() (string, error) {
	return _TokenContract.Contract.Name(&_TokenContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TokenContract.Contract.OwnerOf(&_TokenContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenContract *TokenContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TokenContract.Contract.OwnerOf(&_TokenContract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenContract *TokenContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenContract *TokenContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenContract.Contract.SupportsInterface(&_TokenContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenContract *TokenContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenContract.Contract.SupportsInterface(&_TokenContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenContract *TokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenContract *TokenContractSession) Symbol() (string, error) {
	return _TokenContract.Contract.Symbol(&_TokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenContract *TokenContractCallerSession) Symbol() (string, error) {
	return _TokenContract.Contract.Symbol(&_TokenContract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenContract *TokenContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenContract *TokenContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TokenContract.Contract.TokenByIndex(&_TokenContract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenContract *TokenContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TokenContract.Contract.TokenByIndex(&_TokenContract.CallOpts, index)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TokenContract *TokenContractCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TokenContract *TokenContractSession) TokenFactory() (common.Address, error) {
	return _TokenContract.Contract.TokenFactory(&_TokenContract.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TokenContract *TokenContractCallerSession) TokenFactory() (common.Address, error) {
	return _TokenContract.Contract.TokenFactory(&_TokenContract.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenContract *TokenContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenContract *TokenContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TokenContract.Contract.TokenOfOwnerByIndex(&_TokenContract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenContract *TokenContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TokenContract.Contract.TokenOfOwnerByIndex(&_TokenContract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_TokenContract *TokenContractCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_TokenContract *TokenContractSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _TokenContract.Contract.TokenURI(&_TokenContract.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_TokenContract *TokenContractCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _TokenContract.Contract.TokenURI(&_TokenContract.CallOpts, tokenId_)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenContract *TokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenContract *TokenContractSession) TotalSupply() (*big.Int, error) {
	return _TokenContract.Contract.TotalSupply(&_TokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenContract *TokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenContract.Contract.TotalSupply(&_TokenContract.CallOpts)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x2585a985.
//
// Solidity: function __TokenContract_init((string,string,address,address) initParams_) returns()
func (_TokenContract *TokenContractTransactor) TokenContractInit(opts *bind.TransactOpts, initParams_ ITokenContractTokenContractInitParams) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "__TokenContract_init", initParams_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x2585a985.
//
// Solidity: function __TokenContract_init((string,string,address,address) initParams_) returns()
func (_TokenContract *TokenContractSession) TokenContractInit(initParams_ ITokenContractTokenContractInitParams) (*types.Transaction, error) {
	return _TokenContract.Contract.TokenContractInit(&_TokenContract.TransactOpts, initParams_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x2585a985.
//
// Solidity: function __TokenContract_init((string,string,address,address) initParams_) returns()
func (_TokenContract *TokenContractTransactorSession) TokenContractInit(initParams_ ITokenContractTokenContractInitParams) (*types.Transaction, error) {
	return _TokenContract.Contract.TokenContractInit(&_TokenContract.TransactOpts, initParams_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.Approve(&_TokenContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.Approve(&_TokenContract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenContract *TokenContractSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.Burn(&_TokenContract.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.Burn(&_TokenContract.TransactOpts, tokenId)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_TokenContract *TokenContractTransactor) DeleteAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "deleteAdmin", admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_TokenContract *TokenContractSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _TokenContract.Contract.DeleteAdmin(&_TokenContract.TransactOpts, admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_TokenContract *TokenContractTransactorSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _TokenContract.Contract.DeleteAdmin(&_TokenContract.TransactOpts, admin)
}

// MintToken is a paid mutator transaction binding the contract method 0x3d02d0c9.
//
// Solidity: function mintToken(address to, string tokenURI_) returns(uint256)
func (_TokenContract *TokenContractTransactor) MintToken(opts *bind.TransactOpts, to common.Address, tokenURI_ string) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "mintToken", to, tokenURI_)
}

// MintToken is a paid mutator transaction binding the contract method 0x3d02d0c9.
//
// Solidity: function mintToken(address to, string tokenURI_) returns(uint256)
func (_TokenContract *TokenContractSession) MintToken(to common.Address, tokenURI_ string) (*types.Transaction, error) {
	return _TokenContract.Contract.MintToken(&_TokenContract.TransactOpts, to, tokenURI_)
}

// MintToken is a paid mutator transaction binding the contract method 0x3d02d0c9.
//
// Solidity: function mintToken(address to, string tokenURI_) returns(uint256)
func (_TokenContract *TokenContractTransactorSession) MintToken(to common.Address, tokenURI_ string) (*types.Transaction, error) {
	return _TokenContract.Contract.MintToken(&_TokenContract.TransactOpts, to, tokenURI_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.SafeTransferFrom(&_TokenContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.SafeTransferFrom(&_TokenContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenContract *TokenContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenContract *TokenContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenContract.Contract.SafeTransferFrom0(&_TokenContract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenContract *TokenContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenContract.Contract.SafeTransferFrom0(&_TokenContract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenContract *TokenContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenContract *TokenContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenContract.Contract.SetApprovalForAll(&_TokenContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenContract *TokenContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenContract.Contract.SetApprovalForAll(&_TokenContract.TransactOpts, operator, approved)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address admin) returns()
func (_TokenContract *TokenContractTransactor) SetNewAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "setNewAdmin", admin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address admin) returns()
func (_TokenContract *TokenContractSession) SetNewAdmin(admin common.Address) (*types.Transaction, error) {
	return _TokenContract.Contract.SetNewAdmin(&_TokenContract.TransactOpts, admin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address admin) returns()
func (_TokenContract *TokenContractTransactorSession) SetNewAdmin(admin common.Address) (*types.Transaction, error) {
	return _TokenContract.Contract.SetNewAdmin(&_TokenContract.TransactOpts, admin)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.TransferFrom(&_TokenContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.TransferFrom(&_TokenContract.TransactOpts, from, to, tokenId)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactor) TransferToken(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.contract.Transact(opts, "transferToken", from, to, tokenId)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractSession) TransferToken(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.TransferToken(&_TokenContract.TransactOpts, from, to, tokenId)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address from, address to, uint256 tokenId) returns()
func (_TokenContract *TokenContractTransactorSession) TransferToken(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenContract.Contract.TransferToken(&_TokenContract.TransactOpts, from, to, tokenId)
}

// TokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenContract contract.
type TokenContractApprovalIterator struct {
	Event *TokenContractApproval // Event containing the contract specifics and raw log

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
func (it *TokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractApproval)
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
		it.Event = new(TokenContractApproval)
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
func (it *TokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractApproval represents a Approval event raised by the TokenContract contract.
type TokenContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenContractApprovalIterator{contract: _TokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractApproval)
				if err := _TokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) ParseApproval(log types.Log) (*TokenContractApproval, error) {
	event := new(TokenContractApproval)
	if err := _TokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TokenContract contract.
type TokenContractApprovalForAllIterator struct {
	Event *TokenContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractApprovalForAll)
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
		it.Event = new(TokenContractApprovalForAll)
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
func (it *TokenContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractApprovalForAll represents a ApprovalForAll event raised by the TokenContract contract.
type TokenContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenContract *TokenContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenContractApprovalForAllIterator{contract: _TokenContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenContract *TokenContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractApprovalForAll)
				if err := _TokenContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenContract *TokenContractFilterer) ParseApprovalForAll(log types.Log) (*TokenContractApprovalForAll, error) {
	event := new(TokenContractApprovalForAll)
	if err := _TokenContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenContract contract.
type TokenContractInitializedIterator struct {
	Event *TokenContractInitialized // Event containing the contract specifics and raw log

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
func (it *TokenContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractInitialized)
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
		it.Event = new(TokenContractInitialized)
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
func (it *TokenContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractInitialized represents a Initialized event raised by the TokenContract contract.
type TokenContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenContract *TokenContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenContractInitializedIterator, error) {

	logs, sub, err := _TokenContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenContractInitializedIterator{contract: _TokenContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenContract *TokenContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenContractInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractInitialized)
				if err := _TokenContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenContract *TokenContractFilterer) ParseInitialized(log types.Log) (*TokenContractInitialized, error) {
	event := new(TokenContractInitialized)
	if err := _TokenContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractSuccessfullyMintedIterator is returned from FilterSuccessfullyMinted and is used to iterate over the raw logs and unpacked data for SuccessfullyMinted events raised by the TokenContract contract.
type TokenContractSuccessfullyMintedIterator struct {
	Event *TokenContractSuccessfullyMinted // Event containing the contract specifics and raw log

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
func (it *TokenContractSuccessfullyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractSuccessfullyMinted)
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
		it.Event = new(TokenContractSuccessfullyMinted)
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
func (it *TokenContractSuccessfullyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractSuccessfullyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractSuccessfullyMinted represents a SuccessfullyMinted event raised by the TokenContract contract.
type TokenContractSuccessfullyMinted struct {
	Recipient common.Address
	TokenId   *big.Int
	TokenURI  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSuccessfullyMinted is a free log retrieval operation binding the contract event 0x2dde4817c8eaaee29c7f7bd4d03af4c70f3dc6832746afb9da69d84bd0560f63.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_TokenContract *TokenContractFilterer) FilterSuccessfullyMinted(opts *bind.FilterOpts, recipient []common.Address) (*TokenContractSuccessfullyMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenContract.contract.FilterLogs(opts, "SuccessfullyMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenContractSuccessfullyMintedIterator{contract: _TokenContract.contract, event: "SuccessfullyMinted", logs: logs, sub: sub}, nil
}

// WatchSuccessfullyMinted is a free log subscription operation binding the contract event 0x2dde4817c8eaaee29c7f7bd4d03af4c70f3dc6832746afb9da69d84bd0560f63.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_TokenContract *TokenContractFilterer) WatchSuccessfullyMinted(opts *bind.WatchOpts, sink chan<- *TokenContractSuccessfullyMinted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenContract.contract.WatchLogs(opts, "SuccessfullyMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractSuccessfullyMinted)
				if err := _TokenContract.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
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

// ParseSuccessfullyMinted is a log parse operation binding the contract event 0x2dde4817c8eaaee29c7f7bd4d03af4c70f3dc6832746afb9da69d84bd0560f63.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_TokenContract *TokenContractFilterer) ParseSuccessfullyMinted(log types.Log) (*TokenContractSuccessfullyMinted, error) {
	event := new(TokenContractSuccessfullyMinted)
	if err := _TokenContract.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenContract contract.
type TokenContractTransferIterator struct {
	Event *TokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *TokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractTransfer)
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
		it.Event = new(TokenContractTransfer)
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
func (it *TokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractTransfer represents a Transfer event raised by the TokenContract contract.
type TokenContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenContractTransferIterator{contract: _TokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractTransfer)
				if err := _TokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenContract *TokenContractFilterer) ParseTransfer(log types.Log) (*TokenContractTransfer, error) {
	event := new(TokenContractTransfer)
	if err := _TokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

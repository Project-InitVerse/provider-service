// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package util

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProviderABI is the input ABI used to generate the binding from.
const ProviderABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider_info\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"ChallengeStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MarginAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MarginWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ProviderResourceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"StateChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"new_info\",\"type\":\"string\"}],\"name\":\"changeProviderInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_new_region\",\"type\":\"string\"}],\"name\":\"changeRegion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consume_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consume_mem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consume_storage\",\"type\":\"uint256\"}],\"name\":\"consumeResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDetail\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"used\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"lock\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"challenge\",\"type\":\"bool\"},{\"internalType\":\"enumProviderState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"last_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_margin_time\",\"type\":\"uint256\"}],\"internalType\":\"structproviderInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLeftResource\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalResource\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_challenge_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_margin_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_punish_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"margin_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider_first_margin_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_start_margin_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_start_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consumed_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consumed_mem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consumed_storage\",\"type\":\"uint256\"}],\"name\":\"recoverResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"name\":\"reduceResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"region\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePunish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"whether_start\",\"type\":\"bool\"}],\"name\":\"startChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumProviderState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_sto_count\",\"type\":\"uint256\"}],\"name\":\"updateResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"used\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Provider is an auto generated Go binding around an Ethereum contract.
type Provider struct {
	ProviderCaller     // Read-only binding to the contract
	ProviderTransactor // Write-only binding to the contract
	ProviderFilterer   // Log filterer for contract events
}

// ProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderSession struct {
	Contract     *Provider         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderCallerSession struct {
	Contract *ProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderTransactorSession struct {
	Contract     *ProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRaw struct {
	Contract *Provider // Generic contract binding to access the raw methods on
}

// ProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderCallerRaw struct {
	Contract *ProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderTransactorRaw struct {
	Contract *ProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvider creates a new instance of Provider, bound to a specific deployed contract.
func NewProvider(address common.Address, backend bind.ContractBackend) (*Provider, error) {
	contract, err := bindProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Provider{ProviderCaller: ProviderCaller{contract: contract}, ProviderTransactor: ProviderTransactor{contract: contract}, ProviderFilterer: ProviderFilterer{contract: contract}}, nil
}

// NewProviderCaller creates a new read-only instance of Provider, bound to a specific deployed contract.
func NewProviderCaller(address common.Address, caller bind.ContractCaller) (*ProviderCaller, error) {
	contract, err := bindProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderCaller{contract: contract}, nil
}

// NewProviderTransactor creates a new write-only instance of Provider, bound to a specific deployed contract.
func NewProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderTransactor, error) {
	contract, err := bindProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderTransactor{contract: contract}, nil
}

// NewProviderFilterer creates a new log filterer instance of Provider, bound to a specific deployed contract.
func NewProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderFilterer, error) {
	contract, err := bindProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderFilterer{contract: contract}, nil
}

// bindProvider binds a generic wrapper to an already deployed contract.
func bindProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.ProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transact(opts, method, params...)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(bool)
func (_Provider *ProviderCaller) Challenge(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(bool)
func (_Provider *ProviderSession) Challenge() (bool, error) {
	return _Provider.Contract.Challenge(&_Provider.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(bool)
func (_Provider *ProviderCallerSession) Challenge() (bool, error) {
	return _Provider.Contract.Challenge(&_Provider.CallOpts)
}

// GetDetail is a free data retrieval call binding the contract method 0x5b2a4cff.
//
// Solidity: function getDetail() view returns(((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256))
func (_Provider *ProviderCaller) GetDetail(opts *bind.CallOpts) (providerInfo, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getDetail")

	if err != nil {
		return *new(providerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(providerInfo)).(*providerInfo)

	return out0, err

}

// GetDetail is a free data retrieval call binding the contract method 0x5b2a4cff.
//
// Solidity: function getDetail() view returns(((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256))
func (_Provider *ProviderSession) GetDetail() (providerInfo, error) {
	return _Provider.Contract.GetDetail(&_Provider.CallOpts)
}

// GetDetail is a free data retrieval call binding the contract method 0x5b2a4cff.
//
// Solidity: function getDetail() view returns(((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256))
func (_Provider *ProviderCallerSession) GetDetail() (providerInfo, error) {
	return _Provider.Contract.GetDetail(&_Provider.CallOpts)
}

// GetLeftResource is a free data retrieval call binding the contract method 0x78bd8d8b.
//
// Solidity: function getLeftResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderCaller) GetLeftResource(opts *bind.CallOpts) (poaResource, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getLeftResource")

	if err != nil {
		return *new(poaResource), err
	}

	out0 := *abi.ConvertType(out[0], new(poaResource)).(*poaResource)

	return out0, err

}

// GetLeftResource is a free data retrieval call binding the contract method 0x78bd8d8b.
//
// Solidity: function getLeftResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderSession) GetLeftResource() (poaResource, error) {
	return _Provider.Contract.GetLeftResource(&_Provider.CallOpts)
}

// GetLeftResource is a free data retrieval call binding the contract method 0x78bd8d8b.
//
// Solidity: function getLeftResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderCallerSession) GetLeftResource() (poaResource, error) {
	return _Provider.Contract.GetLeftResource(&_Provider.CallOpts)
}

// GetTotalResource is a free data retrieval call binding the contract method 0x41278c3c.
//
// Solidity: function getTotalResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderCaller) GetTotalResource(opts *bind.CallOpts) (poaResource, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getTotalResource")

	if err != nil {
		return *new(poaResource), err
	}

	out0 := *abi.ConvertType(out[0], new(poaResource)).(*poaResource)

	return out0, err

}

// GetTotalResource is a free data retrieval call binding the contract method 0x41278c3c.
//
// Solidity: function getTotalResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderSession) GetTotalResource() (poaResource, error) {
	return _Provider.Contract.GetTotalResource(&_Provider.CallOpts)
}

// GetTotalResource is a free data retrieval call binding the contract method 0x41278c3c.
//
// Solidity: function getTotalResource() view returns((uint256,uint256,uint256))
func (_Provider *ProviderCallerSession) GetTotalResource() (poaResource, error) {
	return _Provider.Contract.GetTotalResource(&_Provider.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(string)
func (_Provider *ProviderCaller) Info(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "info")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(string)
func (_Provider *ProviderSession) Info() (string, error) {
	return _Provider.Contract.Info(&_Provider.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(string)
func (_Provider *ProviderCallerSession) Info() (string, error) {
	return _Provider.Contract.Info(&_Provider.CallOpts)
}

// LastChallengeTime is a free data retrieval call binding the contract method 0xc827adda.
//
// Solidity: function last_challenge_time() view returns(uint256)
func (_Provider *ProviderCaller) LastChallengeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "last_challenge_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastChallengeTime is a free data retrieval call binding the contract method 0xc827adda.
//
// Solidity: function last_challenge_time() view returns(uint256)
func (_Provider *ProviderSession) LastChallengeTime() (*big.Int, error) {
	return _Provider.Contract.LastChallengeTime(&_Provider.CallOpts)
}

// LastChallengeTime is a free data retrieval call binding the contract method 0xc827adda.
//
// Solidity: function last_challenge_time() view returns(uint256)
func (_Provider *ProviderCallerSession) LastChallengeTime() (*big.Int, error) {
	return _Provider.Contract.LastChallengeTime(&_Provider.CallOpts)
}

// LastMarginTime is a free data retrieval call binding the contract method 0x331cdced.
//
// Solidity: function last_margin_time() view returns(uint256)
func (_Provider *ProviderCaller) LastMarginTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "last_margin_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMarginTime is a free data retrieval call binding the contract method 0x331cdced.
//
// Solidity: function last_margin_time() view returns(uint256)
func (_Provider *ProviderSession) LastMarginTime() (*big.Int, error) {
	return _Provider.Contract.LastMarginTime(&_Provider.CallOpts)
}

// LastMarginTime is a free data retrieval call binding the contract method 0x331cdced.
//
// Solidity: function last_margin_time() view returns(uint256)
func (_Provider *ProviderCallerSession) LastMarginTime() (*big.Int, error) {
	return _Provider.Contract.LastMarginTime(&_Provider.CallOpts)
}

// LastPunishTime is a free data retrieval call binding the contract method 0xf2c5a2d5.
//
// Solidity: function last_punish_time() view returns(uint256)
func (_Provider *ProviderCaller) LastPunishTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "last_punish_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPunishTime is a free data retrieval call binding the contract method 0xf2c5a2d5.
//
// Solidity: function last_punish_time() view returns(uint256)
func (_Provider *ProviderSession) LastPunishTime() (*big.Int, error) {
	return _Provider.Contract.LastPunishTime(&_Provider.CallOpts)
}

// LastPunishTime is a free data retrieval call binding the contract method 0xf2c5a2d5.
//
// Solidity: function last_punish_time() view returns(uint256)
func (_Provider *ProviderCallerSession) LastPunishTime() (*big.Int, error) {
	return _Provider.Contract.LastPunishTime(&_Provider.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCaller) Lock(opts *bind.CallOpts) (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "lock")

	outstruct := new(struct {
		CpuCount     *big.Int
		MemoryCount  *big.Int
		StorageCount *big.Int
	})

	outstruct.CpuCount = out[0].(*big.Int)
	outstruct.MemoryCount = out[1].(*big.Int)
	outstruct.StorageCount = out[2].(*big.Int)

	return *outstruct, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderSession) Lock() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Lock(&_Provider.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCallerSession) Lock() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Lock(&_Provider.CallOpts)
}

// MarginBlock is a free data retrieval call binding the contract method 0x2b9ea20d.
//
// Solidity: function margin_block() view returns(uint256)
func (_Provider *ProviderCaller) MarginBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "margin_block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginBlock is a free data retrieval call binding the contract method 0x2b9ea20d.
//
// Solidity: function margin_block() view returns(uint256)
func (_Provider *ProviderSession) MarginBlock() (*big.Int, error) {
	return _Provider.Contract.MarginBlock(&_Provider.CallOpts)
}

// MarginBlock is a free data retrieval call binding the contract method 0x2b9ea20d.
//
// Solidity: function margin_block() view returns(uint256)
func (_Provider *ProviderCallerSession) MarginBlock() (*big.Int, error) {
	return _Provider.Contract.MarginBlock(&_Provider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Provider *ProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Provider *ProviderSession) Owner() (common.Address, error) {
	return _Provider.Contract.Owner(&_Provider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Provider *ProviderCallerSession) Owner() (common.Address, error) {
	return _Provider.Contract.Owner(&_Provider.CallOpts)
}

// ProviderFirstMarginTime is a free data retrieval call binding the contract method 0x6890c245.
//
// Solidity: function provider_first_margin_time() view returns(uint256)
func (_Provider *ProviderCaller) ProviderFirstMarginTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "provider_first_margin_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderFirstMarginTime is a free data retrieval call binding the contract method 0x6890c245.
//
// Solidity: function provider_first_margin_time() view returns(uint256)
func (_Provider *ProviderSession) ProviderFirstMarginTime() (*big.Int, error) {
	return _Provider.Contract.ProviderFirstMarginTime(&_Provider.CallOpts)
}

// ProviderFirstMarginTime is a free data retrieval call binding the contract method 0x6890c245.
//
// Solidity: function provider_first_margin_time() view returns(uint256)
func (_Provider *ProviderCallerSession) ProviderFirstMarginTime() (*big.Int, error) {
	return _Provider.Contract.ProviderFirstMarginTime(&_Provider.CallOpts)
}

// PunishStartMarginAmount is a free data retrieval call binding the contract method 0x1a1e4d75.
//
// Solidity: function punish_start_margin_amount() view returns(uint256)
func (_Provider *ProviderCaller) PunishStartMarginAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "punish_start_margin_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishStartMarginAmount is a free data retrieval call binding the contract method 0x1a1e4d75.
//
// Solidity: function punish_start_margin_amount() view returns(uint256)
func (_Provider *ProviderSession) PunishStartMarginAmount() (*big.Int, error) {
	return _Provider.Contract.PunishStartMarginAmount(&_Provider.CallOpts)
}

// PunishStartMarginAmount is a free data retrieval call binding the contract method 0x1a1e4d75.
//
// Solidity: function punish_start_margin_amount() view returns(uint256)
func (_Provider *ProviderCallerSession) PunishStartMarginAmount() (*big.Int, error) {
	return _Provider.Contract.PunishStartMarginAmount(&_Provider.CallOpts)
}

// PunishStartTime is a free data retrieval call binding the contract method 0x8405ff46.
//
// Solidity: function punish_start_time() view returns(uint256)
func (_Provider *ProviderCaller) PunishStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "punish_start_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishStartTime is a free data retrieval call binding the contract method 0x8405ff46.
//
// Solidity: function punish_start_time() view returns(uint256)
func (_Provider *ProviderSession) PunishStartTime() (*big.Int, error) {
	return _Provider.Contract.PunishStartTime(&_Provider.CallOpts)
}

// PunishStartTime is a free data retrieval call binding the contract method 0x8405ff46.
//
// Solidity: function punish_start_time() view returns(uint256)
func (_Provider *ProviderCallerSession) PunishStartTime() (*big.Int, error) {
	return _Provider.Contract.PunishStartTime(&_Provider.CallOpts)
}

// Region is a free data retrieval call binding the contract method 0x5062d68d.
//
// Solidity: function region() view returns(string)
func (_Provider *ProviderCaller) Region(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "region")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Region is a free data retrieval call binding the contract method 0x5062d68d.
//
// Solidity: function region() view returns(string)
func (_Provider *ProviderSession) Region() (string, error) {
	return _Provider.Contract.Region(&_Provider.CallOpts)
}

// Region is a free data retrieval call binding the contract method 0x5062d68d.
//
// Solidity: function region() view returns(string)
func (_Provider *ProviderCallerSession) Region() (string, error) {
	return _Provider.Contract.Region(&_Provider.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Provider *ProviderCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Provider *ProviderSession) State() (uint8, error) {
	return _Provider.Contract.State(&_Provider.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Provider *ProviderCallerSession) State() (uint8, error) {
	return _Provider.Contract.State(&_Provider.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCaller) Total(opts *bind.CallOpts) (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "total")

	outstruct := new(struct {
		CpuCount     *big.Int
		MemoryCount  *big.Int
		StorageCount *big.Int
	})

	outstruct.CpuCount = out[0].(*big.Int)
	outstruct.MemoryCount = out[1].(*big.Int)
	outstruct.StorageCount = out[2].(*big.Int)

	return *outstruct, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderSession) Total() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Total(&_Provider.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCallerSession) Total() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Total(&_Provider.CallOpts)
}

// Used is a free data retrieval call binding the contract method 0x63898e2b.
//
// Solidity: function used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCaller) Used(opts *bind.CallOpts) (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "used")

	outstruct := new(struct {
		CpuCount     *big.Int
		MemoryCount  *big.Int
		StorageCount *big.Int
	})

	outstruct.CpuCount = out[0].(*big.Int)
	outstruct.MemoryCount = out[1].(*big.Int)
	outstruct.StorageCount = out[2].(*big.Int)

	return *outstruct, err

}

// Used is a free data retrieval call binding the contract method 0x63898e2b.
//
// Solidity: function used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderSession) Used() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Used(&_Provider.CallOpts)
}

// Used is a free data retrieval call binding the contract method 0x63898e2b.
//
// Solidity: function used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_Provider *ProviderCallerSession) Used() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _Provider.Contract.Used(&_Provider.CallOpts)
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0xaefe36a7.
//
// Solidity: function challengeProvider() returns()
func (_Provider *ProviderTransactor) ChallengeProvider(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "challengeProvider")
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0xaefe36a7.
//
// Solidity: function challengeProvider() returns()
func (_Provider *ProviderSession) ChallengeProvider() (*types.Transaction, error) {
	return _Provider.Contract.ChallengeProvider(&_Provider.TransactOpts)
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0xaefe36a7.
//
// Solidity: function challengeProvider() returns()
func (_Provider *ProviderTransactorSession) ChallengeProvider() (*types.Transaction, error) {
	return _Provider.Contract.ChallengeProvider(&_Provider.TransactOpts)
}

// ChangeProviderInfo is a paid mutator transaction binding the contract method 0x56fdcabe.
//
// Solidity: function changeProviderInfo(string new_info) returns()
func (_Provider *ProviderTransactor) ChangeProviderInfo(opts *bind.TransactOpts, new_info string) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "changeProviderInfo", new_info)
}

// ChangeProviderInfo is a paid mutator transaction binding the contract method 0x56fdcabe.
//
// Solidity: function changeProviderInfo(string new_info) returns()
func (_Provider *ProviderSession) ChangeProviderInfo(new_info string) (*types.Transaction, error) {
	return _Provider.Contract.ChangeProviderInfo(&_Provider.TransactOpts, new_info)
}

// ChangeProviderInfo is a paid mutator transaction binding the contract method 0x56fdcabe.
//
// Solidity: function changeProviderInfo(string new_info) returns()
func (_Provider *ProviderTransactorSession) ChangeProviderInfo(new_info string) (*types.Transaction, error) {
	return _Provider.Contract.ChangeProviderInfo(&_Provider.TransactOpts, new_info)
}

// ChangeRegion is a paid mutator transaction binding the contract method 0x89908892.
//
// Solidity: function changeRegion(string _new_region) returns()
func (_Provider *ProviderTransactor) ChangeRegion(opts *bind.TransactOpts, _new_region string) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "changeRegion", _new_region)
}

// ChangeRegion is a paid mutator transaction binding the contract method 0x89908892.
//
// Solidity: function changeRegion(string _new_region) returns()
func (_Provider *ProviderSession) ChangeRegion(_new_region string) (*types.Transaction, error) {
	return _Provider.Contract.ChangeRegion(&_Provider.TransactOpts, _new_region)
}

// ChangeRegion is a paid mutator transaction binding the contract method 0x89908892.
//
// Solidity: function changeRegion(string _new_region) returns()
func (_Provider *ProviderTransactorSession) ChangeRegion(_new_region string) (*types.Transaction, error) {
	return _Provider.Contract.ChangeRegion(&_Provider.TransactOpts, _new_region)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0x83dcd2ec.
//
// Solidity: function consumeResource(uint256 consume_cpu, uint256 consume_mem, uint256 consume_storage) returns()
func (_Provider *ProviderTransactor) ConsumeResource(opts *bind.TransactOpts, consume_cpu *big.Int, consume_mem *big.Int, consume_storage *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "consumeResource", consume_cpu, consume_mem, consume_storage)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0x83dcd2ec.
//
// Solidity: function consumeResource(uint256 consume_cpu, uint256 consume_mem, uint256 consume_storage) returns()
func (_Provider *ProviderSession) ConsumeResource(consume_cpu *big.Int, consume_mem *big.Int, consume_storage *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.ConsumeResource(&_Provider.TransactOpts, consume_cpu, consume_mem, consume_storage)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0x83dcd2ec.
//
// Solidity: function consumeResource(uint256 consume_cpu, uint256 consume_mem, uint256 consume_storage) returns()
func (_Provider *ProviderTransactorSession) ConsumeResource(consume_cpu *big.Int, consume_mem *big.Int, consume_storage *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.ConsumeResource(&_Provider.TransactOpts, consume_cpu, consume_mem, consume_storage)
}

// Punish is a paid mutator transaction binding the contract method 0x826d3dec.
//
// Solidity: function punish() returns()
func (_Provider *ProviderTransactor) Punish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "punish")
}

// Punish is a paid mutator transaction binding the contract method 0x826d3dec.
//
// Solidity: function punish() returns()
func (_Provider *ProviderSession) Punish() (*types.Transaction, error) {
	return _Provider.Contract.Punish(&_Provider.TransactOpts)
}

// Punish is a paid mutator transaction binding the contract method 0x826d3dec.
//
// Solidity: function punish() returns()
func (_Provider *ProviderTransactorSession) Punish() (*types.Transaction, error) {
	return _Provider.Contract.Punish(&_Provider.TransactOpts)
}

// RecoverResource is a paid mutator transaction binding the contract method 0xacaab27c.
//
// Solidity: function recoverResource(uint256 consumed_cpu, uint256 consumed_mem, uint256 consumed_storage) returns()
func (_Provider *ProviderTransactor) RecoverResource(opts *bind.TransactOpts, consumed_cpu *big.Int, consumed_mem *big.Int, consumed_storage *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "recoverResource", consumed_cpu, consumed_mem, consumed_storage)
}

// RecoverResource is a paid mutator transaction binding the contract method 0xacaab27c.
//
// Solidity: function recoverResource(uint256 consumed_cpu, uint256 consumed_mem, uint256 consumed_storage) returns()
func (_Provider *ProviderSession) RecoverResource(consumed_cpu *big.Int, consumed_mem *big.Int, consumed_storage *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.RecoverResource(&_Provider.TransactOpts, consumed_cpu, consumed_mem, consumed_storage)
}

// RecoverResource is a paid mutator transaction binding the contract method 0xacaab27c.
//
// Solidity: function recoverResource(uint256 consumed_cpu, uint256 consumed_mem, uint256 consumed_storage) returns()
func (_Provider *ProviderTransactorSession) RecoverResource(consumed_cpu *big.Int, consumed_mem *big.Int, consumed_storage *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.RecoverResource(&_Provider.TransactOpts, consumed_cpu, consumed_mem, consumed_storage)
}

// ReduceResource is a paid mutator transaction binding the contract method 0x5245b098.
//
// Solidity: function reduceResource(uint256 cpu_count, uint256 memory_count, uint256 storage_count) returns()
func (_Provider *ProviderTransactor) ReduceResource(opts *bind.TransactOpts, cpu_count *big.Int, memory_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "reduceResource", cpu_count, memory_count, storage_count)
}

// ReduceResource is a paid mutator transaction binding the contract method 0x5245b098.
//
// Solidity: function reduceResource(uint256 cpu_count, uint256 memory_count, uint256 storage_count) returns()
func (_Provider *ProviderSession) ReduceResource(cpu_count *big.Int, memory_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.ReduceResource(&_Provider.TransactOpts, cpu_count, memory_count, storage_count)
}

// ReduceResource is a paid mutator transaction binding the contract method 0x5245b098.
//
// Solidity: function reduceResource(uint256 cpu_count, uint256 memory_count, uint256 storage_count) returns()
func (_Provider *ProviderTransactorSession) ReduceResource(cpu_count *big.Int, memory_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.ReduceResource(&_Provider.TransactOpts, cpu_count, memory_count, storage_count)
}

// RemovePunish is a paid mutator transaction binding the contract method 0x2b65c333.
//
// Solidity: function removePunish() returns()
func (_Provider *ProviderTransactor) RemovePunish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "removePunish")
}

// RemovePunish is a paid mutator transaction binding the contract method 0x2b65c333.
//
// Solidity: function removePunish() returns()
func (_Provider *ProviderSession) RemovePunish() (*types.Transaction, error) {
	return _Provider.Contract.RemovePunish(&_Provider.TransactOpts)
}

// RemovePunish is a paid mutator transaction binding the contract method 0x2b65c333.
//
// Solidity: function removePunish() returns()
func (_Provider *ProviderTransactorSession) RemovePunish() (*types.Transaction, error) {
	return _Provider.Contract.RemovePunish(&_Provider.TransactOpts)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x0ba1b7f8.
//
// Solidity: function startChallenge(bool whether_start) returns()
func (_Provider *ProviderTransactor) StartChallenge(opts *bind.TransactOpts, whether_start bool) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "startChallenge", whether_start)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x0ba1b7f8.
//
// Solidity: function startChallenge(bool whether_start) returns()
func (_Provider *ProviderSession) StartChallenge(whether_start bool) (*types.Transaction, error) {
	return _Provider.Contract.StartChallenge(&_Provider.TransactOpts, whether_start)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x0ba1b7f8.
//
// Solidity: function startChallenge(bool whether_start) returns()
func (_Provider *ProviderTransactorSession) StartChallenge(whether_start bool) (*types.Transaction, error) {
	return _Provider.Contract.StartChallenge(&_Provider.TransactOpts, whether_start)
}

// UpdateResource is a paid mutator transaction binding the contract method 0x1887349e.
//
// Solidity: function updateResource(uint256 new_cpu_count, uint256 new_mem_count, uint256 new_sto_count) returns()
func (_Provider *ProviderTransactor) UpdateResource(opts *bind.TransactOpts, new_cpu_count *big.Int, new_mem_count *big.Int, new_sto_count *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "updateResource", new_cpu_count, new_mem_count, new_sto_count)
}

// UpdateResource is a paid mutator transaction binding the contract method 0x1887349e.
//
// Solidity: function updateResource(uint256 new_cpu_count, uint256 new_mem_count, uint256 new_sto_count) returns()
func (_Provider *ProviderSession) UpdateResource(new_cpu_count *big.Int, new_mem_count *big.Int, new_sto_count *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.UpdateResource(&_Provider.TransactOpts, new_cpu_count, new_mem_count, new_sto_count)
}

// UpdateResource is a paid mutator transaction binding the contract method 0x1887349e.
//
// Solidity: function updateResource(uint256 new_cpu_count, uint256 new_mem_count, uint256 new_sto_count) returns()
func (_Provider *ProviderTransactorSession) UpdateResource(new_cpu_count *big.Int, new_mem_count *big.Int, new_sto_count *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.UpdateResource(&_Provider.TransactOpts, new_cpu_count, new_mem_count, new_sto_count)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_Provider *ProviderTransactor) WithdrawMargin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "withdrawMargin")
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_Provider *ProviderSession) WithdrawMargin() (*types.Transaction, error) {
	return _Provider.Contract.WithdrawMargin(&_Provider.TransactOpts)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_Provider *ProviderTransactorSession) WithdrawMargin() (*types.Transaction, error) {
	return _Provider.Contract.WithdrawMargin(&_Provider.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Provider *ProviderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Provider *ProviderSession) Receive() (*types.Transaction, error) {
	return _Provider.Contract.Receive(&_Provider.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Provider *ProviderTransactorSession) Receive() (*types.Transaction, error) {
	return _Provider.Contract.Receive(&_Provider.TransactOpts)
}

// ProviderChallengeStateChangeIterator is returned from FilterChallengeStateChange and is used to iterate over the raw logs and unpacked data for ChallengeStateChange events raised by the Provider contract.
type ProviderChallengeStateChangeIterator struct {
	Event *ProviderChallengeStateChange // Event containing the contract specifics and raw log

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
func (it *ProviderChallengeStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderChallengeStateChange)
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
		it.Event = new(ProviderChallengeStateChange)
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
func (it *ProviderChallengeStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderChallengeStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderChallengeStateChange represents a ChallengeStateChange event raised by the Provider contract.
type ProviderChallengeStateChange struct {
	Arg0 common.Address
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChallengeStateChange is a free log retrieval operation binding the contract event 0xf1bb76c4f7fb203057fc14be90e480ac15948e4d933fa29203ac92ae78e51dc0.
//
// Solidity: event ChallengeStateChange(address indexed arg0, bool arg1)
func (_Provider *ProviderFilterer) FilterChallengeStateChange(opts *bind.FilterOpts, arg0 []common.Address) (*ProviderChallengeStateChangeIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "ChallengeStateChange", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &ProviderChallengeStateChangeIterator{contract: _Provider.contract, event: "ChallengeStateChange", logs: logs, sub: sub}, nil
}

// WatchChallengeStateChange is a free log subscription operation binding the contract event 0xf1bb76c4f7fb203057fc14be90e480ac15948e4d933fa29203ac92ae78e51dc0.
//
// Solidity: event ChallengeStateChange(address indexed arg0, bool arg1)
func (_Provider *ProviderFilterer) WatchChallengeStateChange(opts *bind.WatchOpts, sink chan<- *ProviderChallengeStateChange, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "ChallengeStateChange", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderChallengeStateChange)
				if err := _Provider.contract.UnpackLog(event, "ChallengeStateChange", log); err != nil {
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

// ParseChallengeStateChange is a log parse operation binding the contract event 0xf1bb76c4f7fb203057fc14be90e480ac15948e4d933fa29203ac92ae78e51dc0.
//
// Solidity: event ChallengeStateChange(address indexed arg0, bool arg1)
func (_Provider *ProviderFilterer) ParseChallengeStateChange(log types.Log) (*ProviderChallengeStateChange, error) {
	event := new(ProviderChallengeStateChange)
	if err := _Provider.contract.UnpackLog(event, "ChallengeStateChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProviderMarginAddIterator is returned from FilterMarginAdd and is used to iterate over the raw logs and unpacked data for MarginAdd events raised by the Provider contract.
type ProviderMarginAddIterator struct {
	Event *ProviderMarginAdd // Event containing the contract specifics and raw log

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
func (it *ProviderMarginAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderMarginAdd)
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
		it.Event = new(ProviderMarginAdd)
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
func (it *ProviderMarginAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderMarginAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderMarginAdd represents a MarginAdd event raised by the Provider contract.
type ProviderMarginAdd struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMarginAdd is a free log retrieval operation binding the contract event 0x61d0b7be5d309e6e638c81efd0b82c49c48c9b716313b189f5ef0f08a985497a.
//
// Solidity: event MarginAdd(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) FilterMarginAdd(opts *bind.FilterOpts, arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (*ProviderMarginAddIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "MarginAdd", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return &ProviderMarginAddIterator{contract: _Provider.contract, event: "MarginAdd", logs: logs, sub: sub}, nil
}

// WatchMarginAdd is a free log subscription operation binding the contract event 0x61d0b7be5d309e6e638c81efd0b82c49c48c9b716313b189f5ef0f08a985497a.
//
// Solidity: event MarginAdd(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) WatchMarginAdd(opts *bind.WatchOpts, sink chan<- *ProviderMarginAdd, arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "MarginAdd", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderMarginAdd)
				if err := _Provider.contract.UnpackLog(event, "MarginAdd", log); err != nil {
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

// ParseMarginAdd is a log parse operation binding the contract event 0x61d0b7be5d309e6e638c81efd0b82c49c48c9b716313b189f5ef0f08a985497a.
//
// Solidity: event MarginAdd(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) ParseMarginAdd(log types.Log) (*ProviderMarginAdd, error) {
	event := new(ProviderMarginAdd)
	if err := _Provider.contract.UnpackLog(event, "MarginAdd", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProviderMarginWithdrawIterator is returned from FilterMarginWithdraw and is used to iterate over the raw logs and unpacked data for MarginWithdraw events raised by the Provider contract.
type ProviderMarginWithdrawIterator struct {
	Event *ProviderMarginWithdraw // Event containing the contract specifics and raw log

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
func (it *ProviderMarginWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderMarginWithdraw)
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
		it.Event = new(ProviderMarginWithdraw)
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
func (it *ProviderMarginWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderMarginWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderMarginWithdraw represents a MarginWithdraw event raised by the Provider contract.
type ProviderMarginWithdraw struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMarginWithdraw is a free log retrieval operation binding the contract event 0x14259ba2b8b85c65d4fc9ed09cc282307ed6011350fd85a025d81cf92eb4d412.
//
// Solidity: event MarginWithdraw(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) FilterMarginWithdraw(opts *bind.FilterOpts, arg0 []common.Address, arg1 []*big.Int) (*ProviderMarginWithdrawIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "MarginWithdraw", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &ProviderMarginWithdrawIterator{contract: _Provider.contract, event: "MarginWithdraw", logs: logs, sub: sub}, nil
}

// WatchMarginWithdraw is a free log subscription operation binding the contract event 0x14259ba2b8b85c65d4fc9ed09cc282307ed6011350fd85a025d81cf92eb4d412.
//
// Solidity: event MarginWithdraw(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) WatchMarginWithdraw(opts *bind.WatchOpts, sink chan<- *ProviderMarginWithdraw, arg0 []common.Address, arg1 []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "MarginWithdraw", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderMarginWithdraw)
				if err := _Provider.contract.UnpackLog(event, "MarginWithdraw", log); err != nil {
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

// ParseMarginWithdraw is a log parse operation binding the contract event 0x14259ba2b8b85c65d4fc9ed09cc282307ed6011350fd85a025d81cf92eb4d412.
//
// Solidity: event MarginWithdraw(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) ParseMarginWithdraw(log types.Log) (*ProviderMarginWithdraw, error) {
	event := new(ProviderMarginWithdraw)
	if err := _Provider.contract.UnpackLog(event, "MarginWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProviderProviderResourceChangeIterator is returned from FilterProviderResourceChange and is used to iterate over the raw logs and unpacked data for ProviderResourceChange events raised by the Provider contract.
type ProviderProviderResourceChangeIterator struct {
	Event *ProviderProviderResourceChange // Event containing the contract specifics and raw log

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
func (it *ProviderProviderResourceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderProviderResourceChange)
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
		it.Event = new(ProviderProviderResourceChange)
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
func (it *ProviderProviderResourceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderProviderResourceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderProviderResourceChange represents a ProviderResourceChange event raised by the Provider contract.
type ProviderProviderResourceChange struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProviderResourceChange is a free log retrieval operation binding the contract event 0xdc969fe8d93d63d43d2eceb6d701f4e21b6dc3f5f593b568cef29da47021958c.
//
// Solidity: event ProviderResourceChange(address arg0)
func (_Provider *ProviderFilterer) FilterProviderResourceChange(opts *bind.FilterOpts) (*ProviderProviderResourceChangeIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "ProviderResourceChange")
	if err != nil {
		return nil, err
	}
	return &ProviderProviderResourceChangeIterator{contract: _Provider.contract, event: "ProviderResourceChange", logs: logs, sub: sub}, nil
}

// WatchProviderResourceChange is a free log subscription operation binding the contract event 0xdc969fe8d93d63d43d2eceb6d701f4e21b6dc3f5f593b568cef29da47021958c.
//
// Solidity: event ProviderResourceChange(address arg0)
func (_Provider *ProviderFilterer) WatchProviderResourceChange(opts *bind.WatchOpts, sink chan<- *ProviderProviderResourceChange) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "ProviderResourceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderProviderResourceChange)
				if err := _Provider.contract.UnpackLog(event, "ProviderResourceChange", log); err != nil {
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

// ParseProviderResourceChange is a log parse operation binding the contract event 0xdc969fe8d93d63d43d2eceb6d701f4e21b6dc3f5f593b568cef29da47021958c.
//
// Solidity: event ProviderResourceChange(address arg0)
func (_Provider *ProviderFilterer) ParseProviderResourceChange(log types.Log) (*ProviderProviderResourceChange, error) {
	event := new(ProviderProviderResourceChange)
	if err := _Provider.contract.UnpackLog(event, "ProviderResourceChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProviderPunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Provider contract.
type ProviderPunishIterator struct {
	Event *ProviderPunish // Event containing the contract specifics and raw log

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
func (it *ProviderPunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderPunish)
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
		it.Event = new(ProviderPunish)
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
func (it *ProviderPunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderPunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderPunish represents a Punish event raised by the Provider contract.
type ProviderPunish struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0x82b4934446d1f4337c42929ced0362455ec8e8462c799682797acbfa914cf0b2.
//
// Solidity: event Punish(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) FilterPunish(opts *bind.FilterOpts, arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (*ProviderPunishIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Punish", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return &ProviderPunishIterator{contract: _Provider.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0x82b4934446d1f4337c42929ced0362455ec8e8462c799682797acbfa914cf0b2.
//
// Solidity: event Punish(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *ProviderPunish, arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}
	var arg2Rule []interface{}
	for _, arg2Item := range arg2 {
		arg2Rule = append(arg2Rule, arg2Item)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Punish", arg0Rule, arg1Rule, arg2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderPunish)
				if err := _Provider.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0x82b4934446d1f4337c42929ced0362455ec8e8462c799682797acbfa914cf0b2.
//
// Solidity: event Punish(address indexed arg0, uint256 indexed arg1, uint256 indexed arg2)
func (_Provider *ProviderFilterer) ParsePunish(log types.Log) (*ProviderPunish, error) {
	event := new(ProviderPunish)
	if err := _Provider.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProviderStateChangeIterator is returned from FilterStateChange and is used to iterate over the raw logs and unpacked data for StateChange events raised by the Provider contract.
type ProviderStateChangeIterator struct {
	Event *ProviderStateChange // Event containing the contract specifics and raw log

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
func (it *ProviderStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderStateChange)
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
		it.Event = new(ProviderStateChange)
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
func (it *ProviderStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderStateChange represents a StateChange event raised by the Provider contract.
type ProviderStateChange struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStateChange is a free log retrieval operation binding the contract event 0x6b97a7b03d79e34b82e2ee319a00e3ea6ed574939b7031de492b148c29ff9439.
//
// Solidity: event StateChange(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) FilterStateChange(opts *bind.FilterOpts, arg0 []common.Address, arg1 []*big.Int) (*ProviderStateChangeIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "StateChange", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &ProviderStateChangeIterator{contract: _Provider.contract, event: "StateChange", logs: logs, sub: sub}, nil
}

// WatchStateChange is a free log subscription operation binding the contract event 0x6b97a7b03d79e34b82e2ee319a00e3ea6ed574939b7031de492b148c29ff9439.
//
// Solidity: event StateChange(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) WatchStateChange(opts *bind.WatchOpts, sink chan<- *ProviderStateChange, arg0 []common.Address, arg1 []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "StateChange", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderStateChange)
				if err := _Provider.contract.UnpackLog(event, "StateChange", log); err != nil {
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

// ParseStateChange is a log parse operation binding the contract event 0x6b97a7b03d79e34b82e2ee319a00e3ea6ed574939b7031de492b148c29ff9439.
//
// Solidity: event StateChange(address indexed arg0, uint256 indexed arg1)
func (_Provider *ProviderFilterer) ParseStateChange(log types.Log) (*ProviderStateChange, error) {
	event := new(ProviderStateChange)
	if err := _Provider.contract.UnpackLog(event, "StateChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

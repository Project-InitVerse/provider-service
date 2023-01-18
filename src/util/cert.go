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

// CertcertRetInfo is an auto generated low-level Go binding around an user-defined struct.
type CertcertRetInfo struct {
	CreateTime *big.Int
	RemainTime *big.Int
	User       common.Address
	State      uint8
	Cert       string
}

// CertABI is the input ABI used to generate the binding from.
const CertABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"cert_user\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumCertState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cert\",\"type\":\"string\"},{\"internalType\":\"enumCertState\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"changeCertState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"new_cert\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"remain_time\",\"type\":\"uint256\"},{\"internalType\":\"enumCertState\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"createNewCert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllUserCert\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumCertState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cert\",\"type\":\"string\"}],\"internalType\":\"structCert.certRetInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserCert\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumCertState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cert\",\"type\":\"string\"}],\"internalType\":\"structCert.certRetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cert\",\"type\":\"string\"}],\"name\":\"user_cert_state\",\"outputs\":[{\"internalType\":\"enumCertState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	CertCaller     // Read-only binding to the contract
	CertTransactor // Write-only binding to the contract
	CertFilterer   // Log filterer for contract events
}

// CertCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertSession struct {
	Contract     *Cert             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertCallerSession struct {
	Contract *CertCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertTransactorSession struct {
	Contract     *CertTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRaw struct {
	Contract *Cert // Generic contract binding to access the raw methods on
}

// CertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertCallerRaw struct {
	Contract *CertCaller // Generic read-only contract binding to access the raw methods on
}

// CertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertTransactorRaw struct {
	Contract *CertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCert creates a new instance of Cert, bound to a specific deployed contract.
func NewCert(address common.Address, backend bind.ContractBackend) (*Cert, error) {
	contract, err := bindCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// NewCertCaller creates a new read-only instance of Cert, bound to a specific deployed contract.
func NewCertCaller(address common.Address, caller bind.ContractCaller) (*CertCaller, error) {
	contract, err := bindCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertCaller{contract: contract}, nil
}

// NewCertTransactor creates a new write-only instance of Cert, bound to a specific deployed contract.
func NewCertTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTransactor, error) {
	contract, err := bindCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTransactor{contract: contract}, nil
}

// NewCertFilterer creates a new log filterer instance of Cert, bound to a specific deployed contract.
func NewCertFilterer(address common.Address, filterer bind.ContractFilterer) (*CertFilterer, error) {
	contract, err := bindCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertFilterer{contract: contract}, nil
}

// bindCert binds a generic wrapper to an already deployed contract.
func bindCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.CertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
}

// CertUser is a free data retrieval call binding the contract method 0x68010ae4.
//
// Solidity: function cert_user(string ) view returns(uint256 createTime, uint256 remainTime, address user, uint8 state)
func (_Cert *CertCaller) CertUser(opts *bind.CallOpts, arg0 string) (struct {
	CreateTime *big.Int
	RemainTime *big.Int
	User       common.Address
	State      uint8
}, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "cert_user", arg0)

	outstruct := new(struct {
		CreateTime *big.Int
		RemainTime *big.Int
		User       common.Address
		State      uint8
	})

	outstruct.CreateTime = out[0].(*big.Int)
	outstruct.RemainTime = out[1].(*big.Int)
	outstruct.User = out[2].(common.Address)
	outstruct.State = out[3].(uint8)

	return *outstruct, err

}

// CertUser is a free data retrieval call binding the contract method 0x68010ae4.
//
// Solidity: function cert_user(string ) view returns(uint256 createTime, uint256 remainTime, address user, uint8 state)
func (_Cert *CertSession) CertUser(arg0 string) (struct {
	CreateTime *big.Int
	RemainTime *big.Int
	User       common.Address
	State      uint8
}, error) {
	return _Cert.Contract.CertUser(&_Cert.CallOpts, arg0)
}

// CertUser is a free data retrieval call binding the contract method 0x68010ae4.
//
// Solidity: function cert_user(string ) view returns(uint256 createTime, uint256 remainTime, address user, uint8 state)
func (_Cert *CertCallerSession) CertUser(arg0 string) (struct {
	CreateTime *big.Int
	RemainTime *big.Int
	User       common.Address
	State      uint8
}, error) {
	return _Cert.Contract.CertUser(&_Cert.CallOpts, arg0)
}

// GetAllUserCert is a free data retrieval call binding the contract method 0x3a705b16.
//
// Solidity: function getAllUserCert(address user) view returns((uint256,uint256,address,uint8,string)[])
func (_Cert *CertCaller) GetAllUserCert(opts *bind.CallOpts, user common.Address) ([]CertcertRetInfo, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "getAllUserCert", user)

	if err != nil {
		return *new([]CertcertRetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertcertRetInfo)).(*[]CertcertRetInfo)

	return out0, err

}

// GetAllUserCert is a free data retrieval call binding the contract method 0x3a705b16.
//
// Solidity: function getAllUserCert(address user) view returns((uint256,uint256,address,uint8,string)[])
func (_Cert *CertSession) GetAllUserCert(user common.Address) ([]CertcertRetInfo, error) {
	return _Cert.Contract.GetAllUserCert(&_Cert.CallOpts, user)
}

// GetAllUserCert is a free data retrieval call binding the contract method 0x3a705b16.
//
// Solidity: function getAllUserCert(address user) view returns((uint256,uint256,address,uint8,string)[])
func (_Cert *CertCallerSession) GetAllUserCert(user common.Address) ([]CertcertRetInfo, error) {
	return _Cert.Contract.GetAllUserCert(&_Cert.CallOpts, user)
}

// GetUserCert is a free data retrieval call binding the contract method 0x33ed01a7.
//
// Solidity: function getUserCert(address user, uint256 index) view returns((uint256,uint256,address,uint8,string))
func (_Cert *CertCaller) GetUserCert(opts *bind.CallOpts, user common.Address, index *big.Int) (CertcertRetInfo, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "getUserCert", user, index)

	if err != nil {
		return *new(CertcertRetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CertcertRetInfo)).(*CertcertRetInfo)

	return out0, err

}

// GetUserCert is a free data retrieval call binding the contract method 0x33ed01a7.
//
// Solidity: function getUserCert(address user, uint256 index) view returns((uint256,uint256,address,uint8,string))
func (_Cert *CertSession) GetUserCert(user common.Address, index *big.Int) (CertcertRetInfo, error) {
	return _Cert.Contract.GetUserCert(&_Cert.CallOpts, user, index)
}

// GetUserCert is a free data retrieval call binding the contract method 0x33ed01a7.
//
// Solidity: function getUserCert(address user, uint256 index) view returns((uint256,uint256,address,uint8,string))
func (_Cert *CertCallerSession) GetUserCert(user common.Address, index *big.Int) (CertcertRetInfo, error) {
	return _Cert.Contract.GetUserCert(&_Cert.CallOpts, user, index)
}

// UserCertState is a free data retrieval call binding the contract method 0x5add417e.
//
// Solidity: function user_cert_state(address user, string cert) view returns(uint8)
func (_Cert *CertCaller) UserCertState(opts *bind.CallOpts, user common.Address, cert string) (uint8, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "user_cert_state", user, cert)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UserCertState is a free data retrieval call binding the contract method 0x5add417e.
//
// Solidity: function user_cert_state(address user, string cert) view returns(uint8)
func (_Cert *CertSession) UserCertState(user common.Address, cert string) (uint8, error) {
	return _Cert.Contract.UserCertState(&_Cert.CallOpts, user, cert)
}

// UserCertState is a free data retrieval call binding the contract method 0x5add417e.
//
// Solidity: function user_cert_state(address user, string cert) view returns(uint8)
func (_Cert *CertCallerSession) UserCertState(user common.Address, cert string) (uint8, error) {
	return _Cert.Contract.UserCertState(&_Cert.CallOpts, user, cert)
}

// ChangeCertState is a paid mutator transaction binding the contract method 0x1a510b30.
//
// Solidity: function changeCertState(string cert, uint8 _state) returns()
func (_Cert *CertTransactor) ChangeCertState(opts *bind.TransactOpts, cert string, _state uint8) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "changeCertState", cert, _state)
}

// ChangeCertState is a paid mutator transaction binding the contract method 0x1a510b30.
//
// Solidity: function changeCertState(string cert, uint8 _state) returns()
func (_Cert *CertSession) ChangeCertState(cert string, _state uint8) (*types.Transaction, error) {
	return _Cert.Contract.ChangeCertState(&_Cert.TransactOpts, cert, _state)
}

// ChangeCertState is a paid mutator transaction binding the contract method 0x1a510b30.
//
// Solidity: function changeCertState(string cert, uint8 _state) returns()
func (_Cert *CertTransactorSession) ChangeCertState(cert string, _state uint8) (*types.Transaction, error) {
	return _Cert.Contract.ChangeCertState(&_Cert.TransactOpts, cert, _state)
}

// CreateNewCert is a paid mutator transaction binding the contract method 0x3457e6e0.
//
// Solidity: function createNewCert(string new_cert, uint256 remain_time, uint8 _state) returns()
func (_Cert *CertTransactor) CreateNewCert(opts *bind.TransactOpts, new_cert string, remain_time *big.Int, _state uint8) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "createNewCert", new_cert, remain_time, _state)
}

// CreateNewCert is a paid mutator transaction binding the contract method 0x3457e6e0.
//
// Solidity: function createNewCert(string new_cert, uint256 remain_time, uint8 _state) returns()
func (_Cert *CertSession) CreateNewCert(new_cert string, remain_time *big.Int, _state uint8) (*types.Transaction, error) {
	return _Cert.Contract.CreateNewCert(&_Cert.TransactOpts, new_cert, remain_time, _state)
}

// CreateNewCert is a paid mutator transaction binding the contract method 0x3457e6e0.
//
// Solidity: function createNewCert(string new_cert, uint256 remain_time, uint8 _state) returns()
func (_Cert *CertTransactorSession) CreateNewCert(new_cert string, remain_time *big.Int, _state uint8) (*types.Transaction, error) {
	return _Cert.Contract.CreateNewCert(&_Cert.TransactOpts, new_cert, remain_time, _state)
}

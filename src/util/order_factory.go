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

// OrderFactoryABI is the input ABI used to generate the binding from.
const OrderFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"order_addr\",\"type\":\"address\"}],\"name\":\"OrderCreation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cert_center\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orderAddress\",\"type\":\"address\"}],\"name\":\"checkIsOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"m_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"m_storage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"m_cert\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"m_trx_id\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"v_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_storage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cert_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"trx_id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"}],\"name\":\"getProviderAllOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"v_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_storage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cert_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"trx_id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnCompleteOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"v_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_storage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cert_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"trx_id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUserAllOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"v_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_storage\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cert_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"trx_id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_minimum_deposit_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_order_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimum_deposit_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_value\",\"type\":\"uint256\"}],\"name\":\"modify_minimum_deposit_amount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"order_base_map\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cert_center_\",\"type\":\"address\"}],\"name\":\"set_cert_center\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory_addr\",\"type\":\"address\"}],\"name\":\"set_provider_factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrderFactory is an auto generated Go binding around an Ethereum contract.
type OrderFactory struct {
	OrderFactoryCaller     // Read-only binding to the contract
	OrderFactoryTransactor // Write-only binding to the contract
	OrderFactoryFilterer   // Log filterer for contract events
}

// OrderFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderFactorySession struct {
	Contract     *OrderFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderFactoryCallerSession struct {
	Contract *OrderFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OrderFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderFactoryTransactorSession struct {
	Contract     *OrderFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OrderFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderFactoryRaw struct {
	Contract *OrderFactory // Generic contract binding to access the raw methods on
}

// OrderFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderFactoryCallerRaw struct {
	Contract *OrderFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OrderFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderFactoryTransactorRaw struct {
	Contract *OrderFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderFactory creates a new instance of OrderFactory, bound to a specific deployed contract.
func NewOrderFactory(address common.Address, backend bind.ContractBackend) (*OrderFactory, error) {
	contract, err := bindOrderFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderFactory{OrderFactoryCaller: OrderFactoryCaller{contract: contract}, OrderFactoryTransactor: OrderFactoryTransactor{contract: contract}, OrderFactoryFilterer: OrderFactoryFilterer{contract: contract}}, nil
}

// NewOrderFactoryCaller creates a new read-only instance of OrderFactory, bound to a specific deployed contract.
func NewOrderFactoryCaller(address common.Address, caller bind.ContractCaller) (*OrderFactoryCaller, error) {
	contract, err := bindOrderFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFactoryCaller{contract: contract}, nil
}

// NewOrderFactoryTransactor creates a new write-only instance of OrderFactory, bound to a specific deployed contract.
func NewOrderFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderFactoryTransactor, error) {
	contract, err := bindOrderFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFactoryTransactor{contract: contract}, nil
}

// NewOrderFactoryFilterer creates a new log filterer instance of OrderFactory, bound to a specific deployed contract.
func NewOrderFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFactoryFilterer, error) {
	contract, err := bindOrderFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFactoryFilterer{contract: contract}, nil
}

// bindOrderFactory binds a generic wrapper to an already deployed contract.
func bindOrderFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFactory *OrderFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFactory.Contract.OrderFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFactory *OrderFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFactory.Contract.OrderFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFactory *OrderFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFactory.Contract.OrderFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFactory *OrderFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFactory *OrderFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFactory *OrderFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFactory.Contract.contract.Transact(opts, method, params...)
}

// CertCenter is a free data retrieval call binding the contract method 0x7a197933.
//
// Solidity: function cert_center() view returns(address)
func (_OrderFactory *OrderFactoryCaller) CertCenter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "cert_center")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CertCenter is a free data retrieval call binding the contract method 0x7a197933.
//
// Solidity: function cert_center() view returns(address)
func (_OrderFactory *OrderFactorySession) CertCenter() (common.Address, error) {
	return _OrderFactory.Contract.CertCenter(&_OrderFactory.CallOpts)
}

// CertCenter is a free data retrieval call binding the contract method 0x7a197933.
//
// Solidity: function cert_center() view returns(address)
func (_OrderFactory *OrderFactoryCallerSession) CertCenter() (common.Address, error) {
	return _OrderFactory.Contract.CertCenter(&_OrderFactory.CallOpts)
}

// CheckIsOrder is a free data retrieval call binding the contract method 0x4bb21d44.
//
// Solidity: function checkIsOrder(address orderAddress) view returns(uint256)
func (_OrderFactory *OrderFactoryCaller) CheckIsOrder(opts *bind.CallOpts, orderAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "checkIsOrder", orderAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckIsOrder is a free data retrieval call binding the contract method 0x4bb21d44.
//
// Solidity: function checkIsOrder(address orderAddress) view returns(uint256)
func (_OrderFactory *OrderFactorySession) CheckIsOrder(orderAddress common.Address) (*big.Int, error) {
	return _OrderFactory.Contract.CheckIsOrder(&_OrderFactory.CallOpts, orderAddress)
}

// CheckIsOrder is a free data retrieval call binding the contract method 0x4bb21d44.
//
// Solidity: function checkIsOrder(address orderAddress) view returns(uint256)
func (_OrderFactory *OrderFactoryCallerSession) CheckIsOrder(orderAddress common.Address) (*big.Int, error) {
	return _OrderFactory.Contract.CheckIsOrder(&_OrderFactory.CallOpts, orderAddress)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((address,uint256,uint256,uint256,string,uint256,uint8))
func (_OrderFactory *OrderFactoryCaller) GetOrder(opts *bind.CallOpts, orderId *big.Int) (Order, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "getOrder", orderId)

	if err != nil {
		return *new(Order), err
	}

	out0 := *abi.ConvertType(out[0], new(Order)).(*Order)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((address,uint256,uint256,uint256,string,uint256,uint8))
func (_OrderFactory *OrderFactorySession) GetOrder(orderId *big.Int) (Order, error) {
	return _OrderFactory.Contract.GetOrder(&_OrderFactory.CallOpts, orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((address,uint256,uint256,uint256,string,uint256,uint8))
func (_OrderFactory *OrderFactoryCallerSession) GetOrder(orderId *big.Int) (Order, error) {
	return _OrderFactory.Contract.GetOrder(&_OrderFactory.CallOpts, orderId)
}

// GetProviderAllOrder is a free data retrieval call binding the contract method 0xd35c294f.
//
// Solidity: function getProviderAllOrder(address providerAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCaller) GetProviderAllOrder(opts *bind.CallOpts, providerAddress common.Address) ([]Order, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "getProviderAllOrder", providerAddress)

	if err != nil {
		return *new([]Order), err
	}

	out0 := *abi.ConvertType(out[0], new([]Order)).(*[]Order)

	return out0, err

}

// GetProviderAllOrder is a free data retrieval call binding the contract method 0xd35c294f.
//
// Solidity: function getProviderAllOrder(address providerAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactorySession) GetProviderAllOrder(providerAddress common.Address) ([]Order, error) {
	return _OrderFactory.Contract.GetProviderAllOrder(&_OrderFactory.CallOpts, providerAddress)
}

// GetProviderAllOrder is a free data retrieval call binding the contract method 0xd35c294f.
//
// Solidity: function getProviderAllOrder(address providerAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCallerSession) GetProviderAllOrder(providerAddress common.Address) ([]Order, error) {
	return _OrderFactory.Contract.GetProviderAllOrder(&_OrderFactory.CallOpts, providerAddress)
}

// GetUnCompleteOrder is a free data retrieval call binding the contract method 0x64ba9d9a.
//
// Solidity: function getUnCompleteOrder() view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCaller) GetUnCompleteOrder(opts *bind.CallOpts) ([]Order, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "getUnCompleteOrder")

	if err != nil {
		return *new([]Order), err
	}

	out0 := *abi.ConvertType(out[0], new([]Order)).(*[]Order)

	return out0, err

}

// GetUnCompleteOrder is a free data retrieval call binding the contract method 0x64ba9d9a.
//
// Solidity: function getUnCompleteOrder() view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactorySession) GetUnCompleteOrder() ([]Order, error) {
	return _OrderFactory.Contract.GetUnCompleteOrder(&_OrderFactory.CallOpts)
}

// GetUnCompleteOrder is a free data retrieval call binding the contract method 0x64ba9d9a.
//
// Solidity: function getUnCompleteOrder() view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCallerSession) GetUnCompleteOrder() ([]Order, error) {
	return _OrderFactory.Contract.GetUnCompleteOrder(&_OrderFactory.CallOpts)
}

// GetUserAllOrder is a free data retrieval call binding the contract method 0x3fe8ffee.
//
// Solidity: function getUserAllOrder(address userAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCaller) GetUserAllOrder(opts *bind.CallOpts, userAddress common.Address) ([]Order, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "getUserAllOrder", userAddress)

	if err != nil {
		return *new([]Order), err
	}

	out0 := *abi.ConvertType(out[0], new([]Order)).(*[]Order)

	return out0, err

}

// GetUserAllOrder is a free data retrieval call binding the contract method 0x3fe8ffee.
//
// Solidity: function getUserAllOrder(address userAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactorySession) GetUserAllOrder(userAddress common.Address) ([]Order, error) {
	return _OrderFactory.Contract.GetUserAllOrder(&_OrderFactory.CallOpts, userAddress)
}

// GetUserAllOrder is a free data retrieval call binding the contract method 0x3fe8ffee.
//
// Solidity: function getUserAllOrder(address userAddress) view returns((address,uint256,uint256,uint256,string,uint256,uint8)[])
func (_OrderFactory *OrderFactoryCallerSession) GetUserAllOrder(userAddress common.Address) ([]Order, error) {
	return _OrderFactory.Contract.GetUserAllOrder(&_OrderFactory.CallOpts, userAddress)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xd7d35f10.
//
// Solidity: function get_minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactoryCaller) GetMinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "get_minimum_deposit_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xd7d35f10.
//
// Solidity: function get_minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactorySession) GetMinimumDepositAmount() (*big.Int, error) {
	return _OrderFactory.Contract.GetMinimumDepositAmount(&_OrderFactory.CallOpts)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xd7d35f10.
//
// Solidity: function get_minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactoryCallerSession) GetMinimumDepositAmount() (*big.Int, error) {
	return _OrderFactory.Contract.GetMinimumDepositAmount(&_OrderFactory.CallOpts)
}

// MaxOrderIndex is a free data retrieval call binding the contract method 0x7274c49f.
//
// Solidity: function max_order_index() view returns(uint256)
func (_OrderFactory *OrderFactoryCaller) MaxOrderIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "max_order_index")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOrderIndex is a free data retrieval call binding the contract method 0x7274c49f.
//
// Solidity: function max_order_index() view returns(uint256)
func (_OrderFactory *OrderFactorySession) MaxOrderIndex() (*big.Int, error) {
	return _OrderFactory.Contract.MaxOrderIndex(&_OrderFactory.CallOpts)
}

// MaxOrderIndex is a free data retrieval call binding the contract method 0x7274c49f.
//
// Solidity: function max_order_index() view returns(uint256)
func (_OrderFactory *OrderFactoryCallerSession) MaxOrderIndex() (*big.Int, error) {
	return _OrderFactory.Contract.MaxOrderIndex(&_OrderFactory.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0xa9d66352.
//
// Solidity: function minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactoryCaller) MinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "minimum_deposit_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0xa9d66352.
//
// Solidity: function minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactorySession) MinimumDepositAmount() (*big.Int, error) {
	return _OrderFactory.Contract.MinimumDepositAmount(&_OrderFactory.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0xa9d66352.
//
// Solidity: function minimum_deposit_amount() view returns(uint256)
func (_OrderFactory *OrderFactoryCallerSession) MinimumDepositAmount() (*big.Int, error) {
	return _OrderFactory.Contract.MinimumDepositAmount(&_OrderFactory.CallOpts)
}

// OrderBaseMap is a free data retrieval call binding the contract method 0xe202e048.
//
// Solidity: function order_base_map(address ) view returns(uint256)
func (_OrderFactory *OrderFactoryCaller) OrderBaseMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "order_base_map", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderBaseMap is a free data retrieval call binding the contract method 0xe202e048.
//
// Solidity: function order_base_map(address ) view returns(uint256)
func (_OrderFactory *OrderFactorySession) OrderBaseMap(arg0 common.Address) (*big.Int, error) {
	return _OrderFactory.Contract.OrderBaseMap(&_OrderFactory.CallOpts, arg0)
}

// OrderBaseMap is a free data retrieval call binding the contract method 0xe202e048.
//
// Solidity: function order_base_map(address ) view returns(uint256)
func (_OrderFactory *OrderFactoryCallerSession) OrderBaseMap(arg0 common.Address) (*big.Int, error) {
	return _OrderFactory.Contract.OrderBaseMap(&_OrderFactory.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address)
func (_OrderFactory *OrderFactoryCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "orders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address)
func (_OrderFactory *OrderFactorySession) Orders(arg0 *big.Int) (common.Address, error) {
	return _OrderFactory.Contract.Orders(&_OrderFactory.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address)
func (_OrderFactory *OrderFactoryCallerSession) Orders(arg0 *big.Int) (common.Address, error) {
	return _OrderFactory.Contract.Orders(&_OrderFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderFactory *OrderFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderFactory *OrderFactorySession) Owner() (common.Address, error) {
	return _OrderFactory.Contract.Owner(&_OrderFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderFactory *OrderFactoryCallerSession) Owner() (common.Address, error) {
	return _OrderFactory.Contract.Owner(&_OrderFactory.CallOpts)
}

// ProviderAddress is a free data retrieval call binding the contract method 0x0e5630c1.
//
// Solidity: function provider_address() view returns(address)
func (_OrderFactory *OrderFactoryCaller) ProviderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderFactory.contract.Call(opts, &out, "provider_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderAddress is a free data retrieval call binding the contract method 0x0e5630c1.
//
// Solidity: function provider_address() view returns(address)
func (_OrderFactory *OrderFactorySession) ProviderAddress() (common.Address, error) {
	return _OrderFactory.Contract.ProviderAddress(&_OrderFactory.CallOpts)
}

// ProviderAddress is a free data retrieval call binding the contract method 0x0e5630c1.
//
// Solidity: function provider_address() view returns(address)
func (_OrderFactory *OrderFactoryCallerSession) ProviderAddress() (common.Address, error) {
	return _OrderFactory.Contract.ProviderAddress(&_OrderFactory.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address new_owner) returns()
func (_OrderFactory *OrderFactoryTransactor) ChangeOwner(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _OrderFactory.contract.Transact(opts, "changeOwner", new_owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address new_owner) returns()
func (_OrderFactory *OrderFactorySession) ChangeOwner(new_owner common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.ChangeOwner(&_OrderFactory.TransactOpts, new_owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address new_owner) returns()
func (_OrderFactory *OrderFactoryTransactorSession) ChangeOwner(new_owner common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.ChangeOwner(&_OrderFactory.TransactOpts, new_owner)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb4840fdb.
//
// Solidity: function createOrder(uint256 m_cpu, uint256 m_memory, uint256 m_storage, string m_cert, uint256 m_trx_id) returns(address)
func (_OrderFactory *OrderFactoryTransactor) CreateOrder(opts *bind.TransactOpts, m_cpu *big.Int, m_memory *big.Int, m_storage *big.Int, m_cert string, m_trx_id *big.Int) (*types.Transaction, error) {
	return _OrderFactory.contract.Transact(opts, "createOrder", m_cpu, m_memory, m_storage, m_cert, m_trx_id)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb4840fdb.
//
// Solidity: function createOrder(uint256 m_cpu, uint256 m_memory, uint256 m_storage, string m_cert, uint256 m_trx_id) returns(address)
func (_OrderFactory *OrderFactorySession) CreateOrder(m_cpu *big.Int, m_memory *big.Int, m_storage *big.Int, m_cert string, m_trx_id *big.Int) (*types.Transaction, error) {
	return _OrderFactory.Contract.CreateOrder(&_OrderFactory.TransactOpts, m_cpu, m_memory, m_storage, m_cert, m_trx_id)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb4840fdb.
//
// Solidity: function createOrder(uint256 m_cpu, uint256 m_memory, uint256 m_storage, string m_cert, uint256 m_trx_id) returns(address)
func (_OrderFactory *OrderFactoryTransactorSession) CreateOrder(m_cpu *big.Int, m_memory *big.Int, m_storage *big.Int, m_cert string, m_trx_id *big.Int) (*types.Transaction, error) {
	return _OrderFactory.Contract.CreateOrder(&_OrderFactory.TransactOpts, m_cpu, m_memory, m_storage, m_cert, m_trx_id)
}

// ModifyMinimumDepositAmount is a paid mutator transaction binding the contract method 0x8e3fbb7a.
//
// Solidity: function modify_minimum_deposit_amount(uint256 new_value) returns()
func (_OrderFactory *OrderFactoryTransactor) ModifyMinimumDepositAmount(opts *bind.TransactOpts, new_value *big.Int) (*types.Transaction, error) {
	return _OrderFactory.contract.Transact(opts, "modify_minimum_deposit_amount", new_value)
}

// ModifyMinimumDepositAmount is a paid mutator transaction binding the contract method 0x8e3fbb7a.
//
// Solidity: function modify_minimum_deposit_amount(uint256 new_value) returns()
func (_OrderFactory *OrderFactorySession) ModifyMinimumDepositAmount(new_value *big.Int) (*types.Transaction, error) {
	return _OrderFactory.Contract.ModifyMinimumDepositAmount(&_OrderFactory.TransactOpts, new_value)
}

// ModifyMinimumDepositAmount is a paid mutator transaction binding the contract method 0x8e3fbb7a.
//
// Solidity: function modify_minimum_deposit_amount(uint256 new_value) returns()
func (_OrderFactory *OrderFactoryTransactorSession) ModifyMinimumDepositAmount(new_value *big.Int) (*types.Transaction, error) {
	return _OrderFactory.Contract.ModifyMinimumDepositAmount(&_OrderFactory.TransactOpts, new_value)
}

// SetCertCenter is a paid mutator transaction binding the contract method 0x306e72ff.
//
// Solidity: function set_cert_center(address cert_center_) returns()
func (_OrderFactory *OrderFactoryTransactor) SetCertCenter(opts *bind.TransactOpts, cert_center_ common.Address) (*types.Transaction, error) {
	return _OrderFactory.contract.Transact(opts, "set_cert_center", cert_center_)
}

// SetCertCenter is a paid mutator transaction binding the contract method 0x306e72ff.
//
// Solidity: function set_cert_center(address cert_center_) returns()
func (_OrderFactory *OrderFactorySession) SetCertCenter(cert_center_ common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.SetCertCenter(&_OrderFactory.TransactOpts, cert_center_)
}

// SetCertCenter is a paid mutator transaction binding the contract method 0x306e72ff.
//
// Solidity: function set_cert_center(address cert_center_) returns()
func (_OrderFactory *OrderFactoryTransactorSession) SetCertCenter(cert_center_ common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.SetCertCenter(&_OrderFactory.TransactOpts, cert_center_)
}

// SetProviderFactory is a paid mutator transaction binding the contract method 0x27c1bb4c.
//
// Solidity: function set_provider_factory(address factory_addr) returns()
func (_OrderFactory *OrderFactoryTransactor) SetProviderFactory(opts *bind.TransactOpts, factory_addr common.Address) (*types.Transaction, error) {
	return _OrderFactory.contract.Transact(opts, "set_provider_factory", factory_addr)
}

// SetProviderFactory is a paid mutator transaction binding the contract method 0x27c1bb4c.
//
// Solidity: function set_provider_factory(address factory_addr) returns()
func (_OrderFactory *OrderFactorySession) SetProviderFactory(factory_addr common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.SetProviderFactory(&_OrderFactory.TransactOpts, factory_addr)
}

// SetProviderFactory is a paid mutator transaction binding the contract method 0x27c1bb4c.
//
// Solidity: function set_provider_factory(address factory_addr) returns()
func (_OrderFactory *OrderFactoryTransactorSession) SetProviderFactory(factory_addr common.Address) (*types.Transaction, error) {
	return _OrderFactory.Contract.SetProviderFactory(&_OrderFactory.TransactOpts, factory_addr)
}

// OrderFactoryOrderCreationIterator is returned from FilterOrderCreation and is used to iterate over the raw logs and unpacked data for OrderCreation events raised by the OrderFactory contract.
type OrderFactoryOrderCreationIterator struct {
	Event *OrderFactoryOrderCreation // Event containing the contract specifics and raw log

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
func (it *OrderFactoryOrderCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderFactoryOrderCreation)
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
		it.Event = new(OrderFactoryOrderCreation)
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
func (it *OrderFactoryOrderCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderFactoryOrderCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderFactoryOrderCreation represents a OrderCreation event raised by the OrderFactory contract.
type OrderFactoryOrderCreation struct {
	OrderNumber *big.Int
	Owner       common.Address
	OrderAddr   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderCreation is a free log retrieval operation binding the contract event 0x61efd0de216a2d10b7f644bfc63a67b801738395f7fb2fb8f45569143f7d9b19.
//
// Solidity: event OrderCreation(uint256 indexed orderNumber, address indexed owner, address indexed order_addr)
func (_OrderFactory *OrderFactoryFilterer) FilterOrderCreation(opts *bind.FilterOpts, orderNumber []*big.Int, owner []common.Address, order_addr []common.Address) (*OrderFactoryOrderCreationIterator, error) {

	var orderNumberRule []interface{}
	for _, orderNumberItem := range orderNumber {
		orderNumberRule = append(orderNumberRule, orderNumberItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var order_addrRule []interface{}
	for _, order_addrItem := range order_addr {
		order_addrRule = append(order_addrRule, order_addrItem)
	}

	logs, sub, err := _OrderFactory.contract.FilterLogs(opts, "OrderCreation", orderNumberRule, ownerRule, order_addrRule)
	if err != nil {
		return nil, err
	}
	return &OrderFactoryOrderCreationIterator{contract: _OrderFactory.contract, event: "OrderCreation", logs: logs, sub: sub}, nil
}

// WatchOrderCreation is a free log subscription operation binding the contract event 0x61efd0de216a2d10b7f644bfc63a67b801738395f7fb2fb8f45569143f7d9b19.
//
// Solidity: event OrderCreation(uint256 indexed orderNumber, address indexed owner, address indexed order_addr)
func (_OrderFactory *OrderFactoryFilterer) WatchOrderCreation(opts *bind.WatchOpts, sink chan<- *OrderFactoryOrderCreation, orderNumber []*big.Int, owner []common.Address, order_addr []common.Address) (event.Subscription, error) {

	var orderNumberRule []interface{}
	for _, orderNumberItem := range orderNumber {
		orderNumberRule = append(orderNumberRule, orderNumberItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var order_addrRule []interface{}
	for _, order_addrItem := range order_addr {
		order_addrRule = append(order_addrRule, order_addrItem)
	}

	logs, sub, err := _OrderFactory.contract.WatchLogs(opts, "OrderCreation", orderNumberRule, ownerRule, order_addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderFactoryOrderCreation)
				if err := _OrderFactory.contract.UnpackLog(event, "OrderCreation", log); err != nil {
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

// ParseOrderCreation is a log parse operation binding the contract event 0x61efd0de216a2d10b7f644bfc63a67b801738395f7fb2fb8f45569143f7d9b19.
//
// Solidity: event OrderCreation(uint256 indexed orderNumber, address indexed owner, address indexed order_addr)
func (_OrderFactory *OrderFactoryFilterer) ParseOrderCreation(log types.Log) (*OrderFactoryOrderCreation, error) {
	event := new(OrderFactoryOrderCreation)
	if err := _OrderFactory.contract.UnpackLog(event, "OrderCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}

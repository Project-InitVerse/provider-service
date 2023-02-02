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

// Order is an auto generated low-level Go binding around an user-defined struct.

// PriceOracle is an auto generated low-level Go binding around an user-defined struct.
type PriceOracle struct {
	Provider     common.Address
	CpuPrice     *big.Int
	MemoryPrice  *big.Int
	StoragePrice *big.Int
}

// OrderBaseABI is the input ABI used to generate the binding from.
const OrderBaseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_order_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider_factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cert_key_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdl_trx_id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"order_number\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CanQuote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"final_price\",\"type\":\"uint256\"}],\"name\":\"ChooseQuote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_memory\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_storage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_memory\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_storage\",\"type\":\"uint256\"}],\"name\":\"DeployMentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memory_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storage_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cert\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"order_number\",\"type\":\"uint256\"}],\"name\":\"OrderCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OrderEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PayBill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpu_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memory_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storage_price\",\"type\":\"uint256\"}],\"name\":\"Quote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_sdl_trx_id\",\"type\":\"uint256\"}],\"name\":\"UpdateSDL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UserCancelOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_trx_hash\",\"type\":\"uint256\"}],\"name\":\"change_sdl_trx_hash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quote_index\",\"type\":\"uint256\"}],\"name\":\"choose_provider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit_balance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"final_choice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"final_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_pay_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_cert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_cpu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_memory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_order_number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_pending_sdl_trx_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_sdl_trx_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"o_storage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"order_info\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"v_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v_storage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cert_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trx_id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"order_status\",\"outputs\":[{\"internalType\":\"enumOrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pay_billing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"provide_quotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider_factory\",\"outputs\":[{\"internalType\":\"contractIProviderFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"query_provide_quotes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_price\",\"type\":\"uint256\"}],\"internalType\":\"structPriceOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"query_provider_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"p_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p_memory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p_storage\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"server_uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"submit_server_uri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"update_deployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrderBase is an auto generated Go binding around an Ethereum contract.
type OrderBase struct {
	OrderBaseCaller     // Read-only binding to the contract
	OrderBaseTransactor // Write-only binding to the contract
	OrderBaseFilterer   // Log filterer for contract events
}

// OrderBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBaseSession struct {
	Contract     *OrderBase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBaseCallerSession struct {
	Contract *OrderBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBaseTransactorSession struct {
	Contract     *OrderBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBaseRaw struct {
	Contract *OrderBase // Generic contract binding to access the raw methods on
}

// OrderBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBaseCallerRaw struct {
	Contract *OrderBaseCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBaseTransactorRaw struct {
	Contract *OrderBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBase creates a new instance of OrderBase, bound to a specific deployed contract.
func NewOrderBase(address common.Address, backend bind.ContractBackend) (*OrderBase, error) {
	contract, err := bindOrderBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBase{OrderBaseCaller: OrderBaseCaller{contract: contract}, OrderBaseTransactor: OrderBaseTransactor{contract: contract}, OrderBaseFilterer: OrderBaseFilterer{contract: contract}}, nil
}

// NewOrderBaseCaller creates a new read-only instance of OrderBase, bound to a specific deployed contract.
func NewOrderBaseCaller(address common.Address, caller bind.ContractCaller) (*OrderBaseCaller, error) {
	contract, err := bindOrderBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBaseCaller{contract: contract}, nil
}

// NewOrderBaseTransactor creates a new write-only instance of OrderBase, bound to a specific deployed contract.
func NewOrderBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBaseTransactor, error) {
	contract, err := bindOrderBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBaseTransactor{contract: contract}, nil
}

// NewOrderBaseFilterer creates a new log filterer instance of OrderBase, bound to a specific deployed contract.
func NewOrderBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBaseFilterer, error) {
	contract, err := bindOrderBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBaseFilterer{contract: contract}, nil
}

// bindOrderBase binds a generic wrapper to an already deployed contract.
func bindOrderBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBase *OrderBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBase.Contract.OrderBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBase *OrderBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBase.Contract.OrderBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBase *OrderBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBase.Contract.OrderBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBase *OrderBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBase *OrderBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBase *OrderBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBase.Contract.contract.Transact(opts, method, params...)
}

// FinalChoice is a free data retrieval call binding the contract method 0xc8c0c097.
//
// Solidity: function final_choice() view returns(uint256)
func (_OrderBase *OrderBaseCaller) FinalChoice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "final_choice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalChoice is a free data retrieval call binding the contract method 0xc8c0c097.
//
// Solidity: function final_choice() view returns(uint256)
func (_OrderBase *OrderBaseSession) FinalChoice() (*big.Int, error) {
	return _OrderBase.Contract.FinalChoice(&_OrderBase.CallOpts)
}

// FinalChoice is a free data retrieval call binding the contract method 0xc8c0c097.
//
// Solidity: function final_choice() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) FinalChoice() (*big.Int, error) {
	return _OrderBase.Contract.FinalChoice(&_OrderBase.CallOpts)
}

// FinalPrice is a free data retrieval call binding the contract method 0x08b309e3.
//
// Solidity: function final_price() view returns(uint256)
func (_OrderBase *OrderBaseCaller) FinalPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "final_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalPrice is a free data retrieval call binding the contract method 0x08b309e3.
//
// Solidity: function final_price() view returns(uint256)
func (_OrderBase *OrderBaseSession) FinalPrice() (*big.Int, error) {
	return _OrderBase.Contract.FinalPrice(&_OrderBase.CallOpts)
}

// FinalPrice is a free data retrieval call binding the contract method 0x08b309e3.
//
// Solidity: function final_price() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) FinalPrice() (*big.Int, error) {
	return _OrderBase.Contract.FinalPrice(&_OrderBase.CallOpts)
}

// LastPayTime is a free data retrieval call binding the contract method 0x91c4d7a6.
//
// Solidity: function last_pay_time() view returns(uint256)
func (_OrderBase *OrderBaseCaller) LastPayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "last_pay_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPayTime is a free data retrieval call binding the contract method 0x91c4d7a6.
//
// Solidity: function last_pay_time() view returns(uint256)
func (_OrderBase *OrderBaseSession) LastPayTime() (*big.Int, error) {
	return _OrderBase.Contract.LastPayTime(&_OrderBase.CallOpts)
}

// LastPayTime is a free data retrieval call binding the contract method 0x91c4d7a6.
//
// Solidity: function last_pay_time() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) LastPayTime() (*big.Int, error) {
	return _OrderBase.Contract.LastPayTime(&_OrderBase.CallOpts)
}

// OCert is a free data retrieval call binding the contract method 0xcc61ac87.
//
// Solidity: function o_cert() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OCert(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_cert")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OCert is a free data retrieval call binding the contract method 0xcc61ac87.
//
// Solidity: function o_cert() view returns(uint256)
func (_OrderBase *OrderBaseSession) OCert() (*big.Int, error) {
	return _OrderBase.Contract.OCert(&_OrderBase.CallOpts)
}

// OCert is a free data retrieval call binding the contract method 0xcc61ac87.
//
// Solidity: function o_cert() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OCert() (*big.Int, error) {
	return _OrderBase.Contract.OCert(&_OrderBase.CallOpts)
}

// OCpu is a free data retrieval call binding the contract method 0x27fb1fb2.
//
// Solidity: function o_cpu() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OCpu(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_cpu")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OCpu is a free data retrieval call binding the contract method 0x27fb1fb2.
//
// Solidity: function o_cpu() view returns(uint256)
func (_OrderBase *OrderBaseSession) OCpu() (*big.Int, error) {
	return _OrderBase.Contract.OCpu(&_OrderBase.CallOpts)
}

// OCpu is a free data retrieval call binding the contract method 0x27fb1fb2.
//
// Solidity: function o_cpu() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OCpu() (*big.Int, error) {
	return _OrderBase.Contract.OCpu(&_OrderBase.CallOpts)
}

// OMemory is a free data retrieval call binding the contract method 0x157b02e7.
//
// Solidity: function o_memory() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OMemory(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_memory")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OMemory is a free data retrieval call binding the contract method 0x157b02e7.
//
// Solidity: function o_memory() view returns(uint256)
func (_OrderBase *OrderBaseSession) OMemory() (*big.Int, error) {
	return _OrderBase.Contract.OMemory(&_OrderBase.CallOpts)
}

// OMemory is a free data retrieval call binding the contract method 0x157b02e7.
//
// Solidity: function o_memory() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OMemory() (*big.Int, error) {
	return _OrderBase.Contract.OMemory(&_OrderBase.CallOpts)
}

// OOrderNumber is a free data retrieval call binding the contract method 0xc2e1d74d.
//
// Solidity: function o_order_number() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OOrderNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_order_number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OOrderNumber is a free data retrieval call binding the contract method 0xc2e1d74d.
//
// Solidity: function o_order_number() view returns(uint256)
func (_OrderBase *OrderBaseSession) OOrderNumber() (*big.Int, error) {
	return _OrderBase.Contract.OOrderNumber(&_OrderBase.CallOpts)
}

// OOrderNumber is a free data retrieval call binding the contract method 0xc2e1d74d.
//
// Solidity: function o_order_number() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OOrderNumber() (*big.Int, error) {
	return _OrderBase.Contract.OOrderNumber(&_OrderBase.CallOpts)
}

// OPendingSdlTrxId is a free data retrieval call binding the contract method 0xd83d7ff3.
//
// Solidity: function o_pending_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OPendingSdlTrxId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_pending_sdl_trx_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPendingSdlTrxId is a free data retrieval call binding the contract method 0xd83d7ff3.
//
// Solidity: function o_pending_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseSession) OPendingSdlTrxId() (*big.Int, error) {
	return _OrderBase.Contract.OPendingSdlTrxId(&_OrderBase.CallOpts)
}

// OPendingSdlTrxId is a free data retrieval call binding the contract method 0xd83d7ff3.
//
// Solidity: function o_pending_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OPendingSdlTrxId() (*big.Int, error) {
	return _OrderBase.Contract.OPendingSdlTrxId(&_OrderBase.CallOpts)
}

// OSdlTrxId is a free data retrieval call binding the contract method 0xdb3e43ce.
//
// Solidity: function o_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OSdlTrxId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_sdl_trx_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OSdlTrxId is a free data retrieval call binding the contract method 0xdb3e43ce.
//
// Solidity: function o_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseSession) OSdlTrxId() (*big.Int, error) {
	return _OrderBase.Contract.OSdlTrxId(&_OrderBase.CallOpts)
}

// OSdlTrxId is a free data retrieval call binding the contract method 0xdb3e43ce.
//
// Solidity: function o_sdl_trx_id() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OSdlTrxId() (*big.Int, error) {
	return _OrderBase.Contract.OSdlTrxId(&_OrderBase.CallOpts)
}

// OStorage is a free data retrieval call binding the contract method 0x8e15c70a.
//
// Solidity: function o_storage() view returns(uint256)
func (_OrderBase *OrderBaseCaller) OStorage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "o_storage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OStorage is a free data retrieval call binding the contract method 0x8e15c70a.
//
// Solidity: function o_storage() view returns(uint256)
func (_OrderBase *OrderBaseSession) OStorage() (*big.Int, error) {
	return _OrderBase.Contract.OStorage(&_OrderBase.CallOpts)
}

// OStorage is a free data retrieval call binding the contract method 0x8e15c70a.
//
// Solidity: function o_storage() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) OStorage() (*big.Int, error) {
	return _OrderBase.Contract.OStorage(&_OrderBase.CallOpts)
}

// OrderInfo is a free data retrieval call binding the contract method 0x2500f01d.
//
// Solidity: function order_info() view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint8,uint256))
func (_OrderBase *OrderBaseCaller) OrderInfo(opts *bind.CallOpts) (Order, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "order_info")

	if err != nil {
		return *new(Order), err
	}

	out0 := *abi.ConvertType(out[0], new(Order)).(*Order)

	return out0, err

}

// OrderInfo is a free data retrieval call binding the contract method 0x2500f01d.
//
// Solidity: function order_info() view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint8,uint256))
func (_OrderBase *OrderBaseSession) OrderInfo() (Order, error) {
	return _OrderBase.Contract.OrderInfo(&_OrderBase.CallOpts)
}

// OrderInfo is a free data retrieval call binding the contract method 0x2500f01d.
//
// Solidity: function order_info() view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint8,uint256))
func (_OrderBase *OrderBaseCallerSession) OrderInfo() (Order, error) {
	return _OrderBase.Contract.OrderInfo(&_OrderBase.CallOpts)
}

// OrderStatus is a free data retrieval call binding the contract method 0x55532c2e.
//
// Solidity: function order_status() view returns(uint8)
func (_OrderBase *OrderBaseCaller) OrderStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "order_status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x55532c2e.
//
// Solidity: function order_status() view returns(uint8)
func (_OrderBase *OrderBaseSession) OrderStatus() (uint8, error) {
	return _OrderBase.Contract.OrderStatus(&_OrderBase.CallOpts)
}

// OrderStatus is a free data retrieval call binding the contract method 0x55532c2e.
//
// Solidity: function order_status() view returns(uint8)
func (_OrderBase *OrderBaseCallerSession) OrderStatus() (uint8, error) {
	return _OrderBase.Contract.OrderStatus(&_OrderBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderBase *OrderBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderBase *OrderBaseSession) Owner() (common.Address, error) {
	return _OrderBase.Contract.Owner(&_OrderBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderBase *OrderBaseCallerSession) Owner() (common.Address, error) {
	return _OrderBase.Contract.Owner(&_OrderBase.CallOpts)
}

// ProvideQuotes is a free data retrieval call binding the contract method 0xd2c3aaf5.
//
// Solidity: function provide_quotes(uint256 ) view returns(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseCaller) ProvideQuotes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Provider     common.Address
	CpuPrice     *big.Int
	MemoryPrice  *big.Int
	StoragePrice *big.Int
}, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "provide_quotes", arg0)

	outstruct := new(struct {
		Provider     common.Address
		CpuPrice     *big.Int
		MemoryPrice  *big.Int
		StoragePrice *big.Int
	})

	outstruct.Provider = out[0].(common.Address)
	outstruct.CpuPrice = out[1].(*big.Int)
	outstruct.MemoryPrice = out[2].(*big.Int)
	outstruct.StoragePrice = out[3].(*big.Int)

	return *outstruct, err

}

// ProvideQuotes is a free data retrieval call binding the contract method 0xd2c3aaf5.
//
// Solidity: function provide_quotes(uint256 ) view returns(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseSession) ProvideQuotes(arg0 *big.Int) (struct {
	Provider     common.Address
	CpuPrice     *big.Int
	MemoryPrice  *big.Int
	StoragePrice *big.Int
}, error) {
	return _OrderBase.Contract.ProvideQuotes(&_OrderBase.CallOpts, arg0)
}

// ProvideQuotes is a free data retrieval call binding the contract method 0xd2c3aaf5.
//
// Solidity: function provide_quotes(uint256 ) view returns(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseCallerSession) ProvideQuotes(arg0 *big.Int) (struct {
	Provider     common.Address
	CpuPrice     *big.Int
	MemoryPrice  *big.Int
	StoragePrice *big.Int
}, error) {
	return _OrderBase.Contract.ProvideQuotes(&_OrderBase.CallOpts, arg0)
}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_OrderBase *OrderBaseCaller) ProviderFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "provider_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_OrderBase *OrderBaseSession) ProviderFactory() (common.Address, error) {
	return _OrderBase.Contract.ProviderFactory(&_OrderBase.CallOpts)
}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_OrderBase *OrderBaseCallerSession) ProviderFactory() (common.Address, error) {
	return _OrderBase.Contract.ProviderFactory(&_OrderBase.CallOpts)
}

// QueryProvideQuotes is a free data retrieval call binding the contract method 0x214c9d35.
//
// Solidity: function query_provide_quotes() view returns((address,uint256,uint256,uint256)[])
func (_OrderBase *OrderBaseCaller) QueryProvideQuotes(opts *bind.CallOpts) ([]PriceOracle, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "query_provide_quotes")

	if err != nil {
		return *new([]PriceOracle), err
	}

	out0 := *abi.ConvertType(out[0], new([]PriceOracle)).(*[]PriceOracle)

	return out0, err

}

// QueryProvideQuotes is a free data retrieval call binding the contract method 0x214c9d35.
//
// Solidity: function query_provide_quotes() view returns((address,uint256,uint256,uint256)[])
func (_OrderBase *OrderBaseSession) QueryProvideQuotes() ([]PriceOracle, error) {
	return _OrderBase.Contract.QueryProvideQuotes(&_OrderBase.CallOpts)
}

// QueryProvideQuotes is a free data retrieval call binding the contract method 0x214c9d35.
//
// Solidity: function query_provide_quotes() view returns((address,uint256,uint256,uint256)[])
func (_OrderBase *OrderBaseCallerSession) QueryProvideQuotes() ([]PriceOracle, error) {
	return _OrderBase.Contract.QueryProvideQuotes(&_OrderBase.CallOpts)
}

// QueryProviderAddress is a free data retrieval call binding the contract method 0x3b3b5147.
//
// Solidity: function query_provider_address() view returns(address)
func (_OrderBase *OrderBaseCaller) QueryProviderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "query_provider_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QueryProviderAddress is a free data retrieval call binding the contract method 0x3b3b5147.
//
// Solidity: function query_provider_address() view returns(address)
func (_OrderBase *OrderBaseSession) QueryProviderAddress() (common.Address, error) {
	return _OrderBase.Contract.QueryProviderAddress(&_OrderBase.CallOpts)
}

// QueryProviderAddress is a free data retrieval call binding the contract method 0x3b3b5147.
//
// Solidity: function query_provider_address() view returns(address)
func (_OrderBase *OrderBaseCallerSession) QueryProviderAddress() (common.Address, error) {
	return _OrderBase.Contract.QueryProviderAddress(&_OrderBase.CallOpts)
}

// ServerUri is a free data retrieval call binding the contract method 0x025b36da.
//
// Solidity: function server_uri() view returns(string)
func (_OrderBase *OrderBaseCaller) ServerUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "server_uri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ServerUri is a free data retrieval call binding the contract method 0x025b36da.
//
// Solidity: function server_uri() view returns(string)
func (_OrderBase *OrderBaseSession) ServerUri() (string, error) {
	return _OrderBase.Contract.ServerUri(&_OrderBase.CallOpts)
}

// ServerUri is a free data retrieval call binding the contract method 0x025b36da.
//
// Solidity: function server_uri() view returns(string)
func (_OrderBase *OrderBaseCallerSession) ServerUri() (string, error) {
	return _OrderBase.Contract.ServerUri(&_OrderBase.CallOpts)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_OrderBase *OrderBaseCaller) TotalSpent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBase.contract.Call(opts, &out, "totalSpent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_OrderBase *OrderBaseSession) TotalSpent() (*big.Int, error) {
	return _OrderBase.Contract.TotalSpent(&_OrderBase.CallOpts)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_OrderBase *OrderBaseCallerSession) TotalSpent() (*big.Int, error) {
	return _OrderBase.Contract.TotalSpent(&_OrderBase.CallOpts)
}

// ChangeSdlTrxHash is a paid mutator transaction binding the contract method 0xd809d3c3.
//
// Solidity: function change_sdl_trx_hash(uint256 new_trx_hash) returns()
func (_OrderBase *OrderBaseTransactor) ChangeSdlTrxHash(opts *bind.TransactOpts, new_trx_hash *big.Int) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "change_sdl_trx_hash", new_trx_hash)
}

// ChangeSdlTrxHash is a paid mutator transaction binding the contract method 0xd809d3c3.
//
// Solidity: function change_sdl_trx_hash(uint256 new_trx_hash) returns()
func (_OrderBase *OrderBaseSession) ChangeSdlTrxHash(new_trx_hash *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.ChangeSdlTrxHash(&_OrderBase.TransactOpts, new_trx_hash)
}

// ChangeSdlTrxHash is a paid mutator transaction binding the contract method 0xd809d3c3.
//
// Solidity: function change_sdl_trx_hash(uint256 new_trx_hash) returns()
func (_OrderBase *OrderBaseTransactorSession) ChangeSdlTrxHash(new_trx_hash *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.ChangeSdlTrxHash(&_OrderBase.TransactOpts, new_trx_hash)
}

// ChooseProvider is a paid mutator transaction binding the contract method 0x8eb3078b.
//
// Solidity: function choose_provider(uint256 quote_index) returns()
func (_OrderBase *OrderBaseTransactor) ChooseProvider(opts *bind.TransactOpts, quote_index *big.Int) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "choose_provider", quote_index)
}

// ChooseProvider is a paid mutator transaction binding the contract method 0x8eb3078b.
//
// Solidity: function choose_provider(uint256 quote_index) returns()
func (_OrderBase *OrderBaseSession) ChooseProvider(quote_index *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.ChooseProvider(&_OrderBase.TransactOpts, quote_index)
}

// ChooseProvider is a paid mutator transaction binding the contract method 0x8eb3078b.
//
// Solidity: function choose_provider(uint256 quote_index) returns()
func (_OrderBase *OrderBaseTransactorSession) ChooseProvider(quote_index *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.ChooseProvider(&_OrderBase.TransactOpts, quote_index)
}

// DepositBalance is a paid mutator transaction binding the contract method 0x8627436b.
//
// Solidity: function deposit_balance() payable returns()
func (_OrderBase *OrderBaseTransactor) DepositBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "deposit_balance")
}

// DepositBalance is a paid mutator transaction binding the contract method 0x8627436b.
//
// Solidity: function deposit_balance() payable returns()
func (_OrderBase *OrderBaseSession) DepositBalance() (*types.Transaction, error) {
	return _OrderBase.Contract.DepositBalance(&_OrderBase.TransactOpts)
}

// DepositBalance is a paid mutator transaction binding the contract method 0x8627436b.
//
// Solidity: function deposit_balance() payable returns()
func (_OrderBase *OrderBaseTransactorSession) DepositBalance() (*types.Transaction, error) {
	return _OrderBase.Contract.DepositBalance(&_OrderBase.TransactOpts)
}

// PayBilling is a paid mutator transaction binding the contract method 0x4e1e2fe1.
//
// Solidity: function pay_billing() returns()
func (_OrderBase *OrderBaseTransactor) PayBilling(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "pay_billing")
}

// PayBilling is a paid mutator transaction binding the contract method 0x4e1e2fe1.
//
// Solidity: function pay_billing() returns()
func (_OrderBase *OrderBaseSession) PayBilling() (*types.Transaction, error) {
	return _OrderBase.Contract.PayBilling(&_OrderBase.TransactOpts)
}

// PayBilling is a paid mutator transaction binding the contract method 0x4e1e2fe1.
//
// Solidity: function pay_billing() returns()
func (_OrderBase *OrderBaseTransactorSession) PayBilling() (*types.Transaction, error) {
	return _OrderBase.Contract.PayBilling(&_OrderBase.TransactOpts)
}

// Quote is a paid mutator transaction binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 p_cpu, uint256 p_memory, uint256 p_storage) returns(uint256)
func (_OrderBase *OrderBaseTransactor) Quote(opts *bind.TransactOpts, p_cpu *big.Int, p_memory *big.Int, p_storage *big.Int) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "quote", p_cpu, p_memory, p_storage)
}

// Quote is a paid mutator transaction binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 p_cpu, uint256 p_memory, uint256 p_storage) returns(uint256)
func (_OrderBase *OrderBaseSession) Quote(p_cpu *big.Int, p_memory *big.Int, p_storage *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.Quote(&_OrderBase.TransactOpts, p_cpu, p_memory, p_storage)
}

// Quote is a paid mutator transaction binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 p_cpu, uint256 p_memory, uint256 p_storage) returns(uint256)
func (_OrderBase *OrderBaseTransactorSession) Quote(p_cpu *big.Int, p_memory *big.Int, p_storage *big.Int) (*types.Transaction, error) {
	return _OrderBase.Contract.Quote(&_OrderBase.TransactOpts, p_cpu, p_memory, p_storage)
}

// SubmitServerUri is a paid mutator transaction binding the contract method 0x9a4339fe.
//
// Solidity: function submit_server_uri(string uri) returns()
func (_OrderBase *OrderBaseTransactor) SubmitServerUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "submit_server_uri", uri)
}

// SubmitServerUri is a paid mutator transaction binding the contract method 0x9a4339fe.
//
// Solidity: function submit_server_uri(string uri) returns()
func (_OrderBase *OrderBaseSession) SubmitServerUri(uri string) (*types.Transaction, error) {
	return _OrderBase.Contract.SubmitServerUri(&_OrderBase.TransactOpts, uri)
}

// SubmitServerUri is a paid mutator transaction binding the contract method 0x9a4339fe.
//
// Solidity: function submit_server_uri(string uri) returns()
func (_OrderBase *OrderBaseTransactorSession) SubmitServerUri(uri string) (*types.Transaction, error) {
	return _OrderBase.Contract.SubmitServerUri(&_OrderBase.TransactOpts, uri)
}

// UpdateDeployment is a paid mutator transaction binding the contract method 0xd26d6c73.
//
// Solidity: function update_deployment(uint256 cpu_, uint256 memory_, uint256 storage_, string uri_) returns()
func (_OrderBase *OrderBaseTransactor) UpdateDeployment(opts *bind.TransactOpts, cpu_ *big.Int, memory_ *big.Int, storage_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "update_deployment", cpu_, memory_, storage_, uri_)
}

// UpdateDeployment is a paid mutator transaction binding the contract method 0xd26d6c73.
//
// Solidity: function update_deployment(uint256 cpu_, uint256 memory_, uint256 storage_, string uri_) returns()
func (_OrderBase *OrderBaseSession) UpdateDeployment(cpu_ *big.Int, memory_ *big.Int, storage_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _OrderBase.Contract.UpdateDeployment(&_OrderBase.TransactOpts, cpu_, memory_, storage_, uri_)
}

// UpdateDeployment is a paid mutator transaction binding the contract method 0xd26d6c73.
//
// Solidity: function update_deployment(uint256 cpu_, uint256 memory_, uint256 storage_, string uri_) returns()
func (_OrderBase *OrderBaseTransactorSession) UpdateDeployment(cpu_ *big.Int, memory_ *big.Int, storage_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _OrderBase.Contract.UpdateDeployment(&_OrderBase.TransactOpts, cpu_, memory_, storage_, uri_)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x5cd30cbf.
//
// Solidity: function withdraw_fund() returns()
func (_OrderBase *OrderBaseTransactor) WithdrawFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBase.contract.Transact(opts, "withdraw_fund")
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x5cd30cbf.
//
// Solidity: function withdraw_fund() returns()
func (_OrderBase *OrderBaseSession) WithdrawFund() (*types.Transaction, error) {
	return _OrderBase.Contract.WithdrawFund(&_OrderBase.TransactOpts)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x5cd30cbf.
//
// Solidity: function withdraw_fund() returns()
func (_OrderBase *OrderBaseTransactorSession) WithdrawFund() (*types.Transaction, error) {
	return _OrderBase.Contract.WithdrawFund(&_OrderBase.TransactOpts)
}

// OrderBaseCanQuoteIterator is returned from FilterCanQuote and is used to iterate over the raw logs and unpacked data for CanQuote events raised by the OrderBase contract.
type OrderBaseCanQuoteIterator struct {
	Event *OrderBaseCanQuote // Event containing the contract specifics and raw log

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
func (it *OrderBaseCanQuoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseCanQuote)
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
		it.Event = new(OrderBaseCanQuote)
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
func (it *OrderBaseCanQuoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseCanQuoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseCanQuote represents a CanQuote event raised by the OrderBase contract.
type OrderBaseCanQuote struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCanQuote is a free log retrieval operation binding the contract event 0xb35acc2c23ccc06831d6d983ba73f53023f6c21e5015e3b2121eeae8fe5c984c.
//
// Solidity: event CanQuote()
func (_OrderBase *OrderBaseFilterer) FilterCanQuote(opts *bind.FilterOpts) (*OrderBaseCanQuoteIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "CanQuote")
	if err != nil {
		return nil, err
	}
	return &OrderBaseCanQuoteIterator{contract: _OrderBase.contract, event: "CanQuote", logs: logs, sub: sub}, nil
}

// WatchCanQuote is a free log subscription operation binding the contract event 0xb35acc2c23ccc06831d6d983ba73f53023f6c21e5015e3b2121eeae8fe5c984c.
//
// Solidity: event CanQuote()
func (_OrderBase *OrderBaseFilterer) WatchCanQuote(opts *bind.WatchOpts, sink chan<- *OrderBaseCanQuote) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "CanQuote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseCanQuote)
				if err := _OrderBase.contract.UnpackLog(event, "CanQuote", log); err != nil {
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

// ParseCanQuote is a log parse operation binding the contract event 0xb35acc2c23ccc06831d6d983ba73f53023f6c21e5015e3b2121eeae8fe5c984c.
//
// Solidity: event CanQuote()
func (_OrderBase *OrderBaseFilterer) ParseCanQuote(log types.Log) (*OrderBaseCanQuote, error) {
	event := new(OrderBaseCanQuote)
	if err := _OrderBase.contract.UnpackLog(event, "CanQuote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseChooseQuoteIterator is returned from FilterChooseQuote and is used to iterate over the raw logs and unpacked data for ChooseQuote events raised by the OrderBase contract.
type OrderBaseChooseQuoteIterator struct {
	Event *OrderBaseChooseQuote // Event containing the contract specifics and raw log

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
func (it *OrderBaseChooseQuoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseChooseQuote)
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
		it.Event = new(OrderBaseChooseQuote)
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
func (it *OrderBaseChooseQuoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseChooseQuoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseChooseQuote represents a ChooseQuote event raised by the OrderBase contract.
type OrderBaseChooseQuote struct {
	Provider   common.Address
	FinalPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChooseQuote is a free log retrieval operation binding the contract event 0x40b763fcbe2138c8d50a909799fd5f1690100032212cbf680152269eabab7831.
//
// Solidity: event ChooseQuote(address indexed provider, uint256 indexed final_price)
func (_OrderBase *OrderBaseFilterer) FilterChooseQuote(opts *bind.FilterOpts, provider []common.Address, final_price []*big.Int) (*OrderBaseChooseQuoteIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var final_priceRule []interface{}
	for _, final_priceItem := range final_price {
		final_priceRule = append(final_priceRule, final_priceItem)
	}

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "ChooseQuote", providerRule, final_priceRule)
	if err != nil {
		return nil, err
	}
	return &OrderBaseChooseQuoteIterator{contract: _OrderBase.contract, event: "ChooseQuote", logs: logs, sub: sub}, nil
}

// WatchChooseQuote is a free log subscription operation binding the contract event 0x40b763fcbe2138c8d50a909799fd5f1690100032212cbf680152269eabab7831.
//
// Solidity: event ChooseQuote(address indexed provider, uint256 indexed final_price)
func (_OrderBase *OrderBaseFilterer) WatchChooseQuote(opts *bind.WatchOpts, sink chan<- *OrderBaseChooseQuote, provider []common.Address, final_price []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var final_priceRule []interface{}
	for _, final_priceItem := range final_price {
		final_priceRule = append(final_priceRule, final_priceItem)
	}

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "ChooseQuote", providerRule, final_priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseChooseQuote)
				if err := _OrderBase.contract.UnpackLog(event, "ChooseQuote", log); err != nil {
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

// ParseChooseQuote is a log parse operation binding the contract event 0x40b763fcbe2138c8d50a909799fd5f1690100032212cbf680152269eabab7831.
//
// Solidity: event ChooseQuote(address indexed provider, uint256 indexed final_price)
func (_OrderBase *OrderBaseFilterer) ParseChooseQuote(log types.Log) (*OrderBaseChooseQuote, error) {
	event := new(OrderBaseChooseQuote)
	if err := _OrderBase.contract.UnpackLog(event, "ChooseQuote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseDeployMentUpdatedIterator is returned from FilterDeployMentUpdated and is used to iterate over the raw logs and unpacked data for DeployMentUpdated events raised by the OrderBase contract.
type OrderBaseDeployMentUpdatedIterator struct {
	Event *OrderBaseDeployMentUpdated // Event containing the contract specifics and raw log

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
func (it *OrderBaseDeployMentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseDeployMentUpdated)
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
		it.Event = new(OrderBaseDeployMentUpdated)
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
func (it *OrderBaseDeployMentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseDeployMentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseDeployMentUpdated represents a DeployMentUpdated event raised by the OrderBase contract.
type OrderBaseDeployMentUpdated struct {
	OldCpu     *big.Int
	OldMemory  *big.Int
	OldStorage *big.Int
	NewCpu     *big.Int
	NewMemory  *big.Int
	NewStorage *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeployMentUpdated is a free log retrieval operation binding the contract event 0x8c679db344e603e7397dad8cd5778fc955e2ae22a43add0bd2cf20ad9c7183cb.
//
// Solidity: event DeployMentUpdated(uint256 old_cpu, uint256 old_memory, uint256 old_storage, uint256 new_cpu, uint256 new_memory, uint256 new_storage)
func (_OrderBase *OrderBaseFilterer) FilterDeployMentUpdated(opts *bind.FilterOpts) (*OrderBaseDeployMentUpdatedIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "DeployMentUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderBaseDeployMentUpdatedIterator{contract: _OrderBase.contract, event: "DeployMentUpdated", logs: logs, sub: sub}, nil
}

// WatchDeployMentUpdated is a free log subscription operation binding the contract event 0x8c679db344e603e7397dad8cd5778fc955e2ae22a43add0bd2cf20ad9c7183cb.
//
// Solidity: event DeployMentUpdated(uint256 old_cpu, uint256 old_memory, uint256 old_storage, uint256 new_cpu, uint256 new_memory, uint256 new_storage)
func (_OrderBase *OrderBaseFilterer) WatchDeployMentUpdated(opts *bind.WatchOpts, sink chan<- *OrderBaseDeployMentUpdated) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "DeployMentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseDeployMentUpdated)
				if err := _OrderBase.contract.UnpackLog(event, "DeployMentUpdated", log); err != nil {
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

// ParseDeployMentUpdated is a log parse operation binding the contract event 0x8c679db344e603e7397dad8cd5778fc955e2ae22a43add0bd2cf20ad9c7183cb.
//
// Solidity: event DeployMentUpdated(uint256 old_cpu, uint256 old_memory, uint256 old_storage, uint256 new_cpu, uint256 new_memory, uint256 new_storage)
func (_OrderBase *OrderBaseFilterer) ParseDeployMentUpdated(log types.Log) (*OrderBaseDeployMentUpdated, error) {
	event := new(OrderBaseDeployMentUpdated)
	if err := _OrderBase.contract.UnpackLog(event, "DeployMentUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseDepositBalanceIterator is returned from FilterDepositBalance and is used to iterate over the raw logs and unpacked data for DepositBalance events raised by the OrderBase contract.
type OrderBaseDepositBalanceIterator struct {
	Event *OrderBaseDepositBalance // Event containing the contract specifics and raw log

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
func (it *OrderBaseDepositBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseDepositBalance)
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
		it.Event = new(OrderBaseDepositBalance)
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
func (it *OrderBaseDepositBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseDepositBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseDepositBalance represents a DepositBalance event raised by the OrderBase contract.
type OrderBaseDepositBalance struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositBalance is a free log retrieval operation binding the contract event 0x8d36bfbc8c8f391b6058824dfedd0b18c4f772f7f0dc517d87b52d957dc85b54.
//
// Solidity: event DepositBalance(uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) FilterDepositBalance(opts *bind.FilterOpts, amount []*big.Int) (*OrderBaseDepositBalanceIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "DepositBalance", amountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBaseDepositBalanceIterator{contract: _OrderBase.contract, event: "DepositBalance", logs: logs, sub: sub}, nil
}

// WatchDepositBalance is a free log subscription operation binding the contract event 0x8d36bfbc8c8f391b6058824dfedd0b18c4f772f7f0dc517d87b52d957dc85b54.
//
// Solidity: event DepositBalance(uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) WatchDepositBalance(opts *bind.WatchOpts, sink chan<- *OrderBaseDepositBalance, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "DepositBalance", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseDepositBalance)
				if err := _OrderBase.contract.UnpackLog(event, "DepositBalance", log); err != nil {
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

// ParseDepositBalance is a log parse operation binding the contract event 0x8d36bfbc8c8f391b6058824dfedd0b18c4f772f7f0dc517d87b52d957dc85b54.
//
// Solidity: event DepositBalance(uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) ParseDepositBalance(log types.Log) (*OrderBaseDepositBalance, error) {
	event := new(OrderBaseDepositBalance)
	if err := _OrderBase.contract.UnpackLog(event, "DepositBalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseOrderCreateIterator is returned from FilterOrderCreate and is used to iterate over the raw logs and unpacked data for OrderCreate events raised by the OrderBase contract.
type OrderBaseOrderCreateIterator struct {
	Event *OrderBaseOrderCreate // Event containing the contract specifics and raw log

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
func (it *OrderBaseOrderCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseOrderCreate)
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
		it.Event = new(OrderBaseOrderCreate)
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
func (it *OrderBaseOrderCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseOrderCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseOrderCreate represents a OrderCreate event raised by the OrderBase contract.
type OrderBaseOrderCreate struct {
	Owner       common.Address
	Cpu         *big.Int
	Memory      *big.Int
	Storage     *big.Int
	Cert        *big.Int
	Sdl         *big.Int
	OrderNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderCreate is a free log retrieval operation binding the contract event 0x54c2886880d2438c34b17297401499ffbc5e9663e148d1703e4658a4a0eefc95.
//
// Solidity: event OrderCreate(address owner_, uint256 cpu, uint256 memory_, uint256 storage_, uint256 cert, uint256 sdl, uint256 order_number)
func (_OrderBase *OrderBaseFilterer) FilterOrderCreate(opts *bind.FilterOpts) (*OrderBaseOrderCreateIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "OrderCreate")
	if err != nil {
		return nil, err
	}
	return &OrderBaseOrderCreateIterator{contract: _OrderBase.contract, event: "OrderCreate", logs: logs, sub: sub}, nil
}

// WatchOrderCreate is a free log subscription operation binding the contract event 0x54c2886880d2438c34b17297401499ffbc5e9663e148d1703e4658a4a0eefc95.
//
// Solidity: event OrderCreate(address owner_, uint256 cpu, uint256 memory_, uint256 storage_, uint256 cert, uint256 sdl, uint256 order_number)
func (_OrderBase *OrderBaseFilterer) WatchOrderCreate(opts *bind.WatchOpts, sink chan<- *OrderBaseOrderCreate) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "OrderCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseOrderCreate)
				if err := _OrderBase.contract.UnpackLog(event, "OrderCreate", log); err != nil {
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

// ParseOrderCreate is a log parse operation binding the contract event 0x54c2886880d2438c34b17297401499ffbc5e9663e148d1703e4658a4a0eefc95.
//
// Solidity: event OrderCreate(address owner_, uint256 cpu, uint256 memory_, uint256 storage_, uint256 cert, uint256 sdl, uint256 order_number)
func (_OrderBase *OrderBaseFilterer) ParseOrderCreate(log types.Log) (*OrderBaseOrderCreate, error) {
	event := new(OrderBaseOrderCreate)
	if err := _OrderBase.contract.UnpackLog(event, "OrderCreate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseOrderEndedIterator is returned from FilterOrderEnded and is used to iterate over the raw logs and unpacked data for OrderEnded events raised by the OrderBase contract.
type OrderBaseOrderEndedIterator struct {
	Event *OrderBaseOrderEnded // Event containing the contract specifics and raw log

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
func (it *OrderBaseOrderEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseOrderEnded)
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
		it.Event = new(OrderBaseOrderEnded)
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
func (it *OrderBaseOrderEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseOrderEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseOrderEnded represents a OrderEnded event raised by the OrderBase contract.
type OrderBaseOrderEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrderEnded is a free log retrieval operation binding the contract event 0xe2fb59d1a8aac83c47f73e8b5c52024d5729c91f381a061d5117dacd4f06c1eb.
//
// Solidity: event OrderEnded()
func (_OrderBase *OrderBaseFilterer) FilterOrderEnded(opts *bind.FilterOpts) (*OrderBaseOrderEndedIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "OrderEnded")
	if err != nil {
		return nil, err
	}
	return &OrderBaseOrderEndedIterator{contract: _OrderBase.contract, event: "OrderEnded", logs: logs, sub: sub}, nil
}

// WatchOrderEnded is a free log subscription operation binding the contract event 0xe2fb59d1a8aac83c47f73e8b5c52024d5729c91f381a061d5117dacd4f06c1eb.
//
// Solidity: event OrderEnded()
func (_OrderBase *OrderBaseFilterer) WatchOrderEnded(opts *bind.WatchOpts, sink chan<- *OrderBaseOrderEnded) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "OrderEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseOrderEnded)
				if err := _OrderBase.contract.UnpackLog(event, "OrderEnded", log); err != nil {
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

// ParseOrderEnded is a log parse operation binding the contract event 0xe2fb59d1a8aac83c47f73e8b5c52024d5729c91f381a061d5117dacd4f06c1eb.
//
// Solidity: event OrderEnded()
func (_OrderBase *OrderBaseFilterer) ParseOrderEnded(log types.Log) (*OrderBaseOrderEnded, error) {
	event := new(OrderBaseOrderEnded)
	if err := _OrderBase.contract.UnpackLog(event, "OrderEnded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBasePayBillIterator is returned from FilterPayBill and is used to iterate over the raw logs and unpacked data for PayBill events raised by the OrderBase contract.
type OrderBasePayBillIterator struct {
	Event *OrderBasePayBill // Event containing the contract specifics and raw log

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
func (it *OrderBasePayBillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBasePayBill)
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
		it.Event = new(OrderBasePayBill)
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
func (it *OrderBasePayBillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBasePayBillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBasePayBill represents a PayBill event raised by the OrderBase contract.
type OrderBasePayBill struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayBill is a free log retrieval operation binding the contract event 0xc509241168a93b52647ae47383362edfc1e268f1d0ae8485703e21fedcdf2709.
//
// Solidity: event PayBill(address indexed provider, uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) FilterPayBill(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*OrderBasePayBillIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "PayBill", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBasePayBillIterator{contract: _OrderBase.contract, event: "PayBill", logs: logs, sub: sub}, nil
}

// WatchPayBill is a free log subscription operation binding the contract event 0xc509241168a93b52647ae47383362edfc1e268f1d0ae8485703e21fedcdf2709.
//
// Solidity: event PayBill(address indexed provider, uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) WatchPayBill(opts *bind.WatchOpts, sink chan<- *OrderBasePayBill, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "PayBill", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBasePayBill)
				if err := _OrderBase.contract.UnpackLog(event, "PayBill", log); err != nil {
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

// ParsePayBill is a log parse operation binding the contract event 0xc509241168a93b52647ae47383362edfc1e268f1d0ae8485703e21fedcdf2709.
//
// Solidity: event PayBill(address indexed provider, uint256 indexed amount)
func (_OrderBase *OrderBaseFilterer) ParsePayBill(log types.Log) (*OrderBasePayBill, error) {
	event := new(OrderBasePayBill)
	if err := _OrderBase.contract.UnpackLog(event, "PayBill", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseQuoteIterator is returned from FilterQuote and is used to iterate over the raw logs and unpacked data for Quote events raised by the OrderBase contract.
type OrderBaseQuoteIterator struct {
	Event *OrderBaseQuote // Event containing the contract specifics and raw log

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
func (it *OrderBaseQuoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseQuote)
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
		it.Event = new(OrderBaseQuote)
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
func (it *OrderBaseQuoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseQuoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseQuote represents a Quote event raised by the OrderBase contract.
type OrderBaseQuote struct {
	Provider     common.Address
	CpuPrice     *big.Int
	MemoryPrice  *big.Int
	StoragePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuote is a free log retrieval operation binding the contract event 0x28be2c5e63173eb72b4c1d64e411524d61c8c22cf40149a0d1320e17e25339d0.
//
// Solidity: event Quote(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseFilterer) FilterQuote(opts *bind.FilterOpts) (*OrderBaseQuoteIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "Quote")
	if err != nil {
		return nil, err
	}
	return &OrderBaseQuoteIterator{contract: _OrderBase.contract, event: "Quote", logs: logs, sub: sub}, nil
}

// WatchQuote is a free log subscription operation binding the contract event 0x28be2c5e63173eb72b4c1d64e411524d61c8c22cf40149a0d1320e17e25339d0.
//
// Solidity: event Quote(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseFilterer) WatchQuote(opts *bind.WatchOpts, sink chan<- *OrderBaseQuote) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "Quote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseQuote)
				if err := _OrderBase.contract.UnpackLog(event, "Quote", log); err != nil {
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

// ParseQuote is a log parse operation binding the contract event 0x28be2c5e63173eb72b4c1d64e411524d61c8c22cf40149a0d1320e17e25339d0.
//
// Solidity: event Quote(address provider, uint256 cpu_price, uint256 memory_price, uint256 storage_price)
func (_OrderBase *OrderBaseFilterer) ParseQuote(log types.Log) (*OrderBaseQuote, error) {
	event := new(OrderBaseQuote)
	if err := _OrderBase.contract.UnpackLog(event, "Quote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseUpdateSDLIterator is returned from FilterUpdateSDL and is used to iterate over the raw logs and unpacked data for UpdateSDL events raised by the OrderBase contract.
type OrderBaseUpdateSDLIterator struct {
	Event *OrderBaseUpdateSDL // Event containing the contract specifics and raw log

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
func (it *OrderBaseUpdateSDLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseUpdateSDL)
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
		it.Event = new(OrderBaseUpdateSDL)
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
func (it *OrderBaseUpdateSDLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseUpdateSDLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseUpdateSDL represents a UpdateSDL event raised by the OrderBase contract.
type OrderBaseUpdateSDL struct {
	NewSdlTrxId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateSDL is a free log retrieval operation binding the contract event 0xea8d9427410dd8e182bfd394ce4dc513628f622b93ca7e32a2713dcfbe0b5b04.
//
// Solidity: event UpdateSDL(uint256 new_sdl_trx_id)
func (_OrderBase *OrderBaseFilterer) FilterUpdateSDL(opts *bind.FilterOpts) (*OrderBaseUpdateSDLIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "UpdateSDL")
	if err != nil {
		return nil, err
	}
	return &OrderBaseUpdateSDLIterator{contract: _OrderBase.contract, event: "UpdateSDL", logs: logs, sub: sub}, nil
}

// WatchUpdateSDL is a free log subscription operation binding the contract event 0xea8d9427410dd8e182bfd394ce4dc513628f622b93ca7e32a2713dcfbe0b5b04.
//
// Solidity: event UpdateSDL(uint256 new_sdl_trx_id)
func (_OrderBase *OrderBaseFilterer) WatchUpdateSDL(opts *bind.WatchOpts, sink chan<- *OrderBaseUpdateSDL) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "UpdateSDL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseUpdateSDL)
				if err := _OrderBase.contract.UnpackLog(event, "UpdateSDL", log); err != nil {
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

// ParseUpdateSDL is a log parse operation binding the contract event 0xea8d9427410dd8e182bfd394ce4dc513628f622b93ca7e32a2713dcfbe0b5b04.
//
// Solidity: event UpdateSDL(uint256 new_sdl_trx_id)
func (_OrderBase *OrderBaseFilterer) ParseUpdateSDL(log types.Log) (*OrderBaseUpdateSDL, error) {
	event := new(OrderBaseUpdateSDL)
	if err := _OrderBase.contract.UnpackLog(event, "UpdateSDL", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderBaseUserCancelOrderIterator is returned from FilterUserCancelOrder and is used to iterate over the raw logs and unpacked data for UserCancelOrder events raised by the OrderBase contract.
type OrderBaseUserCancelOrderIterator struct {
	Event *OrderBaseUserCancelOrder // Event containing the contract specifics and raw log

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
func (it *OrderBaseUserCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBaseUserCancelOrder)
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
		it.Event = new(OrderBaseUserCancelOrder)
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
func (it *OrderBaseUserCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBaseUserCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBaseUserCancelOrder represents a UserCancelOrder event raised by the OrderBase contract.
type OrderBaseUserCancelOrder struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUserCancelOrder is a free log retrieval operation binding the contract event 0x7f33f96d4f3b2c29007fe630eedb88481977bf9038c3c7d5c87bdd10fc11730d.
//
// Solidity: event UserCancelOrder()
func (_OrderBase *OrderBaseFilterer) FilterUserCancelOrder(opts *bind.FilterOpts) (*OrderBaseUserCancelOrderIterator, error) {

	logs, sub, err := _OrderBase.contract.FilterLogs(opts, "UserCancelOrder")
	if err != nil {
		return nil, err
	}
	return &OrderBaseUserCancelOrderIterator{contract: _OrderBase.contract, event: "UserCancelOrder", logs: logs, sub: sub}, nil
}

// WatchUserCancelOrder is a free log subscription operation binding the contract event 0x7f33f96d4f3b2c29007fe630eedb88481977bf9038c3c7d5c87bdd10fc11730d.
//
// Solidity: event UserCancelOrder()
func (_OrderBase *OrderBaseFilterer) WatchUserCancelOrder(opts *bind.WatchOpts, sink chan<- *OrderBaseUserCancelOrder) (event.Subscription, error) {

	logs, sub, err := _OrderBase.contract.WatchLogs(opts, "UserCancelOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBaseUserCancelOrder)
				if err := _OrderBase.contract.UnpackLog(event, "UserCancelOrder", log); err != nil {
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

// ParseUserCancelOrder is a log parse operation binding the contract event 0x7f33f96d4f3b2c29007fe630eedb88481977bf9038c3c7d5c87bdd10fc11730d.
//
// Solidity: event UserCancelOrder()
func (_OrderBase *OrderBaseFilterer) ParseUserCancelOrder(log types.Log) (*OrderBaseUserCancelOrder, error) {
	event := new(OrderBaseUserCancelOrder)
	if err := _OrderBase.contract.UnpackLog(event, "UserCancelOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

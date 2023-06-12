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

// ProviderFactoryproviderInfos is an auto generated low-level Go binding around an user-defined struct.
type ProviderFactoryproviderInfos struct {
	ProviderContract common.Address
	Info             providerInfo
	MarginAmount     *big.Int
	Audits           []common.Address
}

// poaResource is an auto generated low-level Go binding around an user-defined struct.
type poaResource struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}

// providerInfo is an auto generated low-level Go binding around an user-defined struct.
type providerInfo struct {
	Total             poaResource
	Used              poaResource
	Lock              poaResource
	Challenge         bool
	State             uint8
	Owner             common.Address
	Region            string
	Info              string
	LastChallengeTime *big.Int
	LastMarginTime    *big.Int
}

// ProviderFactoryABI is the input ABI used to generate the binding from.
const ProviderFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ProviderCreate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addMargin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auditor_factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"}],\"name\":\"calcProviderAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_audit_factory\",\"type\":\"address\"}],\"name\":\"changeAuditorFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_cpu_decimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_memory_decimal\",\"type\":\"uint256\"}],\"name\":\"changeDecimal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_order_factory\",\"type\":\"address\"}],\"name\":\"changeOrderFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_max\",\"type\":\"uint256\"}],\"name\":\"changeProviderLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lock_time\",\"type\":\"uint256\"}],\"name\":\"changeProviderLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"changeProviderResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whether_start\",\"type\":\"bool\"}],\"name\":\"changeProviderState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"changeProviderUsedResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_punish_address\",\"type\":\"address\"}],\"name\":\"changePunishAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_punish_start_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_punish_interval\",\"type\":\"uint256\"}],\"name\":\"changePunishParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_punish_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_punish_all_percent\",\"type\":\"uint256\"}],\"name\":\"changePunishPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"name\":\"consumeResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider_info\",\"type\":\"string\"}],\"name\":\"createNewProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal_cpu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal_memory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getProvideContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getProvideResource\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getProvideTotalResource\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getProviderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider_contract\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"used\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"lock\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"challenge\",\"type\":\"bool\"},{\"internalType\":\"enumProviderState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"last_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_margin_time\",\"type\":\"uint256\"}],\"internalType\":\"structproviderInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"margin_amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"audits\",\"type\":\"address[]\"}],\"internalType\":\"structProviderFactory.providerInfos[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderInfoLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider_contract\",\"type\":\"address\"}],\"name\":\"getProviderSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider_contract\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"used\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"lock\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"challenge\",\"type\":\"bool\"},{\"internalType\":\"enumProviderState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"last_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_margin_time\",\"type\":\"uint256\"}],\"internalType\":\"structproviderInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"margin_amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"audits\",\"type\":\"address[]\"}],\"internalType\":\"structProviderFactory.providerInfos\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"punish_amount\",\"type\":\"uint256\"}],\"name\":\"getPunishAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"internalType\":\"structpoaResource\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_value_tobe_provider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min_value_tobe_provider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"order_factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider_lock_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providers\",\"outputs\":[{\"internalType\":\"contractIProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_all_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_start_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mem_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"name\":\"recoverResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"removePunishList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_all\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_used\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memory_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storage_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_provider\",\"type\":\"address\"}],\"name\":\"tryPunish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"val_factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_owner\",\"type\":\"address\"}],\"name\":\"whetherCanPOR\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProviderFactory is an auto generated Go binding around an Ethereum contract.
type ProviderFactory struct {
	ProviderFactoryCaller     // Read-only binding to the contract
	ProviderFactoryTransactor // Write-only binding to the contract
	ProviderFactoryFilterer   // Log filterer for contract events
}

// ProviderFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderFactorySession struct {
	Contract     *ProviderFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderFactoryCallerSession struct {
	Contract *ProviderFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProviderFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderFactoryTransactorSession struct {
	Contract     *ProviderFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProviderFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderFactoryRaw struct {
	Contract *ProviderFactory // Generic contract binding to access the raw methods on
}

// ProviderFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderFactoryCallerRaw struct {
	Contract *ProviderFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderFactoryTransactorRaw struct {
	Contract *ProviderFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderFactory creates a new instance of ProviderFactory, bound to a specific deployed contract.
func NewProviderFactory(address common.Address, backend bind.ContractBackend) (*ProviderFactory, error) {
	contract, err := bindProviderFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderFactory{ProviderFactoryCaller: ProviderFactoryCaller{contract: contract}, ProviderFactoryTransactor: ProviderFactoryTransactor{contract: contract}, ProviderFactoryFilterer: ProviderFactoryFilterer{contract: contract}}, nil
}

// NewProviderFactoryCaller creates a new read-only instance of ProviderFactory, bound to a specific deployed contract.
func NewProviderFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProviderFactoryCaller, error) {
	contract, err := bindProviderFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderFactoryCaller{contract: contract}, nil
}

// NewProviderFactoryTransactor creates a new write-only instance of ProviderFactory, bound to a specific deployed contract.
func NewProviderFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderFactoryTransactor, error) {
	contract, err := bindProviderFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderFactoryTransactor{contract: contract}, nil
}

// NewProviderFactoryFilterer creates a new log filterer instance of ProviderFactory, bound to a specific deployed contract.
func NewProviderFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderFactoryFilterer, error) {
	contract, err := bindProviderFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderFactoryFilterer{contract: contract}, nil
}

// bindProviderFactory binds a generic wrapper to an already deployed contract.
func bindProviderFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderFactory *ProviderFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderFactory.Contract.ProviderFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderFactory *ProviderFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ProviderFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderFactory *ProviderFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ProviderFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderFactory *ProviderFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderFactory *ProviderFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderFactory *ProviderFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ProviderFactory *ProviderFactorySession) Admin() (common.Address, error) {
	return _ProviderFactory.Contract.Admin(&_ProviderFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) Admin() (common.Address, error) {
	return _ProviderFactory.Contract.Admin(&_ProviderFactory.CallOpts)
}

// AuditorFactory is a free data retrieval call binding the contract method 0x5ec35663.
//
// Solidity: function auditor_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) AuditorFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "auditor_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuditorFactory is a free data retrieval call binding the contract method 0x5ec35663.
//
// Solidity: function auditor_factory() view returns(address)
func (_ProviderFactory *ProviderFactorySession) AuditorFactory() (common.Address, error) {
	return _ProviderFactory.Contract.AuditorFactory(&_ProviderFactory.CallOpts)
}

// AuditorFactory is a free data retrieval call binding the contract method 0x5ec35663.
//
// Solidity: function auditor_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) AuditorFactory() (common.Address, error) {
	return _ProviderFactory.Contract.AuditorFactory(&_ProviderFactory.CallOpts)
}

// CalcProviderAmount is a free data retrieval call binding the contract method 0xa6817371.
//
// Solidity: function calcProviderAmount(uint256 cpu_count, uint256 memory_count) view returns(uint256, uint256)
func (_ProviderFactory *ProviderFactoryCaller) CalcProviderAmount(opts *bind.CallOpts, cpu_count *big.Int, memory_count *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "calcProviderAmount", cpu_count, memory_count)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CalcProviderAmount is a free data retrieval call binding the contract method 0xa6817371.
//
// Solidity: function calcProviderAmount(uint256 cpu_count, uint256 memory_count) view returns(uint256, uint256)
func (_ProviderFactory *ProviderFactorySession) CalcProviderAmount(cpu_count *big.Int, memory_count *big.Int) (*big.Int, *big.Int, error) {
	return _ProviderFactory.Contract.CalcProviderAmount(&_ProviderFactory.CallOpts, cpu_count, memory_count)
}

// CalcProviderAmount is a free data retrieval call binding the contract method 0xa6817371.
//
// Solidity: function calcProviderAmount(uint256 cpu_count, uint256 memory_count) view returns(uint256, uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) CalcProviderAmount(cpu_count *big.Int, memory_count *big.Int) (*big.Int, *big.Int, error) {
	return _ProviderFactory.Contract.CalcProviderAmount(&_ProviderFactory.CallOpts, cpu_count, memory_count)
}

// DecimalCpu is a free data retrieval call binding the contract method 0x8026a5a2.
//
// Solidity: function decimal_cpu() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) DecimalCpu(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "decimal_cpu")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalCpu is a free data retrieval call binding the contract method 0x8026a5a2.
//
// Solidity: function decimal_cpu() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) DecimalCpu() (*big.Int, error) {
	return _ProviderFactory.Contract.DecimalCpu(&_ProviderFactory.CallOpts)
}

// DecimalCpu is a free data retrieval call binding the contract method 0x8026a5a2.
//
// Solidity: function decimal_cpu() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) DecimalCpu() (*big.Int, error) {
	return _ProviderFactory.Contract.DecimalCpu(&_ProviderFactory.CallOpts)
}

// DecimalMemory is a free data retrieval call binding the contract method 0x422d7be8.
//
// Solidity: function decimal_memory() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) DecimalMemory(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "decimal_memory")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalMemory is a free data retrieval call binding the contract method 0x422d7be8.
//
// Solidity: function decimal_memory() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) DecimalMemory() (*big.Int, error) {
	return _ProviderFactory.Contract.DecimalMemory(&_ProviderFactory.CallOpts)
}

// DecimalMemory is a free data retrieval call binding the contract method 0x422d7be8.
//
// Solidity: function decimal_memory() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) DecimalMemory() (*big.Int, error) {
	return _ProviderFactory.Contract.DecimalMemory(&_ProviderFactory.CallOpts)
}

// GetProvideContract is a free data retrieval call binding the contract method 0x97ddf78c.
//
// Solidity: function getProvideContract(address account) view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) GetProvideContract(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProvideContract", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProvideContract is a free data retrieval call binding the contract method 0x97ddf78c.
//
// Solidity: function getProvideContract(address account) view returns(address)
func (_ProviderFactory *ProviderFactorySession) GetProvideContract(account common.Address) (common.Address, error) {
	return _ProviderFactory.Contract.GetProvideContract(&_ProviderFactory.CallOpts, account)
}

// GetProvideContract is a free data retrieval call binding the contract method 0x97ddf78c.
//
// Solidity: function getProvideContract(address account) view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) GetProvideContract(account common.Address) (common.Address, error) {
	return _ProviderFactory.Contract.GetProvideContract(&_ProviderFactory.CallOpts, account)
}

// GetProvideResource is a free data retrieval call binding the contract method 0x4c01448b.
//
// Solidity: function getProvideResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCaller) GetProvideResource(opts *bind.CallOpts, account common.Address) (poaResource, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProvideResource", account)

	if err != nil {
		return *new(poaResource), err
	}

	out0 := *abi.ConvertType(out[0], new(poaResource)).(*poaResource)

	return out0, err

}

// GetProvideResource is a free data retrieval call binding the contract method 0x4c01448b.
//
// Solidity: function getProvideResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactorySession) GetProvideResource(account common.Address) (poaResource, error) {
	return _ProviderFactory.Contract.GetProvideResource(&_ProviderFactory.CallOpts, account)
}

// GetProvideResource is a free data retrieval call binding the contract method 0x4c01448b.
//
// Solidity: function getProvideResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCallerSession) GetProvideResource(account common.Address) (poaResource, error) {
	return _ProviderFactory.Contract.GetProvideResource(&_ProviderFactory.CallOpts, account)
}

// GetProvideTotalResource is a free data retrieval call binding the contract method 0xd732e0f4.
//
// Solidity: function getProvideTotalResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCaller) GetProvideTotalResource(opts *bind.CallOpts, account common.Address) (poaResource, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProvideTotalResource", account)

	if err != nil {
		return *new(poaResource), err
	}

	out0 := *abi.ConvertType(out[0], new(poaResource)).(*poaResource)

	return out0, err

}

// GetProvideTotalResource is a free data retrieval call binding the contract method 0xd732e0f4.
//
// Solidity: function getProvideTotalResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactorySession) GetProvideTotalResource(account common.Address) (poaResource, error) {
	return _ProviderFactory.Contract.GetProvideTotalResource(&_ProviderFactory.CallOpts, account)
}

// GetProvideTotalResource is a free data retrieval call binding the contract method 0xd732e0f4.
//
// Solidity: function getProvideTotalResource(address account) view returns((uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCallerSession) GetProvideTotalResource(account common.Address) (poaResource, error) {
	return _ProviderFactory.Contract.GetProvideTotalResource(&_ProviderFactory.CallOpts, account)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(uint256 start, uint256 limit) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[])[])
func (_ProviderFactory *ProviderFactoryCaller) GetProviderInfo(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]ProviderFactoryproviderInfos, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProviderInfo", start, limit)

	if err != nil {
		return *new([]ProviderFactoryproviderInfos), err
	}

	out0 := *abi.ConvertType(out[0], new([]ProviderFactoryproviderInfos)).(*[]ProviderFactoryproviderInfos)

	return out0, err

}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(uint256 start, uint256 limit) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[])[])
func (_ProviderFactory *ProviderFactorySession) GetProviderInfo(start *big.Int, limit *big.Int) ([]ProviderFactoryproviderInfos, error) {
	return _ProviderFactory.Contract.GetProviderInfo(&_ProviderFactory.CallOpts, start, limit)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(uint256 start, uint256 limit) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[])[])
func (_ProviderFactory *ProviderFactoryCallerSession) GetProviderInfo(start *big.Int, limit *big.Int) ([]ProviderFactoryproviderInfos, error) {
	return _ProviderFactory.Contract.GetProviderInfo(&_ProviderFactory.CallOpts, start, limit)
}

// GetProviderInfoLength is a free data retrieval call binding the contract method 0xa2b7e6d7.
//
// Solidity: function getProviderInfoLength() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) GetProviderInfoLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProviderInfoLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderInfoLength is a free data retrieval call binding the contract method 0xa2b7e6d7.
//
// Solidity: function getProviderInfoLength() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) GetProviderInfoLength() (*big.Int, error) {
	return _ProviderFactory.Contract.GetProviderInfoLength(&_ProviderFactory.CallOpts)
}

// GetProviderInfoLength is a free data retrieval call binding the contract method 0xa2b7e6d7.
//
// Solidity: function getProviderInfoLength() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) GetProviderInfoLength() (*big.Int, error) {
	return _ProviderFactory.Contract.GetProviderInfoLength(&_ProviderFactory.CallOpts)
}

// GetProviderSingle is a free data retrieval call binding the contract method 0xd8c38449.
//
// Solidity: function getProviderSingle(address _provider_contract) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[]))
func (_ProviderFactory *ProviderFactoryCaller) GetProviderSingle(opts *bind.CallOpts, _provider_contract common.Address) (ProviderFactoryproviderInfos, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getProviderSingle", _provider_contract)

	if err != nil {
		return *new(ProviderFactoryproviderInfos), err
	}

	out0 := *abi.ConvertType(out[0], new(ProviderFactoryproviderInfos)).(*ProviderFactoryproviderInfos)

	return out0, err

}

// GetProviderSingle is a free data retrieval call binding the contract method 0xd8c38449.
//
// Solidity: function getProviderSingle(address _provider_contract) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[]))
func (_ProviderFactory *ProviderFactorySession) GetProviderSingle(_provider_contract common.Address) (ProviderFactoryproviderInfos, error) {
	return _ProviderFactory.Contract.GetProviderSingle(&_ProviderFactory.CallOpts, _provider_contract)
}

// GetProviderSingle is a free data retrieval call binding the contract method 0xd8c38449.
//
// Solidity: function getProviderSingle(address _provider_contract) view returns((address,((uint256,uint256,uint256),(uint256,uint256,uint256),(uint256,uint256,uint256),bool,uint8,address,string,string,uint256,uint256),uint256,address[]))
func (_ProviderFactory *ProviderFactoryCallerSession) GetProviderSingle(_provider_contract common.Address) (ProviderFactoryproviderInfos, error) {
	return _ProviderFactory.Contract.GetProviderSingle(&_ProviderFactory.CallOpts, _provider_contract)
}

// GetPunishAmount is a free data retrieval call binding the contract method 0x342fcbe4.
//
// Solidity: function getPunishAmount(uint256 punish_amount) view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) GetPunishAmount(opts *bind.CallOpts, punish_amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getPunishAmount", punish_amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishAmount is a free data retrieval call binding the contract method 0x342fcbe4.
//
// Solidity: function getPunishAmount(uint256 punish_amount) view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) GetPunishAmount(punish_amount *big.Int) (*big.Int, error) {
	return _ProviderFactory.Contract.GetPunishAmount(&_ProviderFactory.CallOpts, punish_amount)
}

// GetPunishAmount is a free data retrieval call binding the contract method 0x342fcbe4.
//
// Solidity: function getPunishAmount(uint256 punish_amount) view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) GetPunishAmount(punish_amount *big.Int) (*big.Int, error) {
	return _ProviderFactory.Contract.GetPunishAmount(&_ProviderFactory.CallOpts, punish_amount)
}

// GetTotalDetail is a free data retrieval call binding the contract method 0xd6843945.
//
// Solidity: function getTotalDetail() view returns((uint256,uint256,uint256), (uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCaller) GetTotalDetail(opts *bind.CallOpts) (poaResource, poaResource, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "getTotalDetail")

	if err != nil {
		return *new(poaResource), *new(poaResource), err
	}

	out0 := *abi.ConvertType(out[0], new(poaResource)).(*poaResource)
	out1 := *abi.ConvertType(out[1], new(poaResource)).(*poaResource)

	return out0, out1, err

}

// GetTotalDetail is a free data retrieval call binding the contract method 0xd6843945.
//
// Solidity: function getTotalDetail() view returns((uint256,uint256,uint256), (uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactorySession) GetTotalDetail() (poaResource, poaResource, error) {
	return _ProviderFactory.Contract.GetTotalDetail(&_ProviderFactory.CallOpts)
}

// GetTotalDetail is a free data retrieval call binding the contract method 0xd6843945.
//
// Solidity: function getTotalDetail() view returns((uint256,uint256,uint256), (uint256,uint256,uint256))
func (_ProviderFactory *ProviderFactoryCallerSession) GetTotalDetail() (poaResource, poaResource, error) {
	return _ProviderFactory.Contract.GetTotalDetail(&_ProviderFactory.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ProviderFactory *ProviderFactoryCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ProviderFactory *ProviderFactorySession) Initialized() (bool, error) {
	return _ProviderFactory.Contract.Initialized(&_ProviderFactory.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ProviderFactory *ProviderFactoryCallerSession) Initialized() (bool, error) {
	return _ProviderFactory.Contract.Initialized(&_ProviderFactory.CallOpts)
}

// MaxValueTobeProvider is a free data retrieval call binding the contract method 0x83c36af7.
//
// Solidity: function max_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) MaxValueTobeProvider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "max_value_tobe_provider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValueTobeProvider is a free data retrieval call binding the contract method 0x83c36af7.
//
// Solidity: function max_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) MaxValueTobeProvider() (*big.Int, error) {
	return _ProviderFactory.Contract.MaxValueTobeProvider(&_ProviderFactory.CallOpts)
}

// MaxValueTobeProvider is a free data retrieval call binding the contract method 0x83c36af7.
//
// Solidity: function max_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) MaxValueTobeProvider() (*big.Int, error) {
	return _ProviderFactory.Contract.MaxValueTobeProvider(&_ProviderFactory.CallOpts)
}

// MinValueTobeProvider is a free data retrieval call binding the contract method 0x168dc4f0.
//
// Solidity: function min_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) MinValueTobeProvider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "min_value_tobe_provider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValueTobeProvider is a free data retrieval call binding the contract method 0x168dc4f0.
//
// Solidity: function min_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) MinValueTobeProvider() (*big.Int, error) {
	return _ProviderFactory.Contract.MinValueTobeProvider(&_ProviderFactory.CallOpts)
}

// MinValueTobeProvider is a free data retrieval call binding the contract method 0x168dc4f0.
//
// Solidity: function min_value_tobe_provider() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) MinValueTobeProvider() (*big.Int, error) {
	return _ProviderFactory.Contract.MinValueTobeProvider(&_ProviderFactory.CallOpts)
}

// OrderFactory is a free data retrieval call binding the contract method 0xe37fd986.
//
// Solidity: function order_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) OrderFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "order_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderFactory is a free data retrieval call binding the contract method 0xe37fd986.
//
// Solidity: function order_factory() view returns(address)
func (_ProviderFactory *ProviderFactorySession) OrderFactory() (common.Address, error) {
	return _ProviderFactory.Contract.OrderFactory(&_ProviderFactory.CallOpts)
}

// OrderFactory is a free data retrieval call binding the contract method 0xe37fd986.
//
// Solidity: function order_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) OrderFactory() (common.Address, error) {
	return _ProviderFactory.Contract.OrderFactory(&_ProviderFactory.CallOpts)
}

// ProviderLockTime is a free data retrieval call binding the contract method 0xbacc430f.
//
// Solidity: function provider_lock_time() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) ProviderLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "provider_lock_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderLockTime is a free data retrieval call binding the contract method 0xbacc430f.
//
// Solidity: function provider_lock_time() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) ProviderLockTime() (*big.Int, error) {
	return _ProviderFactory.Contract.ProviderLockTime(&_ProviderFactory.CallOpts)
}

// ProviderLockTime is a free data retrieval call binding the contract method 0xbacc430f.
//
// Solidity: function provider_lock_time() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) ProviderLockTime() (*big.Int, error) {
	return _ProviderFactory.Contract.ProviderLockTime(&_ProviderFactory.CallOpts)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "providers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address)
func (_ProviderFactory *ProviderFactorySession) Providers(arg0 common.Address) (common.Address, error) {
	return _ProviderFactory.Contract.Providers(&_ProviderFactory.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) Providers(arg0 common.Address) (common.Address, error) {
	return _ProviderFactory.Contract.Providers(&_ProviderFactory.CallOpts, arg0)
}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) PunishAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "punish_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ProviderFactory *ProviderFactorySession) PunishAddress() (common.Address, error) {
	return _ProviderFactory.Contract.PunishAddress(&_ProviderFactory.CallOpts)
}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) PunishAddress() (common.Address, error) {
	return _ProviderFactory.Contract.PunishAddress(&_ProviderFactory.CallOpts)
}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) PunishAllPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "punish_all_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) PunishAllPercent() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishAllPercent(&_ProviderFactory.CallOpts)
}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) PunishAllPercent() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishAllPercent(&_ProviderFactory.CallOpts)
}

// PunishInterval is a free data retrieval call binding the contract method 0x2d918023.
//
// Solidity: function punish_interval() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) PunishInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "punish_interval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishInterval is a free data retrieval call binding the contract method 0x2d918023.
//
// Solidity: function punish_interval() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) PunishInterval() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishInterval(&_ProviderFactory.CallOpts)
}

// PunishInterval is a free data retrieval call binding the contract method 0x2d918023.
//
// Solidity: function punish_interval() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) PunishInterval() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishInterval(&_ProviderFactory.CallOpts)
}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) PunishPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "punish_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) PunishPercent() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishPercent(&_ProviderFactory.CallOpts)
}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) PunishPercent() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishPercent(&_ProviderFactory.CallOpts)
}

// PunishStartLimit is a free data retrieval call binding the contract method 0x78c330d0.
//
// Solidity: function punish_start_limit() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCaller) PunishStartLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "punish_start_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishStartLimit is a free data retrieval call binding the contract method 0x78c330d0.
//
// Solidity: function punish_start_limit() view returns(uint256)
func (_ProviderFactory *ProviderFactorySession) PunishStartLimit() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishStartLimit(&_ProviderFactory.CallOpts)
}

// PunishStartLimit is a free data retrieval call binding the contract method 0x78c330d0.
//
// Solidity: function punish_start_limit() view returns(uint256)
func (_ProviderFactory *ProviderFactoryCallerSession) PunishStartLimit() (*big.Int, error) {
	return _ProviderFactory.Contract.PunishStartLimit(&_ProviderFactory.CallOpts)
}

// TotalAll is a free data retrieval call binding the contract method 0x306c47ea.
//
// Solidity: function total_all() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactoryCaller) TotalAll(opts *bind.CallOpts) (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "total_all")

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

// TotalAll is a free data retrieval call binding the contract method 0x306c47ea.
//
// Solidity: function total_all() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactorySession) TotalAll() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _ProviderFactory.Contract.TotalAll(&_ProviderFactory.CallOpts)
}

// TotalAll is a free data retrieval call binding the contract method 0x306c47ea.
//
// Solidity: function total_all() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactoryCallerSession) TotalAll() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _ProviderFactory.Contract.TotalAll(&_ProviderFactory.CallOpts)
}

// TotalUsed is a free data retrieval call binding the contract method 0x5ed98228.
//
// Solidity: function total_used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactoryCaller) TotalUsed(opts *bind.CallOpts) (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "total_used")

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

// TotalUsed is a free data retrieval call binding the contract method 0x5ed98228.
//
// Solidity: function total_used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactorySession) TotalUsed() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _ProviderFactory.Contract.TotalUsed(&_ProviderFactory.CallOpts)
}

// TotalUsed is a free data retrieval call binding the contract method 0x5ed98228.
//
// Solidity: function total_used() view returns(uint256 cpu_count, uint256 memory_count, uint256 storage_count)
func (_ProviderFactory *ProviderFactoryCallerSession) TotalUsed() (struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}, error) {
	return _ProviderFactory.Contract.TotalUsed(&_ProviderFactory.CallOpts)
}

// ValFactory is a free data retrieval call binding the contract method 0xd9780266.
//
// Solidity: function val_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCaller) ValFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "val_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValFactory is a free data retrieval call binding the contract method 0xd9780266.
//
// Solidity: function val_factory() view returns(address)
func (_ProviderFactory *ProviderFactorySession) ValFactory() (common.Address, error) {
	return _ProviderFactory.Contract.ValFactory(&_ProviderFactory.CallOpts)
}

// ValFactory is a free data retrieval call binding the contract method 0xd9780266.
//
// Solidity: function val_factory() view returns(address)
func (_ProviderFactory *ProviderFactoryCallerSession) ValFactory() (common.Address, error) {
	return _ProviderFactory.Contract.ValFactory(&_ProviderFactory.CallOpts)
}

// WhetherCanPOR is a free data retrieval call binding the contract method 0xc5611b6a.
//
// Solidity: function whetherCanPOR(address provider_owner) view returns(bool)
func (_ProviderFactory *ProviderFactoryCaller) WhetherCanPOR(opts *bind.CallOpts, provider_owner common.Address) (bool, error) {
	var out []interface{}
	err := _ProviderFactory.contract.Call(opts, &out, "whetherCanPOR", provider_owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhetherCanPOR is a free data retrieval call binding the contract method 0xc5611b6a.
//
// Solidity: function whetherCanPOR(address provider_owner) view returns(bool)
func (_ProviderFactory *ProviderFactorySession) WhetherCanPOR(provider_owner common.Address) (bool, error) {
	return _ProviderFactory.Contract.WhetherCanPOR(&_ProviderFactory.CallOpts, provider_owner)
}

// WhetherCanPOR is a free data retrieval call binding the contract method 0xc5611b6a.
//
// Solidity: function whetherCanPOR(address provider_owner) view returns(bool)
func (_ProviderFactory *ProviderFactoryCallerSession) WhetherCanPOR(provider_owner common.Address) (bool, error) {
	return _ProviderFactory.Contract.WhetherCanPOR(&_ProviderFactory.CallOpts, provider_owner)
}

// AddMargin is a paid mutator transaction binding the contract method 0x483a00e8.
//
// Solidity: function addMargin() payable returns()
func (_ProviderFactory *ProviderFactoryTransactor) AddMargin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "addMargin")
}

// AddMargin is a paid mutator transaction binding the contract method 0x483a00e8.
//
// Solidity: function addMargin() payable returns()
func (_ProviderFactory *ProviderFactorySession) AddMargin() (*types.Transaction, error) {
	return _ProviderFactory.Contract.AddMargin(&_ProviderFactory.TransactOpts)
}

// AddMargin is a paid mutator transaction binding the contract method 0x483a00e8.
//
// Solidity: function addMargin() payable returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) AddMargin() (*types.Transaction, error) {
	return _ProviderFactory.Contract.AddMargin(&_ProviderFactory.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address new_admin) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, new_admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeAdmin", new_admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address new_admin) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeAdmin(new_admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeAdmin(&_ProviderFactory.TransactOpts, new_admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address new_admin) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeAdmin(new_admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeAdmin(&_ProviderFactory.TransactOpts, new_admin)
}

// ChangeAuditorFactory is a paid mutator transaction binding the contract method 0x8fe0ce72.
//
// Solidity: function changeAuditorFactory(address new_audit_factory) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeAuditorFactory(opts *bind.TransactOpts, new_audit_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeAuditorFactory", new_audit_factory)
}

// ChangeAuditorFactory is a paid mutator transaction binding the contract method 0x8fe0ce72.
//
// Solidity: function changeAuditorFactory(address new_audit_factory) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeAuditorFactory(new_audit_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeAuditorFactory(&_ProviderFactory.TransactOpts, new_audit_factory)
}

// ChangeAuditorFactory is a paid mutator transaction binding the contract method 0x8fe0ce72.
//
// Solidity: function changeAuditorFactory(address new_audit_factory) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeAuditorFactory(new_audit_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeAuditorFactory(&_ProviderFactory.TransactOpts, new_audit_factory)
}

// ChangeDecimal is a paid mutator transaction binding the contract method 0xf4be3256.
//
// Solidity: function changeDecimal(uint256 new_cpu_decimal, uint256 new_memory_decimal) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeDecimal(opts *bind.TransactOpts, new_cpu_decimal *big.Int, new_memory_decimal *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeDecimal", new_cpu_decimal, new_memory_decimal)
}

// ChangeDecimal is a paid mutator transaction binding the contract method 0xf4be3256.
//
// Solidity: function changeDecimal(uint256 new_cpu_decimal, uint256 new_memory_decimal) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeDecimal(new_cpu_decimal *big.Int, new_memory_decimal *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeDecimal(&_ProviderFactory.TransactOpts, new_cpu_decimal, new_memory_decimal)
}

// ChangeDecimal is a paid mutator transaction binding the contract method 0xf4be3256.
//
// Solidity: function changeDecimal(uint256 new_cpu_decimal, uint256 new_memory_decimal) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeDecimal(new_cpu_decimal *big.Int, new_memory_decimal *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeDecimal(&_ProviderFactory.TransactOpts, new_cpu_decimal, new_memory_decimal)
}

// ChangeOrderFactory is a paid mutator transaction binding the contract method 0xacbee7e0.
//
// Solidity: function changeOrderFactory(address new_order_factory) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeOrderFactory(opts *bind.TransactOpts, new_order_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeOrderFactory", new_order_factory)
}

// ChangeOrderFactory is a paid mutator transaction binding the contract method 0xacbee7e0.
//
// Solidity: function changeOrderFactory(address new_order_factory) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeOrderFactory(new_order_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeOrderFactory(&_ProviderFactory.TransactOpts, new_order_factory)
}

// ChangeOrderFactory is a paid mutator transaction binding the contract method 0xacbee7e0.
//
// Solidity: function changeOrderFactory(address new_order_factory) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeOrderFactory(new_order_factory common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeOrderFactory(&_ProviderFactory.TransactOpts, new_order_factory)
}

// ChangeProviderLimit is a paid mutator transaction binding the contract method 0xf518147d.
//
// Solidity: function changeProviderLimit(uint256 _new_min, uint256 _new_max) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeProviderLimit(opts *bind.TransactOpts, _new_min *big.Int, _new_max *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeProviderLimit", _new_min, _new_max)
}

// ChangeProviderLimit is a paid mutator transaction binding the contract method 0xf518147d.
//
// Solidity: function changeProviderLimit(uint256 _new_min, uint256 _new_max) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeProviderLimit(_new_min *big.Int, _new_max *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderLimit(&_ProviderFactory.TransactOpts, _new_min, _new_max)
}

// ChangeProviderLimit is a paid mutator transaction binding the contract method 0xf518147d.
//
// Solidity: function changeProviderLimit(uint256 _new_min, uint256 _new_max) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeProviderLimit(_new_min *big.Int, _new_max *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderLimit(&_ProviderFactory.TransactOpts, _new_min, _new_max)
}

// ChangeProviderLockTime is a paid mutator transaction binding the contract method 0xbc3443a5.
//
// Solidity: function changeProviderLockTime(uint256 _lock_time) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeProviderLockTime(opts *bind.TransactOpts, _lock_time *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeProviderLockTime", _lock_time)
}

// ChangeProviderLockTime is a paid mutator transaction binding the contract method 0xbc3443a5.
//
// Solidity: function changeProviderLockTime(uint256 _lock_time) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeProviderLockTime(_lock_time *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderLockTime(&_ProviderFactory.TransactOpts, _lock_time)
}

// ChangeProviderLockTime is a paid mutator transaction binding the contract method 0xbc3443a5.
//
// Solidity: function changeProviderLockTime(uint256 _lock_time) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeProviderLockTime(_lock_time *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderLockTime(&_ProviderFactory.TransactOpts, _lock_time)
}

// ChangeProviderResource is a paid mutator transaction binding the contract method 0x0657026a.
//
// Solidity: function changeProviderResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeProviderResource(opts *bind.TransactOpts, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeProviderResource", cpu_count, mem_count, storage_count, add)
}

// ChangeProviderResource is a paid mutator transaction binding the contract method 0x0657026a.
//
// Solidity: function changeProviderResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeProviderResource(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderResource(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, add)
}

// ChangeProviderResource is a paid mutator transaction binding the contract method 0x0657026a.
//
// Solidity: function changeProviderResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeProviderResource(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderResource(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, add)
}

// ChangeProviderState is a paid mutator transaction binding the contract method 0x00e6fb08.
//
// Solidity: function changeProviderState(address provider_owner, bool whether_start) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeProviderState(opts *bind.TransactOpts, provider_owner common.Address, whether_start bool) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeProviderState", provider_owner, whether_start)
}

// ChangeProviderState is a paid mutator transaction binding the contract method 0x00e6fb08.
//
// Solidity: function changeProviderState(address provider_owner, bool whether_start) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeProviderState(provider_owner common.Address, whether_start bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderState(&_ProviderFactory.TransactOpts, provider_owner, whether_start)
}

// ChangeProviderState is a paid mutator transaction binding the contract method 0x00e6fb08.
//
// Solidity: function changeProviderState(address provider_owner, bool whether_start) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeProviderState(provider_owner common.Address, whether_start bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderState(&_ProviderFactory.TransactOpts, provider_owner, whether_start)
}

// ChangeProviderUsedResource is a paid mutator transaction binding the contract method 0x0d351e73.
//
// Solidity: function changeProviderUsedResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangeProviderUsedResource(opts *bind.TransactOpts, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changeProviderUsedResource", cpu_count, mem_count, storage_count, add)
}

// ChangeProviderUsedResource is a paid mutator transaction binding the contract method 0x0d351e73.
//
// Solidity: function changeProviderUsedResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactorySession) ChangeProviderUsedResource(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderUsedResource(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, add)
}

// ChangeProviderUsedResource is a paid mutator transaction binding the contract method 0x0d351e73.
//
// Solidity: function changeProviderUsedResource(uint256 cpu_count, uint256 mem_count, uint256 storage_count, bool add) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangeProviderUsedResource(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, add bool) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangeProviderUsedResource(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, add)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangePunishAddress(opts *bind.TransactOpts, _punish_address common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changePunishAddress", _punish_address)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ProviderFactory *ProviderFactorySession) ChangePunishAddress(_punish_address common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishAddress(&_ProviderFactory.TransactOpts, _punish_address)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangePunishAddress(_punish_address common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishAddress(&_ProviderFactory.TransactOpts, _punish_address)
}

// ChangePunishParam is a paid mutator transaction binding the contract method 0xf1c53fb5.
//
// Solidity: function changePunishParam(uint256 _new_punish_start_limit, uint256 _new_punish_interval) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangePunishParam(opts *bind.TransactOpts, _new_punish_start_limit *big.Int, _new_punish_interval *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changePunishParam", _new_punish_start_limit, _new_punish_interval)
}

// ChangePunishParam is a paid mutator transaction binding the contract method 0xf1c53fb5.
//
// Solidity: function changePunishParam(uint256 _new_punish_start_limit, uint256 _new_punish_interval) returns()
func (_ProviderFactory *ProviderFactorySession) ChangePunishParam(_new_punish_start_limit *big.Int, _new_punish_interval *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishParam(&_ProviderFactory.TransactOpts, _new_punish_start_limit, _new_punish_interval)
}

// ChangePunishParam is a paid mutator transaction binding the contract method 0xf1c53fb5.
//
// Solidity: function changePunishParam(uint256 _new_punish_start_limit, uint256 _new_punish_interval) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangePunishParam(_new_punish_start_limit *big.Int, _new_punish_interval *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishParam(&_ProviderFactory.TransactOpts, _new_punish_start_limit, _new_punish_interval)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ChangePunishPercent(opts *bind.TransactOpts, _new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "changePunishPercent", _new_punish_percent, _new_punish_all_percent)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ProviderFactory *ProviderFactorySession) ChangePunishPercent(_new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishPercent(&_ProviderFactory.TransactOpts, _new_punish_percent, _new_punish_all_percent)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ChangePunishPercent(_new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ChangePunishPercent(&_ProviderFactory.TransactOpts, _new_punish_percent, _new_punish_all_percent)
}

// CloseProvider is a paid mutator transaction binding the contract method 0x7ed65c87.
//
// Solidity: function closeProvider() returns()
func (_ProviderFactory *ProviderFactoryTransactor) CloseProvider(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "closeProvider")
}

// CloseProvider is a paid mutator transaction binding the contract method 0x7ed65c87.
//
// Solidity: function closeProvider() returns()
func (_ProviderFactory *ProviderFactorySession) CloseProvider() (*types.Transaction, error) {
	return _ProviderFactory.Contract.CloseProvider(&_ProviderFactory.TransactOpts)
}

// CloseProvider is a paid mutator transaction binding the contract method 0x7ed65c87.
//
// Solidity: function closeProvider() returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) CloseProvider() (*types.Transaction, error) {
	return _ProviderFactory.Contract.CloseProvider(&_ProviderFactory.TransactOpts)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0xdb4fce50.
//
// Solidity: function consumeResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactoryTransactor) ConsumeResource(opts *bind.TransactOpts, account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "consumeResource", account, cpu_count, mem_count, storage_count)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0xdb4fce50.
//
// Solidity: function consumeResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactorySession) ConsumeResource(account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ConsumeResource(&_ProviderFactory.TransactOpts, account, cpu_count, mem_count, storage_count)
}

// ConsumeResource is a paid mutator transaction binding the contract method 0xdb4fce50.
//
// Solidity: function consumeResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) ConsumeResource(account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.ConsumeResource(&_ProviderFactory.TransactOpts, account, cpu_count, mem_count, storage_count)
}

// CreateNewProvider is a paid mutator transaction binding the contract method 0x09fe81b3.
//
// Solidity: function createNewProvider(uint256 cpu_count, uint256 mem_count, uint256 storage_count, string region, string provider_info) payable returns(address)
func (_ProviderFactory *ProviderFactoryTransactor) CreateNewProvider(opts *bind.TransactOpts, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, region string, provider_info string) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "createNewProvider", cpu_count, mem_count, storage_count, region, provider_info)
}

// CreateNewProvider is a paid mutator transaction binding the contract method 0x09fe81b3.
//
// Solidity: function createNewProvider(uint256 cpu_count, uint256 mem_count, uint256 storage_count, string region, string provider_info) payable returns(address)
func (_ProviderFactory *ProviderFactorySession) CreateNewProvider(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, region string, provider_info string) (*types.Transaction, error) {
	return _ProviderFactory.Contract.CreateNewProvider(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, region, provider_info)
}

// CreateNewProvider is a paid mutator transaction binding the contract method 0x09fe81b3.
//
// Solidity: function createNewProvider(uint256 cpu_count, uint256 mem_count, uint256 storage_count, string region, string provider_info) payable returns(address)
func (_ProviderFactory *ProviderFactoryTransactorSession) CreateNewProvider(cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int, region string, provider_info string) (*types.Transaction, error) {
	return _ProviderFactory.Contract.CreateNewProvider(&_ProviderFactory.TransactOpts, cpu_count, mem_count, storage_count, region, provider_info)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_ProviderFactory *ProviderFactoryTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_ProviderFactory *ProviderFactorySession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.Initialize(&_ProviderFactory.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.Initialize(&_ProviderFactory.TransactOpts, _admin)
}

// RecoverResource is a paid mutator transaction binding the contract method 0x91d62360.
//
// Solidity: function recoverResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactoryTransactor) RecoverResource(opts *bind.TransactOpts, account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "recoverResource", account, cpu_count, mem_count, storage_count)
}

// RecoverResource is a paid mutator transaction binding the contract method 0x91d62360.
//
// Solidity: function recoverResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactorySession) RecoverResource(account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.RecoverResource(&_ProviderFactory.TransactOpts, account, cpu_count, mem_count, storage_count)
}

// RecoverResource is a paid mutator transaction binding the contract method 0x91d62360.
//
// Solidity: function recoverResource(address account, uint256 cpu_count, uint256 mem_count, uint256 storage_count) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) RecoverResource(account common.Address, cpu_count *big.Int, mem_count *big.Int, storage_count *big.Int) (*types.Transaction, error) {
	return _ProviderFactory.Contract.RecoverResource(&_ProviderFactory.TransactOpts, account, cpu_count, mem_count, storage_count)
}

// RemovePunishList is a paid mutator transaction binding the contract method 0xbe23c44d.
//
// Solidity: function removePunishList(address provider) returns()
func (_ProviderFactory *ProviderFactoryTransactor) RemovePunishList(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "removePunishList", provider)
}

// RemovePunishList is a paid mutator transaction binding the contract method 0xbe23c44d.
//
// Solidity: function removePunishList(address provider) returns()
func (_ProviderFactory *ProviderFactorySession) RemovePunishList(provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.RemovePunishList(&_ProviderFactory.TransactOpts, provider)
}

// RemovePunishList is a paid mutator transaction binding the contract method 0xbe23c44d.
//
// Solidity: function removePunishList(address provider) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) RemovePunishList(provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.RemovePunishList(&_ProviderFactory.TransactOpts, provider)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address new_provider) returns()
func (_ProviderFactory *ProviderFactoryTransactor) TryPunish(opts *bind.TransactOpts, new_provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "tryPunish", new_provider)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address new_provider) returns()
func (_ProviderFactory *ProviderFactorySession) TryPunish(new_provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.TryPunish(&_ProviderFactory.TransactOpts, new_provider)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address new_provider) returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) TryPunish(new_provider common.Address) (*types.Transaction, error) {
	return _ProviderFactory.Contract.TryPunish(&_ProviderFactory.TransactOpts, new_provider)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_ProviderFactory *ProviderFactoryTransactor) WithdrawMargin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderFactory.contract.Transact(opts, "withdrawMargin")
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_ProviderFactory *ProviderFactorySession) WithdrawMargin() (*types.Transaction, error) {
	return _ProviderFactory.Contract.WithdrawMargin(&_ProviderFactory.TransactOpts)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x9e83d5b1.
//
// Solidity: function withdrawMargin() returns()
func (_ProviderFactory *ProviderFactoryTransactorSession) WithdrawMargin() (*types.Transaction, error) {
	return _ProviderFactory.Contract.WithdrawMargin(&_ProviderFactory.TransactOpts)
}

// ProviderFactoryProviderCreateIterator is returned from FilterProviderCreate and is used to iterate over the raw logs and unpacked data for ProviderCreate events raised by the ProviderFactory contract.
type ProviderFactoryProviderCreateIterator struct {
	Event *ProviderFactoryProviderCreate // Event containing the contract specifics and raw log

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
func (it *ProviderFactoryProviderCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderFactoryProviderCreate)
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
		it.Event = new(ProviderFactoryProviderCreate)
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
func (it *ProviderFactoryProviderCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderFactoryProviderCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderFactoryProviderCreate represents a ProviderCreate event raised by the ProviderFactory contract.
type ProviderFactoryProviderCreate struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProviderCreate is a free log retrieval operation binding the contract event 0x7a5f975375bdfb84fb025f54bdc23a48d52dd5dbb07eeee54ef79ad7c7ff0164.
//
// Solidity: event ProviderCreate(address arg0)
func (_ProviderFactory *ProviderFactoryFilterer) FilterProviderCreate(opts *bind.FilterOpts) (*ProviderFactoryProviderCreateIterator, error) {

	logs, sub, err := _ProviderFactory.contract.FilterLogs(opts, "ProviderCreate")
	if err != nil {
		return nil, err
	}
	return &ProviderFactoryProviderCreateIterator{contract: _ProviderFactory.contract, event: "ProviderCreate", logs: logs, sub: sub}, nil
}

// WatchProviderCreate is a free log subscription operation binding the contract event 0x7a5f975375bdfb84fb025f54bdc23a48d52dd5dbb07eeee54ef79ad7c7ff0164.
//
// Solidity: event ProviderCreate(address arg0)
func (_ProviderFactory *ProviderFactoryFilterer) WatchProviderCreate(opts *bind.WatchOpts, sink chan<- *ProviderFactoryProviderCreate) (event.Subscription, error) {

	logs, sub, err := _ProviderFactory.contract.WatchLogs(opts, "ProviderCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderFactoryProviderCreate)
				if err := _ProviderFactory.contract.UnpackLog(event, "ProviderCreate", log); err != nil {
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

// ParseProviderCreate is a log parse operation binding the contract event 0x7a5f975375bdfb84fb025f54bdc23a48d52dd5dbb07eeee54ef79ad7c7ff0164.
//
// Solidity: event ProviderCreate(address arg0)
func (_ProviderFactory *ProviderFactoryFilterer) ParseProviderCreate(log types.Log) (*ProviderFactoryProviderCreate, error) {
	event := new(ProviderFactoryProviderCreate)
	if err := _ProviderFactory.contract.UnpackLog(event, "ProviderCreate", log); err != nil {
		return nil, err
	}
	return event, nil
}

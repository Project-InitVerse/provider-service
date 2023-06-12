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

// ValidatorFactoryValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorFactoryValidatorInfo struct {
	Validator         common.Address
	ValidatorContract common.Address
	State             uint8
	StartTime         *big.Int
}

// ValidatorFactoryproviderChallengeInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorFactoryproviderChallengeInfo struct {
	Provider            common.Address
	ChallengeValidator  common.Address
	Md5Seed             *big.Int
	Url                 string
	CreateChallengeTime *big.Int
	ChallengeFinishTime *big.Int
	State               uint8
	ChallengeAmount     *big.Int
	Seed                *big.Int
	RootHash            *big.Int
	Index               *big.Int
}

// ValidatorFactoryABI is the input ABI used to generate the binding from.
const ValidatorFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ChallengeCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ChallengeEnd\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MarginCalls\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"all_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"all_validators\",\"outputs\":[{\"internalType\":\"contractIValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challenge_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root_hash\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorFactory.ChallengeState\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"challengeFinish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"md5_seed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"challengeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge_all_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge_sdl_trx_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_admin\",\"type\":\"address\"}],\"name\":\"changeAdminAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_trx_id\",\"type\":\"uint256\"}],\"name\":\"changeChallengeSdlTrxID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_challenge_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challenge_all_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_provider_start_challenge_time\",\"type\":\"uint256\"}],\"name\":\"changeMaxChallengeParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_validator_count\",\"type\":\"uint256\"}],\"name\":\"changeMaxValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_punish_address\",\"type\":\"address\"}],\"name\":\"changePunishAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_punish_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_punish_all_percent\",\"type\":\"uint256\"}],\"name\":\"changePunishPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_team_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validator_percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_all_percent\",\"type\":\"uint256\"}],\"name\":\"changeRewardPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team_address\",\"type\":\"address\"}],\"name\":\"changeTeamAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_lock\",\"type\":\"uint256\"}],\"name\":\"changeValidatorLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validator_min_pledgeAmount\",\"type\":\"uint256\"}],\"name\":\"changeValidatorMinPledgeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_interval\",\"type\":\"uint256\"}],\"name\":\"changeValidatorPunishInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_start_limit\",\"type\":\"uint256\"}],\"name\":\"changeValidatorPunishStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumValidatorState\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"changeValidatorState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_challenge_provider_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_validator_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitProduceBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllActiveValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_contract\",\"type\":\"address\"},{\"internalType\":\"enumValidatorState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorFactory.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllActiveValidatorAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPunishValidator\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_contract\",\"type\":\"address\"},{\"internalType\":\"enumValidatorState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorFactory.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_owner\",\"type\":\"address\"}],\"name\":\"getProviderChallengeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenge_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"md5_seed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"create_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challenge_finish_time\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorFactory.ChallengeState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"challenge_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root_hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorFactory.providerChallengeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPunishAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_init_validator\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_challenge_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_challenge_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_provider_start_challenge_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_validator_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owner_validator\",\"outputs\":[{\"internalType\":\"contractIValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"provider_challenge_info\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenge_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"md5_seed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"create_challenge_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challenge_finish_time\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorFactory.ChallengeState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"challenge_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root_hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provider_factory\",\"outputs\":[{\"internalType\":\"contractIProviderFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"provider_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"provider_last_challenge_state\",\"outputs\":[{\"internalType\":\"enumValidatorFactory.ChallengeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_all_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punish_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeRankingList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"tryPunish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"validatorNotSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator_lock_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator_percent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator_pledgeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator_punish_interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator_punish_start_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList_validator\",\"outputs\":[{\"internalType\":\"contractIValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorFactory is an auto generated Go binding around an Ethereum contract.
type ValidatorFactory struct {
	ValidatorFactoryCaller     // Read-only binding to the contract
	ValidatorFactoryTransactor // Write-only binding to the contract
	ValidatorFactoryFilterer   // Log filterer for contract events
}

// ValidatorFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorFactorySession struct {
	Contract     *ValidatorFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorFactoryCallerSession struct {
	Contract *ValidatorFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorFactoryTransactorSession struct {
	Contract     *ValidatorFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorFactoryRaw struct {
	Contract *ValidatorFactory // Generic contract binding to access the raw methods on
}

// ValidatorFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorFactoryCallerRaw struct {
	Contract *ValidatorFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorFactoryTransactorRaw struct {
	Contract *ValidatorFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorFactory creates a new instance of ValidatorFactory, bound to a specific deployed contract.
func NewValidatorFactory(address common.Address, backend bind.ContractBackend) (*ValidatorFactory, error) {
	contract, err := bindValidatorFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorFactory{ValidatorFactoryCaller: ValidatorFactoryCaller{contract: contract}, ValidatorFactoryTransactor: ValidatorFactoryTransactor{contract: contract}, ValidatorFactoryFilterer: ValidatorFactoryFilterer{contract: contract}}, nil
}

// NewValidatorFactoryCaller creates a new read-only instance of ValidatorFactory, bound to a specific deployed contract.
func NewValidatorFactoryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorFactoryCaller, error) {
	contract, err := bindValidatorFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorFactoryCaller{contract: contract}, nil
}

// NewValidatorFactoryTransactor creates a new write-only instance of ValidatorFactory, bound to a specific deployed contract.
func NewValidatorFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorFactoryTransactor, error) {
	contract, err := bindValidatorFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorFactoryTransactor{contract: contract}, nil
}

// NewValidatorFactoryFilterer creates a new log filterer instance of ValidatorFactory, bound to a specific deployed contract.
func NewValidatorFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFactoryFilterer, error) {
	contract, err := bindValidatorFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFactoryFilterer{contract: contract}, nil
}

// bindValidatorFactory binds a generic wrapper to an already deployed contract.
func bindValidatorFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorFactory *ValidatorFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorFactory.Contract.ValidatorFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorFactory *ValidatorFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ValidatorFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorFactory *ValidatorFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ValidatorFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorFactory *ValidatorFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorFactory *ValidatorFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorFactory *ValidatorFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminAddress is a free data retrieval call binding the contract method 0xf2e62772.
//
// Solidity: function admin_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "admin_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xf2e62772.
//
// Solidity: function admin_address() view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) AdminAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.AdminAddress(&_ValidatorFactory.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xf2e62772.
//
// Solidity: function admin_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) AdminAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.AdminAddress(&_ValidatorFactory.CallOpts)
}

// AllPercent is a free data retrieval call binding the contract method 0xd9242ba6.
//
// Solidity: function all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) AllPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "all_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPercent is a free data retrieval call binding the contract method 0xd9242ba6.
//
// Solidity: function all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) AllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.AllPercent(&_ValidatorFactory.CallOpts)
}

// AllPercent is a free data retrieval call binding the contract method 0xd9242ba6.
//
// Solidity: function all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) AllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.AllPercent(&_ValidatorFactory.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xefe0b221.
//
// Solidity: function all_validators(uint256 ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) AllValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "all_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xefe0b221.
//
// Solidity: function all_validators(uint256 ) view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) AllValidators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorFactory.Contract.AllValidators(&_ValidatorFactory.CallOpts, arg0)
}

// AllValidators is a free data retrieval call binding the contract method 0xefe0b221.
//
// Solidity: function all_validators(uint256 ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) AllValidators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorFactory.Contract.AllValidators(&_ValidatorFactory.CallOpts, arg0)
}

// ChallengeAllPercent is a free data retrieval call binding the contract method 0xf6512661.
//
// Solidity: function challenge_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ChallengeAllPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "challenge_all_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeAllPercent is a free data retrieval call binding the contract method 0xf6512661.
//
// Solidity: function challenge_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ChallengeAllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.ChallengeAllPercent(&_ValidatorFactory.CallOpts)
}

// ChallengeAllPercent is a free data retrieval call binding the contract method 0xf6512661.
//
// Solidity: function challenge_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ChallengeAllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.ChallengeAllPercent(&_ValidatorFactory.CallOpts)
}

// ChallengeSdlTrxId is a free data retrieval call binding the contract method 0x5d980e3a.
//
// Solidity: function challenge_sdl_trx_id() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ChallengeSdlTrxId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "challenge_sdl_trx_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeSdlTrxId is a free data retrieval call binding the contract method 0x5d980e3a.
//
// Solidity: function challenge_sdl_trx_id() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ChallengeSdlTrxId() (*big.Int, error) {
	return _ValidatorFactory.Contract.ChallengeSdlTrxId(&_ValidatorFactory.CallOpts)
}

// ChallengeSdlTrxId is a free data retrieval call binding the contract method 0x5d980e3a.
//
// Solidity: function challenge_sdl_trx_id() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ChallengeSdlTrxId() (*big.Int, error) {
	return _ValidatorFactory.Contract.ChallengeSdlTrxId(&_ValidatorFactory.CallOpts)
}

// CurrentChallengeProviderCount is a free data retrieval call binding the contract method 0x7f6a249b.
//
// Solidity: function current_challenge_provider_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) CurrentChallengeProviderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "current_challenge_provider_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentChallengeProviderCount is a free data retrieval call binding the contract method 0x7f6a249b.
//
// Solidity: function current_challenge_provider_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) CurrentChallengeProviderCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.CurrentChallengeProviderCount(&_ValidatorFactory.CallOpts)
}

// CurrentChallengeProviderCount is a free data retrieval call binding the contract method 0x7f6a249b.
//
// Solidity: function current_challenge_provider_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) CurrentChallengeProviderCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.CurrentChallengeProviderCount(&_ValidatorFactory.CallOpts)
}

// CurrentValidatorCount is a free data retrieval call binding the contract method 0x375ce8dc.
//
// Solidity: function current_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) CurrentValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "current_validator_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorCount is a free data retrieval call binding the contract method 0x375ce8dc.
//
// Solidity: function current_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) CurrentValidatorCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.CurrentValidatorCount(&_ValidatorFactory.CallOpts)
}

// CurrentValidatorCount is a free data retrieval call binding the contract method 0x375ce8dc.
//
// Solidity: function current_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) CurrentValidatorCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.CurrentValidatorCount(&_ValidatorFactory.CallOpts)
}

// GetAllActiveValidator is a free data retrieval call binding the contract method 0xf70df47a.
//
// Solidity: function getAllActiveValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactoryCaller) GetAllActiveValidator(opts *bind.CallOpts) ([]ValidatorFactoryValidatorInfo, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getAllActiveValidator")

	if err != nil {
		return *new([]ValidatorFactoryValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ValidatorFactoryValidatorInfo)).(*[]ValidatorFactoryValidatorInfo)

	return out0, err

}

// GetAllActiveValidator is a free data retrieval call binding the contract method 0xf70df47a.
//
// Solidity: function getAllActiveValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactorySession) GetAllActiveValidator() ([]ValidatorFactoryValidatorInfo, error) {
	return _ValidatorFactory.Contract.GetAllActiveValidator(&_ValidatorFactory.CallOpts)
}

// GetAllActiveValidator is a free data retrieval call binding the contract method 0xf70df47a.
//
// Solidity: function getAllActiveValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetAllActiveValidator() ([]ValidatorFactoryValidatorInfo, error) {
	return _ValidatorFactory.Contract.GetAllActiveValidator(&_ValidatorFactory.CallOpts)
}

// GetAllActiveValidatorAddr is a free data retrieval call binding the contract method 0xfcbad075.
//
// Solidity: function getAllActiveValidatorAddr() view returns(address[])
func (_ValidatorFactory *ValidatorFactoryCaller) GetAllActiveValidatorAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getAllActiveValidatorAddr")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllActiveValidatorAddr is a free data retrieval call binding the contract method 0xfcbad075.
//
// Solidity: function getAllActiveValidatorAddr() view returns(address[])
func (_ValidatorFactory *ValidatorFactorySession) GetAllActiveValidatorAddr() ([]common.Address, error) {
	return _ValidatorFactory.Contract.GetAllActiveValidatorAddr(&_ValidatorFactory.CallOpts)
}

// GetAllActiveValidatorAddr is a free data retrieval call binding the contract method 0xfcbad075.
//
// Solidity: function getAllActiveValidatorAddr() view returns(address[])
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetAllActiveValidatorAddr() ([]common.Address, error) {
	return _ValidatorFactory.Contract.GetAllActiveValidatorAddr(&_ValidatorFactory.CallOpts)
}

// GetAllPunishValidator is a free data retrieval call binding the contract method 0x1b9779b6.
//
// Solidity: function getAllPunishValidator() view returns(address[])
func (_ValidatorFactory *ValidatorFactoryCaller) GetAllPunishValidator(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getAllPunishValidator")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllPunishValidator is a free data retrieval call binding the contract method 0x1b9779b6.
//
// Solidity: function getAllPunishValidator() view returns(address[])
func (_ValidatorFactory *ValidatorFactorySession) GetAllPunishValidator() ([]common.Address, error) {
	return _ValidatorFactory.Contract.GetAllPunishValidator(&_ValidatorFactory.CallOpts)
}

// GetAllPunishValidator is a free data retrieval call binding the contract method 0x1b9779b6.
//
// Solidity: function getAllPunishValidator() view returns(address[])
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetAllPunishValidator() ([]common.Address, error) {
	return _ValidatorFactory.Contract.GetAllPunishValidator(&_ValidatorFactory.CallOpts)
}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactoryCaller) GetAllValidator(opts *bind.CallOpts) ([]ValidatorFactoryValidatorInfo, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getAllValidator")

	if err != nil {
		return *new([]ValidatorFactoryValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ValidatorFactoryValidatorInfo)).(*[]ValidatorFactoryValidatorInfo)

	return out0, err

}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactorySession) GetAllValidator() ([]ValidatorFactoryValidatorInfo, error) {
	return _ValidatorFactory.Contract.GetAllValidator(&_ValidatorFactory.CallOpts)
}

// GetAllValidator is a free data retrieval call binding the contract method 0x4a91a2f8.
//
// Solidity: function getAllValidator() view returns((address,address,uint8,uint256)[])
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetAllValidator() ([]ValidatorFactoryValidatorInfo, error) {
	return _ValidatorFactory.Contract.GetAllValidator(&_ValidatorFactory.CallOpts)
}

// GetAllValidatorLength is a free data retrieval call binding the contract method 0x04d3c7bb.
//
// Solidity: function getAllValidatorLength() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) GetAllValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getAllValidatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllValidatorLength is a free data retrieval call binding the contract method 0x04d3c7bb.
//
// Solidity: function getAllValidatorLength() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) GetAllValidatorLength() (*big.Int, error) {
	return _ValidatorFactory.Contract.GetAllValidatorLength(&_ValidatorFactory.CallOpts)
}

// GetAllValidatorLength is a free data retrieval call binding the contract method 0x04d3c7bb.
//
// Solidity: function getAllValidatorLength() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetAllValidatorLength() (*big.Int, error) {
	return _ValidatorFactory.Contract.GetAllValidatorLength(&_ValidatorFactory.CallOpts)
}

// GetProviderChallengeInfo is a free data retrieval call binding the contract method 0x1c66250f.
//
// Solidity: function getProviderChallengeInfo(address provider_owner) view returns((address,address,uint256,string,uint256,uint256,uint8,uint256,uint256,uint256,uint256))
func (_ValidatorFactory *ValidatorFactoryCaller) GetProviderChallengeInfo(opts *bind.CallOpts, provider_owner common.Address) (ValidatorFactoryproviderChallengeInfo, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getProviderChallengeInfo", provider_owner)

	if err != nil {
		return *new(ValidatorFactoryproviderChallengeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorFactoryproviderChallengeInfo)).(*ValidatorFactoryproviderChallengeInfo)

	return out0, err

}

// GetProviderChallengeInfo is a free data retrieval call binding the contract method 0x1c66250f.
//
// Solidity: function getProviderChallengeInfo(address provider_owner) view returns((address,address,uint256,string,uint256,uint256,uint8,uint256,uint256,uint256,uint256))
func (_ValidatorFactory *ValidatorFactorySession) GetProviderChallengeInfo(provider_owner common.Address) (ValidatorFactoryproviderChallengeInfo, error) {
	return _ValidatorFactory.Contract.GetProviderChallengeInfo(&_ValidatorFactory.CallOpts, provider_owner)
}

// GetProviderChallengeInfo is a free data retrieval call binding the contract method 0x1c66250f.
//
// Solidity: function getProviderChallengeInfo(address provider_owner) view returns((address,address,uint256,string,uint256,uint256,uint8,uint256,uint256,uint256,uint256))
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetProviderChallengeInfo(provider_owner common.Address) (ValidatorFactoryproviderChallengeInfo, error) {
	return _ValidatorFactory.Contract.GetProviderChallengeInfo(&_ValidatorFactory.CallOpts, provider_owner)
}

// GetPunishAmount is a free data retrieval call binding the contract method 0xdbeffe26.
//
// Solidity: function getPunishAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) GetPunishAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "getPunishAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishAmount is a free data retrieval call binding the contract method 0xdbeffe26.
//
// Solidity: function getPunishAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) GetPunishAmount() (*big.Int, error) {
	return _ValidatorFactory.Contract.GetPunishAmount(&_ValidatorFactory.CallOpts)
}

// GetPunishAmount is a free data retrieval call binding the contract method 0xdbeffe26.
//
// Solidity: function getPunishAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) GetPunishAmount() (*big.Int, error) {
	return _ValidatorFactory.Contract.GetPunishAmount(&_ValidatorFactory.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ValidatorFactory *ValidatorFactoryCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ValidatorFactory *ValidatorFactorySession) Initialized() (bool, error) {
	return _ValidatorFactory.Contract.Initialized(&_ValidatorFactory.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ValidatorFactory *ValidatorFactoryCallerSession) Initialized() (bool, error) {
	return _ValidatorFactory.Contract.Initialized(&_ValidatorFactory.CallOpts)
}

// MaxChallengePercent is a free data retrieval call binding the contract method 0x82fa9588.
//
// Solidity: function max_challenge_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) MaxChallengePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "max_challenge_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxChallengePercent is a free data retrieval call binding the contract method 0x82fa9588.
//
// Solidity: function max_challenge_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) MaxChallengePercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxChallengePercent(&_ValidatorFactory.CallOpts)
}

// MaxChallengePercent is a free data retrieval call binding the contract method 0x82fa9588.
//
// Solidity: function max_challenge_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) MaxChallengePercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxChallengePercent(&_ValidatorFactory.CallOpts)
}

// MaxChallengeTime is a free data retrieval call binding the contract method 0x7b16da06.
//
// Solidity: function max_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) MaxChallengeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "max_challenge_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxChallengeTime is a free data retrieval call binding the contract method 0x7b16da06.
//
// Solidity: function max_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) MaxChallengeTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxChallengeTime(&_ValidatorFactory.CallOpts)
}

// MaxChallengeTime is a free data retrieval call binding the contract method 0x7b16da06.
//
// Solidity: function max_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) MaxChallengeTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxChallengeTime(&_ValidatorFactory.CallOpts)
}

// MaxProviderStartChallengeTime is a free data retrieval call binding the contract method 0x58c8ca08.
//
// Solidity: function max_provider_start_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) MaxProviderStartChallengeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "max_provider_start_challenge_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxProviderStartChallengeTime is a free data retrieval call binding the contract method 0x58c8ca08.
//
// Solidity: function max_provider_start_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) MaxProviderStartChallengeTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxProviderStartChallengeTime(&_ValidatorFactory.CallOpts)
}

// MaxProviderStartChallengeTime is a free data retrieval call binding the contract method 0x58c8ca08.
//
// Solidity: function max_provider_start_challenge_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) MaxProviderStartChallengeTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxProviderStartChallengeTime(&_ValidatorFactory.CallOpts)
}

// MaxValidatorCount is a free data retrieval call binding the contract method 0xf94fb56b.
//
// Solidity: function max_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) MaxValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "max_validator_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCount is a free data retrieval call binding the contract method 0xf94fb56b.
//
// Solidity: function max_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) MaxValidatorCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxValidatorCount(&_ValidatorFactory.CallOpts)
}

// MaxValidatorCount is a free data retrieval call binding the contract method 0xf94fb56b.
//
// Solidity: function max_validator_count() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) MaxValidatorCount() (*big.Int, error) {
	return _ValidatorFactory.Contract.MaxValidatorCount(&_ValidatorFactory.CallOpts)
}

// OwnerValidator is a free data retrieval call binding the contract method 0xb68f9ece.
//
// Solidity: function owner_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) OwnerValidator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "owner_validator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerValidator is a free data retrieval call binding the contract method 0xb68f9ece.
//
// Solidity: function owner_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) OwnerValidator(arg0 common.Address) (common.Address, error) {
	return _ValidatorFactory.Contract.OwnerValidator(&_ValidatorFactory.CallOpts, arg0)
}

// OwnerValidator is a free data retrieval call binding the contract method 0xb68f9ece.
//
// Solidity: function owner_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) OwnerValidator(arg0 common.Address) (common.Address, error) {
	return _ValidatorFactory.Contract.OwnerValidator(&_ValidatorFactory.CallOpts, arg0)
}

// ProviderChallengeInfo is a free data retrieval call binding the contract method 0x5c4830df.
//
// Solidity: function provider_challenge_info(address , uint256 ) view returns(address provider, address challenge_validator, uint256 md5_seed, string url, uint256 create_challenge_time, uint256 challenge_finish_time, uint8 state, uint256 challenge_amount, uint256 seed, uint256 root_hash, uint256 index)
func (_ValidatorFactory *ValidatorFactoryCaller) ProviderChallengeInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Provider            common.Address
	ChallengeValidator  common.Address
	Md5Seed             *big.Int
	Url                 string
	CreateChallengeTime *big.Int
	ChallengeFinishTime *big.Int
	State               uint8
	ChallengeAmount     *big.Int
	Seed                *big.Int
	RootHash            *big.Int
	Index               *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "provider_challenge_info", arg0, arg1)

	outstruct := new(struct {
		Provider            common.Address
		ChallengeValidator  common.Address
		Md5Seed             *big.Int
		Url                 string
		CreateChallengeTime *big.Int
		ChallengeFinishTime *big.Int
		State               uint8
		ChallengeAmount     *big.Int
		Seed                *big.Int
		RootHash            *big.Int
		Index               *big.Int
	})

	outstruct.Provider = out[0].(common.Address)
	outstruct.ChallengeValidator = out[1].(common.Address)
	outstruct.Md5Seed = out[2].(*big.Int)
	outstruct.Url = out[3].(string)
	outstruct.CreateChallengeTime = out[4].(*big.Int)
	outstruct.ChallengeFinishTime = out[5].(*big.Int)
	outstruct.State = out[6].(uint8)
	outstruct.ChallengeAmount = out[7].(*big.Int)
	outstruct.Seed = out[8].(*big.Int)
	outstruct.RootHash = out[9].(*big.Int)
	outstruct.Index = out[10].(*big.Int)

	return *outstruct, err

}

// ProviderChallengeInfo is a free data retrieval call binding the contract method 0x5c4830df.
//
// Solidity: function provider_challenge_info(address , uint256 ) view returns(address provider, address challenge_validator, uint256 md5_seed, string url, uint256 create_challenge_time, uint256 challenge_finish_time, uint8 state, uint256 challenge_amount, uint256 seed, uint256 root_hash, uint256 index)
func (_ValidatorFactory *ValidatorFactorySession) ProviderChallengeInfo(arg0 common.Address, arg1 *big.Int) (struct {
	Provider            common.Address
	ChallengeValidator  common.Address
	Md5Seed             *big.Int
	Url                 string
	CreateChallengeTime *big.Int
	ChallengeFinishTime *big.Int
	State               uint8
	ChallengeAmount     *big.Int
	Seed                *big.Int
	RootHash            *big.Int
	Index               *big.Int
}, error) {
	return _ValidatorFactory.Contract.ProviderChallengeInfo(&_ValidatorFactory.CallOpts, arg0, arg1)
}

// ProviderChallengeInfo is a free data retrieval call binding the contract method 0x5c4830df.
//
// Solidity: function provider_challenge_info(address , uint256 ) view returns(address provider, address challenge_validator, uint256 md5_seed, string url, uint256 create_challenge_time, uint256 challenge_finish_time, uint8 state, uint256 challenge_amount, uint256 seed, uint256 root_hash, uint256 index)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ProviderChallengeInfo(arg0 common.Address, arg1 *big.Int) (struct {
	Provider            common.Address
	ChallengeValidator  common.Address
	Md5Seed             *big.Int
	Url                 string
	CreateChallengeTime *big.Int
	ChallengeFinishTime *big.Int
	State               uint8
	ChallengeAmount     *big.Int
	Seed                *big.Int
	RootHash            *big.Int
	Index               *big.Int
}, error) {
	return _ValidatorFactory.Contract.ProviderChallengeInfo(&_ValidatorFactory.CallOpts, arg0, arg1)
}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) ProviderFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "provider_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) ProviderFactory() (common.Address, error) {
	return _ValidatorFactory.Contract.ProviderFactory(&_ValidatorFactory.CallOpts)
}

// ProviderFactory is a free data retrieval call binding the contract method 0x64c76465.
//
// Solidity: function provider_factory() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ProviderFactory() (common.Address, error) {
	return _ValidatorFactory.Contract.ProviderFactory(&_ValidatorFactory.CallOpts)
}

// ProviderIndex is a free data retrieval call binding the contract method 0xa7693767.
//
// Solidity: function provider_index(address ) view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ProviderIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "provider_index", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderIndex is a free data retrieval call binding the contract method 0xa7693767.
//
// Solidity: function provider_index(address ) view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ProviderIndex(arg0 common.Address) (*big.Int, error) {
	return _ValidatorFactory.Contract.ProviderIndex(&_ValidatorFactory.CallOpts, arg0)
}

// ProviderIndex is a free data retrieval call binding the contract method 0xa7693767.
//
// Solidity: function provider_index(address ) view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ProviderIndex(arg0 common.Address) (*big.Int, error) {
	return _ValidatorFactory.Contract.ProviderIndex(&_ValidatorFactory.CallOpts, arg0)
}

// ProviderLastChallengeState is a free data retrieval call binding the contract method 0x6cbd328c.
//
// Solidity: function provider_last_challenge_state(address ) view returns(uint8)
func (_ValidatorFactory *ValidatorFactoryCaller) ProviderLastChallengeState(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "provider_last_challenge_state", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProviderLastChallengeState is a free data retrieval call binding the contract method 0x6cbd328c.
//
// Solidity: function provider_last_challenge_state(address ) view returns(uint8)
func (_ValidatorFactory *ValidatorFactorySession) ProviderLastChallengeState(arg0 common.Address) (uint8, error) {
	return _ValidatorFactory.Contract.ProviderLastChallengeState(&_ValidatorFactory.CallOpts, arg0)
}

// ProviderLastChallengeState is a free data retrieval call binding the contract method 0x6cbd328c.
//
// Solidity: function provider_last_challenge_state(address ) view returns(uint8)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ProviderLastChallengeState(arg0 common.Address) (uint8, error) {
	return _ValidatorFactory.Contract.ProviderLastChallengeState(&_ValidatorFactory.CallOpts, arg0)
}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) PunishAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "punish_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) PunishAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.PunishAddress(&_ValidatorFactory.CallOpts)
}

// PunishAddress is a free data retrieval call binding the contract method 0x18f508ce.
//
// Solidity: function punish_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) PunishAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.PunishAddress(&_ValidatorFactory.CallOpts)
}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) PunishAllPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "punish_all_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) PunishAllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.PunishAllPercent(&_ValidatorFactory.CallOpts)
}

// PunishAllPercent is a free data retrieval call binding the contract method 0xbefc4512.
//
// Solidity: function punish_all_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) PunishAllPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.PunishAllPercent(&_ValidatorFactory.CallOpts)
}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) PunishPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "punish_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) PunishPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.PunishPercent(&_ValidatorFactory.CallOpts)
}

// PunishPercent is a free data retrieval call binding the contract method 0xeb3359fc.
//
// Solidity: function punish_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) PunishPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.PunishPercent(&_ValidatorFactory.CallOpts)
}

// TeamAddress is a free data retrieval call binding the contract method 0x8f32cf0c.
//
// Solidity: function team_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) TeamAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "team_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeamAddress is a free data retrieval call binding the contract method 0x8f32cf0c.
//
// Solidity: function team_address() view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) TeamAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.TeamAddress(&_ValidatorFactory.CallOpts)
}

// TeamAddress is a free data retrieval call binding the contract method 0x8f32cf0c.
//
// Solidity: function team_address() view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) TeamAddress() (common.Address, error) {
	return _ValidatorFactory.Contract.TeamAddress(&_ValidatorFactory.CallOpts)
}

// TeamPercent is a free data retrieval call binding the contract method 0xb916847c.
//
// Solidity: function team_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) TeamPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "team_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamPercent is a free data retrieval call binding the contract method 0xb916847c.
//
// Solidity: function team_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) TeamPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.TeamPercent(&_ValidatorFactory.CallOpts)
}

// TeamPercent is a free data retrieval call binding the contract method 0xb916847c.
//
// Solidity: function team_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) TeamPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.TeamPercent(&_ValidatorFactory.CallOpts)
}

// ValidatorLockTime is a free data retrieval call binding the contract method 0x0b8f4329.
//
// Solidity: function validator_lock_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ValidatorLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "validator_lock_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLockTime is a free data retrieval call binding the contract method 0x0b8f4329.
//
// Solidity: function validator_lock_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ValidatorLockTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorLockTime(&_ValidatorFactory.CallOpts)
}

// ValidatorLockTime is a free data retrieval call binding the contract method 0x0b8f4329.
//
// Solidity: function validator_lock_time() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ValidatorLockTime() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorLockTime(&_ValidatorFactory.CallOpts)
}

// ValidatorPercent is a free data retrieval call binding the contract method 0x0b7db192.
//
// Solidity: function validator_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ValidatorPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "validator_percent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPercent is a free data retrieval call binding the contract method 0x0b7db192.
//
// Solidity: function validator_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ValidatorPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPercent(&_ValidatorFactory.CallOpts)
}

// ValidatorPercent is a free data retrieval call binding the contract method 0x0b7db192.
//
// Solidity: function validator_percent() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ValidatorPercent() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPercent(&_ValidatorFactory.CallOpts)
}

// ValidatorPledgeAmount is a free data retrieval call binding the contract method 0x0d3318bc.
//
// Solidity: function validator_pledgeAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ValidatorPledgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "validator_pledgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPledgeAmount is a free data retrieval call binding the contract method 0x0d3318bc.
//
// Solidity: function validator_pledgeAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ValidatorPledgeAmount() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPledgeAmount(&_ValidatorFactory.CallOpts)
}

// ValidatorPledgeAmount is a free data retrieval call binding the contract method 0x0d3318bc.
//
// Solidity: function validator_pledgeAmount() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ValidatorPledgeAmount() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPledgeAmount(&_ValidatorFactory.CallOpts)
}

// ValidatorPunishInterval is a free data retrieval call binding the contract method 0xe9807f63.
//
// Solidity: function validator_punish_interval() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ValidatorPunishInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "validator_punish_interval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPunishInterval is a free data retrieval call binding the contract method 0xe9807f63.
//
// Solidity: function validator_punish_interval() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ValidatorPunishInterval() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPunishInterval(&_ValidatorFactory.CallOpts)
}

// ValidatorPunishInterval is a free data retrieval call binding the contract method 0xe9807f63.
//
// Solidity: function validator_punish_interval() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ValidatorPunishInterval() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPunishInterval(&_ValidatorFactory.CallOpts)
}

// ValidatorPunishStartLimit is a free data retrieval call binding the contract method 0x9c59499d.
//
// Solidity: function validator_punish_start_limit() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCaller) ValidatorPunishStartLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "validator_punish_start_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPunishStartLimit is a free data retrieval call binding the contract method 0x9c59499d.
//
// Solidity: function validator_punish_start_limit() view returns(uint256)
func (_ValidatorFactory *ValidatorFactorySession) ValidatorPunishStartLimit() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPunishStartLimit(&_ValidatorFactory.CallOpts)
}

// ValidatorPunishStartLimit is a free data retrieval call binding the contract method 0x9c59499d.
//
// Solidity: function validator_punish_start_limit() view returns(uint256)
func (_ValidatorFactory *ValidatorFactoryCallerSession) ValidatorPunishStartLimit() (*big.Int, error) {
	return _ValidatorFactory.Contract.ValidatorPunishStartLimit(&_ValidatorFactory.CallOpts)
}

// WhiteListValidator is a free data retrieval call binding the contract method 0x73bc8c7f.
//
// Solidity: function whiteList_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCaller) WhiteListValidator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ValidatorFactory.contract.Call(opts, &out, "whiteList_validator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhiteListValidator is a free data retrieval call binding the contract method 0x73bc8c7f.
//
// Solidity: function whiteList_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactorySession) WhiteListValidator(arg0 common.Address) (common.Address, error) {
	return _ValidatorFactory.Contract.WhiteListValidator(&_ValidatorFactory.CallOpts, arg0)
}

// WhiteListValidator is a free data retrieval call binding the contract method 0x73bc8c7f.
//
// Solidity: function whiteList_validator(address ) view returns(address)
func (_ValidatorFactory *ValidatorFactoryCallerSession) WhiteListValidator(arg0 common.Address) (common.Address, error) {
	return _ValidatorFactory.Contract.WhiteListValidator(&_ValidatorFactory.CallOpts, arg0)
}

// MarginCalls is a paid mutator transaction binding the contract method 0x5b3327d7.
//
// Solidity: function MarginCalls() payable returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) MarginCalls(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "MarginCalls")
}

// MarginCalls is a paid mutator transaction binding the contract method 0x5b3327d7.
//
// Solidity: function MarginCalls() payable returns()
func (_ValidatorFactory *ValidatorFactorySession) MarginCalls() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.MarginCalls(&_ValidatorFactory.TransactOpts)
}

// MarginCalls is a paid mutator transaction binding the contract method 0x5b3327d7.
//
// Solidity: function MarginCalls() payable returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) MarginCalls() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.MarginCalls(&_ValidatorFactory.TransactOpts)
}

// ChallengeFinish is a paid mutator transaction binding the contract method 0x9747b38b.
//
// Solidity: function challengeFinish(address provider, uint256 seed, uint256 challenge_amount, uint256 root_hash, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChallengeFinish(opts *bind.TransactOpts, provider common.Address, seed *big.Int, challenge_amount *big.Int, root_hash *big.Int, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "challengeFinish", provider, seed, challenge_amount, root_hash, _state)
}

// ChallengeFinish is a paid mutator transaction binding the contract method 0x9747b38b.
//
// Solidity: function challengeFinish(address provider, uint256 seed, uint256 challenge_amount, uint256 root_hash, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChallengeFinish(provider common.Address, seed *big.Int, challenge_amount *big.Int, root_hash *big.Int, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChallengeFinish(&_ValidatorFactory.TransactOpts, provider, seed, challenge_amount, root_hash, _state)
}

// ChallengeFinish is a paid mutator transaction binding the contract method 0x9747b38b.
//
// Solidity: function challengeFinish(address provider, uint256 seed, uint256 challenge_amount, uint256 root_hash, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChallengeFinish(provider common.Address, seed *big.Int, challenge_amount *big.Int, root_hash *big.Int, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChallengeFinish(&_ValidatorFactory.TransactOpts, provider, seed, challenge_amount, root_hash, _state)
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0x3fbe464f.
//
// Solidity: function challengeProvider(address provider, uint256 md5_seed, string url) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChallengeProvider(opts *bind.TransactOpts, provider common.Address, md5_seed *big.Int, url string) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "challengeProvider", provider, md5_seed, url)
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0x3fbe464f.
//
// Solidity: function challengeProvider(address provider, uint256 md5_seed, string url) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChallengeProvider(provider common.Address, md5_seed *big.Int, url string) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChallengeProvider(&_ValidatorFactory.TransactOpts, provider, md5_seed, url)
}

// ChallengeProvider is a paid mutator transaction binding the contract method 0x3fbe464f.
//
// Solidity: function challengeProvider(address provider, uint256 md5_seed, string url) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChallengeProvider(provider common.Address, md5_seed *big.Int, url string) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChallengeProvider(&_ValidatorFactory.TransactOpts, provider, md5_seed, url)
}

// ChangeAdminAddress is a paid mutator transaction binding the contract method 0x1021688f.
//
// Solidity: function changeAdminAddress(address _new_admin) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeAdminAddress(opts *bind.TransactOpts, _new_admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeAdminAddress", _new_admin)
}

// ChangeAdminAddress is a paid mutator transaction binding the contract method 0x1021688f.
//
// Solidity: function changeAdminAddress(address _new_admin) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeAdminAddress(_new_admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeAdminAddress(&_ValidatorFactory.TransactOpts, _new_admin)
}

// ChangeAdminAddress is a paid mutator transaction binding the contract method 0x1021688f.
//
// Solidity: function changeAdminAddress(address _new_admin) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeAdminAddress(_new_admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeAdminAddress(&_ValidatorFactory.TransactOpts, _new_admin)
}

// ChangeChallengeSdlTrxID is a paid mutator transaction binding the contract method 0x497440bf.
//
// Solidity: function changeChallengeSdlTrxID(uint256 _new_trx_id) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeChallengeSdlTrxID(opts *bind.TransactOpts, _new_trx_id *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeChallengeSdlTrxID", _new_trx_id)
}

// ChangeChallengeSdlTrxID is a paid mutator transaction binding the contract method 0x497440bf.
//
// Solidity: function changeChallengeSdlTrxID(uint256 _new_trx_id) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeChallengeSdlTrxID(_new_trx_id *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeChallengeSdlTrxID(&_ValidatorFactory.TransactOpts, _new_trx_id)
}

// ChangeChallengeSdlTrxID is a paid mutator transaction binding the contract method 0x497440bf.
//
// Solidity: function changeChallengeSdlTrxID(uint256 _new_trx_id) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeChallengeSdlTrxID(_new_trx_id *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeChallengeSdlTrxID(&_ValidatorFactory.TransactOpts, _new_trx_id)
}

// ChangeMaxChallengeParam is a paid mutator transaction binding the contract method 0xa159d942.
//
// Solidity: function changeMaxChallengeParam(uint256 _max_challenge_percent, uint256 _challenge_all_percent, uint256 _max_challenge_time, uint256 _max_provider_start_challenge_time) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeMaxChallengeParam(opts *bind.TransactOpts, _max_challenge_percent *big.Int, _challenge_all_percent *big.Int, _max_challenge_time *big.Int, _max_provider_start_challenge_time *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeMaxChallengeParam", _max_challenge_percent, _challenge_all_percent, _max_challenge_time, _max_provider_start_challenge_time)
}

// ChangeMaxChallengeParam is a paid mutator transaction binding the contract method 0xa159d942.
//
// Solidity: function changeMaxChallengeParam(uint256 _max_challenge_percent, uint256 _challenge_all_percent, uint256 _max_challenge_time, uint256 _max_provider_start_challenge_time) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeMaxChallengeParam(_max_challenge_percent *big.Int, _challenge_all_percent *big.Int, _max_challenge_time *big.Int, _max_provider_start_challenge_time *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeMaxChallengeParam(&_ValidatorFactory.TransactOpts, _max_challenge_percent, _challenge_all_percent, _max_challenge_time, _max_provider_start_challenge_time)
}

// ChangeMaxChallengeParam is a paid mutator transaction binding the contract method 0xa159d942.
//
// Solidity: function changeMaxChallengeParam(uint256 _max_challenge_percent, uint256 _challenge_all_percent, uint256 _max_challenge_time, uint256 _max_provider_start_challenge_time) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeMaxChallengeParam(_max_challenge_percent *big.Int, _challenge_all_percent *big.Int, _max_challenge_time *big.Int, _max_provider_start_challenge_time *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeMaxChallengeParam(&_ValidatorFactory.TransactOpts, _max_challenge_percent, _challenge_all_percent, _max_challenge_time, _max_provider_start_challenge_time)
}

// ChangeMaxValidatorCount is a paid mutator transaction binding the contract method 0x46bae458.
//
// Solidity: function changeMaxValidatorCount(uint256 _max_validator_count) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeMaxValidatorCount(opts *bind.TransactOpts, _max_validator_count *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeMaxValidatorCount", _max_validator_count)
}

// ChangeMaxValidatorCount is a paid mutator transaction binding the contract method 0x46bae458.
//
// Solidity: function changeMaxValidatorCount(uint256 _max_validator_count) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeMaxValidatorCount(_max_validator_count *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeMaxValidatorCount(&_ValidatorFactory.TransactOpts, _max_validator_count)
}

// ChangeMaxValidatorCount is a paid mutator transaction binding the contract method 0x46bae458.
//
// Solidity: function changeMaxValidatorCount(uint256 _max_validator_count) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeMaxValidatorCount(_max_validator_count *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeMaxValidatorCount(&_ValidatorFactory.TransactOpts, _max_validator_count)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangePunishAddress(opts *bind.TransactOpts, _punish_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changePunishAddress", _punish_address)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangePunishAddress(_punish_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangePunishAddress(&_ValidatorFactory.TransactOpts, _punish_address)
}

// ChangePunishAddress is a paid mutator transaction binding the contract method 0x467a01ad.
//
// Solidity: function changePunishAddress(address _punish_address) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangePunishAddress(_punish_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangePunishAddress(&_ValidatorFactory.TransactOpts, _punish_address)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangePunishPercent(opts *bind.TransactOpts, _new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changePunishPercent", _new_punish_percent, _new_punish_all_percent)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangePunishPercent(_new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangePunishPercent(&_ValidatorFactory.TransactOpts, _new_punish_percent, _new_punish_all_percent)
}

// ChangePunishPercent is a paid mutator transaction binding the contract method 0x4fb6ecda.
//
// Solidity: function changePunishPercent(uint256 _new_punish_percent, uint256 _new_punish_all_percent) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangePunishPercent(_new_punish_percent *big.Int, _new_punish_all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangePunishPercent(&_ValidatorFactory.TransactOpts, _new_punish_percent, _new_punish_all_percent)
}

// ChangeRewardPercent is a paid mutator transaction binding the contract method 0xb0b64768.
//
// Solidity: function changeRewardPercent(uint256 _team_percent, uint256 _validator_percent, uint256 _all_percent) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeRewardPercent(opts *bind.TransactOpts, _team_percent *big.Int, _validator_percent *big.Int, _all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeRewardPercent", _team_percent, _validator_percent, _all_percent)
}

// ChangeRewardPercent is a paid mutator transaction binding the contract method 0xb0b64768.
//
// Solidity: function changeRewardPercent(uint256 _team_percent, uint256 _validator_percent, uint256 _all_percent) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeRewardPercent(_team_percent *big.Int, _validator_percent *big.Int, _all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeRewardPercent(&_ValidatorFactory.TransactOpts, _team_percent, _validator_percent, _all_percent)
}

// ChangeRewardPercent is a paid mutator transaction binding the contract method 0xb0b64768.
//
// Solidity: function changeRewardPercent(uint256 _team_percent, uint256 _validator_percent, uint256 _all_percent) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeRewardPercent(_team_percent *big.Int, _validator_percent *big.Int, _all_percent *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeRewardPercent(&_ValidatorFactory.TransactOpts, _team_percent, _validator_percent, _all_percent)
}

// ChangeTeamAddress is a paid mutator transaction binding the contract method 0x3aee69bb.
//
// Solidity: function changeTeamAddress(address _team_address) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeTeamAddress(opts *bind.TransactOpts, _team_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeTeamAddress", _team_address)
}

// ChangeTeamAddress is a paid mutator transaction binding the contract method 0x3aee69bb.
//
// Solidity: function changeTeamAddress(address _team_address) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeTeamAddress(_team_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeTeamAddress(&_ValidatorFactory.TransactOpts, _team_address)
}

// ChangeTeamAddress is a paid mutator transaction binding the contract method 0x3aee69bb.
//
// Solidity: function changeTeamAddress(address _team_address) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeTeamAddress(_team_address common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeTeamAddress(&_ValidatorFactory.TransactOpts, _team_address)
}

// ChangeValidatorLockTime is a paid mutator transaction binding the contract method 0xe2d42a57.
//
// Solidity: function changeValidatorLockTime(uint256 _new_lock) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeValidatorLockTime(opts *bind.TransactOpts, _new_lock *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeValidatorLockTime", _new_lock)
}

// ChangeValidatorLockTime is a paid mutator transaction binding the contract method 0xe2d42a57.
//
// Solidity: function changeValidatorLockTime(uint256 _new_lock) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeValidatorLockTime(_new_lock *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorLockTime(&_ValidatorFactory.TransactOpts, _new_lock)
}

// ChangeValidatorLockTime is a paid mutator transaction binding the contract method 0xe2d42a57.
//
// Solidity: function changeValidatorLockTime(uint256 _new_lock) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeValidatorLockTime(_new_lock *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorLockTime(&_ValidatorFactory.TransactOpts, _new_lock)
}

// ChangeValidatorMinPledgeAmount is a paid mutator transaction binding the contract method 0xdac0bd20.
//
// Solidity: function changeValidatorMinPledgeAmount(uint256 _validator_min_pledgeAmount) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeValidatorMinPledgeAmount(opts *bind.TransactOpts, _validator_min_pledgeAmount *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeValidatorMinPledgeAmount", _validator_min_pledgeAmount)
}

// ChangeValidatorMinPledgeAmount is a paid mutator transaction binding the contract method 0xdac0bd20.
//
// Solidity: function changeValidatorMinPledgeAmount(uint256 _validator_min_pledgeAmount) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeValidatorMinPledgeAmount(_validator_min_pledgeAmount *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorMinPledgeAmount(&_ValidatorFactory.TransactOpts, _validator_min_pledgeAmount)
}

// ChangeValidatorMinPledgeAmount is a paid mutator transaction binding the contract method 0xdac0bd20.
//
// Solidity: function changeValidatorMinPledgeAmount(uint256 _validator_min_pledgeAmount) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeValidatorMinPledgeAmount(_validator_min_pledgeAmount *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorMinPledgeAmount(&_ValidatorFactory.TransactOpts, _validator_min_pledgeAmount)
}

// ChangeValidatorPunishInterval is a paid mutator transaction binding the contract method 0xdde11573.
//
// Solidity: function changeValidatorPunishInterval(uint256 _new_interval) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeValidatorPunishInterval(opts *bind.TransactOpts, _new_interval *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeValidatorPunishInterval", _new_interval)
}

// ChangeValidatorPunishInterval is a paid mutator transaction binding the contract method 0xdde11573.
//
// Solidity: function changeValidatorPunishInterval(uint256 _new_interval) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeValidatorPunishInterval(_new_interval *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorPunishInterval(&_ValidatorFactory.TransactOpts, _new_interval)
}

// ChangeValidatorPunishInterval is a paid mutator transaction binding the contract method 0xdde11573.
//
// Solidity: function changeValidatorPunishInterval(uint256 _new_interval) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeValidatorPunishInterval(_new_interval *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorPunishInterval(&_ValidatorFactory.TransactOpts, _new_interval)
}

// ChangeValidatorPunishStartTime is a paid mutator transaction binding the contract method 0xee8d7adb.
//
// Solidity: function changeValidatorPunishStartTime(uint256 _new_start_limit) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeValidatorPunishStartTime(opts *bind.TransactOpts, _new_start_limit *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeValidatorPunishStartTime", _new_start_limit)
}

// ChangeValidatorPunishStartTime is a paid mutator transaction binding the contract method 0xee8d7adb.
//
// Solidity: function changeValidatorPunishStartTime(uint256 _new_start_limit) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeValidatorPunishStartTime(_new_start_limit *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorPunishStartTime(&_ValidatorFactory.TransactOpts, _new_start_limit)
}

// ChangeValidatorPunishStartTime is a paid mutator transaction binding the contract method 0xee8d7adb.
//
// Solidity: function changeValidatorPunishStartTime(uint256 _new_start_limit) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeValidatorPunishStartTime(_new_start_limit *big.Int) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorPunishStartTime(&_ValidatorFactory.TransactOpts, _new_start_limit)
}

// ChangeValidatorState is a paid mutator transaction binding the contract method 0xbd6ad9d0.
//
// Solidity: function changeValidatorState(address validator, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ChangeValidatorState(opts *bind.TransactOpts, validator common.Address, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "changeValidatorState", validator, _state)
}

// ChangeValidatorState is a paid mutator transaction binding the contract method 0xbd6ad9d0.
//
// Solidity: function changeValidatorState(address validator, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactorySession) ChangeValidatorState(validator common.Address, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorState(&_ValidatorFactory.TransactOpts, validator, _state)
}

// ChangeValidatorState is a paid mutator transaction binding the contract method 0xbd6ad9d0.
//
// Solidity: function changeValidatorState(address validator, uint8 _state) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ChangeValidatorState(validator common.Address, _state uint8) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ChangeValidatorState(&_ValidatorFactory.TransactOpts, validator, _state)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xd125ca18.
//
// Solidity: function createValidator() payable returns(address)
func (_ValidatorFactory *ValidatorFactoryTransactor) CreateValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "createValidator")
}

// CreateValidator is a paid mutator transaction binding the contract method 0xd125ca18.
//
// Solidity: function createValidator() payable returns(address)
func (_ValidatorFactory *ValidatorFactorySession) CreateValidator() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.CreateValidator(&_ValidatorFactory.TransactOpts)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xd125ca18.
//
// Solidity: function createValidator() payable returns(address)
func (_ValidatorFactory *ValidatorFactoryTransactorSession) CreateValidator() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.CreateValidator(&_ValidatorFactory.TransactOpts)
}

// ExitProduceBlock is a paid mutator transaction binding the contract method 0xab4e9fc9.
//
// Solidity: function exitProduceBlock() returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ExitProduceBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "exitProduceBlock")
}

// ExitProduceBlock is a paid mutator transaction binding the contract method 0xab4e9fc9.
//
// Solidity: function exitProduceBlock() returns()
func (_ValidatorFactory *ValidatorFactorySession) ExitProduceBlock() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ExitProduceBlock(&_ValidatorFactory.TransactOpts)
}

// ExitProduceBlock is a paid mutator transaction binding the contract method 0xab4e9fc9.
//
// Solidity: function exitProduceBlock() returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ExitProduceBlock() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ExitProduceBlock(&_ValidatorFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _init_validator, address _admin) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) Initialize(opts *bind.TransactOpts, _init_validator []common.Address, _admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "initialize", _init_validator, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _init_validator, address _admin) returns()
func (_ValidatorFactory *ValidatorFactorySession) Initialize(_init_validator []common.Address, _admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.Initialize(&_ValidatorFactory.TransactOpts, _init_validator, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _init_validator, address _admin) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) Initialize(_init_validator []common.Address, _admin common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.Initialize(&_ValidatorFactory.TransactOpts, _init_validator, _admin)
}

// RemoveRankingList is a paid mutator transaction binding the contract method 0x54452901.
//
// Solidity: function removeRankingList() returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) RemoveRankingList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "removeRankingList")
}

// RemoveRankingList is a paid mutator transaction binding the contract method 0x54452901.
//
// Solidity: function removeRankingList() returns()
func (_ValidatorFactory *ValidatorFactorySession) RemoveRankingList() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.RemoveRankingList(&_ValidatorFactory.TransactOpts)
}

// RemoveRankingList is a paid mutator transaction binding the contract method 0x54452901.
//
// Solidity: function removeRankingList() returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) RemoveRankingList() (*types.Transaction, error) {
	return _ValidatorFactory.Contract.RemoveRankingList(&_ValidatorFactory.TransactOpts)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address val) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) TryPunish(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "tryPunish", val)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address val) returns()
func (_ValidatorFactory *ValidatorFactorySession) TryPunish(val common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.TryPunish(&_ValidatorFactory.TransactOpts, val)
}

// TryPunish is a paid mutator transaction binding the contract method 0x132d8c25.
//
// Solidity: function tryPunish(address val) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) TryPunish(val common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.TryPunish(&_ValidatorFactory.TransactOpts, val)
}

// ValidatorNotSubmitResult is a paid mutator transaction binding the contract method 0x79d26934.
//
// Solidity: function validatorNotSubmitResult(address provider) returns()
func (_ValidatorFactory *ValidatorFactoryTransactor) ValidatorNotSubmitResult(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.contract.Transact(opts, "validatorNotSubmitResult", provider)
}

// ValidatorNotSubmitResult is a paid mutator transaction binding the contract method 0x79d26934.
//
// Solidity: function validatorNotSubmitResult(address provider) returns()
func (_ValidatorFactory *ValidatorFactorySession) ValidatorNotSubmitResult(provider common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ValidatorNotSubmitResult(&_ValidatorFactory.TransactOpts, provider)
}

// ValidatorNotSubmitResult is a paid mutator transaction binding the contract method 0x79d26934.
//
// Solidity: function validatorNotSubmitResult(address provider) returns()
func (_ValidatorFactory *ValidatorFactoryTransactorSession) ValidatorNotSubmitResult(provider common.Address) (*types.Transaction, error) {
	return _ValidatorFactory.Contract.ValidatorNotSubmitResult(&_ValidatorFactory.TransactOpts, provider)
}

// ValidatorFactoryChallengeCreateIterator is returned from FilterChallengeCreate and is used to iterate over the raw logs and unpacked data for ChallengeCreate events raised by the ValidatorFactory contract.
type ValidatorFactoryChallengeCreateIterator struct {
	Event *ValidatorFactoryChallengeCreate // Event containing the contract specifics and raw log

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
func (it *ValidatorFactoryChallengeCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorFactoryChallengeCreate)
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
		it.Event = new(ValidatorFactoryChallengeCreate)
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
func (it *ValidatorFactoryChallengeCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorFactoryChallengeCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorFactoryChallengeCreate represents a ChallengeCreate event raised by the ValidatorFactory contract.
type ValidatorFactoryChallengeCreate struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChallengeCreate is a free log retrieval operation binding the contract event 0x578637e46a7b8909c55e7d4ae6c9c4e95ff81200abb4b0e194d5aa03302081cc.
//
// Solidity: event ChallengeCreate(address arg0, uint256 arg1)
func (_ValidatorFactory *ValidatorFactoryFilterer) FilterChallengeCreate(opts *bind.FilterOpts) (*ValidatorFactoryChallengeCreateIterator, error) {

	logs, sub, err := _ValidatorFactory.contract.FilterLogs(opts, "ChallengeCreate")
	if err != nil {
		return nil, err
	}
	return &ValidatorFactoryChallengeCreateIterator{contract: _ValidatorFactory.contract, event: "ChallengeCreate", logs: logs, sub: sub}, nil
}

// WatchChallengeCreate is a free log subscription operation binding the contract event 0x578637e46a7b8909c55e7d4ae6c9c4e95ff81200abb4b0e194d5aa03302081cc.
//
// Solidity: event ChallengeCreate(address arg0, uint256 arg1)
func (_ValidatorFactory *ValidatorFactoryFilterer) WatchChallengeCreate(opts *bind.WatchOpts, sink chan<- *ValidatorFactoryChallengeCreate) (event.Subscription, error) {

	logs, sub, err := _ValidatorFactory.contract.WatchLogs(opts, "ChallengeCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorFactoryChallengeCreate)
				if err := _ValidatorFactory.contract.UnpackLog(event, "ChallengeCreate", log); err != nil {
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

// ParseChallengeCreate is a log parse operation binding the contract event 0x578637e46a7b8909c55e7d4ae6c9c4e95ff81200abb4b0e194d5aa03302081cc.
//
// Solidity: event ChallengeCreate(address arg0, uint256 arg1)
func (_ValidatorFactory *ValidatorFactoryFilterer) ParseChallengeCreate(log types.Log) (*ValidatorFactoryChallengeCreate, error) {
	event := new(ValidatorFactoryChallengeCreate)
	if err := _ValidatorFactory.contract.UnpackLog(event, "ChallengeCreate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorFactoryChallengeEndIterator is returned from FilterChallengeEnd and is used to iterate over the raw logs and unpacked data for ChallengeEnd events raised by the ValidatorFactory contract.
type ValidatorFactoryChallengeEndIterator struct {
	Event *ValidatorFactoryChallengeEnd // Event containing the contract specifics and raw log

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
func (it *ValidatorFactoryChallengeEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorFactoryChallengeEnd)
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
		it.Event = new(ValidatorFactoryChallengeEnd)
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
func (it *ValidatorFactoryChallengeEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorFactoryChallengeEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorFactoryChallengeEnd represents a ChallengeEnd event raised by the ValidatorFactory contract.
type ValidatorFactoryChallengeEnd struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnd is a free log retrieval operation binding the contract event 0x1c32ca6921c33b35f5daf9698010cec08295d382f765de3d753088cc7b24a471.
//
// Solidity: event ChallengeEnd(address arg0)
func (_ValidatorFactory *ValidatorFactoryFilterer) FilterChallengeEnd(opts *bind.FilterOpts) (*ValidatorFactoryChallengeEndIterator, error) {

	logs, sub, err := _ValidatorFactory.contract.FilterLogs(opts, "ChallengeEnd")
	if err != nil {
		return nil, err
	}
	return &ValidatorFactoryChallengeEndIterator{contract: _ValidatorFactory.contract, event: "ChallengeEnd", logs: logs, sub: sub}, nil
}

// WatchChallengeEnd is a free log subscription operation binding the contract event 0x1c32ca6921c33b35f5daf9698010cec08295d382f765de3d753088cc7b24a471.
//
// Solidity: event ChallengeEnd(address arg0)
func (_ValidatorFactory *ValidatorFactoryFilterer) WatchChallengeEnd(opts *bind.WatchOpts, sink chan<- *ValidatorFactoryChallengeEnd) (event.Subscription, error) {

	logs, sub, err := _ValidatorFactory.contract.WatchLogs(opts, "ChallengeEnd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorFactoryChallengeEnd)
				if err := _ValidatorFactory.contract.UnpackLog(event, "ChallengeEnd", log); err != nil {
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

// ParseChallengeEnd is a log parse operation binding the contract event 0x1c32ca6921c33b35f5daf9698010cec08295d382f765de3d753088cc7b24a471.
//
// Solidity: event ChallengeEnd(address arg0)
func (_ValidatorFactory *ValidatorFactoryFilterer) ParseChallengeEnd(log types.Log) (*ValidatorFactoryChallengeEnd, error) {
	event := new(ValidatorFactoryChallengeEnd)
	if err := _ValidatorFactory.contract.UnpackLog(event, "ChallengeEnd", log); err != nil {
		return nil, err
	}
	return event, nil
}

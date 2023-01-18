package bid

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

const OrderFactoryABI = `[]`
const OrderBaseABI = `[{
      "inputs": [
        {
          "internalType": "address",
          "name": "_order_factory",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "provider_factory_",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "owner_",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "cpu_",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "cert_key_",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "sdl_trx_id_",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "order_number",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [],
      "name": "CanQuote",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "provider",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "cpu_price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_price",
              "type": "uint256"
            }
          ],
          "indexed": true,
          "internalType": "struct PriceOracle",
          "name": "price",
          "type": "tuple"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "final_price",
          "type": "uint256"
        }
      ],
      "name": "ChooseQuote",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "old_cpu",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "old_memory",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "old_storage",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "new_cpu",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "new_memory",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "new_storage",
          "type": "uint256"
        }
      ],
      "name": "DeployMentUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "DepositBalance",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "owner_",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "cpu",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "memory_",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "storage_",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "cert",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "sdl",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "order_number",
          "type": "uint256"
        }
      ],
      "name": "OrderCreate",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [],
      "name": "OrderEnded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "provider",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "PayBill",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "provider",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "cpu_price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "memory_price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "storage_price",
          "type": "uint256"
        }
      ],
      "name": "Quote",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "new_sdl_trx_id",
          "type": "uint256"
        }
      ],
      "name": "UpdateSDL",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [],
      "name": "UserCancelOrder",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "new_trx_hash",
          "type": "uint256"
        }
      ],
      "name": "change_sdl_trx_hash",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "quote_index",
          "type": "uint256"
        }
      ],
      "name": "choose_provider",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "deposit_balance",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "final_choice",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "final_price",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "last_pay_time",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_cert",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_cpu",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_memory",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_order_number",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_pending_sdl_trx_id",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_sdl_trx_id",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "o_storage",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "order_info",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "owner",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "v_cpu",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "v_memory",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "v_storage",
              "type": "uint256"
            },
            {
              "internalType": "string",
              "name": "cert_key",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "trx_id",
              "type": "uint256"
            },
            {
              "internalType": "uint8",
              "name": "state",
              "type": "uint8"
            }
          ],
          "internalType": "struct Order",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "order_status",
      "outputs": [
        {
          "internalType": "enum OrderStatus",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "pay_billing",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "provide_quotes",
      "outputs": [
        {
          "internalType": "address",
          "name": "provider",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "cpu_price",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_price",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_price",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "provider_factory",
      "outputs": [
        {
          "internalType": "contract IProviderFactory",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "query_provide_quotes",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "provider",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "cpu_price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_price",
              "type": "uint256"
            }
          ],
          "internalType": "struct PriceOracle[]",
          "name": "",
          "type": "tuple[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "query_provider_address",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "p_cpu",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "p_memory",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "p_storage",
          "type": "uint256"
        }
      ],
      "name": "quote",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "server_uri",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "uri",
          "type": "string"
        }
      ],
      "name": "submit_server_uri",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "uri_",
          "type": "string"
        }
      ],
      "name": "update_deployment",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "withdraw_fund",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }]`
const ProviderFactoryABI = `[{
      "inputs": [
        {
          "internalType": "address",
          "name": "_admin",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_order_factory",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_auditor_factory",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [],
      "name": "MIN_VALUE_TO_BE_PROVIDER",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "admin",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "auditor_factory",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "new_admin",
          "type": "address"
        }
      ],
      "name": "changeAdmin",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "new_order_factory",
          "type": "address"
        }
      ],
      "name": "changeOrderFactory",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "add",
          "type": "bool"
        }
      ],
      "name": "changeProviderResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "add",
          "type": "bool"
        }
      ],
      "name": "changeProviderUsedResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "closeProvider",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "name": "consumeResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        },
        {
          "internalType": "string",
          "name": "provider_info",
          "type": "string"
        }
      ],
      "name": "createNewProvider",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "getProvideContract",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "getProvideResource",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "getProvideTotalResource",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "start",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "limit",
          "type": "uint256"
        }
      ],
      "name": "getProviderInfo",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "provider",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "provider_owner",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "total_cpu",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "total_mem",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "total_sto",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "left_cpu",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "left_mem",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "left_sto",
              "type": "uint256"
            },
            {
              "internalType": "string",
              "name": "info",
              "type": "string"
            },
            {
              "internalType": "bool",
              "name": "is_active",
              "type": "bool"
            },
            {
              "internalType": "address[]",
              "name": "audits",
              "type": "address[]"
            }
          ],
          "internalType": "struct ProviderFactory.providerInfo[]",
          "name": "",
          "type": "tuple[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getProviderInfoLength",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "order_factory",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "provider_pledge",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "providers",
      "outputs": [
        {
          "internalType": "contract IProvider",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "reOpenProvider",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "name": "recoverResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_cpu",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_mem",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_storage",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_used_cpu",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_used_mem",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_used_storage",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }]`
const ProviderABI = `[{
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "_owner",
          "type": "address"
        },
        {
          "internalType": "string",
          "name": "provider_info",
          "type": "string"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [
        {
          "internalType": "bool",
          "name": "active",
          "type": "bool"
        }
      ],
      "name": "changeActive",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "new_info",
          "type": "string"
        }
      ],
      "name": "changeProviderInfo",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "consume_cpu",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "consume_mem",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "consume_storage",
          "type": "uint256"
        }
      ],
      "name": "consumeResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getLeftResource",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getTotalResource",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "info",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "isActive",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "consumed_cpu",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "consumed_mem",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "consumed_storage",
          "type": "uint256"
        }
      ],
      "name": "recoverResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_cpu",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_mem",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_storage",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "new_cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "new_mem_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "new_sto_count",
          "type": "uint256"
        }
      ],
      "name": "updateResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "used_cpu",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "used_mem",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "used_storage",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }]`

var (
	OrderFactoryName    = "order_factory"
	OrderBaseName       = "order_base"
	ProviderFactoryName = "provider_factory"
	ProviderName        = "provider"
	abiMap              = map[string]abi.ABI{}
)

func init() {
	abiMap = make(map[string]abi.ABI, 0)
	tmpABI, _ := abi.JSON(strings.NewReader(OrderFactoryABI))
	abiMap[OrderFactoryName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(OrderBaseABI))
	abiMap[OrderBaseName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(ProviderFactoryABI))
	abiMap[ProviderFactoryName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(ProviderABI))
	abiMap[ProviderName] = tmpABI
}
func GetInteractiveABI() map[string]abi.ABI {
	return abiMap
}

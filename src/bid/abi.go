package bid

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

const orderFactoryABI = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"orderNumber","type":"uint256"},{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"order_addr","type":"address"}],"name":"OrderCreation","type":"event"},{"inputs":[],"name":"cert_center","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"new_owner","type":"address"}],"name":"changeOwner","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"orderAddress","type":"address"}],"name":"checkIsOrder","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"m_cpu","type":"uint256"},{"internalType":"uint256","name":"m_memory","type":"uint256"},{"internalType":"uint256","name":"m_storage","type":"uint256"},{"internalType":"uint256","name":"m_cert","type":"uint256"},{"internalType":"uint256","name":"m_trx_id","type":"uint256"}],"name":"createOrder","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"orderId","type":"uint256"}],"name":"getOrder","outputs":[{"components":[{"internalType":"address","name":"contract_address","type":"address"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"v_cpu","type":"uint256"},{"internalType":"uint256","name":"v_memory","type":"uint256"},{"internalType":"uint256","name":"v_storage","type":"uint256"},{"internalType":"uint256","name":"cert_key","type":"uint256"},{"internalType":"uint256","name":"trx_id","type":"uint256"},{"internalType":"uint8","name":"state","type":"uint8"}],"internalType":"structOrder","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"providerAddress","type":"address"}],"name":"getProviderAllOrder","outputs":[{"components":[{"internalType":"address","name":"contract_address","type":"address"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"v_cpu","type":"uint256"},{"internalType":"uint256","name":"v_memory","type":"uint256"},{"internalType":"uint256","name":"v_storage","type":"uint256"},{"internalType":"uint256","name":"cert_key","type":"uint256"},{"internalType":"uint256","name":"trx_id","type":"uint256"},{"internalType":"uint8","name":"state","type":"uint8"}],"internalType":"structOrder[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getUnCompleteOrder","outputs":[{"components":[{"internalType":"address","name":"contract_address","type":"address"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"v_cpu","type":"uint256"},{"internalType":"uint256","name":"v_memory","type":"uint256"},{"internalType":"uint256","name":"v_storage","type":"uint256"},{"internalType":"uint256","name":"cert_key","type":"uint256"},{"internalType":"uint256","name":"trx_id","type":"uint256"},{"internalType":"uint8","name":"state","type":"uint8"}],"internalType":"structOrder[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"userAddress","type":"address"}],"name":"getUserAllOrder","outputs":[{"components":[{"internalType":"address","name":"contract_address","type":"address"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"v_cpu","type":"uint256"},{"internalType":"uint256","name":"v_memory","type":"uint256"},{"internalType":"uint256","name":"v_storage","type":"uint256"},{"internalType":"uint256","name":"cert_key","type":"uint256"},{"internalType":"uint256","name":"trx_id","type":"uint256"},{"internalType":"uint8","name":"state","type":"uint8"}],"internalType":"structOrder[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"get_minimum_deposit_amount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"max_order_index","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"minimum_deposit_amount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"new_value","type":"uint256"}],"name":"modify_minimum_deposit_amount","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"order_base_map","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"orders","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"provider_address","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"cert_center_","type":"address"}],"name":"set_cert_center","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"factory_addr","type":"address"}],"name":"set_provider_factory","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
const orderBaseABI = `[{"inputs":[{"internalType":"address","name":"_order_factory","type":"address"},{"internalType":"address","name":"provider_factory_","type":"address"},{"internalType":"address","name":"owner_","type":"address"},{"internalType":"uint256","name":"cpu_","type":"uint256"},{"internalType":"uint256","name":"memory_","type":"uint256"},{"internalType":"uint256","name":"storage_","type":"uint256"},{"internalType":"uint256","name":"cert_key_","type":"uint256"},{"internalType":"uint256","name":"sdl_trx_id_","type":"uint256"},{"internalType":"uint256","name":"order_number","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[],"name":"CanQuote","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"provider","type":"address"},{"indexed":true,"internalType":"uint256","name":"final_price","type":"uint256"}],"name":"ChooseQuote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"old_cpu","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"old_memory","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"old_storage","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"new_cpu","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"new_memory","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"new_storage","type":"uint256"}],"name":"DeployMentUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"DepositBalance","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"owner_","type":"address"},{"indexed":false,"internalType":"uint256","name":"cpu","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"memory_","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"storage_","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"cert","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"sdl","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"order_number","type":"uint256"}],"name":"OrderCreate","type":"event"},{"anonymous":false,"inputs":[],"name":"OrderEnded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"provider","type":"address"},{"indexed":true,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"PayBill","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"provider","type":"address"},{"indexed":false,"internalType":"uint256","name":"cpu_price","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"memory_price","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"storage_price","type":"uint256"}],"name":"Quote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"new_sdl_trx_id","type":"uint256"}],"name":"UpdateSDL","type":"event"},{"anonymous":false,"inputs":[],"name":"UserCancelOrder","type":"event"},{"inputs":[{"internalType":"uint256","name":"new_trx_hash","type":"uint256"}],"name":"change_sdl_trx_hash","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"quote_index","type":"uint256"}],"name":"choose_provider","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"deposit_balance","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"final_choice","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"final_price","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"last_pay_time","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_cert","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_cpu","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_memory","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_order_number","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_pending_sdl_trx_id","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_sdl_trx_id","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"o_storage","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"order_info","outputs":[{"components":[{"internalType":"address","name":"contract_address","type":"address"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"v_cpu","type":"uint256"},{"internalType":"uint256","name":"v_memory","type":"uint256"},{"internalType":"uint256","name":"v_storage","type":"uint256"},{"internalType":"uint256","name":"cert_key","type":"uint256"},{"internalType":"uint256","name":"trx_id","type":"uint256"},{"internalType":"uint8","name":"state","type":"uint8"}],"internalType":"structOrder","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"order_status","outputs":[{"internalType":"enumOrderStatus","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pay_billing","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"provide_quotes","outputs":[{"internalType":"address","name":"provider","type":"address"},{"internalType":"uint256","name":"cpu_price","type":"uint256"},{"internalType":"uint256","name":"memory_price","type":"uint256"},{"internalType":"uint256","name":"storage_price","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"provider_factory","outputs":[{"internalType":"contractIProviderFactory","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"query_provide_quotes","outputs":[{"components":[{"internalType":"address","name":"provider","type":"address"},{"internalType":"uint256","name":"cpu_price","type":"uint256"},{"internalType":"uint256","name":"memory_price","type":"uint256"},{"internalType":"uint256","name":"storage_price","type":"uint256"}],"internalType":"structPriceOracle[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"query_provider_address","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"p_cpu","type":"uint256"},{"internalType":"uint256","name":"p_memory","type":"uint256"},{"internalType":"uint256","name":"p_storage","type":"uint256"}],"name":"quote","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"server_uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"uri","type":"string"}],"name":"submit_server_uri","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"totalSpent","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"cpu_","type":"uint256"},{"internalType":"uint256","name":"memory_","type":"uint256"},{"internalType":"uint256","name":"storage_","type":"uint256"},{"internalType":"string","name":"uri_","type":"string"}],"name":"update_deployment","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"withdraw_fund","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

const providerFactoryABI = `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "ProviderCreate",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "addMargin",
      "outputs": [],
      "stateMutability": "payable",
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
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        }
      ],
      "name": "calcProviderAmount",
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
          "name": "new_audit_factory",
          "type": "address"
        }
      ],
      "name": "changeAuditorFactory",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "new_cpu_decimal",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "new_memory_decimal",
          "type": "uint256"
        }
      ],
      "name": "changeDecimal",
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
          "name": "_new_min",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_new_max",
          "type": "uint256"
        }
      ],
      "name": "changeProviderLimit",
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
          "internalType": "address",
          "name": "provider_owner",
          "type": "address"
        },
        {
          "internalType": "bool",
          "name": "whether_start",
          "type": "bool"
        }
      ],
      "name": "changeProviderState",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "_punish_address",
          "type": "address"
        }
      ],
      "name": "changePunishAddress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_new_punish_start_limit",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_new_punish_interval",
          "type": "uint256"
        }
      ],
      "name": "changePunishParam",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_new_punish_percent",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_new_punish_all_percent",
          "type": "uint256"
        }
      ],
      "name": "changePunishPercent",
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
          "name": "region",
          "type": "string"
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
      "inputs": [],
      "name": "decimal_cpu",
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
      "name": "decimal_memory",
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
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
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
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
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
              "name": "provider_contract",
              "type": "address"
            },
            {
              "components": [
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "total",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "used",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "lock",
                  "type": "tuple"
                },
                {
                  "internalType": "bool",
                  "name": "challenge",
                  "type": "bool"
                },
                {
                  "internalType": "enum ProviderState",
                  "name": "state",
                  "type": "uint8"
                },
                {
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "internalType": "string",
                  "name": "region",
                  "type": "string"
                },
                {
                  "internalType": "string",
                  "name": "info",
                  "type": "string"
                },
                {
                  "internalType": "uint256",
                  "name": "last_challenge_time",
                  "type": "uint256"
                }
              ],
              "internalType": "struct providerInfo",
              "name": "info",
              "type": "tuple"
            },
            {
              "internalType": "uint256",
              "name": "margin_amount",
              "type": "uint256"
            },
            {
              "internalType": "address[]",
              "name": "audits",
              "type": "address[]"
            }
          ],
          "internalType": "struct ProviderFactory.providerInfos[]",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "_provider_contract",
          "type": "address"
        }
      ],
      "name": "getProviderSingle",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "provider_contract",
              "type": "address"
            },
            {
              "components": [
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "total",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "used",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "uint256",
                      "name": "cpu_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "memory_count",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint256",
                      "name": "storage_count",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct poaResource",
                  "name": "lock",
                  "type": "tuple"
                },
                {
                  "internalType": "bool",
                  "name": "challenge",
                  "type": "bool"
                },
                {
                  "internalType": "enum ProviderState",
                  "name": "state",
                  "type": "uint8"
                },
                {
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "internalType": "string",
                  "name": "region",
                  "type": "string"
                },
                {
                  "internalType": "string",
                  "name": "info",
                  "type": "string"
                },
                {
                  "internalType": "uint256",
                  "name": "last_challenge_time",
                  "type": "uint256"
                }
              ],
              "internalType": "struct providerInfo",
              "name": "info",
              "type": "tuple"
            },
            {
              "internalType": "uint256",
              "name": "margin_amount",
              "type": "uint256"
            },
            {
              "internalType": "address[]",
              "name": "audits",
              "type": "address[]"
            }
          ],
          "internalType": "struct ProviderFactory.providerInfos",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "punish_amount",
          "type": "uint256"
        }
      ],
      "name": "getPunishAmount",
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
      "name": "getTotalDetail",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_admin",
          "type": "address"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialized",
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
      "name": "max_value_tobe_provider",
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
      "name": "min_value_tobe_provider",
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
      "name": "punish_address",
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
      "name": "punish_all_percent",
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
      "name": "punish_interval",
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
      "name": "punish_percent",
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
      "name": "punish_start_limit",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "provider",
          "type": "address"
        }
      ],
      "name": "removePunishList",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_all",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total_used",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
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
          "name": "new_provider",
          "type": "address"
        }
      ],
      "name": "tryPunish",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "provider_owner",
          "type": "address"
        }
      ],
      "name": "whetherCanPOR",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`
const providerABI = ` [
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
          "internalType": "address",
          "name": "_owner",
          "type": "address"
        },
        {
          "internalType": "string",
          "name": "_region",
          "type": "string"
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
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "ProviderResourceChange",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "challenge",
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
      "name": "challengeProvider",
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
          "internalType": "string",
          "name": "_new_region",
          "type": "string"
        }
      ],
      "name": "changeRegion",
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
      "name": "getDetail",
      "outputs": [
        {
          "components": [
            {
              "components": [
                {
                  "internalType": "uint256",
                  "name": "cpu_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "memory_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "storage_count",
                  "type": "uint256"
                }
              ],
              "internalType": "struct poaResource",
              "name": "total",
              "type": "tuple"
            },
            {
              "components": [
                {
                  "internalType": "uint256",
                  "name": "cpu_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "memory_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "storage_count",
                  "type": "uint256"
                }
              ],
              "internalType": "struct poaResource",
              "name": "used",
              "type": "tuple"
            },
            {
              "components": [
                {
                  "internalType": "uint256",
                  "name": "cpu_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "memory_count",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "storage_count",
                  "type": "uint256"
                }
              ],
              "internalType": "struct poaResource",
              "name": "lock",
              "type": "tuple"
            },
            {
              "internalType": "bool",
              "name": "challenge",
              "type": "bool"
            },
            {
              "internalType": "enum ProviderState",
              "name": "state",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "owner",
              "type": "address"
            },
            {
              "internalType": "string",
              "name": "region",
              "type": "string"
            },
            {
              "internalType": "string",
              "name": "info",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "last_challenge_time",
              "type": "uint256"
            }
          ],
          "internalType": "struct providerInfo",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getLeftResource",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
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
          "components": [
            {
              "internalType": "uint256",
              "name": "cpu_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "memory_count",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "storage_count",
              "type": "uint256"
            }
          ],
          "internalType": "struct poaResource",
          "name": "",
          "type": "tuple"
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
      "name": "last_challenge_time",
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
      "name": "last_margin_time",
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
      "name": "last_punish_time",
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
      "name": "lock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "margin_block",
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
      "name": "provider_first_margin_time",
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
      "name": "punish",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "punish_start_margin_amount",
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
      "name": "punish_start_time",
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
      "inputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "name": "reduceResource",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "region",
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
      "name": "removePunish",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bool",
          "name": "whether_start",
          "type": "bool"
        }
      ],
      "name": "startChallenge",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "state",
      "outputs": [
        {
          "internalType": "enum ProviderState",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "total",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "triggerMargin",
      "outputs": [],
      "stateMutability": "nonpayable",
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
      "name": "used",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "cpu_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "memory_count",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "storage_count",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "withdrawMargin",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "stateMutability": "payable",
      "type": "receive"
    }
  ]`

var (
	//OrderFactoryName is abi map name
	OrderFactoryName = "order_factory"
	//OrderBaseName is abi map name
	OrderBaseName = "order_base"
	//ProviderFactoryName is abi map name
	ProviderFactoryName = "provider_factory"
	//ProviderName is abi map name
	ProviderName = "provider"
	abiMap       = map[string]abi.ABI{}
)

func init() {
	abiMap = make(map[string]abi.ABI, 0)
	tmpABI, _ := abi.JSON(strings.NewReader(orderFactoryABI))
	abiMap[OrderFactoryName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(orderBaseABI))
	abiMap[OrderBaseName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(providerFactoryABI))
	abiMap[ProviderFactoryName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(providerABI))
	abiMap[ProviderName] = tmpABI
}

//GetInteractiveABI is interface get all abi
func GetInteractiveABI() map[string]abi.ABI {
	return abiMap
}

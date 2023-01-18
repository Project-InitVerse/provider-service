package util

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"runtime"
	"strings"
)

type BlockData struct {
	Difficulty               string        `json:"difficulty"`
	ExtraData                string        `json:"extraData"`
	GasLimit                 string        `json:"gasLimit"`
	GasUsed                  string        `json:"gasUsed"`
	Hash                     string        `json:"hash"`
	LogsBloom                string        `json:"logsBloom"`
	Miner                    string        `json:"miner"`
	MixHash                  string        `json:"mixHash"`
	Nonce                    string        `json:"nonce"`
	Number                   string        `json:"number"`
	ParentHash               string        `json:"parentHash"`
	ReceiptsRoot             string        `json:"receiptsRoot"`
	Sha3Uncles               string        `json:"sha3Uncles"`
	Size                     string        `json:"size"`
	StateRoot                string        `json:"stateRoot"`
	Timestamp                string        `json:"timestamp"`
	TotalDifficulty          string        `json:"totalDifficulty"`
	Transactions             []interface{} `json:"transactions"`
	TransactionsRoot         string        `json:"transactionsRoot"`
	Uncles                   []interface{} `json:"uncles"`
	BlockReward              string        `json:"blockReward"`
	InternalTransactionCount int           `json:"internalTransactionCount"`
	L1Status                 int64         `json:"l1Status"`
	BaseFee                  string        `json:"baseFeePerGas"`
}

type BlockDataHeader struct {
	ExtraData        string `json:"-"`
	Difficulty       int64  `json:"-"`
	GasLimit         int64  `json:"gasLimit"`
	GasUsed          int64  `json:"gasUsed"`
	Hash             string `json:"blockHash"`
	TransactionCount int    `json:"transactionCount"`
	Validator        string `json:"validator"`
	Nonce            string `json:"-"`
	Number           int64  `json:"blockNumber"`
	ParentHash       string `json:"-"`

	Sha3Uncles string `json:"-"`
	Size       int64  `json:"-"`

	Timestamp                 int64  `json:"blockTime"`
	TotalDifficulty           int64  `json:"totalDifficulty"`
	Uncles                    string `json:"-"`
	BlockReward               string `json:"blockReward"`
	InternalTransactionCount  int    `json:"internalTransactionCount"`
	L1Status                  int64  `json:"l1Status"`
	L1CommitBlockHash         string `json:"-"`
	L1CommitBlockNumber       string `json:"-"`
	L1CommitTransactionHash   string `json:"l1CommitTransactionHash"`
	L1FinalizeBlockHash       string `json:"-"`
	L1FinalizeBlockNumber     string `json:"-"`
	L1FinalizeTransactionHash string `json:"l1FinalizeTransactionHash"`
}
type Bloom [256]byte
type BlockNonce [8]byte

func (n *BlockNonce) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BlockNonce", input, n[:])
}

type ChainHeader struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"            gencodec:"required"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`
}

type InternalTransaction struct {
	BlockHash             string `json:"blockHash"`
	BlockNumber           string `json:"blockNumber"`
	ParentTransactionHash string `json:"parentTransactionHash"`
	From                  string `json:"from"`
	To                    string `json:"to"`
	Value                 string `json:"value"`
	TypeTraceAddress      string `json:"typeTraceAddress"`
	GasLimit              string `json:"gasLimit"`
	Op                    string `json:"op"`
}

type TransactionLogs struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type CrossTransactionLogs struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	EventName        string   `json:"event_name"`
	EventType        int      `json:"event_type"`
	Status           int      `json:"status"`
}

type TrxReceiptData struct {
	BlockHash         string            `json:"blockHash"`
	BlockNumber       string            `json:"blockNumber"`
	ContractAddress   string            `json:"contractAddress"`
	CumulativeGasUsed string            `json:"cumulativeGasUsed"`
	EffectiveGasPrice string            `json:"effectiveGasPrice"`
	From              string            `json:"from"`
	GasUsed           string            `json:"gasUsed"`
	Logs              []TransactionLogs `json:"logs"`
	LogsBloom         string            `json:"logsBloom"`
	Status            string            `json:"status"`
	To                string            `json:"to"`
	TransactionHash   string            `json:"transactionHash"`
	TransactionIndex  string            `json:"transactionIndex"`
	Type              string            `json:"type"`
}

type TokenTransaction struct {
	TransactionHash string `json:"transactionHash"`
	LogIndex        string `json:"logIndex"`
	Contract        string `json:"contract"`
	TokenType       int    `json:"tokenType"`
	Value           string `json:"value"`
	TokenId         string `json:"tokenId"`
	From            string `json:"from"`
	To              string `json:"to"`
	MethodId        string `json:"methodId"`
	BlockHash       string `json:"blockHash"`
	BlockTime       int64  `json:"blockTime"`
}

type ContractObject struct {
	ContractAddress string `json:"contractAddress"`
	CreateTxHash    string `json:"createTxHash"`
	Creator         string `json:"creator"`
	ByteCode        string `json:"byteCode"`
}
type UpdateContractObject struct {
	ContractAddress string `json:"contractAddress"`
	ContractType    int    `json:"contractType"`
	Symbol          string `json:"symbol"`
	Decimals        int    `json:"decimals"`
	TotalSupply     string `json:"totalSupply"`
	Name            string `json:"name"`
}
type ContractNeedInit struct {
	Contract  string `json:"contract"`
	TokenType int    `json:"tokenType"`
}

type BalanceChange struct {
	Address      string   `json:"address"`
	Contract     string   `json:"contract"`
	ContractType int      `json:"contractType"`
	TokenId      string   `json:"tokenId"`
	Value        *big.Int `_`
}
type VoteDetail struct {
	Address         string   `json:"address"`
	Contract        string   `json:"address"`
	Value           *big.Int `_`
	LogIndex        string   `json:"logIndex"`
	TransactionHash string   `json:"transactionHash"`
}
type ValidatorState struct {
	Address  string `json:"address"`
	Contract string `json:"address"`
	State    int    `json:"state"`
	Active   int    `json:"active"`
}

func GetAddress(addr string) string {
	return strings.ToLower(common.HexToAddress(addr).String())
}

func GetBigInt(int_str string, base int) *big.Int {
	if int_str == "" {
		return big.NewInt(0)
	}
	if int_str == "0x" {
		return big.NewInt(0)
	}
	if strings.HasPrefix(int_str, "0x") {
		int_str = int_str[2:]
	}
	big_int, _ := big.NewInt(0).SetString(int_str, base)

	return big_int
}
func GetBigIntString(int_str string, base int) string {
	if strings.HasPrefix(int_str, "0x") {
		int_str = int_str[2:]
	}
	if int_str == "" {
		return "0"
	}
	big_int, _ := big.NewInt(0).SetString(int_str, base)
	//fmt.Println("GetBigIntString",int_str,big_int.String())
	return big_int.String()
}

func GetAbiStrings(origins string) map[string]string {
	res := make(map[string]string, 0)
	len := GetBigInt(origins[66:130], 16).Int64()
	for i := 0; i < int(len); i++ {
		one_addr := "0x" + origins[130+i*64+24:130+(1+i)*64]
		res[one_addr] = ""
	}
	return res
}

type Transaction struct {
	Hash             string `json:"hash"`
	Status           string `json:"status"`
	ErrorInfo        string `json:"errorInfo"`
	BlockNumber      string `json:"blockNumber"`
	BlockTime        int64  `json:"blockTime"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Fee              string `json:"fee"`
	GasLimit         string `json:"gasLimit"`
	GasPrice         string `json:"gasPrice"`
	TransactionType  string `json:"transactionType"`
	MaxPriority      string `json:"maxPriority"`
	MaxFee           string `json:"maxFee"`
	Nonce            string `json:"nonce"`
	InputData        string `json:"inputData"`
	BlockHash        string `json:"blockHash"`
	GasUsed          string `json:"gasUsed"`
	TransactionIndex string `json:"transactionIndex"`
	L1Status         string `json:"l1Status"`
}
type NeedBid struct {
	ContractAddress string   `json:"contract_address"`
	Cpu             *big.Int `json:"cpu"`
	Memory          *big.Int `json:"memory"`
	Storage         *big.Int `json:"storage"`
	Cert            string   `json:"cert"`
	SdlTrxId        string   `json:"sdl_trx_id"`
	State           uint8    `json:"state"`
}
type NeedCreate struct {
	Provider        common.Address `json:"provider"`
	ContractAddress string         `json:"contract_address"`
	CpuPrice        *big.Int       `json:"cpu_price"`
	MemoryPrice     *big.Int       `json:"memory_price"`
	StoragePrice    *big.Int       `json:"storage_price"`
	FinalPrice      *big.Int       `json:"final_price"`
}
type UserCancelOrder struct {
	ContractAddress common.Address `json:"contract_address"`
}

type TransactionHandled struct {
	Hash             string `json:"hash"`
	Status           int64  `json:"status"`
	ErrorInfo        string `json:"errorInfo"`
	BlockNumber      int64  `json:"blockNumber"`
	BlockTime        int64  `json:"blockTime"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Fee              string `json:"fee"`
	GasLimit         int64  `json:"gasLimit"`
	GasPrice         string `json:"gasPrice"`
	TransactionType  int64  `json:"transactionType"`
	MaxPriority      int64  `json:"maxPriority"`
	MaxFee           int64  `json:"maxFee"`
	Nonce            int64  `json:"nonce"`
	InputData        string `json:"inputData"`
	BlockHash        string `json:"blockHash"`
	GasUsed          int64  `json:"gasUsed"`
	TransactionIndex int64  `json:"transactionIndex"`
	L1Status         int64  `json:"l1Status"`
}

func PrintMemStats() {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

}

func ConvertLogToGethLogs(logs TransactionLogs) types.Log {
	bdata, _ := json.Marshal(logs)

	data := new(types.Log)
	err := json.Unmarshal(bdata, data)
	if err != nil {
		fmt.Println("ConvertLogToGethLogs failed", err)
		return types.Log{}
	}
	return *data
}

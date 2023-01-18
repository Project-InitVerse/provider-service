package scan

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"providerService/src/config"
	"providerService/src/util"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Scan struct {
	linkClient      *util.LinkClient
	height          uint
	MaintainChan    chan util.Order
	NeedBidChan     chan util.NeedBid
	NeedCreateChan  chan util.NeedCreate
	UserCancelChan  chan util.UserCancelOrder
	collect_threads int
	handle_threads  int
	wg              sync.WaitGroup
	exit_count      int32
	allEvent        util.AllEvent
	Config          *config.ProviderConfig
	is_done         bool
}

func (self *Scan) QueryExsitMaintainOrder() {
	//todo
	for {
		conn, err := ethclient.Dial(self.Config.NodeUrl)
		if err != nil {
			log.Printf("Failed to connect to the Ethereum client %v", err)
		} else {
			defer conn.Close()
			orderBase, err := util.NewOrderFactory(common.HexToAddress(util.GetOrderFactory(self.Config)), conn)
			if err != nil {
				log.Printf("call orderBase failed")
				continue
			}
			allProviderOrder, err := orderBase.GetProviderAllOrder(nil, common.HexToAddress(self.Config.ProviderContract))
			if err != nil {
				log.Printf("call orderBase failed")
				continue
			}
			for _, one_order := range allProviderOrder {
				if one_order.State == 2 {
					self.MaintainChan <- one_order
				}
			}

			break
		}
		time.Sleep(5 * time.Second)
	}
}

func (self *Scan) MainLoop(ctx context.Context, link_client *util.LinkClient, global_wg *sync.WaitGroup) {
	global_wg.Add(1)
	defer global_wg.Done()
	blockdata_chan := make(chan simplejson.Json, 40)
	height_chan := make(chan int, 40)
	self.wg.Add(self.handle_threads + self.collect_threads)
	for i := 0; i < self.collect_threads; i++ {
		go self.collect_block(link_client, height_chan, blockdata_chan)
	}

	for i := 0; i < self.handle_threads; i++ {
		go self.handle_block_new(link_client, blockdata_chan, 1000)
	}
	old_count := 0

	go func() {
		select {
		case <-ctx.Done():
			self.quit()
			break
		}
	}()
	for !self.is_done {

		param := make([]interface{}, 0)
		json_data := link_client.SafeLinkHttpFunc("eth_blockNumber", &param)
		res_count, _ := json_data.Get("result").String()

		count64, _ := strconv.ParseInt(res_count[2:], 16, 32)
		count := int(count64) - 2
		if old_count == 0 {
			old_count = count - 1
		}
		fmt.Println(old_count, count)
		if old_count < count {
			//fmt.Println("current height:", old_count, "target height", count, "time", time.Now())
			for i := old_count; i < count; i++ {
				height_chan <- i
			}
			old_count = count
		} else {
			time.Sleep(5 * time.Second)
		}

	}
	for i := 0; i < self.collect_threads; i++ {
		height_chan <- -1
	}
	self.wg.Wait()
}
func (self *Scan) quit() {
	self.is_done = true

}

func (self *Scan) InitScan(pConfig *config.ProviderConfig) {
	self.allEvent = util.AllEvent{}
	self.allEvent.Init()
	self.Config = pConfig
	self.NeedBidChan = make(chan util.NeedBid, 20)
	self.UserCancelChan = make(chan util.UserCancelOrder, 20)
	self.NeedCreateChan = make(chan util.NeedCreate, 20)
	self.collect_threads = 1
	self.handle_threads = 1
	self.is_done = false
}

func (scan *Scan) collect_block(link_client *util.LinkClient, height_chan chan int, blockdata_chan chan simplejson.Json) {
	defer scan.wg.Done()

	//vin := &pro.TrxObject_VIN{}
	//trx_object := pro.TrxObject{}
	//back_time := time.Now()

	for {
		param := make([]interface{}, 0, 2)
		once_height := <-height_chan
		if once_height == -1 {
			fmt.Println("collect exit")
			atomic.AddInt32(&scan.exit_count, int32(1))
			json_data, _ := simplejson.NewJson([]byte("{\"result\":\"exit\"}"))
			blockdata_chan <- *json_data

			return
		}
		param = append(param, "0x"+strconv.FormatInt(int64(once_height), 16))
		param = append(param, "true")
		blockdata := link_client.SafeLinkHttpFunc("eth_getBlockByNumber", &param)
		if once_height%1000 == 0 {
			fmt.Println("height", once_height, "chan size", len(blockdata_chan))
		}
		blockdata_chan <- *blockdata
		Hash, _ := blockdata.Get("result").Get("hash").String()
		log.Printf("number:%d hash:%s", once_height, Hash)

	}
}

func (self *Scan) handle_block_new(link_client *util.LinkClient, blockdata_chan chan simplejson.Json, interval int64) {
	defer self.wg.Done()

	for {
		blockchain_data := <-blockdata_chan

		exit_code, err := blockchain_data.Get("result").String()
		if err == nil {

			if exit_code == "exit" {
				for self.exit_count < int32(self.collect_threads) {
					time.Sleep(1 * time.Second)
				}
				if self.exit_count < int32(self.handle_threads) {
					atomic.AddInt32(&self.exit_count, int32(1))
					json_data, _ := simplejson.NewJson([]byte("{\"result\":\"exit\"}"))
					blockdata_chan <- *json_data
				}
				//flush_db_nosync(&trx_cache,&erc20_address_trx_cache)

				fmt.Println("check exit")

				return
			}
		}
		bdata, _ := blockchain_data.Get("result").MarshalJSON()

		var blockData util.BlockData
		err = json.Unmarshal(bdata, &blockData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		blockTime := get_int(blockData.Timestamp)
		blockReward := big.NewInt(0)
		for _, trx_data := range blockData.Transactions {
			blockReward = blockReward.Add(blockReward, self.handle_transaction(link_client, &trx_data, blockTime))
		}

		tmp_height := get_int(blockData.Number)
		if tmp_height%interval == 0 {
			runtime.GC()
		}

	}

}

func get_int(int_str string) int64 {
	if strings.HasPrefix(int_str, "0x") {
		value, _ := strconv.ParseInt(int_str[2:], 16, 32)
		return value
	} else {
		value, _ := strconv.ParseInt(int_str, 10, 32)
		return value
	}
}

type TrxData struct {
	AccessList           []interface{} `json:"accessList"`
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	ChainId              string        `json:"chainId"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Nonce                string        `json:"nonce"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	To                   string        `json:"to"`
	TransactionIndex     string        `json:"transactionIndex"`
	Type                 string        `json:"type"`
	V                    string        `json:"v"`
	Value                string        `json:"value"`
}

func (self *Scan) handle_transaction(link_client *util.LinkClient, trx_data *interface{}, blockTime int64) *big.Int {

	trx_map_obj := make(map[string]interface{})
	reward := big.NewInt(0)
	var trx_simple_data TrxData
	bdata, _ := json.Marshal(*trx_data)
	err := json.Unmarshal(bdata, &trx_simple_data)
	if err != nil {
		fmt.Println("transaction Decode Failed!", err)
		return reward
	}
	param := make([]interface{}, 0, 20)
	param = append(param, trx_simple_data.Hash)
	trxReceiptData, _ := link_client.SafeLinkHttpFunc("eth_getTransactionReceipt", &param).Get("result").MarshalJSON()
	var trxReceipt util.TrxReceiptData
	err = json.Unmarshal(trxReceiptData, &trxReceipt)
	if err != nil {
		fmt.Println(err)
		return reward
	}
	trx_map_obj["hash"] = trx_simple_data.Hash
	trx_map_obj["blockNumber"] = trx_simple_data.BlockNumber
	trx_map_obj["blockHash"] = trx_simple_data.BlockHash
	trx_map_obj["from"] = trx_simple_data.From
	trx_map_obj["to"] = trx_simple_data.To
	trx_map_obj["gasLimit"] = trx_simple_data.Gas
	gasUsed := util.GetBigInt(trxReceipt.GasUsed, 16)
	if gasUsed == nil {
		gasUsed = big.NewInt(0)
	}
	if trx_simple_data.Type == "0x0" {
		trx_map_obj["gasPrice"] = trx_simple_data.GasPrice
		trx_map_obj["maxPriority"] = ""
		trx_map_obj["maxFee"] = ""
		if trxReceipt.EffectiveGasPrice != "" {
			efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
			gasP := big.NewInt(0).Sub(util.GetBigInt(trx_simple_data.GasPrice, 16), efGas)
			reward = reward.Mul(gasUsed, gasP)
		} else {

			reward = reward.Mul(gasUsed, util.GetBigInt(trx_simple_data.GasPrice, 16))
		}

		trx_map_obj["fee"] = big.NewInt(0).Mul(util.GetBigInt(trx_simple_data.GasPrice, 16), gasUsed).String()
	} else if trx_simple_data.Type == "0x2" {
		trx_map_obj["maxPriority"] = trx_simple_data.MaxPriorityFeePerGas
		trx_map_obj["maxFee"] = trx_simple_data.MaxFeePerGas
		trx_map_obj["gasPrice"] = trxReceipt.EffectiveGasPrice
		efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
		maxF := util.GetBigInt(trx_simple_data.MaxFeePerGas, 16)

		reward = reward.Mul(gasUsed, maxF)
		trx_map_obj["fee"] = big.NewInt(0).Mul(gasUsed, efGas).String()
	} else {
		trx_map_obj["gasPrice"] = trx_simple_data.GasPrice
		trx_map_obj["maxPriority"] = ""
		trx_map_obj["maxFee"] = ""
		if trxReceipt.EffectiveGasPrice != "" {
			efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
			gasP := big.NewInt(0).Sub(util.GetBigInt(trx_simple_data.GasPrice, 16), efGas)
			reward = reward.Mul(gasUsed, gasP)
		} else {

			reward = reward.Mul(gasUsed, util.GetBigInt(trx_simple_data.GasPrice, 16))
		}

		trx_map_obj["fee"] = big.NewInt(0).Mul(util.GetBigInt(trx_simple_data.GasPrice, 16), gasUsed).String()
	}
	trx_map_obj["to"] = trx_simple_data.To
	trx_map_obj["blockTime"] = blockTime
	trx_map_obj["transactionType"] = trx_simple_data.Type
	trx_map_obj["value"] = trx_simple_data.Value
	trx_map_obj["inputData"] = trx_simple_data.Input
	trx_map_obj["gasUsed"] = trxReceipt.GasUsed
	trx_map_obj["status"] = trxReceipt.Status
	trx_map_obj["nonce"] = trx_simple_data.Nonce
	trx_map_obj["transactionIndex"] = trx_simple_data.TransactionIndex

	if trxReceipt.Status == "0x1" {
		for _, one_log := range trxReceipt.Logs {
			// CanQuote()

			if len(one_log.Topics) > 0 {
				if self.allEvent.CheckEqual(self.allEvent.CanQuote, one_log.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					contract_addr := one_log.Address
					call_data := util.GenerateOrderCall(contract_addr)
					param := make([]interface{}, 0, 20)
					param = append(param, trx_simple_data.Hash)
					call_res := link_client.EthCall(util.GetOrderFactory(self.Config), call_data)
					orderId := util.GetBigInt(call_res, 16)

					if orderId.Int64() > 0 {
						var order_base util.Order
						for {
							conn, err := ethclient.Dial(self.Config.NodeUrl)
							if err != nil {
								log.Printf("Failed to connect to the Ethereum client %v", err)
							} else {
								defer conn.Close()
								orderBase, err := util.NewOrderBase(common.HexToAddress(contract_addr), conn)
								if err != nil {
									log.Printf("call orderBase failed")
									continue
								}
								order_base, err = orderBase.OrderInfo(nil)
								if err != nil {
									log.Printf("call orderBase failed")
									continue
								}
								break
							}
							time.Sleep(5 * time.Second)
						}
						NeedBidObj := new(util.NeedBid)
						NeedBidObj.Cpu = order_base.VCpu
						NeedBidObj.Cert = order_base.CertKey
						NeedBidObj.Memory = order_base.VMemory
						NeedBidObj.Storage = order_base.VStorage
						NeedBidObj.SdlTrxId = order_base.TrxId.Text(16)
						NeedBidObj.ContractAddress = contract_addr
						NeedBidObj.State = order_base.State
						self.NeedBidChan <- *NeedBidObj
					}

				} else if self.allEvent.CheckEqual(self.allEvent.ChooseQuote, one_log.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					contract_addr := one_log.Address
					call_data := util.GenerateOrderCall(contract_addr)
					param := make([]interface{}, 0, 20)
					param = append(param, trx_simple_data.Hash)
					call_res := link_client.EthCall(util.GetOrderFactory(self.Config), call_data)
					orderId := util.GetBigInt(call_res, 16)

					if orderId.Int64() > 0 {
						OrderBaseFilter, _ := util.NewOrderBaseFilterer(common.HexToAddress(contract_addr), nil)

						CanQuoteEvent, err := OrderBaseFilter.ParseChooseQuote(util.ConvertLogToGethLogs(one_log))
						if err == nil {
							needCreate := new(util.NeedCreate)
							needCreate.CpuPrice = CanQuoteEvent.Price.CpuPrice
							needCreate.MemoryPrice = CanQuoteEvent.Price.MemoryPrice
							needCreate.StoragePrice = CanQuoteEvent.Price.StoragePrice
							needCreate.Provider = CanQuoteEvent.Price.Provider
							needCreate.FinalPrice = CanQuoteEvent.FinalPrice
							needCreate.ContractAddress = contract_addr
							self.NeedCreateChan <- *needCreate
						}
					}

				} else if self.allEvent.CheckEqual(self.allEvent.UserCancelOrder, one_log.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					contract_addr := one_log.Address
					call_data := util.GenerateOrderCall(contract_addr)
					param := make([]interface{}, 0, 20)
					param = append(param, trx_simple_data.Hash)
					call_res := link_client.EthCall(util.GetOrderFactory(self.Config), call_data)
					orderId := util.GetBigInt(call_res, 16)

					if orderId.Int64() > 0 {
						userCancel := new(util.UserCancelOrder)
						userCancel.ContractAddress = common.HexToAddress(contract_addr)
						self.UserCancelChan <- *userCancel
					}

				}

			}

		}
		//trace log

		trace_param := make([]interface{}, 0, 2)
		trace_param = append(trace_param, trx_simple_data.Hash)

	}

	return reward

	//from

}

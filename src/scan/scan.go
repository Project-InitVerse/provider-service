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

// Scan is struct
type Scan struct {
	linkClient     *util.LinkClient
	height         uint
	MaintainChan   chan util.Order
	NeedBidChan    chan util.NeedBid
	NeedCreateChan chan util.NeedCreate
	UserCancelChan chan util.UserCancelOrder
	collectThreads int
	handleThreads  int
	wg             sync.WaitGroup
	exitCount      int32
	allEvent       util.AllEvent
	Config         *config.ProviderConfig
	isDone         bool
}

//QueryExsitMaintainOrder is function
func (sc *Scan) QueryExsitMaintainOrder() {
	//todo
	for {
		conn, err := ethclient.Dial(sc.Config.NodeURL)
		if err != nil {
			log.Printf("Failed to connect to the Ethereum client %v", err)
		} else {
			defer conn.Close()
			orderBase, err := util.NewOrderFactory(common.HexToAddress(util.GetOrderFactory(sc.Config)), conn)
			if err != nil {
				log.Printf("call orderBase failed")
				continue
			}
			allProviderOrder, err := orderBase.GetProviderAllOrder(nil, common.HexToAddress(sc.Config.ProviderContract))
			if err != nil {
				log.Printf("call orderBase failed")
				continue
			}
			for _, oneOrder := range allProviderOrder {
				if oneOrder.State == 2 {
					sc.MaintainChan <- oneOrder
				}
			}

			break
		}
		time.Sleep(5 * time.Second)
	}
}

//MainLoop is service get info from chain
func (sc *Scan) MainLoop(ctx context.Context, linkClient *util.LinkClient, globalWg *sync.WaitGroup) {
	globalWg.Add(1)
	defer globalWg.Done()
	blockDataChan := make(chan simplejson.Json, 40)
	heightChan := make(chan int, 40)
	sc.wg.Add(sc.handleThreads + sc.collectThreads)
	for i := 0; i < sc.collectThreads; i++ {
		go sc.collectBlock(linkClient, heightChan, blockDataChan)
	}

	for i := 0; i < sc.handleThreads; i++ {
		go sc.handleBlockNew(linkClient, blockDataChan, 1000)
	}
	oldCount := 0

	go func() {
		select {
		case <-ctx.Done():
			sc.quit()
			break
		}
	}()
	for !sc.isDone {

		param := make([]interface{}, 0)
		jsonData := linkClient.SafeLinkHTTPFunc("eth_blockNumber", &param)
		resCount, _ := jsonData.Get("result").String()

		count64, _ := strconv.ParseInt(resCount[2:], 16, 32)
		count := int(count64) - 2
		//TODO for test
		if oldCount == 0 {
			oldCount = count - 1
		}
		if oldCount < count {
			//fmt.Println("current height:", old_count, "target height", count, "time", time.Now())
			for i := oldCount; i < count; i++ {
				heightChan <- i
			}
			oldCount = count
		} else {
			time.Sleep(5 * time.Second)
		}

	}
	for i := 0; i < sc.collectThreads; i++ {
		heightChan <- -1
	}
	sc.wg.Wait()
}
func (sc *Scan) quit() {
	sc.isDone = true

}

//InitScan init scan service
func (sc *Scan) InitScan(pConfig *config.ProviderConfig) {
	sc.allEvent = util.AllEvent{}
	sc.allEvent.Init()
	sc.Config = pConfig
	sc.NeedBidChan = make(chan util.NeedBid, 20)
	sc.UserCancelChan = make(chan util.UserCancelOrder, 20)
	sc.NeedCreateChan = make(chan util.NeedCreate, 20)
	sc.collectThreads = 1
	sc.handleThreads = 1
	sc.isDone = false
}

func (sc *Scan) collectBlock(linkClient *util.LinkClient, heightChan chan int, blockDataChan chan simplejson.Json) {
	defer sc.wg.Done()

	//vin := &pro.TrxObject_VIN{}
	//trx_object := pro.TrxObject{}
	//back_time := time.Now()

	for {
		param := make([]interface{}, 0, 2)
		onceHeight := <-heightChan
		if onceHeight == -1 {
			fmt.Println("collect exit")
			atomic.AddInt32(&sc.exitCount, int32(1))
			jsonData, _ := simplejson.NewJson([]byte("{\"result\":\"exit\"}"))
			blockDataChan <- *jsonData

			return
		}
		param = append(param, "0x"+strconv.FormatInt(int64(onceHeight), 16))
		param = append(param, "true")
		blockdata := linkClient.SafeLinkHTTPFunc("eth_getBlockByNumber", &param)
		if onceHeight%1000 == 0 {
			fmt.Println("height", onceHeight, "chan size", len(blockDataChan))
		}
		blockDataChan <- *blockdata
		//Hash, _ := blockdata.Get("result").Get("hash").String()
		//log.Printf("number:%d hash:%s", once_height, Hash)

	}
}

func (sc *Scan) handleBlockNew(linkClient *util.LinkClient, blockdataChan chan simplejson.Json, interval int64) {
	defer sc.wg.Done()

	for {
		blockchainData := <-blockdataChan

		exitCode, err := blockchainData.Get("result").String()
		if err == nil {

			if exitCode == "exit" {
				for sc.exitCount < int32(sc.collectThreads) {
					time.Sleep(1 * time.Second)
				}
				if sc.exitCount < int32(sc.handleThreads) {
					atomic.AddInt32(&sc.exitCount, int32(1))
					jsonData, _ := simplejson.NewJson([]byte("{\"result\":\"exit\"}"))
					blockdataChan <- *jsonData
				}
				//flush_db_nosync(&trx_cache,&erc20_address_trx_cache)

				fmt.Println("check exit")

				return
			}
		}
		bdata, _ := blockchainData.Get("result").MarshalJSON()

		var blockData util.BlockData
		err = json.Unmarshal(bdata, &blockData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		blockTime := getInt(blockData.Timestamp)
		blockReward := big.NewInt(0)
		for _, trxData := range blockData.Transactions {
			blockReward = blockReward.Add(blockReward, sc.handleTransaction(linkClient, &trxData, blockTime))
		}

		tmpHeight := getInt(blockData.Number)
		if tmpHeight%interval == 0 {
			runtime.GC()
		}

	}

}

func getInt(intStr string) int64 {
	if strings.HasPrefix(intStr, "0x") {
		value, _ := strconv.ParseInt(intStr[2:], 16, 32)
		return value
	}
	value, _ := strconv.ParseInt(intStr, 10, 32)
	return value
}

// TrxData is struct
type TrxData struct {
	AccessList           []interface{} `json:"accessList"`
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	ChainID              string        `json:"chainId"`
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

func (sc *Scan) handleTransaction(linkClient *util.LinkClient, trxData *interface{}, blockTime int64) *big.Int {

	trxMapObj := make(map[string]interface{})
	reward := big.NewInt(0)
	var trxSimpleData TrxData
	bdata, _ := json.Marshal(*trxData)
	err := json.Unmarshal(bdata, &trxSimpleData)
	if err != nil {
		fmt.Println("transaction Decode Failed!", err)
		return reward
	}
	param := make([]interface{}, 0, 20)
	param = append(param, trxSimpleData.Hash)
	trxReceiptData, _ := linkClient.SafeLinkHTTPFunc("eth_getTransactionReceipt", &param).Get("result").MarshalJSON()
	var trxReceipt util.TrxReceiptData
	err = json.Unmarshal(trxReceiptData, &trxReceipt)
	if err != nil {
		fmt.Println(err)
		return reward
	}
	trxMapObj["hash"] = trxSimpleData.Hash
	trxMapObj["blockNumber"] = trxSimpleData.BlockNumber
	trxMapObj["blockHash"] = trxSimpleData.BlockHash
	trxMapObj["from"] = trxSimpleData.From
	trxMapObj["to"] = trxSimpleData.To
	trxMapObj["gasLimit"] = trxSimpleData.Gas
	gasUsed := util.GetBigInt(trxReceipt.GasUsed, 16)
	if gasUsed == nil {
		gasUsed = big.NewInt(0)
	}
	if trxSimpleData.Type == "0x0" {
		trxMapObj["gasPrice"] = trxSimpleData.GasPrice
		trxMapObj["maxPriority"] = ""
		trxMapObj["maxFee"] = ""
		if trxReceipt.EffectiveGasPrice != "" {
			efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
			gasP := big.NewInt(0).Sub(util.GetBigInt(trxSimpleData.GasPrice, 16), efGas)
			reward = reward.Mul(gasUsed, gasP)
		} else {

			reward = reward.Mul(gasUsed, util.GetBigInt(trxSimpleData.GasPrice, 16))
		}

		trxMapObj["fee"] = big.NewInt(0).Mul(util.GetBigInt(trxSimpleData.GasPrice, 16), gasUsed).String()
	} else if trxSimpleData.Type == "0x2" {
		trxMapObj["maxPriority"] = trxSimpleData.MaxPriorityFeePerGas
		trxMapObj["maxFee"] = trxSimpleData.MaxFeePerGas
		trxMapObj["gasPrice"] = trxReceipt.EffectiveGasPrice
		efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
		maxF := util.GetBigInt(trxSimpleData.MaxFeePerGas, 16)

		reward = reward.Mul(gasUsed, maxF)
		trxMapObj["fee"] = big.NewInt(0).Mul(gasUsed, efGas).String()
	} else {
		trxMapObj["gasPrice"] = trxSimpleData.GasPrice
		trxMapObj["maxPriority"] = ""
		trxMapObj["maxFee"] = ""
		if trxReceipt.EffectiveGasPrice != "" {
			efGas := util.GetBigInt(trxReceipt.EffectiveGasPrice, 16)
			gasP := big.NewInt(0).Sub(util.GetBigInt(trxSimpleData.GasPrice, 16), efGas)
			reward = reward.Mul(gasUsed, gasP)
		} else {

			reward = reward.Mul(gasUsed, util.GetBigInt(trxSimpleData.GasPrice, 16))
		}

		trxMapObj["fee"] = big.NewInt(0).Mul(util.GetBigInt(trxSimpleData.GasPrice, 16), gasUsed).String()
	}
	trxMapObj["to"] = trxSimpleData.To
	trxMapObj["blockTime"] = blockTime
	trxMapObj["transactionType"] = trxSimpleData.Type
	trxMapObj["value"] = trxSimpleData.Value
	trxMapObj["inputData"] = trxSimpleData.Input
	trxMapObj["gasUsed"] = trxReceipt.GasUsed
	trxMapObj["status"] = trxReceipt.Status
	trxMapObj["nonce"] = trxSimpleData.Nonce
	trxMapObj["transactionIndex"] = trxSimpleData.TransactionIndex

	if trxReceipt.Status == "0x1" {
		for _, oneLog := range trxReceipt.Logs {
			// CanQuote()

			if len(oneLog.Topics) > 0 {
				if sc.allEvent.CheckEqual(sc.allEvent.CanQuote, oneLog.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					log.Printf("find CanQuote event %v", oneLog)
					contractAddr := oneLog.Address
					callData := util.GenerateOrderCall(contractAddr)

					callRes := linkClient.EthCall(util.GetOrderFactory(sc.Config), callData)
					orderID := util.GetBigInt(callRes, 16)

					if orderID.Int64() > 0 {
						var orderBaseO util.Order
						for {
							conn, err := ethclient.Dial(sc.Config.NodeURL)
							if err != nil {
								log.Printf("Failed to connect to the Ethereum client %v", err)
							} else {
								defer conn.Close()
								orderBase, err := util.NewOrderBase(common.HexToAddress(contractAddr), conn)
								if err != nil {
									log.Printf("call orderBase failed %v", err)
									continue
								}
								orderBaseO, err = orderBase.OrderInfo(nil)
								if err != nil {
									log.Printf("call orderBase failed %v", err)
									continue
								}
								break
							}
							time.Sleep(5 * time.Second)
						}
						NeedBidObj := new(util.NeedBid)
						NeedBidObj.CPU = orderBaseO.VCpu
						NeedBidObj.Cert = orderBaseO.CertKey
						NeedBidObj.Memory = orderBaseO.VMemory
						NeedBidObj.Storage = orderBaseO.VStorage
						NeedBidObj.SdlTrxID = orderBaseO.TrxId.Text(16)
						NeedBidObj.ContractAddress = contractAddr
						NeedBidObj.State = orderBaseO.State
						sc.NeedBidChan <- *NeedBidObj
						log.Printf("insert  CanQuote obj %v", *NeedBidObj)
					}

				} else if sc.allEvent.CheckEqual(sc.allEvent.ChooseQuote, oneLog.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					log.Printf("find ChooseQuote event %v", oneLog)
					contractAddr := oneLog.Address
					callData := util.GenerateOrderCall(contractAddr)

					callRes := linkClient.EthCall(util.GetOrderFactory(sc.Config), callData)
					orderID := util.GetBigInt(callRes, 16)
					if orderID.Int64() > 0 {
						OrderBaseFilter, _ := util.NewOrderBaseFilterer(common.HexToAddress(contractAddr), nil)
						CanQuoteEvent, err := OrderBaseFilter.ParseChooseQuote(util.ConvertLogToGethLogs(oneLog))
						if err == nil {
							needCreate := new(util.NeedCreate)

							needCreate.Provider = CanQuoteEvent.Provider
							needCreate.FinalPrice = CanQuoteEvent.FinalPrice
							needCreate.ContractAddress = contractAddr
							sc.NeedCreateChan <- *needCreate
							log.Printf("insert NeedCreateChan event %v", *needCreate)
						}
					}

				} else if sc.allEvent.CheckEqual(sc.allEvent.UserCancelOrder, oneLog.Topics[0]) {
					// Check whether the order is a valid order
					// put in canBid chan
					log.Printf("find UserCancelOrder event %v", oneLog)
					contractAddr := oneLog.Address
					callData := util.GenerateOrderCall(contractAddr)

					callRes := linkClient.EthCall(util.GetOrderFactory(sc.Config), callData)
					orderID := util.GetBigInt(callRes, 16)

					if orderID.Int64() > 0 {
						userCancel := new(util.UserCancelOrder)
						userCancel.ContractAddress = common.HexToAddress(contractAddr)
						sc.UserCancelChan <- *userCancel
						log.Printf("insert UserCancelOrder event %v", *userCancel)
					}

				}

			}

		}
		//trace log

		traceParam := make([]interface{}, 0, 2)
		traceParam = append(traceParam, trxSimpleData.Hash)

	}

	return reward

	//from

}

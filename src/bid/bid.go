package bid

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/config"
	"providerService/src/util"
	ubic_cluster "providerService/ubic-cluster"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	orderStatusQuoting int64 = 1
	orderStatusRunning int64 = 2
	orderStatusEnded   int64 = 3
)

type resourceStorage struct {
	CPUCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}

// Service is bid service
type Service struct {
	BidChan          <-chan util.NeedBid
	BidFinalChan     <-chan util.NeedCreate
	OrderFinish      <-chan util.UserCancelOrder
	BidTimeout       <-chan time.Time
	Total            resourceStorage
	Client           *ethclient.Client
	Conf             *config.ProviderConfig
	Cluster          *ubic_cluster.UbicService
	KeepResource     sync.Map
	KeepResourceTime sync.Map
	MutexRw          *sync.RWMutex
	Abi              map[string]abi.ABI
	Ctx              context.Context
	WgBid            *sync.WaitGroup
	LastPayTime      int64
}

// Init is service initialize function
func (bs *Service) Init(ctx context.Context, config *config.ProviderConfig,
	bidChan <-chan util.NeedBid,
	bidFinal <-chan util.NeedCreate,
	orderFinish <-chan util.UserCancelOrder,
	cluster *ubic_cluster.UbicService) {
	bs.Client, _ = ethclient.Dial(config.NodeURL)
	bs.Conf = config
	bs.BidFinalChan = bidFinal
	bs.BidChan = bidChan
	bs.OrderFinish = orderFinish
	bs.MutexRw = new(sync.RWMutex)
	bs.Abi = GetInteractiveABI()
	bs.Cluster = cluster
	bs.Total = bs.getTotalResource()
	bs.BidTimeout = time.After(time.Duration(config.BidTimeOut) * time.Second)
	bs.Ctx = ctx
	bs.WgBid = new(sync.WaitGroup)
	bs.LastPayTime = time.Now().Unix()
	bs.initExistDeployment()
}
func (bs *Service) updateResource() {
	lens := 0
	bs.KeepResource.Range(func(key, value interface{}) bool {
		lens++
		return true
	})
	zeroBig := new(big.Int).SetInt64(0)
	if lens == 0 &&
		bs.Total.CPUCount.Cmp(zeroBig) == 0 &&
		bs.Total.MemoryCount.Cmp(zeroBig) == 0 &&
		bs.Total.StorageCount.Cmp(zeroBig) == 0 {
		/*
			avaTotalTemp, err := bs.Cluster.GetTotalAvailable()
			if err != nil {
				return
			}*/
		totalTemp, err := bs.Cluster.GetTotalAvailable()
		if err != nil {
			return
		}
		providerContractAddr := common.HexToAddress(bs.Conf.ProviderContract)
		privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		method := "updateResource"
		data, err := bs.Abi[ProviderName].Pack(method, new(big.Int).SetUint64(totalTemp.CPU), new(big.Int).SetUint64(totalTemp.Memory), new(big.Int).SetUint64(totalTemp.StorageEphemeral))
		if err != nil {
			fmt.Println("err is ", err.Error())
		}
		msg := ethereum.CallMsg{
			From: fromAddress,
			To:   &providerContractAddr,
			Gas:  0,
			Data: data,
		}
		bs.Client.CallContract(context.Background(), msg, nil)
	}
}
func (bs *Service) getAllProviderServOrders() []util.Order {
	orderBase, _ := util.NewOrderFactory(common.HexToAddress(util.GetOrderFactory(bs.Conf)), bs.Client)
	allProviderOrder, _ := orderBase.GetProviderAllOrder(nil, common.HexToAddress(bs.Conf.ProviderContract))
	return allProviderOrder
}
func (bs *Service) initExistDeployment() {
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	orderInChain := bs.getAllProviderServOrders()
	leaseLocal := make(map[string]int, 0)
	for _, lease := range allActiveLeases {
		state, orderAddr := bs.getOrderState(lease.OSeq)
		fmt.Println("initExistDeployment", lease, state, orderAddr)
		leaseLocal[strings.ToLower(orderAddr)] = 1
		if state == orderStatusQuoting {
			uri := bs.Cluster.GetURI(lease)
			bs.submitURL(orderAddr, uri)
			_, ok := bs.KeepResource.Load(strings.ToLower(orderAddr))
			if !ok {
				bs.KeepResource.Store(strings.ToLower(orderAddr), resourceStorage{
					CPUCount:     big.NewInt(0),
					MemoryCount:  big.NewInt(0),
					StorageCount: big.NewInt(0),
				})
			}
			_, ok = bs.KeepResourceTime.Load(strings.ToLower(orderAddr))
			if !ok {
				bs.KeepResourceTime.Store(strings.ToLower(orderAddr), time.Now().Unix())
			}
		} else if state == orderStatusRunning {
			_, ok := bs.KeepResource.Load(strings.ToLower(orderAddr))
			if !ok {
				bs.KeepResource.Store(strings.ToLower(orderAddr), resourceStorage{
					CPUCount:     big.NewInt(0),
					MemoryCount:  big.NewInt(0),
					StorageCount: big.NewInt(0),
				})
			}
		} else if state == orderStatusEnded {
			bs.Cluster.CloseManager(lease)
		}
	}
	for _, orderSingle := range orderInChain {
		if _, ok := leaseLocal[strings.ToLower(orderSingle.ContractAddress.String())]; ok {
			continue
		}
		if int64(orderSingle.State) == orderStatusRunning {
			sdlFile := bs.getSdlByID(orderSingle.ContractAddress.String())
			owner := bs.getOwner(orderSingle.ContractAddress.String())
			index := bs.getOrderIndex(orderSingle.ContractAddress.String())
			lid := ctypes.LeaseID{
				Owner:    "0x" + strings.TrimLeft(owner, "000000000000000000000000"),
				OSeq:     index,
				Provider: bs.Conf.ProviderContract,
			}
			uri, _ := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
			bs.submitURL(orderSingle.ContractAddress.String(), uri)
		}
	}
}

func (bs *Service) initTotalResource() {
	avaTotal, err := bs.Cluster.GetTotalAvailable()
	if err != nil {
		log.Println("Bid get Left Total Resource error")
		return
	}
	bs.Total.CPUCount = new(big.Int).SetUint64(avaTotal.CPU)
	bs.Total.MemoryCount = new(big.Int).SetUint64(avaTotal.Memory)
	bs.Total.StorageCount = new(big.Int).SetUint64(avaTotal.StorageEphemeral)
}
func (bs *Service) getProviderAddr() string {
	providerFactoryAddr := common.HexToAddress(bs.Conf.ProviderFactoryContract)
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	method := "getProvideContract"
	data, err := bs.Abi[ProviderFactoryName].Pack(method, fromAddress)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &providerFactoryAddr,
		Gas:  0,
		Data: data,
	}
	result, err := bs.Client.CallContract(context.Background(), msg, nil)
	fmt.Println("result is", common.BytesToAddress(result).String())
	return common.BytesToAddress(result).String()
}
func (bs *Service) getTotalResource() resourceStorage {
	providerContractAddr := common.HexToAddress(bs.Conf.ProviderContract)
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	method := "getLeftResource"
	data, err := bs.Abi[ProviderName].Pack(method)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &providerContractAddr,
		Gas:  0,
		Data: data,
	}
	result, err := bs.Client.CallContract(context.Background(), msg, nil)
	var ret resourceStorage
	ret.CPUCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[0:32]), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[32:64]), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[64:]), 16)
	fmt.Println("result is", ret)
	return ret
}
func (bs *Service) handleResource(resource resourceStorage, add bool) {
	bs.MutexRw.Lock()
	defer bs.MutexRw.Unlock()
	if add {
		bs.Total.CPUCount = bs.Total.CPUCount.Add(bs.Total.CPUCount, resource.CPUCount)
		bs.Total.MemoryCount = bs.Total.MemoryCount.Add(bs.Total.MemoryCount, resource.MemoryCount)
		bs.Total.StorageCount = bs.Total.StorageCount.Add(bs.Total.StorageCount, resource.StorageCount)
	} else {
		bs.Total.CPUCount = bs.Total.CPUCount.Sub(bs.Total.CPUCount, resource.CPUCount)
		bs.Total.MemoryCount = bs.Total.MemoryCount.Sub(bs.Total.MemoryCount, resource.MemoryCount)
		bs.Total.StorageCount = bs.Total.StorageCount.Sub(bs.Total.StorageCount, resource.StorageCount)
	}
}
func (bs *Service) quoteBidOrder(orderContractAddr string) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0)      // in wei (1 eth)
	gasLimit := uint64(3000000) // in units
	gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(orderContractAddr)
	cpuPrice, _ := new(big.Int).SetString(bs.Conf.CPUPrice, 10)
	memoryPrice, _ := new(big.Int).SetString(bs.Conf.MemoryPrice, 10)
	storagePrice, _ := new(big.Int).SetString(bs.Conf.StoragePrice, 10)
	method := "quote"
	data, _ := bs.Abi[OrderBaseName].Pack(method, cpuPrice, memoryPrice, storagePrice)
	c := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}
	tx := types.NewTx(c)
	chainID, _ := new(big.Int).SetString(bs.Conf.NodeChainID, 10)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent: ", signedTx.Hash().Hex())
}
func (bs *Service) submitURL(orderContractAddr string, url string) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0)      // in wei (1 eth)
	gasLimit := uint64(3000000) // in units
	gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(orderContractAddr)

	method := "submit_server_uri"
	data, _ := bs.Abi[OrderBaseName].Pack(method, url)
	c := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}
	tx := types.NewTx(c)
	chainID, _ := new(big.Int).SetString(bs.Conf.NodeChainID, 10)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("submit url tx sent: ", signedTx.Hash().Hex())

}
func (bs *Service) getSdlByID(orderContractAddr string) []byte {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodSdlTrx := "o_sdl_trx_id"
	dataSdlTrx, _ := bs.Abi[OrderBaseName].Pack(methodSdlTrx)
	msgSdl := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataSdlTrx,
	}
	resultSdl, _ := bs.Client.CallContract(context.Background(), msgSdl, nil)
	sdlTrxID := new(common.Hash)
	sdlTrxID.SetBytes(resultSdl)
	sdlTrx, _, _ := bs.Client.TransactionByHash(context.Background(), *sdlTrxID)
	return sdlTrx.Data()
}
func (bs *Service) getOwner(orderContractAddr string) string {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodSdlTrx := "owner"
	dataSdlTrx, _ := bs.Abi[OrderBaseName].Pack(methodSdlTrx)
	msgSdl := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataSdlTrx,
	}
	resultAddress, _ := bs.Client.CallContract(context.Background(), msgSdl, nil)
	return common.Bytes2Hex(resultAddress)
}
func (bs *Service) getOrderState(index uint64) (int64, string) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(bs.Conf.OrderFactory)
	methodOrders := "orders"
	dataOrderTrx, _ := bs.Abi[OrderFactoryName].Pack(methodOrders, new(big.Int).SetUint64(index))
	msgOrder := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataOrderTrx,
	}
	resultOrder, err := bs.Client.CallContract(context.Background(), msgOrder, nil)
	if err != nil {
		fmt.Println("result Order", err.Error())
	}
	hexOrderAddr := common.BytesToAddress(resultOrder)
	methodOrderState := "order_status"
	dataOrderStateTrx, _ := bs.Abi[OrderBaseName].Pack(methodOrderState)
	msgOrderState := ethereum.CallMsg{
		From: fromAddress,
		To:   &hexOrderAddr,
		Gas:  0,
		Data: dataOrderStateTrx,
	}
	resultOrderState, _ := bs.Client.CallContract(context.Background(), msgOrderState, nil)
	state, _ := strconv.ParseInt(common.Bytes2Hex(resultOrderState), 16, 64)
	return state, hexOrderAddr.String()
}
func (bs *Service) getOrderIndex(orderContractAddr string) uint64 {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodSdlTrx := "o_order_number"
	dataSdlTrx, _ := bs.Abi[OrderBaseName].Pack(methodSdlTrx)
	msgSdl := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataSdlTrx,
	}
	result, _ := bs.Client.CallContract(context.Background(), msgSdl, nil)
	hex := common.Bytes2Hex(result)
	index, _ := strconv.ParseUint(hex, 16, 64)
	return index
}
func (bs *Service) getOrderCount(orderContractAddr string) resourceStorage {

	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	methodCPU := "o_cpu"
	methodMemory := "o_memory"
	methodStorage := "o_storage"
	dataCPU, err := bs.Abi[ProviderName].Pack(methodCPU)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	dataMemory, err := bs.Abi[ProviderName].Pack(methodMemory)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	dataStorage, err := bs.Abi[ProviderName].Pack(methodStorage)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	toAddress := common.HexToAddress(orderContractAddr)
	msgCPU := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataCPU,
	}
	msgMemory := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataMemory,
	}
	msgStorage := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataStorage,
	}
	resultCPU, _ := bs.Client.CallContract(context.Background(), msgCPU, nil)
	resultMemory, _ := bs.Client.CallContract(context.Background(), msgMemory, nil)
	resultStorage, _ := bs.Client.CallContract(context.Background(), msgStorage, nil)
	var ret resourceStorage
	ret.CPUCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultCPU), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultMemory), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultStorage), 16)
	fmt.Println("result is", ret)
	return ret
}
func (bs *Service) payBill(orderContractAddr string) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0)      // in wei (1 eth)
	gasLimit := uint64(3000000) // in units
	gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
	toAddress := common.HexToAddress(orderContractAddr)
	methodBillTrx := "_pay_billing"
	dataPayBillTrx, _ := bs.Abi[OrderBaseName].Pack(methodBillTrx, fromAddress)
	c := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     dataPayBillTrx,
	}
	tx := types.NewTx(c)
	chainID, _ := new(big.Int).SetString(bs.Conf.NodeChainID, 10)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pay Bill tx sent: ", signedTx.Hash().Hex())
}
func (bs *Service) handleBid(orderInfo *util.NeedBid) {
	log.Println("in Handle bid")
	log.Println(orderInfo.ContractAddress)
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	_, ok := bs.KeepResource.Load(strings.ToLower(orderInfo.ContractAddress))
	if ok {
		log.Println("This has handled")
		return
	}
	if bs.Total.CPUCount.Cmp(orderInfo.CPU) >= 0 &&
		bs.Total.MemoryCount.Cmp(orderInfo.Memory) >= 0 &&
		bs.Total.StorageCount.Cmp(orderInfo.Storage) >= 0 {
		bs.quoteBidOrder(orderInfo.ContractAddress)
		bs.KeepResource.Store(strings.ToLower(orderInfo.ContractAddress), resourceStorage{orderInfo.CPU, orderInfo.Memory, orderInfo.Storage})
		bs.KeepResourceTime.Store(strings.ToLower(orderInfo.ContractAddress), time.Now().Unix())
		bs.handleResource(resourceStorage{orderInfo.CPU, orderInfo.Memory, orderInfo.Storage}, false)
	}
}
func (bs *Service) handleBidFinal(bidFinalInfo *util.NeedCreate) {
	log.Println("in handle bid final", bidFinalInfo.ContractAddress)
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	privateKey, _ := crypto.HexToECDSA(bs.Conf.SecretKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	providerAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	resource, ok := bs.KeepResource.Load(strings.ToLower(bidFinalInfo.ContractAddress))
	bs.KeepResourceTime.Delete(strings.ToLower(bidFinalInfo.ContractAddress))
	log.Println("in handle bid final providers ", bidFinalInfo.Provider.String(), providerAddr.String())
	if bidFinalInfo.Provider.String() != bs.Conf.ProviderContract {
		if ok {
			bs.handleResource(resource.(resourceStorage), true)
			return
		}
	} else {
		if !ok {
			orderSource := bs.getOrderCount(bidFinalInfo.ContractAddress)
			bs.KeepResource.Store(strings.ToLower(bidFinalInfo.ContractAddress), orderSource)
			bs.handleResource(orderSource, false)
		}
		log.Println(bs.Total)
		sdlFile := bs.getSdlByID(bidFinalInfo.ContractAddress)
		owner := bs.getOwner(bidFinalInfo.ContractAddress)
		index := bs.getOrderIndex(bidFinalInfo.ContractAddress)
		fmt.Println(owner, index, bidFinalInfo.Provider.String())
		lid := ctypes.LeaseID{
			Owner:    "0x" + strings.TrimLeft(owner, "000000000000000000000000"),
			OSeq:     index,
			Provider: bidFinalInfo.Provider.String(),
		}
		uri, _ := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
		bs.submitURL(bidFinalInfo.ContractAddress, uri)
	}
}
func (bs *Service) handleOrderFinish(orderFinishInfo *util.UserCancelOrder) {
	log.Println("in Handle HandleOrderFinish", orderFinishInfo.ContractAddress)
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	resource, ok := bs.KeepResource.Load(strings.ToLower(orderFinishInfo.ContractAddress.String()))
	if ok {
		bs.handleResource(resource.(resourceStorage), true)
	}
	owner := bs.getOwner(orderFinishInfo.ContractAddress.String())
	index := bs.getOrderIndex(orderFinishInfo.ContractAddress.String())
	provider := bs.getProviderAddr()
	lid := ctypes.LeaseID{
		Owner:    "0x" + strings.TrimLeft(owner, "000000000000000000000000"),
		OSeq:     index,
		Provider: provider,
	}
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	for _, lease := range allActiveLeases {
		if lease.Equals(lid) {
			bs.Cluster.CloseManager(lease)
			fmt.Println("HandleOrderFinish CloseManager")
			break
		}
	}
}
func (bs *Service) handleRecoverResource() {
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	contractAdders := make([]string, 0)
	bs.KeepResourceTime.Range(func(key any, value any) bool {
		if value.(int64)+bs.Conf.BidTimeOut < time.Now().Unix() {
			resourceTemp, ok := bs.KeepResource.Load(key)
			if ok {
				bs.handleResource(resourceTemp.(resourceStorage), true)
			}
			contractAdders = append(contractAdders, key.(string))
		}
		return true
	})
	for _, value := range contractAdders {
		bs.KeepResourceTime.Delete(strings.ToLower(value))
	}
	if bs.LastPayTime+10*bs.Conf.BidTimeOut < time.Now().Unix() {
		allActiveLeases := bs.Cluster.GetAllActiveLeases()
		for _, lease := range allActiveLeases {
			_, orderAddr := bs.getOrderState(lease.OSeq)
			bs.payBill(orderAddr)
		}
		bs.LastPayTime = time.Now().Unix()
	}
}

// Run is start bid service
func (bs *Service) Run(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
loop:
	for {
		select {
		case bid := <-bs.BidChan:
			go bs.handleBid(&bid)
		case bidFinalChan := <-bs.BidFinalChan:
			go bs.handleBidFinal(&bidFinalChan)
		case orderFinish := <-bs.OrderFinish:
			go bs.handleOrderFinish(&orderFinish)
		case <-bs.BidTimeout:
			go bs.handleRecoverResource()
			bs.BidTimeout = time.After(time.Duration(bs.Conf.BidTimeOut) * time.Second)
		case <-bs.Ctx.Done():
			bs.WgBid.Wait()
			log.Println("bid service exit")
			break loop
		}
	}
}

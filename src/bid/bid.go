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
	"providerService/ubic-cluster"
	"strconv"
	"sync"
	"time"
)

type ResourceStorage struct {
	CpuCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}

type BidService struct {
	BidChan          <-chan util.NeedBid
	BidFinalChan     <-chan util.NeedCreate
	OrderFinish      <-chan util.UserCancelOrder
	BidTimeout       <-chan time.Time
	Total            ResourceStorage
	Client           *ethclient.Client
	Conf             *config.ProviderConfig
	Cluster          *ubic_cluster.UbicService
	KeepResource     sync.Map
	KeepResourceTime sync.Map
	MutexRw          *sync.RWMutex
	Abi              map[string]abi.ABI
	Ctx              context.Context
	WgGlobal         *sync.WaitGroup
	WgBid            *sync.WaitGroup
}

func (bs *BidService) Init(ctx context.Context, wg *sync.WaitGroup, config *config.ProviderConfig,
	bidChan <-chan util.NeedBid,
	bidFinal <-chan util.NeedCreate,
	orderFinish <-chan util.UserCancelOrder,
	cluster *ubic_cluster.UbicService) {
	bs.Client, _ = ethclient.Dial(config.NodeUrl)
	bs.Conf = config
	bs.BidFinalChan = bidFinal
	bs.BidChan = bidChan
	bs.OrderFinish = orderFinish
	bs.MutexRw = new(sync.RWMutex)
	bs.Abi = GetInteractiveABI()
	bs.Total = bs.GetTotalResource()
	bs.Cluster = cluster
	bs.BidTimeout = time.After(time.Duration(config.BidTimeOut) * time.Second)
	bs.Ctx = ctx
	bs.WgGlobal = wg
	bs.WgGlobal.Add(1)
	bs.WgBid = new(sync.WaitGroup)
}
func (bs *BidService) GetProviderAddr() string {
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
func (bs *BidService) GetTotalResource() ResourceStorage {
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
	var ret ResourceStorage
	ret.CpuCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[0:32]), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[32:64]), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[64:]), 16)
	fmt.Println("result is", ret)
	return ret
}
func (bs *BidService) HandleResource(resource ResourceStorage, add bool) {
	bs.MutexRw.Lock()
	defer bs.MutexRw.Unlock()
	if add {
		bs.Total.CpuCount = bs.Total.CpuCount.Add(bs.Total.CpuCount, resource.CpuCount)
		bs.Total.MemoryCount = bs.Total.MemoryCount.Add(bs.Total.MemoryCount, resource.MemoryCount)
		bs.Total.StorageCount = bs.Total.StorageCount.Add(bs.Total.StorageCount, resource.StorageCount)
	} else {
		bs.Total.CpuCount = bs.Total.CpuCount.Sub(bs.Total.CpuCount, resource.CpuCount)
		bs.Total.MemoryCount = bs.Total.MemoryCount.Sub(bs.Total.MemoryCount, resource.MemoryCount)
		bs.Total.StorageCount = bs.Total.StorageCount.Sub(bs.Total.StorageCount, resource.StorageCount)
	}
}
func (bs *BidService) quoteBidOrder(orderContractAddr string) {
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
	cpuPrice, _ := new(big.Int).SetString(bs.Conf.CpuPrice, 10)
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
	chainId, _ := new(big.Int).SetString(bs.Conf.NodeChainId, 10)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent: ", signedTx.Hash().Hex())
}
func (bs *BidService) SubmitUrl(orderContractAddr string, url string) {
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
	chainId, _ := new(big.Int).SetString(bs.Conf.NodeChainId, 10)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent: ", signedTx.Hash().Hex())

}
func (bs *BidService) GetSdlById(orderContractAddr string) []byte {
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
	sdlTrxId := new(common.Hash)
	sdlTrxId.SetBytes(resultSdl)
	sdlTrx, _, _ := bs.Client.TransactionByHash(context.Background(), *sdlTrxId)
	return sdlTrx.Data()
}
func (bs *BidService) GetOwner(orderContractAddr string) string {
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
func (bs *BidService) GetOrderIndex(orderContractAddr string) uint64 {
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
	methodSdlTrx := "checkIsOrder"
	dataSdlTrx, _ := bs.Abi[OrderFactoryName].Pack(methodSdlTrx)
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
func (bs *BidService) GetOrderCount(orderContractAddr string) ResourceStorage {

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
	methodCpu := "o_cpu"
	methodMemory := "o_memory"
	methodStorage := "o_storage"
	dataCpu, err := bs.Abi[ProviderName].Pack(methodCpu)
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
	msgCpu := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataCpu,
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
	resultCpu, _ := bs.Client.CallContract(context.Background(), msgCpu, nil)
	resultMemory, _ := bs.Client.CallContract(context.Background(), msgMemory, nil)
	resultStorage, _ := bs.Client.CallContract(context.Background(), msgStorage, nil)
	var ret ResourceStorage
	ret.CpuCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultCpu), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultMemory), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultStorage), 16)
	fmt.Println("result is", ret)
	return ret
}
func (bs *BidService) HandleBid(orderInfo *util.NeedBid) {
	log.Println("in Handle bid")
	log.Println(orderInfo.ContractAddress)
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	_, ok := bs.KeepResource.Load(orderInfo.ContractAddress)
	if ok {
		log.Println("This has handled")
		return
	}
	if bs.Total.CpuCount.Cmp(orderInfo.Cpu) >= 0 &&
		bs.Total.MemoryCount.Cmp(orderInfo.Memory) >= 0 &&
		bs.Total.StorageCount.Cmp(orderInfo.Storage) >= 0 {
		bs.quoteBidOrder(orderInfo.ContractAddress)
		bs.KeepResource.Store(orderInfo.ContractAddress, ResourceStorage{orderInfo.Cpu, orderInfo.Memory, orderInfo.Storage})
		bs.KeepResourceTime.Store(orderInfo.ContractAddress, time.Now().Unix())
		bs.HandleResource(ResourceStorage{orderInfo.Cpu, orderInfo.Memory, orderInfo.Storage}, false)
	}
}
func (bs *BidService) HandleBidFinal(bidFinalInfo *util.NeedCreate) {
	log.Println("in handle bid final")
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	privateKey, _ := crypto.HexToECDSA(bs.Conf.SecretKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	providerAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	resource, ok := bs.KeepResource.Load(bidFinalInfo.ContractAddress)
	bs.KeepResourceTime.Delete(bidFinalInfo.ContractAddress)
	if bidFinalInfo.Provider.String() != providerAddr.String() {
		if ok {
			bs.HandleResource(resource.(ResourceStorage), true)
			return
		}
	} else {
		if !ok {
			orderSource := bs.GetOrderCount(bidFinalInfo.ContractAddress)
			bs.KeepResource.Store(bidFinalInfo.ContractAddress, orderSource)
			bs.HandleResource(orderSource, false)
		}
		log.Println(bs.Total)
		sdlFile := bs.GetSdlById(bidFinalInfo.ContractAddress)
		owner := bs.GetOwner(bidFinalInfo.ContractAddress)
		index := bs.GetOrderIndex(bidFinalInfo.ContractAddress)
		lid := ctypes.LeaseID{
			Owner:    owner,
			OSeq:     index,
			Provider: bidFinalInfo.Provider.String(),
		}
		uris, _ := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
		var uri string
		for _, v := range uris {
			uri += v
		}
		bs.SubmitUrl(bidFinalInfo.ContractAddress, uri)
	}
}
func (bs *BidService) HandleOrderFinish(orderFinishInfo *util.UserCancelOrder) {
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	resource, ok := bs.KeepResource.Load(orderFinishInfo.ContractAddress.String())
	if !ok {
		return
	}
	bs.HandleResource(resource.(ResourceStorage), true)
	owner := bs.GetOwner(orderFinishInfo.ContractAddress.String())
	index := bs.GetOrderIndex(orderFinishInfo.ContractAddress.String())
	provider := bs.GetProviderAddr()
	lid := ctypes.LeaseID{
		Owner:    owner,
		OSeq:     index,
		Provider: provider,
	}
	bs.Cluster.CloseManager(lid)
}
func (bs *BidService) HandleRecoverResource() {
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	contractAddrs := make([]string, 0)
	bs.KeepResourceTime.Range(func(key any, value any) bool {
		if value.(int64)+bs.Conf.BidTimeOut < time.Now().Unix() {
			resourceTemp, ok := bs.KeepResource.Load(key)
			if ok {
				bs.HandleResource(resourceTemp.(ResourceStorage), true)
			}
			contractAddrs = append(contractAddrs, key.(string))
		}
		return true
	})
	for _, value := range contractAddrs {
		bs.KeepResourceTime.Delete(value)
	}
}

func (bs *BidService) Run() {
	for {
		select {
		case bid := <-bs.BidChan:
			go bs.HandleBid(&bid)
		case bidFinalChan := <-bs.BidFinalChan:
			go bs.HandleBidFinal(&bidFinalChan)
		case orderFinish := <-bs.OrderFinish:
			go bs.HandleOrderFinish(&orderFinish)
		case <-bs.BidTimeout:
			go bs.HandleRecoverResource()
			bs.BidTimeout = time.After(time.Duration(bs.Conf.BidTimeOut) * time.Second)
		case <-bs.Ctx.Done():
			bs.WgBid.Wait()
			bs.WgGlobal.Done()
			log.Println("bid service exit")
		}
	}
}

package bid

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/util"
	"strconv"
	"strings"
)

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
	orderBase, err := util.NewOrderBase(common.HexToAddress(orderContractAddr), bs.Client)
	if err != nil {
		log.Println("quoteBidOrder error create order base", err.Error())
		return
	}
	orderState, err := orderBase.OrderStatus(nil)
	if err != nil {
		log.Println("quoteBidOrder error get order state", err.Error())
		return
	}
	fmt.Println("order state is ", orderState)
	if int64(orderState) == orderStatusQuoting {
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
}
func (bs *Service) submitURI(orderContractAddr string, uri string) {
	orderBase, err := util.NewOrderBase(common.HexToAddress(orderContractAddr), bs.Client)
	if err != nil {
		log.Println("submitURI error", err.Error())
	}
	resultURI, err := orderBase.ServerUri(nil)
	if err != nil {
		log.Println("submitURI error", err.Error())
	}
	resultState, err := orderBase.OrderStatus(nil)

	fmt.Println("uri is ", uri)
	if resultURI != uri && int64(resultState) == orderStatusRunning {
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
		data, _ := bs.Abi[OrderBaseName].Pack(method, uri)
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
		fmt.Println("submit uri tx sent: ", signedTx.Hash().Hex())
	}
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
func (bs *Service) getOwner(orderContractAddr string) common.Address {
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
	return common.BytesToAddress(resultAddress)
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
	dataCPU, err := bs.Abi[OrderBaseName].Pack(methodCPU)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	dataMemory, err := bs.Abi[OrderBaseName].Pack(methodMemory)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	dataStorage, err := bs.Abi[OrderBaseName].Pack(methodStorage)
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
	orderBase, err := util.NewOrderBase(common.HexToAddress(orderContractAddr), bs.Client)
	if err != nil {
		log.Println("quoteBidOrder error create order base", err.Error())
		return
	}
	orderState, err := orderBase.OrderStatus(nil)
	if err != nil {
		log.Println("quoteBidOrder error get order state", err.Error())
		return
	}
	if int64(orderState) == orderStatusRunning {
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
		methodBillTrx := "pay_billing"
		dataPayBillTrx, _ := bs.Abi[OrderBaseName].Pack(methodBillTrx)
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

}
func (bs *Service) getOrderLastPayTime(orderContractAddr string) int64 {
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
	methodLastPayTime := "last_pay_time"

	dataPayTime, err := bs.Abi[OrderBaseName].Pack(methodLastPayTime)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
	toAddress := common.HexToAddress(orderContractAddr)
	msgLastPayTime := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataPayTime,
	}
	resultLastPayTime, _ := bs.Client.CallContract(context.Background(), msgLastPayTime, nil)
	lastPayTimeHex := common.Bytes2Hex(resultLastPayTime)
	lastPayTime, _ := strconv.ParseInt(lastPayTimeHex, 16, 64)
	return lastPayTime
}
func (bs *Service) updateResource() {
	chainTotalLeft := bs.getTotalResource()
	totalTemp, err := bs.Cluster.GetTotalAvailable()
	if err != nil {
		log.Println("update Resource cluster get error")
	}
	orders := bs.getAllProviderServOrders()
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	if len(orders) != 0 {
		leaseMap := make(map[ctypes.LeaseID]int, len(allActiveLeases))
		for _, lease := range allActiveLeases {
			leaseMap[lease] = 0
		}
		runningCount := 0
		allRunningCheck := false
		for _, order := range orders {
			if int64(order.State) == orderStatusRunning {
				runningCount += 1
				tempLease := ctypes.LeaseID{
					Owner:    order.Owner.String(),
					OSeq:     order.OrderId.Uint64(),
					Provider: common.HexToAddress(bs.Conf.ProviderContract).String(),
				}
				if _, ok := leaseMap[tempLease]; !ok {
					return
				} else {
					leaseMap[tempLease] = 1
				}
			}
		}
		if runningCount != len(allActiveLeases) {
			return
		} else {
			allRunningCheck = true
		}
		for _, value := range leaseMap {
			if value != 1 {
				allRunningCheck = false
				break
			}
		}
		if !allRunningCheck {
			return
		}
	} else {
		if len(allActiveLeases) != 0 {
			return
		}
	}

	if chainTotalLeft.CPUCount.Cmp(new(big.Int).SetUint64(totalTemp.CPU)) != 0 ||
		chainTotalLeft.MemoryCount.Cmp(new(big.Int).SetUint64(totalTemp.Memory)) != 0 ||
		chainTotalLeft.StorageCount.Cmp(new(big.Int).SetUint64(totalTemp.StorageEphemeral)) != 0 {
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
		msg := &types.LegacyTx{
			Nonce:    nonce,
			To:       &providerContractAddr,
			Gas:      gasLimit,
			Value:    value,
			GasPrice: gasPrice,
			Data:     data,
		}
		tx := types.NewTx(msg)
		chainID, _ := new(big.Int).SetString(bs.Conf.NodeChainID, 10)
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}
		err = bs.Client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("update source tx sent: ", signedTx.Hash().Hex())
		contractAdders := make([]string, 0)
		contractTempAddrs := make([]string, 0)
		bs.KeepResourceTime.Range(func(key any, value any) bool {
			contractAdders = append(contractAdders, key.(string))
			return true
		})
		for _, value := range contractAdders {
			bs.KeepResourceTime.Delete(strings.ToLower(value))
		}

		bs.KeepResource.Range(func(key any, value any) bool {
			contractTempAddrs = append(contractTempAddrs, key.(string))
			return true
		})
		for _, value := range contractTempAddrs {
			bs.KeepResource.Delete(strings.ToLower(value))
		}
		bs.Total.CPUCount = new(big.Int).SetUint64(totalTemp.CPU)
		bs.Total.MemoryCount = new(big.Int).SetUint64(totalTemp.Memory)
		bs.Total.StorageCount = new(big.Int).SetUint64(totalTemp.StorageEphemeral)
	}
}
func (bs *Service) getAllProviderServOrders() []util.Order {
	orderBase, err := util.NewOrderFactory(common.HexToAddress(util.GetOrderFactory(bs.Conf)), bs.Client)
	if err != nil {
		log.Fatal("getAllProviderServOrders create new order factory fail", err.Error())
		return nil
	}
	allProviderOrder, err := orderBase.GetProviderAllOrder(nil, common.HexToAddress(bs.Conf.ProviderContract))
	if err != nil {
		log.Fatal("getAllProviderServOrders get all order fail", err.Error())
		return nil
	}
	return allProviderOrder
}

package bid

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/util"
	"strconv"
	"strings"
	"time"
)

func (bs *Service) getProviderAddr() string {
	providerFactoryAddr := common.HexToAddress(bs.Conf.ProviderFactoryContract)
	/*
		privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
		if err != nil {
			log.Println("getProviderAddr:turn private key error", err.Error())
			return bs.Conf.ProviderContract
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Println("getProviderAddr: publicKey is not of type *ecdsa.PublicKey", err.Error())
			return bs.Conf.ProviderContract
		}*/
	fromAddress := common.HexToAddress(bs.Conf.ProviderAddress)
	method := "getProvideContract"
	data, err := bs.Abi[ProviderFactoryName].Pack(method, fromAddress)
	if err != nil {
		log.Println("getProviderAddr:pack abi error", err.Error())
		return bs.Conf.ProviderContract
	}
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &providerFactoryAddr,
		Gas:  0,
		Data: data,
	}
	result, err := bs.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Println("getProviderAddr:call contract error", err.Error())
		return bs.Conf.ProviderContract
	}
	return common.BytesToAddress(result).String()
}
func (bs *Service) getTotalResource() resourceStorage {
	providerContractAddr := common.HexToAddress(bs.Conf.ProviderContract)
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	var ret resourceStorage
	ret.CPUCount = new(big.Int).SetInt64(0)
	ret.MemoryCount = new(big.Int).SetInt64(0)
	ret.StorageCount = new(big.Int).SetInt64(0)
	if err != nil {
		log.Println("getTotalResource: private key error", err.Error())
		return ret
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("getTotalResource: publicKey is not of type *ecdsa.PublicKey", err.Error())
		return ret
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	method := "getLeftResource"
	data, err := bs.Abi[ProviderName].Pack(method)
	if err != nil {
		log.Println("getTotalResource:pack abi error", err.Error())
		return ret
	}
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &providerContractAddr,
		Gas:  0,
		Data: data,
	}
	result, err := bs.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Println("getTotalResource:call contract error", err.Error())
		return ret
	}
	if len(result) < 96 {
		return ret
	}
	ret.CPUCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[0:32]), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[32:64]), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(result[64:]), 16)
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
		orderObj := util.QuoteBidOrder{
			ProviderAddr: bs.Conf.ProviderContract,
			OrderAddr:    orderContractAddr,
			CpuPrice:     bs.Conf.CPUPrice,
			MemoryPrice:  bs.Conf.MemoryPrice,
			StoragePrice: bs.Conf.StoragePrice,
		}
		privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
		if err != nil {
			log.Println("quoteBidOrder:turn ecdsa", err.Error())
			return
		}
		orderObj.Sign(privateKey)
		if orderObj.SignMsg == "" {
			log.Println("quoteBidOrder: Communication sign fail")
			return
		}
		success, trxID := orderObj.Send(bs.Conf.PoolURI)
		if !success {
			log.Println("quoteBidOrder: send trx to pool fail")
			return
		}
		log.Println("quoteBidOrder tx sent: ", trxID)
		/*
			privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
			if err != nil {
				log.Println("quoteBidOrder:turn ecdsa", err.Error())
				return
			}
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Println("quoteBidOrder: publicKey is not of type *ecdsa.PublicKey")
				return
			}
			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Println("quoteBidOrder:get nonce error")
				return
			}
			value := big.NewInt(0)      // in wei (1 eth)
			gasLimit := uint64(3000000) // in units
			gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Println("quoteBidOrder:SuggestGasPrice error")
				return
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
				log.Println("quoteBidOrder:SignTx error")
				return
			}
			err = bs.Client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Println("quoteBidOrder:SendTransaction error")
				return
			}
			log.Println("quoteBidOrder tx sent: ", signedTx.Hash().Hex())*/
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

	log.Println("uri is ", uri)
	if resultURI != uri && int64(resultState) == orderStatusRunning {
		submitObj := util.SubmitURI{
			ProviderAddr: bs.Conf.ProviderContract,
			OrderAddr:    orderContractAddr,
			NewUrl:       uri,
		}
		privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
		if err != nil {
			log.Println("submitURI:turn ecdsa", err.Error())
			return
		}
		submitObj.Sign(privateKey)
		if submitObj.SignMsg == "" {
			log.Println("submitURI: Communication sign fail")
			return
		}
		success, trxID := submitObj.Send(bs.Conf.PoolURI)
		if !success {
			log.Println("submitURI: send trx to pool fail")
			return
		}
		log.Println("submitURI tx sent: ", trxID)
		/*
			privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
			if err != nil {
				log.Println("submitURI:HexToECDSA", err.Error())
				return
			}
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Println("submitURI:is not of type *ecdsa.PublicKey", err.Error())
				return
			}
			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Println("submitURI:PendingNonceAt", err.Error())
				return
			}
			value := big.NewInt(0)      // in wei (1 eth)
			gasLimit := uint64(3000000) // in units
			gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Println("submitURI:SuggestGasPrice", err.Error())
				return
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
				log.Println("submitURI:SignTx", err.Error())
				return
			}
			err = bs.Client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Println("submitURI:SendTransaction", err.Error())
				return
			}
			log.Println("submit uri tx sent: ", signedTx.Hash().Hex())*/
	}
}
func (bs *Service) getSdlByID(orderContractAddr string) []byte {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		log.Println("getSdlByID:HexToECDSA", err.Error())
		return nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("getSdlByID:publicKey is not of type *ecdsa.PublicKey", err.Error())
		return nil
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
func (bs *Service) getOwner(orderContractAddr string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		log.Println("getOwner:HexToECDSA", err.Error())
		return common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("getOwner:publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodSdlTrx := "owner"
	dataSdlTrx, err := bs.Abi[OrderBaseName].Pack(methodSdlTrx)
	if err != nil {
		log.Println("getOwner:Pack abi", err.Error())
		return common.Address{}, err
	}
	msgSdl := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataSdlTrx,
	}
	resultAddress, err := bs.Client.CallContract(context.Background(), msgSdl, nil)
	if err != nil {
		log.Println("getOwner:CallContract", err.Error())
		return common.Address{}, err
	}
	return common.BytesToAddress(resultAddress), nil
}
func (bs *Service) getOrderState(index uint64) (int64, string, error) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		return 0, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, "", errors.New("getOrderState : not ecdsa publicKey type")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(bs.Conf.OrderFactory)
	methodOrders := "orders"
	dataOrderTrx, err := bs.Abi[OrderFactoryName].Pack(methodOrders, new(big.Int).SetUint64(index))
	if err != nil {
		return 0, "", err
	}
	msgOrder := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataOrderTrx,
	}
	resultOrder, err := bs.Client.CallContract(context.Background(), msgOrder, nil)
	if err != nil {
		return 0, "", err
	}
	hexOrderAddr := common.BytesToAddress(resultOrder)
	if hexOrderAddr.String() == "0x0000000000000000000000000000000000000000" {
		return 0, hexOrderAddr.String(), errors.New(noOrderFound)
	}
	methodOrderState := "order_status"
	dataOrderStateTrx, _ := bs.Abi[OrderBaseName].Pack(methodOrderState)
	msgOrderState := ethereum.CallMsg{
		From: fromAddress,
		To:   &hexOrderAddr,
		Gas:  0,
		Data: dataOrderStateTrx,
	}
	resultOrderState, err := bs.Client.CallContract(context.Background(), msgOrderState, nil)
	if err != nil {
		return 0, "", err
	}
	state, err := strconv.ParseInt(common.Bytes2Hex(resultOrderState), 16, 64)
	if err != nil {
		return 0, "", err
	}
	return state, hexOrderAddr.String(), nil
}
func (bs *Service) getOrderIndex(orderContractAddr string) (uint64, error) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		return 0, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, errors.New("getOrderIndex:publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodSdlTrx := "o_order_number"
	dataSdlTrx, err := bs.Abi[OrderBaseName].Pack(methodSdlTrx)
	if err != nil {
		return 0, err
	}
	msgSdl := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataSdlTrx,
	}
	result, err := bs.Client.CallContract(context.Background(), msgSdl, nil)
	if err != nil {
		return 0, err
	}
	hex := common.Bytes2Hex(result)
	index, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return index, nil
}
func (bs *Service) getOrderCount(orderContractAddr string) (resourceStorage, error) {

	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		return resourceStorage{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return resourceStorage{}, errors.New("getOrderCount:publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	methodCPU := "o_cpu"
	methodMemory := "o_memory"
	methodStorage := "o_storage"
	dataCPU, err := bs.Abi[OrderBaseName].Pack(methodCPU)
	if err != nil {
		return resourceStorage{}, err
	}
	dataMemory, err := bs.Abi[OrderBaseName].Pack(methodMemory)
	if err != nil {
		return resourceStorage{}, err
	}
	dataStorage, err := bs.Abi[OrderBaseName].Pack(methodStorage)
	if err != nil {
		return resourceStorage{}, err
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
	resultCPU, err := bs.Client.CallContract(context.Background(), msgCPU, nil)
	if err != nil {
		return resourceStorage{}, err
	}
	resultMemory, err := bs.Client.CallContract(context.Background(), msgMemory, nil)
	if err != nil {
		return resourceStorage{}, err
	}
	resultStorage, err := bs.Client.CallContract(context.Background(), msgStorage, nil)
	if err != nil {
		return resourceStorage{}, err
	}
	var ret resourceStorage
	ret.CPUCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultCPU), 16)
	ret.MemoryCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultMemory), 16)
	ret.StorageCount, _ = new(big.Int).SetString(common.Bytes2Hex(resultStorage), 16)
	fmt.Println("result is", ret)
	return ret, nil
}
func (bs *Service) getOrderChosenProvider(orderContractAddr string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		return common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("getOrderChosenProvider:publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(orderContractAddr)
	methodProviderTrx := "query_provider_address"
	dataProviderTrx, err := bs.Abi[OrderBaseName].Pack(methodProviderTrx)
	if err != nil {
		return common.Address{}, err
	}
	msgProvider := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataProviderTrx,
	}
	resultAddress, err := bs.Client.CallContract(context.Background(), msgProvider, nil)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(resultAddress), nil
}
func (bs *Service) payBill(orderContractAddr string) {
	orderBase, err := util.NewOrderBase(common.HexToAddress(orderContractAddr), bs.Client)
	if err != nil {
		log.Println("payBill error create order base", err.Error())
		return
	}
	orderState, err := orderBase.OrderStatus(nil)
	if err != nil {
		log.Println("payBill error get order state", err.Error())
		return
	}
	if int64(orderState) == orderStatusRunning {
		payBillObj := util.PayBill{
			ProviderAddr: bs.Conf.ProviderContract,
			OrderAddr:    orderContractAddr,
		}
		privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
		if err != nil {
			log.Println("payBill:turn ecdsa", err.Error())
			return
		}
		payBillObj.Sign(privateKey)
		if payBillObj.SignMsg == "" {
			log.Println("payBill: Communication sign fail")
			return
		}
		success, trxID := payBillObj.Send(bs.Conf.PoolURI)
		if !success {
			log.Println("payBill: send trx to pool fail")
			return
		}
		log.Println("payBill tx sent: ", trxID)
		/*
			privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
			if err != nil {
				log.Println("payBill HexToECDSA", err.Error())
				return
			}
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Println("payBill publicKey is not of type *ecdsa.PublicKey")
				return
			}
			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Println("payBill PendingNonceAt", err.Error())
				return
			}
			value := big.NewInt(0)      // in wei (1 eth)
			gasLimit := uint64(3000000) // in units
			gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
			toAddress := common.HexToAddress(orderContractAddr)
			methodBillTrx := "pay_billing"
			dataPayBillTrx, err := bs.Abi[OrderBaseName].Pack(methodBillTrx)
			if err != nil {
				log.Println("payBill pack error", err.Error())
				return
			}
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
				log.Println("payBill SignTx error", err.Error())
				return
			}
			err = bs.Client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Println("payBill SendTransaction error", err.Error())
				return
			}
			log.Println("Pay Bill tx sent: ", signedTx.Hash().Hex())*/
	}

}
func (bs *Service) getOrderLastPayTime(orderContractAddr string) (int64, error) {
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		return 0, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, errors.New("publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	methodLastPayTime := "last_pay_time"

	dataPayTime, err := bs.Abi[OrderBaseName].Pack(methodLastPayTime)
	if err != nil {
		return 0, err
	}
	toAddress := common.HexToAddress(orderContractAddr)
	msgLastPayTime := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Gas:  0,
		Data: dataPayTime,
	}
	resultLastPayTime, err := bs.Client.CallContract(context.Background(), msgLastPayTime, nil)
	if err != nil {
		return 0, err
	}
	lastPayTimeHex := common.Bytes2Hex(resultLastPayTime)
	lastPayTime, err := strconv.ParseInt(lastPayTimeHex, 16, 64)
	if err != nil {
		return 0, err
	}
	return lastPayTime, nil
}
func (bs *Service) updateResource() {
	whetherChallenge, ok := bs.ChallengeLidsMap.Load(challengeState)
	if ok {
		if whetherChallenge.(bool) {
			log.Println("challenge in updateResource")
			return
		}
	}
	chainTotalLeft := bs.getTotalResource()
	totalTemp, err := bs.Cluster.GetTotalAvailable()
	if err != nil {
		log.Println("updateResource:cluster get error")
		return
	}
	orders, err := bs.getAllProviderServOrders()
	if err != nil {
		log.Println("updateResource:getAllProviderServOrders")
		return
	}
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	activeOrder := make(map[string]int)
	if len(orders) != 0 {
		leaseMap := make(map[ctypes.LeaseID]int, len(allActiveLeases))
		for _, lease := range allActiveLeases {
			leaseMap[lease] = 0
		}
		runningCount := 0
		allRunningCheck := false
		for _, order := range orders {
			if int64(order.State) == orderStatusRunning {
				runningCount++
				tempLease := ctypes.LeaseID{
					Owner:    order.Owner.String(),
					OSeq:     order.OrderId.Uint64(),
					Provider: common.HexToAddress(bs.Conf.ProviderContract).String(),
				}
				if _, ok := leaseMap[tempLease]; !ok {
					return
				}
				activeOrder[strings.ToLower(order.ContractAddress.String())] = 1
				leaseMap[tempLease] = 1

			}
		}
		if runningCount != len(allActiveLeases) {
			return
		}
		allRunningCheck = true
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
		//providerContractAddr := common.HexToAddress(bs.Conf.ProviderContract)
		updateResourceObj := util.UpdateResource{
			ProviderAddr: bs.Conf.ProviderContract,
			Cpu:          new(big.Int).SetUint64(totalTemp.CPU).String(),
			Memory:       new(big.Int).SetUint64(totalTemp.Memory).String(),
			Storage:      new(big.Int).SetUint64(totalTemp.StorageEphemeral).String(),
		}
		privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
		if err != nil {
			log.Println("UpdateResource:turn ecdsa", err.Error())
			return
		}
		updateResourceObj.Sign(privateKey)
		if updateResourceObj.SignMsg == "" {
			log.Println("UpdateResource: Communication sign fail")
			return
		}
		success, trxID := updateResourceObj.Send(bs.Conf.PoolURI)
		if !success {
			log.Println("UpdateResource: send trx to pool fail", trxID)
			return
		}
		log.Println("UpdateResource tx sent: ", trxID)
		/*
			privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
			if err != nil {
				log.Println("updateResource:HexToECDSA", err.Error())
				return
			}
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Println("updateResource:publicKey is not of type *ecdsa.PublicKey", err.Error())
				return
			}
			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			method := "updateResource"
			data, err := bs.Abi[ProviderName].Pack(method, new(big.Int).SetUint64(totalTemp.CPU), new(big.Int).SetUint64(totalTemp.Memory), new(big.Int).SetUint64(totalTemp.StorageEphemeral))
			if err != nil {
				fmt.Println("err is ", err.Error())
			}
			nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Println("updateResource:PendingNonceAt", err.Error())
				return
			}
			value := big.NewInt(0)      // in wei (1 eth)
			gasLimit := uint64(3000000) // in units
			gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Println("updateResource:SuggestGasPrice", err.Error())
				return
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
				log.Println("updateResource:SignTx", err.Error())
				return
			}
			err = bs.Client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Println("updateResource:SendTransaction", err.Error())
				return
			}
			log.Println("update source tx sent: ", signedTx.Hash().Hex())
		*/
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
			if _, ok := activeOrder[strings.ToLower(key.(string))]; !ok {
				contractTempAddrs = append(contractTempAddrs, key.(string))
			}
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
func (bs *Service) getAllProviderServOrders() ([]util.Order, error) {
	orderBase, err := util.NewOrderFactory(common.HexToAddress(util.GetOrderFactory(bs.Conf)), bs.Client)
	if err != nil {
		return nil, err
	}
	allProviderOrder, err := orderBase.GetProviderAllOrder(nil, common.HexToAddress(bs.Conf.ProviderContract))
	if err != nil {
		return nil, err
	}
	return allProviderOrder, nil
}

type getSeedStruct struct {
	Md5Seed      *big.Int `json:"md5Seed"`
	ProviderAddr string   `json:"providerAddr"`
	PorCount     uint64   `json:"porCount"`
}
type seedResponse struct {
	Seed   uint64 `json:"seed"`
	TaskID int64  `json:"task_id"`
	Find   bool   `json:"find"`
}

func (bs *Service) calcPORCount() uint64 {
	providerFactory, err := util.NewProviderFactory(common.HexToAddress(bs.Conf.ProviderFactoryContract), bs.Client)
	if err != nil {
		log.Println("calcPORCount error", err.Error())
		return 0
	}
	decimalCPU, err := providerFactory.DecimalCpu(nil)
	if err != nil {
		log.Println("get DecimalCpu error", err.Error())
		return 0
	}
	decimalMemory, err := providerFactory.DecimalMemory(nil)
	if err != nil {
		log.Println("get decimalMemory error", err.Error())
		return 0
	}
	nodeResources, err := bs.Cluster.GetNodesAvailable()
	if err != nil {
		log.Println("GetNodesAvailable error", err.Error())
		return 0
	}
	allCount := uint64(0)
	for _, nodeResource := range nodeResources {
		cpuCount := nodeResource.CPU / decimalCPU.Uint64()
		memCount := nodeResource.Memory / decimalMemory.Uint64()
		fmt.Println(cpuCount, memCount, bs.Total.MemoryCount.Int64(), decimalCPU, decimalMemory)
		if cpuCount > memCount {
			cpuCount = memCount
		}
		allCount = allCount + cpuCount
	}

	return allCount
}
func (bs *Service) getSeedFromValidatorMidWare(md5Seed *big.Int, providerAddr string, validatorURL string) (*seedResponse, uint64) {
	porCount := bs.calcPORCount()
	requestFunction := func(porCount uint64) (*seedResponse, error) {
		method := "POST"
		client := &http.Client{Timeout: 5 * time.Second}

		seedJSONTemp := getSeedStruct{
			Md5Seed:      md5Seed,
			ProviderAddr: providerAddr,
			PorCount:     porCount,
		}
		log.Println(seedJSONTemp)
		reqBody, err := json.Marshal(seedJSONTemp)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(method, validatorURL+"/get_seed", strings.NewReader(string(reqBody)))
		if err != nil {
			return nil, err
		}
		res, err := client.Do(req)

		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(res.Body)
		if err != nil {
			fmt.Println("read all error")
			return nil, err
		}
		var response seedResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("enter Unmarshal error")
			return nil, err
		}
		if !response.Find {
			return nil, errors.New("Not find seed")
		}
		return &response, nil
	}
	for i := 0; i < 100; i++ {
		unHashSeed, err := requestFunction(porCount)
		if err != nil {
			time.Sleep(2 * time.Second)
			log.Println("get seed error ", err.Error())
			continue
		}
		return unHashSeed, porCount
	}
	return nil, porCount
}
func (bs *Service) getChallengeSdl() []byte {
	//path := "./sdl.txt"
	//buf, err := os.ReadFile(path)
	//if err != nil {
	//	fmt.Println(err.Error)
	//	return nil
	//}
	//return buf
	valFactory, err := util.NewValidatorFactory(common.HexToAddress(bs.Conf.ValidatorFactoryContract), bs.Client)
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return nil
	}
	sdlTrxID, err := valFactory.ChallengeSdlTrxId(nil)
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return nil
	}
	sdlTrx, pending, err := bs.Client.TransactionByHash(context.Background(), common.HexToHash(sdlTrxID.Text(16)))
	if pending {
		log.Println("getChallengeInfo transaction is pending")
		return nil
	}
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return nil
	}
	return sdlTrx.Data()
}
func (bs *Service) getChallengeInfo(providerOwner string) *util.ValidatorFactoryproviderChallengeInfo {
	valFactory, err := util.NewValidatorFactory(common.HexToAddress(bs.Conf.ValidatorFactoryContract), bs.Client)
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return nil
	}
	info, err := valFactory.GetProviderChallengeInfo(nil, common.HexToAddress(providerOwner))
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return nil
	}
	return &info
}
func (bs *Service) getChallengeTimeout() int64 {
	valFactory, err := util.NewValidatorFactory(common.HexToAddress(bs.Conf.ValidatorFactoryContract), bs.Client)
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return 0
	}
	maxChallengeTime, err := valFactory.MaxChallengeTime(nil)
	if err != nil {
		log.Println("getChallengeInfo error", err.Error())
		return 0
	}
	return maxChallengeTime.Int64()
}
func (bs *Service) endChallenge() {
	submitObj := util.EndChallenge{
		ProviderAddr: bs.Conf.ProviderContract,
	}
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	if err != nil {
		log.Println("EndChallenge:turn ecdsa", err.Error())
		return
	}
	submitObj.Sign(privateKey)
	if submitObj.SignMsg == "" {
		log.Println("EndChallenge: Communication sign fail")
		return
	}
	success, trxID := submitObj.Send(bs.Conf.PoolURI)
	if !success {
		log.Println("EndChallenge: send trx to pool fail")
		return
	}
	log.Println("EndChallenge tx sent: ", trxID)
	/*
		privateKey, err := crypto.HexToECDSA(bs.Conf.SecretKey)
		if err != nil {
			log.Println("endChallenge HexToECDSA")
			return
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Println("endChallenge publicKey is not of type *ecdsa.PublicKey")
			return
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := bs.Client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Println("endChallenge PendingNonceAt")
			return
		}
		value := big.NewInt(0)      // in wei (1 eth)
		gasLimit := uint64(3000000) // in units
		gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
		method := "validatorNotSubmitResult"
		data, err := bs.Abi[ValidatorFactoryName].Pack(method, common.HexToAddress(bs.Conf.ProviderAddress))
		if err != nil {
			log.Println("endChallenge Pack")
			return
		}
		valFactoryContract := common.HexToAddress(bs.Conf.ValidatorFactoryContract)
		msg := &types.LegacyTx{
			Nonce:    nonce,
			To:       &valFactoryContract,
			Gas:      gasLimit,
			Value:    value,
			GasPrice: gasPrice,
			Data:     data,
		}
		tx := types.NewTx(msg)
		chainID, _ := new(big.Int).SetString(bs.Conf.NodeChainID, 10)
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Println("endChallenge SignTx")
			return
		}
		err = bs.Client.SendTransaction(context.Background(), signedTx)
		log.Println("ValidatorNotSubmitResult tx sent: ", signedTx.Hash().Hex())

	*/
}
func (bs *Service) ChangeProviderInfo() {
	provider, err := util.NewProvider(common.HexToAddress(bs.Conf.ProviderContract), bs.Client)
	if err != nil {
		log.Println("changeProviderInfo error", err.Error())
		return
	}
	privateKey, err := crypto.HexToECDSA(bs.Conf.CommunicationPrivateKey)
	nodeId, _ := strconv.ParseInt(bs.Conf.NodeChainID, 10, 64)
	fmt.Println(nodeId)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(nodeId))
	if err != nil {
		log.Println("changeProviderInfo error1", err.Error())
		return
	}
	auth.GasLimit = uint64(3000000)
	gasPrice, err := bs.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("changeProviderInfo error1", err.Error())
		return
	}
	fmt.Println("changeProviderInfo 4")
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
	auth.GasPrice = gasPrice
	auth.Nonce = new(big.Int).SetUint64(nonce) // the account nonce for the transaction
	tx, err := provider.ChangeProviderInfo(auth, "{\"country\":\"CN\",\"email\":\"365mad@qq.com\",\"website\":\"www.google.com\",\"kubeVersion\":\"1.0\",\"platform\":\"linux/amd64\",\"uri\":\"https://provider.ubicloud.matrixlabs.org:8443\",\"attributes\":{\"region\":\"wuxi\",\"chiaPlotting\":\"false\",\"host\":\"192.168.1.1\",\"cpu\":\"intel\"}}")

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(nodeId)), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = bs.Client.SendTransaction(context.Background(), signedTx)
	fmt.Println("ValidatorNotSubmitResult tx sent: ", signedTx.Hash().Hex())

}

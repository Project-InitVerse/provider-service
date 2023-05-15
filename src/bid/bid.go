package bid

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ovrclk/akash/sdl"
	"log"
	"math/big"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/config"
	"providerService/src/util"
	ubic_cluster "providerService/ubic-cluster"
	"strings"
	"sync"
	"time"
)

var (
	orderStatusQuoting int64 = 1
	orderStatusRunning int64 = 2
	orderStatusEnded   int64 = 3
	orderPayInt        int64 = 3600
)

type resourceStorage struct {
	CPUCount     *big.Int
	MemoryCount  *big.Int
	StorageCount *big.Int
}

// Service is bid service
type Service struct {
	BidChan             <-chan util.NeedBid
	BidFinalChan        <-chan util.NeedCreate
	OrderFinish         <-chan util.UserCancelOrder
	BidTimeout          <-chan time.Time
	SubMitTimeOut       <-chan time.Time
	UpdateSourceTimeout <-chan time.Time
	Total               resourceStorage
	Client              *ethclient.Client
	Conf                *config.ProviderConfig
	Cluster             *ubic_cluster.UbicService
	KeepResource        sync.Map
	KeepResourceTime    sync.Map
	MutexRw             *sync.RWMutex
	Abi                 map[string]abi.ABI
	Ctx                 context.Context
	WgBid               *sync.WaitGroup
	LastPayTime         int64
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
	bs.BidTimeout = time.After(30 * time.Second)
	bs.SubMitTimeOut = time.After(10 * time.Second)
	bs.Ctx = ctx
	bs.WgBid = new(sync.WaitGroup)
	bs.LastPayTime = time.Now().Unix()
	bs.initExistDeployment()
	bs.UpdateSourceTimeout = time.After(30 * time.Second)
}

func (bs *Service) initExistDeployment() {
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	orderInChain := bs.getAllProviderServOrders()
	leaseLocal := make(map[string]int, 0)
	for _, lease := range allActiveLeases {
		state, orderAddr := bs.getOrderState(lease.OSeq)
		fmt.Println("initExistDeployment", lease, state, orderAddr)
		leaseLocal[strings.ToLower(orderAddr)] = 1
		if state == orderStatusRunning {
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
				Owner:    owner.String(),
				OSeq:     index,
				Provider: bs.getProviderAddr(),
			}
			uri, _ := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
			bs.submitURI(orderSingle.ContractAddress.String(), uri)
		}
	}
}
func (bs *Service) validateSdlFile(orderContract string, _CPU, _MEM, _STO *big.Int) bool {
	sdlFile := bs.getSdlByID(orderContract)
	sdlObj, err := sdl.Read(sdlFile)
	if err != nil {
		log.Println("sdl struct error", err.Error())
		return false
	}
	deployments, err := sdlObj.DeploymentGroups()
	if err != nil {
		log.Println("get deployment Group error", err.Error())
		return false
	}
	if len(deployments) != 1 {
		log.Println("deployment length error", len(deployments))
		return false
	}

	sdlCPU := new(big.Int)
	sdlMem := new(big.Int)
	sdlSto := new(big.Int)
	resources := deployments[0].GetResources()
	for _, resource := range resources {
		cpu := new(big.Int).SetUint64(resource.Resources.CPU.Units.Value())
		cpu = cpu.Mul(cpu, new(big.Int).SetInt64(int64(resource.Count)))
		sdlCPU = sdlCPU.Add(sdlCPU, cpu)
		mem := new(big.Int).SetUint64(resource.Resources.Memory.Quantity.Value())
		mem = cpu.Mul(mem, new(big.Int).SetInt64(int64(resource.Count)))
		sdlMem = sdlMem.Add(sdlMem, mem)
		sto := new(big.Int)
		for _, v := range resource.Resources.Storage {
			sto = sto.Add(sto, new(big.Int).SetUint64(v.Quantity.Value()))
		}
		sto = sto.Mul(sto, new(big.Int).SetInt64(int64(resource.Count)))
		sdlSto = sdlSto.Add(sdlSto, sto)
	}
	if _CPU.Cmp(sdlCPU) != 0 || _MEM.Cmp(sdlMem) != 0 || _STO.Cmp(sdlSto) != 0 {
		return false
	}
	return true
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
	if !bs.validateSdlFile(orderInfo.ContractAddress, orderInfo.CPU, orderInfo.Memory, orderInfo.Storage) {
		log.Println("validate sdl fail")
		return
	}
	fmt.Println(bs.Total.CPUCount, bs.Total.MemoryCount, bs.Total.StorageCount, orderInfo)
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
	if strings.ToLower(bidFinalInfo.Provider.String()) != strings.ToLower(bs.Conf.ProviderContract) {
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
			Owner:    owner.String(),
			OSeq:     index,
			Provider: bidFinalInfo.Provider.String(),
		}
		uri, _ := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
		bs.submitURI(bidFinalInfo.ContractAddress, uri)
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
		Owner:    owner.String(),
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
func (bs *Service) handleSubmitURI() {
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	for _, lease := range allActiveLeases {
		_, orderAddr := bs.getOrderState(lease.OSeq)
		uri := bs.Cluster.GetURI(lease)
		bs.submitURI(orderAddr, uri)
	}
}
func (bs *Service) handleRecoverResource() {
	log.Println("in handleRecoverResource")
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
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	for _, lease := range allActiveLeases {
		fmt.Println(lease)
		_, orderAddr := bs.getOrderState(lease.OSeq)
		lastTime := bs.getOrderLastPayTime(orderAddr)
		fmt.Println("last time is ", lastTime)
		if lastTime+orderPayInt < time.Now().Unix() {
			bs.payBill(orderAddr)
		}
	}
	bs.LastPayTime = time.Now().Unix()
}
func (bs *Service) HandleChallenge(challenge *util.NeedChallenge) {

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
		case <-bs.SubMitTimeOut:
			go bs.handleSubmitURI()
			bs.SubMitTimeOut = time.After(30 * time.Second)
		case <-bs.UpdateSourceTimeout:
			go bs.updateResource()
			bs.UpdateSourceTimeout = time.After(6 * time.Hour)
		case <-bs.Ctx.Done():
			bs.WgBid.Wait()
			log.Println("bid service exit")
			break loop
		}
	}
}

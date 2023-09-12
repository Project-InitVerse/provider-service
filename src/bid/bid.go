package bid

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
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
	orderStatusQuoting   int64  = 1
	orderStatusRunning   int64  = 2
	orderStatusEnded     int64  = 3
	orderPayInt          int64  = 3600
	challengeLids        string = "challengeLids"
	challengeState       string = "challengeState"
	challengeProvider    string = "challengeProvider"
	challengeCreateState uint64 = 1
	zeroAddr             string = "0x0000000000000000000000000000000000000000"
	noOrderFound         string = "no orders found"
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
	BidChallenge        <-chan util.NeedChallenge
	ChallengeFinal      <-chan util.ChallengeEnd
	BidTimeout          <-chan time.Time
	SubMitTimeOut       <-chan time.Time
	UpdateSourceTimeout <-chan time.Time
	EndChallengeTimeout <-chan time.Time
	Total               resourceStorage
	Client              *ethclient.Client
	Conf                *config.ProviderConfig
	Cluster             *ubic_cluster.UbicService
	KeepResource        sync.Map
	KeepResourceTime    sync.Map
	ChallengeLidsMap    sync.Map
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
	bidChallenge <-chan util.NeedChallenge,
	challengeFinal <-chan util.ChallengeEnd,
	cluster *ubic_cluster.UbicService) {
	bs.Client, _ = ethclient.Dial(config.NodeURL)
	bs.Conf = config
	bs.BidFinalChan = bidFinal
	bs.BidChan = bidChan
	bs.BidChallenge = bidChallenge
	bs.ChallengeFinal = challengeFinal
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
	bs.UpdateSourceTimeout = time.After(10 * time.Second)

}

func (bs *Service) initExistDeployment() {
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	orderInChain, err := bs.getAllProviderServOrders()
	if err != nil {
		log.Fatal("initExistDeployment getAllProviderServOrders error", err.Error())
	}
	leaseLocal := make(map[string]int, 0)
	leaseChallenge := make([]ctypes.LeaseID, 0)
	for _, lease := range allActiveLeases {
		if lease.Provider == challengeProvider {
			leaseChallenge = append(leaseChallenge, lease)
		} else {
			state, orderAddr, err := bs.getOrderState(lease.OSeq)
			if err != nil {
				log.Println("initExistDeployment get order state error", err.Error())
				if orderAddr == zeroAddr && err.Error() == noOrderFound {
					bs.Cluster.CloseManager(lease)
				}
				continue
			}
			log.Println("initExistDeployment", lease, state, orderAddr)
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

	}
	if len(leaseChallenge) > 0 {
		bs.ChallengeLidsMap.Store(challengeLids, leaseChallenge)
		bs.ChallengeLidsMap.Store(challengeState, true)
		timeout := bs.getChallengeTimeout()
		bs.EndChallengeTimeout = time.After(time.Duration(timeout) * time.Second)
	} else {
		bs.ChallengeLidsMap.Store(challengeState, false)
	}

	for _, orderSingle := range orderInChain {
		if _, ok := leaseLocal[strings.ToLower(orderSingle.ContractAddress.String())]; ok {
			continue
		}
		if int64(orderSingle.State) == orderStatusRunning {
			sdlFile := bs.getSdlByID(orderSingle.ContractAddress.String())
			owner, err := bs.getOwner(orderSingle.ContractAddress.String())
			if err != nil {
				log.Println("initExistDeployment:get order owner error", err.Error())
				continue
			}
			index, err := bs.getOrderIndex(orderSingle.ContractAddress.String())
			if err != nil {
				log.Println("initExistDeployment:getOrderIndex error", err.Error())
				continue
			}
			lid := ctypes.LeaseID{
				Owner:    owner.String(),
				OSeq:     index,
				Provider: bs.getProviderAddr(),
			}
			uri, err := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
			if err != nil {
				log.Println("initExistDeployment:create deployment error", err.Error())
				continue
			}
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
	whetherChallenge, ok := bs.ChallengeLidsMap.Load(challengeState)
	if ok {
		if whetherChallenge.(bool) {
			log.Println("challenge in handle bid")
			return
		}
	}
	log.Println(orderInfo.ContractAddress)
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	_, ok = bs.KeepResource.Load(strings.ToLower(orderInfo.ContractAddress))
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
	whetherChallenge, ok := bs.ChallengeLidsMap.Load(challengeState)
	if ok {
		if whetherChallenge.(bool) {
			log.Println("challenge in handle bid final")
			return
		}
	}
	bs.WgBid.Add(1)
	defer bs.WgBid.Done()
	privateKey, _ := crypto.HexToECDSA(bs.Conf.SecretKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
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
			orderSource, err := bs.getOrderCount(bidFinalInfo.ContractAddress)
			if err != nil {
				log.Println("handleBidFinal:getOrderCount error", err.Error())
				return
			}
			log.Println("handleBidFinal source", bs.Total, orderSource)
			if bs.Total.CPUCount.Cmp(orderSource.CPUCount) >= 0 &&
				bs.Total.MemoryCount.Cmp(orderSource.MemoryCount) >= 0 &&
				bs.Total.StorageCount.Cmp(orderSource.StorageCount) >= 0 {
				bs.KeepResource.Store(strings.ToLower(bidFinalInfo.ContractAddress), orderSource)
				bs.handleResource(orderSource, false)
			} else {
				log.Println("handleBidFinal:left resource not enough")
				return
			}
		}
		sdlFile := bs.getSdlByID(bidFinalInfo.ContractAddress)
		owner, err := bs.getOwner(bidFinalInfo.ContractAddress)
		if err != nil {
			log.Println("handleBidFinal:get order owner error", err.Error())
			return
		}
		index, err := bs.getOrderIndex(bidFinalInfo.ContractAddress)
		if err != nil {
			log.Println("handleBidFinal:getOrderIndex error", err.Error())
			return
		}
		log.Println(owner, index, bidFinalInfo.Provider.String())
		lid := ctypes.LeaseID{
			Owner:    owner.String(),
			OSeq:     index,
			Provider: bidFinalInfo.Provider.String(),
		}
		uri, err := bs.Cluster.NewUbicDeployManager(lid, sdlFile)
		if err != nil {
			log.Println("handleBidFinal:create deployment error", err.Error())
			return
		}
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
	owner, err := bs.getOwner(orderFinishInfo.ContractAddress.String())
	if err != nil {
		log.Println("handleOrderFinish:get order owner error", err.Error())
		return
	}
	index, err := bs.getOrderIndex(orderFinishInfo.ContractAddress.String())
	if err != nil {
		log.Println("handleOrderFinish:getOrderIndex error", err.Error())
		return
	}
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
	//log.Println("in handle submit URL")
	//bs.changeProviderInfo()
	allActiveLeases := bs.Cluster.GetAllActiveLeases()
	for _, lease := range allActiveLeases {
		if lease.Provider != challengeProvider {
			_, orderAddr, err := bs.getOrderState(lease.OSeq)
			if err != nil {
				log.Println("handleSubmitURI: getOrderState error,", err.Error())
				continue
			}
			uri := bs.Cluster.GetURI(lease)
			bs.submitURI(orderAddr, uri)
		}
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
		log.Println(lease)
		if lease.Provider != challengeProvider {
			_, orderAddr, err := bs.getOrderState(lease.OSeq)
			if err != nil {
				log.Println("handleRecoverResource: getOrderState error,", err.Error())
				continue
			}
			providerAddr, err := bs.getOrderChosenProvider(orderAddr)
			if err != nil {
				log.Println("handleRecoverResource getOrderChosenProvider", err.Error())
				continue
			}
			if providerAddr.String() != bs.Conf.ProviderContract {
				log.Println("not my deployment")
				bs.Cluster.CloseManager(lease)
				continue
			}
			lastTime, err := bs.getOrderLastPayTime(orderAddr)
			if err != nil {
				log.Println("handleRecoverResource getOrderChosenProvider", err.Error())
				continue
			}
			log.Println("last time is ", lastTime)
			if lastTime+orderPayInt < time.Now().Unix() {
				bs.payBill(orderAddr)
			}
		}
	}
	bs.LastPayTime = time.Now().Unix()
}
func (bs *Service) handleChallenge(challenge *util.NeedChallenge) {
	log.Println("in handle challenge")
	if challenge.Owner.String() != bs.Conf.ProviderAddress {
		log.Println("not my challenge")
		return
	}
	whetherChallenge, ok := bs.ChallengeLidsMap.Load(challengeState)
	if ok {
		if whetherChallenge.(bool) {
			log.Println("has in challenge state")
			return
		}
	}
	info := bs.getChallengeInfo(challenge.Owner.String())
	if info != nil {
		if info.Index.Cmp(challenge.Index) != 0 {
			log.Println("not this turn challenge")
			return
		}
		if info.Url == "" || uint64(info.State) != challengeCreateState {
			log.Println("no info in challenge")
			return
		}
		//verify seed
		bs.ChallengeLidsMap.Store(challengeState, true)
		seedResponse, porCount := bs.getSeedFromValidatorMidWare(info.Md5Seed, info.Provider.String(), info.Url)
		if seedResponse == nil {
			log.Println("connect mid ware fail")
			bs.ChallengeLidsMap.Store(challengeState, false)
			return
		}
		seedBytes := [8]byte{}
		binary.LittleEndian.PutUint64(seedBytes[:], seedResponse.Seed)
		seedHash := crypto.Keccak256(seedBytes[:])

		if new(big.Int).SetBytes(seedHash).Cmp(info.Md5Seed) != 0 {
			log.Println("seed was not correct")
			bs.ChallengeLidsMap.Store(challengeState, false)
			return
		}

		challengeSdl := bs.getChallengeSdl()
		lids, err := bs.Cluster.NewChallengeDeployManager(challengeSdl, porCount, seedResponse.Seed, seedResponse.TaskID, info.Url)
		if err != nil {
			log.Println("create challengeDeploy fail")
			bs.ChallengeLidsMap.Store(challengeState, false)
			return
		}

		log.Println("store lids: ", lids)
		bs.ChallengeLidsMap.Store(challengeLids, lids)
		timeout := bs.getChallengeTimeout()
		bs.EndChallengeTimeout = time.After(time.Duration(timeout) * time.Second)
	}
}
func (bs *Service) handleChallengeEnd(challenge *util.ChallengeEnd) {
	log.Println("in handle challenge end")
	if challenge.Owner.String() != bs.Conf.ProviderAddress {
		log.Println("not my challenge")
		return
	}
	info := bs.getChallengeInfo(challenge.Owner.String())
	if info != nil {
		if info.Index.Cmp(challenge.Index) != 0 {
			log.Println("not this turn challenge end")
			return
		}
		lids, ok := bs.ChallengeLidsMap.Load(challengeLids)
		log.Println("load lids: ", lids)
		if ok {
			for _, lid := range lids.([]ctypes.LeaseID) {
				bs.Cluster.CloseManager(lid)
			}
			bs.ChallengeLidsMap.Delete(challengeLids)
		}
		bs.ChallengeLidsMap.Store(challengeState, false)
	}
}
func (bs *Service) endChallengeTime() {
	log.Println("in handle end challenge time")
	info := bs.getChallengeInfo(bs.Conf.ProviderAddress)
	if info != nil {
		//if uint64(info.State) == challengeCreateState {
		lids, ok := bs.ChallengeLidsMap.Load(challengeLids)
		if ok {
			for _, lid := range lids.([]ctypes.LeaseID) {
				bs.Cluster.CloseManager(lid)
			}
			bs.ChallengeLidsMap.Delete(challengeLids)
		}
		bs.ChallengeLidsMap.Store("challenge", false)
		bs.endChallenge()
		//}

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
		case <-bs.SubMitTimeOut:
			go bs.handleSubmitURI()
			bs.SubMitTimeOut = time.After(30 * time.Second)
		case <-bs.UpdateSourceTimeout:
			go bs.updateResource()
			bs.UpdateSourceTimeout = time.After(300 * time.Second)
		case challengeFinal := <-bs.ChallengeFinal:
			go bs.handleChallengeEnd(&challengeFinal)
		case challenge := <-bs.BidChallenge:
			go bs.handleChallenge(&challenge)
		case <-bs.EndChallengeTimeout:
			go bs.endChallengeTime()
		case <-bs.Ctx.Done():
			bs.WgBid.Wait()
			log.Println("bid service exit")
			break loop
		}
	}
}

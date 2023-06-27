package ubiccluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ovrclk/akash/sdl"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	"providerService/src/cluster"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/cluster/ubickube"
	cmdutil "providerService/src/cmd/provider-services/cmd/util"
	"providerService/src/config"
	"strconv"
	"sync"
)

// UbicService is struct
type UbicService struct {
	UbicKubeClient cluster.UbicClient
	Managers       map[ctypes.LeaseID]*cluster.UbicDeploymentManager
	UbicLog        log.Logger
	Config         cluster.Config
	ConfBase       config.ProviderConfig
	Hostnames      *cluster.UbicHostnameService
}

// NewService create ubic new service
func (us *UbicService) NewService(ctx context.Context, conf config.ProviderConfig, clusterConf cluster.Config) {
	us.ConfBase = conf
	us.Config = clusterConf
	logger := cmdutil.OpenLogger().With("cmp", "provider")
	us.UbicLog = logger
	cclient, err := kube.NewClient(context.Background(), logger, conf.NameSpace, conf.K8sConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	allHostnames, _ := cclient.AllHostnames(ctx)
	activeHostnames := make(map[string]ctypes.LeaseID, len(allHostnames))
	fmt.Println("allHostnames", allHostnames)
	for _, v := range allHostnames {
		activeHostnames[v.Hostname] = v.ID
	}
	hostnames, _ := cluster.NewUbicHostnameService(ctx, us.Config, activeHostnames)
	us.Hostnames = hostnames
	us.UbicKubeClient = cclient
	us.Managers = make(map[ctypes.LeaseID]*cluster.UbicDeploymentManager, 0)
	us.LoadExistDeployManager()
	fmt.Println(len(us.Managers))
}

// LoadExistDeployManager load deployment from k8s
func (us *UbicService) LoadExistDeployManager() {
	deployments, err := (us.UbicKubeClient).Deployments(context.Background())
	if err != nil {
		return
	}

	for _, dm := range deployments {
		group := dm.ManifestGroup()
		us.Managers[dm.LeaseID()] = cluster.NewUbicDeploymentManager(context.Background(),
			us.UbicKubeClient,
			dm.LeaseID(),
			us.UbicLog,
			&group,
			us.Hostnames,
			us.Config,
			false)
	}

}

// NewChallengeDeployManager create new deployment for POR
func (us *UbicService) NewChallengeDeployManager(
	sdlSteam []byte,
	challengeCount uint64,
	seed uint64, _taskID int64, commitURL string) ([]ctypes.LeaseID, error) {
	sdlFile, err := sdl.Read(sdlSteam)
	if err != nil {
		fmt.Println("exit in sdl conv", err.Error())
		return nil, err
	}
	groups, _ := sdlFile.Manifest()
	if len(groups) != 1 {
		return nil, errors.New("NewChallengeDeployManager:group not support over 1")
	}
	ret := make([]ctypes.LeaseID, 0)
	for i := uint64(0); i < challengeCount; i++ {
		groupsIn, _ := sdlFile.Manifest()
		lidChallenge := ctypes.LeaseID{
			Owner:    "Owner",
			OSeq:     i,
			Provider: "challengeProvider",
		}
		for _, group := range groupsIn {
			for key := range group.Services {
				fmt.Println("enter lease id ", strconv.FormatUint(seed+uint64(i), 10))
				porSeed := "por_seed=" + strconv.FormatUint(seed+uint64(i), 10)
				commit := "commit_url=" + commitURL + ""
				taskID := "task_id=" + strconv.FormatInt(_taskID, 10)
				group.Services[key].Env = append(group.Services[key].Env, porSeed)
				group.Services[key].Env = append(group.Services[key].Env, commit)
				group.Services[key].Env = append(group.Services[key].Env, taskID)
				fmt.Println(group.Services)
			}

			ret = append(ret, lidChallenge)
			deployManager := cluster.NewUbicDeploymentManager(context.Background(),
				us.UbicKubeClient,
				lidChallenge,
				us.UbicLog,
				&group,
				us.Hostnames,
				us.Config,
				true)
			us.Managers[lidChallenge] = deployManager
			/*
				ret := make(map[string]interface{}, 0)
				exist, groups, _ := us.UbicKubeClient.GetManifestGroup(context.Background(), lidChallenge)
				if exist {
					for _, service := range groups.Services {
						s, _ := us.UbicKubeClient.ServiceStatus(context.Background(), lidChallenge, service.Name)
						ret[service.Name] = s.URIs
					}
				}
				result, err := json.Marshal(ret)
				if err != nil {
					fmt.Println(err.Error())
					return "", nil
				}
				return string(result), nil*/
		}
	}
	return ret, nil
}

// NewUbicDeployManager create new deployment
func (us *UbicService) NewUbicDeployManager(lid ctypes.LeaseID, sdlSteam []byte) (string, error) {
	sdlFile, err := sdl.Read(sdlSteam)
	if err != nil {
		fmt.Println("NewUbicDeployManager:exit in sdl conv", err.Error())
		return "", err
	}
	groups, _ := sdlFile.Manifest()
	if len(groups) != 1 {
		return "", errors.New("not support more group")
	}
	for _, group := range groups {
		deployManager := cluster.NewUbicDeploymentManager(context.Background(),
			us.UbicKubeClient,
			lid,
			us.UbicLog,
			&group,
			us.Hostnames,
			us.Config,
			true)
		us.Managers[lid] = deployManager
		ret := make(map[string]interface{}, 0)
		exist, groups, _ := us.UbicKubeClient.GetManifestGroup(context.Background(), lid)
		if exist {
			for _, service := range groups.Services {
				s, _ := us.UbicKubeClient.ServiceStatus(context.Background(), lid, service.Name)
				ret[service.Name] = s.URIs
			}
		}
		result, err := json.Marshal(ret)
		if err != nil {
			fmt.Println(err.Error())
			return "", nil
		}
		return string(result), nil
	}
	return "", nil
}

// GetURI get uri from leaseid
func (us *UbicService) GetURI(id ctypes.LeaseID) string {
	ret := make(map[string]interface{}, 0)
	exist, groups, _ := us.UbicKubeClient.GetManifestGroup(context.Background(), id)
	if exist {
		for _, service := range groups.Services {
			s, _ := us.UbicKubeClient.ServiceStatus(context.Background(), id, service.Name)
			ret[service.Name] = s.URIs
		}
	}
	result, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println(string(result))
	return string(result)
}

// IsActive create deployment active
func (us *UbicService) IsActive(leaseID ctypes.LeaseID) bool {
	_, ok := us.Managers[leaseID]
	return ok
}

// GetAllActiveLeases is function
func (us *UbicService) GetAllActiveLeases() []ctypes.LeaseID {
	ret := make([]ctypes.LeaseID, 0)
	for k := range us.Managers {
		ret = append(ret, k)
		//us.CloseManager(k)
	}
	return ret
}

// CloseAllLease is function close all exist lease for test
func (us *UbicService) CloseAllLease() {
	for k := range us.Managers {
		us.CloseManager(k)
	}
}

// GetTotalResource get k8s total resource
func (us *UbicService) GetTotalResource() (ctypes.InventoryMetricTotal, error) {
	in, err := us.UbicKubeClient.Inventory(context.Background())
	if err != nil {
		fmt.Println("inventory error ", err.Error())
		return ctypes.InventoryMetricTotal{}, err
	}
	return in.Metrics().TotalAllocatable, nil
}

// GetTotalAvailable get k8s available resource
func (us *UbicService) GetTotalAvailable() (ctypes.InventoryMetricTotal, error) {
	in, err := us.UbicKubeClient.Inventory(context.Background())
	if err != nil {
		fmt.Println("inventory error ", err.Error())
		return ctypes.InventoryMetricTotal{}, err
	}
	return in.Metrics().TotalAvailable, nil
}
func (us *UbicService) GetNodesAvailable() ([]ctypes.InventoryNodeMetric, error) {
	in, err := us.UbicKubeClient.Inventory(context.Background())
	if err != nil {
		fmt.Println("inventory error ", err.Error())
		return nil, err
	}
	ret := make([]ctypes.InventoryNodeMetric, 0)
	for _, v := range in.Metrics().Nodes {
		ret = append(ret, v.Available)
	}
	return ret, nil
}

// GetStatus get k8s status
func (us *UbicService) GetStatus() {
	in, err := us.UbicKubeClient.Inventory(context.Background())
	if err != nil {
		fmt.Println("inventory error ", err.Error())
	}
	fmt.Println(in.Metrics().TotalAvailable)
	fmt.Println(in.Metrics().TotalAllocatable)
	deployments, _ := us.UbicKubeClient.Deployments(context.Background())
	for _, v := range deployments {
		for ki, vi := range v.ManifestGroup().GetResources() {

			fmt.Println("deployments", ki, vi.Resources)
		}

	}

}

// CloseManager close deployment manager
func (us *UbicService) CloseManager(lid ctypes.LeaseID) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ubicDeployTemp, ok := us.Managers[lid]
		if ok {
			err := ubicDeployTemp.Teardown()
			if err == nil {
				delete(us.Managers, lid)
			}
		}
	}(wg)
	wg.Wait()
}

/*
func CreateDeploymentBySdl(sdlFile []byte) {
	sdlTemp, err := sdl.Read(sdlFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
*/

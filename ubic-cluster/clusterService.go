package ubic_cluster

import (
	"context"
	"fmt"
	"github.com/ovrclk/akash/sdl"
	"github.com/tendermint/tendermint/libs/log"
	"providerService/src/cluster"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/cluster/ubickube"
	cmdutil "providerService/src/cmd/provider-services/cmd/util"
	"providerService/src/config"
	"sync"
)

type UbicService struct {
	UbicKubeClient cluster.UbicClient
	Managers       map[ctypes.LeaseID]*cluster.UbicDeploymentManager
	UbicLog        log.Logger
	Config         cluster.Config
	ConfBase       config.ProviderConfig
	Hostnames      *cluster.UbicHostnameService
}

func (us *UbicService) NewService(ctx context.Context, conf config.ProviderConfig, kubeSetting map[interface{}]interface{}) {
	us.ConfBase = conf
	clusterConfig := cluster.NewDefaultConfig()
	clusterConfig.ClusterSettings = kubeSetting
	us.Config = clusterConfig
	logger := cmdutil.OpenLogger().With("cmp", "provider")
	us.UbicLog = logger
	cclient, err := kube.NewClient(context.Background(), logger, conf.NameSpace, conf.K8sConfigPath)
	if err != nil {
		fmt.Println("fail")
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("success")
	}

	allHostnames, _ := cclient.AllHostnames(ctx)
	activeHostnames := make(map[string]ctypes.LeaseID, len(allHostnames))
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
func (us *UbicService) NewUbicDeployManager(lid ctypes.LeaseID, sdlSteam []byte) ([]string, error) {
	sdlFile, err := sdl.Read(sdlSteam)
	if err != nil {
		fmt.Println("exit in sdl conv", err.Error())
	}
	groups, _ := sdlFile.Manifest()
	if len(groups) != 1 {
		return nil, nil
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
		s, _ := (us.UbicKubeClient).LeaseStatus(context.Background(), lid)
		ret := make([]string, 0)
		for k, v := range s {
			var uriTemps string
			for _, v1 := range v.URIs {
				uriTemps += v1
			}
			ret = append(ret, k+uriTemps)
		}
		fmt.Println(len(us.Managers))
		return ret, nil
	}
	return nil, nil
}
func (us *UbicService) IsActive(leaseId ctypes.LeaseID) bool {
	_, ok := us.Managers[leaseId]
	return ok
}
func (us *UbicService) CloseManager(lid ctypes.LeaseID) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ubicDeployTemp, ok := us.Managers[lid]
		if ok {
			err := ubicDeployTemp.Teardown()
			if err != nil {
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

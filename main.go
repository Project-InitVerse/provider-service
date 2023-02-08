package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"os/signal"
	"providerService/src/bid"
	"providerService/src/scan"
	"providerService/src/ubic_operator/hostnameoperator"
	"sync"
	"time"

	"providerService/src/cluster"
	//ctypes "providerService/src/cluster/types/v1"
	"providerService/src/cluster/ubickube/builder"
	cmdutil "providerService/src/cmd/provider-services/cmd/util"
	"providerService/src/config"
	gwrest "providerService/src/gateway/rest"
	"providerService/src/ubic_operator/operatorcommon"
	"providerService/src/util"
	ubic_cluster "providerService/ubic-cluster"
	"runtime"
	"strings"
)

func clusterFunc() {
	var wgGlobal sync.WaitGroup
	ctxGlobal, cancel := context.WithCancel(context.Background())
	runtime.GOMAXPROCS(runtime.NumCPU())
	configHandle := viper.New()
	configData := config.LoadConfig(configHandle)
	configHandle.WatchConfig()
	watch := func(e fsnotify.Event) {
		log.Printf("Config file is changed: %s\n", e.String())
		config.ConvertConfigPtr(configHandle, configData)
	}
	configHandle.OnConfigChange(watch)
	cs := new(ubic_cluster.UbicService)
	kubeSettings := builder.NewDefaultSettings()
	kubeSettings.DeploymentIngressDomain = configData.DeploymentIngressDomain
	kubeSettings.DeploymentIngressExposeLBHosts = configData.DeploymentIngressExposeLBHosts
	kubeSettings.DeploymentIngressStaticHosts = configData.DeploymentIngressStaticHosts
	kubeSettings.NetworkPoliciesEnabled = configData.DeploymentNetworkPoliciesEnabled
	kubeSettings.ClusterPublicHostname = configData.ClusterPublicHostname
	kubeSettings.CPUCommitLevel = configData.OvercommitPercentCPU
	kubeSettings.MemoryCommitLevel = configData.OvercommitPercentMemory
	kubeSettings.StorageCommitLevel = configData.OvercommitPercentStorage
	kubeSettings.DeploymentRuntimeClass = configData.DeploymentRuntimeClass
	kubeSettings.DockerImagePullSecretsName = strings.TrimSpace(configData.DockerImagePullSecretsName)
	hostConf := operatorcommon.OperatorConfig{
		ProviderAddress:    configData.ProviderAddress,
		PruneInterval:      time.Duration(configData.HostPruneInterval),
		WebRefreshInterval: time.Duration(configData.HostWebRefreshInterval),
		RetryDelay:         time.Duration(configData.HostRetryDelay),
	}
	// start HostName service
	go hostnameoperator.DoUbicHostnameOperator(ctxGlobal,
		configData.K8sConfigPath, configData.NameSpace, configData.HostNameServiceListenAddr, &hostConf)

	clusterSettings := map[interface{}]interface{}{
		builder.SettingsKey: kubeSettings,
	}
	clusterConfig := cluster.NewDefaultConfig()
	clusterConfig.DeploymentIngressDomain = configData.DeploymentIngressDomain
	clusterConfig.DeploymentIngressStaticHosts = configData.DeploymentIngressStaticHosts
	clusterConfig.ClusterSettings = clusterSettings
	cs.NewService(ctxGlobal, *configData, clusterConfig)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, os.Interrupt)
	var certFromFlag io.Reader

	kpm, err := util.NewKeyPairManager(common.HexToAddress(configData.ProviderAddress))
	exsit, _ := kpm.KeyExists()
	if !exsit {
		startTime := time.Now().Truncate(time.Second)
		validDuration := time.Hour * 24 * 365

		kpm.Generate(startTime, startTime.Add(validDuration), []string{configData.ClusterPublicHostname})
	}
	if err != nil {
		fmt.Println(err)
	}

	_, tlsCert, err := kpm.ReadX509KeyPair(certFromFlag)

	if err != nil {
		fmt.Println(err)
	}

	if err := builder.ValidateSettings(kubeSettings); err != nil {
		fmt.Println(err)
	}

	logger := cmdutil.OpenLogger().With("cmp", "provider")
	gateway, err := gwrest.NewServer(
		ctxGlobal,
		logger,
		cs,
		configData,

		configData.GatewayListenAddress,
		common.HexToAddress(configData.ProviderContract),
		[]tls.Certificate{tlsCert},
		clusterSettings,
	)
	go gateway.ListenAndServeTLS("", "")
	linkClient := util.LinkClient{
		URL: configData.NodeURL,
	}
	mainScan := scan.Scan{}
	mainScan.InitScan(configData)
	bs := new(bid.Service)
	bs.Init(ctxGlobal, configData,
		mainScan.NeedBidChan,
		mainScan.NeedCreateChan,
		mainScan.UserCancelChan, cs)
	go bs.Run(&wgGlobal)
	go mainScan.MainLoop(ctxGlobal, &linkClient, &wgGlobal)
	for {
		select {
		case <-c:
			fmt.Println("get signal exit")
			cancel()
			wgGlobal.Wait()
			os.Exit(1)
		}
	}
}
func main() {
	clusterFunc()
}

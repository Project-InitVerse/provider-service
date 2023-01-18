package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/signal"
	"providerService/src/scan"
	"sync"
	"time"

	//ctypes "providerService/src/cluster/types/v1"
	"providerService/src/cluster/ubickube/builder"
	cmdutil "providerService/src/cmd/provider-services/cmd/util"
	"providerService/src/config"
	gwrest "providerService/src/gateway/rest"
	"providerService/src/util"
	ubic_cluster "providerService/ubic-cluster"
	"runtime"
	"strings"
)

var (
	config_handle viper.Viper
)

/*
func bidfunc() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config_handle := viper.New()
	config_data := config.LoadConfig(config_handle)
	bs := new(bid.BidService)
	NeedBidChan := make(chan util.NeedBid, 20)
	UserCancelChan := make(chan util.UserCancelOrder, 20)
	NeedCreateChan := make(chan util.NeedCreate, 20)
	bs.Init(config_data, NeedBidChan, NeedCreateChan, UserCancelChan)

	go bs.Run(context.Background())
	tempNeedBid := new(util.NeedBid)
	tempNeedBid.Cpu = new(big.Int).SetInt64(3)
	tempNeedBid.Memory = new(big.Int).SetInt64(6)
	tempNeedBid.Storage = new(big.Int).SetInt64(9)
	tempNeedBid.ContractAddress = "0xb3Ece6Fb0376935A699b26159DAd29c3C7f3340D"
	tempNeedBid.Cert = "haha"
	tempNeedBid.SdlTrxId = "3"
	tempNeedBid.State = 1
	NeedBidChan <- *tempNeedBid
	time.Sleep(3 * time.Second)
	tempNeedCreate := new(util.NeedCreate)
	tempNeedCreate.Provider = common.HexToAddress("0x624199661664b9EF607c02D934EA25701076C3B8")
	tempNeedCreate.ContractAddress = "0xb3Ece6Fb0376935A699b26159DAd29c3C7f3340D"
	tempNeedCreate.CpuPrice, _ = new(big.Int).SetString(config_data.CpuPrice, 10)
	tempNeedCreate.MemoryPrice, _ = new(big.Int).SetString(config_data.MemoryPrice, 10)
	tempNeedCreate.StoragePrice, _ = new(big.Int).SetString(config_data.StoragePrice, 10)
	NeedCreateChan <- *tempNeedCreate
	time.Sleep(30 * time.Second)
}
*/

func clusterFunc() {
	var wgGlobal sync.WaitGroup
	ctxGlobal, cancel := context.WithCancel(context.Background())
	fmt.Println(wgGlobal, ctxGlobal)
	runtime.GOMAXPROCS(runtime.NumCPU())
	config_handle := viper.New()
	config_data := config.LoadConfig(config_handle)
	cs := new(ubic_cluster.UbicService)
	kubeSettings := builder.NewDefaultSettings()
	kubeSettings.DeploymentIngressDomain = config_data.DeploymentIngressDomain
	kubeSettings.DeploymentIngressExposeLBHosts = config_data.DeploymentIngressExposeLBHosts
	kubeSettings.DeploymentIngressStaticHosts = config_data.DeploymentIngressStaticHosts
	kubeSettings.NetworkPoliciesEnabled = config_data.DeploymentNetworkPoliciesEnabled
	kubeSettings.ClusterPublicHostname = config_data.ClusterPublicHostname
	kubeSettings.CPUCommitLevel = config_data.OvercommitPercentCPU
	kubeSettings.MemoryCommitLevel = config_data.OvercommitPercentMemory
	kubeSettings.StorageCommitLevel = config_data.OvercommitPercentStorage
	kubeSettings.DeploymentRuntimeClass = config_data.DeploymentRuntimeClass
	kubeSettings.DockerImagePullSecretsName = strings.TrimSpace(config_data.DockerImagePullSecretsName)

	clusterSettings := map[interface{}]interface{}{
		builder.SettingsKey: kubeSettings,
	}
	cs.NewService(context.Background(), *config_data, clusterSettings)
	//path := "./sdl.txt"
	//buf, err := os.ReadFile(path)
	//if err != nil {
	//	fmt.Println("error happen in read sdl")
	//	fmt.Println(err.Error())
	//	return
	//}
	//lid := ctypes.LeaseID{
	//	Owner:    "0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5",
	//	OSeq:     1,
	//	Provider: "0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5",
	//}
	//cs.NewUbicDeployManager(lid, buf)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, os.Interrupt)

	var certFromFlag io.Reader
	if val := config_data.AuthPem; val != "" {
		certFromFlag = bytes.NewBufferString(val)
	}

	kpm, err := util.NewKeyPairManager(common.HexToAddress(config_data.ProviderAddress))
	exsit, _ := kpm.KeyExists()
	if !exsit {
		startTime := time.Now().Truncate(time.Second)
		validDuration := time.Hour * 24 * 365

		kpm.Generate(startTime, startTime.Add(validDuration), []string{config_data.ClusterPublicHostname})
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

	ctx := context.Background()
	logger := cmdutil.OpenLogger().With("cmp", "provider")
	gateway, err := gwrest.NewServer(
		ctx,
		logger,
		cs,
		config_data,

		config_data.GatewayListenAddress,
		common.HexToAddress(config_data.ProviderAddress),
		[]tls.Certificate{tlsCert},
		clusterSettings,
	)
	go gateway.ListenAndServeTLS("", "")
	link_client := util.LinkClient{
		Url: config_data.NodeUrl,
	}
	main_scan := scan.Scan{}
	main_scan.InitScan(config_data)
	go main_scan.MainLoop(ctxGlobal, &link_client, &wgGlobal)
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

/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config_handle := viper.New()
	config_data := config.LoadConfig(config_handle)
	config_handle.WatchConfig()
	watch := func(e fsnotify.Event) {
		log.Printf("Config file is changed: %s\n", e.String())
		config.ConvertConfigPtr(config_handle, config_data)
	}
	ctx := context.Background()
	config_handle.OnConfigChange(watch)
	//fmt.Println(&config_handle)
	//config_data := config.ConvertConfig(&config_handle)
	fmt.Println(*config_data)
	link_client := util.LinkClient{
		Url: config_data.NodeUrl,
	}
	fmt.Println(&link_client)

	main_scan := scan.Scan{}
	bs := new(bid.BidService)
	main_scan.InitScan(config_data)
	bs.Init(config_data, main_scan.NeedBidChan, main_scan.NeedCreateChan)
	bs.Run(context.Background())
	for {
		fmt.Println(util.GetOrderFactory(main_scan.Config))
		time.Sleep(10 * time.Second)
	}
	logger := cmdutil.OpenLogger().With("cmp", "provider")
	service := cluster.UbicService{}
	var certFromFlag io.Reader
	if val := config_data.AuthPem; val != "" {
		certFromFlag = bytes.NewBufferString(val)
	}

	kpm, err := util.NewKeyPairManager(common.HexToAddress( config_data.ProviderAddress))
	if err != nil {
		fmt.Println(err)
	}

	_, tlsCert, err := kpm.ReadX509KeyPair(certFromFlag)
	if err != nil {
		fmt.Println(err)
	}
	kubeSettings := builder.NewDefaultSettings()
	kubeSettings.DeploymentIngressDomain = config_data.DeploymentIngressDomain
	kubeSettings.DeploymentIngressExposeLBHosts = config_data.DeploymentIngressExposeLBHosts
	kubeSettings.DeploymentIngressStaticHosts = config_data.DeploymentIngressStaticHosts
	kubeSettings.NetworkPoliciesEnabled = config_data.DeploymentNetworkPoliciesEnabled
	kubeSettings.ClusterPublicHostname = config_data.ClusterPublicHostname
	kubeSettings.CPUCommitLevel = config_data.OvercommitPercentCPU
	kubeSettings.MemoryCommitLevel = config_data.OvercommitPercentMemory
	kubeSettings.StorageCommitLevel = config_data.OvercommitPercentStorage
	kubeSettings.DeploymentRuntimeClass = config_data.DeploymentRuntimeClass
	kubeSettings.DockerImagePullSecretsName = strings.TrimSpace(config_data.DockerImagePullSecretsName)

	if err := builder.ValidateSettings(kubeSettings); err != nil {
		fmt.Println( err)
	}

	clusterSettings := map[interface{}]interface{}{
		builder.SettingsKey: kubeSettings,
	}
	gateway, err := gwrest.NewServer(
		ctx,
		logger,
		service,
		config_data,

		config_data.GatewayListenAddress,
		common.HexToAddress(config_data.ProviderAddress),
		[]tls.Certificate{tlsCert},
		clusterSettings,
	)
	go gateway.ListenAndServeTLS("", "")
}
*/
//func gatewayFunc() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	config_handle := viper.New()
//	config_data := config.LoadConfig(config_handle)
//	config_handle.WatchConfig()
//	watch := func(e fsnotify.Event) {
//		log.Printf("Config file is changed: %s\n", e.String())
//		config.ConvertConfigPtr(config_handle, config_data)
//	}
//	ctx := context.Background()
//	config_handle.OnConfigChange(watch)
//	//fmt.Println(&config_handle)
//	//config_data := config.ConvertConfig(&config_handle)
//	fmt.Println(*config_data)
//	link_client := util.LinkClient{
//		Url: config_data.NodeUrl,
//	}
//	fmt.Println(&link_client)
//
//	main_scan := scan.Scan{}
//	main_scan.InitScan(config_data)
//	go main_scan.MainLoop(&link_client)
//
//	logger := cmdutil.OpenLogger().With("cmp", "provider")
//	service := ubic_cluster.UbicService{}
//	var certFromFlag io.Reader
//	if val := config_data.AuthPem; val != "" {
//		certFromFlag = bytes.NewBufferString(val)
//	}
//
//	kpm, err := util.NewKeyPairManager(common.HexToAddress( config_data.ProviderAddress))
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	_, tlsCert, err := kpm.ReadX509KeyPair(certFromFlag)
//	if err != nil {
//		fmt.Println(err)
//	}
//	kubeSettings := builder.NewDefaultSettings()
//	kubeSettings.DeploymentIngressDomain = config_data.DeploymentIngressDomain
//	kubeSettings.DeploymentIngressExposeLBHosts = config_data.DeploymentIngressExposeLBHosts
//	kubeSettings.DeploymentIngressStaticHosts = config_data.DeploymentIngressStaticHosts
//	kubeSettings.NetworkPoliciesEnabled = config_data.DeploymentNetworkPoliciesEnabled
//	kubeSettings.ClusterPublicHostname = config_data.ClusterPublicHostname
//	kubeSettings.CPUCommitLevel = config_data.OvercommitPercentCPU
//	kubeSettings.MemoryCommitLevel = config_data.OvercommitPercentMemory
//	kubeSettings.StorageCommitLevel = config_data.OvercommitPercentStorage
//	kubeSettings.DeploymentRuntimeClass = config_data.DeploymentRuntimeClass
//	kubeSettings.DockerImagePullSecretsName = strings.TrimSpace(config_data.DockerImagePullSecretsName)
//
//	if err := builder.ValidateSettings(kubeSettings); err != nil {
//		fmt.Println( err)
//	}
//
//	clusterSettings := map[interface{}]interface{}{
//		builder.SettingsKey: kubeSettings,
//	}
//	gateway, err := gwrest.NewServer(
//		ctx,
//		logger,
//		service,
//		config_data,
//
//		config_data.GatewayListenAddress,
//		common.HexToAddress(config_data.ProviderAddress),
//		[]tls.Certificate{tlsCert},
//		clusterSettings,
//	)
//	go gateway.ListenAndServeTLS("", "")
//
//	for {
//		fmt.Println(util.GetOrderFactory(main_scan.Config))
//		time.Sleep(10 * time.Second)
//	}
//}

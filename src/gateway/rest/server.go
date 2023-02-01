package rest

import (
	"context"
	"crypto/tls"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/libs/log"
	"net"
	"net/http"
	"providerService/src/config"
	gwutils "providerService/src/gateway/utils"
	ubic_cluster "providerService/ubic-cluster"
)

// NewServer is function create new server
func NewServer(
	ctx context.Context,
	log log.Logger,
	pclient *ubic_cluster.UbicService,
	cConfig *config.ProviderConfig,
	address string,
	pid common.Address,
	certs []tls.Certificate,
	clusterConfig map[interface{}]interface{}) (*http.Server, error) {

	// fixme ovrclk/engineering#609
	// nolint: gosec
	srv := &http.Server{
		Addr:    address,
		Handler: newRouter(log, pid, pclient, clusterConfig),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	var err error

	srv.TLSConfig, err = gwutils.NewServerTLSConfig(context.Background(), certs, cConfig)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

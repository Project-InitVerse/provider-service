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

//
//func NewJwtServer(ctx context.Context,
//	cquery ctypes.QueryClient,
//	jwtGatewayAddr string,
//	providerAddr sdk.Address,
//	cert tls.Certificate,
//	certSerialNumber string,
//	jwtExpiresAfter time.Duration,
//) (*http.Server, error) {
//	// fixme ovrclk/engineering#609
//	// nolint: gosec
//	srv := &http.Server{
//		Addr:    jwtGatewayAddr,
//		Handler: newJwtServerRouter(providerAddr, cert.PrivateKey, jwtExpiresAfter, certSerialNumber),
//		BaseContext: func(_ net.Listener) context.Context {
//			return ctx
//		},
//	}
//
//	var err error
//	srv.TLSConfig, err = gwutils.NewServerTLSConfig(ctx, []tls.Certificate{cert}, cquery)
//	if err != nil {
//		return nil, err
//	}
//
//	return srv, nil
//}
//
//func NewResourceServer(ctx context.Context,
//	log log.Logger,
//	serverAddr string,
//	providerAddr sdk.Address,
//	pubkey *ecdsa.PublicKey,
//	lokiGwAddr string,
//) (*http.Server, error) {
//	// fixme ovrclk/engineering#609
//	// nolint: gosec
//	srv := &http.Server{
//		Addr:        serverAddr,
//		Handler:     newResourceServerRouter(log, providerAddr, pubkey, lokiGwAddr),
//		BaseContext: func(_ net.Listener) context.Context { return ctx },
//	}
//
//	return srv, nil
//}

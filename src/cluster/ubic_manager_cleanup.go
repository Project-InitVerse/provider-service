package cluster

import (
	"context"
	"github.com/ovrclk/akash/manifest/v2beta1"
	logger "github.com/tendermint/tendermint/libs/log"
	ctypes "providerService/src/cluster/types/v1"
)

type ubicPurgeIPEntry struct {
	serviceName string
	port        uint32
	proto       v2beta1.ServiceProtocol
}

type ubicDeployCleanupHelper struct {
	client UbicClient
	log    logger.Logger
	lease  ctypes.LeaseID

	hostnamesToPurge []string
	ipsToPurge       []ubicPurgeIPEntry
}

func NewUbicDeployCleanupHelper(lease ctypes.LeaseID, client UbicClient, log logger.Logger) *ubicDeployCleanupHelper {
	return &ubicDeployCleanupHelper{
		client: client,
		log:    log,
		lease:  lease,
	}
}

func (dch *ubicDeployCleanupHelper) addHostname(hostname string) {
	dch.hostnamesToPurge = append(dch.hostnamesToPurge, hostname)
}

func (dch *ubicDeployCleanupHelper) addIP(serviceName string, port uint32, proto v2beta1.ServiceProtocol) {
	dch.ipsToPurge = append(
		dch.ipsToPurge,
		ubicPurgeIPEntry{
			serviceName: serviceName,
			port:        port,
			proto:       proto,
		})
}

func (dch *ubicDeployCleanupHelper) purgeAll(ctx context.Context) {
	for _, hostname := range dch.hostnamesToPurge {
		err := dch.client.PurgeDeclaredHostname(ctx, dch.lease, hostname)
		if err != nil {
			dch.log.Error("could not purge hostname",
				"lease", dch.lease, "hsotname", hostname, "error", err)
		}
	}

	for _, ipEntry := range dch.ipsToPurge {
		err := dch.client.PurgeDeclaredIP(ctx, dch.lease, ipEntry.serviceName, ipEntry.port, ipEntry.proto)
		if err != nil {
			dch.log.Error("could not purge IP",
				"lease", dch.lease,
				"serviceName", ipEntry.serviceName,
				"port", ipEntry.port,
				"proto", ipEntry.proto,
				"error", err)
		}
	}
}

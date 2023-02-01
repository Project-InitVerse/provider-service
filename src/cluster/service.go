package cluster

import (
	"context"
	"github.com/ethereum/go-ethereum/common"

	crd "providerService/src/ubicpkg/api/ubicnet/v1"

	"github.com/pkg/errors"

	atypes "github.com/ovrclk/akash/types/v1beta2"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"

	ctypes "providerService/src/cluster/types/v1"
)

// ErrNotRunning is the error when service is not running
var ErrNotRunning = errors.New("not running")

var (
//deploymentManagerGauge = promauto.NewGauge(prometheus.GaugeOpts{
//	// fixme provider_deployment_manager
//	Name:        "provider_deploymetn_manager",
//	Help:        "",
//	ConstLabels: nil,
//})
)

// Cluster is the interface that wraps Reserve and Unreserve methods
type Cluster interface {
	Reserve(mtypes.OrderID, atypes.ResourceGroup) (ctypes.Reservation, error)
	Unreserve(mtypes.OrderID) error
}

// StatusClient is the interface which includes status of service
type StatusClient interface {
	Status(context.Context) (*ctypes.Status, error)
	FindActiveLease(ctx context.Context, owner common.Address, dseq uint64, gseq uint32) (bool, mtypes.LeaseID, crd.ManifestGroup, error)
}

// Service manage compute cluster for the provider.  Will eventually integrate with kubernetes, etc...
type Service interface {
	StatusClient
	Cluster
	Close() error
	Ready() <-chan struct{}
	Done() <-chan struct{}
	HostnameService() ctypes.HostnameServiceClient
	TransferHostname(ctx context.Context, leaseID mtypes.LeaseID, hostname string, serviceName string, externalPort uint32) error
}

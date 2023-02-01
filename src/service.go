package provider

import (
	"context"

	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"
	clustertypes "providerService/src/cluster/types/v1"

	"providerService/src/cluster"
	"providerService/src/manifest"
)

// ValidateClient is the interface to check if provider will bid on given groupspec
type ValidateClient interface {
	Validate(context.Context, dtypes.GroupSpec) (ValidateGroupSpecResult, error)
}

// StatusClient is the interface which includes status of service
type StatusClient interface {
	Status(context.Context) (*Status, error)
}

// Client is interface
type Client interface {
	StatusClient
	ValidateClient
	Manifest() manifest.Client
	Cluster() cluster.UbicClient
	Hostname() clustertypes.HostnameServiceClient
	ClusterService() cluster.Service
}

// Service is the interface that includes StatusClient interface.
// It also wraps ManifestHandler, Close and Done methods.
type Service interface {
	Client

	Close() error
	Done() <-chan struct{}
}

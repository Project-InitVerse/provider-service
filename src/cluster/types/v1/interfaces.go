package v1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// HostnameServiceClient is interFace
type HostnameServiceClient interface {
	ReserveHostnames(ctx context.Context, hostnames []string, leaseID LeaseID) ([]string, error)
	ReleaseHostnames(leaseID LeaseID) error
	CanReserveHostnames(hostnames []string, ownerAddr common.Address) error
	PrepareHostnamesForTransfer(ctx context.Context, hostnames []string, leaseID LeaseID) error
}

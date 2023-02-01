package v1

import (
	manifest "github.com/ovrclk/akash/manifest/v2beta1"
)

// IPResourceEvent is interFace
type IPResourceEvent interface {
	GetLeaseID() LeaseID
	GetServiceName() string
	GetExternalPort() uint32
	GetPort() uint32
	GetSharingKey() string
	GetProtocol() manifest.ServiceProtocol
	GetEventType() ProviderResourceEvent
}

// IPPassthrough is interFace
type IPPassthrough interface {
	GetLeaseID() LeaseID
	GetServiceName() string
	GetExternalPort() uint32
	GetPort() uint32
	GetSharingKey() string
	GetProtocol() manifest.ServiceProtocol
}

// IPLeaseState is interFace
type IPLeaseState interface {
	IPPassthrough
	GetIP() string
}

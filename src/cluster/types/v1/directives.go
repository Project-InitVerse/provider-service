package v1

import (
	manifest "github.com/ovrclk/akash/manifest/v2beta1"
)

type ConnectHostnameToDeploymentDirective struct {
	Hostname    string
	LeaseID     LeaseID
	ServiceName string
	ServicePort int32
	ReadTimeout uint32
	SendTimeout uint32
	NextTimeout uint32
	MaxBodySize uint32
	NextTries   uint32
	NextCases   []string
}

type ClusterIPPassthroughDirective struct {
	LeaseID      LeaseID
	ServiceName  string
	Port         uint32
	ExternalPort uint32
	SharingKey   string
	Protocol     manifest.ServiceProtocol
}

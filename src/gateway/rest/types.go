package rest

import (
	cltypes "providerService/src/cluster/types/v1"
)

//LeasedIPStatus is struct
type LeasedIPStatus struct {
	Port         uint32
	ExternalPort uint32
	Protocol     string
	IP           string
}

//LeaseStatus is struct
type LeaseStatus struct {
	Services       map[string]*cltypes.ServiceStatus        `json:"services"`
	ForwardedPorts map[string][]cltypes.ForwardedPortStatus `json:"forwarded_ports"` // Container services that are externally accessible
	IPs            map[string][]LeasedIPStatus              `json:"ips"`
}

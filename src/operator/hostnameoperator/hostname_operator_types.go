package hostnameoperator

import (
	"time"

	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"

	ctypes "providerService/src/cluster/types/v1beta2"
)

type managedHostname struct {
	lastEvent    ctypes.HostnameResourceEvent
	presentLease mtypes.LeaseID

	presentServiceName  string
	presentExternalPort uint32
	lastChangeAt        time.Time
}

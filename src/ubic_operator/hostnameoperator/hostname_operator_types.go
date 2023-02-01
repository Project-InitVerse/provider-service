package hostnameoperator

import (
	"time"

	ctypes "providerService/src/cluster/types/v1"
)

type managedHostname struct {
	lastEvent    ctypes.HostnameResourceEvent
	presentLease ctypes.LeaseID

	presentServiceName  string
	presentExternalPort uint32
	lastChangeAt        time.Time
}

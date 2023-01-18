package v1

type LeaseIDHostnameConnection interface {
	GetLeaseID() LeaseID
	GetHostname() string
	GetExternalPort() int32
	GetServiceName() string
}

type ActiveHostname struct {
	ID       LeaseID
	Hostname string
}

type ProviderResourceEvent string

const (
	ProviderResourceAdd    = ProviderResourceEvent("add")
	ProviderResourceUpdate = ProviderResourceEvent("update")
	ProviderResourceDelete = ProviderResourceEvent("delete")
)

type HostnameResourceEvent interface {
	GetLeaseID() LeaseID
	GetEventType() ProviderResourceEvent
	GetHostname() string
	GetServiceName() string
	GetExternalPort() uint32
}

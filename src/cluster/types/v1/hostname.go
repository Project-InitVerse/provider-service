package v1

//LeaseIDHostnameConnection is interFace
type LeaseIDHostnameConnection interface {
	GetLeaseID() LeaseID
	GetHostname() string
	GetExternalPort() int32
	GetServiceName() string
}

//ActiveHostname is struct
type ActiveHostname struct {
	ID       LeaseID
	Hostname string
}

//ProviderResourceEvent is string
type ProviderResourceEvent string

const (
	//ProviderResourceAdd is add event string
	ProviderResourceAdd = ProviderResourceEvent("add")
	//ProviderResourceUpdate is update event string
	ProviderResourceUpdate = ProviderResourceEvent("update")
	//ProviderResourceDelete is delete event string
	ProviderResourceDelete = ProviderResourceEvent("delete")
)

//HostnameResourceEvent is interFace
type HostnameResourceEvent interface {
	GetLeaseID() LeaseID
	GetEventType() ProviderResourceEvent
	GetHostname() string
	GetServiceName() string
	GetExternalPort() uint32
}

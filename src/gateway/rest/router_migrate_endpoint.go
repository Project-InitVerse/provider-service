package rest

import (
	manifest "github.com/ovrclk/akash/manifest/v2beta1"

	"providerService/src/ubicpkg/api/ubicnet/v1"
)

type endpointMigrateRequestBody struct {
	EndpointsToMigrate []string `json:"endpoints_to_migrate"`
	DestinationDSeq    uint64   `json:"destination_dseq"`
	DestinationGSeq    uint32   `json:"destination_gseq"`
}

type serviceExposeWithName struct {
	expose      v1.ManifestServiceExpose
	serviceName string
	proto       manifest.ServiceProtocol
}

package ubicutil

import atypes "github.com/ovrclk/akash/types/v1beta2"

// GetEndpointQuantityOfResourceGroup is function
func GetEndpointQuantityOfResourceGroup(resources atypes.ResourceGroup, kind atypes.Endpoint_Kind) uint {
	endpoints := make(map[uint32]struct{})
	for _, resource := range resources.GetResources() {
		accumEndpointsOfResources(resource.Resources, kind, endpoints)

	}
	return uint(len(endpoints))
}

func accumEndpointsOfResources(r atypes.ResourceUnits, kind atypes.Endpoint_Kind, accum map[uint32]struct{}) {
	for _, endpoint := range r.Endpoints {
		if endpoint.Kind == kind {
			accum[endpoint.SequenceNumber] = struct{}{}
		}
	}
}

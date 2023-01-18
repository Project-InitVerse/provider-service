package clientcommon

import (
	"errors"
	"fmt"
	ctype "providerService/src/cluster/types/v1"
	"strconv"

	"providerService/src/cluster/ubickube/builder"
)

var (
	errMissingLabel      = errors.New("kube: missing label")
	errInvalidLabelValue = errors.New("kube: invalid label value")
)

// TODO - move to provider/cluster/util since this is generic
func RecoverLeaseIDFromLabels(labels map[string]string) (ctype.LeaseID, error) {
	oseqS, ok := labels[builder.AkashLeaseOSeqLabelName]
	if !ok {
		return ctype.LeaseID{}, fmt.Errorf("%w: %q", errMissingLabel, builder.AkashLeaseOSeqLabelName)
	}
	owner, ok := labels[builder.AkashLeaseOwnerLabelName]
	if !ok {
		return ctype.LeaseID{}, fmt.Errorf("%w: %q", errMissingLabel, builder.AkashLeaseOwnerLabelName)
	}

	provider, ok := labels[builder.AkashLeaseProviderLabelName]
	if !ok {
		return ctype.LeaseID{}, fmt.Errorf("%w: %q", errMissingLabel, builder.AkashLeaseProviderLabelName)
	}

	oseq, err := strconv.ParseUint(oseqS, 10, 32)
	if err != nil {
		return ctype.LeaseID{}, fmt.Errorf("%w: oesq %q not a uint", errInvalidLabelValue, oseqS)
	}

	return ctype.LeaseID{
		Owner:    owner,
		OSeq:     oseq,
		Provider: provider,
	}, nil
}

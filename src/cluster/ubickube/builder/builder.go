package builder

// nolint:deadcode,golint

import (
	"fmt"
	ctypes "providerService/src/cluster/types/v1"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/util/intstr"

	manifesttypes "github.com/ovrclk/akash/manifest/v2beta1"
	clusterUtil "providerService/src/cluster/ubicutil"
)

const (
	AkashManagedLabelName         = "ubic.net"
	AkashManifestServiceLabelName = "ubic.net/manifest-service"
	AkashNetworkStorageClasses    = "ubic.net/storageclasses"
	AkashServiceTarget            = "ubic.net/service-target"
	AkashMetalLB                  = "metal-lb"

	akashNetworkNamespace = "ubic.net/namespace"

	AkashLeaseOwnerLabelName    = "ubic.net/lease.id.owner"
	AkashLeaseDSeqLabelName     = "ubic.net/lease.id.dseq"
	AkashLeaseGSeqLabelName     = "ubic.net/lease.id.gseq"
	AkashLeaseOSeqLabelName     = "ubic.net/lease.id.oseq"
	AkashLeaseProviderLabelName = "ubic.net/lease.id.provider"
	AkashLeaseManifestVersion   = "ubic.net/manifest.version"
	akashDeploymentPolicyName   = "akash-deployment-restrictions"
)

const runtimeClassNoneValue = "none"

const (
	envVarAkashGroupSequence         = "AKASH_GROUP_SEQUENCE"
	envVarAkashDeploymentSequence    = "AKASH_DEPLOYMENT_SEQUENCE"
	envVarAkashOrderSequence         = "AKASH_ORDER_SEQUENCE"
	envVarAkashOwner                 = "AKASH_OWNER"
	envVarAkashProvider              = "AKASH_PROVIDER"
	envVarAkashClusterPublicHostname = "AKASH_CLUSTER_PUBLIC_HOSTNAME"
)

var (
	dnsPort     = intstr.FromInt(53)
	udpProtocol = corev1.Protocol("UDP")
	tcpProtocol = corev1.Protocol("TCP")
)

type builderBase interface {
	NS() string
	Name() string
	Validate() error
}

type builder struct {
	log      log.Logger
	settings Settings
	lid      ctypes.LeaseID
	group    *manifesttypes.Group
}

var _ builderBase = (*builder)(nil)

func (b *builder) NS() string {
	return LidNS(b.lid)
}

func (b *builder) Name() string {
	return b.NS()
}

func (b *builder) labels() map[string]string {
	return map[string]string{
		AkashManagedLabelName: "true",
		akashNetworkNamespace: LidNS(b.lid),
	}
}

func (b *builder) Validate() error {
	return nil
}

func addIfNotPresent(envVarsAlreadyAdded map[string]int, env []corev1.EnvVar, key string, value interface{}) []corev1.EnvVar {
	_, exists := envVarsAlreadyAdded[key]
	if exists {
		return env
	}

	env = append(env, corev1.EnvVar{Name: key, Value: fmt.Sprintf("%v", value)})
	return env
}

const SuffixForNodePortServiceName = "-np"

func makeGlobalServiceNameFromBasename(basename string) string {
	return fmt.Sprintf("%s%s", basename, SuffixForNodePortServiceName)
}

// LidNS generates a unique sha256 sum for identifying a provider's object name.
func LidNS(lid ctypes.LeaseID) string {
	return clusterUtil.LeaseIDToNamespace(lid)
}

func AppendLeaseLabels(lid ctypes.LeaseID, labels map[string]string) map[string]string {
	labels[AkashLeaseOwnerLabelName] = lid.Owner
	labels[AkashLeaseOSeqLabelName] = strconv.FormatUint(uint64(lid.OSeq), 10)
	labels[AkashLeaseProviderLabelName] = lid.Provider
	return labels
}

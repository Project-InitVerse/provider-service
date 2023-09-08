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

// ubic kube const val
const (
	UbicManagedLabelName         = "ubic.net"
	UbicManifestServiceLabelName = "ubic.net/manifest-service"
	UbicNetworkStorageClasses    = "ubic.net/storageclasses"
	UbicServiceTarget            = "ubic.net/service-target"
	UbicMetalLB                  = "metal-lb"

	ubicNetworkNamespace = "ubic.net/namespace"

	UbicLeaseOwnerLabelName    = "ubic.net/lease.id.owner"
	UbicLeaseDSeqLabelName     = "ubic.net/lease.id.dseq"
	UbicLeaseOSeqLabelName     = "ubic.net/lease.id.oseq"
	UbicLeaseProviderLabelName = "ubic.net/lease.id.provider"
	ubicDeploymentPolicyName   = "ubic-deployment-restrictions"
)

const runtimeClassNoneValue = "none"

const (
	envVarIniOrderSequence         = "INI_ORDER_SEQUENCE"
	envVarIniOwner                 = "INI_OWNER"
	envVarIniProvider              = "INI_PROVIDER"
	envVarIniClusterPublicHostname = "INI_CLUSTER_PUBLIC_HOSTNAME"
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
		UbicManagedLabelName: "true",
		ubicNetworkNamespace: LidNS(b.lid),
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

// SuffixForNodePortServiceName is string
const SuffixForNodePortServiceName = "-np"

func makeGlobalServiceNameFromBasename(basename string) string {
	return fmt.Sprintf("%s%s", basename, SuffixForNodePortServiceName)
}

// LidNS generates a unique sha256 sum for identifying a provider's object name.
func LidNS(lid ctypes.LeaseID) string {
	return clusterUtil.LeaseIDToNamespace(lid)
}

// AppendLeaseLabels is function to generate lables from lease
func AppendLeaseLabels(lid ctypes.LeaseID, labels map[string]string) map[string]string {
	labels[UbicLeaseOwnerLabelName] = lid.Owner
	labels[UbicLeaseOSeqLabelName] = strconv.FormatUint(uint64(lid.OSeq), 10)
	labels[UbicLeaseProviderLabelName] = lid.Provider
	return labels
}

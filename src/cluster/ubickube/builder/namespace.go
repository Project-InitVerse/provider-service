package builder

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctypes "providerService/src/cluster/types/v1"

	manitypes "github.com/ovrclk/akash/manifest/v2beta1"
)

// NS is interface
type NS interface {
	builderBase
	Create() (*corev1.Namespace, error)
	Update(obj *corev1.Namespace) (*corev1.Namespace, error)
}

type ns struct {
	builder
}

var _ NS = (*ns)(nil)

// BuildNS is function build namespace
func BuildNS(settings Settings, lid ctypes.LeaseID, group *manitypes.Group) NS {
	return &ns{builder: builder{settings: settings, lid: lid, group: group}}
}

func (b *ns) labels() map[string]string {
	return AppendLeaseLabels(b.lid, b.builder.labels())
}

func (b *ns) Create() (*corev1.Namespace, error) { // nolint:golint,unparam
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   b.NS(),
			Labels: b.labels(),
		},
	}, nil
}

func (b *ns) Update(obj *corev1.Namespace) (*corev1.Namespace, error) { // nolint:golint,unparam
	obj.Name = b.NS()
	obj.Labels = b.labels()
	return obj, nil
}

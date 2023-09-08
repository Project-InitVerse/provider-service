package kube

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"

	manifest "github.com/ovrclk/akash/manifest/v2beta1"
	kubeErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/pager"

	"providerService/src/cluster/types/v1"
	ctypes "providerService/src/cluster/types/v1"
	"providerService/src/cluster/ubickube/builder"
	kubeclienterrors "providerService/src/cluster/ubickube/errors"
	ubictypes "providerService/src/ubicpkg/api/ubicnet/v1"
)

const (
	serviceNameLabel  = "service-name"
	externalPortLabel = "external-port"
	protoLabel        = "proto"
)

func (c *client) GetDeclaredIPs(ctx context.Context, leaseID ctypes.LeaseID) ([]ubictypes.ProviderLeasedIPSpec, error) {
	labelSelector := &strings.Builder{}
	kubeSelectorForLease(labelSelector, leaseID)
	fmt.Println(labelSelector.String())
	fmt.Println(c.ns)
	results, err := c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return nil, err
	}
	retval := make([]ubictypes.ProviderLeasedIPSpec, 0, len(results.Items))
	for _, item := range results.Items {
		retval = append(retval, item.Spec)
	}

	return retval, nil
}

func (c *client) PurgeDeclaredIP(ctx context.Context, leaseID ctypes.LeaseID, serviceName string, externalPort uint32, proto manifest.ServiceProtocol) error {
	labelSelector := &strings.Builder{}
	kubeSelectorForLease(labelSelector, leaseID)
	_, _ = fmt.Fprintf(labelSelector, ",%s=%s", serviceNameLabel, serviceName)
	_, _ = fmt.Fprintf(labelSelector, ",%s=%s", protoLabel, proto.ToString())
	_, _ = fmt.Fprintf(labelSelector, ",%s=%d", externalPortLabel, externalPort)
	return c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=true", builder.UbicManagedLabelName),
	})
}

func (c *client) DeclareIP(ctx context.Context, lID ctypes.LeaseID, serviceName string, port uint32, externalPort uint32, proto manifest.ServiceProtocol, sharingKey string, overwrite bool) error {
	// Note: This interface expects sharing key to contain a value that is unique per deployment owner, in this
	// case it is the bech32 address, or a derivative thereof
	resourceName := strings.ToLower(fmt.Sprintf("%s-%s-%d", sharingKey, proto.ToString(), externalPort))

	c.log.Debug("checking for resource", "resource-name", resourceName)
	foundEntry, err := c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).Get(ctx, resourceName, metav1.GetOptions{})
	exists := true
	if err != nil {
		if !kubeErrors.IsNotFound(err) {
			return err
		}
		exists = false
	}

	if exists && !overwrite {
		return kubeclienterrors.ErrAlreadyExists
	}

	labels := map[string]string{
		builder.UbicManagedLabelName: "true",
		serviceNameLabel:             serviceName,
		externalPortLabel:            fmt.Sprintf("%d", externalPort),
		protoLabel:                   proto.ToString(),
	}
	builder.AppendLeaseLabels(lID, labels)

	obj := ubictypes.ProviderLeasedIP{
		ObjectMeta: metav1.ObjectMeta{
			Name:   resourceName,
			Labels: labels,
		},
		Spec: ubictypes.ProviderLeasedIPSpec{
			LeaseID:      ubictypes.LeaseIDFromInitype(lID),
			ServiceName:  serviceName,
			ExternalPort: externalPort,
			SharingKey:   sharingKey,
			Protocol:     proto.ToString(),
			Port:         port,
		},
		Status: ubictypes.ProviderLeasedIPStatus{},
	}

	c.log.Info("declaring leased ip", "lease", lID,
		"service-name", serviceName,
		"port", port,
		"external-port", externalPort,
		"sharing-key", sharingKey,
		"exists", exists)
	// Create or update the entry
	if exists {
		obj.ObjectMeta.ResourceVersion = foundEntry.ResourceVersion
		_, err = c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).Update(ctx, &obj, metav1.UpdateOptions{})
	} else {
		_, err = c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).Create(ctx, &obj, metav1.CreateOptions{})
	}

	return err
}

func (c *client) PurgeDeclaredIPs(ctx context.Context, lID ctypes.LeaseID) error {
	labelSelector := &strings.Builder{}
	_, err := fmt.Fprintf(labelSelector, "%s=true,", builder.UbicManagedLabelName)
	if err != nil {
		return err
	}
	kubeSelectorForLease(labelSelector, lID)
	result := c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})

	return result
}

func (c *client) ObserveIPState(ctx context.Context) (<-chan v1.IPResourceEvent, error) {
	var lastResourceVersion string
	phpager := pager.New(func(ctx context.Context, opts metav1.ListOptions) (runtime.Object, error) {
		resources, err := c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).List(ctx, opts)

		if err == nil && len(resources.GetResourceVersion()) != 0 {
			lastResourceVersion = resources.GetResourceVersion()
		}
		return resources, err
	})

	data := make([]ubictypes.ProviderLeasedIP, 0, 128)
	err := phpager.EachListItem(ctx, metav1.ListOptions{}, func(obj runtime.Object) error {
		plip := obj.(*ubictypes.ProviderLeasedIP)
		data = append(data, *plip)
		return nil
	})

	if err != nil {
		return nil, err
	}

	c.log.Info("starting ip passthrough watch", "resourceVersion", lastResourceVersion)
	watcher, err := c.ac.UbicnetV1().ProviderLeasedIPs(c.ns).Watch(ctx, metav1.ListOptions{
		TypeMeta:             metav1.TypeMeta{},
		LabelSelector:        "",
		FieldSelector:        "",
		Watch:                false,
		AllowWatchBookmarks:  false,
		ResourceVersion:      lastResourceVersion,
		ResourceVersionMatch: "",
		TimeoutSeconds:       nil,
		Limit:                0,
		Continue:             "",
	})
	if err != nil {
		return nil, err
	}

	evData := make([]ipResourceEvent, len(data))
	for i, v := range data {
		ownerAddr := common.HexToAddress(v.Spec.LeaseID.Owner)
		providerAddr := common.HexToAddress(v.Spec.LeaseID.Provider)
		leaseID, err := v.Spec.LeaseID.ToIniType()
		if err != nil {
			return nil, err
		}

		proto, err := manifest.ParseServiceProtocol(v.Spec.Protocol)
		if err != nil {
			return nil, err
		}

		ev := ipResourceEvent{
			eventType:    ctypes.ProviderResourceAdd,
			lID:          leaseID,
			serviceName:  v.Spec.ServiceName,
			port:         v.Spec.Port,
			externalPort: v.Spec.ExternalPort,
			ownerAddr:    ownerAddr,
			providerAddr: providerAddr,
			sharingKey:   v.Spec.SharingKey,
			protocol:     proto,
		}
		evData[i] = ev
	}

	data = nil

	output := make(chan v1.IPResourceEvent)

	go func() {
		defer close(output)
		for _, v := range evData {
			output <- v
		}
		evData = nil // do not hold the reference

		results := watcher.ResultChan()
		for {
			select {
			case result, ok := <-results:
				if !ok { // Channel closed when an error happens
					return
				}
				plip := result.Object.(*ubictypes.ProviderLeasedIP)

				ownerAddr := common.HexToAddress(plip.Spec.LeaseID.Owner)
				providerAddr := common.HexToAddress(plip.Spec.LeaseID.Provider)
				leaseID, err := plip.Spec.LeaseID.ToIniType()
				if err != nil {
					c.log.Error("invalid lease ID", "err", err)
					continue // Ignore event
				}
				proto, err := manifest.ParseServiceProtocol(plip.Spec.Protocol)
				if err != nil {
					c.log.Error("invalid protocol", "err", err)
					continue
				}

				ev := ipResourceEvent{
					lID:          leaseID,
					serviceName:  plip.Spec.ServiceName,
					port:         plip.Spec.Port,
					externalPort: plip.Spec.ExternalPort,
					sharingKey:   plip.Spec.SharingKey,
					providerAddr: providerAddr,
					ownerAddr:    ownerAddr,
					protocol:     proto,
				}
				switch result.Type {

				case watch.Added:
					ev.eventType = ctypes.ProviderResourceAdd
				case watch.Modified:
					ev.eventType = ctypes.ProviderResourceUpdate
				case watch.Deleted:
					ev.eventType = ctypes.ProviderResourceDelete

				case watch.Error:
					// Based on examination of the implementation code, this is basically never called anyways
					c.log.Error("watch error", "err", result.Object)

				default:
					continue
				}

				output <- ev

			case <-ctx.Done():
				return
			}
		}
	}()

	return output, nil
}

type ipResourceEvent struct {
	lID          ctypes.LeaseID
	eventType    ctypes.ProviderResourceEvent
	serviceName  string
	port         uint32
	externalPort uint32
	sharingKey   string
	providerAddr common.Address
	ownerAddr    common.Address
	protocol     manifest.ServiceProtocol
}

func (ev ipResourceEvent) GetLeaseID() ctypes.LeaseID {
	return ev.lID
}

func (ev ipResourceEvent) GetEventType() ctypes.ProviderResourceEvent {
	return ev.eventType
}

func (ev ipResourceEvent) GetServiceName() string {
	return ev.serviceName
}

func (ev ipResourceEvent) GetPort() uint32 {
	return ev.port
}

func (ev ipResourceEvent) GetExternalPort() uint32 {
	return ev.externalPort
}

func (ev ipResourceEvent) GetSharingKey() string {
	return ev.sharingKey
}

func (ev ipResourceEvent) GetProtocol() manifest.ServiceProtocol {
	return ev.protocol
}

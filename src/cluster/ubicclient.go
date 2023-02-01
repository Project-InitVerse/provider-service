package cluster

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"

	crd "providerService/src/ubicpkg/api/ubicnet/v1"

	"github.com/ovrclk/akash/sdl"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/remotecommand"

	eventsv1 "k8s.io/api/events/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"

	manifest "github.com/ovrclk/akash/manifest/v2beta1"

	ctypes "providerService/src/cluster/types/v1"

	types "github.com/ovrclk/akash/types/v1beta2"

	"github.com/ovrclk/akash/types/unit"
	ubictypes "providerService/src/ubicpkg/api/ubicnet/v1"
)

// Errors types returned by the Exec function on the client interface
var (
	ErrUbicExec                        = errors.New("remote command execute error")
	ErrUbicExecNoServiceWithName       = fmt.Errorf("%w: no such service exists with that name", ErrUbicExec)
	ErrUbicExecServiceNotRunning       = fmt.Errorf("%w: service with that name is not running", ErrUbicExec)
	ErrUbicExecCommandExecutionFailed  = fmt.Errorf("%w: command execution failed", ErrUbicExec)
	ErrUbicExecCommandDoesNotExist     = fmt.Errorf("%w: command could not be executed because it does not exist", ErrUbicExec)
	ErrUbicExecDeploymentNotYetRunning = fmt.Errorf("%w: deployment is not yet active", ErrUbicExec)
	ErrUbicExecPodIndexOutOfRange      = fmt.Errorf("%w: pod index out of range", ErrUbicExec)
	ErrUbicUnknownStorageClass         = errors.New("inventory: unknown storage class")
	errUbicNotImplemented              = errors.New("not implemented")
)

var _ UbicClient = (*nullUbicClient)(nil)

// UbicReadClient is interface to get k8s info
type UbicReadClient interface {
	LeaseStatus(context.Context, ctypes.LeaseID) (map[string]*ctypes.ServiceStatus, error)
	ForwardedPortStatus(context.Context, ctypes.LeaseID) (map[string][]ctypes.ForwardedPortStatus, error)
	LeaseEvents(context.Context, ctypes.LeaseID, string, bool) (ctypes.EventsWatcher, error)
	LeaseLogs(context.Context, ctypes.LeaseID, string, bool, *int64) ([]*ctypes.ServiceLog, error)
	ServiceStatus(context.Context, ctypes.LeaseID, string) (*ctypes.ServiceStatus, error)

	AllHostnames(context.Context) ([]ctypes.ActiveHostname, error)
	GetManifestGroup(context.Context, ctypes.LeaseID) (bool, crd.ManifestGroup, error)

	ObserveHostnameState(ctx context.Context) (<-chan ctypes.HostnameResourceEvent, error)
	GetHostnameDeploymentConnections(ctx context.Context) ([]ctypes.LeaseIDHostnameConnection, error)

	ObserveIPState(ctx context.Context) (<-chan ctypes.IPResourceEvent, error)
	GetDeclaredIPs(ctx context.Context, leaseID ctypes.LeaseID) ([]ubictypes.ProviderLeasedIPSpec, error)
}

// UbicClient interface lease and deployment methods
type UbicClient interface {
	UbicReadClient
	Deploy(ctx context.Context, lID ctypes.LeaseID, mgroup *manifest.Group) error
	TeardownLease(context.Context, ctypes.LeaseID) error
	Deployments(context.Context) ([]ctypes.Deployment, error)
	Inventory(context.Context) (ctypes.Inventory, error)
	Exec(ctx context.Context,
		lID ctypes.LeaseID,
		service string,
		podIndex uint,
		cmd []string,
		stdin io.Reader,
		stdout io.Writer,
		stderr io.Writer,
		tty bool,
		tsq remotecommand.TerminalSizeQueue) (ctypes.ExecResult, error)

	// Connect a given hostname to a deployment
	ConnectHostnameToDeployment(ctx context.Context, directive ctypes.ConnectHostnameToDeploymentDirective) error
	// Remove a given hostname from a deployment
	RemoveHostnameFromDeployment(ctx context.Context, hostname string, leaseID ctypes.LeaseID, allowMissing bool) error

	// Declare that a given deployment should be connected to a given hostname
	DeclareHostname(ctx context.Context, lID ctypes.LeaseID, host string, serviceName string, externalPort uint32) error
	// Purge any hostnames associated with a given deployment
	PurgeDeclaredHostnames(ctx context.Context, lID ctypes.LeaseID) error

	PurgeDeclaredHostname(ctx context.Context, lID ctypes.LeaseID, hostname string) error

	// KubeVersion returns the version information of kubernetes running in the cluster
	KubeVersion() (*version.Info, error)

	DeclareIP(ctx context.Context, lID ctypes.LeaseID, serviceName string, port uint32, externalPort uint32, proto manifest.ServiceProtocol, sharingKey string, overwrite bool) error
	PurgeDeclaredIP(ctx context.Context, lID ctypes.LeaseID, serviceName string, externalPort uint32, proto manifest.ServiceProtocol) error
	PurgeDeclaredIPs(ctx context.Context, lID ctypes.LeaseID) error
}

// UbicErrorIsOkToSendToClient is function
func UbicErrorIsOkToSendToClient(err error) bool {
	return errors.Is(err, ErrUbicExec)
}

type ubicResourcePair struct {
	allocatable sdk.Int
	allocated   sdk.Int
}

type ubicStorageClassState struct {
	ubicResourcePair
	isDefault bool
}

func (rp *ubicResourcePair) dup() ubicResourcePair {
	return ubicResourcePair{
		allocatable: rp.allocatable.AddRaw(0),
		allocated:   rp.allocated.AddRaw(0),
	}
}

func (rp *ubicResourcePair) subNLZ(val types.ResourceValue) bool {
	avail := rp.available()

	res := avail.Sub(val.Val)
	if res.IsNegative() {
		return false
	}

	*rp = ubicResourcePair{
		allocatable: rp.allocatable.AddRaw(0),
		allocated:   rp.allocated.Add(val.Val),
	}

	return true
}

func (rp ubicResourcePair) available() sdk.Int {
	return rp.allocatable.Sub(rp.allocated)
}

type ubicNode struct {
	id               string
	cpu              ubicResourcePair
	memory           ubicResourcePair
	ephemeralStorage ubicResourcePair
}

type ubicClusterStorage map[string]*ubicStorageClassState

func (cs ubicClusterStorage) dup() ubicClusterStorage {
	res := make(ubicClusterStorage)
	for k, v := range cs {
		res[k] = &ubicStorageClassState{
			ubicResourcePair: v.ubicResourcePair.dup(),
			isDefault:        v.isDefault,
		}
	}

	return res
}

type ubicInventory struct {
	storage ubicClusterStorage
	nodes   []*ubicNode
}

var _ ctypes.Inventory = (*ubicInventory)(nil)

func (inv *ubicInventory) Adjust(reservation ctypes.Reservation) error {
	resources := make([]types.Resources, len(reservation.Resources().GetResources()))
	copy(resources, reservation.Resources().GetResources())

	currInventory := inv.dup()

nodes:
	for nodeName, nd := range currInventory.nodes {
		// with persistent storage go through iff there is capacity available
		// there is no point to go through any other node without available storage
		currResources := resources[:0]

		for _, res := range resources {
			for ; res.Count > 0; res.Count-- {
				var adjusted bool

				cpu := nd.cpu.dup()
				if adjusted = cpu.subNLZ(res.Resources.CPU.Units); !adjusted {
					continue nodes
				}

				memory := nd.memory.dup()
				if adjusted = memory.subNLZ(res.Resources.Memory.Quantity); !adjusted {
					continue nodes
				}

				ephemeralStorage := nd.ephemeralStorage.dup()
				storageClasses := currInventory.storage.dup()

				for idx, storage := range res.Resources.Storage {
					attr := storage.Attributes.Find(sdl.StorageAttributePersistent)

					if persistent, _ := attr.AsBool(); !persistent {
						if adjusted = ephemeralStorage.subNLZ(storage.Quantity); !adjusted {
							continue nodes
						}
						continue
					}

					attr = storage.Attributes.Find(sdl.StorageAttributeClass)
					class, _ := attr.AsString()

					if class == sdl.StorageClassDefault {
						for name, params := range storageClasses {
							if params.isDefault {
								class = name

								for i := range storage.Attributes {
									if storage.Attributes[i].Key == sdl.StorageAttributeClass {
										res.Resources.Storage[idx].Attributes[i].Value = class
										break
									}
								}
								break
							}
						}
					}

					cstorage, activeStorageClass := storageClasses[class]
					if !activeStorageClass {
						continue nodes
					}

					if adjusted = cstorage.subNLZ(storage.Quantity); !adjusted {
						// cluster storage does not have enough space thus break to error
						break nodes
					}
				}

				// all requirements for current group have been satisfied
				// commit and move on
				currInventory.nodes[nodeName] = &ubicNode{
					id:               nd.id,
					cpu:              cpu,
					memory:           memory,
					ephemeralStorage: ephemeralStorage,
				}
			}

			if res.Count > 0 {
				currResources = append(currResources, res)
			}
		}

		resources = currResources
	}

	if len(resources) == 0 {
		*inv = *currInventory

		return nil
	}

	return ctypes.ErrInsufficientCapacity
}

func (inv *ubicInventory) Metrics() ctypes.InventoryMetrics {
	cpuTotal := uint64(0)
	memoryTotal := uint64(0)
	storageEphemeralTotal := uint64(0)
	storageTotal := make(map[string]int64)

	cpuAvailable := uint64(0)
	memoryAvailable := uint64(0)
	storageEphemeralAvailable := uint64(0)
	storageAvailable := make(map[string]int64)

	ret := ctypes.InventoryMetrics{
		Nodes: make([]ctypes.InventoryNode, 0, len(inv.nodes)),
	}

	for _, nd := range inv.nodes {
		invNode := ctypes.InventoryNode{
			Name: nd.id,
			Allocatable: ctypes.InventoryNodeMetric{
				CPU:              nd.cpu.allocatable.Uint64(),
				Memory:           nd.memory.allocatable.Uint64(),
				StorageEphemeral: nd.ephemeralStorage.allocatable.Uint64(),
			},
		}

		cpuTotal += nd.cpu.allocatable.Uint64()
		memoryTotal += nd.memory.allocatable.Uint64()
		storageEphemeralTotal += nd.ephemeralStorage.allocatable.Uint64()

		tmp := nd.cpu.allocatable.Sub(nd.cpu.allocated)
		invNode.Available.CPU = tmp.Uint64()
		cpuAvailable += invNode.Available.CPU

		tmp = nd.memory.allocatable.Sub(nd.memory.allocated)
		invNode.Available.Memory = tmp.Uint64()
		memoryAvailable += invNode.Available.Memory

		tmp = nd.ephemeralStorage.allocatable.Sub(nd.ephemeralStorage.allocated)
		invNode.Available.StorageEphemeral = tmp.Uint64()
		storageEphemeralAvailable += invNode.Available.StorageEphemeral

		ret.Nodes = append(ret.Nodes, invNode)
	}

	ret.TotalAllocatable = ctypes.InventoryMetricTotal{
		CPU:              cpuTotal,
		Memory:           memoryTotal,
		StorageEphemeral: storageEphemeralTotal,
		Storage:          storageTotal,
	}

	ret.TotalAvailable = ctypes.InventoryMetricTotal{
		CPU:              cpuAvailable,
		Memory:           memoryAvailable,
		StorageEphemeral: storageEphemeralAvailable,
		Storage:          storageAvailable,
	}

	return ret
}

func (inv *ubicInventory) dup() *ubicInventory {
	res := &ubicInventory{
		nodes: make([]*ubicNode, 0, len(inv.nodes)),
	}

	for _, nd := range inv.nodes {
		res.nodes = append(res.nodes, &ubicNode{
			id:               nd.id,
			cpu:              nd.cpu.dup(),
			memory:           nd.memory.dup(),
			ephemeralStorage: nd.ephemeralStorage.dup(),
		})
	}

	return res
}

const (
	// 5 CPUs, 5Gi memory for null client.
	nullUbicClientCPU     = 5000
	nullUbicClientMemory  = 32 * unit.Gi
	nullUbicClientStorage = 512 * unit.Gi
)

type nullUbicLease struct {
	ctx    context.Context
	cancel func()
	group  *manifest.Group
}

type nullUbicClient struct {
	leases map[string]*nullUbicLease
	mtx    sync.Mutex
}

// NewUbicServiceLog creates and returns a service log with provided details
func NewUbicServiceLog(name string, stream io.ReadCloser) *ctypes.ServiceLog {
	return &ctypes.ServiceLog{
		Name:    name,
		Stream:  stream,
		Scanner: bufio.NewScanner(stream),
	}
}

// NullUbicClient returns nullClient instance
func NullUbicClient() UbicClient {
	return &nullUbicClient{
		leases: make(map[string]*nullUbicLease),
		mtx:    sync.Mutex{},
	}
}

func (c *nullUbicClient) RemoveHostnameFromDeployment(ctx context.Context, hostname string, leaseID ctypes.LeaseID, allowMissing bool) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) ObserveHostnameState(ctx context.Context) (<-chan ctypes.HostnameResourceEvent, error) {
	return nil, errUbicNotImplemented
}
func (c *nullUbicClient) GetDeployments(ctx context.Context, dID dtypes.DeploymentID) ([]ctypes.Deployment, error) {
	return nil, errUbicNotImplemented
}
func (c *nullUbicClient) GetHostnameDeploymentConnections(ctx context.Context) ([]ctypes.LeaseIDHostnameConnection, error) {
	return nil, errUbicNotImplemented
}

// Connect a given hostname to a deployment
func (c *nullUbicClient) ConnectHostnameToDeployment(ctx context.Context, directive ctypes.ConnectHostnameToDeploymentDirective) error {
	return errUbicNotImplemented
}

// Declare that a given deployment should be connected to a given hostname
func (c *nullUbicClient) DeclareHostname(ctx context.Context, lID ctypes.LeaseID, host string, serviceName string, externalPort uint32) error {
	return errUbicNotImplemented
}

// Purge any hostnames associated with a given deployment
func (c *nullUbicClient) PurgeDeclaredHostnames(ctx context.Context, lID ctypes.LeaseID) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) PurgeDeclaredHostname(ctx context.Context, lID ctypes.LeaseID, hostname string) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) Deploy(ctx context.Context, lid ctypes.LeaseID, mgroup *manifest.Group) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	ctx, cancel := context.WithCancel(ctx)
	c.leases[ctypes.LeasePath(lid)] = &nullUbicLease{
		ctx:    ctx,
		cancel: cancel,
		group:  mgroup,
	}

	return nil
}

func (*nullUbicClient) ForwardedPortStatus(context.Context, ctypes.LeaseID) (map[string][]ctypes.ForwardedPortStatus, error) {
	return nil, errUbicNotImplemented
}

func (c *nullUbicClient) LeaseStatus(_ context.Context, lid ctypes.LeaseID) (map[string]*ctypes.ServiceStatus, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	lease, ok := c.leases[ctypes.LeasePath(lid)]
	if !ok {
		return nil, nil
	}

	resp := make(map[string]*ctypes.ServiceStatus)
	for _, svc := range lease.group.Services {
		resp[svc.Name] = &ctypes.ServiceStatus{
			Name:      svc.Name,
			Available: int32(svc.Count),
			Total:     int32(svc.Count),
		}
	}

	return resp, nil
}

func (c *nullUbicClient) LeaseEvents(ctx context.Context, lid ctypes.LeaseID, _ string, follow bool) (ctypes.EventsWatcher, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	lease, ok := c.leases[ctypes.LeasePath(lid)]
	if !ok {
		return nil, nil
	}

	if lease.ctx.Err() != nil {
		return nil, nil
	}

	feed := ctypes.NewEventsFeed(ctx)
	go func() {
		defer feed.Shutdown()

		tm := time.NewTicker(7 * time.Second)
		tm.Stop()

		genEvent := func() *eventsv1.Event {
			return &eventsv1.Event{
				EventTime:           v1.NewMicroTime(time.Now()),
				ReportingController: lease.group.Name,
			}
		}

		nfollowCh := make(chan *eventsv1.Event, 1)
		count := 0
		if !follow {
			count = rand.Intn(9) // nolint: gosec
			nfollowCh <- genEvent()
		} else {
			tm.Reset(time.Second)
		}

		for {
			select {
			case <-lease.ctx.Done():
				return
			case evt := <-nfollowCh:
				if !feed.SendEvent(evt) || count == 0 {
					return
				}
				count--
				nfollowCh <- genEvent()
				break
			case <-tm.C:
				tm.Stop()
				if !feed.SendEvent(genEvent()) {
					return
				}
				tm.Reset(time.Duration(rand.Intn(9)+1) * time.Second) // nolint: gosec
				break
			}
		}
	}()

	return feed, nil
}

func (c *nullUbicClient) ServiceStatus(_ context.Context, _ ctypes.LeaseID, _ string) (*ctypes.ServiceStatus, error) {
	return nil, nil
}

func (c *nullUbicClient) LeaseLogs(_ context.Context, _ ctypes.LeaseID, _ string, _ bool, _ *int64) ([]*ctypes.ServiceLog, error) {
	return nil, nil
}

func (c *nullUbicClient) TeardownLease(_ context.Context, lid ctypes.LeaseID) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if lease, ok := c.leases[ctypes.LeasePath(lid)]; ok {
		delete(c.leases, ctypes.LeasePath(lid))
		lease.cancel()
	}

	return nil
}

func (c *nullUbicClient) Deployments(context.Context) ([]ctypes.Deployment, error) {
	return nil, nil
}

func (c *nullUbicClient) Inventory(context.Context) (ctypes.Inventory, error) {
	inv := &ubicInventory{
		nodes: []*ubicNode{
			{
				id: "solo",
				cpu: ubicResourcePair{
					allocatable: sdk.NewInt(nullUbicClientCPU),
					allocated:   sdk.NewInt(nullUbicClientCPU - 100),
				},
				memory: ubicResourcePair{
					allocatable: sdk.NewInt(nullUbicClientMemory),
					allocated:   sdk.NewInt(nullUbicClientMemory - unit.Gi),
				},
				ephemeralStorage: ubicResourcePair{
					allocatable: sdk.NewInt(nullUbicClientStorage),
					allocated:   sdk.NewInt(nullUbicClientStorage - (10 * unit.Gi)),
				},
			},
		},
		storage: map[string]*ubicStorageClassState{
			"beta2": {
				ubicResourcePair: ubicResourcePair{
					allocatable: sdk.NewInt(nullUbicClientStorage),
					allocated:   sdk.NewInt(nullUbicClientStorage - (10 * unit.Gi)),
				},
				isDefault: true,
			},
		},
	}

	return inv, nil
}

func (c *nullUbicClient) Exec(context.Context, ctypes.LeaseID, string, uint, []string, io.Reader, io.Writer, io.Writer, bool, remotecommand.TerminalSizeQueue) (ctypes.ExecResult, error) {
	return nil, errUbicNotImplemented
}

func (c *nullUbicClient) GetManifestGroup(context.Context, ctypes.LeaseID) (bool, crd.ManifestGroup, error) {
	return false, crd.ManifestGroup{}, nil
}

func (c *nullUbicClient) AllHostnames(context.Context) ([]ctypes.ActiveHostname, error) {
	return nil, nil
}

func (c *nullUbicClient) KubeVersion() (*version.Info, error) {
	return nil, nil
}

func (c *nullUbicClient) DeclareIP(ctx context.Context, lID ctypes.LeaseID, serviceName string, port uint32, externalPort uint32, proto manifest.ServiceProtocol, sharingKey string, overwrite bool) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) PurgeDeclaredIPs(ctx context.Context, lID ctypes.LeaseID) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) ObserveIPState(ctx context.Context) (<-chan ctypes.IPResourceEvent, error) {
	return nil, errUbicNotImplemented
}

func (c *nullUbicClient) CreateIPPassthrough(ctx context.Context, lID ctypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) PurgeIPPassthrough(ctx context.Context, lID ctypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) PurgeDeclaredIP(ctx context.Context, lID ctypes.LeaseID, serviceName string, externalPort uint32, proto manifest.ServiceProtocol) error {
	return errUbicNotImplemented
}

func (c *nullUbicClient) GetDeclaredIPs(ctx context.Context, leaseID ctypes.LeaseID) ([]ubictypes.ProviderLeasedIPSpec, error) {
	return nil, errUbicNotImplemented
}

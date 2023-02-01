package cluster

import (
	"context"
	"errors"
	"fmt"
	"github.com/ovrclk/provider-services/cluster/util"
	"sync"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/avast/retry-go"
	"github.com/boz/go-lifecycle"
	manifest "github.com/ovrclk/akash/manifest/v2beta1"
	sdlutil "github.com/ovrclk/akash/sdl/util"
	clustertypes "providerService/src/cluster/types/v1"
	kubeclienterrors "providerService/src/cluster/ubickube/errors"
	clusterutil "providerService/src/cluster/ubicutil"
)

const (
	ubicDsDeployActive     ubicDeploymentState = "deploy-active"
	ubicDsDeployPending    ubicDeploymentState = "deploy-pending"
	ubicDsDeployComplete   ubicDeploymentState = "deploy-complete"
	ubicDsTeardownActive   ubicDeploymentState = "teardown-active"
	ubicDsTeardownPending  ubicDeploymentState = "teardown-pending"
	ubicDsTeardownComplete ubicDeploymentState = "teardown-complete"
)

const ubicUncleanShutdownGracePeriod = 30 * time.Second

type ubicDeploymentState string

var (
	/*
		ubicDeploymentCounter = promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "provider_deployment",
		}, []string{"action", "result"})

		ubicMonitorCounter = promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "provider_deployment_monitor",
		}, []string{"action"})
	*/

	errUbicLeaseInactive = errors.New("inactive Lease")
)

// UbicDeploymentManager is struct
type UbicDeploymentManager struct {
	client UbicClient

	state ubicDeploymentState

	lease  clustertypes.LeaseID
	mgroup *manifest.Group

	monitor          *deploymentMonitor
	wg               sync.WaitGroup
	updatech         chan *manifest.Group
	teardownch       chan struct{}
	currentHostnames map[string]struct{}

	log             log.Logger
	lc              lifecycle.Lifecycle
	hostnameService clustertypes.HostnameServiceClient

	config Config

	serviceShuttingDown <-chan struct{}
}

// NewUbicDeploymentManager create deployment manager
func NewUbicDeploymentManager(ctx context.Context,
	client UbicClient,
	lease clustertypes.LeaseID,
	ubicLog log.Logger,
	mgroup *manifest.Group,
	hostServ *UbicHostnameService,
	conf Config,
	isNewLease bool) *UbicDeploymentManager {

	logger := ubicLog.With("cmp", "deployment-manager", "lease", lease, "manifest-group", mgroup.Name)

	dm := &UbicDeploymentManager{
		client:              client,
		state:               ubicDsDeployActive,
		lease:               lease,
		mgroup:              mgroup,
		wg:                  sync.WaitGroup{},
		updatech:            make(chan *manifest.Group),
		teardownch:          make(chan struct{}),
		log:                 logger,
		lc:                  lifecycle.New(),
		hostnameService:     hostServ,
		config:              conf,
		serviceShuttingDown: ctx.Done(),
		currentHostnames:    make(map[string]struct{}),
	}

	go dm.lc.WatchChannel(ctx.Done())

	go dm.run(context.Background())

	go func() {
		<-dm.lc.Done()
		dm.log.Debug("sending manager into channel")

	}()
	return dm
}

func (dm *UbicDeploymentManager) update(mgroup *manifest.Group) error {
	select {
	case dm.updatech <- mgroup:
		return nil
	case <-dm.lc.ShuttingDown():
		return ErrNotRunning
	}
}

// Teardown is shutdown deployment
func (dm *UbicDeploymentManager) Teardown() error {
	select {
	case dm.teardownch <- struct{}{}:
		return nil
	case <-dm.lc.ShuttingDown():
		return ErrNotRunning
	}
}

func (dm *UbicDeploymentManager) handleUpdate(ctx context.Context) <-chan error {
	switch dm.state {
	case ubicDsDeployActive:
		dm.state = ubicDsDeployPending
	case ubicDsDeployComplete:
		// start update
		return dm.startDeploy(ctx)
	case ubicDsDeployPending, ubicDsTeardownActive, ubicDsTeardownPending, ubicDsTeardownComplete:
		// do nothing
	}

	return nil
}

func (dm *UbicDeploymentManager) run(ctx context.Context) {
	defer dm.lc.ShutdownCompleted()
	var shutdownErr error

	runch := dm.startDeploy(ctx)

	defer func() {
		err := dm.hostnameService.ReleaseHostnames(dm.lease)
		if err != nil {
			dm.log.Error("failed releasing hostnames", "err", err)
		}
		dm.log.Debug("hostnames released")
	}()

	var teardownErr error

loop:
	for {
		select {
		case shutdownErr = <-dm.lc.ShutdownRequest():
			break loop

		case mgroup := <-dm.updatech:
			dm.mgroup = mgroup
			newch := dm.handleUpdate(ctx)
			if newch != nil {
				runch = newch
			}

		case result := <-runch:
			runch = nil
			if result != nil {
				dm.log.Error("execution error", "state", dm.state, "err", result)
			}
			switch dm.state {
			case ubicDsDeployActive:
				if result != nil {
					// Run the teardown code to get rid of anything created that might be hanging out
					runch = dm.startTeardown()
				} else {

					dm.log.Debug("deploy complete")
					dm.state = ubicDsDeployComplete
					dm.startMonitor()
				}
			case ubicDsDeployPending:
				if result != nil {
					break loop
				}
				// start update
				runch = dm.startDeploy(ctx)
			case ubicDsDeployComplete:
				//panic(fmt.Sprintf("INVALID STATE: runch read on %v", dm.state))
			case ubicDsTeardownActive:
				teardownErr = result
				dm.state = ubicDsTeardownComplete
				dm.log.Debug("teardown complete")
				break loop
			case ubicDsTeardownPending:
				// start teardown
				runch = dm.startTeardown()
			case ubicDsTeardownComplete:
				//panic(fmt.Sprintf("INVALID STATE: runch read on %v", dm.state))
			}

		case <-dm.teardownch:
			dm.log.Debug("teardown request")
			dm.stopMonitor()
			switch dm.state {
			case ubicDsDeployActive:
				dm.state = ubicDsTeardownPending
			case ubicDsDeployPending:
				dm.state = ubicDsTeardownPending
			case ubicDsDeployComplete:
				// start teardown
				runch = dm.startTeardown()
			case ubicDsTeardownActive, ubicDsTeardownPending, ubicDsTeardownComplete:
			}
		}
	}

	dm.log.Debug("shutting down")
	dm.lc.ShutdownInitiated(shutdownErr)
	if runch != nil {
		<-runch
		dm.log.Debug("read from runch during shutdown")
	}

	dm.log.Debug("waiting on dm.wg")
	dm.wg.Wait()

	if dm.state != ubicDsDeployComplete {
		dm.log.Info("shutting down unclean, running teardown now")
		ctx, cancel := context.WithTimeout(context.Background(), ubicUncleanShutdownGracePeriod)
		defer cancel()
		teardownErr = dm.doTeardown(ctx)
	}

	if teardownErr != nil {
		dm.log.Error("lease teardown failed", "err", teardownErr)
	}

	dm.log.Info("shutdown complete")
}

func (dm *UbicDeploymentManager) startMonitor() {
	//TODO
	/*
		dm.wg.Add(1)
		dm.monitor = newDeploymentMonitor(dm)
		monitorCounter.WithLabelValues("start").Inc()
		go func(m *deploymentMonitor) {
			defer dm.wg.Done()
			<-m.done()
		}(dm.monitor)*/
}

func (dm *UbicDeploymentManager) stopMonitor() {
	if dm.monitor != nil {
		//monitorCounter.WithLabelValues("stop").Inc()
		dm.monitor.shutdown()
	}
}

func (dm *UbicDeploymentManager) startDeploy(ctx context.Context) <-chan error {
	dm.stopMonitor()
	dm.state = ubicDsDeployActive

	chErr := make(chan error, 1)

	go func() {
		hostnames, endpoints, err := dm.doDeploy(ctx)
		if err != nil {
			chErr <- err
			return
		}

		if len(hostnames) != 0 {
			// Some hostnames have been withheld
			dm.log.Info("hostnames withheld from deployment", "cnt", len(hostnames), "lease", dm.lease)
		}

		if len(endpoints) != 0 {
			// Some endpoints have been withheld
			dm.log.Info("endpoints withheld from deployment", "cnt", len(endpoints), "lease", dm.lease)
		}
		//TODO
		/*
			groupCopy := *dm.mgroup
			ev := event.ClusterDeployment{
				LeaseID: dm.lease,
				Group:   &groupCopy,
				Status:  event.ClusterDeploymentUpdated,
			}
			err = dm.bus.Publish(ev)
			if err != nil {
				dm.log.Error("failed publishing event", "err", err)
			}*/

		close(chErr)
	}()
	return chErr
}

func (dm *UbicDeploymentManager) startTeardown() <-chan error {
	dm.stopMonitor()
	dm.state = ubicDsTeardownActive
	return dm.do(func() error {
		// Don't use a context tied to the lifecycle, as we don't want to cancel Kubernetes operations
		return dm.doTeardown(context.Background())
	})
}

type ubicServiceExposeWithServiceName struct {
	expose manifest.ServiceExpose
	name   string
}

func (dm *UbicDeploymentManager) doDeploy(ctx context.Context) ([]string, []string, error) {
	cleanupHelper := NewUbicDeployCleanupHelper(dm.lease, dm.client, dm.log)

	var err error
	ctx, cancel := context.WithCancel(context.Background())

	// Weird hack to tie this context to the lifecycle of the parent service, so this doesn't
	// block forever or anything like that
	go func() {
		select {
		case <-dm.serviceShuttingDown:
			cancel()
		case <-ctx.Done():
		}
	}()

	defer func() {
		// TODO - run on an isolated context
		cleanupHelper.purgeAll(ctx)
		cancel()
	}()

	if err = dm.checkLeaseActive(ctx); err != nil {
		return nil, nil, err
	}

	currentIPs, err := dm.client.GetDeclaredIPs(ctx, dm.lease)
	fmt.Println(currentIPs)
	if err != nil {
		return nil, nil, err
	}
	// Either reserve the hostnames, or confirm that they already are held
	allHostnames := sdlutil.AllHostnamesOfManifestGroup(*dm.mgroup)
	withheldHostnames, err := dm.hostnameService.ReserveHostnames(ctx, allHostnames, dm.lease)
	if err != nil {
		//deploymentCounter.WithLabelValues("reserve-hostnames", "err").Inc()
		dm.log.Error("deploy hostname reservation error", "state", dm.state, "err", err)
		return nil, nil, err
	}
	//deploymentCounter.WithLabelValues("reserve-hostnames", "success").Inc()

	dm.log.Info("hostnames withheld", "cnt", len(withheldHostnames))

	hostnamesInThisRequest := make(map[string]struct{})
	for _, hostname := range allHostnames {
		hostnamesInThisRequest[hostname] = struct{}{}
	}

	// Figure out what hostnames were removed from the manifest if any
	for hostnameInUse := range dm.currentHostnames {
		_, stillInUse := hostnamesInThisRequest[hostnameInUse]
		if !stillInUse {
			cleanupHelper.addHostname(hostnameInUse)
		}
	}

	// Don't use a context tied to the lifecycle, as we don't want to cancel Kubernetes operations
	deployCtx := util.ApplyToContext(context.Background(), dm.config.ClusterSettings)

	err = dm.client.Deploy(deployCtx, dm.lease, dm.mgroup)
	//label := "success"
	//if err != nil {
	//	label = "fail"
	//}
	//deploymentCounter.WithLabelValues("deploy", label).Inc()

	// Figure out what hostnames to declare
	blockedHostnames := make(map[string]struct{})
	for _, hostname := range withheldHostnames {
		blockedHostnames[hostname] = struct{}{}
	}
	hosts := make(map[string]manifest.ServiceExpose)
	leasedIPs := make([]ubicServiceExposeWithServiceName, 0)
	hostToServiceName := make(map[string]string)

	ipsInThisRequest := make(map[string]ubicServiceExposeWithServiceName)
	// clear this out so it gets repopulated
	dm.currentHostnames = make(map[string]struct{})
	// Iterate over each entry, extracting the ingress services & leased IPs
	fmt.Println("service length", len(dm.mgroup.Services))
	for _, service := range dm.mgroup.Services {
		fmt.Println("expose is ", service.Expose)
		for _, expose := range service.Expose {
			if sdlutil.ShouldBeIngress(expose) {
				fmt.Println("enter ingress")
				if dm.config.DeploymentIngressStaticHosts {
					fmt.Println("enter this")
					uid := clustertypes.IngressHost(dm.lease, service.Name)
					host := fmt.Sprintf("%s.%s", uid, dm.config.DeploymentIngressDomain)
					hosts[host] = expose
					hostToServiceName[host] = service.Name
				}

				for _, host := range expose.Hosts {
					_, blocked := blockedHostnames[host]
					if !blocked {
						dm.currentHostnames[host] = struct{}{}
						hosts[host] = expose
						hostToServiceName[host] = service.Name
					}
				}
			}

			if expose.Global && len(expose.IP) != 0 {
				v := ubicServiceExposeWithServiceName{expose: expose, name: service.Name}
				leasedIPs = append(leasedIPs, v)
				sharingKey := clusterutil.MakeIPSharingKey(dm.lease, expose.IP)
				ipsInThisRequest[sharingKey] = v
			}
		}
	}
	for _, currentIP := range currentIPs {
		// Check if the IP exists in the compute cluster but not in the presently used set of IPs
		_, stillInUse := ipsInThisRequest[currentIP.SharingKey]
		if !stillInUse {
			proto, err := manifest.ParseServiceProtocol(currentIP.Protocol)
			if err != nil {
				return withheldHostnames, nil, err
			}
			cleanupHelper.addIP(currentIP.ServiceName, currentIP.ExternalPort, proto)
		}
	}
	fmt.Println("hosts", hosts)
	for host, serviceExpose := range hosts {
		externalPort := uint32(sdlutil.ExposeExternalPort(serviceExpose))
		err = dm.client.DeclareHostname(ctx, dm.lease, host, hostToServiceName[host], externalPort)
		if err != nil {
			// TODO - counter
			return withheldHostnames, nil, err
		}
	}

	withheldEndpoints := make([]string, 0)
	fmt.Println("leasedIPs", leasedIPs)
	for _, serviceExpose := range leasedIPs {
		endpointName := serviceExpose.expose.IP
		sharingKey := clusterutil.MakeIPSharingKey(dm.lease, endpointName)

		externalPort := sdlutil.ExposeExternalPort(serviceExpose.expose)
		port := serviceExpose.expose.Port

		err = dm.client.DeclareIP(ctx, dm.lease, serviceExpose.name, uint32(port), uint32(externalPort), serviceExpose.expose.Proto, sharingKey, false)
		if err != nil {
			if !errors.Is(err, kubeclienterrors.ErrAlreadyExists) {
				dm.log.Error("failed adding IP declaration", "service", serviceExpose.name, "port", externalPort, "endpoint", serviceExpose.expose.IP, "err", err)
				return withheldHostnames, nil, err
			}
			dm.log.Info("IP declaration already exists", "service", serviceExpose.name, "port", externalPort, "endpoint", serviceExpose.expose.IP, "err", err)
			withheldEndpoints = append(withheldEndpoints, sharingKey)

		} else {
			dm.log.Debug("added IP declaration", "service", serviceExpose.name, "port", externalPort, "endpoint", serviceExpose.expose.IP)
		}
	}
	return withheldHostnames, withheldEndpoints, nil
}

func (dm *UbicDeploymentManager) getCleanupRetryOpts(ctx context.Context) []retry.Option {
	retryFn := func(err error) bool {
		isCanceled := errors.Is(err, context.Canceled)
		isDeadlineExceeeded := errors.Is(err, context.DeadlineExceeded)
		return !isCanceled && !isDeadlineExceeeded
	}
	return []retry.Option{
		retry.Attempts(50),
		retry.Delay(100 * time.Millisecond),
		retry.MaxDelay(3000 * time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.RetryIf(retryFn),
		retry.Context(ctx),
	}
}

func (dm *UbicDeploymentManager) doTeardown(ctx context.Context) error {
	const teardownActivityCount = 3
	teardownResults := make(chan error, teardownActivityCount)

	go func() {
		result := retry.Do(func() error {
			err := dm.client.TeardownLease(ctx, dm.lease)
			if err != nil {
				dm.log.Error("lease teardown failed", "err", err)
			}
			return err
		}, dm.getCleanupRetryOpts(ctx)...)

		//label := "success"
		//if result != nil {
		//	label = "fail"
		//}
		//deploymentCounter.WithLabelValues("teardown", label).Inc()
		teardownResults <- result
	}()

	go func() {
		result := retry.Do(func() error {
			err := dm.client.PurgeDeclaredHostnames(ctx, dm.lease)
			if err != nil {
				dm.log.Error("purge declared hostname failure", "err", err)
			}
			return err
		}, dm.getCleanupRetryOpts(ctx)...)
		// TODO - counter

		if result == nil {
			dm.log.Debug("purged hostnames")
		}
		teardownResults <- result
	}()

	go func() {
		result := retry.Do(func() error {
			err := dm.client.PurgeDeclaredIPs(ctx, dm.lease)
			if err != nil {
				dm.log.Error("purge declared ips failure", "err", err)
			}
			return err
		}, dm.getCleanupRetryOpts(ctx)...)
		// TODO - counter

		if result == nil {
			dm.log.Debug("purged ips")
		}
		teardownResults <- result
	}()

	var firstError error
	for i := 0; i != teardownActivityCount; i++ {
		select {
		case err := <-teardownResults:
			if err != nil && firstError == nil {
				firstError = err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return firstError
}

func (dm *UbicDeploymentManager) checkLeaseActive(ctx context.Context) error {
	//TODO handle
	/*
		var lease *mtypes.QueryLeaseResponse

				err := retry.Do(func() error {
					var err error
					lease, err = dm.session.Client().Query().Lease(ctx, &mtypes.QueryLeaseRequest{
						ID: dm.lease,
					})
					if err != nil {
						dm.log.Error("lease query failed", "err")
					}
					return err
				},
					retry.Attempts(50),
					retry.Delay(100*time.Millisecond),
					retry.MaxDelay(3000*time.Millisecond),
					retry.DelayType(retry.BackOffDelay),
					retry.LastErrorOnly(true))

				if err != nil {
					return err
				}

			if lease.GetLease().State != mtypes.LeaseActive {
				dm.log.Error("lease not active, not deploying")
				return fmt.Errorf("%w: %s", errUbicLeaseInactive, dm.lease)
			}*/

	return nil
}

func (dm *UbicDeploymentManager) do(fn func() error) <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- fn()
	}()
	return ch
}

func ubicTieContextToChannel(parentCtx context.Context, donech <-chan struct{}) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		select {
		case <-donech:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

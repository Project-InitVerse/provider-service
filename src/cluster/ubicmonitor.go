package cluster

import (
	"context"
	"math/rand"
	ctypes "providerService/src/cluster/types/v1"
	"time"

	util "providerService/src/cluster/ubicutil"

	"github.com/boz/go-lifecycle"
	"github.com/tendermint/tendermint/libs/log"

	manifest "github.com/ovrclk/akash/manifest/v2beta1"
	"github.com/ovrclk/akash/util/runner"
)

const (
	monitorMaxRetries        = 40
	monitorRetryPeriodMin    = time.Second * 4 // nolint revive
	monitorRetryPeriodJitter = time.Second * 15

	monitorHealthcheckPeriodMin    = time.Second * 10 // nolint revive
	monitorHealthcheckPeriodJitter = time.Second * 5
)

var (
//deploymentHealthCheckCounter = promauto.NewCounterVec(prometheus.CounterOpts{
//	Name: "provider_deployment_monitor_health",
//}, []string{"state"})
)

type deploymentMonitor struct {
	client UbicClient

	lease  ctypes.LeaseID
	mgroup *manifest.Group

	attempts int
	log      log.Logger
	lc       lifecycle.Lifecycle

	clusterSettings map[interface{}]interface{}
}

func newDeploymentMonitor(dm *UbicDeploymentManager) *deploymentMonitor {
	m := &deploymentMonitor{
		client:          dm.client,
		lease:           dm.lease,
		mgroup:          dm.mgroup,
		log:             dm.log.With("cmp", "deployment-monitor"),
		lc:              lifecycle.New(),
		clusterSettings: dm.config.ClusterSettings,
	}

	go m.lc.WatchChannel(dm.lc.ShuttingDown())
	go m.run()

	return m
}

func (m *deploymentMonitor) shutdown() {
	m.lc.ShutdownAsync(nil)
}

func (m *deploymentMonitor) done() <-chan struct{} {
	return m.lc.Done()
}

func (m *deploymentMonitor) run() {
	defer m.lc.ShutdownCompleted()
	ctx, cancel := context.WithCancel(context.Background())

	var (
		runch   <-chan runner.Result
		closech <-chan runner.Result
	)

	tickch := m.scheduleRetry()

loop:
	for {
		select {
		case err := <-m.lc.ShutdownRequest():
			m.log.Debug("shutting down")
			m.lc.ShutdownInitiated(err)
			break loop

		case <-tickch:
			tickch = nil
			runch = m.runCheck(ctx)

		case result := <-runch:
			runch = nil

			if err := result.Error(); err != nil {
				//deploymentHealthCheckCounter.WithLabelValues("err").Inc()
				m.log.Error("monitor check", "err", err)
			}

			ok := result.Value().(bool)

			m.log.Info("check result", "ok", ok, "attempt", m.attempts)

			if ok {
				// healthy
				m.attempts = 0
				tickch = m.scheduleHealthcheck()
				//deploymentHealthCheckCounter.WithLabelValues("up").Inc()

				break
			} else {
				//deploymentHealthCheckCounter.WithLabelValues("down").Inc()
			}

			if m.attempts <= monitorMaxRetries {
				// unhealthy.  retry
				tickch = m.scheduleRetry()
				break
			}

			m.log.Error("deployment failed.  closing lease.")
			//deploymentHealthCheckCounter.WithLabelValues("failed").Inc()

		case <-closech:
			closech = nil
		}
	}
	cancel()

	if runch != nil {
		m.log.Debug("read runch")
		<-runch
	}

	if closech != nil {
		m.log.Debug("read closech")
		<-closech
	}

	// TODO
	// Check that we got here
	m.log.Debug("shutdown complete")
}

func (m *deploymentMonitor) runCheck(ctx context.Context) <-chan runner.Result {
	m.attempts++
	m.log.Debug("running check", "attempt", m.attempts)
	return runner.Do(func() runner.Result {
		return runner.NewResult(m.doCheck(ctx))
	})
}

func (m *deploymentMonitor) doCheck(ctx context.Context) (bool, error) {
	clientCtx := util.ApplyToContext(ctx, m.clusterSettings)

	status, err := m.client.LeaseStatus(clientCtx, m.lease)

	if err != nil {
		m.log.Error("lease status", "err", err)
		return false, err
	}

	badsvc := 0

	for _, spec := range m.mgroup.Services {
		service, foundService := status[spec.Name]
		if foundService {
			if uint32(service.Available) < spec.Count {
				badsvc++
				m.log.Debug("service available replicas below target",
					"service", spec.Name,
					"available", service.Available,
					"target", spec.Count,
				)
			}
		}

		if !foundService {
			badsvc++
			m.log.Debug("service status not found", "service", spec.Name)
		}
	}

	return badsvc == 0, nil
}

func (m *deploymentMonitor) scheduleRetry() <-chan time.Time {
	return m.schedule(monitorRetryPeriodMin, monitorRetryPeriodJitter)
}

func (m *deploymentMonitor) scheduleHealthcheck() <-chan time.Time {
	return m.schedule(monitorHealthcheckPeriodMin, monitorHealthcheckPeriodJitter)
}

func (m *deploymentMonitor) schedule(min, jitter time.Duration) <-chan time.Time {
	period := min + time.Duration(rand.Int63n(int64(jitter))) // nolint: gosec
	return time.After(period)
}

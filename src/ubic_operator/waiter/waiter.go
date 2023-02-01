package waiter

import (
	"context"
	"time"

	"github.com/tendermint/tendermint/libs/log"
)

// OperatorWaiter is interface
type OperatorWaiter interface {
	WaitForAll(ctx context.Context) error
}

//Waitable is interface
type Waitable interface {
	Check(ctx context.Context) error
	String() string
}

type nullWaiter struct{}

func (nw nullWaiter) WaitForAll(ctx context.Context) error {
	return nil
}

//NewNullWaiter create Null Waiter
func NewNullWaiter() OperatorWaiter {
	return nullWaiter{}
}

type operatorWaiter struct {
	waitables   []Waitable
	log         log.Logger
	delayPeriod time.Duration
	allReady    chan struct{}
}

// NewOperatorWaiter create new wait
func NewOperatorWaiter(ctx context.Context, logger log.Logger, waitOn ...Waitable) OperatorWaiter {
	waiter := &operatorWaiter{
		waitables:   waitOn,
		log:         logger.With("cmp", "waiter"),
		delayPeriod: 2 * time.Second,
		allReady:    make(chan struct{}),
	}

	go waiter.run(ctx)

	return waiter
}

func (w *operatorWaiter) WaitForAll(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-w.allReady:
		return nil
	}
}

func (w *operatorWaiter) run(ctx context.Context) {
	for _, waitable := range w.waitables {
		for {
			err := waitable.Check(ctx)
			if err != nil {
				w.log.Error("not yet ready", "waitable", waitable, "error", err)

				select {
				case <-ctx.Done():
					return
				case <-time.After(w.delayPeriod):
				}

				continue
			}
			break
		}

		w.log.Info("ready", "waitable", waitable)
	}
	w.log.Info("all waitables ready")

	close(w.allReady)
}

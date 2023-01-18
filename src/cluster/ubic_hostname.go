package cluster

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"sync"

	"github.com/boz/go-lifecycle"
	clustertypes "providerService/src/cluster/types/v1"

	"github.com/pkg/errors"
)

// hostnameID type exists to identify the target of a reservation. The lease ID type is not used directly because
// there is no need to consider order ID or provider ID for the purposes oft this
type ubicHostnameID struct {
	owner common.Address
	osep  uint64
}

func (hID ubicHostnameID) Equals(other ubicHostnameID) bool {
	return (hID.owner.String() == other.owner.String()) && (hID.osep == other.osep)
}

func ubicHostnameIDFromLeaseID(lID clustertypes.LeaseID) (ubicHostnameID, error) {
	ownerAddr := common.HexToAddress(lID.Owner)
	return ubicHostnameID{
		owner: ownerAddr,
		osep:  lID.OSeq,
	}, nil
}

type ubicSimpleHostnames struct {
	Hostnames map[string]ubicHostnameID
	lock      sync.Mutex
} /* Used in test code */

func NewUbicSimpleHostnames() clustertypes.HostnameServiceClient {
	return &ubicSimpleHostnames{
		Hostnames: make(map[string]ubicHostnameID),
	}
}

type UbicReservationResult struct {
	ChErr               <-chan error
	ChWithheldHostnames <-chan []string
}

func (rr UbicReservationResult) Wait(wait <-chan struct{}) ([]string, error) {
	select {
	case err := <-rr.ChErr:
		return nil, err
	case v := <-rr.ChWithheldHostnames:
		return v, nil
	case <-wait:
		return nil, errors.New("bob")
	}
}

func (sh *ubicSimpleHostnames) PrepareHostnamesForTransfer(ctx context.Context, hostnames []string, leaseID clustertypes.LeaseID) error {
	sh.lock.Lock()
	defer sh.lock.Unlock()
	errCh := make(chan error, 1)
	hID, err := ubicHostnameIDFromLeaseID(leaseID)
	if err != nil {
		return err
	}

	prepareUbicHostnamesImpl(sh.Hostnames, hostnames, hID, errCh)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}

func prepareUbicHostnamesImpl(store map[string]ubicHostnameID, hostnames []string, hID ubicHostnameID, errCh chan<- error) {
	toChange := make([]string, 0, len(hostnames))
	for _, hostname := range hostnames {
		existingID, ok := store[hostname]
		if ok {
			if existingID.owner.String() == hID.owner.String() {
				toChange = append(toChange, hostname)
			} else {
				errCh <- fmt.Errorf("%w: host %q in use", UbicErrHostnameNotAllowed, hostname)
				return
			}
		}
	}

	// Swap over each hostname
	for _, hostname := range toChange {
		store[hostname] = hID
	}
	errCh <- nil
}

func (sh *ubicSimpleHostnames) ReserveHostnames(ctx context.Context, hostnames []string, leaseID clustertypes.LeaseID) ([]string, error) {
	sh.lock.Lock()
	defer sh.lock.Unlock()
	errCh := make(chan error, 1)
	resultCh := make(chan []string, 1)

	hID, err := ubicHostnameIDFromLeaseID(leaseID)
	if err != nil {
		return nil, err
	}
	reserveUbicHostnamesImpl(sh.Hostnames, hostnames, hID, errCh, resultCh)

	select {
	case err := <-errCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func reserveUbicHostnamesImpl(store map[string]ubicHostnameID, hostnames []string, hID ubicHostnameID, ch chan<- error, resultCh chan<- []string) {
	withheldHostnamesMap := make(map[string]struct{})
	withheldHostnames := make([]string, 0)

	requestedHostnames := make(map[string]struct{})

	for _, hostname := range hostnames {
		requestedHostnames[hostname] = struct{}{}
		// Check if in use
		existingID, inUse := store[hostname]
		if inUse {
			// Check to see if the same address already is using this hostname
			if existingID.owner.String() != hID.owner.String() {
				// The owner is not the same, this can't be done
				ch <- fmt.Errorf("%w: host %q in use", UbicErrHostnameNotAllowed, hostname)
				return
			}

			// Check for a deployment replacing another one
			if !existingID.Equals(hID) {
				// Record that the hostname is being replaced
				withheldHostnames = append(withheldHostnames, hostname)
				withheldHostnamesMap[hostname] = struct{}{}
			}
		}
	}

	// Check to see if any hostnames that were previously in use by this ID
	// are no longer used
	removeHostnames := make([]string, 0)
	for hostname, existingID := range store {
		// Skip anything marked as in use still
		_, requested := requestedHostnames[hostname]
		if requested {
			continue
		}
		// If it is equal to this, add it to the list to be removed
		// it is no longer in use
		if existingID.Equals(hID) {
			removeHostnames = append(removeHostnames, hostname)
		}
	}

	// There was no error, mark everything as in use that is not withheld
	for _, hostname := range hostnames {
		_, withheld := withheldHostnamesMap[hostname]
		if !withheld {
			store[hostname] = hID
		}
	}

	// Remove everything that is no longer in use
	for _, removeHostname := range removeHostnames {
		delete(store, removeHostname)

	}

	resultCh <- withheldHostnames
}

func (sh *ubicSimpleHostnames) CanReserveHostnames(hostnames []string, ownerAddr common.Address) error {
	sh.lock.Lock()
	defer sh.lock.Unlock()
	ch := make(chan error, 1)
	canUbicReserveHostnamesImpl(sh.Hostnames, hostnames, ownerAddr, ch)
	return <-ch
}

func canUbicReserveHostnamesImpl(store map[string]ubicHostnameID, hostnames []string, ownerAddr common.Address, chErr chan<- error) {
	for _, hostname := range hostnames {
		existingID, inUse := store[hostname]

		if inUse {
			if existingID.owner.String() != ownerAddr.String() {
				chErr <- fmt.Errorf("%w: host %q in use", UbicErrHostnameNotAllowed, hostname)
				return
			}
		}
	}

	chErr <- nil
}

func (sh *ubicSimpleHostnames) ReleaseHostnames(leaseID clustertypes.LeaseID) error {
	sh.lock.Lock()
	defer sh.lock.Unlock()

	hID, err := ubicHostnameIDFromLeaseID(leaseID)
	if err != nil {
		return err
	}

	releaseUbicHostnamesImpl(sh.Hostnames, hID)
	return nil
}

func releaseUbicHostnamesImpl(store map[string]ubicHostnameID, hID ubicHostnameID) {
	var toDelete []string
	for hostname, existing := range store {
		if existing.Equals(hID) {
			toDelete = append(toDelete, hostname)
		}
	}

	for _, hostname := range toDelete {
		delete(store, hostname)
	}
}

type ubicReserveRequest struct {
	chErr               chan<- error
	chReplacedHostnames chan<- []string
	hostnames           []string
	hID                 ubicHostnameID
}

type ubicCanReserveRequest struct {
	hostnames []string
	result    chan<- error
	ownerAddr common.Address
}

type ubicPrepareTransferRequest struct {
	hostnames []string
	hID       ubicHostnameID
	chErr     chan<- error
}

type UbicHostnameService struct {
	inUse map[string]ubicHostnameID

	requests       chan ubicReserveRequest
	canRequest     chan ubicCanReserveRequest
	prepareRequest chan ubicPrepareTransferRequest
	releases       chan ubicHostnameID
	lc             lifecycle.Lifecycle

	blockedHostnames []string
	blockedDomains   []string
}

const UbicHostnameSeparator = '.'

func NewUbicHostnameService(ctx context.Context, cfg Config, initialData map[string]clustertypes.LeaseID) (*UbicHostnameService, error) {
	blockedHostnames := make([]string, 0)
	blockedDomains := make([]string, 0)
	for _, name := range cfg.BlockedHostnames {
		if len(name) != 0 && name[0] == UbicHostnameSeparator {
			blockedDomains = append(blockedDomains, name)
			blockedHostnames = append(blockedHostnames, name[1:])
		} else {
			blockedHostnames = append(blockedHostnames, name)
		}
	}

	hs := &UbicHostnameService{
		inUse:            make(map[string]ubicHostnameID, len(initialData)),
		blockedHostnames: blockedHostnames,
		blockedDomains:   blockedDomains,
		requests:         make(chan ubicReserveRequest),
		canRequest:       make(chan ubicCanReserveRequest),
		releases:         make(chan ubicHostnameID),
		lc:               lifecycle.New(),
		prepareRequest:   make(chan ubicPrepareTransferRequest),
	}
	for k, v := range initialData {
		hID, err := ubicHostnameIDFromLeaseID(v)
		if err != nil {
			return nil, err
		}
		hs.inUse[k] = hID
	}

	go hs.lc.WatchContext(ctx)
	go hs.run()

	return hs, nil
}

func (hs *UbicHostnameService) run() {
	defer hs.lc.ShutdownCompleted()

loop:
	for {

		// Wait for any service to finish
		select {
		case <-hs.lc.ShutdownRequest():
			hs.lc.ShutdownInitiated(nil)
			break loop
		case rr := <-hs.requests:
			reserveUbicHostnamesImpl(hs.inUse, rr.hostnames, rr.hID, rr.chErr, rr.chReplacedHostnames)
		case crr := <-hs.canRequest:
			canUbicReserveHostnamesImpl(hs.inUse, crr.hostnames, crr.ownerAddr, crr.result)
		case v := <-hs.releases:
			releaseUbicHostnamesImpl(hs.inUse, v)
		case request := <-hs.prepareRequest:
			prepareUbicHostnamesImpl(hs.inUse, request.hostnames, request.hID, request.chErr)

		}
	}

}

var UbicErrHostnameNotAllowed = errors.New("hostname not allowed")

func (hs *UbicHostnameService) PrepareHostnamesForTransfer(ctx context.Context, hostnames []string, leaseID clustertypes.LeaseID) error {
	chErr := make(chan error, 1)

	hID, err := ubicHostnameIDFromLeaseID(leaseID)
	if err != nil {
		return err
	}

	v := ubicPrepareTransferRequest{
		hostnames: hostnames,
		hID:       hID,
		chErr:     chErr,
	}
	select {
	case hs.prepareRequest <- v:
	case <-hs.lc.ShuttingDown():
		chErr <- ErrNotRunning
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case err = <-chErr:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-hs.lc.ShuttingDown():
		return ErrNotRunning
	}
}

func (hs *UbicHostnameService) isHostnameBlocked(hostname string) error {
	for _, blockedHostname := range hs.blockedHostnames {
		if blockedHostname == hostname {
			return fmt.Errorf("%w: %q is blocked by this provider", UbicErrHostnameNotAllowed, hostname)
		}
	}

	for _, blockedDomain := range hs.blockedDomains {
		if strings.HasSuffix(hostname, blockedDomain) {
			return fmt.Errorf("%w: domain %q is blocked by this provider", UbicErrHostnameNotAllowed, hostname)
		}
	}

	return nil
}

func (hs *UbicHostnameService) ReserveHostnames(ctx context.Context, hostnames []string, leaseID clustertypes.LeaseID) ([]string, error) {
	lowercaseHostnames := make([]string, len(hostnames))
	for i, hostname := range hostnames {
		lowercaseHostnames[i] = strings.ToLower(hostname)
	}

	// check if hostname is blocked
	for _, hostname := range lowercaseHostnames {
		blockedErr := hs.isHostnameBlocked(hostname)
		if blockedErr != nil {
			return nil, blockedErr
		}
	}

	chErr := make(chan error, 1)                  // Buffer of one so service does not block
	chWithheldHostnames := make(chan []string, 1) // Buffer of one so service does not block

	hID, err := ubicHostnameIDFromLeaseID(leaseID)

	if err != nil {
		return nil, err
	}
	request := ubicReserveRequest{
		chErr:               chErr,
		chReplacedHostnames: chWithheldHostnames,
		hostnames:           lowercaseHostnames,
		hID:                 hID,
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case hs.requests <- request:

	case <-hs.lc.ShuttingDown():
		return nil, ErrNotRunning
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-hs.lc.ShuttingDown():
		return nil, ErrNotRunning
	case err := <-chErr:
		return nil, err
	case result := <-chWithheldHostnames:
		return result, nil
	}
}

func (hs *UbicHostnameService) ReleaseHostnames(leaseID clustertypes.LeaseID) error {
	hID, err := ubicHostnameIDFromLeaseID(leaseID)
	if err != nil {
		return err
	}
	select {
	case hs.releases <- hID:
	case <-hs.lc.ShuttingDown():
		// service is shutting down, so release doesn't matter
	}
	return nil
}

func (hs *UbicHostnameService) CanReserveHostnames(hostnames []string, ownerAddr common.Address) error {
	returnValue := make(chan error, 1) // Buffer of one so service does not block
	lowercaseHostnames := make([]string, len(hostnames))
	for i, hostname := range hostnames {
		lowercaseHostnames[i] = strings.ToLower(hostname)
	}

	// check if hostname is blocked
	for _, hostname := range lowercaseHostnames {
		blockedErr := hs.isHostnameBlocked(hostname)
		if blockedErr != nil {
			return blockedErr
		}
	}

	request := ubicCanReserveRequest{ // do not actually reserve hostnames
		hostnames: lowercaseHostnames,
		result:    returnValue,
		ownerAddr: ownerAddr,
	}

	select {
	case hs.canRequest <- request:

	case <-hs.lc.ShuttingDown():
		returnValue <- ErrNotRunning
	}

	return <-returnValue
}

package cluster

import (
	"errors"
	"fmt"
)

var (
	// errReservationNotFound is the new error with message "not found"
	errReservationNotFound      = errors.New("reservation not found")
	errInventoryNotAvailableYet = errors.New("inventory status not available yet")
	errInventoryReservation     = errors.New("inventory error")
	errNoLeasedIPsAvailable     = fmt.Errorf("%w: no leased IPs available", errInventoryReservation)
	errInsufficientIPs          = fmt.Errorf("%w: insufficient number of IPs", errInventoryReservation)
)

var (
//inventoryRequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
//	Name:        "provider_inventory_requests",
//	Help:        "",
//	ConstLabels: nil,
//}, []string{"action", "result"})
//
//inventoryReservations = promauto.NewGaugeVec(prometheus.GaugeOpts{
//	Name: "provider_inventory_reservations_total",
//	Help: "",
//}, []string{"classification", "quantity"})
//
//clusterInventoryAllocateable = promauto.NewGaugeVec(prometheus.GaugeOpts{
//	Name: "provider_inventory_allocateable_total",
//	Help: "",
//}, []string{"quantity"})
//
//clusterInventoryAvailable = promauto.NewGaugeVec(prometheus.GaugeOpts{
//	Name: "provider_inventory_available_total",
//	Help: "",
//}, []string{"quantity"})
)

package types

import (
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"
)

// IPReservationDelete stores orderid
type IPReservationDelete struct {
	OrderID mtypes.OrderID
}

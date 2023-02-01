package ubicutil

import (
	"errors"
)

// Define error
var (
	ErrShuttingDown     = errors.New("shutting down")
	errServiceDiscovery = errors.New("service discovery failure")
	errServiceClient    = errors.New("service client failure")
)

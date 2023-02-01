package manifest

import "time"

// ServiceConfig is struct
type ServiceConfig struct {
	HTTPServicesRequireAtLeastOneHost bool
	ManifestTimeout                   time.Duration
	RPCQueryTimeout                   time.Duration
	CachedResultMaxAge                time.Duration
}

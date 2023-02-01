package operatorcommon

import (
	"time"
)

//OperatorConfig is struct
type OperatorConfig struct {
	PruneInterval      time.Duration
	WebRefreshInterval time.Duration
	RetryDelay         time.Duration
	ProviderAddress    string
}

package operatorcommon

import (
	"github.com/spf13/viper"
	"time"
)

// Define flag
const (
	FlagIgnoreListEntryLimit = "ignore-list-entry-limit"
	FlagIgnoreListAgeLimit   = "ignore-list-age-limit"
	FlagEventFailureLimit    = "event-failure-limit"
)

// IgnoreListConfig is struct
type IgnoreListConfig struct {
	// This is a config object, so it isn't exported as an interface
	FailureLimit uint
	EntryLimit   uint
	AgeLimit     time.Duration
}

// IgnoreListConfigFromViper is function get ignorelist from viper
func IgnoreListConfigFromViper() IgnoreListConfig {
	return IgnoreListConfig{
		FailureLimit: viper.GetUint(FlagEventFailureLimit),
		EntryLimit:   viper.GetUint(FlagIgnoreListEntryLimit),
		AgeLimit:     viper.GetDuration(FlagIgnoreListAgeLimit),
	}
}

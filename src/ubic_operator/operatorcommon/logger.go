package operatorcommon

import (
	// nolint: golint, staticcheck
	"github.com/go-kit/kit/log/term"
	"github.com/tendermint/tendermint/libs/log"
	"os"
)

//OpenLogger is logger create function
func OpenLogger() log.Logger {
	// logger with no color output - current debug colors are invisible for me.
	return log.NewTMLoggerWithColorFn(log.NewSyncWriter(os.Stdout), func(_ ...interface{}) term.FgBgColor {
		return term.FgBgColor{}
	})
}

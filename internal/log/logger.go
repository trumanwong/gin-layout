package log

import (
	"github.com/google/wire"
	"github.com/trumanwong/go-tools/log"
)

func NewLogger() *log.Logger {
	return log.NewLogger(nil)
}

var ProviderSet = wire.NewSet(NewLogger)

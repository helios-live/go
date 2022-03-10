package log

import (
	"fmt"

	"github.com/rs/zerolog"
)

// Logger is the bare minimum logging interface
type Logger interface {
	Debug(v ...interface{})
	Log(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

// Zero is a basic re-implementation of zerolog to match the logging interface
type Zero struct {
	zerolog.Logger
}

func (z Zero) Debug(v ...interface{}) { z.Logger.Debug().Msg(fmt.Sprint(v...)) }
func (z Zero) Info(v ...interface{})  { z.Logger.Info().Msg(fmt.Sprint(v...)) }
func (z Zero) Log(v ...interface{})   { z.Logger.Log().Msg(fmt.Sprint(v...)) }
func (z Zero) Warn(v ...interface{})  { z.Logger.Warn().Msg(fmt.Sprint(v...)) }
func (z Zero) Error(v ...interface{}) { z.Logger.Error().Msg(fmt.Sprint(v...)) }
func (z Zero) Fatal(v ...interface{}) { z.Logger.Fatal().Msg(fmt.Sprint(v...)) }

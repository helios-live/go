package log

import (
	"fmt"

	stdlog "log"

	"github.com/fatih/color"
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

// Std is a basic re-implementation of the standard library log to match the logging interface
type Std struct{}

func (s Std) Debug(v ...interface{}) { stdlog.Print("[DEBUG]", fmt.Sprint(v...)) }
func (s Std) Info(v ...interface{})  { stdlog.Print("[Info]", fmt.Sprint(v...)) }
func (s Std) Log(v ...interface{})   { stdlog.Print("[Log]", fmt.Sprint(v...)) }
func (s Std) Warn(v ...interface{})  { stdlog.Print("[Warn]", fmt.Sprint(v...)) }
func (s Std) Error(v ...interface{}) { stdlog.Print("[Error]", fmt.Sprint(v...)) }
func (s Std) Fatal(v ...interface{}) { stdlog.Print("[Fatal]", fmt.Sprint(v...)) }

// Color is a basic re-implementation of the standard library log to match the logging interface
// with added colors
type Color struct{}

func (Color) Debug(v ...interface{}) { stdlog.Print(color.GreenString("[DEBUG]"), fmt.Sprint(v...)) }
func (Color) Info(v ...interface{})  { stdlog.Print(color.HiBlueString("[Info]"), fmt.Sprint(v...)) }
func (Color) Log(v ...interface{})   { stdlog.Print(color.WhiteString("[Log]"), fmt.Sprint(v...)) }
func (Color) Warn(v ...interface{})  { stdlog.Print(color.YellowString("[Warn]"), fmt.Sprint(v...)) }
func (Color) Error(v ...interface{}) { stdlog.Print(color.RedString("[Error]"), fmt.Sprint(v...)) }
func (Color) Fatal(v ...interface{}) { stdlog.Print(color.HiRedString("[Fatal]"), fmt.Sprint(v...)) }

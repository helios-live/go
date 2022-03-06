package dev

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	er "github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

var (
	// ErrInternalServer global error values can be useful when wrapping errors or inspecting error types
	ErrInternalServer = er.New("error internal server")
)

const timeFormat = "2006-01-02 15:04:05"

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zl.Logger = zl.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat})

	log.SetFlags(0)
	log.SetOutput(zl.Logger)

	log.Println("asdf")

	zl.Info().
		Str("foo", "bar").
		Float64("f", 2.22).
		Msg("Hello world")

	zl.Debug().
		Str("foo", "bar").
		Float64("f", 2.22).
		Msg("Hello world")

	zl.Warn().
		Str("foo", "bar").
		Float64("f", 2.22).
		Msg("Hello world")

	// zl.Error().
	// 	Str("foo", "bar").
	// 	Float64("f", 2.22).
	// 	Msg("Hello world")

	err := er.New("error bad request")
	err = er.Wrapf(err, "this is a level 2 error '%v'", "oupss")
	err = er.Wrapf(err, "this is a level 3 error '%v'", "hopa")
	err = er.Wrapf(err, "this is a top level error '%v'", "hopa")

	Trace(err)

}

// Trace prints a pretty trace of the error
func Trace(err error) {
	uErr := er.Unpack(err)

	out := ""

	for i := len(uErr.ErrChain) - 1; i >= 0; i-- {
		e := uErr.ErrChain[i]
		out += perr(e.Msg, er.Stack{e.Frame}, len(uErr.ErrChain)-i)
	}
	out += perr(uErr.ErrRoot.Msg, uErr.ErrRoot.Stack, 0)

	fmt.Println(out)
}

func perr(msg string, st er.Stack, level int) string {

	out := ""

	if level == 0 {
		out += fmt.Sprintln("   " + color.RedString(msg))
	} else if level == 1 {
		out += fmt.Sprintln(time.Now().Format(timeFormat), color.HiRedString("Error: ")+color.YellowString(msg))
	} else {
		out += fmt.Sprintln("   " + color.HiBlackString(msg))
	}
	for _, s := range st {
		out += fmt.Sprintf("   %s: %s:%s\n", color.CyanString(s.Name), s.File, color.HiBlueString("%d", s.Line))
	}
	out += fmt.Sprintln()
	return out
}

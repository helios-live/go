package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/go-kit/log"
	"github.com/ideatocode/go/config"
	"github.com/ideatocode/go/config/pkg/flags"
	"github.com/ideatocode/go/config/pkg/marshal"
	"github.com/ideatocode/go/config/pkg/repository"
)

type mainConfig struct {
	*config.Default `json:"-" yaml:"-"`
	Token           string
	Counter         int
}

func main() {

	flag.Parse()
	logger := log.NewLogfmtLogger(os.Stderr)

	cfg := mainConfig{
		Default: config.New("data", "config"),
	}

	err := config.Load(&cfg)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "default", "Token", cfg.Token, "Counter", cfg.Counter)

	h := mainConfig{
		Default: config.New("data", "config", func(d *config.Default) {
			d.Repo = repository.HTTP{
				Token: "ZemExincRT6FgfvQWflCB8t1MTC8xOl4y1SwyAjmx7nl7WpdRzv0mZrgTr7nm0GJ",
				URL:   "https://peertonet.test/api/static-proxy/config",
			}
		}),
	}

	err = config.Load(&h)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "http", "Token", h.Token, "Counter", h.Counter)

	y := mainConfig{
		Default: config.New("data", "config", func(d *config.Default) {
			d.Repo = repository.File{
				Extension: ".yaml",
			}
			d.Marsh = marshal.YAML{}
		}),
	}

	err = config.Load(&y)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "yaml", "Token", y.Token, "Counter", y.Counter)

	// increase the counter
	y.Counter++
	config.Save(y)

	wrap := mainConfig{
		Default: flags.Wrap(config.New("data", "config")),
	}

	err = config.Load(&wrap)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "flags wrapped", "Token", wrap.Token, "Counter", wrap.Counter)

}

package flags

import (
	"flag"

	"github.com/ideatocode/go/config"
	"github.com/ideatocode/go/config/pkg/marshal"
	"github.com/ideatocode/go/config/pkg/repository"
)

var repo = flag.String("repository", "", "Forcefully use this repository [file, http]")
var marsh = flag.String("marshaler", "", "Forcefully use this marshaler [jsonc, yaml]")
var httpSource = flag.String("http-ds", "", "Forcefully use HTTP source to get and save the config from")
var httpToken = flag.String("http-auth", "", "Forcefully use HTTP Authorization: Bearer Token Header")

// Wrap returns a new config.Default with the options changed by the flags
func Wrap(cfg *config.Default) *config.Default {
	switch *repo {
	case "file":
		cfg.Repo = repository.File{}

	case "http":
		cfg.Repo = repository.HTTP{
			Token: *httpToken,
			URL:   *httpSource,
		}
	}

	switch *marsh {
	case "jsonc":
		cfg.Marsh = marshal.JSONC{}

	case "yaml":
		cfg.Marsh = marshal.YAML{}
	}
	return cfg
}

package config

import (
	"sync"

	"github.com/go-kit/log"

	"github.com/ideatocode/go/config/pkg/interfaces"
	"github.com/ideatocode/go/config/pkg/marshal"
	"github.com/ideatocode/go/config/pkg/repository"
)

// New returns a pointer to a filled new instance of Configuration
func New(group, item string, options ...func(*Default)) *Default {
	cfg := &Default{
		Mutex: &sync.Mutex{},
		Group: group,
		Item:  item,
	}

	for _, option := range options {
		option(cfg)
	}

	// use the fileRepository by default
	if cfg.Repo == nil {
		cfg.Repo = repository.File{
			Extension: ".jsonc",
		}
	}

	// use the fileRepository by default
	if cfg.Marsh == nil {
		cfg.Marsh = marshal.JSONC{}
	}

	// use the fileRepository by default
	if cfg.Log == nil {
		cfg.Log = log.NewNopLogger()
	}

	return cfg
}

// Save .
func Save(config interfaces.Config) error {

	return config.Repository().Save(config)
}

// Load .
func Load(config interfaces.Config) error {
	return config.Repository().Load(config)
}

// LoadConfig loads the config and return the fiilled object
//
// Deprecated: LoadConfig is deprecated use Load instead
//
func LoadConfig(config interfaces.Config) error {

	return Load(config)
}

// Sync Writes the config to disk
//
// Deprecated: Sync is deprecated use Save instead
//
func Sync(config interfaces.Config) {
	Save(config)
}

// Sync Writes the config to disk
//
// Deprecated: Sync is deprecated
//
func (c *Default) Sync() {
	c.Repository().Save(c)
}

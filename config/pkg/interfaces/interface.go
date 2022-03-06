package interfaces

import (
	"sync"

	"github.com/go-kit/log"
)

// Config .
// must be able to Save options,
// must be able to Load options,
// must be able to use multiple "drivers"
// must be thread safe
// must be able to use multiple logical/physical configs
type Config interface {

	// ConfigurationInterface needs to implement Locker
	sync.Locker

	Repository() Repository

	Marshaler() Marshaler

	Logger() log.Logger

	// Path should return the path to save to
	// for jsonc the resulting file will be Path.json
	// for http the path will be appended to the url as a query string parameter
	Path() string
}

// Repository is responsible for interacting with the storage system
type Repository interface {

	// Load is called to load options
	Load(Config) error

	// Save is called to save options
	Save(Config) error
}

// Marshaler is responsible for serializing and deserializing data
type Marshaler interface {
	Marshal(v interface{}) (data []byte, err error)
	Unmarshal(data []byte, v interface{}) error
}

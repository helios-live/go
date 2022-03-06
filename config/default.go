package config

import (
	"sync"

	"github.com/go-kit/log"
	"github.com/ideatocode/go/config/pkg/interfaces"
)

// Default is the base Default object
type Default struct {
	*sync.Mutex `json:"-" yaml:"-"`

	// Group represents a logical grouping
	// for JSON it'll be used as the folder
	// for Sql it'll be used as the database
	Group string `json:"-" yaml:"-"`

	// Item represents this whole configuration object
	// for JSON it'll be the file name
	// for Sql it'll be the table name
	Item string `json:"-" yaml:"-"`

	// Repo is responsible for interracting with the storage medium
	Repo interfaces.Repository `json:"-" yaml:"-"`

	// Marsh is responsible for serializing and deserializing data
	Marsh interfaces.Marshaler `json:"-" yaml:"-"`

	Log log.Logger `json:"-" yaml:"-"`
}

// Repository returns the repository
func (c Default) Repository() interfaces.Repository {
	return c.Repo
}

// Marshaler returns the repository
func (c Default) Marshaler() interfaces.Marshaler {
	return c.Marsh
}

// Logger returns the logger
func (c Default) Logger() log.Logger {
	return c.Log
}

// Path returns the path of the current config
func (c Default) Path() string {
	return c.Group + "/" + c.Item
}

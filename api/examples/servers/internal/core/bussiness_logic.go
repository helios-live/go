package core

import (
	"errors"

	"github.com/go-kit/kit/log"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("s: empty string")

// APIService s
type APIService struct {
	logger log.Logger
}

// New returns a new APIService instance
func (as APIService) New(logger log.Logger) APIService {
	return APIService{
		logger: logger,
	}
}

// Servers returns a list
func (as APIService) Servers(s string) ([]string, error) {
	if s == "" {
		return []string{}, ErrEmpty
	}
	return []string{"a", "b", "c"}, nil
}

// Server returns a list
func (as APIService) Server(sid string) ([]string, error) {
	if sid == "" {
		return []string{}, ErrEmpty
	}
	return []string{sid}, nil
}

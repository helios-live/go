package authenticator

import (
	"fmt"
	"net/http"
)

// Error represents an authorization error.
type Error struct {
	Details string
}

// StatusCode is an implementation of the StatusCoder interface in go-kit/http.
func (ae Error) StatusCode() int {
	return http.StatusUnauthorized
}

// Error is an implementation of the Error interface.
func (ae Error) Error() string {
	return fmt.Sprintf("%s(%s)", http.StatusText(http.StatusUnauthorized), ae.Details)
}

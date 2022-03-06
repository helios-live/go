package api

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/ideatocode/go/api/transport"
)

// Endpoint is one apiEndpoint
type Endpoint interface {
	Entry() endpoint.Endpoint
	Decoder() transport.DecodeFunc
	// Decoder() httptransport.DecodeRequestFunc
	Path() string
}

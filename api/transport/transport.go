package transport

import (
	"context"
	"io"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Transport .
type Transport interface {
	Register(ep endpoint.Endpoint, dec httptransport.DecodeRequestFunc, path string) interface{}
	Decode(df DecodeFunc) httptransport.DecodeRequestFunc
}

// DecodeFunc is used to decode the raw request information
type DecodeFunc func(ctx context.Context, rc io.ReadCloser, vars map[string]interface{}) (interface{}, error)

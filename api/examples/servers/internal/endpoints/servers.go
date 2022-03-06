package endpoints

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/ideatocode/go/api/api"
	"github.com/ideatocode/go/api/examples/servers/internal/core"
	"github.com/ideatocode/go/api/transport"
)

type listServers struct {
	svc core.APIService
}
type listServersRequest struct {
	S string `json:"s"`
}
type listServersResponse struct {
	V   []string `json:"v"`
	Err string   `json:"err,omitempty"` // errors don't define JSON marshaling
}

func (l listServersResponse) StatusCode() int {
	if l.Err != "" {
		return http.StatusInternalServerError
	}
	return 200
}

// ListServersEndpoint returns the api endpoint responsible for listing the servers
func ListServersEndpoint(svc core.APIService) api.Endpoint {
	return &listServers{svc: svc}
}

func (l listServers) Entry() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(listServersRequest)
		v, err := l.svc.Servers(req.S)
		if err != nil {
			return listServersResponse{v, err.Error()}, nil
		}
		return listServersResponse{v, ""}, nil
	}
}
func (l listServers) Decoder() transport.DecodeFunc {
	return func(ctx context.Context, rc io.ReadCloser, vars map[string]interface{}) (interface{}, error) {
		var request listServersRequest

		if err := json.NewDecoder(rc).Decode(&request); err != nil {
			return nil, err
		}

		return request, nil
	}
}

func (l listServers) Path() string {
	return "/servers"
}

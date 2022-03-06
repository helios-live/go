package endpoints

import (
	"context"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/ideatocode/go/api/api"
	"github.com/ideatocode/go/api/examples/servers/internal/core"
	"github.com/ideatocode/go/api/transport"
)

type getServer struct {
	svc core.APIService
}
type getServerRequest struct {
	ServerID string `json:"ServerID"`
}
type getServerResponse struct {
	V   []string `json:"v"`
	Err string   `json:"err,omitempty"` // errors don't define JSON marshaling
}

func (l getServerResponse) StatusCode() int {
	if l.Err != "" {
		return http.StatusInternalServerError
	}
	return 200
}

// GetServerEndpoint returns the api endpoint responsible for listing the servers
func GetServerEndpoint(svc core.APIService) api.Endpoint {
	return &getServer{svc: svc}
}

func (l getServer) Entry() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(getServerRequest)
		v, err := l.svc.Server(req.ServerID)
		if err != nil {
			return getServerResponse{v, err.Error()}, nil
		}
		return getServerResponse{v, ""}, nil
	}
}
func (l getServer) Decoder() transport.DecodeFunc {
	return func(ctx context.Context, rc io.ReadCloser, vars map[string]interface{}) (interface{}, error) {

		request := getServerRequest{
			ServerID: vars["ServerID"].(string),
		}

		return request, nil
	}
}

func (l getServer) Path() string {
	return "/servers/{ServerID:[a-zA-Z0-9]+}"
}

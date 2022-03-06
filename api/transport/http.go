package transport

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/ideatocode/go/api/authenticator"
)

type httpTransport struct {
	srv    *mux.Router
	logger log.Logger
	auth   authenticator.Interface
}

// NewHTTP returns a new httpTransport
func NewHTTP(srv *mux.Router, logger log.Logger, auth authenticator.Interface) Transport {

	return &httpTransport{
		srv:    srv,
		logger: logger,
		auth:   auth,
	}

}

// Register
func (ht *httpTransport) Register(ep endpoint.Endpoint, dec httptransport.DecodeRequestFunc, path string) interface{} {

	ep = authMiddleware(ht.auth)(ep)
	ep = loggingMiddleware(log.With(ht.logger, "endpoint", path))(ep)

	handler := httptransport.NewServer(
		ep,
		dec,
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	)
	return ht.srv.Handle(path, handler)
}

// Decode
func (ht *httpTransport) Decode(df DecodeFunc) httptransport.DecodeRequestFunc {

	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		m2 := make(map[string]interface{}, len(vars))

		for k, v := range vars {
			m2[k] = v
		}

		return df(ctx, r.Body, m2)
	}
}

package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/ideatocode/go/api/api"
	"github.com/ideatocode/go/api/authenticator"
	"github.com/ideatocode/go/api/examples/servers/internal/core"
	"github.com/ideatocode/go/api/examples/servers/internal/endpoints"
	"github.com/ideatocode/go/api/transport"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	svc := core.APIService{}
	svc = svc.New(logger)

	authenticator := authenticator.HTTPBearerToken(
		// authenticator.NopTokenRepository{Ret: false},
		tokenRepository{},
	)

	mux := mux.NewRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	transport := transport.NewHTTP(mux, logger, authenticator)

	i := api.New()
	i.Add(endpoints.ListServersEndpoint(svc))
	i.Add(endpoints.GetServerEndpoint(svc))

	i.Use(transport)

	srv.ListenAndServe()
}

type tokenRepository struct{}

func (tr tokenRepository) Check(token string) (identity interface{}, err error) {
	if token != "test" {
		return "", errors.New("Token not found")
	}
	return "admin", nil
}

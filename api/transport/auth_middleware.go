package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ideatocode/go/api/authenticator"
)

func authMiddleware(auth authenticator.Interface) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {

			if _, err := auth.Check(ctx); err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}

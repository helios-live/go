package transport

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			logger.Log("msg", "calling endpoint")
			defer func(begin time.Time) {
				logger.Log(
					"msg", "called endpoint",
					"took", time.Since(begin),
					"err", err,
				)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

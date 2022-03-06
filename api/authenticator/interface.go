package authenticator

import "context"

// Interface is used to authenticate incoming API Requests
type Interface interface {
	Check(context.Context) (identity interface{}, err error)
}

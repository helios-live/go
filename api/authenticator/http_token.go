package authenticator

import (
	"context"
	"errors"
	"fmt"

	httptransport "github.com/go-kit/kit/transport/http"
)

// TokenRepository is used to check the validity of tokens when token authentication is used
type TokenRepository interface {
	Check(token string) (identity interface{}, err error)
}

// HTTPToken is used for authentication whenever Tokens are used
type HTTPToken struct {
	repository TokenRepository
}

// HTTPBearerToken returns a new http token authenticator that taps into the TokenRepository to check tokens
func HTTPBearerToken(tr TokenRepository) *HTTPToken {
	return &HTTPToken{
		repository: tr,
	}
}

// Check checks the underlying repository for the current authorization bearer
// token and returns an identity, error touple
func (ht HTTPToken) Check(ctx context.Context) (identity interface{}, err error) {

	auth, ok := ctx.Value(httptransport.ContextKeyRequestAuthorization).(string)
	if !ok {
		return nil, Error{"Could not get Authorization header value"}
	}

	if auth[0:7] != "Bearer " {
		return nil, Error{"Only Authorization: Bearer is allowed"}
	}

	auth = auth[7:]

	if identity, err = ht.repository.Check(auth); err != nil {
		return nil, Error{fmt.Sprintf("Token \"%s\" not recognized e: %s", auth, err)}
	}

	return identity, nil
}

// NopTokenRepository is a token repository that does no actual checking of the tokens
type NopTokenRepository struct {
	Ret bool
}

// Check returns n.Ret
func (n NopTokenRepository) Check(token string) (interface{}, error) {
	if !n.Ret {
		return nil, errors.New("Nop: false")
	}
	return nil, nil
}

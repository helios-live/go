package authenticator

import (
	"context"
	"errors"
)

type Nop struct {
	Ret bool
}

func (n Nop) Check(ctx context.Context) (interface{}, error) {
	if !n.Ret {
		return nil, errors.New("Nop: false")
	}
	return nil, nil
}

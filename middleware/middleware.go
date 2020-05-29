package middleware

import (
	"github.com/csiabb/donation-service/context"
)

type Middleware struct {
	ctx *context.Context
}

func NewMiddleware(c *context.Context) (*Middleware, error) {
	return &Middleware{ctx: c}, nil
}

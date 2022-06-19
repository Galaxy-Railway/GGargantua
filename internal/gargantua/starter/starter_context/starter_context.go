package starter_context

import "context"

type StarterContext interface {
	GetContext() context.Context
}

type starterContext struct {
	context.Context
}

func (c *starterContext) GetContext() context.Context {
	return c.Context
}

var CancelAllServer func()

func NewServerContext() (StarterContext, error) {
	ctx, cancel := context.WithCancel(context.Background())
	CancelAllServer = cancel
	return &starterContext{ctx}, nil
}

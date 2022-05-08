package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
)

type RequestService interface {
	SendRequest(ctx context.Context, request *module.Request) (*module.Response, error)
}

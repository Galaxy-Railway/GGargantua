package sender

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
)

type ISender interface {
	SendOnce(ctx context.Context, request interface{}) (*module.AllResponse, error)
}

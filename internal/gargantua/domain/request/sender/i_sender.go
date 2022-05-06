package sender

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"

type ISender interface {
	SendOnce(request interface{}) (*module.AllResponse, error)
}

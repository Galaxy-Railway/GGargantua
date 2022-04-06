package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"

type RequestService interface {
	SendSingleRequest(request *module.Request) (*module.Response, error)
}

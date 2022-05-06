package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"

type RequestService interface {
	SendRequest(request *module.Request) (*module.Response, error)
}

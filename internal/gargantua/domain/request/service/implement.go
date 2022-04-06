package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
)

type RequestServiceImpl struct {
}

func NewRequestService() RequestService {
	return &RequestServiceImpl{}
}

func (s *RequestServiceImpl) SendSingleRequest(request *module.Request) (*module.Response, error) {
	var sender module.ISender
	switch request.RequestSchema {
	case module.HTTP:
		sender = module.NewHttpSender()
	case module.HTTPS:

	}

	resBytes, err := sender.SendOnce(request.RequestContent)
	if err != nil {
		return nil, err
	}
	return &module.Response{
		ResponseSchema:  request.RequestSchema,
		ResponseContent: resBytes,
	}, nil
}

package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
)

type RequestServiceImpl struct {
}

func NewRequestService() RequestService {
	return &RequestServiceImpl{}
}

func (s *RequestServiceImpl) SendRequest(request *module.Request) (*module.Response, error) {
	var (
		sender  module.ISender
		content interface{}
	)

	switch request.RequestSchema {
	case module.HTTP:
		sender = module.HttpSender()
		content = request.HttpRequest
	case module.HTTPS:
		sender = module.HttpsSender()
		content = request.HttpsRequest
	}

	// todo: make goroutine pool, do parallel request
	_, err := sender.SendOnce(content)
	if err != nil {
		return nil, err
	}
	return &module.Response{
		ResponseSchema: request.RequestSchema,
		//todo: complete response
		ResponseContent: nil,
	}, nil
}

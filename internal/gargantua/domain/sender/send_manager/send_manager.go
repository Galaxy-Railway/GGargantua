package send_manager

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/module"
	"sync"
)

type SendManager struct {
}

var (
	once = &sync.Once{}

	sendManager *SendManager
)

func NewSendManager() *SendManager {
	once.Do(func() {
		sendManager = &SendManager{}
	})
	return sendManager
}

func (s *SendManager) SendSingleRequest(request *module.Request) (*module.Response, error) {
	var sender module.Sender
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

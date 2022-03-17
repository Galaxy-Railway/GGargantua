package module

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/module"

func BuildSenderBySingleRequest(sr *SingleRequest) *module.Sender{
	return &module.Sender{
		Req: &module.Request{
			RequestSchema:  sr.RequestSchema,
			RequestContent: sr.RequestContent,
		},
		Res: &module.Response{},
	}
}

func GetSingleResponseFromSender(sender *module.Sender) *SingleResponse{
	return &SingleResponse{
		ResponseSchema:  sender.Res.ResponseSchema,
		ResponseContent: sender.Res.ResponseContent,
	}
}

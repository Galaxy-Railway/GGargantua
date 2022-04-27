package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type RequestStepType struct {
	Request *module.Request

	Result *StepResult
}

func TransRequestStepTypeByPb(stepType *protobuf.RequestStepType) *RequestStepType {
	if stepType == nil {
		return nil
	}
	return &RequestStepType{
		Request: module.TransRequestByPb(stepType.Request),
	}
}

func (f *RequestStepType) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

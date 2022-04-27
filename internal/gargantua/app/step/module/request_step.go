package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
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

func (f *RequestStepType) Execute(ctx context.Context) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

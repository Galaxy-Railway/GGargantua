package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type ScriptStepType struct {
	Script string

	Result *StepResult
}

func TransScriptStepByDb(stepType *protobuf.ScriptStepType) *ScriptStepType {
	if stepType == nil {
		return nil
	}
	return &ScriptStepType{
		Script: stepType.Script,
	}
}

func (f *ScriptStepType) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

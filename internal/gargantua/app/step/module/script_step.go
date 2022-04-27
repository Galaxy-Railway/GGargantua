package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
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

func (f *ScriptStepType) Execute(ctx context.Context) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
)

type IfStepType struct {
	ScriptStep *ScriptStepType
	TrueStep   *Step
	FalseStep  *Step

	Result *StepResult
}

func TransIfStepTypeByPb(stepType *protobuf.IfStepType) *IfStepType {
	if stepType == nil {
		return nil
	}
	return &IfStepType{
		ScriptStep: TransScriptStepByDb(stepType.ScriptStep),
		TrueStep:   TransStepFromPB(stepType.TrueStep),
		FalseStep:  TransStepFromPB(stepType.FalseStep),
	}
}

func (f *IfStepType) Execute(ctx context.Context) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
)

type ForStepType struct {
	Times   int
	SubStep *Step

	Result *StepResult
}

func (f *ForStepType) Execute(ctx context.Context) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

func TransForStepTypeByPb(stepType *protobuf.ForStepType) *ForStepType {
	if stepType == nil {
		return nil
	}
	return &ForStepType{
		Times:   int(stepType.Times),
		SubStep: TransStepFromPB(stepType.SubStep),
	}
}

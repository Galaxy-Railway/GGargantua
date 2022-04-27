package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
)

type SequenceStepType struct {
	SubSteps []*Step

	Result *StepResult
}

func TransSequenceStepTypeByPb(stepType *protobuf.SequenceStepType) *SequenceStepType {
	if stepType == nil {
		return nil
	}
	subSteps := make([]*Step, len(stepType.SubSteps))
	for i := range subSteps {
		subSteps[i] = TransStepFromPB(stepType.SubSteps[i])
	}
	return &SequenceStepType{
		SubSteps: subSteps,
	}
}

func (f *SequenceStepType) Execute(ctx context.Context) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

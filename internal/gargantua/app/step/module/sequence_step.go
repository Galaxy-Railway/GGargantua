package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
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

func (f *SequenceStepType) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type ForStepType struct {
	Times   int
	SubStep *Step

	Result *StepResult
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

func (f *ForStepType) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	//TODO implement me
	panic("implement me")
}

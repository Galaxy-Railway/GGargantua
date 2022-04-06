package multiple_steps

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
)

func TranferSteps(steps *step_pb.MultiSteps) []*module.Step {
	l := len(steps.Steps)
	result := make([]*module.Step, l)
	for i, s := range steps.GetSteps() {
		result[i] = &module.Step{
			Type:     module.TypeOfStep(int(s.TypeOfStep)),
			Content:  s.Content,
			Executor: nil,
		}
	}
	return result
}

func TransferStepResult(result *module.StepResult) *step_pb.StepResult {
	return &step_pb.StepResult{
		Result:    result.Result,
		SubResult: TransFerStepResultsSlice(result.SubStepResult),
	}
}

func TransFerStepResultsSlice(result []*module.StepResult) []*step_pb.StepResult {
	l := len(result)
	pbResult := make([]*step_pb.StepResult, l)
	for i, r := range result {
		pbResult[i] = TransferStepResult(r)
	}
	return pbResult
}

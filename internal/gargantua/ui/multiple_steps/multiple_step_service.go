package multiple_steps

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
)

type MultipleStepServiceUI struct {
	Multi service.MultipleStepServiceApp
}

func NewMultipleStepServiceUI(s service.MultipleStepServiceApp) step_pb.MultipleStepServiceServer {
	return &MultipleStepServiceUI{Multi: s}
}

func (m *MultipleStepServiceUI) ExecuteMultipleStep(ctx context.Context, steps *step_pb.MultiSteps) (*step_pb.MultiResults, error) {
	appSteps := TranferSteps(steps)
	appResult, err := m.Multi.ExecuteSteps(appSteps, ctx)
	if err != nil {
		multiple_steps.Logger().Errorf("failed to ExecuteMultipleStep, err: %v", err)
		return nil, err
	}
	pbResult := TransFerStepResultsSlice(appResult)
	return &step_pb.MultiResults{
		Results: pbResult,
	}, nil
}

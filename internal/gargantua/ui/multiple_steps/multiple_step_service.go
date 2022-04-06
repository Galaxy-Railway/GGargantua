package multiple_steps

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
)

type MultipleStepService struct {
	Multi service.MultipleSteps
}

func NewMultipleStepService(s service.MultipleSteps) step_pb.MultipleStepServiceServer {
	return &MultipleStepService{Multi: s}
}

func (m *MultipleStepService) ExecuteMultipleStep(ctx context.Context, steps *step_pb.MultiSteps) (*step_pb.MultiResults, error) {
	appSteps := TranferSteps(steps)
	appResult, err := m.Multi.ExecuteSteps(appSteps)
	if err != nil {
		common.Logger().Errorf("failed to ExecuteMultipleStep, err: %v", err)
		return nil, err
	}
	pbResult := TransFerStepResultsSlice(appResult)
	return &step_pb.MultiResults{
		Results: pbResult,
	}, nil
}

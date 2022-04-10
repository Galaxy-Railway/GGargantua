package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MultipleStepsImpl struct {
	logger         *zap.SugaredLogger
	requestService request.RequestService
}

func (m *MultipleStepsImpl) ExecuteSteps(steps []*module.Step) ([]*module.StepResult, error) {
	l := len(steps)
	result := make([]*module.StepResult, l)
	for i, s := range steps {
		r, err := s.MarshalAndExecute()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute step %d/%d, content: %d", i, l, s.Content)
		}
		result[i] = r
	}
	return result, nil
}

func NewMultipleSteps(logger *zap.SugaredLogger, rs request.RequestService) MultipleSteps {
	return &MultipleStepsImpl{
		logger:         logger,
		requestService: rs,
	}
}
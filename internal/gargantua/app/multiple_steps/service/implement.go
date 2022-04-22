package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MultipleStepServiceAppImpl struct {
	logger         *zap.SugaredLogger
	requestService request.RequestService
	scriptService  script.ScriptService
}

func (m *MultipleStepServiceAppImpl) ExecuteSteps(steps []*module.Step, ctx context.Context) ([]*module.StepResult, error) {
	l := len(steps)
	result := make([]*module.StepResult, l)
	for i, s := range steps {
		select {
		case <-ctx.Done():
			// todo : deal with cancel
			return nil, common.CanceledError
		default:
			err := s.Marshal(m.requestService, m.scriptService)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to execute step %d/%d, content: %d", i, l, s.Content)
			}
			r, err := s.Execute(ctx)
			if err != nil {
				return nil, err
			}
			result[i] = r
		}
	}
	return result, nil
}

func NewMultipleStepsServiceApp(logger *zap.SugaredLogger, rs request.RequestService, ss script.ScriptService) MultipleStepServiceApp {
	return &MultipleStepServiceAppImpl{
		logger:         logger,
		requestService: rs,
		scriptService:  ss,
	}
}

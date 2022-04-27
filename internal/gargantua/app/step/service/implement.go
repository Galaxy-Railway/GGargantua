package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
	"go.uber.org/zap"
)

type StepServiceAppImpl struct {
	logger         *zap.SugaredLogger
	requestService request.RequestService
	scriptService  script.ScriptService
}

func (m *StepServiceAppImpl) ExecuteSteps(step *module.Step, ctx context.Context) (*module.StepResult, error) {
	return step.Execute(ctx)
}

func NewStepsServiceApp(logger *zap.SugaredLogger, rs request.RequestService, ss script.ScriptService) StepServiceApp {
	return &StepServiceAppImpl{
		logger:         logger,
		requestService: rs,
		scriptService:  ss,
	}
}

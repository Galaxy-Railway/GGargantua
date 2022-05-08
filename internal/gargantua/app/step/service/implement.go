package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type StepServiceAppImpl struct {
	requestService request.RequestService
	scriptService  script.ScriptService
}

func (m *StepServiceAppImpl) ExecuteStep(step *module.Step, ctx context.Context) (*module.StepResult, error) {
	return step.Execute(ctx, m.requestService, m.scriptService)
}

func NewStepsServiceApp(rs request.RequestService, ss script.ScriptService) StepServiceApp {
	return &StepServiceAppImpl{
		requestService: rs,
		scriptService:  ss,
	}
}

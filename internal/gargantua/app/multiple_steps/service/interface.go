package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
)

type MultipleStepServiceApp interface {
	ExecuteSteps([]*module.Step, context.Context) ([]*module.StepResult, error)
}

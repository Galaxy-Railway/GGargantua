package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
)

type StepServiceApp interface {
	ExecuteStep(*module.Step, context.Context) (*module.StepResult, error)
}

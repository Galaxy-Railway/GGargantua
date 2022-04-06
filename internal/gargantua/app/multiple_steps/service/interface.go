package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"

type MultipleSteps interface {
	ExecuteSteps([]*module.Step) ([]*module.StepResult, error)
}

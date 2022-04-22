package app

import (
	jobService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/service"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps"
	multiStepService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InjectAppLayer(container *dig.Container) error {
	if err := InjectLoggers(container); err != nil {
		return errors.Wrap(err, "failed to inject loggers in app layer")
	}
	if err := InjectService(container); err != nil {
		return errors.Wrap(err, "failed to inject services in app layer")
	}
	return nil
}

func InjectLoggers(container *dig.Container) error {
	if err := container.Invoke(multiple_steps.InjectLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of multiple_steps")
	}
	return nil
}

func InjectService(container *dig.Container) error {
	if err := container.Provide(multiStepService.NewMultipleStepsServiceApp); err != nil {
		return errors.Wrap(err, "failed to provide NewMultipleStepsServiceApp()")
	}

	if err := container.Provide(jobService.NewJobServiceApp); err != nil {
		return errors.Wrap(err, "failed to provide NewJobServiceApp()")
	}
	return nil
}

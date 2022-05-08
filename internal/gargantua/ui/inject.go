package ui

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/job"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InjectUiLayer(container *dig.Container) error {

	if err := InjectLoggers(container); err != nil {
		return errors.Wrap(err, "failed to inject loggers in ui layer")
	}

	if err := InvokeRegister(container); err != nil {
		return errors.Wrap(err, "failed to invoke register in ui layer")
	}
	return nil
}

func InvokeRegister(container *dig.Container) error {
	return container.Invoke(RegisterAllGrpc)
}

func InjectLoggers(container *dig.Container) error {
	if err := container.Invoke(job.InjectLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of job")
	}
	return nil
}

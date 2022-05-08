package domain

import (
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request"
	requestService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script"
	scriptService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InjectDomainLayer(container *dig.Container) error {
	if err := InjectService(container); err != nil {
		return errors.Wrap(err, "failed to inject service in layer domain")
	}
	if err := InjectLoggers(container); err != nil {
		return errors.Wrap(err, "failed to inject loggers in layer domain")
	}
	return nil
}

func InjectService(container *dig.Container) error {
	if err := container.Provide(requestService.NewRequestService); err != nil {
		return errors.Wrap(err, "failed to inject NewRequestService")
	}
	if err := container.Provide(scriptService.NewScriptService); err != nil {
		return errors.Wrap(err, "failed to inject NewScriptService")
	}
	return nil
}

func InjectLoggers(container *dig.Container) error {
	if err := container.Invoke(request.InjectLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of request")
	}
	if err := container.Invoke(script.InjectLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of script")
	}
	return nil
}

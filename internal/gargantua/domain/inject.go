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
		return errors.Wrap(err, "failed to inject logger of step")
	}
	if err := container.Invoke(script.InjectLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of step")
	}
	return nil
}

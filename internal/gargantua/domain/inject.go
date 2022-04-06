package domain

import (
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InjectDomainLayer(container *dig.Container) error {
	if err := container.Provide(request.NewRequestService); err != nil {
		return errors.Wrap(err, "failed to inject NewRequestService")
	}
	if err := container.Provide(script.NewScriptService); err != nil {
		return errors.Wrap(err, "failed to inject NewScriptService")
	}
	return nil
}

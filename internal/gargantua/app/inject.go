package app

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InjectAppLayer(container *dig.Container) error {
	if err := container.Provide(service.NewMultipleSteps); err != nil {
		return errors.Wrap(err, "failed to provide NewMultipleSteps()")
	}
	return nil
}

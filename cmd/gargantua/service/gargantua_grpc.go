package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

// GargantuaGrpc is main server starter of gargantua grpc service
func GargantuaGrpc() (err error) {
	// inject all layers of gargantua
	err = InjectLayers(container)
	if err != nil {
		logger.Errorf("inject layers failed, err: %v", err)
		return err
	}

	// start grpc server
	err = starter.InvokeGrpc(container)
	if err != nil {
		logger.Errorf("invoke grpc server failed, err: %v", err)
		return err
	}
	return nil
}

func InjectLayers(container *dig.Container) error {
	if err := starter.InjectInfraLayer(container); err != nil {
		return err
	}
	if err := domain.InjectDomainLayer(container); err != nil {
		return err
	}
	if err := app.InjectAppLayer(container); err != nil {
		return err
	}
	if err := ui.InjectUiLayer(container); err != nil {
		return err
	}
	return nil
}

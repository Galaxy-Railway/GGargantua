package service

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

// GargantuaGrpc is main server starter of gargantua
func GargantuaGrpc(configFile string) {
	common.SetConfigPath(configFile)

	// start a dig container
	container := dig.New()

	// generate config
	err := container.Provide(common.NewConfig)
	if err != nil {
		fmt.Printf("Get config failed! error: %v\nconfig path: %s\n", err, configFile)
	}

	// generate basic(global) logger
	err = container.Provide(common.NewLogger)
	if err != nil {
		fmt.Printf("provide logger failed! error: %v\n", err)
	}

	// invoke makes logger initiated immediately
	err = container.Invoke(common.LoggerInjectTrigger)
	if err != nil {
		fmt.Printf("trigger logger init failed! error: %v\n", err)
	}

	// set a logger for main function
	logger = common.GlobalLogger.Named("main")
	defer logger.Sync()
	logger.Info("get config success")
	logger.Info("init logger success")

	// inject all layers of gargantua
	err = InjectLayers(container)
	if err != nil {
		logger.Fatalf("inject layers failed, err: %v", err)
	}

	// start grpc server
	err = starter.InvokeGrpc(container)
	if err != nil {
		logger.Fatalf("invoke grpc server failed, err: %v", err)
	}
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

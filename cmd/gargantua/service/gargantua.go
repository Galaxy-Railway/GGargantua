package service

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain"
	GGrpc "github.com/Galaxy-Railway/GGargantua/internal/gargantua/infra/grpc"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

// Gargantua is main server starter of gargantua
func Gargantua(configFile string) {
	common.SetConfigPath(configFile)
	container := dig.New()
	err := container.Provide(common.NewConfig)
	if err != nil {
		fmt.Printf("Get config failed! error: %v\nconfig path: %s\n", err, configFile)
	}
	err = container.Provide(common.NewLogger)
	if err != nil {
		fmt.Printf("provide logger failed! error: %v\n", err)
	}
	err = container.Invoke(common.LoggerInjectTrigger)
	if err != nil {
		fmt.Printf("trigger logger init failed! error: %v\n", err)
	}

	logger = common.GlobalLogger.Named("main")
	defer logger.Sync()
	logger.Info("get config success")
	logger.Info("init logger success")
	err = GGrpc.RegisterGrpcServer(container)
	if err != nil {
		logger.Fatalf("get grpc server failed, err: %v", err)
	}

	err = InjectLayers(container)
	if err != nil {
		logger.Fatalf("inject layers failed, err: %v", err)
	}

	err = GGrpc.InvokeGrpc(container)
	if err != nil {
		logger.Fatalf("invoke grpc server failed, err: %v", err)
	}
}

func InjectLayers(container *dig.Container) error {
	if err := domain.InjectDomainLayer(container); err != nil {
		return err
	}
	if err := app.InjectAppLayer(container); err != nil {
		return err
	}
	if err := ui.InvokeRegister(container); err != nil {
		return err
	}
	return nil
}

package service

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
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
	err = container.Invoke(common.NewLogger)
	if err != nil {
		fmt.Printf("Get logger failed! error: %v\n", err)
	}
	logger = common.Logger().Named("main")
	defer logger.Sync()
	logger.Info("get config success")
	logger.Info("init logger success")
	err = container.Provide(NewGrpcServer)
	if err != nil {
		logger.Fatalf("get grpc server failed, err: %v", err)
	}

	err = InjectLayers(container)
	if err != nil {
		logger.Fatalf("inject layers failed, err: %v", err)
	}

	err = container.Invoke(StartGargantuaGrpcServer)
	if err != nil {
		logger.Fatalf("start grpc server failed, err: %v", err)
	}
}

// StartGargantuaGrpcServer will start a grpc server of gargantua
func StartGargantuaGrpcServer(config *common.ProjectConfig, server *grpc.Server) error {
	lis, err := net.Listen(config.Listen.Network, config.Listen.Address)
	if err != nil {
		return err
	}
	err = server.Serve(lis)
	return err
}

func NewGrpcServer() *grpc.Server {
	return grpc.NewServer()
}

func InjectLayers(container *dig.Container) error {
	if err := domain.InjectDomainLayer(container); err != nil {
		return err
	}
	if err := app.InjectAppLayer(container); err != nil {
		return err
	}
	return nil
}
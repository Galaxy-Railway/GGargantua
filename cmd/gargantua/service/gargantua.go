package service

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var (
	logger *zap.SugaredLogger
)

// Gargantua is main server starter of gargantua
func Gargantua(configFile string) {
	config, err := common.GetConfig(configFile)
	if err != nil {
		fmt.Printf("failed to get config from %s, error: %v", configFile, err)
	}
	err = common.InitLogger(config.Log)
	if err != nil {
		fmt.Printf("Get logger failed! error: %v\nconfig: %v\n", err, config)
	}
	logger = common.Logger().Named("main")
	defer logger.Sync()
	logger.Info("get config success")
	logger.Info("init logger success")
	grpcServer := GargantuaGrpcServer(logger)
	lis, err := net.Listen(config.Listen.Network, config.Listen.Address)
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatalf("start grpc server failed, err: %v", err)
	}
}

// GargantuaGrpcServer will start a grpc server of gargantua
func GargantuaGrpcServer(logger *zap.SugaredLogger) *grpc.Server {
	server := grpc.NewServer()
	ui.RegisterAllGrpc(server, logger)
	return server
}

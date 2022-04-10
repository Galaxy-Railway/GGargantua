package grpc

import (
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"google.golang.org/grpc"
	"net"
)

// StartGargantuaGrpcServer will start a grpc server of gargantua
func StartGargantuaGrpcServer(config *common.ProjectConfig, server *grpc.Server) error {
	lis, err := net.Listen(config.Listen.Network, config.Listen.Address)
	if err != nil {
		return err
	}
	err = server.Serve(lis)
	return err
}

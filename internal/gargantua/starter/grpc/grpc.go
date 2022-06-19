package grpc

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/starter_context"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"google.golang.org/grpc"
	"net"
)

func NewGrpcServer() *grpc.Server {
	return grpc.NewServer(SetGrpcInterceptor()...)
}

// StartGargantuaGrpcServer will start a grpc server of gargantua
func StartGargantuaGrpcServer(config *common.ProjectConfig, server *grpc.Server, ctx starter_context.StarterContext) error {
	var lc net.ListenConfig
	lis, err := lc.Listen(ctx.GetContext(), config.Grpc.Network, fmt.Sprintf("%s:%d", config.Grpc.Host, config.Grpc.Port))
	if err != nil {
		return err
	}
	err = server.Serve(lis)
	return err
}

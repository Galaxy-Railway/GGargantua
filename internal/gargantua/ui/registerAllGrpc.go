package ui

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/request_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/single_request"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterAllGrpc(server *grpc.Server, logger *zap.SugaredLogger){
	request_pb.RegisterSingleRequestSenderServer(server, single_request.NewSingleRequestSenderImpl(logger))
}

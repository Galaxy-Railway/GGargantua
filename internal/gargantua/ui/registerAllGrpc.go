package ui

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	jobService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/service"
	service2 "github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/job/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterAllGrpc(
	server *grpc.Server,
	logger *zap.SugaredLogger,
	jobApp jobService.JobServiceApp) {
	protobuf.RegisterJobServiceServer(server, service2.NewJobServiceUI(jobApp, logger))
}

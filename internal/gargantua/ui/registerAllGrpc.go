package ui

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	jobService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/service"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/job"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/multiple_steps"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterAllGrpc(
	server *grpc.Server,
	logger *zap.SugaredLogger,
	multiApp service.MultipleStepServiceApp,
	jobApp jobService.JobServiceApp) {
	step_pb.RegisterMultipleStepServiceServer(server, multiple_steps.NewMultipleStepServiceUI(multiApp))
	step_pb.RegisterJobServiceServer(server, job.NewJobServiceUI(jobApp, logger))
}

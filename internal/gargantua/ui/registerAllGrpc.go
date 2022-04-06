package ui

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/multiple_steps"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterAllGrpc(
	server *grpc.Server,
	logger *zap.SugaredLogger,
	multiApp service.MultipleSteps) {
	step_pb.RegisterMultipleStepServiceServer(server, multiple_steps.NewMultipleStepService(multiApp))
}

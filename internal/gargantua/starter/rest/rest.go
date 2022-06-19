package rest

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/starter_context"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func StartGargantuaRestServer(config *common.ProjectConfig, ctx starter_context.StarterContext) error {
	grpcEndPoint := fmt.Sprintf("localhost:%d", config.Grpc.Port)
	restEndPoint := fmt.Sprintf("%s:%d", config.Rest.Host, config.Rest.Port)

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := protobuf.RegisterJobServiceHandlerFromEndpoint(ctx.GetContext(), mux, grpcEndPoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(restEndPoint, mux)
}

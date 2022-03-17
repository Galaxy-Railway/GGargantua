package single_request

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/request_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/single_request/module"
	domainModule "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/module"
)

func TransportSingleRequest(input *request_pb.SingleRequest) *module.SingleRequest{
	return &module.SingleRequest{
		RequestSchema:  domainModule.SchemaType(input.RequestSchema),
		RequestContent: input.RequestContent,
	}
}

func TransportSingleResponse(input *module.SingleResponse) *request_pb.SingleResponse{
	return &request_pb.SingleResponse{
		ResponseSchema:  request_pb.SchemaType(input.ResponseSchema),
		ResponseContent: input.ResponseContent,
	}
}

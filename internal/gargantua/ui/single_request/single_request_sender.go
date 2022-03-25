package single_request

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/request_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/single_request"
	"go.uber.org/zap"
)

var _ request_pb.SingleRequestSenderServer = &SingleRequestSenderImpl{}

type SingleRequestSenderImpl struct {
	logger *zap.SugaredLogger
}

func NewSingleRequestSenderImpl(logger *zap.SugaredLogger) request_pb.SingleRequestSenderServer {
	return &SingleRequestSenderImpl{
		logger: logger.Named("SingleRequestSender"),
	}
}

func (s SingleRequestSenderImpl) SendSingleRequest(_ context.Context, request *request_pb.SingleRequest) (*request_pb.SingleResponse, error) {
	req := TransportSingleRequest(request)
	res, err := single_request.SendSingleRequest(req)
	if err != nil {
		s.logger.Errorf("send a single request error, err: %v", err)
		return nil, err
	}
	return TransportSingleResponse(res), nil
}

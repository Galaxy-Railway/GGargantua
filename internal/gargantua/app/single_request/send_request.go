package single_request

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/single_request/module"
	domainModule "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/send_manager"
	"github.com/pkg/errors"
)

func SendSingleRequest(sRequest *module.SingleRequest) (*module.SingleResponse, error) {
	sendManager := send_manager.NewSendManager()
	req := &domainModule.Request{
		RequestSchema:  sRequest.RequestSchema,
		RequestContent: sRequest.RequestContent,
	}
	res, err := sendManager.SendSingleRequest(req)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to send singleRequest, schema: %d", sRequest.RequestSchema)
	}
	return &module.SingleResponse{
		ResponseSchema:  res.ResponseSchema,
		ResponseContent: res.ResponseContent,
	}, nil
}

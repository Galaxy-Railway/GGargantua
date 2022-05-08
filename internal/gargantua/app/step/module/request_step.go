package module

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type RequestStepType struct {
	Request *module.Request

	Result *StepResult
}

func TransRequestStepTypeByPb(stepType *protobuf.RequestStepType) *RequestStepType {
	if stepType == nil {
		return nil
	}
	return &RequestStepType{
		Request: module.TransRequestFromPb(stepType.Request),
	}
}

func (f *RequestStepType) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	stepResult := &StepResult{}
	// todo: send ctx
	result, err := requestService.SendRequest(ctx, f.Request)
	if err != nil {
		stepResult.Success = false
		stepResult.Reason = err.Error()
	} else {
		stepResult.Success = true
		stepResult.Reason = "request send successful"
	}
	
	bytes, marshalErr := json.Marshal(result)
	if marshalErr != nil {
		stepResult.Success = false
		stepResult.Reason = fmt.Sprintf("got response: %+v, \nbut failed to marshal into json, err: %+v", result, marshalErr)
	} else {
		stepResult.Result = string(bytes)
	}
	return stepResult, err
}

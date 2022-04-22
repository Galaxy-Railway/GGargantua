package module

import (
	"context"
	"encoding/json"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"time"
)

type RequestStepExecutor struct {
	RequestSchema   module.SchemaType
	RequestContent  []byte
	ResponseContent []byte
	Times           int
	Interval        time.Duration
	requestService  request.RequestService
}

func (r *RequestStepExecutor) Execute(ctx context.Context) (*StepResult, error) {
	for count := 0; count < r.Times; count++ {
		select {
		case <-ctx.Done():
			return nil, common.CanceledError
		default:
			// todo: execute
			multiple_steps.Logger().Debugf("request %d/%d times", count, r.Times)
		}
	}
	return &StepResult{Result: []byte{}}, nil
}

func BuildRequestStepExecutor(bytes []byte, rs request.RequestService) (StepExecutor, error) {
	exec := &RequestStepExecutor{}
	exec.requestService = rs
	err := json.Unmarshal(bytes, exec)
	return exec, err
}

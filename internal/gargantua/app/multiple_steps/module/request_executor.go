package module

import (
	"encoding/json"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"time"
)

type RequestStepExecutor struct {
	RequestSchema   module.SchemaType
	RequestContent  []byte
	ResponseContent []byte
	Times           int
	Interval        time.Duration
}

func (r *RequestStepExecutor) Execute() (*StepResult, error) {
	for count := 0; count < r.Times; count++ {
		common.Logger().Debugf("request %d/%d times", count, r.Times)
		// todo: send multiple request
	}
	return nil, nil
}

func BuildRequestStepExecutor(str string) (*RequestStepExecutor, error) {
	exec := &RequestStepExecutor{}
	err := json.Unmarshal([]byte(str), exec)
	return exec, err
}

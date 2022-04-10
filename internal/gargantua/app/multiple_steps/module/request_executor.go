package module

import (
	"encoding/json"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
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
		multiple_steps.Logger().Debugf("request %d/%d times", count, r.Times)
		// todo: send multiple request
	}
	return &StepResult{Result: []byte{}}, nil
}

func BuildRequestStepExecutor(str string) (*RequestStepExecutor, error) {
	exec := &RequestStepExecutor{}
	err := json.Unmarshal([]byte(str), exec)
	return exec, err
}

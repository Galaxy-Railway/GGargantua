package module

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type TypeOfStep int

const (
	Unknown = iota
	RequestStep
	ScriptStep
	ForStep
	IfStep
)

type Step struct {
	Type     TypeOfStep
	Content  []byte
	Executor StepExecutor
}

var (
	UnknownStepTypeError = errors.New("unknown step type")
)

func (s *Step) MarshalAndExecute() (*StepResult, error) {
	switch s.Type {
	case RequestStep:
		tmp := &RequestStepExecutor{}
		if err := json.Unmarshal(s.Content, tmp); err != nil {
			return nil, errors.Wrapf(err, "failed to marshal step, content: %s", string(s.Content))
		}
		s.Executor = tmp
	default:
		return nil, errors.Wrapf(UnknownStepTypeError, "got step type which is not supported: %s", s.Type)
	}
	return s.Executor.Execute()
}

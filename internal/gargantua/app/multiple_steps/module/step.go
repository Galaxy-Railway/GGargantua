package module

import (
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
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

func (s *Step) MarshalAndExecute(rs request.RequestService, ss script.ScriptService) (*StepResult, error) {
	switch s.Type {
	case RequestStep:
		exec, err := BuildRequestStepExecutor(s.Content, rs)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to marshal step, content: %s", string(s.Content))
		}
		s.Executor = exec
	default:
		return nil, errors.Wrapf(UnknownStepTypeError, "got step type which is not supported: %s", s.Type)
	}
	return s.Executor.Execute()
}

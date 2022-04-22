package module

import (
	"context"
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

func (s *Step) Marshal(rs request.RequestService, ss script.ScriptService) error {
	switch s.Type {
	case RequestStep:
		exec, err := BuildRequestStepExecutor(s.Content, rs)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal step, content: %s", string(s.Content))
		}
		s.Executor = exec
	default:
		return errors.Wrapf(UnknownStepTypeError, "got step type which is not supported: %s", s.Type)
	}
	return nil
}

func (s *Step) Execute(ctx context.Context) (*StepResult, error) {
	return s.Executor.Execute(ctx)
}

package module

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
	"github.com/pkg/errors"
)

type StepType int

const (
	StepTypeUnknown = iota
	StepTypeRequest
	StepTypeScript
	StepTypeFor
	StepTypeIf
	StepTypeSequence
)

type Step struct {
	Type StepType

	RequestStep  *RequestStepType
	ScriptStep   *ScriptStepType
	ForStep      *ForStepType
	IfStep       *IfStepType
	SequenceStep *SequenceStepType

	Result *StepResult
}

func TransStepFromPB(step *protobuf.Step) *Step {
	if step == nil {
		return nil
	}
	return &Step{
		Type:         StepType(int(step.Type)),
		RequestStep:  TransRequestStepTypeByPb(step.RequestStep),
		ScriptStep:   TransScriptStepByDb(step.ScriptStep),
		ForStep:      TransForStepTypeByPb(step.ForStep),
		IfStep:       TransIfStepTypeByPb(step.IfStep),
		SequenceStep: TransSequenceStepTypeByPb(step.SequenceStep),
	}
}

var (
	UnknownStepTypeError = errors.New("unknown step type")
)

func (s *Step) Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error) {
	switch s.Type {
	case StepTypeRequest:
		return s.RequestStep.Execute(ctx, requestService, scriptService)
	case StepTypeScript:
		return s.ScriptStep.Execute(ctx, requestService, scriptService)
	case StepTypeFor:
		return s.ForStep.Execute(ctx, requestService, scriptService)
	case StepTypeIf:
		return s.IfStep.Execute(ctx, requestService, scriptService)
	case StepTypeSequence:
		return s.SequenceStep.Execute(ctx, requestService, scriptService)
	case StepTypeUnknown:
		return nil, UnknownStepTypeError
	default:
		return nil, UnknownStepTypeError
	}
}

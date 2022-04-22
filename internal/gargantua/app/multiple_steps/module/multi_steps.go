package module

import (
	"context"
)

type MultiSteps struct {
	Steps   []*Step
	Context context.Context
	Results []*StepResult
}

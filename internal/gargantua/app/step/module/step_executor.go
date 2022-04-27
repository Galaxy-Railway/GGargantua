package module

import "context"

type StepExecutor interface {
	Execute(ctx context.Context) (*StepResult, error)
}

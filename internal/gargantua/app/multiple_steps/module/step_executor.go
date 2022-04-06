package module

type StepExecutor interface {
	Execute() (*StepResult, error)
}

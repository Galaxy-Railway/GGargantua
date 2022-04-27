package module

type StepResult struct {
	Result        string
	SubStepResult []*StepResult
}

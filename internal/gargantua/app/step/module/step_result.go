package module

type StepResult struct {
	Success       bool
	Reason        string
	Result        string
	SubStepResult []*StepResult
}

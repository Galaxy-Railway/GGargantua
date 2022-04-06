package module

type StepResult struct {
	Result        []byte
	SubStepResult []*StepResult
}

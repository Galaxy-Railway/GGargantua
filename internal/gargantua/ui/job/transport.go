package job

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	module2 "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
)

func TranferSteps(steps *step_pb.MultiSteps) []*module.Step {
	l := len(steps.Steps)
	result := make([]*module.Step, l)
	for i, s := range steps.GetSteps() {
		result[i] = &module.Step{
			Type:     module.TypeOfStep(int(s.TypeOfStep)),
			Content:  []byte(s.Content),
			Executor: nil,
		}
	}
	return result
}

func TransStringToUuid(id string) *step_pb.JobUuid {
	return &step_pb.JobUuid{Uuid: id}
}

func TransUuidToString(id *step_pb.JobUuid) string {
	return id.Uuid
}

func TransferStepResult(result *module.StepResult) *step_pb.StepResult {
	return &step_pb.StepResult{
		Result:    result.Result,
		SubResult: TransferStepResultsSlice(result.SubStepResult),
	}
}

func TransferStepResultsSlice(result []*module.StepResult) []*step_pb.StepResult {
	l := len(result)
	pbResult := make([]*step_pb.StepResult, l)
	for i, r := range result {
		pbResult[i] = TransferStepResult(r)
	}
	return pbResult
}

func TransferMultiStepResults(result []*module.StepResult) *step_pb.MultiResults {
	l := len(result)
	pbResult := make([]*step_pb.StepResult, l)
	for i, r := range result {
		pbResult[i] = TransferStepResult(r)
	}
	return &step_pb.MultiResults{
		Results: pbResult,
	}
}

func TransferUpdateJobContent(content *step_pb.UpdateJobContent) (string, []*module.Step) {
	return content.Uuid.Uuid, TranferSteps(content.Steps)
}

func GetJobStatus(job *module2.Job) *step_pb.JobStatus {
	return &step_pb.JobStatus{
		Uuid:     job.Uuid,
		Progress: step_pb.JobProgress(int(job.Status)),
		Reason:   job.Reason,
	}
}

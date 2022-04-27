package transport

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	module2 "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
)

func TransStringToUuid(id string) *protobuf.JobUuid {
	return &protobuf.JobUuid{Uuid: id}
}

func TransUuidToString(id *protobuf.JobUuid) string {
	return id.Uuid
}

func TransferStepResult(result *module.StepResult) *protobuf.StepResult {
	return &protobuf.StepResult{
		Success:   result.Success,
		Reason:    result.Reason,
		Result:    result.Result,
		SubResult: TransferStepResultsSlice(result.SubStepResult),
	}
}

func TransferStepResultsSlice(result []*module.StepResult) []*protobuf.StepResult {
	l := len(result)
	pbResult := make([]*protobuf.StepResult, l)
	for i, r := range result {
		pbResult[i] = TransferStepResult(r)
	}
	return pbResult
}

func TransferUpdateJobContent(content *protobuf.UpdateJobContent) (string, *module.Step) {
	return content.Uuid.Uuid, module.TransStepFromPB(content.MainStep)
}

func GetJobStatus(job *module2.Job) *protobuf.JobStatus {
	return &protobuf.JobStatus{
		Uuid:     job.Uuid,
		Progress: protobuf.JobProgress(int(job.Status)),
		Reason:   job.Reason,
	}
}

func GetJobResult(job *module2.Job) *protobuf.JobResult {
	return &protobuf.JobResult{
		Status: GetJobStatus(job),
		Result: TransferStepResult(job.MainStep.Result),
	}
}

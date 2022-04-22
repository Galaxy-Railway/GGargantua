package job

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/step_pb"
	jobService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobServiceUI struct {
	JobServiceApp jobService.JobServiceApp
	logger        *zap.SugaredLogger
}

func NewJobServiceUI(jobServiceApp jobService.JobServiceApp, logger *zap.SugaredLogger) step_pb.JobServiceServer {
	return &JobServiceUI{JobServiceApp: jobServiceApp, logger: logger.Named("JobServiceUI")}
}

func (j JobServiceUI) CreateAJob(_ context.Context, _ *emptypb.Empty) (*step_pb.JobUuid, error) {
	job := j.JobServiceApp.CreateJob()
	return TransStringToUuid(job.Uuid), nil
}

func (j JobServiceUI) StartAJob(_ context.Context, content *step_pb.UpdateJobContent) (*step_pb.JobStatus, error) {
	uuid, steps := TransferUpdateJobContent(content)
	job, err := j.JobServiceApp.StartAJob(uuid, steps)
	if err != nil {
		j.logger.Errorf("failed to start a job, err: %+v", err)
	}
	return GetJobStatus(job), nil
}

func (j JobServiceUI) CancelAJob(ctx context.Context, uuid *step_pb.JobUuid) (*step_pb.JobStatus, error) {
	job, err := j.JobServiceApp.CancelAJob(TransUuidToString(uuid))
	if err != nil {
		j.logger.Errorf("failed to cancel a job, err: %+v", err)
	}
	return GetJobStatus(job), nil
}

func (j JobServiceUI) GetJobResult(ctx context.Context, uuid *step_pb.JobUuid) (*step_pb.MultiResults, error) {
	result, err := j.JobServiceApp.GetJobResult(TransUuidToString(uuid))
	if err != nil {
		j.logger.Errorf("failed to result of a job, err: %+v", err)
	}
	return TransferMultiStepResults(result), nil
}

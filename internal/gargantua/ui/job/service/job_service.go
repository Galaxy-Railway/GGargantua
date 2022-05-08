package service

import (
	"context"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	jobService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/service"
	job2 "github.com/Galaxy-Railway/GGargantua/internal/gargantua/ui/job/transport"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobServiceUI struct {
	JobServiceApp jobService.JobServiceApp
	logger        *zap.SugaredLogger
}

func NewJobServiceUI(jobServiceApp jobService.JobServiceApp, logger *zap.SugaredLogger) protobuf.JobServiceServer {
	return &JobServiceUI{JobServiceApp: jobServiceApp, logger: logger.Named("ui | jobService")}
}

func (j JobServiceUI) CreateAJob(_ context.Context, _ *emptypb.Empty) (*protobuf.JobUuid, error) {
	job := j.JobServiceApp.CreateJob()
	return job2.TransStringToUuid(job.Uuid), nil
}

func (j JobServiceUI) StartAJob(_ context.Context, content *protobuf.UpdateJobContent) (*emptypb.Empty, error) {
	uuid, step := job2.TransferUpdateJobContent(content)
	_, err := j.JobServiceApp.StartAJob(uuid, step)
	return &emptypb.Empty{}, err
}

func (j JobServiceUI) CancelAJob(ctx context.Context, uuid *protobuf.JobUuid) (*emptypb.Empty, error) {
	_, err := j.JobServiceApp.CancelAJob(job2.TransUuidToString(uuid))
	if err != nil {
		j.logger.Errorf("failed to cancel a job, err: %+v", err)
	}
	return nil, err
}

func (j JobServiceUI) GetJobResult(ctx context.Context, uuid *protobuf.JobUuid) (*protobuf.JobResult, error) {
	job, err := j.JobServiceApp.GetAJob(job2.TransUuidToString(uuid))
	if err != nil {
		j.logger.Errorf("failed to result of a job, err: %+v", err)
	}
	return job2.GetJobResult(job), nil
}

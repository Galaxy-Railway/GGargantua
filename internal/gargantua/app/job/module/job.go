package module

import (
	"context"
	"errors"
	multiStep "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"github.com/google/uuid"
	"time"
)

type Job struct {
	Uuid   string
	Status JobStatus
	Reason string

	CtxCancel context.CancelFunc

	MultiSteps *multiStep.MultiSteps

	CreateTime time.Time
	StartTime  time.Time
	EndTime    time.Time
}

func NewJob() *Job {
	u := uuid.New()
	return &Job{
		Uuid:       u.String(),
		Status:     CREATED,
		Reason:     "job created",
		CreateTime: time.Now(),
	}
}

func (j *Job) GoJob(ctx context.Context, msService service.MultipleStepServiceApp) {
	go func() {
		j.Status = PROCESSING
		result, err := msService.ExecuteSteps(j.MultiSteps.Steps, ctx)
		if err != nil {
			if errors.Is(err, common.CanceledError) {
				j.Status = CANCELED
				j.Reason = "this job was canceled"
			} else {
				j.Status = FAILED
				j.Reason = err.Error()
			}
		} else {
			j.Status = FINISHED
			j.Reason = "this job was finished"
			j.MultiSteps.Results = result
		}
		j.EndTime = time.Now()
	}()
}

func (j *Job) Cancel() {
	j.Status = CANCELED
	j.CtxCancel()
	j.EndTime = time.Now()
	j.Reason = "job canceled"
}

type JobStatus int

const (
	CREATED JobStatus = iota
	PROCESSING
	FINISHED
	FAILED

	CANCELING
	CANCELED
)

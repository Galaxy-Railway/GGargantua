package module

import (
	"context"
	"errors"
	StepModule "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"github.com/google/uuid"
	"time"
)

type Job struct {
	Uuid   string
	Status JobProgress
	Reason string

	CtxCancel context.CancelFunc

	MainStep *StepModule.Step

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

func (j *Job) GoJob(ctx context.Context, msService service.StepServiceApp) {
	go func() {
		j.Status = PROCESSING
		result, err := msService.ExecuteSteps(j.MainStep, ctx)
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
			j.MainStep.Result = result
		}
		j.EndTime = time.Now()
	}()
}

func (j *Job) Cancel() {
	j.Status = CANCELING
	j.CtxCancel()
	j.Status = CANCELED
	j.EndTime = time.Now()
	j.Reason = "job canceled"
}

type JobProgress int

const (
	CREATED JobProgress = iota
	PROCESSING
	FINISHED
	FAILED

	CANCELING
	CANCELED
)

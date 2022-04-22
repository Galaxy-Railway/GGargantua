package service

import (
	"context"
	"errors"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/module"
	stepModule "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
	stepService "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/service"
	"sync"
	"time"
)

type JobServiceAppImpl struct {
	Jobs             map[string]*module.Job
	MultiStepService stepService.MultipleStepServiceApp
}

var (
	once = &sync.Once{}

	JobNotFoundError = errors.New("uuid not found in job service")
	StepsNotSetError = errors.New("steps have not set to this job")
)

func NewJobServiceApp(mss stepService.MultipleStepServiceApp) JobServiceApp {
	return &JobServiceAppImpl{
		Jobs:             make(map[string]*module.Job),
		MultiStepService: mss,
	}
}

func (j JobServiceAppImpl) CreateJob() *module.Job {
	return module.NewJob()
}

func (j JobServiceAppImpl) StartAJob(uu string, steps []*stepModule.Step) (*module.Job, error) {
	ctx, cancel := context.WithCancel(context.Background())
	job, ok := j.Jobs[uu]
	if !ok {
		cancel()
		return nil, JobNotFoundError
	}
	job.MultiSteps = &stepModule.MultiSteps{
		Steps:   steps,
		Context: ctx,
	}
	job.CtxCancel = cancel
	job.StartTime = time.Now()
	job.GoJob(ctx, j.MultiStepService)
	return job, nil
}

func (j JobServiceAppImpl) CancelAJob(uu string) (*module.Job, error) {
	job, ok := j.Jobs[uu]
	if !ok {
		return nil, JobNotFoundError
	}
	job.Cancel()
	return job, nil
}

func (j JobServiceAppImpl) GetJobResult(uu string) ([]*stepModule.StepResult, error) {
	job, ok := j.Jobs[uu]
	if !ok {
		return nil, JobNotFoundError
	}
	if job.MultiSteps == nil {
		return nil, StepsNotSetError
	}
	return job.MultiSteps.Results, nil
}

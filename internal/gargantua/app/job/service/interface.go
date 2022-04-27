package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/module"
	stepModule "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/step/module"
)

type JobServiceApp interface {
	CreateJob() *module.Job
	StartAJob(uuid string, steps *stepModule.Step) (*module.Job, error)
	CancelAJob(uuid string) (*module.Job, error)
	GetJobResult(uuid string) (*stepModule.StepResult, error)
	GetAJob(uu string) (*module.Job, error)
}

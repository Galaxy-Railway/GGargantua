package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/job/module"
	multiStep "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/multiple_steps/module"
)

type JobServiceApp interface {
	CreateJob() *module.Job
	StartAJob(uuid string, steps []*multiStep.Step) (*module.Job, error)
	CancelAJob(uuid string) (*module.Job, error)
	GetJobResult(uuid string) ([]*multiStep.StepResult, error)
}

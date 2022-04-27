package module

import (
	"context"
	request "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/service"
	script "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/service"
)

type StepExecutor interface {
	Execute(ctx context.Context, requestService request.RequestService, scriptService script.ScriptService) (*StepResult, error)
}

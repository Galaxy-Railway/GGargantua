package starter

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/grpc"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/rest"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/starter_context"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InvokeGrpc(container *dig.Container) error {
	return container.Invoke(grpc.StartGargantuaGrpcServer)
}

func InvokeRest(container *dig.Container) error {
	return container.Invoke(rest.StartGargantuaRestServer)
}

func InjectInfraLayer(container *dig.Container) error {
	if err := InjectLoggers(container); err != nil {
		return errors.Wrap(err, "failed to inject loggers in infra layer")
	}
	if err := InjectGrpcServer(container); err != nil {
		return errors.Wrap(err, "failed to inject grpcServer in infra layer")
	}
	return nil
}

func InjectGrpcServer(container *dig.Container) error {
	if err := container.Provide(grpc.NewGrpcServer); err != nil {
		return err
	}
	return nil
}

func InjectLoggers(container *dig.Container) error {
	if err := container.Invoke(grpc.InitLogger); err != nil {
		return errors.Wrap(err, "failed to inject logger of grpc")
	}
	return nil
}

func InjectContext(container *dig.Container) error {
	if err := container.Provide(starter_context.NewServerContext); err != nil {
		return errors.Wrap(err, "failed to provide context of starters")
	}
	return nil
}

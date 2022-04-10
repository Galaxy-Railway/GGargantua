package grpc

import (
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(container *dig.Container) error {
	if err := container.Invoke(InitLogger); err != nil {
		return err
	}
	if err := container.Provide(NewGrpcServer); err != nil {
		return err
	}
	return nil
}

func InvokeGrpc(container *dig.Container) error {
	return container.Invoke(StartGargantuaGrpcServer)
}

func NewGrpcServer() *grpc.Server {
	return grpc.NewServer(SetGrpcInceptor()...)
}

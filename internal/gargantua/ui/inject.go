package ui

import "go.uber.org/dig"

func InvokeRegister(container *dig.Container) error {
	return container.Invoke(RegisterAllGrpc)
}

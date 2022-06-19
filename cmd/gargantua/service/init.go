package service

import (
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"go.uber.org/dig"
)

var (
	container *dig.Container
)

func Init(configPath string) error {
	common.SetConfigPath(configPath)
	// start a dig container
	container = dig.New()

	// generate config
	err := container.Provide(common.NewConfig)
	if err != nil {
		fmt.Printf("Get config failed! error: %v\nconfig path: %s\n", err, configPath)
		return err
	}

	// generate basic(global) logger
	err = container.Provide(common.NewLogger)
	if err != nil {
		fmt.Printf("provide logger failed! error: %v\n", err)
		return err
	}

	// invoke makes logger initiated immediately
	err = container.Invoke(common.LoggerInjectTrigger)
	if err != nil {
		fmt.Printf("trigger logger init failed! error: %v\n", err)
		return err
	}

	err = starter.InjectContext(container)
	if err != nil {
		fmt.Printf("inject context failed! error: %v\n", err)
		return err
	}

	// set a logger for main function
	logger = common.GlobalLogger.Named("init")
	defer logger.Sync()
	logger.Info("get config success")
	logger.Info("init logger success")

	return nil
}

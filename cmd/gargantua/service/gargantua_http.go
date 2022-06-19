package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter"

// GargantuaRest is main server starter of gargantua rest service
func GargantuaRest() (err error) {
	// start grpc server
	err = starter.InvokeRest(container)
	if err != nil {
		logger.Errorf("invoke rest server failed, err: %v", err)
		return err
	}
	return nil
}

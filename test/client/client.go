package client

import (
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConnection() (*grpc.ClientConn, error) {
	config, err := common.GetConfig("../test.yaml")
	if err != nil {
		return nil, err
	}
	return grpc.Dial(config.Listen.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

package server

import (
	"github.com/Galaxy-Railway/GGargantua/cmd/gargantua/service"
	"testing"
)

func TestGargantua(t *testing.T) {
	service.GargantuaGrpc("../test.yaml")
}

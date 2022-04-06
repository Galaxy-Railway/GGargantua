package module

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/module"

type SingleRequest struct {
	RequestSchema  module.SchemaType
	RequestContent []byte
}

type SingleResponse struct {
	ResponseSchema  module.SchemaType
	ResponseContent []byte
}

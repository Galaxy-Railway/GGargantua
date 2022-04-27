package module

import "github.com/Galaxy-Railway/GGargantua/api/protobuf"

type Request struct {
	RequestSchema  SchemaType
	RequestContent []byte
}

func TransRequestByPb(request *protobuf.SingleRequest) *Request {
	if request == nil {
		return nil
	}
	return &Request{
		RequestSchema:  SchemaType(int(request.RequestSchema)),
		RequestContent: request.RequestContent,
	}
}

type Response struct {
	ResponseSchema  SchemaType
	ResponseContent []byte
}

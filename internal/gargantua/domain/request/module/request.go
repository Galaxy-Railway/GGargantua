package module

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/http"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/https"
	"time"
)

type Request struct {
	RequestSchema SchemaType
	Times         int
	Concurrency   int

	HttpRequest  *http.RequestContent
	HttpsRequest *https.RequestContent
}

func TransRequestFromPb(request *protobuf.Request) *Request {
	if request == nil {
		return nil
	}
	return &Request{
		RequestSchema: SchemaType(int(request.RequestSchema)),
		Times:         int(request.Times),
		Concurrency:   int(request.Concurrency),

		HttpRequest:  http.TransRequestContentFromPb(request.HttpRequest),
		HttpsRequest: https.TransRequestContentFromPb(request.HttpsRequest),
	}
}

type AllResponse struct {
	HttpResponse  *http.ResponseContent
	HttpsResponse *https.ResponseContent
}

type Response struct {
	ResponseSchema  SchemaType
	TotalTime       time.Duration
	ResponseContent []*AllResponse
}

func TransResponseFromPb(response *protobuf.Response) *Response {
	if response == nil {
		return nil
	}
	return &Response{
		ResponseSchema:  SchemaType(int(response.ResponseSchema)),
		TotalTime:       response.TotalTime.AsDuration(),
		ResponseContent: TransAllResponsesFromPb(response.AllResponse),
	}
}

func TransAllResponseFromPb(response *protobuf.SingleResponse) *AllResponse {
	return &AllResponse{
		HttpResponse:  http.TransResponseContentFromPb(response.HttpResponse),
		HttpsResponse: https.TransResponseContentFromPb(response.HttpsResponse),
	}
}

func TransAllResponsesFromPb(response []*protobuf.SingleResponse) []*AllResponse {
	result := make([]*AllResponse, len(response))
	for i := range response {
		result[i] = TransAllResponseFromPb(response[i])
	}
	return result
}

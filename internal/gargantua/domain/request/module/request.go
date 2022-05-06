package module

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

type Request struct {
	RequestSchema SchemaType
	Times         int
	Concurrency   int

	HttpRequest  *HttpRequestContent
	HttpsRequest *HttpsRequestContent
}

func TransRequestFromPb(request *protobuf.Request) *Request {
	if request == nil {
		return nil
	}
	return &Request{
		RequestSchema: SchemaType(int(request.RequestSchema)),
		Times:         int(request.Times),
		Concurrency:   int(request.Concurrency),

		HttpRequest:  TransHttpRequestContentFromPb(request.HttpRequest),
		HttpsRequest: TransHttpsRequestContentFromPb(request.HttpsRequest),
	}
}

type AllResponse struct {
	HttpResponse  *HttpResponseContent
	HttpsResponse *HttpsResponseContent
}

type Response struct {
	ResponseSchema  SchemaType
	TotalTime       time.Duration
	ResponseContent []*AllResponse
}

func TransResponseToPb(response *Response) *protobuf.Response {
	if response == nil {
		return nil
	}
	return &protobuf.Response{
		ResponseSchema: protobuf.SchemaType(response.ResponseSchema),
		TotalTime:      durationpb.New(response.TotalTime),
		AllResponse:    TransAllResponsesToPb(response.ResponseContent),
	}
}

func TransAllResponseToPb(response *AllResponse) *protobuf.SingleResponse {
	return &protobuf.SingleResponse{
		HttpResponse:  TransHttpResponseContentToPb(response.HttpResponse),
		HttpsResponse: TransHttpsResponseContentToPb(response.HttpsResponse),
	}
}

func TransAllResponsesToPb(response []*AllResponse) []*protobuf.SingleResponse {
	result := make([]*protobuf.SingleResponse, len(response))
	for i := range response {
		result[i] = TransAllResponseToPb(response[i])
	}
	return result
}

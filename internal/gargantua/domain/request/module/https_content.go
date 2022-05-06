package module

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type HttpsRequestContent struct {
	Method             string              `json:"Method,omitempty"`
	Url                string              `json:"Url,omitempty"`
	Headers            map[string][]string `json:"Headers,omitempty"`
	Params             map[string]string   `json:"Params,omitempty"`
	Body               []byte              `json:"Body,omitempty"`
	SkipInsecureVerify bool                `json:"SkipInsecureVerify,omitempty"`
	CaCert             []byte              `json:"CaCert,omitempty"`
	Cert               []byte              `json:"Cert,omitempty"`
	Key                []byte              `json:"Key,omitempty"`
}

type HttpsResponseContent struct {
	Body    []byte
	Cookies []*http.Cookie
	Headers map[string][]string
	Status  int

	StartTime    time.Time
	FistByteTime time.Duration
	CompleteTime time.Duration

	Error error
}

func TransHttpsRequestContentFromPb(request *protobuf.HttpsRequest) *HttpsRequestContent {
	return &HttpsRequestContent{
		Method:  request.Method,
		Url:     request.Url,
		Headers: TransMapFromListOfString(request.Headers),
		Params:  request.Params,
		Body:    request.Body,
	}
}

func TransHttpsResponseContentToPb(response *HttpsResponseContent) *protobuf.HttpsResponse {
	return &protobuf.HttpsResponse{
		Body:         response.Body,
		Cookies:      TransCookieToPb(response.Cookies),
		Headers:      TransMapToListOfString(response.Headers),
		Status:       int32(response.Status),
		StartTime:    timestamppb.New(response.StartTime),
		FistByteTime: durationpb.New(response.FistByteTime),
		CompleteTime: durationpb.New(response.CompleteTime),
		Error:        response.Error.Error(),
	}
}

package module

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type HttpRequestContent struct {
	Method  string
	Url     string
	Headers map[string][]string
	// fixme: params: map[string][]string ???
	Params map[string]string
	Body   []byte
}

type HttpResponseContent struct {
	Body    []byte
	Cookies []*http.Cookie
	Headers map[string][]string
	Status  int

	StartTime    time.Time
	FistByteTime time.Duration
	CompleteTime time.Duration

	Error error
}

func TransHttpRequestContentFromPb(request *protobuf.HttpRequest) *HttpRequestContent {
	if request == nil {
		return nil
	}
	return &HttpRequestContent{
		Method:  request.Method,
		Url:     request.Url,
		Headers: TransMapFromListOfString(request.Headers),
		Params:  request.Params,
		Body:    request.Body,
	}
}

func TransHttpResponseContentToPb(response *HttpResponseContent) *protobuf.HttpResponse {
	return &protobuf.HttpResponse{
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

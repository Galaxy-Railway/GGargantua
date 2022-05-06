package http

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"net/http"
)

func TransRequestContentFromPb(request *protobuf.HttpRequest) *RequestContent {
	return &RequestContent{
		Method:  request.Method,
		Url:     request.Url,
		Headers: TransMapFromListOfString(request.Headers),
		Params:  request.Params,
		Body:    request.Body,
	}
}

func TransResponseContentFromPb(response *protobuf.HttpResponse) *ResponseContent {
	return &ResponseContent{
		Body:         response.Body,
		Cookies:      TransCookieFromPb(response.Cookies),
		Headers:      TransMapFromListOfString(response.Headers),
		Status:       int(response.Status),
		StartTime:    response.StartTime.AsTime(),
		FistByteTime: response.FistByteTime.AsDuration(),
		CompleteTime: response.CompleteTime.AsDuration(),
	}
}

func TransMapFromListOfString(ofString map[string]*protobuf.ListOfString) map[string][]string {
	result := make(map[string][]string, len(ofString))
	for k, v := range ofString {
		result[k] = v.Strings
	}
	return result
}

func TransCookieFromPb(cookie []*protobuf.Cookie) []*http.Cookie {
	result := make([]*http.Cookie, len(cookie))
	for i, c := range cookie {
		result[i] = &http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires.AsTime(),
			RawExpires: c.RawExpires,
			MaxAge:     int(c.MaxAge),
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   http.SameSite(int(c.SameSite)),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		}
	}
	return result
}

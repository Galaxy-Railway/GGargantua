package https

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

type Sender struct {
}

var WrongContentTypeError = errors.New("this is not a https request content")

func (h *Sender) SendOnce(input interface{}) ([]byte, error) {
	content, ok := input.(*RequestContent)
	if !ok {
		return nil, WrongContentTypeError
	}
	resp, err := sendHttpsRequest(content)
	if err != nil {
		return nil, err
	}
	output, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func sendHttpsRequest(httpsRequest *RequestContent) (*ResponseContent, error) {
	var (
		client        *http.Client
		request       *http.Request
		resp          *http.Response
		body          []byte
		err           error
		startTime     time.Time
		firstByteTime time.Duration
		completeTime  time.Duration
	)

	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	request, err = http.NewRequest(httpsRequest.Method, httpsRequest.Url, bytes.NewBuffer(httpsRequest.Body))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			firstByteTime = time.Since(startTime)
		},
	}
	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

	for k, vs := range httpsRequest.Headers {
		for _, v := range vs {
			request.Header.Set(k, v)
		}
	}

	startTime = time.Now()
	resp, err = client.Do(request)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	completeTime = time.Since(startTime)

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	httpsResponse := &ResponseContent{
		Body:         body,
		Cookies:      resp.Cookies(),
		Headers:      resp.Header,
		Status:       resp.StatusCode,
		StartTime:    time.Time{},
		FistByteTime: firstByteTime,
		CompleteTime: completeTime,
	}
	return httpsResponse, nil
}

package http

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

type HttpSender struct {
}

func (h *HttpSender) SendOnce(bytes []byte) ([]byte, error) {
	content := &HttpRequestContent{}
	err := json.Unmarshal(bytes, content)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func sendHttpRequest(httpRequest *HttpRequestContent) (*HttpResponseContent, error) {
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

	request, err = http.NewRequest(httpRequest.Method, httpRequest.Url, bytes.NewBuffer(httpRequest.Body))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			firstByteTime = time.Since(startTime)
		},
	}
	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

	for k, v := range httpRequest.Headers {
		request.Header.Set(k, v)
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

	httpResponse := &HttpResponseContent{
		Body:         body,
		Cookies:      resp.Cookies(),
		Headers:      resp.Header,
		Status:       resp.StatusCode,
		StartTime:    time.Time{},
		FistByteTime: firstByteTime,
		CompleteTime: completeTime,
	}
	return httpResponse, nil
}
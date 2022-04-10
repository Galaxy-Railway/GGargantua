package http

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
	HttpClient *http.Client
}

func (s *Sender) Init() {
	// todo: set max conn, timeout, etc.
	s.HttpClient = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
}

// SendOnce will send a message for once
func (s *Sender) SendOnce(input []byte) ([]byte, error) {
	content := &RequestContent{}
	err := json.Unmarshal(input, content)
	if err != nil {
		return nil, err
	}
	resp, err := s.sendHttpRequest(content)
	if err != nil {
		return nil, err
	}
	output, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// sendHttpRequest is the core part to send a request with http schema
func (s *Sender) sendHttpRequest(httpRequest *RequestContent) (*ResponseContent, error) {
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

	client = s.HttpClient

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

	httpResponse := &ResponseContent{
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
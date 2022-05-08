package http

import (
	"bytes"
	"crypto/tls"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
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

var WrongContentTypeError = errors.New("this is not a http request content")

// SendOnce will send a message for once
func (h *Sender) SendOnce(input interface{}) (resp *module.AllResponse, err error) {
	h.Init()
	resp = &module.AllResponse{}
	resp.HttpResponse = &module.HttpResponseContent{}
	defer func() {
		if err != nil {
			resp.HttpResponse.Error = err
		}
	}()
	content, ok := input.(*module.HttpRequestContent)
	if !ok {
		return nil, WrongContentTypeError
	}
	response, err := h.sendHttpRequest(content)
	if err != nil {
		return nil, err
	}
	resp.HttpResponse = response
	return
}

// sendHttpRequest is the core part to send a request with http schema
func (s *Sender) sendHttpRequest(httpRequest *module.HttpRequestContent) (*module.HttpResponseContent, error) {
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

	for k, vs := range httpRequest.Headers {
		for _, v := range vs {
			request.Header.Add(k, v)
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

	httpResponse := &module.HttpResponseContent{
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

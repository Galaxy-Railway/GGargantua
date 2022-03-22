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

type HttpsSender struct {
}

func (h *HttpsSender) SendOnce(intput []byte) ([]byte, error) {
	content := &HttpsRequestContent{}
	err := json.Unmarshal(intput, content)
	if err != nil {
		return nil, err
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

func sendHttpsRequest(httpsRequest *HttpsRequestContent) (*HttpsResponseContent, error) {
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

	for k, v := range httpsRequest.Headers {
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

	httpsResponse := &HttpsResponseContent{
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

package service

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/sender"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type RequestServiceImpl struct {
}

func NewRequestService() RequestService {
	return &RequestServiceImpl{}
}

func (s *RequestServiceImpl) SendRequest(request *module.Request) (*module.Response, error) {
	var (
		iSender sender.ISender
		content interface{}
	)

	switch request.RequestSchema {
	case module.HTTP:
		iSender = sender.HttpSender()
		content = request.HttpRequest
	case module.HTTPS:
		iSender = sender.HttpsSender()
		content = request.HttpsRequest
	}

	// todo: make goroutine pool, do parallel request
	pool, err := ants.NewPool(request.Concurrency)
	if err != nil {
		return nil, errors.Wrap(err, "get ants pool failed")
	}

	var (
		wg        sync.WaitGroup
		responses []*module.AllResponse
		lock      sync.Mutex
	)
	startTime := time.Now()
	for i := 0; i < request.Times; i++ {
		wg.Add(1)
		//todo: deal with error, log it
		_ = pool.Submit(func() {
			resp, err := iSender.SendOnce(content)
			if err != nil {
				//todo: log
			}
			lock.Lock()
			defer lock.Unlock()
			responses = append(responses, resp)
			wg.Done()
		})
	}
	wg.Wait()
	totalTime := time.Since(startTime)
	return &module.Response{
		ResponseSchema:  request.RequestSchema,
		TotalTime:       totalTime,
		ResponseContent: responses,
	}, nil
}

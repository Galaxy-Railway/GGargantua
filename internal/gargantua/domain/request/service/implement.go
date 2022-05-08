package service

import (
	"context"
	requestDomain "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/module"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/sender"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
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

var SendRequestFailedError = errors.New("failed to send request")

func (s *RequestServiceImpl) SendRequest(ctx context.Context, request *module.Request) (response *module.Response, err error) {
	var (
		iSender   sender.ISender
		content   interface{}
		wg        sync.WaitGroup
		errorList []error
		lock      sync.Mutex
	)

	response = &module.Response{
		ResponseSchema:  request.RequestSchema,
		TotalRequests:   request.Times,
		SuccessRequests: 0,
		ResponseContent: make([]*module.AllResponse, 0),
	}

	switch request.RequestSchema {
	case module.HTTP:
		iSender = sender.HttpSender()
		content = request.HttpRequest
	case module.HTTPS:
		iSender = sender.HttpsSender()
		content = request.HttpsRequest
	}

	pool, err := ants.NewPool(request.Concurrency)
	if err != nil {
		err = errors.Wrap(err, "get ants pool failed")
		return
	}

	errorList = make([]error, 0)
	startTime := time.Now()
	for i := 0; i < request.Times; i++ {
		select {
		case <-ctx.Done():
			err = common.CanceledError
			break
		default:
			wg.Add(1)
			//todo: deal with error, log it
			err = pool.Submit(func() {
				resp, err := iSender.SendOnce(ctx, content)
				if err != nil {
					requestDomain.Logger().Infof("failed request, err: %+v", err)
					errorList = append(errorList, err)
				} else {
					lock.Lock()
					defer lock.Unlock()
					response.ResponseContent = append(response.ResponseContent, resp)
				}
				wg.Done()
			})
			if err != nil {
				requestDomain.Logger().Errorf("failed to submit a job into pool, err: %+v", err)
			}
		}
	}
	wg.Wait()

	response.SuccessRequests = len(response.ResponseContent)
	response.TotalTime = time.Since(startTime)
	if len(errorList) != 0 {
		err = errors.Wrapf(SendRequestFailedError, "errors: %+v", errorList)
		return
	}

	return
}

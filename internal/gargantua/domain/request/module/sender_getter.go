package module

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/http"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/https"
	"sync"
)

var (
	once = &sync.Once{}

	httpSender  *http.Sender
	httpsSender *https.Sender
)

func HttpSender() ISender {
	once.Do(func() {
		httpSender = &http.Sender{}
	})
	return httpSender
}

func HttpsSender() ISender {
	once.Do(func() {
		httpsSender = &https.Sender{}
	})
	return httpsSender
}

package module

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/http"
	"sync"
)

var (
	once = &sync.Once{}

	httpSender *http.Sender
)

func NewHttpSender() Sender {
	once.Do(func() {
		httpSender = &http.Sender{}
	})
	return httpSender
}

package module

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/sender/http"
	"sync"
)

var (
	once *sync.Once

	httpSender *http.HttpSender
)

func NewHttpSender() Sender {
	once.Do(func() {
		httpSender = &http.HttpSender{}
	})
	return httpSender
}

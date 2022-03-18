package http

import (
	"net/http"
	"time"
)

type HttpResponseContent struct {
	Body         []byte
	Cookies      []*http.Cookie
	Headers      map[string][]string
	Status       int
	StartTime    time.Time
	FistByteTime time.Duration
	CompleteTime time.Duration
}

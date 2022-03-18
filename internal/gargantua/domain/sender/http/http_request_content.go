package http

type HttpRequestContent struct {
	Method  string
	Url     string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

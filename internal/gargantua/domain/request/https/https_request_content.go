package https

type RequestContent struct {
	Method             string              `json:"Method,omitempty"`
	Url                string              `json:"Url,omitempty"`
	Headers            map[string][]string `json:"Headers,omitempty"`
	Params             map[string]string   `json:"Params,omitempty"`
	Body               []byte              `json:"Body,omitempty"`
	SkipInsecureVerify bool                `json:"SkipInsecureVerify,omitempty"`
	CaCert             []byte              `json:"CaCert,omitempty"`
	Cert               []byte              `json:"Cert,omitempty"`
	Key                []byte              `json:"Key,omitempty"`
}

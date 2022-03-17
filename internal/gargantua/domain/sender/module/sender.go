package module

type Sender struct {
	Req *Request
	Res *Response
}

func (s *Sender)Send() error {
	// todo: implement me
	return nil
}
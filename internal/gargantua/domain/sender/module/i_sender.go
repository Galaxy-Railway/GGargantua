package module

type Sender interface {
	SendOnce(request []byte) ([]byte, error)
}

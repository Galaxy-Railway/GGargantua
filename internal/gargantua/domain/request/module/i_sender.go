package module

type ISender interface {
	SendOnce(request interface{}) ([]byte, error)
}

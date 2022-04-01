package module

type ISender interface {
	SendOnce(request []byte) ([]byte, error)
}

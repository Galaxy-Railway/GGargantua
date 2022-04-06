package module

// SchemaType is type of request schema.
// Such as HTTP, HTTPS.
// And GRPC, TCP, UDP, etc. will support soon too.
type SchemaType int32

const (
	HTTP SchemaType = iota
	HTTPS
)

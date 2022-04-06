package module

type Request struct {
	RequestSchema  SchemaType
	RequestContent []byte
}

type Response struct {
	ResponseSchema  SchemaType
	ResponseContent []byte
}

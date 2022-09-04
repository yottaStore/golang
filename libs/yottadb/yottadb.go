package yottadb

type ReadRequest struct {
	Path string
	Mode string
}

type ReadResponse struct {
	Data []byte
}

type WriteRequest struct {
	Path             string
	Data             []byte
	CreateCollection bool
}

type WriteResponse struct {
}

type Interface interface {
	Read(ReadRequest) (ReadResponse, error)
	Write(WriteRequest) (WriteResponse, error)
	Update(WriteRequest) (WriteResponse, error)
	Append(WriteRequest) (WriteResponse, error)
	Delete(WriteRequest) error
}

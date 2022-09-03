package iodrivers

type IoReadRequest struct {
	Path string
}

type IoReadResponse struct {
	Data []byte
	Aba  string
}

type IoWriteResponse struct {
	Aba string
}

type IoWriteRequest struct {
	Path   string
	Data   []byte
	Method string
}

type IoDriverInterface interface {
	Read(IoReadRequest) (IoReadResponse, error)
	Write(IoWriteRequest) (IoWriteResponse, error)
	Append(IoWriteRequest) (IoWriteResponse, error)
	Delete(IoWriteRequest) error
	//CompareAndSwap(IoWriteRequest) error
	//CompareAndAppend(IoWriteRequest) error
	//Verify(string) error
}

type IoDriver struct {
	IoDriverInterface
	NameSpace string
}

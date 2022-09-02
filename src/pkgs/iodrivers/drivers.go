package iodrivers

type IoReadRequest struct {
	Path   string
	Method string
}

type IoReadResponse struct {
	Data []byte
	Aba  string
}

type IoWriteRequest struct {
	Path   string
	Data   []byte
	Method string
}

type FSDev string

type BlkDev int64

type IoDriverInterface interface {
	Init() error
	Read(IoReadRequest) (IoReadResponse, error)
	ReadAll(IoReadRequest) (IoReadResponse, error)
	Write(IoWriteRequest) error
	Append(IoWriteRequest) error
	Delete(IoWriteRequest) error
	CompareAndSwap(IoWriteRequest) error
	CompareAndAppend(IoWriteRequest) error
	Verify(string, []byte) error
}

type Config struct {
	NameSpace string
	Driver    string
}

type IoDriver struct {
	IoDriverInterface
	NameSpace string
}

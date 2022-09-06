package iodriver

type RequestMethod string

const (
	Read   RequestMethod = "read"
	Write                = "write"
	Delete               = "delete"
	Append               = "append"
)

type DriverType string

const (
	Dummy  = "dummy"
	Direct = "direct"
)

type Request struct {
	Path    string        `json:"Path"`
	Method  RequestMethod `json:"Method"`
	Data    []byte        `json:"Data"`
	Options interface{}   `json:"Options"`
}

type Response struct {
	Path     string
	Method   RequestMethod
	Data     []byte
	AbaToken string
}

type Interface interface {
	Read(Request) (Response, error)
	Write(Request) (Response, error)
	Append(Request) (Response, error)
	Delete(Request) error
	//CompareAndSwap(IoWriteRequest) error
	//CompareAndAppend(IoWriteRequest) error
	//Verify(string) error
}

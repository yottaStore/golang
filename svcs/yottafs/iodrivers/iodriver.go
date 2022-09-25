package iodrivers

type Request struct {
	Path    string `cbor:"Path"`
	Data    []byte `cbor:"Data"`
	Options []byte `cbor:"Options"`
}

type Response struct {
	Path     string
	Data     []byte
	AbaToken []byte
}

type DataBlock struct {
	Data []byte
}

type Interface interface {
	Create(Request) (Response, error)
	Read(Request) (Response, error)
	Update(Request) (Response, error)
	Delete(Request) error
	CompareAndSwap(Request) error
	CompareAndAppend(Request) error
	Verify(Request) error
}

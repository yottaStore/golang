package iodrivers

type Interface interface {
	// Read methods
	Read(Request) (Response, error)
	Compare(Request) (Response, error)

	// Write methods
	Write(Request) (Response, error)
	Delete(Request) (Response, error)
	CompareAndSwap(Request) (Response, error)

	// Utility methods
	Verify(Request) (Response, error)
	Check(Request) (Response, error)
}

type Method string

const (
	Read    Method = "read"
	Compare        = "compare"

	Write          = "write"
	Append         = "append"
	CompareAndSwap = "cas"
	Delete         = "delete"

	Verify = "verify"
	Check  = "check"
)

// TODO: use this
type DataBlock struct {
	Data []byte
}

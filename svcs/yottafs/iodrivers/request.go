package iodrivers

type Request struct {
	Path    string
	Method  Method
	Options uint32
	Data    []byte
}

type Method string

const (
	Read   Method = "read"
	Write         = "write"
	Delete        = "delete"
)

type DataBlock struct {
	Data []byte
}

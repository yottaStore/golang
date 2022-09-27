package iodrivers

type Response struct {
	Path       string
	Data       []byte
	Generation []byte
}

type Request struct {
	Method  Method
	Path    string
	Data    []byte
	Options Option
	Offset  Offset
}

type Offset struct {
	Start uint32
	End   uint32
}

type Option uint32

const (
	CREATE Option = 0x1
	UPDATE        = 0x2
	UPSERT        = 0x3
	APPEND        = 0x4
	OFFSET        = 0x8
)

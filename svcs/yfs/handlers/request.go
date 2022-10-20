package handlers

type Method string

const (
	READ    Method = "READ"
	CREATE         = "CREATE"
	DELETE         = "DELETE"
	APPEND         = "APPEND"
	COMPACT        = "COMPACT"
	MERGE          = "MERGE"
)

type RequestFlag uint32

const (
	FLAG_NONE RequestFlag = 0
)

type Request struct {
	Method Method
	Record string
	Flags  RequestFlag
	Data   []byte
}

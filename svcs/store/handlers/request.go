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

type Driver string

const (
	KEY_VALUE Driver = "KEY_VALUE"
	DOCUMENT         = "DOCUMENT"
	PUBSUB           = "PUBSUB"
)

type Request struct {
	Record string
	Method Method
	Driver Driver
	Flags  RequestFlag
	Data   []byte
}

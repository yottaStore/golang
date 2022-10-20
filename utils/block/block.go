package block

type Block struct {
	Version   uint8
	Type      Type
	Flags     Flags
	Length    uint16
	Reserved1 uint8
	Reserved2 uint8
	// Total 64 bits
	Body []byte
	Hash uint64
}

// Versions
const v0 uint8 = 0

// Types
type Type uint8

const (
	BodyType   Type = 1
	TailType   Type = 2
	AppendType Type = 3
	SkipType   Type = 17
)

// Flags
type Flags uint16

const (
	F_COMPRESSED Flags = 1 << 10
)

package block

type Block struct {
	Version   uint8
	Type      BlockType
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
type BlockType uint8

const (
	BodyType   BlockType = 1
	TailType   BlockType = 2
	AppendType BlockType = 3
)

// Flags
type Flags uint16

const (
	F_COMPRESSED Flags = 1 << 10
)

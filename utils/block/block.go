package block

const (
	Size = 4096
)

type Block [Size]byte

type BlockPointer struct {
	Offset uint64
	Length uint16
}

type Tails []BlockPointer

type Record struct {
	Body    []Block
	Tails   Tails
	Appends [][]Block
}

func (b *Record) Create(body []byte) error { return nil }

func (b *Record) Read() ([]byte, error) {

	// Read Tails
	// Read Body
	// Read Appends

	return nil, nil
}

func (b *Record) Append(tail []byte, append [][]byte) error { return nil }

func (b *Record) Delete() error { return nil }

func (b *Record) Compact() error {

	// Write attempt to log
	// Read record
	// Compact record
	// Write temporary body

	// Replace body
	// Collapse tails
	// Write success to log

	return nil
}

func (b *Record) Merge() error { return nil }

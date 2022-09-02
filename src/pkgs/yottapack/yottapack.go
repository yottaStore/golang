package yottapack

// Read
// Write
// Append
// Update

type PackerInterface[T any] interface {
	Pack(T) ([]byte, error)
	Unpack([]byte) (T, error)
}

type Packer[T any] struct {
	PackerInterface[T]
}

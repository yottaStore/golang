package hasher

type Interface interface {
	Init(pointer string, seed uint64) error
	Read(pointer string, level, length int) ([]byte, error)
}

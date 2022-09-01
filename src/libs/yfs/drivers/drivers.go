package drivers

type IoDriver interface {
	Init() error
	Read(string) ([]byte, error)
	ReadAll(string) ([]byte, error)
	Write(string, []byte) error
	Append(string, []byte) error
	Delete(string) error
	CompareAndSwap(string, []byte) error
	CompareAndAppend(string, []byte) error
	Verify(string, []byte) error
}

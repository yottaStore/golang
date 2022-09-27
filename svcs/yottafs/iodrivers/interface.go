package iodrivers

type Interface interface {
	// Record methods
	Read(Request) (Response, error)
	Write(Request) (Response, error)
	Update(Request) (Response, error)
	CompareAndSwap(Request) (Response, error)

	// Queue methods
	Seek(Request) (Response, error)
	Append(Request) (Response, error)
	CompareAndAppend(Request) (Response, error)

	// Common methods
	Delete(Request) (Response, error)
	Verify(Request) (Response, error)
	Fscheck(Request) (Response, error)
}

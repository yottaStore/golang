package io_driver

type IODriver interface {
	Read(string) ([]byte, error)
	Create(string, []byte) error
	Delete(string) error
	Append(string, []byte) error
	Compact(string) error
	Merge(string, []byte) error
}

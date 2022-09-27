package direct

import (
	"errors"
	"yottafs/iodrivers"
	"yottafs/iodrivers/direct/methods/read"
	"yottafs/iodrivers/direct/methods/write"
)

type Driver struct {
	iodrivers.Interface
	namespace string
	dataspace string
}

var (
	NotImplemented = errors.New("method not implemented yet")
)

func (d Driver) Read(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.dataspace + r.Path
	return read.Read(path)
}

func (d Driver) Compare(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, NotImplemented
}

func (d Driver) Write(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.dataspace + r.Path
	return write.Write(path, r.Data, true)
}

func (d Driver) CompareAndSwap(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, NotImplemented
}

func (d Driver) Delete(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.dataspace + r.Path
	return write.Delete(path)
}

func (d Driver) Verify(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, NotImplemented
}

func (d Driver) Check(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, NotImplemented
}

package dummy

import (
	"yottafs/iodrivers"
)

type Driver struct {
	iodrivers.Interface
}

func (d Driver) Read(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) Compare(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) Write(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) CompareAndSwap(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) Delete(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) Verify(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

func (d Driver) Check(r iodrivers.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response
	return resp, nil
}

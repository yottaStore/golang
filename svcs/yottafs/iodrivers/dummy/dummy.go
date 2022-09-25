package dummy

import "yottafs/iodrivers"

type Driver struct {
	iodrivers.Interface
}

func (d Driver) Create(r iodrivers.Request) (iodrivers.Response, error) {

	resp := iodrivers.Response{
		Path: r.Path}
	return resp, nil
}

func (d Driver) Read(r iodrivers.Request) (iodrivers.Response, error) {

	resp := iodrivers.Response{
		Path: r.Path}
	return resp, nil
}

func (d Driver) Update(r iodrivers.Request) (iodrivers.Response, error) {

	resp := iodrivers.Response{
		Path: r.Path}
	return resp, nil
}

func (d Driver) Delete(r iodrivers.Request) error {

	return nil

}

func New() (Driver, error) {

	var d Driver

	return d, nil
}

package direct

import (
	"yottafs/iodrivers"
	"yottafs/iodrivers/direct/methods"
)

type Driver struct {
	iodrivers.Interface
	namespace string
}

func (d Driver) Write(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.namespace + "/data/" + r.Path
	return methods.Write(path, r.Data, true)
}

func (d Driver) Read(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.namespace + "/data/" + r.Path
	return methods.Read(path)
}

func (d Driver) Update(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.namespace + "/data/" + r.Path
	return methods.Write(path, r.Data, true)
}

func (d Driver) Delete(r iodrivers.Request) (iodrivers.Response, error) {

	path := d.namespace + "/data/" + r.Path
	return methods.Delete(path)
}

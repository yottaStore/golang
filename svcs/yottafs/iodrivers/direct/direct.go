package direct

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers"
	"yottafs/iodrivers/direct/methods"
)

type Driver struct {
	iodrivers.Interface
	namespace string
}

func (d Driver) Create(r iodrivers.Request) (iodrivers.Response, error) {

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

func (d Driver) Delete(r iodrivers.Request) error {

	path := d.namespace + "/data/" + r.Path
	return methods.Delete(path)
}

func New(namespace string) (Driver, error) {

	var d Driver

	err := unix.Access(namespace, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace, 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return d, err
		}
	default:
		return d, err
	}

	err = unix.Access(namespace+"/data/", unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace+"/data/", 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return d, err
		}
	default:
		log.Println("Error instantiating driver: ", err)
		return d, err
	}

	driver := Driver{
		namespace: namespace,
	}

	return driver, nil
}

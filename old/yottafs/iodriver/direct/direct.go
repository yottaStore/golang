package direct

import (
	"golang.org/x/sys/unix"
	"yottafs/iodriver"
	"yottafs/iodriver/direct/methods"
)

type Driver struct {
	iodriver.Interface
	namespace string
}

func (d Driver) Read(req iodriver.Request) (iodriver.Response, error) {

	path := d.namespace + "/data" + req.Path
	resp, err := methods.Read(path)

	// TODO: removing trailing zeros
	return resp, err
}

func (d Driver) Write(req iodriver.Request) (iodriver.Response, error) {
	path := d.namespace + "/data" + req.Path
	return methods.Write(path, []byte(req.Data), true)

}

func (d Driver) Append(req iodriver.Request) (iodriver.Response, error) {

	path := d.namespace + "/data" + req.Path
	return methods.Append(path, []byte(req.Data))
}

func (d Driver) Delete(req iodriver.Request) error {

	path := d.namespace + "/data" + req.Path
	return methods.Delete(path)
}

func New(namespace string) (iodriver.Interface, error) {

	err := unix.Access(namespace, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace, 0766); err != nil {
			return nil, err
		}
	default:
		return nil, err

	}

	err = unix.Access(namespace+"/data", unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace+"/data", 0766); err != nil {
			return nil, err
		}
	default:
		return nil, err

	}

	// TODO: Create file

	driver := Driver{
		namespace: namespace,
	}

	return driver, nil

}

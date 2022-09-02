package direct

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/read"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/write"
)

type DirectDriver struct {
	Namespace string
}

func (d DirectDriver) Init() error {

	var stats unix.Stat_t
	err := unix.Stat(d.Namespace+"/data", &stats)
	if err == unix.ENOENT {
		// TODO: check permission
		err = unix.Mkdir(d.Namespace+"/data", 0777)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func (d DirectDriver) Read(record string) ([]byte, error) {

	path := d.Namespace + "data/" + record
	buff, err := read.ReadAll(path)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (d DirectDriver) ReadAll(record string) ([]byte, error) {

	path := d.Namespace + "data/" + record
	buff, err := read.ReadAll(path)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (d DirectDriver) Write(record string, data []byte) error {

	path := d.Namespace + "data/" + record
	err := write.Write(path, data)
	if err != nil {
		return err
	}

	return nil
}

func (d DirectDriver) Append(record string, data []byte) error {

	path := d.Namespace + "data/" + record
	err := write.Append(path, data)
	if err != nil {
		return err
	}

	return nil
}

func (d DirectDriver) CompareAndSwap(path string, data []byte) error {

	return errors.New("method not implemented")
}

func (d DirectDriver) CompareAndAppend(path string, data []byte) error {

	return errors.New("method not implemented")
}

func (d DirectDriver) Delete(record string) error {

	path := d.Namespace + "data/" + record
	err := write.Delete(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d DirectDriver) Verify(path string, data []byte) error {

	return errors.New("method not implemented")
}

func New(opts iodrivers.Config) iodrivers.IoDriver {

	directDriver := DirectDriver{
		Namespace: opts.NameSpace,
	}

	return directDriver
}

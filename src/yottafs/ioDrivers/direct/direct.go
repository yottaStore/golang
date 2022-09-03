package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
	"yottafs/ioDrivers"
	"yottafs/ioDrivers/direct/src"
)

type Driver struct {
	ioDrivers.IoDriverInterface
	NameSpace     string
	DataNameSpace string
}

func (d Driver) Read(req ioDrivers.IoReadRequest) (ioDrivers.IoReadResponse, error) {

	var resp ioDrivers.IoReadResponse
	path := d.NameSpace + "/data" + req.Path
	buff, err := src.Read(path)
	if err != nil {
		return resp, err
	}
	resp.Data = buff
	resp.Aba = "123"

	// TODO: removing trailing zeros
	return resp, nil
}

func (d Driver) Write(req ioDrivers.IoWriteRequest) (ioDrivers.IoWriteResponse, error) {

	var resp ioDrivers.IoWriteResponse
	path := d.NameSpace + "/data" + req.Path
	err := src.Write(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Append(req ioDrivers.IoWriteRequest) (ioDrivers.IoWriteResponse, error) {

	var resp ioDrivers.IoWriteResponse
	path := d.NameSpace + "/data" + req.Path
	err := src.AppendTo(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Delete(req ioDrivers.IoWriteRequest) error {

	fmt.Println("Deleting: ", req.Path, "\n")

	path := d.NameSpace + "/data" + req.Path
	err := src.Delete(path)
	if err != nil {
		return err
	}

	return nil
}

func New(nameSpace string) (ioDrivers.IoDriverInterface, error) {

	if err := unix.Access(nameSpace, unix.O_RDWR); err != nil {
		return nil, err
	}

	if err := unix.Access(nameSpace+"/data", unix.O_RDWR); err != nil {
		return nil, err
	}

	ioDriver := Driver{
		NameSpace:     nameSpace,
		DataNameSpace: nameSpace,
	}

	return ioDriver, nil
}

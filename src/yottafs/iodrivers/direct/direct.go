package direct

import (
	"golang.org/x/sys/unix"
	"yottafs/iodrivers"
	"yottafs/iodrivers/direct/src"
)

type Driver struct {
	iodrivers.IoDriverInterface
	NameSpace     string
	DataNameSpace string
}

func (d Driver) Read(req iodrivers.IoReadRequest) (iodrivers.IoReadResponse, error) {

	var resp iodrivers.IoReadResponse
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

func (d Driver) Write(req iodrivers.IoWriteRequest) (iodrivers.IoWriteResponse, error) {

	var resp iodrivers.IoWriteResponse
	path := d.NameSpace + "/data" + req.Path
	err := src.Write(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Append(req iodrivers.IoWriteRequest) (iodrivers.IoWriteResponse, error) {

	var resp iodrivers.IoWriteResponse
	path := d.NameSpace + "/data" + req.Path
	err := src.AppendTo(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Delete(req iodrivers.IoWriteRequest) error {

	path := d.NameSpace + "/data" + req.Path
	err := src.Delete(path)
	if err != nil {
		return err
	}

	return nil
}

func New(nameSpace string) (iodrivers.IoDriverInterface, error) {

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

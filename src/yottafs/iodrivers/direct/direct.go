package direct

import (
	"yottafs/iodrivers"
)

type Driver struct {
	iodrivers.IoDriverInterface
	NameSpace     string
	DataNameSpace string
}

func (d Driver) Read(req iodrivers.IoReadRequest) (iodrivers.IoReadResponse, error) {

	var resp iodrivers.IoReadResponse
	path := d.NameSpace + "/data" + req.Path
	buff, err := read(path)
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
	err := write(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Append(req iodrivers.IoWriteRequest) (iodrivers.IoWriteResponse, error) {

	var resp iodrivers.IoWriteResponse
	path := d.NameSpace + "/data" + req.Path
	err := appendTo(path, req.Data)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d Driver) Delete(req iodrivers.IoWriteRequest) error {

	path := d.NameSpace + "/data" + req.Path
	err := delete(path)
	if err != nil {
		return err
	}

	return nil
}

func New(nameSpace string) (iodrivers.IoDriverInterface, error) {

	ioDriver := Driver{
		NameSpace:     nameSpace,
		DataNameSpace: nameSpace,
	}

	return ioDriver, nil
}

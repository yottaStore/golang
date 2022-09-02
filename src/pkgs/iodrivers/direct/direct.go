package direct

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/read"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/utils"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/write"
)

type Driver struct {
	iodrivers.IoDriver
	DataNameSpace string
}

func (d Driver) Init() error {

	log.Println("Initializing driver")
	var stats unix.Stat_t
	if err := unix.Stat(d.NameSpace, &stats); err != nil {
		return err
	}
	utils.CreateDirIfNotExists(d.DataNameSpace)

	return nil
}

func (d Driver) Read(req iodrivers.IoReadRequest) (iodrivers.IoReadResponse, error) {

	var resp iodrivers.IoReadResponse
	path := d.DataNameSpace + req.Path
	buff, err := read.ReadAll(path)
	if err != nil {
		return resp, err
	}
	resp.Data = buff
	return resp, nil
}

func (d Driver) ReadAll(req iodrivers.IoReadRequest) (iodrivers.IoReadResponse, error) {

	var resp iodrivers.IoReadResponse
	path := d.DataNameSpace + req.Path
	buff, err := read.ReadAll(path)
	if err != nil {
		return resp, err
	}
	resp.Data = buff
	return resp, nil
}

func (d Driver) Write(req iodrivers.IoWriteRequest) error {

	path := d.DataNameSpace + req.Path
	if err := write.Write(path, req.Data); err != nil {
		return err
	}
	return nil
}

func (d Driver) Append(req iodrivers.IoWriteRequest) error {

	path := d.DataNameSpace + req.Path
	if err := write.Append(path, req.Data); err != nil {
		return err
	}
	return nil
}

func (d Driver) CompareAndSwap(req iodrivers.IoWriteRequest) error {

	return errors.New("method not implemented")
}

func (d Driver) CompareAndAppend(req iodrivers.IoWriteRequest) error {

	return errors.New("method not implemented")
}

func (d Driver) Delete(req iodrivers.IoWriteRequest) error {

	path := d.DataNameSpace + req.Path
	err := write.Delete(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d Driver) Verify(path string, data []byte) error {

	return errors.New("method not implemented")
}

func New(conf iodrivers.Config) (iodrivers.IoDriverInterface, error) {

	driver := Driver{
		DataNameSpace: conf.NameSpace + "/data/",
	}
	driver.NameSpace = conf.NameSpace

	err := driver.Init()

	return driver, err
}

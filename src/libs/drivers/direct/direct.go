package direct

import (
	"errors"
	"yottaStore/yottaStore-go/src/libs/drivers"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/read"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/write"
)

type DirectDriver struct {
	Path string
}

func (d DirectDriver) Init() error {

	return nil
}

func (d DirectDriver) Read(path string) ([]byte, error) {

	buff, err := read.ReadAll(path)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (d DirectDriver) ReadAll(path string) ([]byte, error) {

	buff, err := read.ReadAll(path)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (d DirectDriver) Write(path string, data []byte) error {

	err := write.Write(path, data)
	if err != nil {
		return err
	}

	return nil
}

func (d DirectDriver) Append(path string, data []byte) error {

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

func (d DirectDriver) Delete(path string) error {

	err := write.Delete(path)
	if err != nil {
		return err
	}

	return nil
}

func (d DirectDriver) Verify(path string, data []byte) error {

	return errors.New("method not implemented")
}

func New() drivers.IoDriver {

	directDriver := DirectDriver{}

	return directDriver
}

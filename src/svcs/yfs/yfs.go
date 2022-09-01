package yfs

import (
	"errors"
	"yottaStore/yottaStore-go/src/pkgs/drivers"
	"yottaStore/yottaStore-go/src/pkgs/drivers/direct"
)

const (
	DirectIO = "direct"
)

func New(config drivers.Config) (drivers.IoDriver, error) {

	var ioDriver drivers.IoDriver

	switch config.Driver {

	case DirectIO:
		ioDriver = direct.New(config)
	default:
		return ioDriver, errors.New("No driver specified")
	}

	if err := ioDriver.Init(); err != nil {
		return ioDriver, err
	}

	return ioDriver, nil
}

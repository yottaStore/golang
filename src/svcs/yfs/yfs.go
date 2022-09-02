package yfs

import (
	"errors"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct"
)

const (
	DirectIO = "direct"
)

func New(config iodrivers.Config) (iodrivers.IoDriver, error) {

	var ioDriver iodrivers.IoDriver

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

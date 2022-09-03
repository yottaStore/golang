package yottafs

import (
	"errors"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct"
)

const (
	DirectIO = "direct"
)

func New(config iodrivers.Config) (ioDriver iodrivers.IoDriverInterface, err error) {

	switch config.Driver {

	case DirectIO:
		ioDriver, err = direct.New(config)
	default:
		return ioDriver, errors.New("No driver specified")
	}

	return ioDriver, err
}

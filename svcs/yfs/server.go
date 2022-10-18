package yfs

import (
	"errors"
	"github.com/yottaStore/golang/svcs/yfs/iodriver"
	"github.com/yottaStore/golang/svcs/yfs/iodriver/unix_xfs"
	"log"
)

type Config struct {
	Namespace string
	IoDriver  string
	Protocol  string
	Port      string
}

type Request struct {
	Op   string
	Path string
	Data []byte
}

func Start(c Config) error {

	var iod iodriver.Iodriver
	var err error

	switch c.IoDriver {
	case "unix_xfs":
		iod, err = unix_xfs.New(c.Namespace)
		if err != nil {
			log.Println("Error creating iodriver: ", err)
			return err
		}
	default:
		return errors.New("invalid iodriver: " + c.IoDriver)

	}

	switch c.Protocol {
	case "http":
		err = New(c, iod)
		if err != nil {
			log.Println("Error creating handler: ", err)
			return err
		}
	case "quic":
		return errors.New("quic protocol not supported yet")
	default:
		return errors.New("protocol not supported: " + c.Protocol)
	}

	return nil
}

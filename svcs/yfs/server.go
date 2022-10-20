package yfs

import (
	"errors"
	"github.com/yottaStore/golang/svcs/yfs/handlers/http"
	"github.com/yottaStore/golang/svcs/yfs/io_driver"
	"github.com/yottaStore/golang/svcs/yfs/io_driver/unix_xfs"
	"log"
)

type Config struct {
	Namespace string
	IoDriver  string
	Protocol  string
	Port      string
}

func Start(c Config) error {

	var iod io_driver.IODriver
	var err error

	switch c.IoDriver {
	case "unix_xfs":
		iod, err = unix_xfs.New(c.Namespace)
		if err != nil {
			log.Println("Error creating io_driver: ", err)
			return err
		}
	default:
		return errors.New("invalid io_driver: " + c.IoDriver)

	}

	switch c.Protocol {
	case "http":
		hc := http.Config{
			Port: c.Port,
		}
		err = http.New(hc, iod)
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

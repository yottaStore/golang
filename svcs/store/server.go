package store

import (
	"errors"
	"github.com/yottaStore/golang/svcs/store/handlers/http"
	"log"
)

type Config struct {
	Protocol string
	Port     string
}

func Start(c Config) error {

	switch c.Protocol {
	case "http":
		hc := http.Config{
			Port: c.Port,
		}
		err := http.Start(hc)
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

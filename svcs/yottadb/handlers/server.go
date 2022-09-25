package handlers

import (
	"errors"
	"log"
)

type Config struct {
	Nodetree *[]string
	Port     string
	Protocol string
	Hashkey  string
}

func StartServer(c Config) error {

	log.Println("Starting yottadb...")

	if c.Hashkey == "" {
		return errors.New("invalid hash key")
	}
	if len(*c.Nodetree) == 0 {
		return errors.New("invalid node tree")
	}
	if c.Protocol == "" {
		return errors.New("invalid protocol")
	}

	switch c.Protocol {
	case "http":
		err := StartHttpServer(c)
		if err != nil {
			return err
		}
	case "quic":

	default:
		return errors.New("invalid protocol")
	}

	return nil

}

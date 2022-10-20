package http

import (
	"github.com/yottaStore/golang/svcs/yfs/io_driver"
	"log"
	"net/http"
)

type Config struct {
	Port string
}

func New(c Config, d io_driver.IODriver) error {

	h, err := HttpHandlerFactory(d)
	if err != nil {
		log.Fatal("Error creating http handler: ", err)
	}

	http.HandleFunc("/yfs", h)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", notFoundHandler)

	log.Println("Listening on port: ", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Println("Error starting server: ", err)
		return err
	}

	return nil
}

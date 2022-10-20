package http

import (
	"log"
	"net/http"
)

type Config struct {
	Port string
}

func New(c Config) error {

	h, err := HttpHandlerFactory()
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

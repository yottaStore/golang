package yfs

import (
	"github.com/yottaStore/golang/svcs/yfs/iodriver"
	"log"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func HttpHandlerFactory(d iodriver.Iodriver) (HttpHandler, error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))

	}

	return handler, nil
}

func New(c Config, d iodriver.Iodriver) error {

	h, err := HttpHandlerFactory(d)
	if err != nil {
		log.Fatal("Error creating http handler: ", err)
	}

	http.HandleFunc("/", h)
	log.Println("Listening on port: ", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Println("Error starting server: ", err)
		return err
	}

	return nil
}

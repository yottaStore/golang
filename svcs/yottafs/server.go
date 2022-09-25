package yottafs

import (
	"errors"
	"log"
	"net/http"
	"yottafs/handlers"
	"yottafs/iodrivers"
	"yottafs/iodrivers/direct"
	"yottafs/iodrivers/dummy"
)

type Config struct {
	Namespace string
	Driver    string
	Port      string
}

func pickDriver(c Config) (iodrivers.Interface, error) {

	switch c.Driver {
	case "dummy":
		return dummy.New()
	case "direct":
		return direct.New(c.Namespace)
	default:
		return nil, errors.New("")

	}

}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello from yottafs-go v 0.0.1!"))
	if err != nil {
		log.Println("Error: version handler failed: ", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("Error: Endpoint not found"))
	if err != nil {
		log.Println("Error: not found handler failed: ", err)
	}
}

func StartServer(c Config) error {

	log.Println("Starting yottafs...")

	d, err := pickDriver(c)
	if err != nil {
		log.Println("Error instantiating driver: ", err)
		return err
	}

	httpHandler, err := handlers.HttpHandlerFactory(d)
	if err != nil {
		log.Println("Error instantiating handler: ", err)
		return err
	}

	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/yottafs/", httpHandler)
	http.HandleFunc("/", notFoundHandler)

	// Start HTTP server.
	log.Printf("listening on port %s", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Println(err)
		return err
	}

	return nil

}

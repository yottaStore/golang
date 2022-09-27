package yottafs

import (
	"log"
	"net/http"
	"yottafs/handlers"
	"yottafs/iodrivers/direct"
)

type Config struct {
	Namespace string
	Driver    string
	Port      string
	Protocol  string
}

func StartServer(c Config) error {

	log.Println("Starting yottafs...")

	dd, err := direct.New(c.Namespace)
	if err != nil {
		log.Fatal("Error instantiating driver: ", err)
	}

	httpHandler, err := handlers.HttpHandlerFactory(dd)
	if err != nil {
		log.Fatal("Error instantiating http handler: ", err)
	}

	http.HandleFunc("/yottafs", httpHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/", notFoundHandler)

	log.Println("Listening on port: ", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Println("Error starting server: ", err)
		return err
	}

	return nil
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

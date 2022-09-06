package yottadb

import (
	"log"
	"net/http"
	"yottadb/dbdriver/keyvalue"
	"yottadb/handlers"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello from yottadb-go v 0.0.1!"))
	if err != nil {
		log.Println("Error with version handler", err)
	}
}

type Config struct {
	NodeTree *[]string
	Port     string
}

func StartServer(config Config) error {

	log.Println("Starting yottadb...")

	// TODO: Switch between dbdriver
	driver, err := keyvalue.New()
	if err != nil {
		log.Println("Error instantiating driver: ", err)
		return err
	}

	// TODO: Write config to disk

	httpHandler, err := handlers.HttpHandlerFactory(driver)
	if err != nil {
		log.Println("Error instantiating handler: ", err)
		return err
	}

	http.HandleFunc("/", versionHandler)
	http.HandleFunc("/yottadb/", httpHandler)

	// Start HTTP server.
	log.Printf("listening on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

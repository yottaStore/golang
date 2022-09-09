package yottafs

import (
	"errors"
	"log"
	"net/http"
	"yottafs/handlers"
	"yottafs/iodriver"
	"yottafs/iodriver/direct"
)

func driverPicker(namespace string, driver string) (iodriver.Interface, error) {

	switch driver {
	case iodriver.Dummy:
		return nil, errors.New("Not implemented yet")
	case iodriver.Direct:
		d, err := direct.New(namespace)
		return d, err
	}
	return nil, errors.New("Driver not found")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottafs-go v 0.0.1!"))
}

type Config struct {
	NameSpace string
	Driver    string
	Port      string
}

func StartServer(config Config) error {

	log.Println("Starting yottafs...")

	var ioDriver iodriver.Interface

	// TODO: Switch between dbdriver
	ioDriver, err := driverPicker(config.NameSpace, config.Driver)
	if err != nil {
		log.Println("Error instantiating driver: ", err)
		return err
	}

	// TODO: Write config to disk

	httpHandler, err := handlers.HttpHandlerFactory(ioDriver)
	if err != nil {
		log.Println("Error instantiating handler: ", err)
		return err
	}

	http.HandleFunc("/", versionHandler)
	http.HandleFunc("/yottafs/", httpHandler)

	// Start HTTP server.
	log.Printf("listening on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

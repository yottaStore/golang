package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"yottadb/dbdrivers/document"
	"yottadb/dbdrivers/dummy"
	"yottadb/dbdrivers/keyval"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello from yottadb-go v 0.0.1!"))
	if err != nil {
		log.Println("Error with version handler", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("Error: Endpoint not found"))
	if err != nil {
		log.Println("Error: not found handler failed: ", err)
	}
}

func StartHttpServer(c Config) error {

	httpHandler, err := HttpHandlerFactory(c)
	if err != nil {
		log.Println("Error instantianting handler: ", err)
		return nil
	}

	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/yottadb/", httpHandler)
	http.HandleFunc("/", notFoundHandler)

	log.Printf("listening on port %s", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func HttpHandlerFactory(c Config) (func(http.ResponseWriter, *http.Request), error) {

	dd, err := dummy.New()
	if err != nil {
		return nil, errors.New("failed instantiating dummy driver")
	}

	kvd, err := keyval.New(c.Nodetree, c.Hashkey)
	if err != nil {
		return nil, errors.New("failed instantiating dummy driver")
	}

	docd, err := document.New(c.Nodetree, c.Hashkey)
	if err != nil {
		return nil, errors.New("failed instantiating dummy driver")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		h := strings.Split(r.URL.Path, "/")[2]
		log.Println("Url Head: ", h)

		switch h {

		case "dummy":
			dummy.Handler(w, r, dd)

		case "keyval":
			keyval.Handler(w, r, kvd)

		case "document":
			document.Handler(w, r, docd)

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("YottaDB driver not found")); err != nil {
				log.Println("ERROR: ", err)
			}
		}

	}

	return handler, nil

}

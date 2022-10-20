package http

import (
	"log"
	"net/http"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Yottafs v0.0.1!")); err != nil {
		log.Println("Error writing version handler: ", err)
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("I'm fine, thanks")); err != nil {
		log.Println("Error writing healthz handler: ", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("Can't handle the request")); err != nil {
		log.Println("Error writing healthz handler: ", err)
	}
}

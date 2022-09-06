package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottadb/dbdriver"
)

func handleHttpError(reason string, statusCode int, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(statusCode)
		if _, err := w.Write([]byte(reason)); err != nil {
			log.Println("ERROR: ", err)
		}
	}

}

func HttpHandlerFactory(d dbdriver.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		badRequest := func(err error) {
			handleHttpError(
				"Malformed YottaFs request",
				http.StatusBadRequest,
				err, w)
		}

		var req dbdriver.Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Malformed YottaFs request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		switch req.Method {

		default:
			badRequest(err)

		}

	}

	return handler, nil

}

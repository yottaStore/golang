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

		case "read":
			resp, err := d.Read(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "write":
			resp, err := d.Write(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "delete":
			resp, err := d.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("YottaDB method not found")); err != nil {
				log.Println("ERROR: ", err)
			}

		}

	}

	return handler, nil

}

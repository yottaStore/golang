package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottafs/handlers/methods"
	"yottafs/iodriver"
)

func handleHttpError2(reason string, statusCode int, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(statusCode)
		if _, err := w.Write([]byte(reason)); err != nil {
			log.Println("ERROR: ", err)
		}
	}

}

func HttpHandlerFactory(d iodriver.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req iodriver.Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			log.Println(req, err)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Malformed YottaFs request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		switch req.Method {
		case iodriver.Read:
			buff, err := methods.Read(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buff)

			if err != nil {
				log.Println("ERROR: ", err)
			}

		case iodriver.Write:
			buff, err := methods.Write(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buff)

			if err != nil {
				log.Println("ERROR: ", err)
			}
		case iodriver.Delete:
			err := d.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

		case iodriver.Append:

		default:
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				if _, err := w.Write([]byte("Method not found")); err != nil {
					log.Println("ERROR: ", err)
				}
			}

		}

	}

	return handler, nil

}

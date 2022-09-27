package handlers

import (
	"github.com/fxamacker/cbor/v2"
	"log"
	"net/http"
	"yottafs/handlers/methods"
	"yottafs/iodrivers"
)

func HttpHandlerFactory(d iodrivers.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req iodrivers.Request
		decoder := cbor.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			log.Println("Malformed request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Malformed YottaFS request")); err != nil {
				log.Println("Error responding to client: ", err)
			}
			return
		}
		log.Println("New request with method ", req.Method, " for record: ", req.Path)

		switch req.Method {
		case iodrivers.Read:
			buff, err := methods.Read(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("Error responding to client: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(buff); err != nil {
				log.Println("Error responding to client: ", err)
			}

		case iodrivers.Write:
			buff, err := methods.Write(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(buff); err != nil {
				log.Println("Error responding to client: ", err)
			}

		case iodrivers.Delete:
			_, err := d.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(nil); err != nil {
				log.Println("Error responding to client: ", err)
			}

		case "append":

		default:
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				if _, err := w.Write([]byte("Method not found")); err != nil {
					log.Println("Error responding to client: ", err)
				}
			}

		}

	}

	return handler, nil

}

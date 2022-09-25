package handlers

import (
	"github.com/fxamacker/cbor/v2"
	"log"
	"net/http"
	"strings"
	"yottafs/handlers/methods"
	"yottafs/iodrivers"
)

func findMethod(path string) string {

	tmp := strings.Split(path, "/")
	return tmp[len(tmp)-1]
}

func HttpHandlerFactory(d iodrivers.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		log.Println("New request")

		var req iodrivers.Request
		decoder := cbor.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			log.Println(req, err)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Malformed YottaFs request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		method := findMethod(r.URL.Path)

		log.Println("New request with method: ", method)

		switch method {
		case "read":
			buff, err := methods.Read(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buff)

			if err != nil {
				log.Println("ERROR: ", err)
			}

		case "create":
			buff, err := methods.Create(req, d)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buff)

			if err != nil {
				log.Println("ERROR: ", err)
			}
		case "delete":
			err := d.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

		case "append":

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

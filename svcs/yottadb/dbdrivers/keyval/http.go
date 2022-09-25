package keyval

import (
	"github.com/fxamacker/cbor/v2"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request, d Driver) {

	// TODO: handle out of index error
	h := strings.Split(r.URL.Path, "/")[3]
	log.Println("Method Head: ", h)

	var req Request
	decoder := cbor.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Request error: ", err)
		if _, err := w.Write([]byte("Malformed YottaDB keyval request")); err != nil {
			log.Println("ERROR: ", err)
		}
		return
	}

	switch h {

	case "create":
		res, err := d.Create(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Keyval create request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving keyval create request: ", err)
		}

	case "read":
		res, err := d.Read(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Keyval read request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving keyval read request: ", err)
		}

	case "update":
		res, err := d.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Keyval update request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving keyval update request: ", err)
		}

	case "delete":
		err := d.Delete(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Keyval delete request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("Error: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if err != nil {
			log.Println("Error serving keyval delete request: ", err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("YottaDB keyval method not found")); err != nil {
			log.Println("ERROR: ", err)
		}

	}

}

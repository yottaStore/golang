package document

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
		if _, err := w.Write([]byte("Malformed YottaDB document request")); err != nil {
			log.Println("ERROR: ", err)
		}
		return
	}

	switch h {

	case "createDocument":
		res, err := d.Create(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Document create request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving document create request: ", err)
		}

	case "readDocument":
		res, err := d.Read(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Document read request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving document read request: ", err)
		}

	case "updateDocument":
		res, err := d.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Document update request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving document update request: ", err)
		}

	case "deleteDocument":
		err := d.Delete(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Document delete request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("Error: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if err != nil {
			log.Println("Error serving document delete request: ", err)
		}

	case "createCollection":
		res, err := d.CollDriver.Create(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Collection create request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving collection create request: ", err)
		}

	case "readCollection":
		res, err := d.CollDriver.Read(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Collection read request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving collection read request: ", err)
		}

	case "updateCollection":
		res, err := d.CollDriver.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Collection update request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder := cbor.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			log.Println("Error serving collection update request: ", err)
		}

	case "deleteCollection":
		err := d.CollDriver.Delete(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Collection delete request error: ", err)
			if _, err := w.Write([]byte("Error serving request")); err != nil {
				log.Println("Error: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if err != nil {
			log.Println("Error serving collection delete request: ", err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("YottaDB document method not found")); err != nil {
			log.Println("ERROR: ", err)
		}

	}

}

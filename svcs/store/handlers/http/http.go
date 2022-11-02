package http

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/yottaStore/golang/svcs/store/handlers"
	"log"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func HttpHandlerFactory() (HttpHandler, error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req handlers.Request
		decoder := cbor.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			log.Println("Malformed request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Malformed YottaStore request")); err != nil {
				log.Println("Error responding to client: ", err)
			}
			return
		}
		log.Println("New request with method ", req.Method, " for record: ", req.Record)

		switch req.Driver {
		case handlers.KEY_VALUE:
		case handlers.DOCUMENT:
		case handlers.PUBSUB:

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Unknown YottaStore driver")); err != nil {
				log.Println("Error responding to client: ", err)
			}

		}

	}

	return handler, nil
}

package http

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/yottaStore/golang/svcs/yfs/handlers"
	"github.com/yottaStore/golang/svcs/yfs/io_driver"
	"log"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func HttpHandlerFactory(d io_driver.IODriver) (HttpHandler, error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req handlers.Request
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
		log.Println("New request with method ", req.Method, " for record: ", req.Record)

		switch req.Method {
		case handlers.READ:
			buff, err := d.Read(req.Record)
			if err != nil {
				// TODO: improve error handling
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("Error responding to client: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write(buff); err != nil {
				log.Println("Error responding to client: ", err)
			}
		case handlers.CREATE:
			err := d.Create(req.Record, req.Data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Create request failed")); err != nil {
					log.Println("Error responding to client:", err)
				}
			}
		case handlers.DELETE:
			err := d.Delete(req.Record)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("Error responding to client:", err)
				}
			}

		case handlers.APPEND:
			err := d.Append(req.Record, req.Data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Append request failed")); err != nil {
					log.Println("Error responding to client:", err)
				}
			}
		case handlers.COMPACT:
			err := d.Compact(req.Record)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Compact request failed")); err != nil {
					log.Println("Error responding to client:", err)
				}
			}
		case handlers.MERGE:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Method not implemented yet")); err != nil {
				log.Println("Error responding to client: ", err)
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Unknown YottaFS method")); err != nil {
				log.Println("Error responding to client: ", err)
			}

		}

	}

	return handler, nil
}

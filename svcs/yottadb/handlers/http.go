package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/dbdriver/document"
	"yottadb/dbdriver/keyvalue"
)

type Config struct {
	NodeTree *[]string
	Port     string
	HashKey  string
}

func HttpHandlerFactory(config Config) (func(http.ResponseWriter, *http.Request), error) {

	dd, err := document.New(config.HashKey, config.NodeTree)
	kvd, err := keyvalue.New(config.HashKey, config.NodeTree)
	if err != nil {
		log.Println("Error instantiating driver: ", err)
		return nil, err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req dbdriver.Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Request error: ", err)
			if _, err := w.Write([]byte("Malformed YottaFs request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		log.Println("Request: ", req)

		switch req.Method {

		case "readDocument":
			resp, err := dd.ReadDocument(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "writeDocument":
			resp, err := dd.WriteDocument(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "deleteDocument":
			resp, err := dd.DeleteDocument(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "createCollection":
			resp, err := dd.CreateCollection(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Create collection request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "deleteCollection":
			resp, err := dd.DeleteCollection(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete collection request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "readKV":
			resp, err := kvd.Read(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "writeKV":
			resp, err := kvd.Write(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
				log.Println("ERROR: ", err)
			}

		case "deleteKV":
			resp, err := kvd.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write([]byte(resp.Data)); err != nil {
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

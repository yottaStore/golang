package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottadb/dbdriver"
	KVD "yottadb/dbdriver/keyvalue"
	"yottadb/handlers/keyvalue"
)

type Config struct {
	NodeTree *[]string
	Port     string
	HashKey  string
}

func HttpHandlerFactory(config Config) (func(http.ResponseWriter, *http.Request), error) {

	//dd, err := document.New(config.HashKey, config.NodeTree)
	kvd, err := KVD.New(config.HashKey, config.NodeTree)
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
			if _, err := w.Write([]byte("Malformed YottaDB request")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		log.Println("Request: ", req)

		switch req.Driver {

		case "document":

		case "collection":

		case "keyvalue":
			keyvalue.Handler(w, req, kvd)

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("YottaDB driver not found")); err != nil {
				log.Println("ERROR: ", err)
			}
		}

	}

	return handler, nil

}

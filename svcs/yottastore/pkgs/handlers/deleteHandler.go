package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottadb"
)

func DeleteHandlerFactory(dbDriver yottadb.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var req WriteRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaFs write request"))
			return
		}

		ioReq := yottadb.WriteRequest{
			Path:             req.Path,
			Data:             req.Data,
			CreateCollection: req.CreatePath,
		}

		if err := dbDriver.Delete(ioReq); err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaFs delete failed for: " + req.Path))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delete successful"))
	}

	return handler, nil
}

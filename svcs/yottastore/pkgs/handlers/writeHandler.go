package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yottastore/pkgs/yottadb"
)

type WriteRequest struct {
	Path       string `json:"Path"`
	Data       []byte `json:"Data"`
	Append     bool
	CreatePath bool
}

func WriteHandlerFactory(dbDriver yottadb.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var req WriteRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaStore write request"))
			return
		}

		ioReq := yottadb.WriteRequest{
			Path:       req.Path,
			Data:       req.Data,
			CreatePath: req.CreatePath,
		}

		var err error
		if req.Append {
			fmt.Println("Appending")
			_, err = dbDriver.Append(ioReq)
		} else {
			_, err = dbDriver.Write(ioReq)
		}

		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaStore write failed for: " + req.Path))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Write successful"))
	}

	return handler, nil
}

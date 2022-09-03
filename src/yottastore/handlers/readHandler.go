package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottastore/dbdrivers"
)

type ReadRequest struct {
	Path    string      `json:"Path"`
	Options interface{} `json:"Options"`
}

type ReadResponse struct {
	Path    string      `json:"Path"`
	Data    string      `json:"Data"`
	Options interface{} `json:"Options"`
}

func ReadHandlerFactory(driverInterface dbdrivers.Interface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req ReadRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaStore read request"))
			return
		}

		ioReq := dbdrivers.ReadRequest{
			Path: req.Path,
		}

		resp, err := driverInterface.Read(ioReq)
		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaStore read failed for: " + req.Path))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(resp.Data)
	}

	return handler, nil
}

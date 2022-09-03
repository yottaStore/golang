package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"yottafs/iodrivers"
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

func ReadHandlerFactory(ioDriver iodrivers.IoDriverInterface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req ReadRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaFs read request"))
			return
		}

		ioReq := iodrivers.IoReadRequest{
			Path: req.Path,
		}

		resp, err := ioDriver.Read(ioReq)
		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaFs read failed for: " + req.Path))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(resp.Data)
	}

	return handler, nil
}

package net

import (
	"encoding/json"
	"log"
	"net/http"
	"yottafs/iodrivers"
)

type YfsWriteRequest struct {
	Path string `json:"Path"`
	Data []byte `json:"Data"`
}

func WriteHandlerFactory(ioDriver iodrivers.IoDriverInterface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var req YfsWriteRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaFs write request"))
		}

		ioReq := iodrivers.IoWriteRequest{
			Path: req.Path,
			Data: req.Data,
		}

		_, err := ioDriver.Write(ioReq)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaFs read failed for: " + req.Path))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Write successful"))
	}

	return handler, nil
}

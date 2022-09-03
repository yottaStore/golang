package handlers

import (
	"encoding/json"
	"fmt"
	"libs/iodrivers"
	"log"
	"net/http"
)

type WriteRequest struct {
	Path       string `json:"Path"`
	Data       []byte `json:"Data"`
	Append     bool
	CreatePath bool
}

func WriteHandlerFactory(ioDriver iodrivers.IoDriverInterface) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var req WriteRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaFs write request"))
			return
		}

		ioReq := iodrivers.IoWriteRequest{
			Path:       req.Path,
			Data:       req.Data,
			CreatePath: req.CreatePath,
		}

		var err error
		if req.Append {
			fmt.Println("Appending")
			_, err = ioDriver.Append(ioReq)
		} else {
			_, err = ioDriver.Write(ioReq)
		}

		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("YottaFs write failed for: " + req.Path))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Write successful"))
	}

	return handler, nil
}

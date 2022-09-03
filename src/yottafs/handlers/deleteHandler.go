package handlers

import (
	"encoding/json"
	"libs/iodrivers"
	"log"
	"net/http"
)

func DeleteHandlerFactory(ioDriver iodrivers.IoDriverInterface) (func(http.ResponseWriter, *http.Request), error) {

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

		if err := ioDriver.Delete(ioReq); err != nil {
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

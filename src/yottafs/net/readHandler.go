package net

import (
	"encoding/json"
	"net/http"
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

func ReadHandlerFactory(ioDriver interface{}) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		var req ReadRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YottaFs read request"))
		}

	}

	return handler, nil
}

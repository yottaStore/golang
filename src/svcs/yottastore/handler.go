package yottastore

import (
	"encoding/json"
	"net/http"
	"strings"
)

type StoreRequest struct {
	Record string
	Data   string
}

func HttpHandlerFactory() (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		endpoint := strings.Split(r.URL.String(), "/")[2]

		decoder := json.NewDecoder(r.Body)
		var storeReq StoreRequest
		err := decoder.Decode(&storeReq)
		if err != nil {
			w.Write([]byte("Malformed request"))
			return
		}

		switch endpoint {
		case "read":

		default:
			w.Write([]byte("Method not found"))
		}

	}

	return handler, nil

}

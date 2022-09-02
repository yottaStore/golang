package yottastore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"yottaStore/yottaStore-go/src/pkgs/rendezvous"
	"yottaStore/yottaStore-go/src/svcs/yottastore/read"
	"yottaStore/yottaStore-go/src/svcs/yottastore/write"
)

type StoreRequest struct {
	Record string
	Data   []byte
}

func HttpHandlerFactory(nodes *[]string, decoder interface{}) (func(http.ResponseWriter, *http.Request), error) {

	// TODO: pick decoder
	// TODO: pick driver

	handler := func(w http.ResponseWriter, r *http.Request) {

		endpoint := strings.Split(r.URL.String(), "/")[2]

		decoder := json.NewDecoder(r.Body)
		var storeReq StoreRequest
		err := decoder.Decode(&storeReq)
		if err != nil {
			w.Write([]byte("Malformed request"))
			return
		}

		node, err := rendezvous.Rendezvous(storeReq.Record, *nodes)
		if err != nil {
			w.Write([]byte("Error with rendezvous"))
			return
		}

		switch endpoint {
		case "read":

			record, err := read.Read(storeReq.Record, node)
			if err != nil {
				w.Write([]byte("Error with read"))
				return
			}
			fmt.Println(string(record))

			w.Write(record)

		case "write":

			record, err := write.WriteNew(storeReq.Record, node, storeReq.Data)
			if err != nil {
				w.Write([]byte("Error with read"))
				return
			}

			fmt.Println(string(record.([]byte)))

		default:
			w.Write([]byte("Method not found"))
		}

	}

	return handler, nil

}

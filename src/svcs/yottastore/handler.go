package yottastore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"yottaStore/yottaStore-go/src/pkgs/rendezvous"
	"yottaStore/yottaStore-go/src/pkgs/yottadb"
	"yottaStore/yottaStore-go/src/pkgs/yottapack"
	"yottaStore/yottaStore-go/src/svcs/yottastore/read"
	"yottaStore/yottaStore-go/src/svcs/yottastore/write"
)

type StoreRequest struct {
	Record string
	Data   string
}

type HandlerConfig[T any] struct {
	Nodes  *[]string
	Driver yottadb.DbDriver
	Packer yottapack.Packer[T]
}

func HttpHandlerFactory[T any](nodes *[]string, config HandlerConfig[T]) (func(http.ResponseWriter, *http.Request), error) {

	// TODO: pick decoder
	// TODO: pick driver

	//dbDriver := config.Driver
	//packer := config.Packer

	handler := func(w http.ResponseWriter, r *http.Request) {

		endpoint := strings.Split(r.URL.String(), "/")[2]
		decoder := json.NewDecoder(r.Body)
		var storeReq StoreRequest
		err := decoder.Decode(&storeReq)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Malformed request"))
			return
		}

		parsedRecord, err := rendezvous.ParseRecord(storeReq.Record)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Malformed record"))
			return
		}
		fmt.Println(parsedRecord)

		node, err := rendezvous.Rendezvous(parsedRecord, *nodes)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Error with rendezvous"))
			return
		}

		fmt.Println(node)

		switch endpoint {
		case "read":

			record, err := read.Read(parsedRecord.RecordIdentifier, node)
			if err != nil {
				log.Println(err)
				w.Write([]byte("Error with read"))
				return
			}
			fmt.Println(string(record))

			w.Write(record)

		case "write":

			// TODO: handle non existing dir
			record, err := write.WriteNew(parsedRecord.RecordIdentifier, node, []byte(storeReq.Data))
			if err != nil {
				log.Println(err)
				w.Write([]byte("Error with write"))
				return
			}

			w.Write([]byte("Write successful"))

			fmt.Println(record)

		default:
			w.Write([]byte("Method not found"))
		}

	}

	return handler, nil

}

package yfs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
)

type Request struct {
	Path    string
	Data    string
	Method  string
	Options struct {
	}
}

func HttpHandlerFactory(ioDriver iodrivers.IoDriverInterface) (handler func(http.ResponseWriter, *http.Request), err error) {

	// TODO: handle different node types (linear,random)
	// TODO: handle locks

	handler = func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var storeReq Request
		err := decoder.Decode(&storeReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed request"))
			return
		}

		switch storeReq.Method {
		case "read":
			req := iodrivers.IoReadRequest{
				storeReq.Path,
				"read",
			}
			if resp, err := ioDriver.ReadAll(req); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Read failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(resp.Data)
			}
		case "write":
			req := iodrivers.IoWriteRequest{
				Path:   storeReq.Path,
				Data:   []byte(storeReq.Data),
				Method: "write",
			}
			if err := ioDriver.Write(req); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Write failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Write successful"))
			}
		case "append":
			req := iodrivers.IoWriteRequest{
				Path:   storeReq.Path,
				Data:   []byte(storeReq.Data),
				Method: "append",
			}
			if err := ioDriver.Append(req); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Append failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Append successful"))
			}
		case "delete":
			req := iodrivers.IoWriteRequest{
				Path:   storeReq.Path,
				Method: "delete",
			}
			if err := ioDriver.Delete(req); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Delete failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Delete successful"))
			}
		case "cas":
			req := iodrivers.IoWriteRequest{
				Path:   storeReq.Path,
				Data:   []byte(storeReq.Data),
				Method: "append",
			}
			if err := ioDriver.CompareAndSwap(req); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("CompareAndSwap failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("CompareAndSwap successful"))
			}
		case "caa":
			req := iodrivers.IoWriteRequest{
				Path:   storeReq.Path,
				Data:   []byte(storeReq.Data),
				Method: "append",
			}
			if err := ioDriver.CompareAndAppend(req); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("CompareAndAppend failed"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("CompareAndAppend successful"))
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Method not found"))
		}

	}

	return handler, nil
}

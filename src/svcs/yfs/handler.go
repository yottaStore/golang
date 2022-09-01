package yfs

import (
	"encoding/json"
	"net/http"
	"strings"
	"yottaStore/yottaStore-go/src/pkgs/drivers"
)

type yfsPath interface {
	string | int
}

type yfsRequest[T yfsPath] struct {
	Path T
	Data string
}

func HttpHandlerFactory(ioDriver drivers.IoDriver) (handler func(http.ResponseWriter, *http.Request), err error) {

	// TODO: handle different node types (linear,random)
	// TODO: handle locks

	handler = func(w http.ResponseWriter, r *http.Request) {

		endpoint := strings.Split(r.URL.String(), "/")[2]
		decoder := json.NewDecoder(r.Body)
		// TOOD: switch according to ioDriver
		var storeReq yfsRequest[string]
		err := decoder.Decode(&storeReq)
		if err != nil {
			w.Write([]byte("Malformed request"))
			return
		}

		switch endpoint {
		case "read":
			if b, err := ioDriver.ReadAll(storeReq.Path); err != nil {
				w.Write([]byte("Read failed"))
			} else {
				w.Write(b)
			}
		case "write":
			if err := ioDriver.Write(storeReq.Path, []byte(storeReq.Data)); err != nil {
				w.Write([]byte("Write failed"))
			} else {
				w.Write([]byte("Write successful"))
			}
		case "append":
			if err := ioDriver.Append(storeReq.Path, []byte(storeReq.Data)); err != nil {
				w.Write([]byte("Append failed"))
			} else {
				w.Write([]byte("Append successful"))
			}
		case "delete":
			if err := ioDriver.Delete(storeReq.Path); err != nil {
				w.Write([]byte("Delete failed"))
			} else {
				w.Write([]byte("Delete successful"))
			}
		case "cas":
			if err := ioDriver.CompareAndSwap(storeReq.Path, []byte(storeReq.Data)); err != nil {
				w.Write([]byte("CompareAndSwap failed"))
			} else {
				w.Write([]byte("CompareAndSwap successful"))
			}
		case "caa":
			if err := ioDriver.CompareAndAppend(storeReq.Path, []byte(storeReq.Data)); err != nil {
				w.Write([]byte("CompareAndAppend failed"))
			} else {
				w.Write([]byte("CompareAndAppend successful"))
			}
		default:
			w.Write([]byte("Method not found"))
		}

	}

	return handler, nil
}

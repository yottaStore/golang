package yottastore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	direct2 "yottaStore/yottaStore-go/src/libs/yfs/drivers/direct"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/read"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/write"
)

type StoreRequest struct {
	Path string
	Data string
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	endpoint := strings.Split(r.URL.String(), "/")[2]

	decoder := json.NewDecoder(r.Body)
	var storeReq StoreRequest
	err := decoder.Decode(&storeReq)
	if err != nil {
		w.Write([]byte("Malformed request"))
		return
	}

	fmt.Println(storeReq, endpoint)

	switch endpoint {
	case "read":
		{
			b, err := read.ReadAll(storeReq.Path)

			if err != nil {
				w.Write([]byte("Read failed"))
				return
			}
			w.Write(b)
		}
	case "write":
		{
			err := direct2.Write(storeReq.Path, []byte(storeReq.Data))
			if err != nil {
				w.Write([]byte("Write failed"))
				return
			}
			w.Write([]byte("Write successful"))
		}
	case "append":
		{
			err := direct2.Append(storeReq.Path, []byte(storeReq.Data))
			if err != nil {
				fmt.Println(err)
				w.Write([]byte("Append failed"))
				return
			}
			w.Write([]byte("Append successful"))
		}
	case "delete":
		{
			err := write.Delete(storeReq.Path)
			if err != nil {
				w.Write([]byte("Delete failed"))
				return
			}
			w.Write([]byte("Delete Successful"))
		}
	case "":
		{
			w.Write([]byte("Hello yottastore!"))

		}
	default:
		w.Write([]byte("Method not found!"))
	}

}

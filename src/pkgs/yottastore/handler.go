package yottastore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
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
			b, err := direct.ReadAll(storeReq.Path)

			if err != nil {
				w.Write([]byte("Read failed"))
				return
			}
			w.Write(b)
		}
	case "write":
		{
			err := direct.Write(storeReq.Path, []byte(storeReq.Data))
			if err != nil {
				w.Write([]byte("Write failed"))
				return
			}
			w.Write([]byte("Write successful"))
		}
	case "append":
		{
			err := direct.Append(storeReq.Path, []byte(storeReq.Data))
			if err != nil {
				fmt.Println(err)
				w.Write([]byte("Append failed"))
				return
			}
			w.Write([]byte("Append successful"))
		}
	case "delete":
		{
			err := direct.Delete(storeReq.Path)
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

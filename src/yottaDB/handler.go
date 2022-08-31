package yottaDB

import (
	"net/http"
	"strings"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	endpoints := strings.Split(r.URL.String(), "/")
	endpoint := endpoints[len(endpoints)-1]

	switch endpoint {
	case "read":
		{

		}
	case "write":
		{

		}
	case "update":
		{

		}
	case "delete":
		{

		}
	case "":
		{
			w.Write([]byte("Hello ydb!"))

		}
	default:
		w.Write([]byte("Method not found!"))
	}

}

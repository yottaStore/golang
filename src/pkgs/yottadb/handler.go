package yottadb

import (
	"net/http"
	"strings"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	endpoint := strings.Split(r.URL.String(), "/")[2]

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

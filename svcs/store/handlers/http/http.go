package http

import (
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func HttpHandlerFactory() (HttpHandler, error) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}

	return handler, nil
}

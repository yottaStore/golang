package yfs

import "net/http"

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello yfs!"))

}

package gossip

import "net/http"

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	healthString := []byte("ok!")
	w.Write(healthString)
}

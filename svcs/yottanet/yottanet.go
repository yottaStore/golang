package yottanet

import "net/http"

func YottanetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottafs-go Gossip!"))
}

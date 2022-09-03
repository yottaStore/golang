package gossip

import "net/http"

func GossipHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottafs-go Gossip!"))
}

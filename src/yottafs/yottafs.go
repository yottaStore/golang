package main

import (
	"gossip"
	"log"
	"net/http"
	"os"
	"yottafs/net"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottafs-go v 0.0.1!"))
}

func main() {

	// Parse config

	//

	readHandler, err := net.ReadHandlerFactory(0)
	if err != nil {
		log.Fatal(err)
	}
	writeHandler, err := net.ReadHandlerFactory(0)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/yottafs/read", readHandler)
	http.HandleFunc("/yottafs/write", writeHandler)

	http.HandleFunc("/", versionHandler)
	http.HandleFunc("/gossip/", gossip.GossipHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

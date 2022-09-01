package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/yfs"
	"yottaStore/yottaStore-go/src/pkgs/gossip"
	"yottaStore/yottaStore-go/src/pkgs/yottadb"
)

func main() {
	log.Print("starting yottadb...")
	http.HandleFunc("/ydb/", yottadb.HttpHandler)
	http.HandleFunc("/yfs/", yfs.HttpHandler)
	http.HandleFunc("/gossip/", gossip.HttpHandler)
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	helloString := []byte("Hello from yottadb-go v 0.0.1!")
	w.Write(helloString)
}

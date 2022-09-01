package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/gossip"
	"yottaStore/yottaStore-go/src/pkgs/yottastore"
)

func main() {
	log.Print("starting yottaStore...")

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		helloString := []byte("Hello from yottaStore-go v 0.0.1!")
		w.Write(helloString)
	}

	http.HandleFunc("/store/", yottastore.HttpHandler)
	http.HandleFunc("/gossip/", gossip.HttpHandler)
	http.HandleFunc("/", versionHandler)

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

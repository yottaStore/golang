package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/gossip"
	"yottaStore/yottaStore-go/src/pkgs/yfs"
)

func main() {

	log.Print("starting yottaStore...")

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		helloString := []byte("Hello from yfs-go v 0.0.1!")
		w.Write(helloString)
	}

	opts := yfs.YfsSetupOptions{}
	// TODO: setup io driver here
	yfsHandler, err := yfs.HttpHandlerFactory(opts)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/yfs/", yfsHandler)
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

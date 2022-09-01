package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/config"
	"yottaStore/yottaStore-go/src/pkgs/drivers"
	"yottaStore/yottaStore-go/src/pkgs/gossip"
	"yottaStore/yottaStore-go/src/svcs/yfs"
)

func main() {
	log.Print("starting yottaStore...")

	config, err := config.ParseConfig[drivers.Config]()
	if err != nil {
		log.Fatal(err)
	}

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		helloString := []byte("Hello from yfs-go v 0.0.1!")
		w.Write(helloString)
	}

	ioDriver, err := yfs.New(config)
	if err != nil {
		log.Fatal(err)
	}

	yfsHandler, err := yfs.HttpHandlerFactory(ioDriver)
	if err != nil {
		log.Fatal(err)
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

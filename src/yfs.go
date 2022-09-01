package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/drivers"
	"yottaStore/yottaStore-go/src/libs/gossip"
	"yottaStore/yottaStore-go/src/pkgs/yfs"
	"yottaStore/yottaStore-go/src/utils/config"
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

	yfsHandler, err := yfs.HttpHandlerFactory(config, ioDriver)
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

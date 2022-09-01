package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/config"
	"yottaStore/yottaStore-go/src/pkgs/drivers"
	"yottaStore/yottaStore-go/src/pkgs/gossip"
	"yottaStore/yottaStore-go/src/svcs/yottastore"
)

func main() {
	log.Print("starting yottaStore...")

	_, err := config.ParseConfig[drivers.Config]()
	if err != nil {
		log.Fatal(err)
	}

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		helloString := []byte("Hello from yottaStore-go v 0.0.1!")
		w.Write(helloString)
	}

	// TODO: parse config
	yottastore.New()

	handler, err := yottastore.HttpHandlerFactory()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/store/", handler)
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

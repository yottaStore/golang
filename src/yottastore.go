package main

import (
	"log"
	"net/http"
	"os"
	"yottaStore/yottaStore-go/src/libs/config"
	"yottaStore/yottaStore-go/src/pkgs/gossip"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers"
	"yottaStore/yottaStore-go/src/pkgs/yottadb"
	"yottaStore/yottaStore-go/src/pkgs/yottapack"
	"yottaStore/yottaStore-go/src/svcs/yottastore"
)

func main() {
	log.Print("starting yottaStore...")

	_, err := config.ParseConfig[iodrivers.Config]()
	if err != nil {
		log.Fatal(err)
	}

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		helloString := []byte("Hello from yottaStore-go v 0.0.1!")
		w.Write(helloString)
	}

	// TODO: parse config
	// TODO: pick decoder
	yottastore.New()
	// TODO: get list of nodes
	nodes := []string{"http://localhost:8081/yottafs/"}

	config := yottastore.HandlerConfig[interface{}]{
		Nodes:  &nodes,
		Driver: yottadb.DbDriver{},
		Packer: yottapack.Packer[interface{}]{},
	}
	handler, err := yottastore.
		HttpHandlerFactory(&nodes, config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/yottastore/", handler)
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

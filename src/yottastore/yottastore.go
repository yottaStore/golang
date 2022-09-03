package main

import (
	"libs/yottadb/keyvalue"
	"log"
	"net/http"
	"os"
	"yottanet"
	"yottastore/handlers"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottastore-go v 0.0.1!"))
}

func main() {

	log.Println("Starting yottastore...")

	// TODO: parse config

	// TODO: gossip peers
	nodes := []string{"http://localhost:8080"}

	// TODO: Switch between drivers
	dbDriver, err := keyvalue.New(&nodes)
	if err != nil {
		log.Fatal("Error instantiating driver: ", err)
	}

	// TODO: aggregate handlers
	readHandler, err := handlers.ReadHandlerFactory(dbDriver)
	if err != nil {
		log.Fatal("Error instantiating read handler: ", err)
	}
	writeHandler, err := handlers.WriteHandlerFactory(dbDriver)
	if err != nil {
		log.Fatal("Error instantiating write handler: ", err)
	}
	deleteHandler, err := handlers.DeleteHandlerFactory(dbDriver)
	if err != nil {
		log.Fatal("Error instantiating delete handler: ", err)
	}

	http.HandleFunc("/yottastore/read", readHandler)
	http.HandleFunc("/yottastore/write", writeHandler)
	http.HandleFunc("/yottastore/delete", deleteHandler)

	http.HandleFunc("/", versionHandler)
	http.HandleFunc("/gossip/", yottanet.YottanetHandler)

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

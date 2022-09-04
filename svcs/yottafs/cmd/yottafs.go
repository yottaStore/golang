package main

import (
	"log"
	"net/http"
	"os"
	"yottafs/pkgs/handlers"
	"yottafs/pkgs/iodrivers/direct"
	"yottanet"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottafs-go v 0.0.1!"))
}

func main() {

	log.Println("Starting yottafs...")

	// TODO: Parse config
	nameSpace := "/tmp/yottafs"
	// driver := "direct"

	// TODO: Switch between drivers
	ioDriver, err := direct.New(nameSpace)
	if err != nil {
		log.Fatal("Error instantiating driver: ", err)
	}

	// TODO: aggregate handlers
	readHandler, err := handlers.ReadHandlerFactory(ioDriver)
	if err != nil {
		log.Fatal("Error instantiating read handler: ", err)
	}
	writeHandler, err := handlers.WriteHandlerFactory(ioDriver)
	if err != nil {
		log.Fatal("Error instantiating write handler: ", err)
	}
	deleteHandler, err := handlers.DeleteHandlerFactory(ioDriver)
	if err != nil {
		log.Fatal("Error instantiating delete handler: ", err)
	}

	http.HandleFunc("/yottafs/read", readHandler)
	http.HandleFunc("/yottafs/write", writeHandler)
	http.HandleFunc("/yottafs/delete", deleteHandler)

	http.HandleFunc("/", versionHandler)
	http.HandleFunc("/gossip/", yottanet.YottanetHandler)

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

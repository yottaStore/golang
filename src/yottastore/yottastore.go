package yottastore

import (
	"log"
	"net/http"
	"os"
	"yottanet"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from yottastore-go v 0.0.1!"))
}

func main() {

	log.Println("Starting yottastore...")

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

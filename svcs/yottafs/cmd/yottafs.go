package main

import (
	"log"
	"os"
	"yottafs"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

	conf := yottafs.Config{
		"/tmp/yottafs",
		"direct",
		port,
	}

	err := yottafs.StartServer(conf)
	if err != nil {
		log.Fatal("Error with server: ", err)
	}

}

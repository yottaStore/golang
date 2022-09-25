package main

import (
	"log"
	"yottadb/handlers"
)

func main() {

	nodetree := []string{"http://localhost:8081"}

	c := handlers.Config{
		Port:     "8080",
		Hashkey:  "38udhjdhd",
		Nodetree: &nodetree,
		Protocol: "http",
	}

	err := handlers.StartServer(c)
	if err != nil {
		log.Println("Error starting server: ", err)
	}

}

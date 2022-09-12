package main

import (
	"log"
	"yottadb"
	"yottadb/handlers"
)

func main() {

	hashKey := "SCtnOxT8NgNRXFkO"
	nodeTree := []string{"http://localhost:8081"}

	conf := handlers.Config{
		Port:     "8080",
		NodeTree: &nodeTree,
		HashKey:  hashKey,
	}

	err := yottadb.StartServer(conf)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}

}

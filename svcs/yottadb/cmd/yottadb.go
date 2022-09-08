package main

import (
	"log"
	"yottadb"
)

func main() {

	hashKey := "SCtnOxT8NgNRXFkO"
	nodeTree := []string{"http://localhost:8081"}

	conf := yottadb.Config{
		Port:     "8080",
		NodeTree: &nodeTree,
		HashKey:  hashKey,
	}

	err := yottadb.StartServer(conf)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}

}

package main

import (
	"log"
	"yottadb"
)

func main() {

	conf := yottadb.Config{
		Port: "8080",
	}
	
	err := yottadb.StartServer(conf)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}

}

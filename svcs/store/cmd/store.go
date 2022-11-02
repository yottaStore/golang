package main

import (
	"github.com/yottaStore/golang/svcs/store"
	"log"
)

func main() {

	c := store.Config{
		Protocol: "htt[",
		Port:     "8080",
	}

	err := store.Start(c)
	if err != nil {
		log.Fatal("Error starting store: ", err)
	}

}

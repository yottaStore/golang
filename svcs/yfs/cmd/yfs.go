package main

import (
	"github.com/yottaStore/golang/svcs/yfs"
	"log"
)

func main() {

	c := yfs.Config{
		Namespace: "/tmp/yfs",
		Port:      ":8081",
	}

	err := yfs.Start(c)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
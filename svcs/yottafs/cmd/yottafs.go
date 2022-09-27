package main

import (
	"log"
	"yottafs"
)

func main() {

	conf := yottafs.Config{
		"/tmp/yottafs",
		"direct",
		"8081",
		"http"}

	err := yottafs.StartServer(conf)
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}

}

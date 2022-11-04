package main

import (
	"github.com/yottaStore/golang/svcs/self"
	"log"
)

func main() {

	var c self.Config
	err := self.Start(c)
	if err != nil {
		log.Fatal("Error starting self service: ", err)
	}

}

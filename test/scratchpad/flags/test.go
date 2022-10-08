package main

import (
	"log"
)

func main() {

	ver := uint16(4088)
	log.Println("Version: ", ver)
	log.Printf("Version: %b", 1)

	log.Printf("Version: %b", ver)

}

package main

import (
	"log"
	"yottafs/client/methods"
)

func main() {

	node := "http://localhost:8081"
	path := "/test"
	payload := "Hello worlder @#$! \n From me!"
	data := []byte(payload)

	log.Println("Writing...")

	res, err := methods.Write(data, path, node)
	if err != nil {
		log.Fatal("Error writing: ", err)
	}
	log.Println("Write response: ", res)

	log.Println("Reading...")

	res, err = methods.Read(path, node)
	if err != nil {
		log.Fatal("Error Reading: ", err)
	}
	log.Println("Read response: ", res)
	log.Println("Data response: ", string(res.Data))

}

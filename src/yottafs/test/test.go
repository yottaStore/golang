package main

import (
	"log"
	"yottaclient"
)

func main() {

	path := "/test.txt"
	node := "http://localhost:8081"
	data := []byte("Hello from yottaclient write! \n")
	appendData := []byte("Hello from append! \n")

	if err := yottaclient.YfsWrite(
		path, data, node); err != nil {
		log.Fatal(err)
	}

	buff, err := yottaclient.YfsRead(
		path, node)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buff))

	if err := yottaclient.YfsAppend(
		path, appendData, node); err != nil {
		log.Fatal(err)
	}

	buff, err = yottaclient.YfsRead(
		path, node)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buff))

	err = yottaclient.YfsDelete(
		path, node)
	if err != nil {
		log.Fatal(err)
	}

	buff, err = yottaclient.YfsRead(
		path, node)
	if err != nil {
		log.Println("Delete successful!\n")
		log.Println(err)
	}
	log.Println("After delete: ", string(buff))
}

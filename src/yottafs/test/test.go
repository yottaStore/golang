package main

import (
	"fmt"
	"log"
	"yottaclient"
)

func main() {

	path := "/test.txt"
	node := "http://localhost:8081"
	data := []byte("Hello from yottaclient write! \n")

	if err := yottaclient.YfsWrite(
		path, data, node); err != nil {
		log.Fatal(err)
	}

	buff, err := yottaclient.YfsRead(
		path, node)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(buff))
}

package main

import (
	"github.com/yottaStore/golang/svcs/yfs/client"
	"log"
)

func main() {

	record := "test"
	url := "http://localhost:8081"

	data, err := client.Read(record, url, 0)
	if err != nil {
		log.Fatal("Error reading record: ", err)
	}

	log.Println("Data: ", string(data))

}

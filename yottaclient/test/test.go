package main

import (
	"log"
	"yottaclient"
)

func main() {

	client, err := yottaclient.NewClient("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	record := "yotta@testCollection/testRecord"
	data := []byte("Hello borld")

	if err := client.Write(record, data); err != nil {
		log.Fatalln(err)
	}

	readData, err := client.Read(record)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(readData))
}

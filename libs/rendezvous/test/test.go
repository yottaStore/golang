package main

import (
	"log"
	"rendezvous"
)

func main() {

	record := "account@driver:collectionName/recordName/recordRow"

	p, err := rendezvous.ParseRecord(record)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(p)

}

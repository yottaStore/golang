package main

import (
	"libs/rendezvous"
	"log"
)

func main() {

	recordString := "account@driver:tableName/recordName:recordRow"
	nodes := []string{"hello", "world"}

	parsedRecord, err := rendezvous.ParseRecord(recordString)
	if err != nil {
		log.Fatalln(err)
	}
	rendezvousFunction := rendezvous.NewRendezvous("83838jdjdhha")

	node, err := rendezvousFunction(parsedRecord, nodes)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(node)
}

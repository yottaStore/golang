package main

import (
	"fmt"
	"log"
	"yottadb/rendezvous"
)

func main() {

	recordString := "account@driver:collectionName/recordName/recordRow"
	nodesMap := [][]string{{"eu-west-1", "eu-east-1"}, {"london", "frankfurt"}}

	finder := rendezvous.Finder{"83838383"}

	parsedRecord, err := rendezvous.ParseRecord(recordString)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(parsedRecord)

	nodes, err := finder.FindPool(parsedRecord, nodesMap[0], 2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(nodes)
	shards, err := finder.FindShard(parsedRecord, nodes, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(shards)

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"yottaclient/yottadb"
	"yottadb/dbdriver"
)

func main() {

	client, err := yottadb.New("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	recordPath := "testAccount@testCollection/testRecord"
	opts := dbdriver.RendezvousOpts{
		Sharding:    1,
		Replication: 1}

	log.Println("Create collection")
	resp, err := client.CreateCollection(recordPath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)

	log.Println("Write document")
	data := []byte("helloworld")
	resp, err = client.WriteDocument(recordPath, data, opts)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Read document")
	resp, err = client.ReadDocument(recordPath, opts)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedResp struct {
		Path   string
		Method string
		Data   []byte
	}

	err = json.Unmarshal(resp, &parsedResp)
	if err != nil {
		log.Fatal(err)
	}

	res := make([]byte, 30)
	//_, err = base64.StdEncoding.Decode(res, parsedResp.Data)

	fmt.Println(parsedResp.Path)
	fmt.Println(parsedResp.Method)
	fmt.Println(string(parsedResp.Data))
	fmt.Println(string(res))

	err = client.DeleteDocument(recordPath, opts)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.ReadDocument(recordPath, opts)
	if err == nil {
		log.Fatalln("Record shouldn't exist")
	}

	_, err = client.WriteDocument(recordPath, data, opts)
	if err != nil {
		log.Fatalln(err)
	}

}

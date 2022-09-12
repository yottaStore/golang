package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
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
	data := []byte("helloworld @(#(#!")
	opts := dbdriver.RendezvousOpts{
		Sharding:    1,
		Replication: 1}

	resp, err := client.Write(recordPath, data, opts)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = client.Read(recordPath, opts)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedResp struct {
		Path   string
		Method string
		Data   []byte
	}

	err = cbor.Unmarshal(resp, &parsedResp)
	if err != nil {
		log.Fatal(err)
	}

	res := make([]byte, 30)
	//_, err = base64.StdEncoding.Decode(res, parsedResp.Data)

	fmt.Println(parsedResp.Path)
	fmt.Println(parsedResp.Method)
	fmt.Println(string(parsedResp.Data))
	fmt.Println(string(res))

	err = client.Delete(recordPath, opts)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Read(recordPath, opts)
	if err == nil {
		log.Fatalln("Record shouldn't exist")
	}

	_, err = client.Write(recordPath, data, opts)
	if err != nil {
		log.Fatalln(err)
	}

}

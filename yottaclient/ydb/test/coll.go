package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"log"
	"yottaclient/ydb"
)

func main() {

	server := "http://localhost:8080"

	client, err := ydb.NewDoc(server)
	if err != nil {
		log.Fatal("Error instantiating client: ", err)
	}

	payload := "Hello worlder @#$! \n From me!"
	data := []byte(payload)

	record := "account@keyval:coll/record"

	log.Println("Creating collection")

	resp, err := client.CreateColl(record, data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Writing")

	resp, err = client.Create(record, data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("reading")

	resp, err = client.Read(record)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedResp struct {
		Record   string
		Data     []byte
		AbaToken []byte
	}

	log.Println("Resp:\n ", resp)

	log.Println("Resp:\n ", string(resp))

	err = cbor.Unmarshal(resp, &parsedResp)
	if err != nil {
		log.Fatal(err)
	}

	// res := make([]byte, 30)
	//_, err = base64.StdEncoding.Decode(res, parsedResp.Data)

	fmt.Println("Record: ")
	fmt.Println(parsedResp.Record)
	fmt.Println("Data: ")
	fmt.Println(string(parsedResp.Data))
	fmt.Println("Aba token: ")
	fmt.Println(string(parsedResp.AbaToken))

}

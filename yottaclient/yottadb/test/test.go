package main

import (
	"encoding/json"
	"fmt"
	"log"
	"yottaclient/yottadb"
)

func main() {

	client, err := yottadb.New("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	recordPath := "testAccount@testCollection/testRecord"
	data := []byte("helloworld")

	resp, err := client.Write(recordPath, data)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = client.Read(recordPath)
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

	err = client.Delete(recordPath)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Read(recordPath)
	if err == nil {
		log.Fatalln("Record shouldn't exist")
	}

}

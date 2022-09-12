package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"log"
	"yottaclient/yfs"
)

func main() {

	client, err := yfs.New("http://localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	payload := "Hello worlder @#$! \n From me!"

	//data := make([]byte, 30)
	//base64.StdEncoding.Encode(data, []byte(payload))
	data := []byte(payload)

	log.Println("Writing")

	resp, err := client.Write("/test2.txt", data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("reading")

	resp, err = client.Read("/test2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var parsedResp struct {
		Path   string
		Method string
		Data   []byte
	}

	//log.Println("Resp:\n ", resp)

	err = cbor.Unmarshal(resp, &parsedResp)
	if err != nil {
		log.Fatal(err)
	}

	// res := make([]byte, 30)
	//_, err = base64.StdEncoding.Decode(res, parsedResp.Data)

	fmt.Println(parsedResp.Path)
	fmt.Println(parsedResp.Method)
	fmt.Println(string(parsedResp.Data))
	//fmt.Println(parsedResp.Data)

}

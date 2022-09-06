package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"yottaclient/yfs"
)

func main() {

	client, err := yfs.New("http://localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	payload := "@$(!\n#$%^&*"

	data := make([]byte, 30)
	base64.StdEncoding.Encode(data, []byte(payload))

	resp, err := client.Write("/test.txt", data)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = client.Read("/test.txt")
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

	_, err = base64.StdEncoding.Decode(res, parsedResp.Data)

	fmt.Println(parsedResp.Path)
	fmt.Println(parsedResp.Method)
	fmt.Println(string(parsedResp.Data))
	fmt.Println(string(res))

}

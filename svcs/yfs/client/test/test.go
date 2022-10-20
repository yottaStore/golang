package main

import (
	"bytes"
	"github.com/fxamacker/cbor/v2"
	"github.com/yottaStore/golang/svcs/yfs/handlers"
	"io"
	"log"
	"net/http"
)

func main() {

	req := handlers.Request{
		Method: handlers.READ,
		Record: "test",
		Flags:  handlers.FLAG_NONE,
		Data:   nil,
	}

	buff, err := cbor.Marshal(req)
	if err != nil {
		log.Fatal("Error marshaling: ", err)
	}
	url := "http://localhost:8081/yfs"
	resp, err := http.Post(url, "application/cbor", bytes.NewReader(buff))
	if err != nil {
		log.Fatal("Error posting: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error posting: ", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}

	log.Println("Response: ", string(b))

}

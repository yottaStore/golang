package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func produce(msg string) {
	resp, err := http.Post(
		"http://localhost:8090/gossip/produce",
		"application/octet-stream",
		strings.NewReader(msg))

	if err != nil || resp.StatusCode != 200 {
		log.Println("Error calling: ", err)
		log.Println("Resp is: ", resp)
	}
}

func consume() {
	resp, err := http.Get(
		"http://localhost:8090/gossip/consume")

	if err != nil || resp.StatusCode != 200 {
		log.Println("Error calling: ", err)
		log.Println("Resp is: ", resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body: ", err)
	}

	log.Println("Current tree: ", string(body))

}

func main() {

	consume()
	produce("world")
	consume()

}

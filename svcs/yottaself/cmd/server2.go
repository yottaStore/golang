package main

import (
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

func main() {

	port := ":8091"

	var ptr atomic.Value

	ptr.Store([]string{port})

	time.Sleep(1 * time.Second)
	log.Println("Posting")
	resp, err := http.Post(
		"http://localhost:8090/gossip",
		"application/octet-stream",
		strings.NewReader(port))

	if err != nil {
		log.Println("Error posting: ", err)
	}
	log.Println("Resp: ", resp)
}

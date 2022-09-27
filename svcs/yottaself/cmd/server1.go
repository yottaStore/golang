package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func gossipProducer(ptr *atomic.Value) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println("Error reading body: ", err)
			return
		}

		tree := ptr.Load().([]string)
		tree = append(tree, string(body))
		ptr.Store(tree)

	}

}

func gossipConsumer(ptr *atomic.Value) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {

		tree := ptr.Load().([]string)
		_, err := fmt.Fprint(w, "Tree is: ", tree)
		if err != nil {
			log.Println("Error consuming gossip: ", err)
		}

	}

}

func main() {

	tree := []string{"hello"}
	var ptr atomic.Value

	ptr.Store(tree)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			tree := ptr.Load().([]string)
			log.Println("Tree: ", tree)

		}
	}()

	http.HandleFunc("/gossip/produce", gossipProducer(&ptr))
	http.HandleFunc("/gossip/consume", gossipConsumer(&ptr))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

package handlers

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

func GossipHandler(ptr *atomic.Value) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println("Error reading body: ", err)
			return
		}

		log.Println("Body: ", string(body))

		tree := ptr.Load().([]string)
		log.Println("Old tree: ", tree)
		tree = append(tree, string(body))
		log.Println("New tree: ", tree)
		ptr.Store(tree)

	}

}

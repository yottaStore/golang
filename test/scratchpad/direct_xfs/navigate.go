package main

import (
	"log"
	"os"
)

func main() {

	path := "/tmp/yfs/data/test/"

	err := os.RemoveAll(path)
	if err != nil {
		log.Println("Error removing path: ", err)
	}
}

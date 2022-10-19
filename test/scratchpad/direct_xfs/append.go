package main

import (
	"github.com/yottaStore/golang/svcs/yfs/iodriver/unix_xfs"
	"log"
)

func main() {

	iod, err := unix_xfs.New("/tmp/yfs")
	if err != nil {
		log.Fatal("Error creating iodriver: ", err)
	}

	record := "test"

	payload := []byte("Hello, world 3! \n")
	err = iod.Append(record, payload)
	if err != nil {
		log.Fatal("Error appending to file: ", err)
	}
}

package main

import (
	"github.com/yottaStore/golang/svcs/yfs/io_driver/unix_xfs"
	"log"
)

func main() {

	iod, err := unix_xfs.New("/tmp/yfs")
	if err != nil {
		log.Fatal("Error creating io_driver: ", err)
	}

	record := "test"

	payload := []byte("Hello, world 3! \n")
	err = iod.Append(record, payload)
	if err != nil {
		log.Fatal("Error appending to file: ", err)
	}
}

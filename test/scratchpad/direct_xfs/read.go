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

	b, err := iod.Read(record)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	log.Println("Read: ", string(b))

}

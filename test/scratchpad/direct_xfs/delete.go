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

	err = iod.Delete(record)
	if err != nil {
		log.Fatal("Error deleting file: ", err)
	}
}

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

	err = iod.Compact(record)
	if err != nil {
		log.Fatal("Error compacting file: ", err)
	}

}

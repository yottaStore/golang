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

	err = iod.Compact(record)
	if err != nil {
		log.Fatal("Error compacting file: ", err)
	}

}

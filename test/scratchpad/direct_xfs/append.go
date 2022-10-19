package main

import (
	"github.com/yottaStore/golang/svcs/yfs/iodriver/unix_xfs"
	"github.com/yottaStore/golang/utils/block"
	"log"
)

func main() {

	iod, err := unix_xfs.New("/tmp/yfs")
	if err != nil {
		log.Fatal("Error creating iodriver: ", err)
	}

	record := "test"

	payload := block.Alloc(1)
	copy(payload, []byte("Hello, world 3! \n"))
	err = iod.Append(record, payload)
	if err != nil {
		log.Fatal("Error appending to file: ", err)
	}
}

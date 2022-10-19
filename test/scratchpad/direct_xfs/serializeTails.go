package main

import (
	"github.com/yottaStore/golang/svcs/yfs/iodriver/unix_xfs"
	"github.com/yottaStore/golang/utils/block"
	"golang.org/x/sys/unix"
	"log"
)

func main() {

	tailPath := "/tmp/yfs/data/test/tails2"
	td, err := unix.Open(tailPath, unix_xfs.AppendOpts, 0766)
	if err != nil {
		log.Fatal("error opening tails: " + err.Error())
	}

	tail := block.Tail{
		Pointer: []byte("/test"),
		Length:  18,
		Hash:    []byte("hello"),
	}

	b, err := block.SerializeTails([]block.Tail{tail}, 0)

	_, err = unix.Write(td, b)
	if err != nil {
		log.Fatal("error writing tails: " + err.Error())
	}

}

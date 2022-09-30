package main

import (
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func main() {
	type Block struct {
		Data []byte
	}

	path := "/tmp/test/blockAppend"

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	buff := utils.CallocAlignedBlock(2)
	_, err = unix.Read(fd, buff)
	if err != nil {
		log.Println("Error reading file: ", err)
	}

	log.Println("Buffer: ", buff)
	log.Println("Buffer: ", string(buff))

	for i := 0; i < 2; i++ {
		var block Block
		err = cbor.Unmarshal(buff[4096*i:4096*(i+1)], &block)
		if err != nil {
			log.Println("Error decoding: ", err)
			break
		}
		log.Println("Block: ", block)
		log.Println("Content: ", string(block.Data))
	}
}

package main

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func main() {

	path := "/tmp/test/big"

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

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error reading file stats: ", err)
	}

	blockSize := (int(stat.Size)-1)/utils.BlockSize + 1
	res := utils.CallocAlignedBlock(blockSize)
	log.Println("Size: ", len(res))

	n, err := unix.Read(fd, res)
	if err != nil {
		log.Fatal("Error reading result: ", err)
	}
	log.Println("Read result: ", res)
	log.Println("Read bytes: ", n)
	log.Println("Read result: ", string(res))
	log.Println("Size: ", len(res))
	log.Println("Stats: ", stat)
	log.Println("Block size: ", blockSize)
}

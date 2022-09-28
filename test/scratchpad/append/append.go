package main

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func main() {

	path := "/tmp/test/test"

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT|unix.O_TRUNC|unix.O_APPEND, 0766)

	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	buff := "Hello world \n"

	writeSize := (len(buff)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(writeSize)

	copy(file, buff)

	_, err = unix.Write(fd, file)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}

}

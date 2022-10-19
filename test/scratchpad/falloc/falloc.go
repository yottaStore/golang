package main

import (
	"fmt"
	"github.com/yottaStore/golang/utils/block"
	"golang.org/x/sys/unix"
	"log"
)

func main() {

	fileOpts := unix.O_RDWR | unix.O_CREAT | unix.O_DIRECT | unix.O_SYNC | unix.O_APPEND | unix.O_TRUNC

	fd, err := unix.Open("/tmp/yfs/data/test/tails2", fileOpts, 0766)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)

	for i := 0; i < 4; i++ {
		buff := block.Alloc(1)
		copy(buff, fmt.Sprintf("Hello world %d \n", i))
		_, err = unix.Write(fd, buff)
		if err != nil {
			log.Println("Error writing file: ", err)
		}
	}

	err = unix.Fallocate(fd, unix.FALLOC_FL_ZERO_RANGE, 0, 4096*4)
	if err != nil {
		log.Println("Error collapsing: ", err)
	}

}

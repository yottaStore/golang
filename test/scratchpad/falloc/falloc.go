package main

import (
	"fmt"
	"github.com/yottaStore/golang/libs/alloc"
	"golang.org/x/sys/unix"
	"log"
)

func main() {

	fileOpts := unix.O_RDWR | unix.O_CREAT | unix.O_DIRECT | unix.O_SYNC | unix.O_APPEND

	fd, err := unix.Open("/tmp/test", fileOpts, 0766)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)

	for i := 0; i < 2; i++ {
		block := alloc.New(1)
		copy(block, fmt.Sprintf("Hello world %d \n", i))
		_, err = unix.Write(fd, block)
		if err != nil {
			log.Println("Error writing file: ", err)
		}
	}

	err = unix.Fallocate(fd, unix.FALLOC_FL_COLLAPSE_RANGE, 0, 4096)
	if err != nil {
		log.Println("Error punching hole: ", err)
	}

}

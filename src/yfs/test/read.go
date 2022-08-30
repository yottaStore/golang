package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {

	path := "/home/mamluk/yotta/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDWR, 0)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 512)

	n, readErr := unix.Read(fd, file)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Println("Content is: ", string(file))
}

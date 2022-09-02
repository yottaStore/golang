package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/svcs/yottafs/test/utils"
)

func main() {

	BlockSize := 512
	AlignSize := 512

	path := "/home/mamluk/Projects/yottaStore-go/src/yottafs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDWR|unix.O_APPEND|unix.O_NOATIME|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	message := []byte("\nHello Append!\n")

	file := make([]byte, 512*2)

	a := utils.Alignment(file, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	file = file[offset : offset+BlockSize]

	copy(file, message)
	copy(file[len(message):], message)

	fmt.Println(a, offset, len(file))

	n, readErr := unix.Write(fd, file)

	//err = unix.Ftruncate(fd, len(message))

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

}

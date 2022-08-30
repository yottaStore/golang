package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/yfs/test/utils"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	message := []byte("Hello world\n")

	file := make([]byte, 4096*2)

	a := utils.Alignment(file, utils.AlignSize)

	offset := 0
	if a != 0 {
		offset = utils.AlignSize - a
	}

	file = file[offset : offset+utils.BlockSize]

	copy(file, message)

	fmt.Println(file, a, offset, len(file))

	n, readErr := unix.Write(fd, file)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

}

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/svcs/yottafs/test/utils"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/svcs/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 4096*2)

	a := utils.Alignment(file, utils.AlignSize)

	offset := 0
	if a != 0 {
		offset = utils.AlignSize - a
	}

	file = file[offset : offset+utils.BlockSize]
	//file = file[0:4096]

	fmt.Println(a, offset, offset+utils.BlockSize, len(file))

	n, readErr := unix.Pread(fd, file, 0)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Println("Content is: ", string(file))
}

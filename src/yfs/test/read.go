package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

const (
	AlignSize = 4096
	BlockSize = 4096
)

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 4096+4096)

	a := alignment(file, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	file = file[offset : offset+BlockSize]

	fmt.Println(a, offset, len(file))

	n, readErr := unix.Read(fd, file)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Println("Content is: ", string(file))
}

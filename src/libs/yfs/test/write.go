package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

func Alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func main() {

	BlockSize := 512
	AlignSize := 512

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	message := []byte("Hello world\n")

	file := make([]byte, 512*2)

	a := Alignment(file, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	file = file[offset : offset+BlockSize]

	copy(file, message)

	fmt.Println(file, a, offset, len(file))

	n, readErr := unix.Write(fd, file)

	//err = unix.Ftruncate(fd, len(message))

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

}

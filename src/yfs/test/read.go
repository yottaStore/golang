package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
	"yottaStore/yottaStore-go/src/yfs/test/utils"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 4096)

	fmt.Println("Pointer: ", &file[0])
	fmt.Println("Unsafe pointer: ", unsafe.Pointer(&file[0]))
	temp := uintptr(unsafe.Pointer(&file[0]))
	fmt.Println("Unsafe pointer, uintptr: ", temp)
	temp = uintptr(unsafe.Pointer(&file[0])) & uintptr(utils.AlignSize-1)
	fmt.Println("Unpersand: ", temp)
	temp2 := int(temp)
	fmt.Println("Cast to int: ", temp2)

	a := utils.Alignment(file, utils.AlignSize)

	offset := 0
	if a != 0 {
		offset = utils.AlignSize - a
	}

	//file = file[offset : offset+utils.BlockSize]
	file = file[0:4096]

	fmt.Println(a, offset, offset+utils.BlockSize, len(file))

	n, readErr := unix.Pread(fd, file, 0)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Println("Content is: ", string(file))
}

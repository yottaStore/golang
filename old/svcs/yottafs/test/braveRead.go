package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {

	path := "/home/mamluk/yotta/yottaStore-go/svcs/pkgs/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 4096)

	fmt.Println(&file[0])
	//fmt.Println("Unsafe pointer: ", unsafe.Pointer(&file[0]))
	//temp := uintptr(unsafe.Pointer(&file[0]))
	//fmt.Println("Unsafe pointer, uintptr: ", temp)
	//temp = uintptr(unsafe.Pointer(&file[0])) & uintptr(4095)
	//fmt.Println("Unpersand: ", temp)
	//temp2 := int(temp)
	//fmt.Println("Cast to int: ", temp2)

	n, readErr := unix.Pread(fd, file, 0)

	fmt.Println("Return is: ", n)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Println("Content is: ", string(file))
}

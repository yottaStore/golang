package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDWR|unix.O_APPEND|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		panic(err)
	}

	fmt.Println(stat)

}

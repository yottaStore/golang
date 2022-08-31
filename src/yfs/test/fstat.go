package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err == unix.ENOENT {
		fmt.Println("Doesn't exist")
		return
	} else if err != nil {
		panic(err)
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		panic(err)
	}

	fmt.Println(stat)

}

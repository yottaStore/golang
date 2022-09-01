package main

import (
	"fmt"
	"io"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/read"
)

func main() {

	pr, pw := io.Pipe()

	path := "/home/mamluk/yotta/yottaStore-go/src/pkgs/yfs/drivers/direct/test/test.txt"
	go read.Read(path, *pw)
	b := make([]byte, 4096)
	for {
		n, err := pr.Read(b)
		fmt.Println(n, string(b))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Emergency break")
			break
		}

	}

}

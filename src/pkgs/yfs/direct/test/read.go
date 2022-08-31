package main

import (
	"fmt"
	"io"
	"yottaStore/yottaStore-go/src/pkgs/yfs/direct"
)

func main() {

	pr, pw := io.Pipe()

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.atxt"
	go direct.Read(path, *pw)
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

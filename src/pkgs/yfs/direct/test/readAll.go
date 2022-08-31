package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/yfs/direct"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/test/readTest.txt"
	buf, err := direct.ReadAll(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))

}

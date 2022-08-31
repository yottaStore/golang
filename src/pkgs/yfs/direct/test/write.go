package main

import (
	"yottaStore/yottaStore-go/src/pkgs/yfs/direct"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/direct/test/writeTest.txt"
	data := []byte("Hello from yfs!\n")
	direct.Write(path, data)

}

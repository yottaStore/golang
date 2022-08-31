package main

import (
	"yottaStore/yottaStore-go/src/yfs/direct"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/direct/test/writeTest.txt"
	data := []byte("\nHello again from yfs!")
	direct.Append(path, data)

}

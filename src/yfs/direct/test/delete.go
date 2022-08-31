package main

import "yottaStore/yottaStore-go/src/yfs/direct"

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yfs/direct/test/writeTest.txt"
	direct.Delete(path)

}

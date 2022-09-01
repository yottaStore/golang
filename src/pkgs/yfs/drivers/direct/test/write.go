package main

import (
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func main() {

	path := "./test.txt"
	data := []byte("Hello from yfs!\n")
	direct.Write(path, data)

}

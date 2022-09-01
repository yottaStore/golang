package main

import (
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func main() {

	path := "./test.txt"
	data := []byte("\nHello again from yfs!")
	direct.Append(path, data)

}

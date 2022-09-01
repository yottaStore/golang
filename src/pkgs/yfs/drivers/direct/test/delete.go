package main

import (
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func main() {

	path := "/home/mamluk/yotta/yottaStore-go/src/pkgs/yfs/drivers/direct/test/test.txt"
	direct.Delete(path)

}

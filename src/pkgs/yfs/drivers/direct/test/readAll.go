package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func main() {

	path := "/home/mamluk/yotta/yottaStore-go/src/pkgs/yfs/drivers/direct/test/test.txt"
	buf, err := direct.ReadAll(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))

}

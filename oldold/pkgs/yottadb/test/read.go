package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/yottadb/keyvalue"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/svcs/yottadb/test/record.txt"

	item, _ := keyvalue.Read(path)

	fmt.Println(item)
}

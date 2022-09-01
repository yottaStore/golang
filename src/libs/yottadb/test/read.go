package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/libs/yottadb"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yottadb/test/record.txt"

	item, _ := yottadb.Read(path)

	fmt.Println(item)
}

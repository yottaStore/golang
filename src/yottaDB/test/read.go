package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/yottaDB"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yottaDB/test/record.txt"

	item, _ := yottaDB.Read(path)

	fmt.Println(item)
}

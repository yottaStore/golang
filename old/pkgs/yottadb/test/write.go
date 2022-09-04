package main

import (
	"yottaStore/yottaStore-go/src/pkgs/yottadb/keyvalue"
)

func main() {

	type Data struct {
		Hello string
	}

	path := "/home/mamluk/Projects/yottaStore-go/src/yottadb/test/record.txt"
	data := Data{
		Hello: "world",
	}
	keyvalue.Write(path, data)

}

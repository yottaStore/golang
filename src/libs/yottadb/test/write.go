package main

import (
	"yottaStore/yottaStore-go/src/libs/yottadb"
)

func main() {

	type Data struct {
		Hello string
	}

	path := "/home/mamluk/Projects/yottaStore-go/src/yottadb/test/record.txt"
	data := Data{
		Hello: "world",
	}
	yottadb.Write(path, data)

}

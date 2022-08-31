package main

import "yottaStore/yottaStore-go/src/yottaDB"

func main() {

	type Data struct {
		Hello string
	}

	path := "/home/mamluk/Projects/yottaStore-go/src/yottaDB/test/record.txt"
	data := Data{
		Hello: "world",
	}
	yottaDB.Write(path, data)

}

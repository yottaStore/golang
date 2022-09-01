package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/yottadb/keyvalue"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yottadb/test/record.txt"

	type Data struct {
		Hello string
		Count int
	}

	/*data := Data{
		Hello: "world",
		Count: 5,
	}
	yottadb.Write(path, data)*/

	item, _ := keyvalue.ReadOf[Data](path)
	fmt.Println(item)

	updates := make(map[string]interface{})

	updates["Count"] = item.Count + 255
	updates["Test"] = "kawabonga"

	keyvalue.Update(path, updates)
	item, _ = keyvalue.ReadOf[Data](path)
	fmt.Println(item)

}
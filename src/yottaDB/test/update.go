package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/yottaDB"
)

func main() {

	path := "/home/mamluk/Projects/yottaStore-go/src/yottaDB/test/record.txt"

	type Data struct {
		Hello string
		Count int
	}

	/*data := Data{
		Hello: "world",
		Count: 5,
	}
	yottaDB.Write(path, data)*/

	item, _ := yottaDB.ReadOf[Data](path)
	fmt.Println(item)

	updates := make(map[string]interface{})

	updates["Count"] = item.Count + 255
	updates["Test"] = "kawabonga"

	yottaDB.Update(path, updates)
	item, _ = yottaDB.ReadOf[Data](path)
	fmt.Println(item)

}

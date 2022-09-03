package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/rendezvous"
)

func main() {

	recordString := "account@driver:tableName/recordName:recordRow"
	nodes := []string{"hello", "world"}

	node, err := rendezvous.Rendezvous(recordString, nodes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(node)
}

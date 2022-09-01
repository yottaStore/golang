package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/rendezvous"
)

func main() {

	recordString := "account@driver:tableName/recordName:recordRow"

	parsedRecord, err := rendezvous.ParseRecord(recordString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parsedRecord)
}

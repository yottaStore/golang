package main

import (
	"fmt"
	"yottadb/rendezvous"
)

func main() {

	recordStrings := []string{
		"account@driver:tableName/recordName/recordRow/subRow",
		"account@tableName/recordName",
	}

	for _, recordString := range recordStrings {
		parsedRecord, err := rendezvous.ParseRecord(recordString)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(parsedRecord)
	}

}

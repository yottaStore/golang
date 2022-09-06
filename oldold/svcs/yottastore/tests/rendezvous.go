package main

import (
	"fmt"
	"log"
)

func main() {

	recordString := "account@driver:tableName/recordName:recordRow"

	atIndex := -1
	columnCounter := 0
	columnIndexs := [2]int{0, len(recordString)}
	slashIndexs := make([]int, 0)

	for idx, char := range recordString {

		switch char {
		case '@':
			if atIndex != -1 {
				log.Fatal("Too many @s")
			}
			atIndex = idx
			columnIndexs[0] = idx
		case ':':
			if columnCounter > 1 {
				log.Fatal("Too many :s")
			}
			columnIndexs[columnCounter] = idx
			columnCounter++
		case '/':
			slashIndexs = append(slashIndexs, idx)
		}

	}

	account := recordString[:atIndex]
	driver := recordString[atIndex+1 : columnIndexs[0]]
	tableName := recordString[columnIndexs[0]+1 : slashIndexs[0]]
	recordName := recordString[slashIndexs[0]+1 : columnIndexs[1]]
	tableIdentifier := recordString[:slashIndexs[0]]
	recordIdentifier := recordString[:columnIndexs[1]]
	recordRow := recordString[columnIndexs[1]+1 : len(recordString)]

	fmt.Println("Account: ", account)
	fmt.Println("Driver: ", driver)
	fmt.Println("Table Name: ", tableName)
	fmt.Println("Record Name: ", recordName)
	fmt.Println("Table Identifier: ", tableIdentifier)
	fmt.Println("Record Identifier: ", recordIdentifier)
	fmt.Println("Record row: ", recordRow)

}

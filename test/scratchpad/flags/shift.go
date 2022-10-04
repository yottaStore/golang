package main

import (
	"log"
	"strconv"
)

func main() {

	test := int64(0b100)
	log.Println("Test is: ", test)

	test = test<<0b1 | 0b10

	log.Println("Test is: ", test)
	log.Println(strconv.FormatInt(test, 2))
}

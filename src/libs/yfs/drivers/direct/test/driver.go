package main

import (
	"fmt"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct"
)

func main() {

	driver := direct.New()

	payload := []byte("Testing direct driver\n")
	err := driver.Write("./test.txt", payload)
	if err != nil {
		panic(err)
	}

	buff, err := driver.Read("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff), len(buff))

	payload = []byte("Testing direct driver append\n")
	err = driver.Append("./test.txt", payload)
	if err != nil {
		panic(err)
	}

	buff, err = driver.Read("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff), len(buff))

	err = driver.Delete("./test.txt")
	if err != nil {
		panic(err)
	}

	err = driver.Delete("./test.txt")
	if err != nil {
		//fmt.Println(err)
	}

}

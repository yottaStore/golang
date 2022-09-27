package main

import (
	"github.com/fxamacker/cbor/v2"
	"log"
)

type Test struct {
	Word    string
	Counter int
}

func main() {

	test1 := Test{"hello", 2}
	buff, err := cbor.Marshal(test1)
	if err != nil {
		log.Fatal("Error marshalling: ")
	}

	log.Println("Buffer 1: ", buff)

	var test2 []interface{}

	test2 = append(test2, "hello")
	test2 = append(test2, 2)

	log.Println("Original: ", test2)

	buff, err = cbor.Marshal(test2)
	if err != nil {
		log.Fatal("Error marshalling: ", err)
	}

	log.Println("Buffer 2: ", buff)

	log.Println("Buffer 3: ", []byte("hello"))

	var test4 []interface{}
	err = cbor.Unmarshal(buff, &test4)
	if err != nil {
		log.Fatal("Error unmarshalling: ", err)
	}
	log.Println("Unmarshal: ", test4)
}

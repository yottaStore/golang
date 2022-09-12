package main

import (
	"github.com/fxamacker/cbor/v2"
	"log"
)

type TestStruct struct {
	Data string
}

func main() {

	buff, err := cbor.Marshal(TestStruct{Data: ""})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(": ", buff)

	buff, err = cbor.Marshal(TestStruct{Data: "1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("1: ", buff)

	buff, err = cbor.Marshal(TestStruct{Data: "12"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("12: ", buff)

}

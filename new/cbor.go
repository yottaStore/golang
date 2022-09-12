package main

import (
	"bytes"
	"github.com/fxamacker/cbor/v2"
	"log"
)

func main() {

	type Record struct {
		Payload string
		Counter int
	}

	r1 := Record{
		"hello", 1}
	r2 := Record{
		" world", 2}

	var buff []byte

	b, err := cbor.Marshal(r1)
	if err != nil {
		log.Fatal(err)
	}
	buff = append(buff, b...)

	b, err = cbor.Marshal(r2)
	if err != nil {
		log.Fatal(err)
	}
	buff = append(buff, b...)

	log.Println(buff)

	var out1, out2 Record
	decoder := cbor.NewDecoder(bytes.NewReader(buff))

	err = decoder.Decode(&out1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out1)

	err = decoder.Decode(&out2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out2)

}

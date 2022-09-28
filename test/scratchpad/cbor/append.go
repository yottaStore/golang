package main

import (
	"bytes"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
)

type Append struct {
	Word    string
	Counter int
}

func main() {

	var buff []byte

	tmp, err := cbor.Marshal(Append{"hello", 1})
	if err != nil {
		log.Fatal("Error marshaling: ", err)
	}
	buff = append(buff, tmp...)
	tmp, err = cbor.Marshal(Append{"world", 2})
	if err != nil {
		log.Fatal("Error marshaling: ", err)
	}
	buff = append(buff, tmp...)
	buff = append(buff, []byte{0, 0, 0, 0, 0}...)

	decoder := cbor.NewDecoder(bytes.NewReader(buff))

	log.Println("Buffer: ", buff)
	var token Append
	err = decoder.Decode(&token)
	if err != nil {
		log.Fatal("Error unmarshaling: ", err)
	}
	log.Println("Token 1: ", token)

	err = decoder.Decode(&token)
	if err != nil {
		log.Fatal("Error unmarshaling: ", err)
	}
	log.Println("Token 2: ", token)

	err = decoder.Decode(&token)
	if err == io.EOF {
		log.Println("Buffer done")
	}
	if err != nil {
		log.Fatal("Error unmarshaling: ", err)
	}
	log.Println("Token 3: ", token)

}

package main

import (
	"github.com/fxamacker/cbor/v2"
	"log"
)

type Block struct {
	Data []byte
	Hash uint64
}

func logger(msg string, buff []byte, size int) {

	e := float64(1+size) / float64(len(buff))

	log.Println("Case: ", msg)
	log.Println("Size: ", len(buff))
	log.Println("Efficiency: ", e)
	//log.Println("Buffer: ", buff)

}

func main() {

	size := 3997
	payload := make([]byte, size)
	for i := 0; i < size; i++ {
		payload[i] = uint8(i)
	}

	first := Block{payload, uint64(0)}
	buff, err := cbor.Marshal(first)
	if err != nil {
		log.Fatal("Error marshalling: ", err)
	}
	logger("first", buff, size)

	buff, err = cbor.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling: ", err)
	}
	logger("second", buff, size)

}

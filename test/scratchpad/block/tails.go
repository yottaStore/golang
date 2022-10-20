package main

import (
	"github.com/yottaStore/golang/utils/block"
	"log"
)

func main() {

	t := block.Tail{Pointer: []byte("hello"), Hash: []byte("world")}

	buff, err := block.SerializeTails([]block.Tail{t}, block.F_COMPRESSED)
	if err != nil {
		log.Fatal("Error serializing tail: ", err)
	}

	log.Println("Buffer: ", buff)
	log.Println("Buffer:", string(buff[8:25]))

}

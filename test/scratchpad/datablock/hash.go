package main

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"log"
)

func main() {

	payload := "hello"

	hash := xxh3.Hash([]byte(payload))
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, hash)
	h1 := binary.LittleEndian.Uint64(b)

	opts := uint32(0xFFFFFFFF)
	b2 := make([]byte, 4)
	binary.LittleEndian.PutUint32(b2, opts)

	log.Println("Hash: ", hash)
	log.Println("Buffer: ", b)
	log.Println("Buffer 2: ", b2)

	log.Println("Hash: ", h1)

}

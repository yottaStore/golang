package main

import (
	"encoding/binary"
	"github.com/fxamacker/cbor/v2"
	"github.com/zeebo/xxh3"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func main() {

	b := utils.CallocAlignedBlock(1)

	header := uint32(0xFF)
	payload := "hello world"
	pb, err := cbor.Marshal(payload)
	if err != nil {
		log.Fatal("error marshaling: ", err)
	}

	binary.BigEndian.PutUint32(b, header)
	copy(b[4:], pb)

	hash := xxh3.Hash(b[:4088])
	binary.BigEndian.PutUint64(b[4088:], hash)

	log.Println("Header: ", b)
	log.Println("Payload: ", pb)
	log.Println("Buffer: ", b)
}

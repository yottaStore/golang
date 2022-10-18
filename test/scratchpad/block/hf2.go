package main

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"log"
)

type Block2 struct {
	Version   uint8
	Type      uint8
	Flags     uint16
	Length    uint16
	Reserved1 uint8
	Reserved2 uint8
	// Total 64 bits
	Body string
	Hash uint64
} // Total 128 bits + body

func serialize2(payload []byte, t uint8, flags uint16) []byte {

	buff := make([]byte, 128, 128)

	pll := len(payload)

	log.Println("Payload length: ", pll)

	buff[0] = 1
	buff[1] = t
	buff[2] = byte(flags >> 8)
	buff[3] = byte(flags)
	buff[4] = byte(pll >> 8)
	buff[5] = byte(pll)

	copy(buff[8:], payload)

	h := xxh3.Hash(buff[:])
	binary.BigEndian.PutUint64(buff[120:], h)

	return buff

}

func deserialize2(buff []byte) Block2 {

	var block Block2
	block.Version = buff[0]
	block.Type = buff[1]
	block.Flags = uint16(buff[2])<<8 | uint16(buff[3])
	block.Length = uint16(buff[4])<<8 | uint16(buff[5])
	block.Reserved1 = buff[6]
	block.Reserved2 = buff[7]
	block.Body = string(buff[8 : 8+block.Length])
	block.Hash = binary.BigEndian.Uint64(buff[120:])

	return block

}

func main() {

	b := serialize2([]byte("Hello World!"), 13, 32767)

	block := deserialize2(b)

	log.Println("Buffer: ", b)
	log.Println("Block: ", block)

}

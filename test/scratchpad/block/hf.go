package main

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"log"
)

type Block struct {
	Version uint8
	Type    uint8
	Flags   uint16
	Length  uint16
	Body    string
	Hash    uint64
}

func serialize(payload []byte) []byte {
	buff := make([]byte, 128, 128)

	var version uint8 = 15
	var blockType uint8 = 13

	var length uint16 = 511
	var flags uint16 = 32767

	buff[0] = blockType<<4 | version
	buff[1] = byte(flags)
	buff[2] = byte(flags>>7) | byte(length&0x100>>8)
	buff[3] = byte(length)

	log.Printf("Flags: %b", buff[1])
	log.Printf("Middle: %b", buff[2])
	log.Printf("Length: %b", buff[3])

	copy(buff[4:], payload)
	h := xxh3.Hash(buff[:])
	log.Println("Hash: ", h)
	binary.BigEndian.PutUint64(buff[120:], h)

	return buff
}

func deserialize(buff []byte) Block {
	var block Block
	block.Version = buff[0] & 0xF
	block.Type = buff[0] >> 4
	block.Flags = uint16(buff[1]) | uint16(buff[2]&0xFE)<<7

	block.Length = uint16(buff[2])&0x1<<8 | uint16(buff[3])
	//block.Body = string(buff[4 : 4+block.Length])
	block.Body = string(buff[4:120])
	block.Hash = binary.BigEndian.Uint64(buff[120:])

	return block
}

func main() {

	payload := []byte("Hello World!")

	b := serialize(payload)
	block := deserialize(b)

	log.Println("Buff: ", b)
	log.Println("Block: ", block)

}

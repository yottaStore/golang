package main

import (
	"log"
)

type Head struct {
	Version    uint8
	Format     uint8
	Size       uint16
	Flags      uint16
	BlockCount uint16
	Reserved   uint16
}

func Serialize(h Head) ([]uint8, error) {

	buff := make([]byte, 8)

	buff[0] = h.Version<<5 |
		h.Format<<1 |
		uint8(h.Size>>8)

	buff[1] = uint8(h.Size & 0xFF)

	buff[2] = uint8(h.Flags >> 8)
	buff[3] = uint8(h.Flags)

	buff[4] = uint8(h.BlockCount >> 8)
	buff[5] = uint8(h.BlockCount)

	buff[6] = uint8(h.Reserved >> 8)
	buff[7] = uint8(h.Reserved)

	return buff, nil
}

func Deserialize(buff []uint8) (Head, error) {

	var h Head

	h.Version = buff[0] >> 5
	h.Format = buff[0] >> 1 & 0xF
	h.Size = uint16(buff[0]&1)<<8 | uint16(buff[1])
	h.Flags = uint16(buff[2])<<8 | uint16(buff[3])
	h.BlockCount = uint16(buff[4])<<8 | uint16(buff[5])
	h.Reserved = uint16(buff[6])<<8 | uint16(buff[7])

	return h, nil
}

func main() {

	h := Head{
		Version:    7,
		Format:     15,
		Size:       511,
		Flags:      4,
		BlockCount: 65535,
	}

	buff, err := Serialize(h)
	if err != nil {
		log.Fatal("Error serializing: ", err)
	}

	log.Println("Buffer: ", buff)
	log.Printf("Version Buffer: %b", buff[0])

	h, err = Deserialize(buff)
	if err != nil {
		log.Fatal("Error deserializing: ", err)
	}

	log.Println("Head: ", h)

}

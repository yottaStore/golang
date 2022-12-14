package block

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
)

func Serialize(payload []byte, t Type, f Flags) ([]byte, error) {

	// Get the size of the block
	size, remainder := GetSize(len(payload))

	// Create the block
	buff := Alloc(size)

	// Set the body
	for i := 0; i < size; i++ {

		index := i * BlockSize

		// Set the head
		buff[index] = v0
		buff[index+1] = uint8(t)
		buff[index+2] = byte(f >> 8)
		buff[index+3] = byte(f)

		if i == size-1 {
			buff[index+4] = byte(remainder >> 8)
			buff[index+5] = byte(remainder)
			copy(buff[i*BlockSize+HeadSize:], payload[i*BodySize:])
		} else {
			buff[index+4] = byte(BodySize >> 8)
			buff[index+5] = byte(BodySize & 0xFF)
			//buff[index+5] = byte(BodySize)
			copy(buff[i*BlockSize+HeadSize:], payload[i*BodySize:(i+1)*BodySize])
		}

		// Set the hash
		// TODO: verify the hash
		// TODO: hash is bugged
		h := xxh3.Hash(buff[i*BlockSize : (i+1)*BlockSize+HeadSize])
		binary.BigEndian.PutUint64(buff[(i+1)*BlockSize-FootSize:], h)
	}

	return buff, nil

}

func Deserialize(buff []byte) ([]Block, error) {

	count := len(buff) / BlockSize

	blocks := make([]Block, count)

	for i := 0; i < count; i++ {
		index := i * BlockSize
		length := binary.BigEndian.Uint16(buff[index+4:])
		blocks[i] = Block{
			Version: buff[index],
			Type:    Type(buff[index+1]),
			Flags:   Flags(binary.BigEndian.Uint16(buff[index+2:])),
			Length:  length,
			Body:    buff[index+HeadSize : index+HeadSize+int(length)],
			Hash:    binary.BigEndian.Uint64(buff[index+BlockSize-FootSize:]),
		}
	}

	return blocks, nil
}

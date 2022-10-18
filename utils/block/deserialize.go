package block

import "encoding/binary"

func Deserialize(buff []byte) ([]Block, error) {

	count := len(buff) / BlockSize

	blocks := make([]Block, count)

	for i := 0; i < count; i++ {
		index := i * BlockSize
		length := binary.BigEndian.Uint16(buff[index+4:])
		blocks[i] = Block{
			Version: buff[index],
			Type:    BlockType(buff[index+1]),
			Flags:   Flag(binary.BigEndian.Uint16(buff[index+2:])),
			Length:  length,
			Body:    buff[index+HeadSize : index+HeadSize+int(length)],
			Hash:    binary.BigEndian.Uint64(buff[index+BlockSize-FootSize:]),
		}
	}

	return blocks, nil
}

package yottablock

import "errors"

type Head struct {
	Version  uint8
	Format   uint8
	Length   uint16
	Flags    uint16
	Size     uint16
	Reserved uint16
}

func Serialize(h Head) ([]uint8, error) {

	if h.Version > 7 {

		return nil, errors.New("version is too big")
	}

	if h.Format > 15 {

		return nil, errors.New("format is too big")
	}

	if h.Length > 511 {

		return nil, errors.New("size is too big")
	}

	buff := make([]byte, 8)

	buff[0] = h.Version<<5 |
		h.Format<<1 |
		uint8(h.Length>>8)

	buff[1] = uint8(h.Length & 0xFF)

	buff[2] = uint8(h.Flags >> 8)
	buff[3] = uint8(h.Flags)

	buff[4] = uint8(h.Size >> 8)
	buff[5] = uint8(h.Size)

	buff[6] = uint8(h.Reserved >> 8)
	buff[7] = uint8(h.Reserved)

	return buff, nil
}

func Deserialize(buff []uint8) (Head, error) {

	var h Head

	h.Version = buff[0] >> 5
	h.Format = buff[0] >> 1 & 0xF
	h.Length = uint16(buff[0]&1)<<8 | uint16(buff[1])
	h.Flags = uint16(buff[2])<<8 | uint16(buff[3])
	h.Size = uint16(buff[4])<<8 | uint16(buff[5])
	h.Reserved = uint16(buff[6])<<8 | uint16(buff[7])

	return h, nil
}

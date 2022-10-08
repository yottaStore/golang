package yottablock

const (
	headSize = 4
)

func getHashes(buff []byte) ([]byte, int) {

	hashCount := (len(buff)-1)/(4096*4) + 1

	hb := make([]byte, hashCount*8)

	for i := 0; i < hashCount; i++ {

	}

	return hb, hashCount
}

func getPayloadSize(payloadLen int) (uint16, uint16) {

	payloadLen = payloadLen + headSize
	length := payloadLen % 4096
	hashLen := (payloadLen-1)/2048 + 8
	payloadLen = payloadLen + hashLen
	size := (payloadLen-1)/4096 + 1

	return uint16(size), uint16(length)
}

func SerializePayload(payload []byte, opts uint16) ([]byte, error) {

	size, length := getPayloadSize(len(payload))

	h := Head{
		Version:  0,
		Format:   1,
		Length:   length,
		Flags:    0b10100,
		Size:     size,
		Reserved: 0,
	}

	hb, err := Serialize(h)
	if err != nil {
		return nil, err
	}

	buff := CallocAlignedBlock(int(size))
	copy(buff, hb)
	copy(buff[4:], payload)

	hb, hl := getHashes(buff)
	copy(buff[len(buff)-hl:], hb)

	return buff, nil
}

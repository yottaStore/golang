package rendezvous

import (
	"fmt"
	"github.com/zeebo/xxh3"
)

const (
	hashKey = "949949494djjdjdjd^"
)

func Uint16(u xxh3.Uint128) [8]uint16 {
	return [8]uint16{
		uint16(u.Hi >> 0x30), uint16(u.Hi >> 0x20),
		uint16(u.Hi >> 0x10), uint16(u.Hi),
		uint16(u.Lo >> 0x30), uint16(u.Lo >> 0x20),
		uint16(u.Lo >> 0x10), uint16(u.Lo),
	}
}

func Rendezvous(record ParsedRecord, nodes []string) (string, error) {

	tmp := xxh3.HashString128(record.RecordIdentifier + hashKey)
	tmpArray := Uint16(tmp)

	fmt.Println(tmp, tmpArray)

	max := uint16(0)
	maxIndex := 0

	for idx, value := range tmpArray {
		if value > max {
			maxIndex = idx
			max = value
		}
	}

	nodeIndex := maxIndex % len(nodes)

	return nodes[nodeIndex], nil
}

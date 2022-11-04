package rebar

import "github.com/zeebo/xxh3"

func Hash(record string, seed uint64) [16]byte {

	h := xxh3.Hash128Seed([]byte(record), seed)

	return h.Bytes()
}

func HashNU16(record string, count int, verticalSeed, horizontalSeed uint64) []uint16 {

	rounds := (count + 1) / 8

	buff := make([]uint16, 0, count)

	for i := 0; i < rounds; i++ {
		hash := xxh3.Hash128Seed([]byte(record), verticalSeed).Bytes()
		buff = append(buff, bytesToUint16(hash[:])...)
	}

	return nil

}

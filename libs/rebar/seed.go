package rebar

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
)

func SeedTable(seed uint64) [16][4]uint64 {

	var table [16][4]uint64

	h1 := xxh3.Hash128(binary.BigEndian.AppendUint64(nil, seed)).Bytes()
	h2 := xxh3.Hash128(h1[:]).Bytes()
	table[0][0] = binary.BigEndian.Uint64(h1[:8])
	table[0][1] = binary.BigEndian.Uint64(h1[8:])
	table[0][2] = binary.BigEndian.Uint64(h2[:8])
	table[0][3] = binary.BigEndian.Uint64(h2[8:])

	for i := 1; i < 16; i++ {
		buff := binary.BigEndian.AppendUint64(nil, table[i-1][0])
		buff = binary.BigEndian.AppendUint64(buff, table[i-1][1])

		h1 := xxh3.Hash128Seed(buff, table[i-1][2]).Bytes()
		h2 := xxh3.Hash128Seed(buff, table[i-1][3]).Bytes()

		table[i][0] = binary.BigEndian.Uint64(h1[:8])
		table[i][1] = binary.BigEndian.Uint64(h1[8:])
		table[i][2] = binary.BigEndian.Uint64(h2[:8])
		table[i][3] = binary.BigEndian.Uint64(h2[8:])
	}

	return table
}

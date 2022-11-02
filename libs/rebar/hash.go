package rebar

import "github.com/zeebo/xxh3"

func Hash(record string, seed uint64) [16]byte {

	h := xxh3.Hash128Seed([]byte(record), seed)

	return h.Bytes()
}

package hasher

import "github.com/zeebo/xxh3"

type Xxh3Hasher struct {
	Seed      [2]uint64
	SeedTable [][]uint64
}

func NewXxh3Hasher(seeds [2]uint64) (Interface, error) {

	table := SeedTable(seeds[0], seeds[1])

	return &Xxh3Hasher{
		Seed:      seeds,
		SeedTable: table,
	}, nil
}

func (h *Xxh3Hasher) Init(pointer string, seed uint64) error {

	return nil
}

func (h *Xxh3Hasher) Read(pointer string, level, length int) ([]byte, error) {

	var buff []byte

	for i := 0; i < length; i++ {
		seed := h.SeedTable[level][i]
		tmp := xxh3.Hash128Seed([]byte(pointer), seed).Bytes()
		buff = append(buff, tmp[:]...)
	}

	return buff, nil
}

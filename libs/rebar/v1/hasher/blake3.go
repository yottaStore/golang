package hasher

type Blake3Hasher struct {
	Seed  [2]uint64
	Cache [][]byte
}

func NewBlake3Hasher(seeds [2]uint64) Interface {

	return &Blake3Hasher{
		Seed: seeds,
	}
}

func (h *Blake3Hasher) Init(pointer string, seed uint64) error {

	return nil
}

func (h *Blake3Hasher) Read(pointer string, level, length int) ([]byte, error) {

	return nil, nil
}

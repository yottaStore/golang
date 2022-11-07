package v1

func New(seed uint64) (*Navigator, error) {

	var seedTable [16][8]uint64

	n := &Navigator{
		Seed:      seed,
		SeedTable: &seedTable,
	}

	return n, nil

}

func NewWithOpts(seed uint64, hasher Hasher, scorer Scorer) (*Navigator, error) {

	var seedTable [16][8]uint64

	n := &Navigator{
		Seed:      seed,
		Hasher:    hasher,
		Scorer:    scorer,
		SeedTable: &seedTable,
	}

	return n, nil

}
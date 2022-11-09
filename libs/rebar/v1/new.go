package v1

import (
	"github.com/yottaStore/golang/libs/rebar/v1/hasher"
	"github.com/yottaStore/golang/libs/rebar/v1/scorer"
)

func New(seed [2]uint64) (*Navigator, error) {

	h, err := hasher.NewXxh3Hasher(seed)
	if err != nil {
		return nil, err
	}

	s, err := scorer.NewLogScore()
	if err != nil {
		return nil, err
	}

	n := &Navigator{
		Seed:   seed,
		Hasher: h,
		Scorer: s,
	}

	return n, nil

}

func NewWithOpts(seed [2]uint64, hasher hasher.Interface, scorer scorer.Interface) (*Navigator, error) {

	n := &Navigator{
		Seed:   seed,
		Hasher: hasher,
		Scorer: scorer,
	}

	return n, nil

}

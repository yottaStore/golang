package horizontal

import "sort"

type Sortable interface {
	uint8 | uint16 | uint32
}

type Slice struct {
	ints []uint32
	idx  []int
}

func (s Slice) Len() int {
	return len(s.ints)
}
func (s Slice) Swap(i, j int) {
	s.ints[i], s.ints[j] = s.ints[j], s.ints[i]
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s Slice) Less(i, j int) bool {
	return s.ints[i] > s.ints[j]
}

func Round[T Sortable](hashes []T, weights []uint32) []int {

	buff := make([]uint32, len(hashes))
	idxs := make([]int, len(hashes))

	is_weights := weights != nil

	for i, hash := range hashes {
		weight := uint32(1)
		if is_weights {
			weight = weights[i]
		}
		buff[i] = uint32(hash) * weight
		idxs[i] = i
	}

	s := Slice{ints: buff, idx: idxs}

	sort.Sort(s)

	return s.idx

}

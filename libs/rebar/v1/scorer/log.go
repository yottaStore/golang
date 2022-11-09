package scorer

import (
	"encoding/binary"
	"errors"
	"github.com/yottaStore/golang/utils/htree"
	"math"
	"sort"
)

type LogScores struct {
	scores []float64
	idx    []int
}

func (s LogScores) Len() int {
	return len(s.scores)
}
func (s LogScores) Swap(i, j int) {
	s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s LogScores) Less(i, j int) bool {
	return s.scores[i] > s.scores[j]
}

func bytesToUint16(b []byte) []uint16 {
	buff := make([]uint16, len(b)/2)
	for i := 0; i < 8; i++ {

		buff[i] = binary.BigEndian.Uint16(b[2*i : 2*(i+1)])
	}

	return buff
}

func log_round(nodes []*htree.Node, hash []byte, shards int) ([]*htree.Node, error) {

	count := len(nodes)

	if count == 1 {
		return nodes[0].Children, nil
	}

	if count > len(hash)/2 {
		return nil, errors.New("not enough hash bytes to round")
	}

	buff := bytesToUint16(hash)

	scores := make([]float64, count)
	idxs := make([]int, count)

	for i := 0; i < count; i++ {
		tmp := (float64(buff[i]) + 1) / float64(0xffff)
		tmp = 1.0 / (-math.Log(tmp))
		scores[i] = tmp * float64(nodes[i].Weight)
		idxs[i] = i
	}

	s := LogScores{scores: scores, idx: idxs}
	sort.Sort(s)

	output := make([]*htree.Node, shards)

	for i := 0; i < shards; i++ {
		output[i] = nodes[s.idx[i]]
	}

	return output, nil
}

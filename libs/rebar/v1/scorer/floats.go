package scorer

type FloatScores struct {
	scores []float64
	idx    []int
}

func (s FloatScores) Len() int {
	return len(s.scores)
}
func (s FloatScores) Swap(i, j int) {
	s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s FloatScores) Less(i, j int) bool {
	return s.scores[i] > s.scores[j]
}

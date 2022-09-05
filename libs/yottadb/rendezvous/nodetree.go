package rendezvous

type HashTree struct {
	hashes []uint16
	idxs   []int
}

func (n HashTree) Len() int {
	return len(n.hashes)
}

func (n HashTree) Less(i, j int) bool {
	return n.hashes[i] > n.hashes[j]
}

func (n HashTree) Swap(i, j int) {
	n.hashes[i], n.hashes[j] = n.hashes[j], n.hashes[i]
	n.idxs[i], n.idxs[j] = n.idxs[j], n.idxs[i]
}

func NewNodeTree(hashes []uint16) HashTree {

	nodesCount := len(hashes)
	indexes := make([]int, nodesCount)
	for i := 0; i < nodesCount; i++ {
		indexes[i] = i
	}

	return HashTree{
		hashes,
		indexes,
	}
}

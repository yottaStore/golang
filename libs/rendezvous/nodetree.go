package rendezvous

import (
	"github.com/zeebo/xxh3"
	"sort"
)

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

func findNodes(token string, nodes Nodemap, opts Options, hashkey string) (Nodemap, error) {

	hashedRecord := xxh3.HashString128(token + hashkey)
	tmpArray := NewNodeTree(toUint16(hashedRecord))
	sort.Sort(tmpArray)

	idxs := tmpArray.idxs[:opts.Sharding]
	pickedNodes := make(Nodemap, opts.Sharding)

	for i := 0; i < opts.Replication; i++ {
		index := idxs[i] % len(nodes)
		pickedNodes[i] = nodes[index]
		pickedNodes[i] = nodes[index]
	}

	return pickedNodes, nil
}

func toUint16(u xxh3.Uint128) []uint16 {
	return []uint16{
		uint16(u.Hi >> 0x30), uint16(u.Hi >> 0x20),
		uint16(u.Hi >> 0x10), uint16(u.Hi),
		uint16(u.Lo >> 0x30), uint16(u.Lo >> 0x20),
		uint16(u.Lo >> 0x10), uint16(u.Lo),
	}
}

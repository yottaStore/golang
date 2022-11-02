package htree

import "github.com/zeebo/xxh3"

type Tree struct {
	Root *Node
}

type Node struct {
	Pointer  string
	Weight   uint32
	Children []*Node
	Hash     [16]byte
}

type Shard struct {
	Pointer string
	Coords  []string
	Weight  uint32
}

func (t *Tree) Insert(s Shard) error {

	lastNode := t.Root

	hash := xxh3.Hash128([]byte(s.Pointer)).Bytes()

	newLeaf := Node{
		Pointer: s.Pointer,
		Weight:  s.Weight,
		Hash:    hash,
	}

	visitedNodes := make([]*Node, len(s.Coords)-1)

	for idx, coord := range s.Coords {

		var found bool
		for _, node := range lastNode.Children {
			if node.Pointer == coord {
				found = true
				node.Weight += s.Weight
				visitedNodes[idx] = node
				lastNode = node
				break
			}
		}

		if !found {

			lastNode.Children = append(lastNode.Children, &newLeaf)
			lastNode = &newLeaf
		}

	}
	return nil
}

func (t *Tree) Delete() error {
	return nil
}

func (t *Tree) Compare(hash [16]byte) error {
	return nil
}

func (t *Tree) Verify() error {
	return nil
}

func (t *Tree) Print() {
}

package hTrie

import (
	"encoding/binary"
	"errors"
	"github.com/zeebo/xxh3"
	"log"
)

type Node struct {
	Pointer  string
	Weight   uint32
	Children []*Node
	Parent   *Node
	Hash     [16]byte
}

func (n *Node) Insert(s Shard, prefix string) error {

	var parent *Node
	children := []*Node{n}

	coords := s.GetCoords(prefix)
	levels := len(coords)
	visitedNodes := make([]*Node, levels)

	for idx, coord := range coords {

		found := false
		for cidx, child := range children {

			if child.Pointer > coord {
				found = true
				newNode := Node{
					Pointer: coord,
					Weight:  s.Weight,
					Parent:  parent,
				}

				// Make sure insertion keep lexicographical order
				parent.Children = append(children[:cidx+1], children[cidx:]...)
				parent.Children[cidx] = &newNode
				visitedNodes[idx] = &newNode
				parent = &newNode
				children = nil
				break
			}

			if child.Pointer == coord {
				found = true
				child.Weight += s.Weight
				visitedNodes[idx] = child
				parent = child
				children = child.Children
				break
			}
		}

		if !found {
			newNode := Node{
				Pointer: coord,
				Weight:  s.Weight,
				Parent:  parent,
			}

			parent.Children = append(children, &newNode)
			visitedNodes[idx] = &newNode
			parent = &newNode
			children = nil
		}

	}

	for i := levels - 1; i >= 0; i-- {

		err := visitedNodes[i].UpdateHash()
		if err != nil {
			return err
		}
	}

	return nil

}

func (n *Node) Update(s Shard, prefix string) error {

	return nil
}

func (n *Node) Delete(s Shard, prefix string) error {

	return nil
}

func (n *Node) UpdateHash() error {

	if len(n.Children) == 0 {
		n.Hash = xxh3.Hash128([]byte(n.Pointer)).Bytes()
		return nil
	}

	buff := make([]byte, 0, len(n.Children)*16+4)
	for _, child := range n.Children {
		buff = append(buff, child.Hash[:]...)
	}

	buff = binary.BigEndian.AppendUint32(buff, n.Weight)

	hash := xxh3.Hash128(buff).Bytes()
	n.Hash = hash

	return nil

}

func (n *Node) Verify() error {

	if len(n.Children) == 0 {
		hash := xxh3.Hash128([]byte(n.Pointer)).Bytes()
		for i, b := range hash {
			if b != n.Hash[i] {
				return errors.New("hash mismatch")
			}
		}
		return nil
	}

	buff := make([]byte, 0, len(n.Children)*16+4)
	for _, child := range n.Children {
		buff = append(buff, child.Hash[:]...)
	}

	buff = binary.BigEndian.AppendUint32(buff, n.Weight)

	hash := xxh3.Hash128(buff).Bytes()

	for i, b := range hash {
		if b != n.Hash[i] {
			return errors.New("hash mismatch")
		}
	}

	for _, child := range n.Children {
		if err := child.Verify(); err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) Print() {

	log.Println("Node: ", n)
	for _, child := range n.Children {
		child.Print()
	}

}

package htree

import (
	"errors"
	"log"
)

type NodeType uint8

const (
	T_SHARD NodeType = 1
	T_PNODE NodeType = 2
	T_VNODE NodeType = 3
	T_ROOT  NodeType = 4
)

const (
	HASH_SIZE = 16
)

type Node struct {
	Pointer  string
	Type     NodeType
	Weight   uint64
	Load     uint32
	Children []*Node
	Parent   *Node
	Hash     [HASH_SIZE]byte
}

func (n *Node) Update() error {
	hash, err := computeHash(n)
	if err != nil {
		return err
	}
	n.Hash = hash
	return nil
}

func (n *Node) Verify(children, integrity bool) error {

	hash, err := computeHash(n)
	if err != nil {
		return errors.New("error computing hash: " + err.Error())
	}

	if hash != n.Hash {
		return errors.New("hash mismatch")
	}

	if integrity {
		if err = n.Integrity(); err != nil {
			return errors.New("integrity error: " + err.Error())
		}
	}

	if children {
		for _, child := range n.Children {
			if err = child.Verify(children, integrity); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *Node) Integrity() error {
	// TODO: check node respect the structure

	return nil
}

func (n *Node) Print() {
	log.Println("Node: ", n)
	for _, child := range n.Children {
		child.Print()
	}
}

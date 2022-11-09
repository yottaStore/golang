package v1

import (
	"errors"
	"github.com/yottaStore/golang/libs/rebar/v1/hasher"
	"github.com/yottaStore/golang/libs/rebar/v1/scorer"
	"github.com/yottaStore/golang/utils/htree"
)

var ErrRoundAtLeaf = errors.New("round at leaf")

func round(pointer string, children []*htree.Node, level, count int,
	h hasher.Interface, s scorer.Interface) (
	[]*htree.Node, int, error) {

	// Repeat until you have enough children
	for {
		level++

		if len(children) == 0 {
			return nil, 0, ErrRoundAtLeaf
		}

		if len(children) == count {
			return children, level, nil
		}

		if len(children) < count {
			children = gatherChildren(children)
			continue
		}

		if len(children) > count {
			break
		}
	}

	hash, err := h.Read(pointer, level, len(children))
	if err != nil {
		return nil, 0, err
	}
	nodes, err := s.Compute(children, hash, count)

	return nodes, level, err

}

func gatherChildren(nodes []*htree.Node) []*htree.Node {
	var children []*htree.Node
	for _, node := range nodes {
		children = append(children, node.Children...)
	}
	return children
}

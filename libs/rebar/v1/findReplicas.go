package v1

import (
	"github.com/yottaStore/golang/libs/rebar/v1/hasher"
	"github.com/yottaStore/golang/libs/rebar/v1/scorer"
	"github.com/yottaStore/golang/utils/htree"
)

func findReplicas(pointer string, root *htree.Node, opts Options,
	h hasher.Interface, s scorer.Interface) (
	[]*htree.Node, int, error) {

	// Find replicas
	var replicas []*htree.Node
	//var hash []byte
	children := root.Children
	var level int
	// Repeat until you have enough children
	for level = 0; ; level++ {

		hash, err := h.Read(pointer, level, len(children))
		if err != nil {
			return nil, 0, err
		}
		replicas, err = s.Compute(children, hash, opts.Replication)

		// Repeat if error == ErrNotEnoughChildren
		if err == nil {
			break
		}
	}

	return replicas, level, nil
}

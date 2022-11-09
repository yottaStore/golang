package v1

import (
	"errors"
	"github.com/yottaStore/golang/libs/rebar/v1/hasher"
	"github.com/yottaStore/golang/libs/rebar/v1/scorer"
	"github.com/yottaStore/golang/utils/htree"
)

func findPools(pointer string, root *htree.Node, opts Options,
	h hasher.Interface, s scorer.Interface) (
	[][]*htree.Node, int, error) {

	// Find replicas
	replicas, level, err := round(pointer, root.Children, 0, opts.Replication, h, s)
	if err != nil {
		return nil, 0, err
	}

	// For each replica, find pool
	pools := make([][]*htree.Node, opts.Replication)
	for idx, replica := range replicas {

		poolRoot, rl, err := round(pointer, replica.Children, level, 2, h, s)
		if err != nil {
			return nil, rl, err
		}

		lowCount := opts.Sharding / 2
		highCount := lowCount + opts.Sharding%2
		highPool, hl, err := findPool(pointer, poolRoot[0], highCount, rl, h, s)
		if err != nil {
			return nil, hl, err
		}
		lowPool, ll, err := findPool(pointer, poolRoot[1], lowCount, rl, h, s)
		if err != nil {
			return nil, ll, err
		}

		level = hl
		if ll > level {
			level = ll
		}

		pools[idx] = append(highPool, lowPool...)
	}

	return pools, level, nil

}

func findPool(pointer string, root *htree.Node, count, level int,
	h hasher.Interface, s scorer.Interface) ([]*htree.Node, int, error) {

	if len(root.Children) == 0 {
		if count == 1 {
			return root.Children, level, nil
		} else {
			return nil, level, errors.New("not enough nodes")
		}
	}

	children := root.Children

	for {
		if len(children[0].Children) == 0 {

			return round(pointer, children, level, count, h, s)
		}
		nextChildren, nextLevel, err := round(pointer, children, level, 1, h, s)
		level = nextLevel
		children = nextChildren
		if err != nil {
			return nil, nextLevel, err
		}

	}
}

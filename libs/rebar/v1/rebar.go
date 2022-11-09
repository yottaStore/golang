package v1

import (
	"github.com/yottaStore/golang/libs/rebar/v1/hasher"
	"github.com/yottaStore/golang/libs/rebar/v1/scorer"
	"github.com/yottaStore/golang/utils/htree"
	"github.com/yottaStore/golang/utils/record"
)

type Scorer func(nodes []*htree.Node, hash []byte, shards int) ([]*htree.Node, error)

type Navigator struct {
	Seed   [2]uint64
	Hasher hasher.Interface
	Scorer scorer.Interface
}

type Options struct {
	Sharding    int
	Replication int
}

func (n *Navigator) FindPools(pointer string, tree *htree.Root, opts Options) ([][]*htree.Node, int, error) {

	// Init hasher
	err := n.Hasher.Init(pointer, n.Seed[0])
	if err != nil {
		return nil, 0, err
	}

	return findPools(pointer, tree.Root, opts, n.Hasher, n.Scorer)
}

func (n *Navigator) FindRecord(pointer string, pool []*htree.Node, level int) (*htree.Node, error) {

	return findRecord(pointer, pool)

}

func (n *Navigator) Find(record record.Record, tree *htree.Root, opts Options) ([]*htree.Node, error) {

	pools, level, err := n.FindPools(record.PoolPointer, tree, opts)
	if err != nil {
		return nil, err
	}

	replicas := make([]*htree.Node, opts.Replication)

	for idx, pool := range pools {
		replicas[idx], err = n.FindRecord(record.ShardPointer, pool, level)
		if err != nil {
			return nil, err
		}
	}

	return replicas, nil

}

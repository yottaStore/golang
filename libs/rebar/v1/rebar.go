package v1

import (
	"github.com/yottaStore/golang/utils/htree"
	"github.com/yottaStore/golang/utils/record"
)

type Hasher func()
type Scorer func(nodes []*htree.Node, hash []byte, shards int) ([]*htree.Node, error)

type Navigator struct {
	Seed      uint64
	SeedTable *[16][8]uint64
	Hasher    Hasher
	Scorer    Scorer
}

type Options struct {
	Sharding    int
	Replication int
}

func (n *Navigator) FindPools(pointer string, tree *htree.Root, opts Options) ([][]*htree.Node, error) {

	// Find replicas

	replicas, err := findReplicas(pointer, tree, opts)
	if err != nil {
		return nil, err
	}
	// For each replica, find pool
	pools := make([][]*htree.Node, opts.Replication)
	for idx, replica := range replicas {
		pool, err := findPool(replica, opts)
		if err != nil {
			return nil, err
		}
		pools[idx] = pool
	}

	return pools, nil
}

func (n *Navigator) FindRecord(pointer string, pool []*htree.Node) (*htree.Node, error) {

	return findRecord(pointer, pool)

}

func (n *Navigator) Find(record record.Record, tree *htree.Root, opts Options) ([]*htree.Node, error) {

	pools, err := n.FindPools(record.PoolPointer, tree, opts)
	if err != nil {
		return nil, err
	}

	replicas := make([]*htree.Node, 0, opts.Replication)

	for _, pool := range pools {
		node, err := n.FindRecord(record.ShardPointer, pool)
		if err != nil {
			return nil, err
		}
		replicas = append(replicas, node)
	}

	return replicas, nil

}

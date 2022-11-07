package rebar

import (
	"github.com/yottaStore/golang/utils/htree"
	"github.com/yottaStore/golang/utils/record"
)

type Opts struct {
	Sharding    int
	Replication int
	Seeds       *[16][4]uint64
}

func FindPool(record string, tree *htree.Root, opts Opts) ([]*htree.Node, error) {

	return findPool(record, tree.Root, opts)
}

func FindShard(record2 string, pool []*htree.Node, opts Opts) (*htree.Node, error) {

	return nil, nil
}

func Find(record record.Record, tree *htree.Root, opts Opts) (*htree.Node, error) {

	pool, err := FindPool(record.PoolPointer, tree, opts)
	if err != nil {
		return nil, err
	}

	npool, err := FindShard(record.ShardPointer, pool, opts)
	if err != nil {
		return nil, err
	}

	return npool, nil
}

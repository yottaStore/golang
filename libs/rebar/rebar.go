package rebar

import (
	"github.com/yottaStore/golang/libs/gap"
	"github.com/yottaStore/golang/utils/record"
)

type Opts struct {
	Sharding    int
	Replication int
}

func FindCollectionPool(record record.Record, tree gap.NodePool, opts Opts) (gap.NodePool, error) {

	return nil, nil
}

func FindNodePool(record2 record.Record, pool gap.NodePool, opts Opts) (gap.NodePool, error) {

	return nil, nil
}

func Find(record record.Record, tree gap.NodePool, opts Opts) (gap.NodePool, error) {

	pool, err := FindCollectionPool(record, tree, opts)
	if err != nil {
		return nil, err
	}

	npool, err := FindNodePool(record, pool, opts)
	if err != nil {
		return nil, err
	}

	return npool, nil
}

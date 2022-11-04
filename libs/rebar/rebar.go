package rebar

import (
	"github.com/yottaStore/golang/utils/hTrie"
	"github.com/yottaStore/golang/utils/record"
)

type Opts struct {
	Sharding       int
	Replication    int
	VerticalSeed   uint64
	HorizontalSeed uint64
}

func FindPool(record record.Record, tree hTrie.Trie, opts Opts) ([]*hTrie.Node, error) {

	return nil, nil
}

func FindShard(record2 record.Record, pool []*hTrie.Node, opts Opts) (*hTrie.Node, error) {

	return nil, nil
}

func Find(record record.Record, tree hTrie.Trie, opts Opts) (*hTrie.Node, error) {

	pool, err := FindPool(record, tree, opts)
	if err != nil {
		return nil, err
	}

	npool, err := FindShard(record, pool, opts)
	if err != nil {
		return nil, err
	}

	return npool, nil
}

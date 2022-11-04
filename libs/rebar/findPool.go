package rebar

import (
	"github.com/yottaStore/golang/utils/hTrie"
	"github.com/zeebo/xxh3"
)

func navigateTree(nodes []*hTrie.Node, hash []byte, count int) ([]*hTrie.Node, error) {

	for {

		break
	}

	return nil, nil
}

func findReplicas(nodes []*hTrie.Node, hash []byte, count int) ([]*hTrie.Node, error) {

	if len(nodes) >= count {
		return round(nodes, hash, count)
	}

	var newNodes []*hTrie.Node

	for _, node := range nodes {
		newNodes = append(newNodes, node.Children...)
	}

	return findReplicas(newNodes, hash, count)
}

func findPool(record string, tree hTrie.Trie, opts Opts) ([][]*hTrie.Node, error) {

	hash := xxh3.Hash128Seed([]byte(record), opts.VerticalSeed).Bytes()

	roots, err := round(tree.Root.Children, hash[:], opts.Replication)
	if err != nil {
		return nil, err
	}

	pool := make([][]*hTrie.Node, 0, opts.Replication)

	for idx, root := range roots {
		tmp, err := navigateTree(root.Children, hash[:], opts.Sharding)
		if err != nil {
			return nil, err
		}

		pool[idx] = tmp
	}

	return pool, nil
}

type Suggestion interface {
	uint8 | uint16
}

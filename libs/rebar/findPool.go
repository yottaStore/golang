package rebar

import (
	"github.com/yottaStore/golang/utils/htree"
	"github.com/zeebo/xxh3"
)

func navigateTree(nodes []*htree.Node, hash []byte, count int) ([]*htree.Node, error) {

	for {

		break
	}

	return nil, nil
}

func findReplicas(nodes []*htree.Node, hash []byte, count int) ([]*htree.Node, error) {

	if len(nodes) >= count {
		return round(nodes, hash, count)
	}

	var replicaPool []*htree.Node

	for _, node := range nodes {
		replicaPool = append(replicaPool, node.Children...)
	}

	return findReplicas(replicaPool, hash, count)
}

func findPool(record string, replica *htree.Node, opts Opts) ([]*htree.Node, error) {

	children := replica.Children
	var err error

	for i := 0; ; i++ {

		if children[0].Children == nil {
			hash := xxh3.Hash128Seed([]byte(record), opts.Seeds[i][0]).Bytes()
			children, err = round(children, hash[:], opts.Sharding)
			if err != nil {
				return nil, err
			}

			break
		}

		hash := xxh3.Hash128Seed([]byte(record), opts.Seeds[i][0]).Bytes()
		children, err = round(children, hash[:], 1)
		if err != nil {
			return nil, err
		}
	}

	return children, nil

}

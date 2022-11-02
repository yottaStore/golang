package horizontal

import (
	"encoding/binary"
	"github.com/yottaStore/golang/libs/rebar"
	"github.com/yottaStore/golang/utils/record"
	"github.com/zeebo/xxh3"
)

func Find_pool(record record.Record, tree [][]string, sharding, replication int, seed uint64) ([]string, error) {

	// TODO: check count is not greater than tree size
	//count := sharding * replication
	levels := len(tree)

	buff := make([]byte, 8)

	prefix := ""

	max_depth := levels - 1
	for i := 0; i < max_depth; i++ {
		h := rebar.Hash(record.PoolPointer, seed)
		idx := Round(h[:], nil)

		prefix += tree[i][idx[0]]

		binary.BigEndian.PutUint64(buff[:], seed)
		seed = xxh3.Hash(buff)
	}

	return []string{prefix}, nil
}

package horizontal

import (
	"github.com/yottaStore/golang/utils/hTrie"
	"github.com/yottaStore/golang/utils/record"
)

func Find_pool(record record.Record, t hTrie.Trie, sharding, replication int, seed uint64) ([]string, error) {

	// TODO: check count is not greater than tree size
	//count := sharding * replication

	//buff := make([]byte, 8)

	prefix := ""

	return []string{prefix}, nil
}

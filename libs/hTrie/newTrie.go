package hTrie

func NewTrie(shards []Shard, prefix string) (*Trie, error) {
	tree := &Trie{
		Prefix: prefix,
		Root: &Node{
			Pointer: prefix,
		},
		Ops: make([]Update, 10),
	}
	for _, shard := range shards {
		if err := tree.Insert(shard); err != nil {
			return nil, err
		}
	}
	return tree, nil
}

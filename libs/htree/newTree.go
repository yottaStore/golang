package htree

func NewTree(shards []Shard) (Tree, error) {

	var t Tree

	for _, shard := range shards {
		err := t.Insert(shard)
		if err != nil {
			return t, err
		}
	}

	return t, nil
}

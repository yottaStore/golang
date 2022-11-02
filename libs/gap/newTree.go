package gap

import (
	"errors"
)

func NewTree(shards []Shard, levels uint8) (Node, error) {

	root := Node{
		Pointer: "root.com",
	}

	for _, shard := range shards {

		if len(shard.Coords) != int(levels) {
			//log.Println("Shard has wrong number of coords:", shard, int(levels), len(shard.Coords))
			return root, errors.New("shard has wrong number of coordinates")
		}

		err := root.Insert(shard)
		if err != nil {
			return root, err
		}

	}

	return root, nil
}

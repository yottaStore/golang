package gap

import (
	"testing"
)

func getShards() []Shard {

	shards := make([]Shard, 2)

	shards[0] = Shard{
		Weight: 1,
		Coords: []string{".root.com", "zone1.", "dc1.", "rack1.", "node1.", "shard1."},
	}

	shards[1] = Shard{
		Weight: 1,
		Coords: []string{"root.com", "zone1.", "dc1.", "rack1.", "node1.", "shard2."},
	}

	return shards
}

func TestNewTree(t *testing.T) {

	shards := getShards()
	tree, err := NewTree(shards, 6)
	if err != nil {
		t.Error(err)
	}

	t.Log("Tree:", tree)
	t.Log("Child:", tree.Children[0])

}

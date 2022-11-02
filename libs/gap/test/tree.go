package main

import (
	"github.com/yottaStore/golang/libs/gap"
	"log"
)

func printChildren(node *gap.Node) {
	for _, child := range node.Children {
		log.Println("Child:", child)
		printChildren(child)
	}
}

func main() {

	shards := make([]gap.Shard, 3)

	shards[0] = gap.Shard{
		Weight: 1,
		Coords: []string{"root.com", "zone1.", "dc1.", "rack1.", "node1.", "shard1."},
	}

	shards[1] = gap.Shard{
		Weight: 1,
		Coords: []string{"root.com", "zone1.", "dc1.", "rack1.", "node1.", "shard2."},
	}

	shards[2] = gap.Shard{
		Weight: 1,
		Coords: []string{"root.com", "zone1.", "dc1.", "rack1.", "node2.", "shard1."},
	}

	tree, err := gap.NewTree(shards, 6)
	if err != nil {
		log.Fatal(err)
	}

	printChildren(&tree)

	log.Println("Embedded method")

	tree.Print()

}

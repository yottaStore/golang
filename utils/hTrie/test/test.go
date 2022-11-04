package main

import (
	"github.com/yottaStore/golang/utils/hTrie"
	"log"
)

func main() {

	shards := make([]hTrie.Shard, 3)

	shards[0] = hTrie.Shard{
		Weight:  1,
		Pointer: "shard1.node2.dc1.root.com:8082",
	}

	shards[1] = hTrie.Shard{
		Weight:  1,
		Pointer: "shard2.node1.dc1.root.com:8081",
	}

	shards[2] = hTrie.Shard{
		Weight:  1,
		Pointer: "shard1.node1.dc1.root.com:8081",
	}

	t, err := hTrie.NewTrie(shards, "root.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tree: ", t)
	t.Print()

	if err = t.Verify(); err != nil {
		log.Fatal("Verification failed: ", err)
	}

}

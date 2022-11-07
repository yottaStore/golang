package main

import (
	"github.com/yottaStore/golang/libs/rebar"
	"github.com/yottaStore/golang/utils/htree"
	"log"
)

func main() {

	seed := uint64(0xfedcba9876543210)
	table := rebar.SeedTable(seed)

	r, err := htree.SampleTree()
	if err != nil {
		log.Fatal("Error creating tree: ", err)
	}

	//r.Print()

	record := "account@test/recorda"
	opts := rebar.Opts{
		Sharding:    2,
		Replication: 1,
		Seeds:       &table,
	}

	pool, err := rebar.FindPool(record, r, opts)

	if err != nil {
		log.Fatal("Error finding pool: ", err)
	}

	log.Println("Pool: ", pool)
	for _, node := range pool {
		log.Println(node.Pointer)
	}

}

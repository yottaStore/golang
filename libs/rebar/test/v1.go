package main

import (
	"github.com/yottaStore/golang/libs/rebar/v1"
	"github.com/yottaStore/golang/utils/htree"
	"log"
)

func main() {

	var seeds [2]uint64
	seeds[0] = uint64(0xfedcba9876543210)
	seeds[1] = uint64(0xfedcba9876543210)

	n, err := v1.New(seeds)
	if err != nil {
		log.Fatal("Error creating rebar: ", err)
	}

	r, err := htree.SampleTree()
	if err != nil {
		log.Fatal("Error creating tree: ", err)
	}

	//r.Print()

	record := "account@test/recorda"
	opts := v1.Options{
		Sharding:    3,
		Replication: 2,
	}

	pools, l, err := n.FindPools(record, r, opts)

	if err != nil {
		log.Fatal("Error finding pool: ", err)
	}

	log.Println("Level: ", l)
	for idx, pool := range pools {
		log.Println("Pool: ", idx)
		for _, node := range pool {
			log.Println(node)
		}
	}

}

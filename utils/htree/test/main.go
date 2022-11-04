package main

import (
	"github.com/yottaStore/golang/utils/htree"
	"log"
)

func main() {

	r, err := htree.New("root.com")
	if err != nil {
		log.Fatal("Error creating tree: ", err)
	}

	us := []htree.Update{
		{
			Type:   htree.INSERT_SHARD,
			Url:    "s1.n1.dc1.root.com",
			Weight: 10,
			Load:   0,
		},
	}

	err = r.Insert(us)

	r.Print()

	if err != nil {
		log.Fatal("Error inserting nodes: ", err)
	}

	err = r.Verify()
	if err != nil {
		log.Fatal("Error verifying tree: ", err)
	}

	r.Print()

}

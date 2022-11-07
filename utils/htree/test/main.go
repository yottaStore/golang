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
			Url:    "s2.n1.dc1.root.com:8081",
			Weight: 5,
			Load:   0,
		},
		{
			Type:   htree.INSERT_SHARD,
			Url:    "s1.n1.dc1.root.com:8081",
			Weight: 10,
			Load:   0,
		},
		{
			Type:   htree.INSERT_SHARD,
			Url:    "s3.n1.dc1.root.com:8081",
			Weight: 5,
			Load:   0,
		},
		{
			Type:   htree.UPDATE_LOAD,
			Url:    "n1.dc1.root.com",
			Weight: 0,
			Load:   5,
		},
		{
			Type: htree.DELETE_NODE,
			Url:  "s3.n1.dc1.root.com:8081",
		},
	}

	err = r.Update(us)

	//r.Print()

	if err != nil {
		log.Fatal("Error updating: ", err)
	}

	/*err = r.Root.Update()
	if err != nil {
		log.Fatal("Error updating tree: ", err)
	}*/

	err = r.Verify()
	if err != nil {
		log.Fatal("Error verifying tree: ", err)
	}

	r.Print()

}

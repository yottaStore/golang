package rendezvous

import (
	"errors"
	"log"
)

func FindRecord(r ParsedRecord, nodemap Nodemap, opts Options, hashkey string) (Nodemap, error) {

	if hashkey == "" {
		return nil, errors.New("hashkey cannot be empty")
	}

	log.Println("Node map: ", nodemap)

	pool, err := FindPool(r, nodemap, opts, hashkey)
	if err != nil {
		log.Println("Error: couldn't find the pool")
		return nil, err
	}

	log.Println("Pool: ", pool)

	nodes, err := FindRecordInPool(r, pool, opts, hashkey)
	if err != nil {
		log.Println("Error: couldn't find the record in the pool")
		return nil, err
	}

	log.Println("Nodes: ", nodes)

	return nodes, nil

}

func FindPool(r ParsedRecord, nodemap Nodemap, opts Options, hashkey string) (Nodemap, error) {

	return findNodes(r.Collection, nodemap, opts, hashkey)

}

func FindRecordInPool(r ParsedRecord, pool Nodemap, opts Options, hashkey string) (Nodemap, error) {

	return findNodes(r.Record, pool, opts, hashkey)

}

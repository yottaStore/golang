package document

import (
	"yottadb/dbdriver"
	"yottadb/rendezvous"
)

type Driver struct {
	dbdriver.Interface
	Finder   rendezvous.Finder
	NodeTree *[]string
}

func New(hashKey string, nodeTree *[]string) (dbdriver.Interface, error) {

	f := rendezvous.Finder{
		HashKey: hashKey,
	}

	d := Driver{
		Finder:   f,
		NodeTree: nodeTree,
	}

	return d, nil

}

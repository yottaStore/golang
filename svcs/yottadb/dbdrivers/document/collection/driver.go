package collection

import (
	"github.com/fxamacker/cbor/v2"
	"log"
	"rendezvous"
	"yottadb/dbdrivers/document"
	"yottadb/dbdrivers/keyval"
	"yottafs/client"
	"yottafs/iodrivers"
)

const (
	collectionSharding    = 1
	collectionReplication = 1
)

func getCollectionOpts() rendezvous.Options {

	return rendezvous.Options{1, 1}
}

type Driver struct {
	Nodemap  *[]string
	Hashkey  string
	Kvdriver keyval.Driver
}

func findNode(req document.Request, d Driver) ([]string, string, error) {
	parsedRecord, err := rendezvous.ParseRecord(req.Record)
	if err != nil {
		log.Println("Error parsing record driver call: ", err)
		return nil, "", err
	}

	nodes, err := rendezvous.FindRecord(parsedRecord, *d.Nodemap, getCollectionOpts(), d.Hashkey)
	if err != nil {
		log.Println("Error during read rendezvous")
		return nil, "", err
	}

	path := "collectionAdmin/" + parsedRecord.Account + "/" + parsedRecord.Collection

	return nodes, path, nil
}

func (d Driver) Create(req document.Request) (iodrivers.Response, error) {

	var resp iodrivers.Response

	nodes, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return resp, err
	}

	// TODO: write all nodes
	node := nodes[0]

	opts := rendezvous.Options{1, 1}
	data, err := cbor.Marshal(opts)
	if err != nil {
		log.Println("Error: Couldn't marshal collection!")

	}

	return client.Create(path, node, data)

}

func (d Driver) Read(req document.Request) (rendezvous.Options, error) {

	var opts rendezvous.Options

	nodes, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return opts, err
	}

	// TODO: pick randomly
	node := nodes[0]

	resp, err := client.Read(path, node)
	if err != nil {
		log.Println("Error: Couldn't find collection!")
		return opts, err
	}

	err = cbor.Unmarshal(resp.Data, &opts)
	if err != nil {
		log.Println("Error: Couldn't unmarshal collection!")
		return opts, err
	}

	return opts, nil

}

func (d Driver) Update(req document.Request) (iodrivers.Response, error) {
	var res iodrivers.Response

	// TODO: update collection

	return res, nil
}

func (d Driver) Delete(req document.Request) error {

	nodes, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return err
	}

	// TODO: pick randomly
	node := nodes[0]

	return client.Delete(path, node)
}

func New(nodes *[]string, hashkey string, kvd keyval.Driver) (Driver, error) {

	d := Driver{nodes, hashkey, kvd}

	return d, nil

}

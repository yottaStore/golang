package document

import (
	"github.com/fxamacker/cbor/v2"
	"log"
	"rendezvous"
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

type CollDriver struct {
	Nodemap  *[]string
	Hashkey  string
	Kvdriver keyval.Driver
}

func findCollNode(req Request, d CollDriver) ([]string, string, error) {
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

func (d CollDriver) Create(req Request) (iodrivers.Response, error) {

	var resp iodrivers.Response

	nodes, path, err := findCollNode(req, d)
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

func (d CollDriver) Read(req Request) (rendezvous.Options, error) {

	var opts rendezvous.Options

	nodes, path, err := findCollNode(req, d)
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

func (d CollDriver) Update(req Request) (iodrivers.Response, error) {
	var res iodrivers.Response

	// TODO: update collection

	return res, nil
}

func (d CollDriver) Delete(req Request) error {

	nodes, path, err := findCollNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return err
	}

	// TODO: pick randomly
	node := nodes[0]

	return client.Delete(path, node)
}

func NewColl(nodes *[]string, hashkey string, kvd keyval.Driver) (CollDriver, error) {

	d := CollDriver{nodes, hashkey, kvd}

	return d, nil

}

package document

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/dbdriver/keyvalue"
	"yottadb/rendezvous"
)

type Driver struct {
	NodeTree         *[]string
	HashKey          string
	CollectionDriver CollDriver
	KVDriver         keyvalue.Driver
}

type CReq struct {
	Path   string
	Method string
	Data   []byte
}

func New(hashKey string, nodeTree *[]string) (Driver, error) {

	var d Driver

	cd, err := NewColl(hashKey, nodeTree)
	if err != nil {
		return d, err
	}

	kvd, err := keyvalue.New(hashKey, nodeTree)
	if err != nil {
		return d, err
	}

	d = Driver{
		HashKey:          hashKey,
		NodeTree:         nodeTree,
		CollectionDriver: cd,
		KVDriver:         kvd}

	return d, nil

}

func (d Driver) Read(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	// Find nodes
	// TODO: fix rendezvous
	// TODO: improve API shape

	// Get collection detail

	coll, err := d.CollectionDriver.Read(req)
	if err != nil {
		return resp, err
	}
	log.Println(coll)

	// Use collection as input

	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
		HashKey:     d.HashKey,
	}

	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick a shard at random and verify others
	node := shards[0]

	// Issue read
	buff, err := cbor.Marshal(CReq{
		Path:   parsedRecord.RecordIdentifier,
		Method: "read"})
	if err != nil {
		return resp, err
	}
	if err != nil {
		return resp, err
	}
	result, err := http.Post(node+"/yottafs/", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		return resp, err
	}
	body, err := io.ReadAll(result.Body)

	if result.StatusCode != http.StatusOK {
		return resp, errors.New(string(body))
	}

	resp.Data = body

	// Return result
	return resp, nil
}

func (d Driver) Write(req dbdriver.Request) (dbdriver.Response, error) {

	log.Println("Writing...")
	log.Println(req)

	var resp dbdriver.Response

	// Get collection detail

	coll, err := d.CollectionDriver.Read(req)
	if err != nil {
		return resp, err
	}

	log.Println(coll)

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
		HashKey:     d.HashKey,
	}

	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		log.Println("Error: ", err)
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	// Issue write
	buff, err := cbor.Marshal(CReq{
		Path:   parsedRecord.RecordIdentifier,
		Method: "write",
		Data:   []byte(req.Data)})
	if err != nil {
		return resp, err
	}

	result, err := http.Post(node+"/yottafs/", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		return resp, err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		return resp, err
	}
	if result.StatusCode != http.StatusOK {
		return resp, errors.New(string(buff))
	}

	return resp, nil
}

func (d Driver) Delete(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response

	coll, err := d.CollectionDriver.Read(req)
	if err != nil {
		return resp, err
	}
	log.Println(coll)

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
		HashKey:     d.HashKey,
	}

	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	// Issue write
	buff, err := cbor.Marshal(CReq{
		Path:   parsedRecord.RecordIdentifier,
		Method: "delete"})
	if err != nil {
		return resp, err
	}
	if err != nil {
		return resp, err
	}

	result, err := http.Post(node+"/yottafs/", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		return resp, err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		return resp, err
	}
	if result.StatusCode != http.StatusOK {
		return resp, errors.New(string(buff))
	}

	return resp, nil

}

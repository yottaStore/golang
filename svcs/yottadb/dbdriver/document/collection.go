package document

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/rendezvous"
)

const (
	CollectionSharding    = 1
	CollectionReplication = 1
)

type Collection struct {
	Replication int
	Sharding    int
}

type CollDriver struct {
	NodeTree *[]string
	HashKey  string
}

func NewColl(hashKey string, nodeTree *[]string) (CollDriver, error) {

	d := CollDriver{
		HashKey:  hashKey,
		NodeTree: nodeTree,
	}

	return d, nil

}

func (d CollDriver) Read(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	// Find nodes
	// TODO: fix rendezvous
	// TODO: improve API shape
	opts := rendezvous.RendezvousOptions{
		Replication: CollectionReplication,
		Sharding:    CollectionSharding,
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

func (d CollDriver) Write(req dbdriver.Request) (dbdriver.Response, error) {

	log.Println("Writing...")

	var resp dbdriver.Response

	log.Println(req)

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: CollectionReplication,
		Sharding:    CollectionSharding,
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

func (d CollDriver) Delete(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: CollectionReplication,
		Sharding:    CollectionSharding,
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

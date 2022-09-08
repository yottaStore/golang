package document

import (
	"encoding/json"
	"strings"
	"yottadb/dbdriver"
	"yottadb/dbdriver/keyvalue"
	"yottadb/rendezvous"
)

type Driver struct {
	KVdriver keyvalue.Driver
	Finder   rendezvous.Finder
	NodeTree *[]string
}

type CollectionResponse struct {
	Sharding    int
	Replication int
}

const (
	collectionPath        = "yotta@collections/"
	collectionSharding    = 1
	collectionReplication = 1
)

func (d Driver) ReadRecord(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	collectionName := strings.Split(req.Path, "/")[1]

	cReq := dbdriver.Request{
		Path: collectionPath + collectionName,
	}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cJson CollectionResponse
	if err := json.Unmarshal([]byte(cResp.Data), &cJson); err != nil {
		return resp, err
	}

	req.Rendezvous.Sharding = cJson.Sharding
	req.Rendezvous.Replication = cJson.Replication

	resp, err = d.KVdriver.Read(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) WriteRecord(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	collectionName := strings.Split(req.Path, "/")[1]

	cReq := dbdriver.Request{
		Path: collectionPath + collectionName,
	}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cJson CollectionResponse
	if err := json.Unmarshal([]byte(cResp.Data), &cJson); err != nil {
		return resp, err
	}

	req.Rendezvous.Sharding = cJson.Sharding
	req.Rendezvous.Replication = cJson.Replication

	resp, err = d.KVdriver.Write(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) UpdateRecord() {

}

func (d Driver) DeleteRecord(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	collectionName := strings.Split(req.Path, "/")[0]

	cReq := dbdriver.Request{
		Path: collectionPath + collectionName,
	}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cJson CollectionResponse
	if err := json.Unmarshal([]byte(cResp.Data), &cJson); err != nil {
		return resp, err
	}

	req.Rendezvous.Sharding = cJson.Sharding
	req.Rendezvous.Replication = cJson.Replication

	resp, err = d.KVdriver.Delete(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) CreateCollection(req dbdriver.Request) (dbdriver.Response, error) {

	// TODO: check if collection exists already
	var resp dbdriver.Response

	req.Path = ""
	req.Rendezvous.Sharding = collectionSharding
	req.Rendezvous.Replication = collectionReplication

	resp, err := d.KVdriver.Write(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) UpdateCollection() {

}

func (d Driver) DeleteCollection() {

}

func New(hashKey string, nodeTree *[]string) (Driver, error) {

	f := rendezvous.Finder{
		HashKey: hashKey,
	}

	d := Driver{
		Finder:   f,
		NodeTree: nodeTree,
	}

	return d, nil

}

package document

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"yottadb/dbdriver"
	"yottadb/dbdriver/keyvalue"
	"yottadb/rendezvous"
)

type Driver struct {
	KVdriver dbdriver.Interface
	Finder   rendezvous.Finder
	NodeTree *[]string
}

type CollectionRecord struct {
	Sharding    int
	Replication int
}

const (
	collectionPath        = "yotta@collections"
	collectionSharding    = 1
	collectionReplication = 1
)

func (d Driver) ReadDocument(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	parsedRecord, err := d.Finder.ParseRecord(req.Path)
	if err != nil {
		return resp, err
	}

	cOpts := dbdriver.RendezvousOpts{
		Sharding:    collectionSharding,
		Replication: collectionReplication,
	}
	cReq := dbdriver.Request{
		Path:       collectionPath + parsedRecord.CollectionPointer,
		Rendezvous: cOpts}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cResult dbdriver.Response
	err = json.Unmarshal([]byte(cResp.Data), &cResult)
	if err != nil {
		return resp, err
	}

	log.Println("Result: ", cResult)

	cRecord := strings.Replace(cResult.Data, "A", "", -1)
	cRecord = strings.Replace(cRecord, "=", "", -1)

	cData, err := base64.StdEncoding.DecodeString(cRecord)
	if err != nil {
		return resp, err
	}

	var cJson CollectionRecord
	if err := json.Unmarshal(cData, &cJson); err != nil {
		return resp, err
	}

	log.Println("Cjson: ", cJson)

	req.Rendezvous.Sharding = cJson.Sharding
	req.Rendezvous.Replication = cJson.Replication

	resp, err = d.KVdriver.Read(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) WriteDocument(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	parsedRecord, err := d.Finder.ParseRecord(req.Path)
	if err != nil {
		return resp, err
	}

	cOpts := dbdriver.RendezvousOpts{
		Sharding:    collectionSharding,
		Replication: collectionReplication,
	}

	cReq := dbdriver.Request{
		Path:       collectionPath + parsedRecord.CollectionPointer,
		Rendezvous: cOpts}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cResult dbdriver.Response
	err = json.Unmarshal([]byte(cResp.Data), &cResult)
	if err != nil {
		return resp, err
	}

	log.Println("Result: ", cResult)

	cRecord := strings.Replace(cResult.Data, "A", "", -1)
	cRecord = strings.Replace(cRecord, "=", "", -1)

	cData, err := base64.StdEncoding.DecodeString(cRecord)
	if err != nil {
		return resp, err
	}

	var cJson CollectionRecord
	if err := json.Unmarshal(cData, &cJson); err != nil {
		return resp, err
	}

	log.Println("Cjson: ", cJson)

	req.Rendezvous.Sharding = cJson.Sharding
	req.Rendezvous.Replication = cJson.Replication

	resp, err = d.KVdriver.Write(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) UpdateDocument() {

}

func (d Driver) DeleteDocument(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	parsedRecord, err := d.Finder.ParseRecord(req.Path)
	if err != nil {
		return resp, err
	}

	cOpts := dbdriver.RendezvousOpts{
		Sharding:    collectionSharding,
		Replication: collectionReplication,
	}
	cReq := dbdriver.Request{
		Path:       collectionPath + parsedRecord.CollectionPointer,
		Rendezvous: cOpts}

	cResp, err := d.KVdriver.Read(cReq)
	if err != nil {
		return resp, err
	}

	var cResult dbdriver.Response
	err = json.Unmarshal([]byte(cResp.Data), &cResult)
	if err != nil {
		return resp, err
	}

	log.Println("Result: ", cResult)

	cRecord := strings.Replace(cResult.Data, "A", "", -1)
	cRecord = strings.Replace(cRecord, "=", "", -1)

	cData, err := base64.StdEncoding.DecodeString(cRecord)
	if err != nil {
		return resp, err
	}

	var cJson CollectionRecord
	if err := json.Unmarshal(cData, &cJson); err != nil {
		return resp, err
	}

	log.Println("Cjson: ", cJson)

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

	parsedRecord, err := d.Finder.ParseRecord(req.Path)

	req.Path = collectionPath + parsedRecord.CollectionPointer
	req.Rendezvous.Sharding = collectionSharding
	req.Rendezvous.Replication = collectionReplication

	recordData := CollectionRecord{
		Sharding:    1,
		Replication: 1,
	}

	buff, err := json.Marshal(recordData)
	if err != nil {
		return resp, err
	}

	req.Data = string(buff)

	resp, err = d.KVdriver.Write(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (d Driver) UpdateCollection() {

}

func (d Driver) DeleteCollection(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response

	parsedRecord, err := d.Finder.ParseRecord(req.Path)

	req.Path = collectionPath + parsedRecord.CollectionPointer
	req.Rendezvous.Sharding = collectionSharding
	req.Rendezvous.Replication = collectionReplication

	resp, err = d.KVdriver.Delete(req)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func New(hashKey string, nodeTree *[]string) (Driver, error) {

	f := rendezvous.Finder{
		HashKey: hashKey,
	}

	kvDriver, err := keyvalue.New(hashKey, nodeTree)

	d := Driver{
		Finder:   f,
		NodeTree: nodeTree,
		KVdriver: kvDriver,
	}
	if err != nil {
		return d, err
	}

	return d, nil

}

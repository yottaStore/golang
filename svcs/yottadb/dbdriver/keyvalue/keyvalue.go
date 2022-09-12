package keyvalue

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/rendezvous"
)

type Driver struct {
	dbdriver.Interface
	NodeTree *[]string
	HashKey  string
}

type CReq struct {
	Path   string
	Method string
	Data   []byte
}

func New(hashKey string, nodeTree *[]string) (Driver, error) {

	d := Driver{
		HashKey:  hashKey,
		NodeTree: nodeTree,
	}

	return d, nil

}

func (d Driver) Read(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	// Find nodes
	// TODO: fix rendezvous
	// TODO: improve API shape
	opts := rendezvous.RendezvousOptions{
		Replication: req.Rendezvous.Replication,
		Sharding:    req.Rendezvous.Sharding,
		HashKey:     d.HashKey,
	}

	if opts.Sharding == 0 || opts.Replication == 0 {
		log.Println("Error with rendezvous options")
		return resp, errors.New("Error with rendezvous options")
	}

	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick a shard at random and verify others
	node := shards[0]

	/*fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)
	fmt.Println("Node picked: ", node)*/

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
	//log.Println("Json data: ", string(json_data))
	result, err := http.Post(node+"/yottafs/", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		return resp, err
	}
	body, err := io.ReadAll(result.Body)

	if result.StatusCode != http.StatusOK {
		return resp, errors.New(string(body))
	}

	resp.Data = string(body)

	// Return result
	return resp, nil
}

func (d Driver) Write(req dbdriver.Request) (dbdriver.Response, error) {

	log.Println("Writing...")

	var resp dbdriver.Response

	log.Println(req)

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: req.Rendezvous.Replication,
		Sharding:    req.Rendezvous.Sharding,
		HashKey:     d.HashKey,
	}

	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		log.Println("Error: ", err)
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	/*fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)*/
	fmt.Println("Node picked: ", node)

	// Issue write
	/*values := map[string]interface{}{
		"Path":   parsedRecord.RecordIdentifier,
		"Method": "write",
		"Data":   req.Data}
	json_data, err := json.Marshal(values)*/

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

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: req.Rendezvous.Replication,
		Sharding:    req.Rendezvous.Sharding,
		HashKey:     d.HashKey,
	}
	shards, _, parsedRecord, err := rendezvous.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	/*fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)
	fmt.Println("Node picked: ", node)*/

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

	//log.Println("Json data: ", string(json_data))
	//fmt.Println(node)
	result, err := http.Post(node+"/yottafs/", "application/json",
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

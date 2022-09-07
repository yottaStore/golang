package keyvalue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/rendezvous"
)

type Driver struct {
	dbdriver.Interface
	Finder   rendezvous.Finder
	NodeTree *[]string
}

func New() (dbdriver.Interface, error) {

	var d Driver

	return d, nil

}

func (d Driver) Read(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response
	// Find nodes
	// TODO: fix rendezvous
	// TODO: improve API shape
	// TODO: check collection for size
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick a shard at random and verify others
	node := shards[0]

	fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)
	fmt.Println("Node picked: ", node)

	// Issue read
	values := map[string]interface{}{"Path": parsedRecord.RecordIdentifier, "Method": "read"}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}
	fmt.Println("Json data: ", string(json_data))
	result, err := http.Post(node+"/yottafs/", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return resp, err
	}
	body, err := io.ReadAll(result.Body)

	resp.Data = body

	// Return result
	return resp, nil
}

func (d Driver) Write(req dbdriver.Request) (dbdriver.Response, error) {

	var resp dbdriver.Response

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)
	fmt.Println("Node picked: ", node)

	// Issue write
	values := map[string]interface{}{
		"Path":   parsedRecord.RecordIdentifier,
		"Method": "write",
		"Data":   req.Data}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}

	fmt.Println("Json data: ", string(json_data))
	fmt.Println(node)
	result, err := http.Post(node+"/yottafs/", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return resp, err
	}

	buff, err := io.ReadAll(result.Body)
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
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.NodeTree, opts)
	if err != nil {
		return resp, err
	}

	// TODO: pick all shards
	node := shards[0]

	fmt.Println("Record: ", parsedRecord)
	fmt.Println("Node tree: ", nodes)
	fmt.Println("Shards pool:", shards)
	fmt.Println("Node picked: ", node)

	// Issue write
	values := map[string]interface{}{
		"Path":   parsedRecord.RecordIdentifier,
		"Method": "delete"}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}

	fmt.Println("Json data: ", string(json_data))
	fmt.Println(node)
	result, err := http.Post(node+"/yottafs/", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return resp, err
	}

	buff, err := io.ReadAll(result.Body)
	if err != nil {
		return resp, err
	}
	if result.StatusCode != http.StatusOK {
		return resp, errors.New(string(buff))
	}

	return resp, nil

}

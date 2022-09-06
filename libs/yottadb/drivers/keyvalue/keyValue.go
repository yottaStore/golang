package keyvalue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"yottadb"
	"yottadb/rendezvous"
)

type KVDriver yottadb.DbDriver

func (d KVDriver) Read(req yottadb.ReadRequest) (yottadb.ReadResponse, error) {

	var resp yottadb.ReadResponse
	// Find nodes
	// TODO: fix rendezvous
	// TODO: improve API shape
	// TODO: check collection for size
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.Nodes, opts)
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
	values := map[string]interface{}{"Path": parsedRecord.RecordIdentifier}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}
	fmt.Println("Json data: ", string(json_data))
	result, err := http.Post(node+"/yottafs/read", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return resp, err
	}
	body, err := io.ReadAll(result.Body)

	resp.Data = body

	// Return result
	return resp, nil
}

func (d KVDriver) Write(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.Nodes, opts)
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
	values := map[string]interface{}{"Path": parsedRecord.RecordIdentifier, "Data": req.Data}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}

	fmt.Println("Json data: ", string(json_data))
	fmt.Println(node)
	result, err := http.Post(node+"/yottafs/write", "application/json",
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

func (d KVDriver) Update(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, errors.New("Method not implemented yet")
}

func (d KVDriver) Append(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, nodes, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.Nodes, opts)
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
	values := map[string]interface{}{"Path": parsedRecord.RecordIdentifier, "Data": req.Data}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}

	fmt.Println("Json data: ", string(json_data))
	fmt.Println(node)
	result, err := http.Post(node+"/yottafs/write", "application/json",
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

func (d KVDriver) Delete(req yottadb.WriteRequest) error {

	// Find nodes
	opts := rendezvous.RendezvousOptions{
		Replication: 1,
		Sharding:    1,
	}
	shards, _, parsedRecord, err := d.Finder.FindRecord(req.Path, *d.Nodes, opts)
	if err != nil {
		return err
	}

	// TODO: pick all shards
	node := shards[0]

	// Issue write
	values := map[string]interface{}{"Path": parsedRecord.RecordIdentifier}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	fmt.Println("Json data: ", string(json_data))
	fmt.Println(node)
	result, err := http.Post(node+"/yottafs/delete", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}
	if result.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil
}

func New(nodes *[]string, hashKey string) (yottadb.Interface, error) {

	finder := rendezvous.Finder{
		hashKey,
	}

	dbDriver := KVDriver{
		Nodes:  nodes,
		Finder: finder,
	}

	return dbDriver, nil

}

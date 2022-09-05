package document

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

type Driver struct {
	Nodes  *[]string
	Finder rendezvous.Finder
}

func (d Driver) Read(req yottadb.ReadRequest) (yottadb.ReadResponse, error) {

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

	// Return result
	return resp, nil
}

func (d Driver) Write(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

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

	return resp, nil

}

func (d Driver) Update(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, errors.New("Method not implemented yet")
}

func (d Driver) Append(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

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

	return resp, nil
}

func (d Driver) Delete(req yottadb.WriteRequest) error {

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

func (d Driver) GetCollection(req yottadb.ReadRequest) (yottadb.ReadResponse, error) {

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

	// Return result
	return resp, nil
}

func (d Driver) CreateCollection(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

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

	return resp, nil

}

func (d Driver) UpdateCollection(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, errors.New("Method not implemented yet")
}

func (d Driver) DeleteCollection(req yottadb.WriteRequest) error {

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

	dbDriver := Driver{
		nodes,
		finder,
	}

	return dbDriver, nil

}

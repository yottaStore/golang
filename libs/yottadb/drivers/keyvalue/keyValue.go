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

type Driver struct {
	Nodes  *[]string
	Finder rendezvous.Finder
}

func (d Driver) Read(req yottadb.ReadRequest) (yottadb.ReadResponse, error) {

	var resp yottadb.ReadResponse
	// Find nodes
	parsedRecord, err := d.Finder.ParseRecord(req.Path)
	if err != nil {
		return resp, err
	}

	// TODO: fix rendezvous
	nodes, err := d.Finder.FindNodes(parsedRecord, *d.Nodes, 1)
	// Find shard
	node := nodes[1]

	// Issue read
	values := map[string]string{"Path": req.Path}
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}
	url := node + "/read/"
	result, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return resp, err
	}
	body, err := io.ReadAll(result.Body)

	resp.Data = body

	// Return result
	return resp, nil
}

func (d Driver) Write(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	// Find nodes
	parsedRecord, err := d.Finder.ParseRecord(req.Path)
	if err != nil {
		return resp, err
	}

	nodes, err := d.Finder.FindNodes(parsedRecord, *d.Nodes, 1)
	// Find shard
	node := nodes[1]

	// Issue write
	values := ""
	json_data, err := json.Marshal(values)
	if err != nil {
		return resp, err
	}

	fmt.Println("Json data: ", string(json_data))
	result, err := http.Post(node, "application/json",
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

func (d Driver) Update(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, nil
}

func (d Driver) Append(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, nil
}

func (d Driver) Delete(req yottadb.WriteRequest) error {

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

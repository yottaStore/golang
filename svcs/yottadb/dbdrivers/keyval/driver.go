package keyval

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"rendezvous"
	"yottafs/iodrivers"
)

type Driver struct {
	Nodemap *[]string
	Hashkey string
}

type Request struct {
	Record     string
	Rendezvous rendezvous.Options
	Data       []byte
}

type Response struct {
	Record   string
	Data     []byte
	AbaToken []byte
}

func findNode(req Request, d Driver) (string, string, error) {
	parsedRecord, err := rendezvous.ParseRecord(req.Record)
	if err != nil {
		log.Println("Error parsing record driver call: ", err)
		return "", "", err
	}

	nodes, err := rendezvous.FindRecord(parsedRecord, *d.Nodemap, req.Rendezvous, d.Hashkey)
	if err != nil {
		log.Println("Error during read rendezvous")
		return "", "", err
	}

	// TODO: pick node randomly
	node := nodes[0]

	path := parsedRecord.Account + "/" + parsedRecord.Collection + parsedRecord.Record

	return node, path, nil
}

func (d Driver) Create(req Request) (Response, error) {
	var resp Response

	node, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return Response{}, err
	}

	log.Println("Node: ", node)
	log.Println("Path: ", path)

	yfsReq := iodrivers.Request{
		Path: path,
		Data: req.Data,
	}

	buff, err := cbor.Marshal(yfsReq)
	if err != nil {
		log.Println("Error marshalling during create request: ", err)
		return resp, err
	}

	result, err := http.Post(node+"/yottafs/create", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return resp, err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		log.Println("Error reading yottafs response: ", err)
		return resp, err
	}
	if result.StatusCode != http.StatusOK {
		log.Println("Create request failed: ", err)
		return resp, errors.New(string(buff))
	}

	return resp, nil
}

func (d Driver) Read(req Request) (Response, error) {
	var resp Response

	node, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return Response{}, err
	}

	log.Println("Node: ", node)
	log.Println("Path: ", path)

	yfsReq := iodrivers.Request{
		Path: path,
		Data: req.Data,
	}

	buff, err := cbor.Marshal(yfsReq)
	if err != nil {
		log.Println("Error marshalling during read request: ", err)
		return resp, err
	}

	result, err := http.Post(node+"/yottafs/read", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return resp, err
	}

	if result.StatusCode != http.StatusOK {
		log.Println("Read request failed: ", err)
		buff, err = io.ReadAll(result.Body)
		if err != nil {
			log.Println("Error reading yottafs response: ", err)
			return resp, err
		}
		return resp, errors.New(string(buff))
	}

	decoder := cbor.NewDecoder(result.Body)

	err = decoder.Decode(&resp)
	if err != nil {
		log.Println("Error unmarshaling yottafs response")
		return resp, err
	}

	return resp, nil
}

func (d Driver) Update(req Request) (Response, error) {
	var res Response
	return res, nil
}

func (d Driver) CompareAndSwap(req Request) (Response, error) {
	var res Response
	return res, nil
}

func (d Driver) Delete(req Request) error {

	node, path, err := findNode(req, d)
	if err != nil {
		log.Println("Error: Couldn't find node!")
		return err
	}

	yfsReq := iodrivers.Request{
		Path: path,
	}

	buff, err := cbor.Marshal(yfsReq)
	if err != nil {
		log.Println("Error marshalling during read request: ", err)
		return err
	}

	result, err := http.Post(node+"/yottafs/delete", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		log.Println("Error reading yottafs response: ", err)
		return err
	}
	if result.StatusCode != http.StatusOK {
		log.Println("Read request failed: ", err)
		return errors.New(string(buff))
	}

	return nil
}

func New(nodes *[]string, hashkey string) (Driver, error) {

	var d Driver

	if len(*nodes) == 0 {
		return d, errors.New("node map is empty")
	}

	if hashkey == "" {
		return d, errors.New("hash key is null")
	}

	d = Driver{
		nodes,
		hashkey,
	}

	return d, nil

}

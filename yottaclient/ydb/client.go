package ydb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"io"
	"net/http"
)

type Request struct {
	Record     string
	Data       []byte
	Rendezvous Options
}

type Options struct {
	Replication int
	Sharding    int
}

type Client struct {
	Url string
}

func (c Client) Read(record string) ([]byte, error) {
	// TODO: switch from interface to yottadb.Request

	// var req keyval.Request

	req := Request{
		Record: record,
		Rendezvous: Options{
			1, 1,
		},
	}

	/*values := map[string]string{
	"Record": record,
	"Method": "read"}*/
	buff, err := cbor.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/keyval/read",
		"application/json",
		bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	buff, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil
}

func (c Client) Create(record string, data []byte) ([]byte, error) {
	req := Request{
		Record: record,
		Data:   data,
		Rendezvous: Options{
			1, 1,
		},
	}
	buff, err := cbor.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/keyval/create",
		"application/json",
		bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	buff, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil
}

func (c Client) Delete(record string) ([]byte, error) {
	req := Request{
		Record: record,
		Rendezvous: Options{
			1, 1,
		},
	}
	buff, err := cbor.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/keyval/delete",
		"application/json",
		bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	buff, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil
}

func New(url string) (Client, error) {

	client := Client{
		url,
	}

	resp, err := http.Get(url + "/version")
	if err != nil {
		return client, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return client, err
	}

	if resp.StatusCode != http.StatusOK {
		return client, errors.New(string(buff))
	}

	fmt.Println("Connection string: ", string(buff))

	return client, nil
}

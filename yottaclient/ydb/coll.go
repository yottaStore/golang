package ydb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"io"
	"net/http"
)

type CollClient struct {
	Url string
}

func (c CollClient) Read(record string) ([]byte, error) {
	// TODO: switch from interface to yottadb.Request

	// var req keyval.Request

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

	resp, err := http.Post(c.Url+"/yottadb/document/readDocument",
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

func (c CollClient) Create(record string, data []byte) ([]byte, error) {
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

	resp, err := http.Post(c.Url+"/yottadb/document/createDocument",
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

func (c CollClient) CreateColl(record string, data []byte) ([]byte, error) {
	req := Request{
		Record: record,
	}
	buff, err := cbor.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/document/createCollection",
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

func (c CollClient) Delete(record string) ([]byte, error) {
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

	resp, err := http.Post(c.Url+"/yottadb/document/deleteDocument",
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

func NewDoc(url string) (CollClient, error) {

	client := CollClient{
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

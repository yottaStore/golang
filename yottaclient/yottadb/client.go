package yottadb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"yottadb/dbdriver"
)

type Client struct {
	Url string
}

func (c Client) Read(path string, opts dbdriver.RendezvousOpts) ([]byte, error) {
	values := map[string]interface{}{
		"Path":       path,
		"Method":     "read",
		"Rendezvous": opts,
		"Driver":     "keyvalue"}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil
}

func (c Client) Write(path string, data []byte, opts dbdriver.RendezvousOpts) ([]byte, error) {
	values := map[string]interface{}{
		"Path":       path,
		"Method":     "write",
		"Driver":     "keyvalue",
		"Data":       string(data),
		"Rendezvous": opts}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil

}

func (c Client) Delete(path string, opts dbdriver.RendezvousOpts) error {
	values := map[string]interface{}{
		"Path":       path,
		"Method":     "delete",
		"Driver":     "keyvalue",
		"Rendezvous": opts}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil

}

func (c Client) ReadDocument(path string, opts dbdriver.RendezvousOpts) ([]byte, error) {
	values := map[string]interface{}{
		"Path":       path,
		"Method":     "readDocument",
		"Rendezvous": opts}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil
}

func (c Client) WriteDocument(path string, data []byte, opts dbdriver.RendezvousOpts) ([]byte, error) {

	values := map[string]interface{}{
		"Path":       path,
		"Driver":     "document",
		"Method":     "create",
		"Data":       data,
		"Rendezvous": opts}

	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	log.Println("Json data: ", string(json_data))

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil

}

func (c Client) DeleteDocument(path string, opts dbdriver.RendezvousOpts) error {
	values := map[string]interface{}{
		"Path":       path,
		"Driver":     "document",
		"Method":     "delete",
		"Rendezvous": opts}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil

}

func (c Client) CreateCollection(path string) ([]byte, error) {
	values := map[string]interface{}{
		"Path":   path,
		"Method": "create",
		"Driver": "collection"}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(buff))
	}

	return buff, nil

}

func (c Client) DeleteCollection(path string) error {
	values := map[string]interface{}{
		"Path":   path,
		"Driver": "collection",
		"Method": "delete"}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.Url+"/yottadb/",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil

}

func New(url string) (Client, error) {

	client := Client{
		url,
	}

	resp, err := http.Get(url)
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

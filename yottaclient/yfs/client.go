package yfs

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"io"
	"net/http"
)

type Client struct {
	Url string
}

func (c Client) Read(path string) ([]byte, error) {
	values := map[string]string{
		"Path":   path,
		"Method": "read"}
	json_data, err := cbor.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottafs/",
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

func (c Client) Write(path string, data []byte) ([]byte, error) {
	values := map[string]interface{}{
		"Path":   path,
		"Method": "write",
		"Data":   data}
	json_data, err := cbor.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottafs/",
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

func (c Client) Delete(path string, data []byte) ([]byte, error) {
	values := map[string]string{
		"Path":   path,
		"Method": "delete"}
	json_data, err := cbor.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Url+"/yottafs/",
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

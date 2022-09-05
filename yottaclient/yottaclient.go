package yottaclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Client struct {
	Server string
}

func (c Client) Read(record string) ([]byte, error) {

	values := map[string]string{"Record": record}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.Server+"/yottastore/read",
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

func (c Client) Write(record string, data []byte) error {

	values := map[string]interface{}{"Record": record, "Data": data}
	json_data, err := json.Marshal(values)

	if err != nil {
		return err
	}

	resp, err := http.Post(c.Server+"/yottastore/write",
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

	log.Println(string(buff))

	return nil
}

func NewClient(server string) (Client, error) {

	client := Client{
		server,
	}

	return client, nil

}
package yottaclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func YfsRead(path string, node string) ([]byte, error) {
	values := map[string]string{"Path": path}
	json_data, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(node+"/yottafs/read",
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

func YfsWrite(path string, data []byte, node string) error {

	values := map[string]interface{}{"Path": path, "Data": data}
	json_data, err := json.Marshal(values)

	if err != nil {
		return err
	}

	resp, err := http.Post(node+"/yottafs/write",
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

func YfsAppend(path string, data []byte, node string) error {

	values := map[string]interface{}{"Path": path, "Data": data, "Append": true}
	json_data, err := json.Marshal(values)

	if err != nil {
		return err
	}

	resp, err := http.Post(node+"/yottafs/write",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(buff))

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil
}

func YfsDelete(path string, node string) error {

	values := map[string]interface{}{"Path": path}
	json_data, err := json.Marshal(values)

	if err != nil {
		return err
	}

	resp, err := http.Post(node+"/yottafs/delete",
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(buff))

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(buff))
	}

	return nil
}

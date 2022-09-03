package yottaclient

import (
	"bytes"
	"encoding/json"
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

	log.Println(string(buff))

	return nil
}

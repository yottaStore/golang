package read

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Read(record string, node string) (interface{}, error) {

	values := map[string]string{"Path": record}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	url := node + "/read/"
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func ReadOf[T any](record string) (T, error) {

	var result T

	return result, nil

}

package write

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"net/http"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/write"
)

func Write(record string, data interface{}) {

	b, err := msgpack.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	write.Write(record, b)
}

func WriteNew(record string, node string, data []byte) (interface{}, error) {

	values := map[string]string{"Path": record, "Data": string(data)}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	url := node + "/write/"
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	return resp, nil

}

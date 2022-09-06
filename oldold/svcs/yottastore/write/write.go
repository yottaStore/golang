package write

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"io"
	"net/http"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/write"
	"yottaStore/yottaStore-go/src/svcs/yottafs"
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

	/*WriteOptions := yfs.WriteOptions{
		CreateDir: true,
	}*/

	values := yottafs.Request{
		Path:   record,
		Data:   string(data),
		Method: "write",
	}

	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	fmt.Println("Json data: ", string(json_data))
	resp, err := http.Post(node, "application/json",
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

	return string(buff), nil

}

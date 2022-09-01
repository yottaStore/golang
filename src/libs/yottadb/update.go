package yottadb

import (
	"github.com/vmihailenco/msgpack/v5"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/read"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/write"
)

func Update(recordPath string, updates map[string]interface{}) (bool, error) {

	buff, err := read.ReadAll(recordPath)
	if err != nil {
		return false, err
	}

	var record map[string]interface{}
	err = msgpack.Unmarshal(buff, &record)
	if err != nil {
		return false, err
	}

	for k, v := range updates {
		record[k] = v
	}

	b, err := msgpack.Marshal(record)

	err = write.Write(recordPath, b)
	if err != nil {
		return false, err
	}

	return true, nil

}

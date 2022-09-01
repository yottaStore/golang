package yottadb

import (
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/read"
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

	_, err = direct.Write(recordPath, b)
	if err != nil {
		return false, err
	}

	return true, nil

}
